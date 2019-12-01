# Go语言学习笔记

## 基础

### 变量声明

```go
i := 1
//短变量声明，只能用于函数内部，不能用于包级变量
//在for循环体内的initialization只能使用这种方法来声明变量
//短变量声明最少必须要声明一个变量，已经声明过得实际上是赋值行为，必须是当前词法域

var s string
//初始化为类型的初值

var s = ""
//根据初值，自动判断类型

var i uint32 = 1
//手动设置类型，并且给初值，默认判断类型会有疑问，例如1到底是什么类型，int？uint32？int64？。Go是强类型语言，所有类型都不能自动转换。
```

所以短变量的申明的默认初始类型是什么呢？

```go
import ."reflect"

i := 1
//TypeOf(i) == int
f := 1.0
//TypeOf(f) == float64

//要明确这个，否则会有下面的错误.
var i uint64 = 1;
a := 0
a = i //错误，不能隐式uint64转int


i := 0      // untyped integer;        implicit int(0)
r := '\000' // untyped rune;           implicit rune('\000')
f := 0.0    // untyped floating-point; implicit float64(0.0)
c := 0i     // untyped complex;        implicit complex128(0i)
```

### 作用域

首先作用域与生命周期是两个东西，后者是变量存活的时间，特别是go这种带gc的。

代码块指定了作用域，句法块内部声明的名字是无法被外部块访问的。存在一个整体的词法块，称为全局词法块；即有一个全局作用域。**编译器看到一个名字是会从内部作用域向外(全局)作用域查找的**。

像类似for循环，if和switch语句，会在条件部分创建隐式词法域，也就是说其实有两个词法域，条件部分，循环体/执行体部分，前者在后者的外部(也就是后面的可以访问前面的)

```go
//容易犯的错误
var cwd string

func init() {
    cwd, err := os.Getwd() // compile error: unused: cwd
    //如果是短变量声明，他在本地作用域没有发现cwd这个变量，那就会自己声明一个，而不是cwd是赋值行为。
    /*修改为
    var err error
    cwd, err = os.Getwd()
    */
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

### 数据类型

#### rune 
代表Unicode一个码点(Code Point，以U+开头，A是U+0061)，底层是int32，也就是4个字节大小个数字，涉及到Unicode版本的问题。

在go中`\uhhhh`对应16bit的码点值，`\Uhhhhhhhh`对应32bit的码点，小于256码点值可以直接用16进制数，例如\x41对应字符'A'，\xe4\xb8\x96并不是一个合法的rune字符。

Unicode并不涉及字符是怎么在字节中表示的，它仅仅指定了字符对应的数字，仅此而已。
Unicode是给所有字符一个编号，而类似UTF-8这种是怎么储存在计算机里面(考虑到节省存储空间)，UTF-8每个码点储存1-3个字节。

#### 结构体

strcut 为关键字，type就是定义别名类型，嵌入结构体q可以直接调用底层结构体p，但是q **is not a** p，Go对于类型是十分严格的，基本上只要是一个类型，那他就是这个类型，不会是同时是多个类型。**除了interface**，


####  字符串处理

Go中储存的数据都是utf-8编码，所以我们对字符串的操作可以直接感受底层的字节

```go
//string的强转是生成对应的Unicode码点，utf-8字符，
//底层每一个元素类型是 int32/rune
string(65)// 
//range 遍历字符串得到的是 rune

//字符串是只读的，不能修改底层内存存储的值
func HasPrefix(s, prefix string) bool {
    return len(s) > len(prefix) && s[:len(prefix)] == prefix
    //len返回的是字节个数
}

