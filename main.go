package main

import (
	"io/ioutil"
	"os"

	"github.com/kr/pretty"

	"github.com/gocc/parser"
	"github.com/gocc/tokenizer"
)

func main() {
	code, _ := ioutil.ReadAll(os.Stdin)

	tokenizer := tokenizer.Init(code)
	tokenList := tokenizer.Tokenize()
	// pretty.Println(tokenList)

	parser := parser.Init(tokenList)
	ast := parser.Parse()

	pretty.Println(ast)
	// generator.Generate()
}
