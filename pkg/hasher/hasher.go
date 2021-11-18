package hasher

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/crypto/argon2"
)

type HashConfig struct {
	Time        uint32
	Memory      uint32
	Threads     uint8
	KeyLen      uint32
	prependSalt bool
}

//hashes and salts password using argon2 and returns b64 hash and salt
func HashPassword(password string, cfg *HashConfig) (hash, salt string, err error) {
	if cfg == nil {
		cfg = &HashConfig{
			Time:        1,
			Memory:      64 * 1024,
			Threads:     4,
			KeyLen:      32,
			prependSalt: false,
		}
	}

	rawSalt := make([]byte, 32)
	_, err = rand.Read(rawSalt)
	if err != nil {
		return
	}

	generateHash([]byte(password), rawSalt, cfg)
	salt = base64.RawStdEncoding.EncodeToString(rawSalt)

	return
}

func generateHash(input []byte, salt []byte, cfg *HashConfig) string {
	rawHash := argon2.IDKey(input, salt, cfg.Time, cfg.Memory, cfg.Threads, cfg.KeyLen)

	return base64.RawStdEncoding.EncodeToString(rawHash)

}

func ComparePasswords(hash, password, salt string, cfg *HashConfig) (same bool, err error) {
	if cfg == nil {
		cfg = &HashConfig{
			Time:        1,
			Memory:      64 * 1024,
			Threads:     4,
			KeyLen:      32,
			prependSalt: false,
		}
	}

	rawSalt, err := base64.RawStdEncoding.DecodeString(salt)
	if err != nil {
		return false, err
	}
	hashedInput := generateHash([]byte(password), rawSalt, cfg)

	return hashedInput == hash, err
}
