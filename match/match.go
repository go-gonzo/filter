package match

import "github.com/bmatcuk/doublestar"

func Good(patterns ...string) error {
	for _, p := range patterns {
		if _, err := doublestar.Match(p, ""); err != nil {
			return err
		}
	}
	return nil
}

// Any returns whatever the path matches at least one of the patterns.
// pattern errors are ignored.
func Any(path string, patterns ...string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, path); match {
			return true
		}
	}
	return false
}

// All returns whatever the path matches all of the patterns.
// pattern errors are ignored.
func All(path string, patterns ...string) bool {
	for _, pattern := range patterns {
		if match, _ := doublestar.Match(pattern, path); !match {
			return false
		}
	}
	return true
}