func HasSuffix(s, suffix string) bool {
    return len(s) > len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
    for i := 0; i < len(s); i++ {
        if HasPrefix(s[i:], substr) {
            return true
        }
    }
    return false
}
```



- 浮点数 和其他语言差不多，浮点类型，注意精度，注意范围
- bool   c99之前c语言没有bool类型，1为true，0为false，go语言区分所有类型，bool就是bool，所有条件语句条件内都要是bool类型。
- 常量 在编译时计算，可以理解为define，常量可以使无类型的，

### 复合数据类型

1. 数组
   1. 数组的长度是数组的一部分，即不同长度数组为不同数据类型
   2. 数组的长度必须是常量，编译阶段就要确定。
   3. 如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的，即长度和元素类型都一样的数组才可以比较。

### 文本和html模块

主要是两个package的使用，html和text。

### 运算符

Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此`-5%3`和`-5%-3`结果都是-2。

位操作运算符`^`作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；也就是说，它返回一个每个bit位都取反的数。位操作运算符`&^`用于按位置零（AND NOT）：如果对应y中bit位为1的话, 表达式z = x `&^` y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。



### 函数

函数是一等公民，可以作为函数的参数，**也可以作为type，实现interface**

#### 闭包

将函数作为参数，或者返回值，这在有些时候会很有意义，例如我们遍历二叉树，遍历作为一个函数，传入一个utility函数，utility函数作为对每个节点的行为，这样的话，遍历函数可以一直复用。

对函数实现interface

#### Panic和Recover

## 测试

### External Test Packages

#### export_test.go/并非一定export

我们写测试时，测试文件以`_test.go`结尾，测试函数以`Test`+大写单词开头，一般我们写黑盒测试，测试文件的package名字为`原包名_test`，有时候我们需要使用待测试包未导出的包，白盒测试。

整理一下，我们的需求是这样的:

1. 不将未导出标识符导出到生产代码中
2. 将一些未导出标识符导出到测试代码中

很简单我们写一个_test.go文件，把未导出的标识符，命名中间变量导出，而`test.go`为后缀的文件，只有`go test`才会读取，也就保护了原代码，也可以测试了。fmt就使用了。

## net相关

### Handler

```go
func ListenAndServe(address string, h Handler) error
//Handler是一个interface
type Handler interface {
    ServerHttp(w ResponseWriter, r *Request)
}
//HandlerFunc是一个是函数，pattern是endpoint，无返回值。
//HandleFunc registers the handler function for the given pattern in the DefaultServeMux. 
func HandleFunc(pattern string, handler func(ResponseWriter, *Request))

//HandlerFunc是一个类型,HandlerFunc是一个让函数值满足一个接口的适配器
type HandlerFunc func(w ResponseWriter, r *Request)
func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```



### 小问题

1. 为什么**Response.Body**每次都要手动关闭。

   底层的实现是当Body被读完之后，即读到EOF就会自动关闭Body，但是没有读完的话，是不会关闭的。
```

```

## gin源码

### [routergroup.go](https://github.com/gin-gonic/gin/blob/master/routergroup.go)

> The router routes you to a route.

关于路由结尾是否加`/`，可以参考flask[文档](https://flask.palletsprojects.com/en/1.1.x/quickstart/)。有`/`类似于文件夹(folder)，所以我们写API时(纯后端)，不加`/`，但是前端会加，因为不仅仅是一个文件。肤浅的理解。

一个路由除了url，还要有http method，也即是method+path+HandlerFunc == router ==> function 。

routergroup是gin实现的路由管理。

```go
// RouterGroup is used internally to configure router, a RouterGroup is associated with a prefix
// and an array of handlers (middleware).
type RouterGroup struct {
	Handlers HandlersChain
	basePath string
	engine   *Engine	//后面提到的Engine，必须是指针，否则就套娃了。
	root     bool		//TODO 什么意思??
}
// Group creates a new router group. You should add all the routes that have common middlwares or the same path prefix.
// For example, all the routes that use a common middlware for authorization could be grouped.
func (group *RouterGroup) Group(relativePath string, handlers ...HandlerFunc) *RouterGroup {
	return &RouterGroup{
		Handlers: group.combineHandlers(handlers),
		basePath: group.calculateAbsolutePath(relativePath),
		engine:   group.engine,
	}
}
//把group指定的relativePath加到bashPath上面去
func (group *RouterGroup) calculateAbsolutePath(relativePath string) string {
	return joinPaths(group.basePath, relativePath)
}
//也就是说Group是一个方法，返回*RouterGroup，是一个router的group，一般是同一个prefix，用相同的middleware。

//也就是我们router使用的，但是我们传入的是gin.Engine，这是一个struct，是gin的核心struct。
//*Engine是实现了IRoutes的，是*Engine.
type Engine struct {
    RouterGroup
    //....
}

//IRoutes defines all router handle interface.
type IRoutes interface {
	Use(...HandlerFunc) IRoutes

	Handle(string, string, ...HandlerFunc) IRoutes
	Any(string, ...HandlerFunc) IRoutes
	GET(string, ...HandlerFunc) IRoutes
	POST(string, ...HandlerFunc) IRoutes
    //....
    //*RouterGroup实现了这些方法，Engine含有的是RouterGroup,没有实现接口IRoutes
}
// IRouter defines all router handle interface includes single and group router.
//Routes加Group可以作为Router
type IRouter interface {
	IRoutes
	Group(string, ...HandlerFunc) *RouterGroup
}

