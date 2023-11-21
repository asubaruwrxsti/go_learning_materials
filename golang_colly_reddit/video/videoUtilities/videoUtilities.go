package videoUtilities

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/golang/freetype/truetype"
)

func IsDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func EmptyDir(path string) error {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		err := os.RemoveAll(filepath.Join(path, file.Name()))
		if err != nil {
			return err
		}
	}

	return nil
}

func LoadFont(path string) (*truetype.Font, error) {
	// Read the font file
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	// Parse the font data
	f, err := truetype.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font data: %v", err)
	}

	return f, nil
}
