package models

// JWK（JSON Web Key）の構造体
type JWK struct {
	Kid string `json:"kid"`
	Kty string `json:"kty"`
	Alg string `json:"alg"`
	N   string `json:"n"`
	E   string `json:"e"`
}

// JWKsレスポンスの構造体
type JWKs struct {
	Keys []JWK `json:"keys"`
}
