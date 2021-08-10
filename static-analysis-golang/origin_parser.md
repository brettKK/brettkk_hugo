+ go compiler的实现入口在src/cmd/compile
    + cmd/compile/internal/gc.Main
    + 初始化ssa配置，cmd/compile/internal/gc.initssaconfig
    + cmd/compile/internal/gc.funccompile
    + 优化中间代码的过程在cmd/compile/internal/ssa.Compile
+ cmd/internal/obj
+ golang compiler
    + go gccgo
    + closed project tulgo
    + open source llgo GopherJS TARDIS Go
        + ssa --> llgo --> llvm module --> llgo-build --> executable 
        + github.com/axw/llgo
    + Go SSA viewer  http://golang-ssaview.herokuapp.com/

+ file.cpp --clang--> clang AST -> LLVM IR -> Selection DAG --llc--> LLVM MIR --llc--> Machine Code
+ golang code -> AST -> LLVM IR -> LLVM bytecode -> ASM -> Native
+ source code -> AST
    + parser.Parse* (ParseExpr, ParseExprFrom, ParseFile, ParseDir)
+ llvm api(godoc), binding go (module, type, values, Basic Blocks, )
  + module := llvm.NewModule("main")
  + import "llvm.org/llvm/bindings/go/llvm"

+ Passes in llvm (levels of granularity)
  + Module Pass, Call Graph Pass, Function Pass, Basic Block Pass
+ Identifier
  + @ global, % local

+ 有用package
  + ast
  + 
+ go 工具链， go build, go test, go fmt, go get, go vet, go tool cover, go tool pprof, go -race
    + GOSSAFUNC=HelloWorld go build
    + cmd/compile/README, cmd/compile/internal/ssa/README
    + go tool compile <=> go tool objdump 
    + obj files ==> go tool link ==> binary file
+ go ssa (how to read go SSA intermediate representation)
    + 
+ go/ https://golang.org/s/types-tutorial
    + go/doc go/printer go/scanner
    + token,  FileSets
    + ast
        + Package File Scope Object
        + interface: Node, Decl, Spec, Stmt, Expr
        + https://golang.org/ref/spec
    + parser
        + ParseExpr(x string)
        + ParseFile
        + ParseDir
    
+ golang.org/x/tools/go
    + packages (recommend~)

    + ssa
        + Member, Value, Instruction, Node
    + callgraph
        + go get golang.org/x/tools/cmd/callgraph
        + Graph Node Edge
    + analysis
        + /singlechecker

  + golang.org/x/tools/go/ssa (2015年 用到这个package的项目)
    + stackcheck
    + go-llvm/llgo 如何挖掘出有用的东西
    + go-type-switch-gen
    + stripe/safesql, 检查sql injection
    + tardisgo

+ 如何静态分析golang代码
    + go/scanner包 词法分析 变为token
    + ast和parser的实现，源码在<go sdk>/src/go/parser包下

+ 开源项目例子 llgo

+ create own static analysis tools, pay more attention to algorithm, design and performance
+ static analysis tools for go
    + gofmt/goimports
    + go vet/golint github.com/golang/go/tree/master/src/cmd/vet
    + guru
    + gocode
    + errcheck
    + gorename/gomvpkg
+ static analysis background
    + https://github.com/mre/awesome-static-analysis#go
    + https://github.com/golang/lint
    + https://github.com/golangci/golangci-lint
    + https://github.com/360EntSecGroup-Skylar/goreporter
    + https://github.com/alecthomas/gometalinter
    + http://golang.org/lib/godoc/analysis/help.html
    + Oracle (not open source) 
    + github.com/golang/example/tree/master/gotypes
    + github.com/retailnext/stan
    + godoc.org/golang.org/x/tools/go/loader
    + godoc.org/golang.org/x/tools/go/analysis
    + godoc.org/golang.org/x/tools/go/ssa
    + godoc.org/golang.org/x/tools/go/pointer
    + golang.org/x/tools/go/analysis
        + defines the interface between a modular static analysis and an analysis driver program
    + golang.org/x/tools/go/callgraph
    + golang.org/x/tools/go/callgraph/cha
    + golang.org/x/tools/go/callgraph/rta
    + golang.org/x/tools/go/callgraph/static
+ dynamic analysis 
    + dmitry vyukov

+ 反编译二进制文件
    + linux objdump
    + mac otool


+ 内置 Parser 打造轻量级规则引擎
    + 业务系统就是对PM的需求堆if else，基础服务就是对OS资源和网络问题堆if else
    + 造成了代码的膨胀，如何来管理这坨if else，就衍生出了设计模式：通过各种手段来把if else划分到更小的粒度
    + 不论用什么方法，if else不会凭空消失，只会从一个地方转移到另一个地方 可维护性以及扩展性是可以得到提升的


+ 符号执行
    + 约束规划的求解问题 constraint programming, SMT(satisfiability modulo theory) ,  SAT


+ 软件静态漏洞检测
    + 二进制漏洞检测
    + 源码漏洞检测
        + 基于代码相似性...
        + 基于符号执行
        + 基于规则
        + 基于机器学习

+ wangcong15

+ llvm 中文资料  白马负金羁 梦在哪里 csdn
+ http://llvm.org/docs/WritingAnLLVMPass.html
+ http://llvm.org/docs/tutorial/index.html
+ http://www.aosabook.org/en/index.html, http://www.aosabook.org/en/llvm.html
+ MIT Computer Systems Security
https://ocw.mit.edu/courses/electrical-engineering-and-computer-science/6-858-computer-systems-security-fall-2014/index.htm


+ https://sitano.github.io/2018/03/18/howto-read-gossa/
+ https://github.com/golang/tools/blob/master/go/ssa/doc.go
+ https://github.com/golang/go/blob/master/src/cmd/compile/internal/gc/ssa_test.go
+ https://github.com/golang/go/tree/master/src/cmd/compile/internal/ssa
+ https://github.com/golang/go/blob/master/src/cmd/compile/internal/ssa/gen/genericOps.go
+ https://golang.org/pkg/cmd/compile/internal/ssa/
+ https://github.com/golang/tools

源伞科技

三板斧：DELF, DALF, TCF, TEF

你必须只有内心丰富，才能摆脱这些生活表面的相似。 王朔 

【sticksology欧乐集】皇家玫瑰红茶棒

Golang并发缺陷分析

生活中其实没有那么多皆大欢喜的故事，更没有那么多浮想联翩的奇迹，你只能拼命努力，去换一个还不错的结局。

	你要解决什么问题

知行合一： 知道且做到才是知道。知道但做不到等于不知道。
