//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/kadenSeaman/lumaCompiler/lexer"
	"github.com/kadenSeaman/lumaCompiler/parser"
)

func parse(this js.Value, args []js.Value) any {
	// Get the source code from the JavaScript argument
	source := args[0].String()

	// Tokenize and parse the source code
	tokens := lexer.Tokenize(source)
	p := parser.NewParser(tokens)

	// Parse and handle errors
	ast, err := p.Parse()
	if err != nil {
		// Return the error message as a string to JavaScript
		return js.ValueOf(fmt.Sprintf("Error parsing: %v", err))
	}

	astJSON, err := json.Marshal(ast)
	if err != nil {
		// Return the error message as a string to JavaScript
		return js.ValueOf(fmt.Sprintf("Error parsing: %v", err))
	}

	// Return the AST as a string representation to JavaScript
	return js.ValueOf(string(astJSON))
}

func main() {
	// Export the parse function to JavaScript
	js.Global().Set("parse", js.FuncOf(parse))

	// Prevent the Go runtime from exiting
	select {}
}
