// Copyright yet.another.artem 2025

package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

var extensions = []string{".go", ".js", ".sh", ".py", ".c", ".cpp", ".ts"}
var licensePrefix = map[string]string{
	".go":  "// ",
	".js":  "// ",
	".ts":  "// ",
	".py":  "# ",
	".sh":  "# ",
	".c":   "// ",
	".cpp": "// ",
}

func hasValidExtension(filename string) bool {
	for _, ext := range extensions {
		if strings.HasSuffix(filename, ext) {
			return true
		}
	}
	return false
}

func getCommentedLicense(ext, license string) string {
	if prefix, ok := licensePrefix[ext]; ok {
		return prefix + license + "\n\n"
	}
	return "// " + license + "\n\n"
}

func fileHasLicense(path string, license string) (bool, error) {
	f, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	if scanner.Scan() {
		return strings.Contains(scanner.Text(), license), nil
	}
	return false, scanner.Err()
}

func injectLicense(path string, license string) error {
	has, err := fileHasLicense(path, license)
	if err != nil || has {
		return err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	ext := filepath.Ext(path)
	newContent := []byte(getCommentedLicense(ext, license))
	newContent = append(newContent, content...)

	if err := os.WriteFile(path, newContent, 0644); err != nil {
		return err
	}

	fmt.Println("✅ Licensed:", path)
	return nil
}

// InjectLicenses this fuction injects licenses into indicated files
func InjectLicenses(root string, license string) error {
	return filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() {
			return err
		}
		if hasValidExtension(path) {
			if err := injectLicense(path, license); err != nil {
				fmt.Println("❌ Error in", path, ":", err)
			}
		}
		return nil
	})
}
