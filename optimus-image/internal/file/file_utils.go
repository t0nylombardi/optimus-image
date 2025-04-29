package file

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
)

// FileUtils defines the contract for file-related operations.
type FileUtils interface {
	GetFilePath() (string, error)
	GetDirectoryPath() (string, error) // ✅ Now implemented
	GetFilesInDirectory(dirPath string) ([]string, error)
}

// FileUtilsImpl is the default implementation of FileUtils.
type FileUtilsImpl struct{}

// NewFileUtils creates and returns a new FileUtils instance.
func NewFileUtils() FileUtils {
	return &FileUtilsImpl{}
}

// GetFilePath prompts the user to enter a file path.
func (f *FileUtilsImpl) GetFilePath() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter the image file path",
	}

	filePath, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get file path: %w", err)
	}

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return "", fmt.Errorf("file does not exist: %s", filePath)
	}

	return filePath, nil
}

// GetDirectoryPath prompts the user to enter a directory path.
func (f *FileUtilsImpl) GetDirectoryPath() (string, error) {
	prompt := promptui.Prompt{
		Label: "Enter the directory path",
	}

	dirPath, err := prompt.Run()
	if err != nil {
		return "", fmt.Errorf("failed to get directory path: %w", err)
	}

	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return "", fmt.Errorf("directory does not exist: %s", dirPath)
	}

	return dirPath, nil
}

// GetFilesInDirectory returns all image files in the specified directory.
func (f *FileUtilsImpl) GetFilesInDirectory(dirPath string) ([]string, error) {
	var files []string

	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
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

// IsValidImage checks if the file has a valid image extension.
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
