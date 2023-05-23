package utils

import (
	"github.com/blessium/porking/model"
	"github.com/lestrrat-go/jwx/v2/jwk"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/lestrrat-go/jwx/v2/jwa"
	"time"
)

func GenerateJWT(u model.User) (string, error) {
    now := time.Now()
	t, err := jwt.NewBuilder().
		Claim(`user_id`, u.ID).
		Claim(`email`, u.Email).
		IssuedAt(now).
        Expiration(now.Add(5 * time.Minute)).
		Build()

	if err != nil {
		return "", err
	}

	rawKey := []byte(`al`)
	jwkKey, err := jwk.FromRaw(rawKey)
	if err != nil {
		return "", err
	}

    token, err := jwt.Sign(t, jwt.WithKey(jwa.HS256, jwkKey))
    if err != nil {
        return "", err
    }

	return string(token), nil
}
