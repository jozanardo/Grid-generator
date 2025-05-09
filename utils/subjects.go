package utils

// Maps all valid subjects (normalized).
var AllowedSubjects = map[string]bool{
	"matematica": true, "fisica": true, "quimica": true, "biologia": true,
	"redacao": true, "portugues": true, "geografia": true,
}

// Checks if a given subject is valid.
func IsValidSubject(subject string) bool {
	return AllowedSubjects[subject]
}
