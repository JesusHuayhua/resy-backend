package crypto

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/json"
	"errors"
	"io"
	"log"

	"golang.org/x/crypto/pbkdf2"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type Envelope struct {
	KmsCiphertext string `json:"kmsCiphertext"`
	WrapNonce     string `json:"wrapNonce"`
	WrappedKey    string `json:"wrappedKey"`
	PayloadNonce  string `json:"payloadNonce"`
	Ciphertext    string `json:"ciphertext"`
	Salt          string `json:"salt"`
	Iter          int    `json:"iter"`
}

type EnvelopeCrypto struct {
	kmsClient  *kms.Client
	keyID      *string
	passphrase string
	iter       int
}

/*
Cuando se esta ejecutando en maquina local, se requiere dumpear las credenciales, esto se puede hacer con.
aws configure get aws_secret_access_key
aws configure get aws_session_token
aws configure get aws_access_key_id
T.B.D:
*/
func New(keyAlias string, region string, passphraseName string, iter int) (*EnvelopeCrypto, error) {
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
			"ASIAW3MEC65RY7HDSFZQ",
			"d0K3TDqeQMZImW1Ps/GsWFSaewR+WHPJCh6qd55w",
			"IQoJb3JpZ2luX2VjEAEaCXVzLXdlc3QtMiJGMEQCIENCmSZuKqY3e+eaI6QXObRHgF20hYKItzx/Q3pzsqpcAiAGy4bLhbe5C5NKrxx16VJdhIKUhydm+K87m3qToYOmXiq4Agj6//////////8BEAAaDDQ3MTExMjgwODI5MSIMLoF3zprfFsdiyjriKowCRciP7qkjIW2NuzF3mTdlO6mih8pLaUKLMwdSd6XaV8yQrj4Rj09XoVtSyjJkH1j7WkWhR84GBcHhQfOPVruWrkRFgDuymoNHaxk0MCb2oX+/hDpb4zsuiNy6E0f9q4jrVbHu0jEkwVfqYs82JvWwiVErCVlfwNOzMIM/2VDvo3ZpNkl7K6LEQ4nZ3E6pxwy/sy8Rn9/mYjh/vS5fnays2JntwhMKVbDNRi+6MScgSqBGagf0+eKM7sc3xh3FPQBM8ao1S16yXWrWFqwTbGEjs641IQuF34mCRu7u+wStsew46H9uOzl0ted//uTGOO4J9OfPhh6TLpkqie4OiqRAyqk5Bg01yrHgM2YafDCHsJfDBjqeAXdinOoVJjeosVsJSOjMeBNePex4z0qWAdf8UOWrvOtZEitsyT8DbTMMuPhfy1yG5YTIPgDRXoQGsff6jiyc7iUurVRQt0AZvGGVOWHXrJUZEXFy+vlQ8VHJE7EWubJhOsii6nUCzl/J0RatzeHgxxrthbOKHYSh2NEONWpIQpn/XjKP2ZiZ7v2kT0EzX9rzcj+haywxYgRu4W8nMPoN")),
	)
	if err != nil {
		return nil, err
	}
	svc := sm.NewFromConfig(cfg)
	input := &sm.GetSecretValueInput{
		SecretId:     aws.String(passphraseName),
		VersionStage: aws.String("AWSCURRENT"),
	}
	out, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		log.Fatal(err.Error())
	}
	aws_json := aws.ToString(out.SecretString)
	var m map[string]interface{}
	if err := json.Unmarshal([]byte(aws_json), &m); err != nil {
		log.Fatalf("invalid JSON from Secrets Manager: %v", err)
	}
	pass, ok := m["passphrase"].(string)
	if !ok {
		log.Fatal("secret JSON does not contain a string passphrase")
	}
	return &EnvelopeCrypto{
		kmsClient:  kms.NewFromConfig(cfg),
		keyID:      aws.String(keyAlias),
		passphrase: pass,
		iter:       iter,
	}, nil
}

