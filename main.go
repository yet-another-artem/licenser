// Copyright yet.another.artem 2025

// Copyright . 2025

// Copyright yet.another.artem 2025

package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	dir := os.Args[1]
	copyrightHolder := os.Args[2]
	currentYear := time.Now().Year()
	license := "Copyright " + copyrightHolder + " " + strconv.Itoa(currentYear)
	fmt.Println("Copyright will be injected in the following file or directory:", dir)
	if err := InjectLicenses(dir, license); err != nil {
		fmt.Println("Oh, there is an error:", err)
	}
}
