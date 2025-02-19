package utils

// AllowedSubjects stores all valid subjects (normalized).
var AllowedSubjects = map[string]bool{
	"matematica": true, "fisica": true, "quimica": true, "biologia": true,
	"redacao": true, "portugues": true, "geografia": true,
}

// IsValidSubject checks if a given subject is valid.
func IsValidSubject(subject string) bool {
	return AllowedSubjects[subject]
}
