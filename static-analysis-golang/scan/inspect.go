package main

import (
  "fmt"
  "go/ast"
  "go/parser"
  "go/printer"
  "go/token"
  "log"
  "os"
  "golang.org/x/tools/go/packages"
)

func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)
  packages.Load

  fset := token.NewFileSet()

  file, err := parser.ParseFile(fset, "", src, 0)
  if err != nil {
     log.Fatal(err)
  }

  ast.Inspect(file, func(n ast.Node) bool {
     call, ok := n.(*ast.CallExpr)
     if !ok {
        return true
     }

     printer.Fprint(os.Stdout, fset, call.Fun)
     fmt.Println()

     return false
  })
}
//fmt.Println

