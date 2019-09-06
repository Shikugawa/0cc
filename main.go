package main

import (
	"io/ioutil"
	"os"

	"github.com/kr/pretty"

	"github.com/gocc/generator"
	"github.com/gocc/lexer"
	"github.com/gocc/parser"
)

func main() {
	code, _ := ioutil.ReadAll(os.Stdin)

	lexer := lexer.Init(code)
	tokenList := lexer.Tokenize()

	parser := parser.Init(tokenList)
	ast := parser.Parse()
	// pretty.Println(ast)
	gen := generator.Init(ast)

	pretty.Println(gen.Asm)
}
