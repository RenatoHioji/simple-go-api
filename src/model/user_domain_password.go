package model

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"

	"github.com/RenatoHioji/simple-go-api/src/configuration/rest_err"
	"golang.org/x/crypto/argon2"
)

func (user *userDomain) EncryptPassword() *rest_err.RestErr {
	salt := make([]byte, 16)

	if _, err := rand.Read(salt); err != nil {
		return rest_err.NewInternalServerErr("error reading salt from encryption")
	}

	hash := argon2.IDKey([]byte(user.password), salt, 2, 19456, 1, 32)

	user.password = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, 19456, 2, 1,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return nil
}

func (user *userDomain) VerifyPassword(encodedHash string) (bool, error) {
	var version int
	var saltB64, hashB64 string

	_, err := fmt.Sscanf(encodedHash, "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		&version, 19456, 2, 1, &saltB64, &hashB64)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(saltB64)
	if err != nil {
		return false, err
	}
	storedHash, err := base64.RawStdEncoding.DecodeString(hashB64)
	if err != nil {
		return false, err
	}

	computedHash := argon2.IDKey([]byte(user.password), salt, 2, 19456, 1, uint32(len(storedHash)))

	if subtle.ConstantTimeCompare(storedHash, computedHash) == 1 {
		return true, nil
	}
	return false, nil
}
