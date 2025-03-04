package utils

import (
	"os"
	"path/filepath"
)

// GetFilesInDirectory returns a list of image files in the directory
// GetFilesInDirectory returns all image files in the specified directory.
func GetFilesInDirectory(dirPath string) ([]string, error) {
	var files []string
	// Walk through the directory
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		// If an error occurs while walking, return it
		if err != nil {
			return err
		}

		// Check if the file is an image based on its extension
		if !info.IsDir() && IsValidImage(info.Name()) {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}

// IsValidImage checks if the file has a valid image extension
func IsValidImage(filePath string) bool {
	ext := filepath.Ext(filePath)
	validExtensions := []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".tiff", ".webp"}
	for _, validExt := range validExtensions {
		if ext == validExt {
			return true
		}
	}
	return false
}
