package supportingfunctions

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"regexp"
)

// CheckStringHash определеяет тип хеш суммы по ее длинне
func CheckStringHash(value string) (string, int, error) {
	size := len([]byte(value))

	reg := regexp.MustCompile(`^[a-fA-F0-9]+$`)
	if !reg.MatchString(value) {
		return "", size, errors.New("the value must consist of hexadecimal characters only")
	}

	switch size {
	case md5.Size:
		return "md5", size, nil
	case sha1.Size:
		return "sha1", size, nil
	case sha256.Size:
		return "sha256", size, nil
	case sha512.Size:
		return "sha512", size, nil
	}

	return "other", size, nil
}

/*
func CheckHashSum(hsum string) string {
	switch len(hsum) {
	case 32:
		return "md5"
	case 40:
		return "sha1"
	case 64:
		return "sha256"
	}

	return "other"
}
*/
