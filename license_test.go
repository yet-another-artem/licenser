// Copyright yet.another.artem 2025

package main

import (
    "os"
    "path/filepath"
    "strings"
    "testing"
)

func setupTestFile(t *testing.T, content string) string {
    t.Helper()
    dir := "testdata"
    os.MkdirAll(dir, 0755)

    file := filepath.Join(dir, "test.go")
    err := os.WriteFile(file, []byte(content), 0644)
    if err != nil {
        t.Fatalf("Failed to write test file: %v", err)
    }
    return file
}

func TestInjectLicense(t *testing.T) {
    const license = "Copyright DOLBAEB 2025"

    path := setupTestFile(t, "package main\n\nfunc main() {}")
    err := InjectLicenses("testdata", license)
    if err != nil {
        t.Fatalf("Inject failed: %v", err)
    }

    data, err := os.ReadFile(path)
    if err != nil {
        t.Fatalf("Failed to read file: %v", err)
    }
    if !strings.Contains(string(data), license) {
        t.Errorf("License not injected properly")
    }

    err = InjectLicenses("testdata", license)
    if err != nil {
        t.Fatalf("Inject second time failed: %v", err)
    }
    data2, _ := os.ReadFile(path)
    lines := strings.Split(string(data2), "\n")
    count := 0
    for _, line := range lines {
        if strings.Contains(line, license) {
            count++
        }
    }
    if count > 1 {
        t.Errorf("License line duplicated (%d times)", count)
    }
}