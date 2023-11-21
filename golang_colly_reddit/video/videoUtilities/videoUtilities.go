package videoUtilities

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"golang.org/x/image/font/sfnt"
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

func LoadFont(path string) (font.Face, error) {
	// Check if the path is a valid file
	if fileInfo, err := os.Stat(path); err != nil {
		return nil, fmt.Errorf("failed to access font file: %v", err)
	} else if fileInfo.IsDir() {
		return nil, errors.New("provided path is a directory, not a font file")
	}

	// Read the font file
	fontBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read font file: %v", err)
	}

	// Parse the font data
	f, err := sfnt.Parse(fontBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font data: %v", err)
	}

	// Create a font.Face from the parsed font
	face, err := opentype.NewFace(f, &opentype.FaceOptions{
		Size:    12, // Adjust the size based on your requirements
		DPI:     72,
		Hinting: font.HintingFull,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create font face: %v", err)
	}

	return face, nil
}
