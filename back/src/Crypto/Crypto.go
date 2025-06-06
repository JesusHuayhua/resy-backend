package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
)

type Config struct {
	EncryptionKey string `json:"encryption_key"`
}

// Versión determinista (mismo resultado para mismos inputs)
func getDeterministicNonce(key []byte) []byte {
	hash := sha256.Sum256(key)
	return hash[:12] // 12 bytes para GCM nonce
}

func loadKey() ([]byte, error) {
	configPath := filepath.Join("config", "crip.json")
	file, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var config Config
	if err := json.Unmarshal(file, &config); err != nil {
		return nil, err
	}

	if config.EncryptionKey == "" {
		return nil, errors.New("la clave de encriptación no puede estar vacía")
	}

	key := sha256.Sum256([]byte(config.EncryptionKey))
	return key[:], nil
}

func Encrypt(plaintext string) (string, error) {
	plaintextBytes := []byte(plaintext)
	key, err := loadKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	// Nonce determinista derivado de la clave
	nonce := getDeterministicNonce(key)

	ciphertext := gcm.Seal(nil, nonce, plaintextBytes, nil)
	return base64.URLEncoding.EncodeToString(append(nonce, ciphertext...)), nil
}

func Decrypt(ciphertext string) (string, error) {
	ciphertextBytes, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	key, err := loadKey()
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertextBytes) < nonceSize {
		return "", errors.New("texto cifrado demasiado corto")
	}

	nonce, ciphertextBytes := ciphertextBytes[:nonceSize], ciphertextBytes[nonceSize:]
	plaintextBytes, err := gcm.Open(nil, nonce, ciphertextBytes, nil)
	if err != nil {
		return "", err
	}

	return string(plaintextBytes), nil
}