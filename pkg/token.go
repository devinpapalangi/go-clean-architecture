package pkg

import (
	"encoding/base64"
	"fmt"
	jwt "github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(payload string) (string, error) {
	config := GetConfig()

	secretKey := config.JWT.Secret

	decodedPrivateKey, err := base64.StdEncoding.DecodeString(secretKey)
	if err != nil {
		return "", err
	}

	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	claims := make(jwt.MapClaims)

	now := time.Now()
	expireIn := time.Duration(config.JWT.Expire)

	var aud, issuer = config.Server.Host, config.Server.Host

	claims["sub"] = payload
	claims["exp"] = now.Add(expireIn).Unix()
	claims["iat"] = now.Unix()
	claims["iss"] = issuer
	claims["aud"] = aud
	claims["nbf"] = now.Unix()

	algorithm := jwt.SigningMethodES256
	token, err := jwt.NewWithClaims(algorithm, claims).SignedString(key)
	if err != nil {
		return "", err
	}

	return token, nil
}

func VerifyToken(token string, publicKey string) (interface{}, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)
	if err != nil {
		return nil, err
	}
	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return nil, err
	}

	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, err
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	return claims["sub"], nil
}