//重点是下面的。
func (group *RouterGroup) handle(httpMethod, relativePath string, handlers HandlersChain) IRoutes {
	absolutePath := group.calculateAbsolutePath(relativePath)
	handlers = group.combineHandlers(handlers)
	group.engine.addRoute(httpMethod, absolutePath, handlers)
    return group.returnObj()//return group(*RouterGroup) 或者 group.engine(*Engine)
}
```

```go
//关于*Engine是否满足IRoutes接口
type test interface{
    vv()int
}

type base struct {
    *advance
}
type advance struct {
    base					//false 如果换成*base也满足了接口test
    //*base 				  true
    //b *base 				  flase
}

func (b *base) vv() int {
    return 0
}

func main() {
    var b,a interface{};
    b = &base{}
    _, ok := b.(test)
    fmt.Println(ok)			//true
    a = advance{}
    //a= &advance{}			true 它实现了实现了......
    _, ok = a.(test)
    fmt.Println(ok)			//false
}

/*-------------------------*/
//然而
func main() {
    var test interface{}
    test = &gin.Engine{}
    _, ok := test.(gin.IRoutes)
    fmt.Println(ok)		//true
}
```


### [gin.go](https://github.com/gin-gonic/gin/blob/master/gin.go)

```go
// Default returns an Engine instance with the Logger and Recovery middleware already attached.
func Default() *Engine {
	debugPrintWARNINGDefault()
	engine := New()						//这就是gin框架的核心，核心引擎
	engine.Use(Logger(), Recovery())	//加上两个中间件，middleware
	return engine
}

//New函数做一些初始化操作，对于Engine的操作，我们先注意几点
engine.RouterGroup.engine = engine

//下面是初始化pool对象，以及allocateContext()初始化上下文对象。
//pool是临时对象池，用于储存context对象
engine.pool.New = func() interface{} {
	return engine.allocateContext()	
}
```

gin中重新定义了HandlerFunc

```go
type HandlerFunc func(c *gin.Context)
//so
g.Get("/", foo)	//相当于http中的函数HandleFunc
func foo(c *gin.Context) {
    //....
}
```

## interface与reflect

interface申明一些方法，如果其他类型实现了这些方法，那么就实现了这个interface，interface只能使用自己申明的方法，如果其他类型有其他方法，也是不能**直接**调用的，但是可以使用类型断言去扩展他的方法，代码示例

```go
var r io.Reader
r = io.Stdout			//r只有Read方法
rw = r.(io.ReadWriter)	//rw有Read和Write方法

w, ok := rw.(io.Writer)	//可以通过判断ok，决定时候使用某些方法，比如说我要使用String方法，定义一个只有一个方法的interface，利用类型断言。

//使用i.(type)类型断言，可以利用类型分支
```

interface有两个值，一个是动态type，一个是动态value，也即是该接口真正的指向对象类型是什么，真正的指向对象是什么。

如果struct A has a struct B，这种情况到底是否实现了interface呢?

```go
type test interface{
    vv()int
}

type base struct {
    *advance
}

type advance struct {
    base					//false 如果换成*base也满足了接口test
    //*base 				  true
    //b *base 				  flase
}

func (b *base) vv() int {
    return 0
}

func main() {
    var b,a interface{};
    b = &base{}
    _, ok := b.(test)
    fmt.Println(ok)			//true
    a = advance{}
    //a= &advance{}			true 它实现了实现了...... 也就是如果是advance指针，那么里面的base也当做指针，也是实现了vv方法的，或者advance底层就是有*base那么也实现了。
    _, ok = a.(test)
    fmt.Println(ok)			//false
}
```



### reflect

但是我们会发现，type assertion 只会返回底层类型，而不会返回我们自定义类型，例如

```go
type values map[string][]string
var x values
x.(type)			//map[]
reflect.TypeOf(x)	//main.values，返回的是具体类型
```

reflect下面有Type，Value，Kind。

- Type是具体类型，例如main.Values
- Value Value底层是一个strcuct，包含了我们需要的数据，例如Type，Kind。
- Kind只关心底层类型，map，int.... 

reflect可以访问未导出的字段，通过Value。



## 杂乱

### Addressable

