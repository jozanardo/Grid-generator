package utils

import "regexp"

// Regular expression to validate names (only letters, spaces and periods)
var NameRegex = regexp.MustCompile(`^[a-zA-ZÀ-ú\s.]+$`)
