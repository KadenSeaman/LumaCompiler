//go:build js && wasm

package main

import (
	"encoding/json"

	"github.com/kadenSeaman/lumaCompiler/lexer"
	"github.com/kadenSeaman/lumaCompiler/parser"
)

//export parseLuma
func parseLuma(source string) string {
	// Tokenize and parse the source code
	tokens := lexer.Tokenize(source)
	p := parser.NewParser(tokens)

	// Parse and handle errors
	ast, err := p.Parse()
	if err != nil {
		// Return the error message as a string to JavaScript
		return "Error parsing: " + err.Error()
	}

	astJSON, err := json.Marshal(ast)
	if err != nil {
		// Return the error message as a string to JavaScript
		return "Error parsing: " + err.Error()
	}

	// Return the AST as a string representation to JavaScript
	return string(astJSON)
}

func main() {
	// Create a channel to keep the program running
	c := make(chan struct{})
	<-c
}
