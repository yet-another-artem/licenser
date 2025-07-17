package main

import (
    "fmt"
    "os"
)

func main() {
    dir := "."
    if len(os.Args) > 1 {
        dir = os.Args[1]
    }
    license := "Copyright DOLBAEB 2025"

    fmt.Println("💉 Injecting license into files in:", dir)
    if err := InjectLicenses(dir, license); err != nil {
        fmt.Println("❌ Error:", err)
    }
}