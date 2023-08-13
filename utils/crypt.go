package utils

import (
	"encoding/base64"
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/scrypt"
)

// SaltRange is the range of salt
var SaltRange []byte = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxy@#$")

func init() {
	// init rand, seed
	rand.NewSource(time.Now().UnixNano()).Seed(time.Now().UnixNano())
}

// CryptPassword is a func to crypt password
func CryptPassword(password, salt []byte) (hash string, err error) {
	hashByte, err := scrypt.Key(password, salt, 32768, 8, 1, 32)
	if err != nil {
		return "", err
	}
	hash = base64.StdEncoding.EncodeToString(hashByte)
	return
}

// ParsePassword is a func to parse crypt password
func ParsePassword(hash string, password, salt []byte) (err error) {
	computedHash, err := CryptPassword(password, salt)
	if err != nil {
		return
	}
	if hash != computedHash {
		return fmt.Errorf("ParsePassword: password fail")
	}
	return
}

// RandSalt is a func to generate rand salt
func RandSalt(saltRange []byte, large int) (salt []byte, err error) {
	for i := 0; i <= large; i++ {
		salt = append(salt, saltRange[rand.Intn(len(saltRange))])
	}
	if len(saltRange) < 6 || large < 6 {
		return nil, fmt.Errorf("RandSalt: len of saltRang and largen must greater than 6")
	}
	return
}

// RandHashedPassword is a func to generate rand hashed password
func RandHashedPassword(hashed string, randSalt []byte) (password string) {
	return fmt.Sprintf("%v:%v", hashed, base64.StdEncoding.EncodeToString(randSalt))
}

// ParseRandPassword is a func to parse crypt password with rand salt
func ParseRandPassword(hash string, password []byte) (err error) {
	salt, err := base64.StdEncoding.DecodeString(strings.Split(hash, ":")[1])
	if err != nil {
		return fmt.Errorf("ParsePassword: hash fail")
	}
	computedHash, err := CryptPassword(password, salt)
	if err != nil {
		return
	}
	if strings.Split(hash, ":")[0] != computedHash {
		return fmt.Errorf("ParsePassword: password fail")
	}
	return
}

// GenerateRandID is a func to generate 10 bit rand id
func GenerateRandID(digits int) string {
	rang := math.Pow10(digits)
	return fmt.Sprintf("%0"+strconv.Itoa(digits)+"d", rand.Int63n(int64(rang)))
}
