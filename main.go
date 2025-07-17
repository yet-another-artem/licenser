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
    license := "Copyright ya.eblan.pomogite.mne 2025"

    fmt.Println("ватермарка отправлена в файл:", dir)
    if err := InjectLicenses(dir, license); err != nil {
        fmt.Println("А ПОШЕЛ ТЫ НАХУЙ :)):", err)
    }
}
