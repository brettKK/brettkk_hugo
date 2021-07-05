


词法分析： code -> token

lex, flex 

递归下降

LALR(1)

词法分析+语法分析： promql

---

语法分析： token -> ASR

yacc


EBNF grammar



--

## goyacc

安装： go get -u github.com/golang/tools/tree/master/cmd/goyacc

```
type yyLexer interface {
	Lex(lval *yySymType) int
	Error(e string)
}

type yyParser interface {
	Parse(yyLex) int
	Lookahead() int
}
```
