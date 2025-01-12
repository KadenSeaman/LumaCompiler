//go:build js && wasm

package main

import (
	"encoding/json"
	"fmt"
	"syscall/js"

	"github.com/kadenSeaman/lumaCompiler/lexer"
	"github.com/kadenSeaman/lumaCompiler/parser"
)

// Wrapper for uintptr mapping to the AST for tracking in Go
var astMap = make(map[uintptr]any)
var astCounter uintptr = 1 // Start counting from 1 to ensure non-zero pointers

func main() {
	js.Global().Set("parse", js.FuncOf(func(this js.Value, args []js.Value) any {
		source := args[0].String()

		tokens := lexer.Tokenize(source)
		p := parser.NewParser(tokens)

		ast, err := p.Parse()
		if err != nil {
			fmt.Println("Error parsing:", err)
			return nil
		}

		// Store AST in map and get a unique uintptr
		ptr := astCounter
		astMap[ptr] = ast
		astCounter++

		// Return the uintptr as the identifier
		return uintptr(ptr)
	}))

	js.Global().Set("getAST", js.FuncOf(func(this js.Value, args []js.Value) any {
		ptr := uintptr(args[0].Int())
		if ast, exists := astMap[ptr]; exists {
			astJSON, err := json.Marshal(ast)

			if err != nil {
				fmt.Println("Error converting ast to json")
				return nil
			}

			return js.ValueOf(string(astJSON))
		}
		return js.ValueOf("AST not found")
	}))

	js.Global().Set("releaseAST", js.FuncOf(func(this js.Value, args []js.Value) any {
		ptr := uintptr(args[0].Int())
		delete(astMap, ptr)
		return nil
	}))

	select {}
}
