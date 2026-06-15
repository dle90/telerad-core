package jwtchecker

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var _config Config
var publicKey *rsa.PublicKey
var privateKey *rsa.PrivateKey

const AUTHORIZATION string = "Authorization"

/*
REQUIRED(Any middleware must have this)

For every middleware we need a config.
In config we also need to define a function which allows us to skip the middleware if return true.
By convention it should be named as "Filter" but any other name will work too.
*/
type Config struct {
	// function to run when there is error decoding jwt
	Unauthorized fiber.Handler
	// function to decode our jwt token
	Decode func(c *fiber.Ctx) (*jwt.MapClaims, error)
	// set jwt expiry in seconds
	Expiry int64
}

/*
Middleware specific

Our middleware's config default values if not passed
*/
var ConfigDefault = Config{
	Decode:       nil,
	Unauthorized: nil,
	Expiry:       60,
}

func InitPrivateKey(keyPath string) {
	key, err := GetPrivateKeyFromFile(keyPath)
	if err != nil {
		panic("This panic is caught by settings jwt-private-key-path fail")
	}

	_privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(key))
	if err != nil {
		panic("This panic is caught by parsing jwt-private-key fail")
	}

	privateKey = _privateKey
}

func InitPublicKey(keyPath string) {
	key, err := GetPublicKeyFromFile(keyPath)
	if err != nil {
		panic("This panic is caught by settings jwt-public-key-path fail")
	}

	_publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(key))
	if err != nil {
		panic("This panic is caught by parsing jwt-public-key fail")
	}

	publicKey = _publicKey
}

// InitPrivateKeyFromContent nạp private key trực tiếp từ nội dung PEM (vd lấy từ
// biến môi trường JWT_PRIVATE_KEY) thay vì đọc file.
func InitPrivateKeyFromContent(pemContent string) {
	_privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(pemContent))
	if err != nil {
		panic("This panic is caught by parsing jwt-private-key content fail")
	}

	privateKey = _privateKey
}

// InitPublicKeyFromContent nạp public key trực tiếp từ nội dung PEM (env JWT_PUBLIC_KEY).
func InitPublicKeyFromContent(pemContent string) {
	_publicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(pemContent))
	if err != nil {
		panic("This panic is caught by parsing jwt-public-key content fail")
	}

	publicKey = _publicKey
}

// validate jwt token
func ValidateJwtToken(tokenString string) (bool, error) {
	if tokenString == "" || len(tokenString) < 8 {
		return false, errors.New("AUTHORIZATION HEADER IS REQUIRED")
	}
	// we parse our jwt token and check for validity against our secret
	token, err := jwt.Parse(
		tokenString[7:],
		func(token *jwt.Token) (interface{}, error) {
			// IMPORTANT: Validating the algorithm per https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
			if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
				return nil, fmt.Errorf(
					"EXPECTED TOKEN ALGORITHM '%v' BUT GOT '%v'",
					jwt.SigningMethodRS256.Name,
					token.Header)
			}
			untypedKeyId, found := token.Header["alg"]
			if !found {
				return nil, fmt.Errorf("NO KEY ID KEY '%v' FOUND IN TOKEN HEADER", "alg")
			}

			_, ok := untypedKeyId.(string)
			if !ok {
				return nil, fmt.Errorf("FOUND KEY ID, BUT VALUE WAS NOT A STRING")
			}

			return publicKey, nil
		},
	)

	if err != nil {
		return false, errors.New("ERROR PARSING TOKEN")
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	// if !(ok && token.Valid) {
	if !(ok) {
		return false, errors.New("INVALID TOKEN")
	}

	if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
		return false, errors.New("jwt is expired")
	}

	return true, nil
}

