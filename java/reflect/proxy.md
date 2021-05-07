java动态代理

作用
+ 增加方法，加日志，加监控等
+ RPC的基础。 接口的实现是动态代理对象，方法名和参数传送到远程服务器，远程服务器执行响应方法返回结果

java.lang.reflect.Proxy类 提供一些静态方法来创建动态代理类，并且是所有动态代理类的父类

Proxy类与普通类的唯一区别是： proxy类的字节码不预存在.class文件中，而是有JVM在运行时动态生成的，每次生成动态代理类时都需要classloader

// 获取指定代理关联的调用处理器
public static InvocationHandler getInvocationHandler(Object proxy)

// 获取动态代理类的类对象
public static Class<?> getProxyClass(ClassLoader loader, Class<?>... interfaces)

// 判断一个类对象是否为动态代理类
public static boolean isProxyClass(Class<?> cl)

// 生成动态代理类的实例
public static Object newProxyInstance(ClassLoader loader, Class<?>[] interfaces, InvocationHandler h)
1 通过Proxy的getProxyClass获取代理类的类对象
2 通过反射从代理类的类对象上获取构造函数对象
3 通过构造函数对象创建代理类的实例
----

java.lang.reflect.InvocationHandler：调用处理器的接口
// 在该方法上实现对代理类的代理访问
Object invoke(Object proxy, Method method, Object[] args)

java.lang.ClassLoader: 类加载器。 定义类对象，之后类才能被使用
