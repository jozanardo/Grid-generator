package utils

import (
	"strings"
	"unicode"
)

// Function to remove accents and convert to lowercase
func NormalizeText(text string) string {
	var result strings.Builder
	for _, char := range text {
		switch char {
		case 'á', 'à', 'â', 'ã', 'ä', 'Á', 'À', 'Â', 'Ã', 'Ä':
			result.WriteRune('a')
		case 'é', 'è', 'ê', 'ë', 'É', 'È', 'Ê', 'Ë':
			result.WriteRune('e')
		case 'í', 'ì', 'î', 'ï', 'Í', 'Ì', 'Î', 'Ï':
			result.WriteRune('i')
		case 'ó', 'ò', 'ô', 'õ', 'ö', 'Ó', 'Ò', 'Ô', 'Õ', 'Ö':
			result.WriteRune('o')
		case 'ú', 'ù', 'û', 'ü', 'Ú', 'Ù', 'Û', 'Ü':
			result.WriteRune('u')
		case 'ç', 'Ç':
			result.WriteRune('c')
		default:
			result.WriteRune(unicode.ToLower(char)) // Convert to lowercase
		}
	}
	return result.String()
}
