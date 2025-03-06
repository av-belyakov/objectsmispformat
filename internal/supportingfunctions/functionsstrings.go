package supportingfunctions

import (
	"errors"
	"regexp"
)

// CheckStringHash определеяет тип хеш суммы по ее длинне
func CheckStringHash(v string) (string, int, error) {
	size := len(v)

	reg := regexp.MustCompile(`^[a-fA-F0-9]+$`)
	if !reg.MatchString(v) {
		return "", size, errors.New("the value must consist of hexadecimal characters only")
	}

	switch size {
	case 32:
		return "md5", size, nil
	case 40:
		return "sha1", size, nil
	case 64:
		return "sha256", size, nil
	case 128:
		return "sha512", size, nil
	}

	return "other", size, nil
}
