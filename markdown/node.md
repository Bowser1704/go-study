## Q&A

1. 为什么引用类型叫引用类型

   因为返回的是对底层数据类型的引用，自己是没有真正的数据的，例如slice，其实是对底层数组的引用。

## go 工具 

`go test`

该命令用于Go的单元测试，它也是接受一个包名作为参数，如果没有指定，使用当前目录，运行的单元测试必须符合go的测试要求。

1. 写有单元测试的文件名，必须以`_test.go`结尾。
2. 测试文件要包含若干个测试函数。
3. 这些测试函数要以Test为前缀，还要接收一个`*testing.T`类型的参数。



## 1.变量

byte--uint8

rune--int32

### 1.1.运算符

- 运算符优先级，从上至小，从高到低

  ```
  优先级 	运算符
   7 		^ !
   6 		* / % << >> & &^
   5 		+ - | ^
   4 		== != < <= >= >
   3 		<-
   2 		&&
   1 		||
  ```

- 位移运算

  ```go
  func main() {
      a := 10
      b := 1<<a	//无效，移动距离必须是unsigned，或者常量
  }
  ```

  ​	

### 1.2.常量

- 枚举

  没有明确意义上的`enum`，借助iota

    ```go
   const (
        x = iota	//0
        y	//1
        z	//2
  )
    ```
  
- 数字常量不会分配内存

### 1.3.Slice

- 函数之间传递，传递的是slice的副本，但是底层数组不会变

- 左闭右开

- 指针指向第一个slice元素对应的底层数组元素的地址，要注意的是slice的第一个元素并不一定就是数组的第一个元素。

```go
//array
var array []int
array := [5]int{1,2,3,4,5}
array := [...]int{1,2,3,4,5}
//slice
slice := make([]int,5,10)	//后面的是底层数组的长度,可以不指定。
slice := []int{1,2,3,4,5}
slice := []int{4:1}		//指定第五个元素为1

```

```go
//slice 基于现有的切片或者数组创建，会共用同一个底层数组
//对于底层数组容量是k的切片slice[i:j]来说
//长度：j-i
//容量:k-i
newSlice := slice[i:j]
```

此外还有一种3个索引的方法，第3个用来限定新切片的容量，其用法为`slice[i:j:k]`。

```go
slice := []int{1, 2, 3, 4, 5}
newSlice := slice[1:2:3]
//len=2-1=1，cap=3-1=2，第三个值不能大过原长度5
```

切片还有nil切片和空切片，它们的长度和容量都是0，但是它们指向底层数组的指针不一样，nil切片意味着指向底层数组的指针为nil，而空切片对应的指针是个地址。

nil切片表示不存在的切片，而空切片表示一个空集合，它们各有用处。

```go
//nil
var nilSLice []int

//empty
slice := []int{}
```

#### append

```go
slice := []int{1, 2, 3, 4, 5}
newSlice := slice[1:3]
//newSlice := slice[1:3:3]
//output又是怎样？
newSlice=append(newSlice,10)
fmt.Println(newSlice)
fmt.Println(slice)
//Output
[2 3 10]
[1 2 3 10 5]
```

append函数会智能的增长底层数组的容量，目前的算法是：容量小于1000个时，总是成倍的增长，一旦容量超过1000个，增长因子设为1.25，也就是说每次会增加25%的容量。

此外，我们还可以通过`...`操作符，把一个切片追加到另一个切片里

```go
slice := []int{1, 2, 3, 4, 5}
newSlice := slice[1:2:3]

newSlice=append(newSlice,slice...)
fmt.Println(newSlice)
fmt.Println(slice)
```

### 1.4.Map

```go
dict := make(map[string]int)
dict["bowser"] = 1
dict := map[string]int{"bowser":1,"bobo":2}
dict := map[string]int{}
//nil map
var dict map[string]int	//还不能用
dict = make(map[string]int)
```

删除键值对用delete函数

```go
delete(dict, "张三")
for key, value := range dict {
    fmt.Println(key, value)
}
//一个返回值，就是map的键
for key := range dict {
    ...
}
```

函数之间传递map，不是副本。

### 1.5数组

- 一个数组名，是一个数组变量，区别于c语言，不是首地址的指针，函数传参依然是复制，如果传一个数组就会很大了。

- 长度为0的数组在内存中并不占用空间。空数组虽然很少直接使用，但是可以用于强调某种特有类型的操作时避免分配额外的内存空间，比如用于管道的同步操作：
- 我们可以用`fmt.Printf`函数提供的%T或%#v谓词语法来打印数组的类型和详细信息

### 1.6.进制

`strconv`可以不同进制之间转换

### 1.7.type

type与typedef其实类似，自定义类型。

## 2.函数

- 在函数调用时，Go语言没有默认参数值，也没有任何方法可以通过参数名指定形参，因此形参和返回值的变量名对于函数调用者而言没有意义。

- 你可能会偶尔遇到没有函数体的函数声明，这表示该函数不是以Go实现、的。这样的声明定义了函数标识符。

- 自己最好吧参数列表，返回值列表都写出来，并且加上括号

- 没有条件的switch可以用

  ```go
  package main
  
  import (
  	"fmt"
  	"time"
  )
  
  func main() {
  	t := time.Now()
  	switch {
  	case t.Hour() < 12:
  		fmt.Println("Good morning!")
  	case t.Hour() < 17:
  		fmt.Println("Good afternoon.")
  	default:
  		fmt.Println("Good evening.")
  	}
  }
  ```

  - defer语句虽然推迟到外层函数返回之后执行。但是：

  `推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。`

- 可变参数

  ```go
  func print (a ...interface{}){ 
     for _,v:=range a{
          fmt.Print(v)
      }
      fmt.Println()
  }
  ```

  

### 2.1.方法

方法有一个接受对象

- 最好对象那里写指针，传递时可以不用写地址

  ```go
  func main() {
      p:=person{name:"张三"}
      p.modify() //指针接收者，修改有效
      fmt.Println(p.String())
  }
  type person struct {
      name string
  }
  func (p person) String() string{ 
     return "the person name is "+p.name
  }
  func (p *person) modify(){
      p.name = "李四"
  }
  ```

### 2.2.接口

- 实体类型以指针接收者实现接口的时候，只有指向这个类型的指针才被认为实现了该接口

  | Methods Receivers | Values   |
  | ----------------- | -------- |
  | (t T)             | T and *T |
  | (t *T)            | *T       |

  如果是值接收者，实体类型的值和指针都可以实现对应的接口；如果是指针接收者，那么只有类型的指针能够实现对应的接口。

  | Values | Methods Receivers |
  | ------ | ----------------- |
  | T      | (t T)             |
  | *T     | (t T) and (t *T)  |

  类型的值只能实现值接收者的接口；指向类型的指针，既可以实现值接收者的接口，也可以实现指针接收者的接口。

### 2.3.goroutine

```go
func main() {
    //runtime.GOMAXPROCS(1)
    var wg sync.WaitGroup
    wg.Add(2)    
    
    go func(){
            defer wg.Done()        
            for i:=1;i<100;i++ {
                fmt.Println("A:",i)
            }
    }()    
       
    go func(){
            defer wg.Done()
            for i:=1;i<100;i++ {
                fmt.Println("B:",i)
            }
    }()
    wg.Wait()
}
```

#### channel

单向通道，无缓冲，有缓存

```go
var send chan<- int		//send是针对于channel使用者来说的，使用者只能send到channel中
var receive <-chan int	//即说明这个通道只能发送数据给使用者
```

#### Context

使用context来控制`goroutine`，

## 3. go和web

### Middleware

- Handle
- `HandleFunc`





