package validators

import "unicode"

func Password(pass string) bool {
	var (
		upp, low, num bool
		tot           uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || tot < 8 {
		return false
	}

	return true
}
