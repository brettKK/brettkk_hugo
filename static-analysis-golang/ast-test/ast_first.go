package main

import (
	"go/token"
	"go/parser"
	"log"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"go/printer"
	"os"
	"go/ast"
	"strings"
)

func main() {
	fs := token.NewFileSet()
	f, err := parser.ParseFile(fs, "", "package main;var a = 3", parser.AllErrors)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(f)
	fmt.Println("``````")
	//fmt.Printf("%#v", f)
	spew.Dump(f)
	fmt.Println("111")
	printer.Fprint(os.Stdout, fs, f)

	var v visitor
	ast.Walk(v, f)

	var v1 visitor1
	ast.Walk(v1, f)

	var v2 visitor2
	ast.Walk(v2, f)
	fmt.Printf("%+v", v2.locals)
}

type visitor struct {}

func (v visitor) Visit(n ast.Node) ast.Visitor {
	fmt.Printf("%T\n", n)
	return v
}

type visitor1 int
func (v visitor1) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	fmt.Printf("%s%T\n", strings.Repeat("\t", int(v)), n)
	return v + 1
}

type visitor2 struct {
	locals map[string]int
}
func (v visitor2) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	switch d := n.(type) {
	case *ast.AssignStmt:
		for _, name := range d.Lhs {
			if ident, ok := name.(*ast.Ident); ok {
				if ident.Name == "_" {
					continue
				}
				if ident.Obj != nil && ident.Obj.Pos() == ident.Pos() {
					v.locals[ident.Name]++
				}
			}
		}
	}
	return v
}