package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/kadenSeaman/lumaCompiler/lexer"
	"github.com/kadenSeaman/lumaCompiler/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/example00.lang")

	tokens := lexer.Tokenize(string(bytes))
	p := parser.NewParser(tokens)

	ast, err := p.Parse()

	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	// fmt.Printf("Parsed AST: \n%+v\n", ast)

	jsonFormat, err := json.Marshal(ast)

	if err != nil {
		fmt.Println("Error converting ast to json")
		return
	}

	fmt.Println(string(jsonFormat))
}
