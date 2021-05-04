package main

import (
  "fmt"
  "go/scanner"
  "go/token"
)
//go/scanner, go/parser, go/token, go/ast
func main() {
  src := []byte(`package main
import "fmt"
func main() {
  fmt.Println("Hello, world!")
}
`)

  var s scanner.Scanner
  fset := token.NewFileSet()
  file := fset.AddFile("", fset.Base(), len(src))
  s.Init(file, src, nil, 0)

  for {
     pos, tok, lit := s.Scan()
     fmt.Printf("%-6s%-8s%q\n", fset.Position(pos), tok, lit)

     if tok == token.EOF {
        break
     }
  }
}

/*
1:1   package "package"
1:9   IDENT   "main"
1:13  ;       "\n"
2:1   import  "import"
2:8   STRING  "\"fmt\""
2:13  ;       "\n"
3:1   func    "func"
3:6   IDENT   "main"
3:10  (       ""
3:11  )       ""
3:13  {       ""
4:3   IDENT   "fmt"
4:6   .       ""
4:7   IDENT   "Println"
4:14  (       ""
4:15  STRING  "\"Hello, world!\""
4:30  )       ""
4:31  ;       "\n"
5:1   }       ""
5:2   ;       "\n"
5:3   EOF     ""
*/