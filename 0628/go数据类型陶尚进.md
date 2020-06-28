# 数据类型


## 1、整型   
   按长度：int8 int16 int32 int64
   对应的无符号整型：uint8 uint16...

## 2、浮点型
   float32、float64
   打印浮点数时，可以使用fmt包配合动词“%f”:

```go
package main

import (
        "fmt"
        "math"
        )

func main() {

fmt.Printf("%f\n",math.Pi)

fmt.Printf("%.2f\n",math.Pi)

}
```


float64(x)表示将x变量转化为float64类型，之后运算的所有表达式都以float64类型进行。

## 布尔类型
var n bool//bool类型无法参与数值运算，也无法与其他数据类型进行转换。

## 字符串
字符串以原生数据类型出现，直接使用：str := "hello"   不能跨行，如果需要跨行，则需要"`",反引号之间的转义字符无效。一般内嵌源码数据
                                            ch ：= "中文"                         ` `之间所有代码作为字符串，不会被编译器识别。
	
字符串转义符：
\r    回车
\n    换行
\t    制表
\'    单引号
\"    双引号
\\    反斜杠

uint8=byte，ASCII码
rune ： UTF-8字符，实际上是int32
%T；输出变量实际类型

## 切片
拥有相同类型元素的可变长度的序列
`var name []T`

```go
a := make([]int,3)//创建一个容量为3的切片
a[0] = 1 //赋值
a[1] = 2
a[2] = 3
```

