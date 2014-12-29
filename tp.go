package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

const version = "0.0.1"

func main() {
	os.Exit(_main())
}

func usage() {
	fmt.Fprintf(os.Stderr, `tp version %s

Usage:
  e.g.
    tp template.tmpl < variables.json
    curl http://example.com/resource.json | tp template.tmpl
`, version)
}

func _main() int {
	if len(os.Args) != 2 {
		usage()
		return 1
	}

	t := template.New("tp")
	t.Funcs(template.FuncMap{
		"int": func(f float64) string { return fmt.Sprintf("%0.0f", f) },
	})

	if _, err := t.ParseFiles(os.Args[1]); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse template: %s\n", err)
		return 1
	}

	var v interface{}
	dec := json.NewDecoder(os.Stdin)
	if err := dec.Decode(&v); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse json: %s\n", err)
		return 1
	}

	if err := t.ExecuteTemplate(os.Stdout, os.Args[1], v); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to execute template: %s\n", err)
		return 1
	}
	return 0
}