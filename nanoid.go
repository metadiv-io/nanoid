package nanoid

import (
	gonanoid "github.com/matoous/go-nanoid"
)

const (
	NUMBERS    = "0123456789"
	LOWERCASE  = "abcdefghijklmnopqrstuvwxyz"
	UPPERCASE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	SYMBOLS    = "_-"
	LOOK_ALIKE = "1lI0Oouv5Ss"

	SAFE_LENGTH   = 21
	SAFE_ALPHABET = "2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ" // numbers + lowercase + uppercase - look-alike
)

// New returns a new nanoid with the given length and alphabet
// Please ensure that the length is long enough to avoid collisions
// You can check the collision probability here: https://zelark.github.io/nano-id-cc/
func New(length int, alphabet string) string {
	nano, err := gonanoid.Generate(alphabet, length)
	if err != nil {
		panic(err)
	}
	return nano
}

// NewSafe returns a new nanoid with numbers, lowercase and uppercase but without symbols and look-alike characters
// The length of the nanoid is 21 characters
// For brute-force cracking, it will take 4 millions of years to have a 1% chance of success for generating 1000 ids per second:
// https://alex7kom.github.io/nano-nanoid-cc/?alphabet=2346789abcdefghijkmnpqrtwxyzABCDEFGHJKLMNPQRTUVWXYZ&size=21&speed=1000&speedUnit=second
func NewSafe() string {
	return New(SAFE_LENGTH, SAFE_ALPHABET)
}
