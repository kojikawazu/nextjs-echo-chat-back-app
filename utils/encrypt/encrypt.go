package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/url"
	"nextjs-echo-chat-back-app/utils/logger"
	"os"
	"strconv"
	"strings"
)

// Decrypt は `text` を復号化する
func Decrypt(encryptedText string) (string, error) {
	secretKey := os.Getenv("ALGO_KEY")

	if len(secretKey) != 32 {
		logger.ErrorLog.Printf("Invalid secret key length: expected 32 bytes, got %d", len(secretKey))
		return "", errors.New("invalid secret key length")
	}

	ivLength, err := strconv.Atoi(os.Getenv("IV_LENGTH"))
	if err != nil || ivLength <= 0 {
		logger.ErrorLog.Printf("Invalid IV_LENGTH: using default 12")
		ivLength = 12
	}

	// `encryptedText` を `url.QueryUnescape` でデコード
	decodedText, err := url.QueryUnescape(encryptedText)
	if err != nil {
		logger.ErrorLog.Printf("Failed to URL decode text: %v", err)
		return "", errors.New("invalid URL encoding")
	}

	// `decodedText` を `":"` で分割
	parts := strings.SplitN(decodedText, ":", 3)
	if len(parts) != 3 {
		logger.ErrorLog.Printf("Invalid encrypted data format: %+v", parts)
		return "", errors.New("invalid encrypted data format")
	}

	// `IV` を `hex` でデコード
	iv, err := hex.DecodeString(parts[0])
	if err != nil {
		logger.ErrorLog.Printf("Failed to decode IV: %v", err)
		return "", errors.New("invalid IV format")
	}

	// `IV` の長さが `ivLength` と一致しない場合、エラーを返す
	if len(iv) != ivLength {
		logger.ErrorLog.Printf("Invalid IV length: expected %d, got %d", ivLength, len(iv))
		return "", errors.New("invalid IV length")
	}

	// `ciphertext` を `base64` でデコード
	ciphertext, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		logger.ErrorLog.Printf("Failed to decode ciphertext (base64): %v", err)
		return "", errors.New("invalid ciphertext format")
	}

	// `authTag` を `hex` でデコード
	authTag, err := hex.DecodeString(parts[2])
	if err != nil {
		logger.ErrorLog.Printf("Failed to decode authTag: %v", err)
		return "", errors.New("invalid authTag format")
	}

	// `authTag` の長さが 16 バイトでない場合、エラーを返す
	if len(authTag) != 16 {
		logger.ErrorLog.Printf("Invalid authTag length: expected 16, got %d", len(authTag))
		return "", errors.New("invalid authTag length")
	}

	// `secretKey` を `aes.NewCipher` で暗号化キーとして初期化
	block, err := aes.NewCipher([]byte(secretKey))
	if err != nil {
		logger.ErrorLog.Printf("Failed to create cipher: %v", err)
		return "", err
	}

	// `block` を `cipher.NewGCM` で `GCM` モードの暗号化キーとして初期化
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		logger.ErrorLog.Printf("Failed to create GCM: %v", err)
		return "", err
	}

	// `aesGCM` を `cipher.Open` で復号処理を実行
	plaintext, err := aesGCM.Open(nil, iv, append(ciphertext, authTag...), nil)
	if err != nil {
		logger.ErrorLog.Printf("Failed to decrypt: %v", err)
		return "", errors.New("decryption failed")
	}

	logger.InfoLog.Printf("Decrypted roomId.")
	return string(plaintext), nil
}
