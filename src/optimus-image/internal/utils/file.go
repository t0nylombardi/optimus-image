package utils

import (
	"path/filepath"
)

// IsValidImage checks if the file has a valid image extension
func IsValidImage(filePath string) bool {
	ext := filepath.Ext(filePath)
	validExtensions := []string{".jpg", ".jpeg", ".png", ".webp"}
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
