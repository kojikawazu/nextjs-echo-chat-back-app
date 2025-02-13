package utils

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"nextjs-echo-chat-back-app/models"
	"nextjs-echo-chat-back-app/utils/logger"
	"os"
	"sync"
)

// ClerkのJWKsエンドポイント
var clerkJwksURL string

// 公開鍵のキャッシュ用
var (
	clerkPublicKeys map[string]*rsa.PublicKey
	mu              sync.Mutex
)

func init() {
	clerkJwksURL = os.Getenv("CLERK_JWT_ENDPOINT")
	clerkPublicKeys = make(map[string]*rsa.PublicKey)
}

// Clerkの公開鍵を取得
func GetClerkPublicKey(kid string) (*rsa.PublicKey, error) {
	mu.Lock()
	defer mu.Unlock()

	// キャッシュがある場合、それを使用
	if key, exists := clerkPublicKeys[kid]; exists {
		return key, nil
	}

	// JWKsを取得
	resp, err := http.Get(clerkJwksURL)
	if err != nil {
		logger.ErrorLog.Printf("failed to get Clerk JWKs: %v", err)
		return nil, fmt.Errorf("failed to get Clerk JWKs: %v", err)
	}
	defer resp.Body.Close()

	var jwks models.JWKs
	if err := json.NewDecoder(resp.Body).Decode(&jwks); err != nil {
		logger.ErrorLog.Printf("failed to parse Clerk JWKs: %v", err)
		return nil, fmt.Errorf("failed to parse Clerk JWKs: %v", err)
	}

	// `kid` に一致する鍵を探す
	for _, key := range jwks.Keys {
		if key.Kid == kid {
			// RSA公開鍵に変換
			pubKey, err := ConvertJWKToRSAPublicKey(key)
			if err != nil {
				logger.ErrorLog.Printf("failed to convert JWK to RSA public key: %v", err)
				return nil, fmt.Errorf("failed to convert JWK to RSA public key: %v", err)
			}

			// キャッシュに保存
			clerkPublicKeys[kid] = pubKey

			return pubKey, nil
		}
	}

	return nil, errors.New("public key not found")
}

// JWKをRSA公開鍵に変換
func ConvertJWKToRSAPublicKey(jwk models.JWK) (*rsa.PublicKey, error) {
	// `n` と `e` をBase64URLデコード
	nBytes, err := base64.RawURLEncoding.DecodeString(jwk.N)
	if err != nil {
		logger.ErrorLog.Printf("failed to decode n: %v", err)
		return nil, fmt.Errorf("failed to decode n: %v", err)
	}
	eBytes, err := base64.RawURLEncoding.DecodeString(jwk.E)
	if err != nil {
		logger.ErrorLog.Printf("failed to decode e: %v", err)
		return nil, fmt.Errorf("failed to decode e: %v", err)
	}

	// `e` はバイト配列なので数値に変換
	e := int(binary.BigEndian.Uint64(append(make([]byte, 8-len(eBytes)), eBytes...)))

	// `n` を big.Int に変換
	n := new(big.Int).SetBytes(nBytes)

	// RSA公開鍵を生成
	pubKey := &rsa.PublicKey{
		N: n,
		E: e,
	}

	// 公開鍵の正当性を検証
	_, err = x509.MarshalPKIXPublicKey(pubKey)
	if err != nil {
		logger.ErrorLog.Printf("invalid RSA public key: %v", err)
		return nil, fmt.Errorf("invalid RSA public key: %v", err)
	}

	return pubKey, nil
}
