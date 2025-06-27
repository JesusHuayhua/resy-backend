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

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
	sm "github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"golang.org/x/crypto/pbkdf2"
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
	keyID      string
	passphrase string
	iter       int
}

func New(kmsKeyARN, passphraseName string, iter int) (*EnvelopeCrypto, error) {
	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	smClient := sm.NewFromConfig(cfg)
	out, err := smClient.GetSecretValue(ctx, &sm.GetSecretValueInput{
		SecretId: aws.String(passphraseName),
	})
	if err != nil {
		return nil, err
	}
	pass := aws.ToString(out.SecretString)
	return &EnvelopeCrypto{
		kmsClient:  kms.NewFromConfig(cfg),
		keyID:      kmsKeyARN,
		passphrase: pass,
		iter:       iter,
	}, nil
}

func (e *EnvelopeCrypto) Encrypt(plaintext string) (string, error) {
	salt := make([]byte, 16)
	ctx := context.Background()
	if _, err := io.ReadFull(rand.Reader, salt); err != nil {
		return "", err
	}
	localKey := pbkdf2.Key([]byte(e.passphrase), salt, e.iter, 32, sha512.New)
	genOut, err := e.kmsClient.GenerateDataKey(ctx, &kms.GenerateDataKeyInput{
		KeyId:   &e.keyID,
		KeySpec: types.DataKeySpecAes256,
	})
	if err != nil {
		return "", err
	}

	wrapBlock, _ := aes.NewCipher(genOut.Plaintext)
	wrapGCM, _ := cipher.NewGCM(wrapBlock)
	wrapNonce := make([]byte, wrapGCM.NonceSize())
	io.ReadFull(rand.Reader, wrapNonce)
	wrappedKey := wrapGCM.Seal(nil, wrapNonce, localKey, nil)

	payloadBlock, _ := aes.NewCipher(localKey)
	payloadGCM, _ := cipher.NewGCM(payloadBlock)
	payloadNonce := make([]byte, payloadGCM.NonceSize())
	io.ReadFull(rand.Reader, payloadNonce)
	ct := payloadGCM.Seal(nil, payloadNonce, []byte(plaintext), nil)

	env := Envelope{
		KmsCiphertext: base64.StdEncoding.EncodeToString(genOut.CiphertextBlob),
		WrapNonce:     base64.StdEncoding.EncodeToString(wrapNonce),
		WrappedKey:    base64.StdEncoding.EncodeToString(wrappedKey),
		PayloadNonce:  base64.StdEncoding.EncodeToString(payloadNonce),
		Ciphertext:    base64.StdEncoding.EncodeToString(ct),
		Salt:          base64.StdEncoding.EncodeToString(salt),
		Iter:          e.iter,
	}
	data, err := json.Marshal(env)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func (e *EnvelopeCrypto) Decrypt(envJSON string) (string, error) {
	var env Envelope
	if err := json.Unmarshal([]byte(envJSON), &env); err != nil {
		return "", err
	}

	kmsCt, _ := base64.StdEncoding.DecodeString(env.KmsCiphertext)
	wrapNonce, _ := base64.StdEncoding.DecodeString(env.WrapNonce)
	wrappedKey, _ := base64.StdEncoding.DecodeString(env.WrappedKey)
	payloadNonce, _ := base64.StdEncoding.DecodeString(env.PayloadNonce)
	ct, _ := base64.StdEncoding.DecodeString(env.Ciphertext)
	salt, _ := base64.StdEncoding.DecodeString(env.Salt)
	e.iter = env.Iter
	ctx := context.Background()

	decOut, err := e.kmsClient.Decrypt(ctx, &kms.DecryptInput{
		CiphertextBlob: kmsCt,
	})

	if err != nil {
		return "", err
	}

	wrapBlock, _ := aes.NewCipher(decOut.Plaintext)
	wrapGCM, _ := cipher.NewGCM(wrapBlock)
	localKey, err := wrapGCM.Open(nil, wrapNonce, wrappedKey, nil)
	if err != nil {
		return "", err
	}

	expected := pbkdf2.Key([]byte(e.passphrase), salt, e.iter, 32, sha512.New)
	if !hmac.Equal(localKey, expected) {
		return "", errors.New("passphrase mismatch")
	}

	payloadBlock, _ := aes.NewCipher(localKey)
	payloadGCM, _ := cipher.NewGCM(payloadBlock)
	pt, err := payloadGCM.Open(nil, payloadNonce, ct, nil)
	if err != nil {
		return "", err
	}
	return string(pt), nil
}
