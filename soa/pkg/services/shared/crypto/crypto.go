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
	"fmt"
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
Explicacion para burros de lo que pasa aca (envelope encryption):
Referencias:
	- https://docs.aws.amazon.com/encryption-sdk/latest/developer-guide/concepts.html
	- https://www.mongodb.com/docs/manual/core/csfle/fundamentals/automatic-encryption/
	- https://cloud.google.com/kms/docs/envelope-encryption
*/

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
			"ASIAW3MEC65RQDVR76XU",
			"sIkpmYqgIErRXnhMYTI6JrGfbXRsTTlJDQejuoOe",
			"IQoJb3JpZ2luX2VjEBQaCXVzLXdlc3QtMiJHMEUCIQCBZ+zrnksUubDQuuc7WFeEIJ3Nd8jqxEUG1x9LyOCRLgIgYzNy9sF6C35aLoo1jlI2FjC91F9XfarLQhSmGwq4dZsqrwIIHRAAGgw0NzExMTI4MDgyOTEiDAU9CXNt+Tp0fQv6hiqMAnP67gAkpArtMigANzCZQRSgaLe5wyxfODw2la7ta9iGHJxnSt/rxFlTI/Mu9WcJPLZIvUgB2IRPSZt0NdOqcyv3/h+80lHafubdHIYvPNWDl397MxF4smmGFx4rDtE1fGBl6JugZBg1qwmloNOoI6elvtYrofjlvu5Q/7XRlFS67mSA3L3+1JvUYTGbwE5eF6CoScmAE6xFV0KBhgAb1h2bc7u+E1BI0bVIah/cyTWmZrW7Kr17vT5tLNuFF5UkT6k/97x0PFv1cEK88Vmrj3UD/WbW3F/pOPbqVEd89yKqYMPc17Ponijk4PQS9helUeF6MGgs0hnI0gtPLz0ciAalAQGK49NbQYdqVCYw9MGbwwY6nQEX5JHaK9DZnxhPZFUhDwchRTLmCzH/2QK/B9lqDcjkfS65yvoZKrbBxQpHxHl+4AhYBgosNAla/u6Dphh0BwZ9n++RvntBSH6HvN4FiqJmkLrtGyWPE9SLALK2CLImbkXAAV9ObBKJA+Dh9qCNTIkBB/R8ko10CYW3T9XendBzaIsAgMFGHFOh+9xwKndKUhTIi8Qp21RIt6yStEjt"),
		),
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
	ciphertext := base64.StdEncoding.EncodeToString(out)
	return string(ciphertext), err
}

func (e *EnvelopeCrypto) Decrypt(enc string) (string, error) {
	var env Envelope
	envJSON, err := base64.StdEncoding.DecodeString(enc)
	if err != nil {
		return "", fmt.Errorf("error al desencodear b64 config, error: %v", err)
	}
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
