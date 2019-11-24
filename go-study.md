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

- rune 
  代表Unicode一个码点(Code Point，以U+开头，A是U+0061)，底层是int32，也就是4个字节大小个数字，涉及到Unicode版本的问题。

  在go中`\uhhhh`对应16bit的码点值，`\Uhhhhhhhh`对应32bit的码点，小于256码点值可以直接用16进制数，例如\x41对应字符'A'，\xe4\xb8\x96并不是一个合法的rune字符。

  Unicode并不涉及字符是怎么在字节中表示的，它仅仅指定了字符对应的数字，仅此而已。
  Unicode是给所有字符一个编号，而类似UTF-8这种是怎么储存在计算机里面(考虑到节省存储空间)，UTF-8每个码点储存1-3个字节。

####  字符串处理

Go中储存的数据都是utf-8编码，所以我们对字符串的操作可以直接感受底层的字节

```go
//string的强转是生成对应的Unicode码点，utf-8字符，
string(65)// A

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

### 运算符

Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此`-5%3`和`-5%-3`结果都是-2。

位操作运算符`^`作为二元运算符时是按位异或（XOR），当用作一元运算符时表示按位取反；也就是说，它返回一个每个bit位都取反的数。位操作运算符`&^`用于按位置零（AND NOT）：如果对应y中bit位为1的话, 表达式z = x `&^` y结果z的对应的bit位为0，否则z对应的bit位等于x相应的bit位的值。

## net相关

### 小问题

1. 为什么**Response.Body**每次都要手动关闭。

   底层的实现是当Body被读完之后，即读到EOF就会自动关闭Body，但是没有读完的话，是不会关闭的。
```

```