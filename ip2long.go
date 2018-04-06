package ip2long

import (
	"errors"
)

const (
	digitMask int = 1
	dotMask   int = 2
	spaceMask int = 4
)

// Errors
var (
	// ErrInvalidIPv4Address means that the string is not a valid ipv4 address
	ErrInvalidIPv4Address = errors.New("invalid ipv4 address")

	// ErrMalformedIPv4Address means that the string contains invalid characters other than digits, dots or spaces
	ErrMalformedIPv4Address = errors.New("malformed ipv4 address")

	// ErrOverflowedIPv4Segment means that the string contains segments bigger than 255
	ErrOverflowedIPv4Segment = errors.New("overflowed ipv4 segment")
)

var typeDescMap = map[int]string{
	digitMask: "digit",
	dotMask:   "dot",
	spaceMask: "space",
}

// IPv42long exports function that convert ipv4 string to an integer
func IPv42long(ipStr string) (int, error) {
	seg := -1
	segs := 0
	ipInt := 0

	// Allowed next character
	bitmask := digitMask | spaceMask

	// Iterate the input string
	for _, char := range ipStr {
		// Check if the current character matches our expection,
		// and also get the type (digit, dot, or space)
		category, err := validateChar(char, bitmask, segs)
		if err != nil {
			return 0, err
		}

		switch category {
		case digitMask:
			if seg == -1 {
				seg = 0
			}
			seg = seg*10 + int(char-'0')

			if seg > 255 {
				return 0, ErrOverflowedIPv4Segment
			}
			bitmask = digitMask | spaceMask | dotMask
		case dotMask:
			segs++
			ipInt = (seg & 255) + (ipInt << 8)
			seg = -1
			bitmask = digitMask | spaceMask
		case spaceMask:
			if seg == -1 {
				bitmask = digitMask | spaceMask
			} else {
				bitmask = dotMask | spaceMask
			}
			continue
		}
	}
	ipInt = (seg & 255) + (ipInt << 8)

	return ipInt, nil
}

func validateChar(char rune, mask int, segs int) (int, error) {
	if segs > 3 {
		return 0, ErrInvalidIPv4Address
	}

	var bit int
	switch true {
	case (char >= '0' && char <= '9'):
		bit = digitMask
	case (char == '.'):
		bit = dotMask
	case (char == ' '):
		bit = spaceMask
	default:
		return 0, ErrMalformedIPv4Address
	}

	if mask&bit == 0 {
		return 0, ErrInvalidIPv4Address
	}

	return bit, nil
}
