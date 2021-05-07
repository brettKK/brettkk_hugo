4中IR
+ Baf
  + 在bytecode上进行的抽象，忽略字节码对于常量池和type的依赖
+ Jimple
  + 重要的IR，三地址表达式res = num1 op num2，一个表达式最多有三个引用地址
  + 15条指令
  + SSA,Static Single-Assignment。 每个变量仅被赋值一次，保证每个被使用的变量都有唯一的定义
+ Grimple
+ Shimple

Soot中的数据结构
+ Scene类, 环境上下文
+ SootClass， 代表类
+ SootMethod，类的方法
+ SootField，类的属性
+ Body，方法体

Soot支持的分析：
  + Call-graph构建
  + Points-to分析
  + Def/use链
  + 模板驱动的过程间数据流分析
  + 模板驱动的跨过程数据流分析（与heros组合使用）
  + 污点分析（与FlowDroid组合使用）

Body类 （用来存放方法的代码）
  + BafBody, JimpleBody, ShimpleBody, GrimpBody 四种不同IR所对应的Body类
  + 3中主要的chain， Unints chain, Locals chain, Traps chain

Jimple 代码

locals
方法内的局部变量在方法的顶部
$开头的变量为原java代码中没有的变量

traps: 异常处理被表示为一个tuple(元组，异常，开始，结束，处理器)

units: body中包含的实际代码, 语句
  + interface Unit, Jimple使用Stmt， Grimp使用Inst

Value接口（代表数据）： Locals, Constants, Expressions, ParameterRefs

Boxes(代表指针)： ValueBox, UnitBox
