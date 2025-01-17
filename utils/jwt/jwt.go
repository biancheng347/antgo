package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

//New Jwt parameter
type Jwt struct {
	PrivateKey []byte //Private key
	PublicKey  []byte //Public key
	Exp        int64  //Expiration timestamp Default 15 days
}

//New function
func New(PublicKey, PrivateKey []byte, exp ...int64) *Jwt {
	var Exp = time.Now().Add(time.Hour * 168).Unix()
	if len(exp) > 0 {
		Exp = exp[0]
	}

	return &Jwt{
		PublicKey:  PublicKey,
		PrivateKey: PrivateKey,
		Exp:        Exp,
	}
}

//Encrypt json web token encryption<json web token 加密>
func (j *Jwt) Encrypt(row map[string]interface{}) (string, error) {
	Key, _ := jwt.ParseRSAPrivateKeyFromPEM(j.PrivateKey)
	if j.Exp == 0 {
		j.Exp = time.Now().Add(time.Hour * 168).Unix()
	}
	MapClaims := jwt.MapClaims{}
	MapClaims = row
	return jwt.NewWithClaims(jwt.SigningMethodRS256, MapClaims).SignedString(Key)
}

//Decode json web token decryption<json web token解密>
func (j *Jwt) Decode(tokenStr string) (map[string]interface{}, error) {
	result := map[string]interface{}{}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(j.PublicKey)

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, errors.New("token encryption type error")
		}
		return publicKey, nil
	})

	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		result = claims
		return result, nil
	}
	return result, errors.New("token invalid")
}
