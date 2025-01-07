package main

import (
	"fmt"
	"os"

	"github.com/kadenSeaman/lumaCompiler/lexer"
	"github.com/kadenSeaman/lumaCompiler/parser"
)

func main() {
	bytes, _ := os.ReadFile("./examples/test.lang")

	tokens := lexer.Tokenize(string(bytes))
	p := parser.NewParser(tokens)

	ast, err := p.Parse()

	if err != nil {
		fmt.Println("Error parsing:", err)
		return
	}

	fmt.Printf("Parsed AST: \n%+v\n", ast)

}
