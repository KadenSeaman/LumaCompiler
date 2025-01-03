package main

import (
	"os"

	"github.com/kadenSeaman/lumaCompiler/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/test.lang")

	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		lexer.Debug(token)
	}
}
