package main

import (
	"github.com/gocc/generator/x64"
	"github.com/kr/pretty"

	"github.com/gocc/lexer"
	"github.com/gocc/parser"
)

func main() {
	// code, _ := ioutil.ReadAll(os.Stdin)
	code := []byte("2 == (2+4)/2")

	lexer := lexer.Init(code)
	tokenList := lexer.Tokenize()
	parser := parser.Init(tokenList)
	ast := parser.ParseExpr()
	gen := x64.Init(ast)
	pretty.Println(gen.Asm)
}
