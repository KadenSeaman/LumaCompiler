package main

import (
	"fmt"
	"os"

	"github.com/kadenSeaman/lumaCompiler/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/test.lang")

	tokens := lexer.Tokenize(string(bytes))

	p := parser.newParser(tokens)

	ast, err := p.parse()

	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	fmt.Printf("Parsed AST: %+v\n", ast)
}
