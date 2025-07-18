package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

func main() {
	dir := os.Args[1]
	copyright_holder := os.Args[2]
	current_year := time.Now().Year()
	license := "Copyright " + copyright_holder + " " + strconv.Itoa(current_year)
	fmt.Println("Copyright will be injected in the following file or directory:", dir)
	if err := InjectLicenses(dir, license); err != nil {
		fmt.Println("Oh, there is an error:", err)
	}
}
