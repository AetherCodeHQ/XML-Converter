package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: xml-converter <file.xml>")
		os.Exit(1)
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	s := string(data)
	// basit string tabanli XML element/attribute sayma
	tags := 0
	attrs := 0
	depth := 0
	maxDepth := 0
	i := 0
	for i < len(s) {
		if s[i] == '<' {
			if i+1 < len(s) && s[i+1] == '/' {
				tags++
				depth--
				i += 2
				for i < len(s) && s[i] != '>' { i++ }
				i++
				continue
			}
			// self-closing?
			end := strings.Index(s[i:], ">")
			if end < 0 {
				break
			}
			tagContent := s[i+1 : i+end]
			if strings.HasSuffix(tagContent, "/") {
				tags++
				i += end + 1
				continue
			}
			tags++
			depth++
			if depth > maxDepth {
				maxDepth = depth
			}
			// count attributes: words with =
			parts := strings.Fields(tagContent)
			for _, p := range parts[1:] {
				if strings.Contains(p, "=") {
					attrs++
				}
			}
			i += end + 1
		} else {
			i++
		}
	}
	fmt.Printf("elements=%d attributes=%d max_depth=%d\n", tags, attrs, maxDepth)
	fmt.Printf("size=%d bytes\n", len(s))
}
