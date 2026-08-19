package main

import (
	"fmt"
	"os"
)

// xml_converter - XML parsing and conversion
func xml_converter(path string) {
	fmt.Println("========================================")
	fmt.Println("  XML-Converter")
	fmt.Println("  XML parsing and conversion")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	xml_converter(path)
}