func (e *EnvelopeCrypto) Encrypt(plaintext string) (string, error) {
	ctx := context.Background()
	genOut, err := e.kmsClient.GenerateDataKey(ctx, &kms.GenerateDataKeyInput{
		KeyId:   e.keyID,
		KeySpec: types.DataKeySpecAes256,
	})
	if err != nil {
		return "", err
	}
	dek := genOut.Plaintext
	salt := make([]byte, 16)
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	kek := pbkdf2.Key([]byte(e.passphrase), salt, e.iter, 32, sha512.New)
	wrapBlock, _ := aes.NewCipher(kek)
	wrapGCM, _ := cipher.NewGCM(wrapBlock)
	wrapNonce := make([]byte, wrapGCM.NonceSize())
	io.ReadFull(rand.Reader, wrapNonce)
	wrappedKey := wrapGCM.Seal(nil, wrapNonce, dek, nil)
	payloadBlock, _ := aes.NewCipher(dek)
	payloadGCM, _ := cipher.NewGCM(payloadBlock)
	payloadNonce := make([]byte, payloadGCM.NonceSize())
	io.ReadFull(rand.Reader, payloadNonce)
	ct := payloadGCM.Seal(nil, payloadNonce, []byte(plaintext), nil)
	for i := range dek {
		dek[i] = 0
	}
	for i := range kek {
		kek[i] = 0
	}
	env := Envelope{
		KmsCiphertext: base64.StdEncoding.EncodeToString(genOut.CiphertextBlob),
		Salt:          base64.StdEncoding.EncodeToString(salt),
		Iter:          e.iter,
		WrapNonce:     base64.StdEncoding.EncodeToString(wrapNonce),
		WrappedKey:    base64.StdEncoding.EncodeToString(wrappedKey),
		PayloadNonce:  base64.StdEncoding.EncodeToString(payloadNonce),
		Ciphertext:    base64.StdEncoding.EncodeToString(ct),
	}
	out, err := json.Marshal(env)
	return string(out), err
}

func (e *EnvelopeCrypto) Decrypt(envJSON string) (string, error) {
	var env Envelope
	if err := json.Unmarshal([]byte(envJSON), &env); err != nil {
		return "", err
	}
	kmsCt, _ := base64.StdEncoding.DecodeString(env.KmsCiphertext)
	salt, _ := base64.StdEncoding.DecodeString(env.Salt)
	wrapNonce, _ := base64.StdEncoding.DecodeString(env.WrapNonce)
	wrappedKey, _ := base64.StdEncoding.DecodeString(env.WrappedKey)
	payloadNonce, _ := base64.StdEncoding.DecodeString(env.PayloadNonce)
	ct, _ := base64.StdEncoding.DecodeString(env.Ciphertext)
	kek := pbkdf2.Key([]byte(e.passphrase), salt, env.Iter, 32, sha512.New)

	wrapBlock, _ := aes.NewCipher(kek)
	wrapGCM, _ := cipher.NewGCM(wrapBlock)
	localDEK, err := wrapGCM.Open(nil, wrapNonce, wrappedKey, nil)
	if err != nil {
		return "", errors.New("invalid passphrase or corrupted data")
	}

	decOut, err := e.kmsClient.Decrypt(context.Background(), &kms.DecryptInput{
		CiphertextBlob: kmsCt,
	})
	if err != nil {
		return "", err
	}

	if !hmac.Equal(decOut.Plaintext, localDEK) {
		return "", errors.New("passphrase mismatch")
	}

	block, _ := aes.NewCipher(localDEK)
	gcm, _ := cipher.NewGCM(block)
	pt, err := gcm.Open(nil, payloadNonce, ct, nil)

	for i := range localDEK {
		localDEK[i] = 0
	}
	for i := range kek {
		kek[i] = 0
	}

	if err != nil {
		return "", err
	}
	return string(pt), nil
}