/*
Middleware specific
Function for generating default config
*/
func configDefault(config ...Config) Config {
	// Return default config if nothing provided
	if len(config) < 1 {
		return ConfigDefault
	}

	// Override default config
	cfg := config[0]

	// Set default expiry if not passed
	if cfg.Expiry == 0 {
		cfg.Expiry = ConfigDefault.Expiry
	}

	// this is the main jwt decode function of our middleware
	if cfg.Decode == nil {
		// Set default Decode function if not passed
		cfg.Decode = func(c *fiber.Ctx) (*jwt.MapClaims, error) {
			authHeader := c.Get(AUTHORIZATION)
			if authHeader == "" || len(authHeader) < 8 {
				return nil, errors.New("AUTHORIZATION HEADER IS REQUIRED")
			}

			//tokenString := authHeader[7:]
			/*claims := serviceCaches.GetJwtCache(tokenString)
			getClaimsFromCache := false*/
			var claims jwt.MapClaims

			//if claims == nil {
			if true {
				// we parse our jwt token and check for validity against our secret
				token, err := jwt.Parse(
					authHeader[7:],
					func(token *jwt.Token) (interface{}, error) {
						// IMPORTANT: Validating the algorithm per https://godoc.org/github.com/dgrijalva/jwt-go#example-Parse--Hmac
						if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
							return nil, fmt.Errorf(
								"EXPECTED TOKEN ALGORITHM '%v' BUT GOT '%v'",
								jwt.SigningMethodRS256.Name,
								token.Header)
						}
						untypedKeyId, found := token.Header["alg"]
						if !found {
							return nil, fmt.Errorf("NO KEY ID KEY '%v' FOUND IN TOKEN HEADER", "alg")
						}

						_, ok := untypedKeyId.(string)
						if !ok {
							return nil, fmt.Errorf("FOUND KEY ID, BUT VALUE WAS NOT A STRING")
						}

						return publicKey, nil
					},
				)

				if err != nil {
					return nil, errors.New("ERROR PARSING TOKEN")
				}

				claimss, ok := token.Claims.(jwt.MapClaims)

				// if !(ok && token.Valid) {
				if !(ok) {
					return nil, errors.New("INVALID TOKEN")
				}

				claims = claimss
				for key, value := range claims {
					c.Locals(key, value)
				}
			} else {
				//getClaimsFromCache = true
			}

			/*if expiresAt, ok := claims["exp"]; ok && int64(expiresAt.(float64)) < time.Now().UTC().Unix() {
				if getClaimsFromCache {
					serviceCaches.DeleteJwtCache(tokenString)
				}
				return nil, errors.New("jwt is expired")
			}

			if !getClaimsFromCache {
				serviceCaches.SetJwtCache(tokenString, claims)
			}*/

			return &claims, nil
		}
	}

	// Set default Unauthorized if not passed
	if cfg.Unauthorized == nil {
		cfg.Unauthorized = func(c *fiber.Ctx) error {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
	}

	return cfg
}

/*
Middleware specific
Function to generate a jwt token
*/
func Encode(claims *jwt.MapClaims, expiryAfter int64) (string, error) {
	// setting default expiryAfter
	if expiryAfter == 0 {
		expiryAfter = ConfigDefault.Expiry
	}

	// or you can use time.Now().Add(time.Second * time.Duration(expiryAfter)).UTC().Unix()
	(*claims)["exp"] = time.Now().UTC().Unix() + expiryAfter

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// our signed jwt token string
	signedToken, err := token.SignedString(privateKey)

	if err != nil {
		return "", errors.New("ERROR CREATING A TOKEN")
	}

	return signedToken, nil
}

/*
REQUIRED(Any middleware must have this)
Our main middleware function used to initialize our middleware.
By convention we name it "New" but any other name will work too.
*/
func New(config Config) fiber.Handler {
	// For setting default config
	_config = configDefault(config)

	return func(c *fiber.Ctx) error {
		claims, err := _config.Decode(c)
		if err == nil {
			c.Locals("jwtClaims", *claims)
			return c.Next()
		}

		return _config.Unauthorized(c)
	}
}

func GetPrivateKeyFromFile(filePath string) (string, error) {
	// we need to pass in a secret otherwise default secret is used
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// return jwt.ParseRSAPublicKeyFromPEM(keyData)
	return string(keyData), nil
}

func GetPublicKeyFromFile(filePath string) (string, error) {
	// we need to pass in a secret otherwise default secret is used
	keyData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	// return jwt.ParseRSAPublicKeyFromPEM(keyData)
	return string(keyData), nil
}

func ExtractClaimValue(c *fiber.Ctx, key string) any {
	/*claims, err := extractClaimsUnverified(c)
	if err != nil {
		return ""
	}

	return claims[key].(string)*/

	return c.Locals(key)
}
