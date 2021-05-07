AnalysisScope
+ 用于指定用户静态代码分析的程序或库
+ format: Classloader, Language,  Type, Location

CallGraphBuilder
  + makeCallGraph() = CallGraph
  + getPointerAnalysis() = PointerAnalysis

Representation of Code Structure
+ AnalysisScope
+ IClassHierarchy

---
IR -> 代码生成器 -> 目标代码

代码生成器： 指令选择，寄存器分配和指派，指令排序
