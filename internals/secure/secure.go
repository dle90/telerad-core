package secure

import (
	"crypto/rand"
	"math/big"

	"telerad-core-module/jwtchecker"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

const (
	passwordLength   = 8
	passwordDigits   = "0123456789"
	passwordUppers   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordLowers   = "abcdefghijklmnopqrstuvwxyz"
	passwordSpecials = "!@#$%^&*()-_=+"
	passwordAllChars = passwordDigits + passwordUppers + passwordLowers + passwordSpecials
)

func CheckAuthorization() fiber.Handler {
	return jwtchecker.New(jwtchecker.Config{})
}

func GetUsernameFromJwt(c *fiber.Ctx) string {
	return extractClaimString(c, JWT_KEY_USER_NAME)
}

func GetUserUuidFromJwt(c *fiber.Ctx) uuid.UUID {
	_uuid, err := uuid.Parse(extractClaimString(c, JWT_KEY_UUID))
	if err != nil {
		return uuid.Nil
	}

	return _uuid
}

func GetCompanyUuidFromJwt(c *fiber.Ctx) uuid.UUID {
	_companyUuid, err := uuid.Parse(extractClaimString(c, JWT_KEY_COMPANY_UUID))
	if err != nil {
		return uuid.Nil
	}

	return _companyUuid
}

func GetFacilityUuidFromJwt(c *fiber.Ctx) uuid.UUID {
	_facilityUuid, err := uuid.Parse(extractClaimString(c, JWT_KEY_FACILITY_UUID))
	if err != nil {
		return uuid.Nil
	}

	return _facilityUuid
}

func GetFacilityUuidFromHeader(c *fiber.Ctx) uuid.UUID {
	_facilityUuid, err := uuid.Parse(c.Get(HEADER_FACILITY_UUID, ""))
	if err != nil {
		return uuid.Nil
	}

	return _facilityUuid
}

func GenerateRandomPassword() string {
	password := make([]byte, passwordLength)

	// Guarantee at least one char from each required category.
	password[0] = pickRandomByte(passwordDigits)
	password[1] = pickRandomByte(passwordUppers)
	password[2] = pickRandomByte(passwordLowers)
	password[3] = pickRandomByte(passwordSpecials)

	// Fill remaining slots from the combined pool.
	for i := 4; i < passwordLength; i++ {
		password[i] = pickRandomByte(passwordAllChars)
	}

	// Shuffle so the required-category chars aren't always at fixed positions.
	shuffleBytes(password)

	return string(password)
}

func VerifyBcryptPassword(hash string, passwordText string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordText)); err == nil {
		return true
	} else {
		return false
	}
}

func GenerateBcryptHash(passwordText string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(passwordText), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	storedHash := string(hash)
	return storedHash, nil
}

func pickRandomByte(charset string) byte {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
	if err != nil {
		// crypto/rand only fails if the OS entropy source is broken — treat as fatal.
		panic(err)
	}
	return charset[n.Int64()]
}

func shuffleBytes(b []byte) {
	for i := len(b) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			panic(err)
		}
		k := j.Int64()
		b[i], b[k] = b[k], b[i]
	}
}

func extractClaimString(c *fiber.Ctx, key string) string {
	value := jwtchecker.ExtractClaimValue(c, key)

	output, ok := value.(string)
	if !ok {
		return ""
	} else {
		return output
	}
}
