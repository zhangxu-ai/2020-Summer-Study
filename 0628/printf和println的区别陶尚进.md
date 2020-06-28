# Printf和Println

1. println会根据你输入格式原样输出，printf需要格式化输出并带输出格式
2. Println :可以打印出字符串，和变量 
   Printf : 只可以打印出格式化的字符串,可以输出字符串类型的变量，不可以输出整形变量和整形
也就是说，当需要格式化输出信息时一般选择 Printf，其他时候用 Println 就可以了，比如：
```go
a := 10
fmt.Println(a)　　//right
fmt.Println("abc")　　//right
fmt.Printf("%d",a)　　//right
fmt.Printf(a)　　//error
```
3. fmt.println与fmt.printf除了格式化输出以外，有一个小差别，在最后的换行上；
* fmt.printf，如果不带 ，则不会自动加换行； 
* println 会在参数之前加空格 及最后的参数后面加换行