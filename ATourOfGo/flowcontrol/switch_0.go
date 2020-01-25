package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
		// break
	case "linux":
		fmt.Println("Linux.")
		// break
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}

// "switch" in GoLang
// Go automatically provides the "break" statement required after each "case" in these languages
