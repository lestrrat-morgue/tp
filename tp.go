package main

import (
	"encoding/json"
	"fmt"
	"os"
	"text/template"
)

func main() {
	os.Exit(_main())
}

func _main() int {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage: tp template < variables.json\n")
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