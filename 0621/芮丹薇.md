#python语言元素
##变量和类型
整型（int）：Python中可以处理任意大小的整数，而且支持二进制（如0b100，换算成十进制是4）。  
八进制（如0o100，换算成十进制是64）、十进制（100）和十六进制（0x100，换算成十进制是256）的表示法。  
浮点型（float）：浮点数也就是小数，之所以称为浮点数，是因为按照科学记数法表示时，一个浮点数的小数点位置是可变的，浮点数除了数学写法（如123.456）之外还支持科学计数法（如1.23456e2）。  
字符串型（str）：字符串是以单引号或双引号括起来的任意文本，比如'hello'和"hello"。  
布尔型（bool）：布尔值只有True、False两种值，要么是True，要么是False。
##变量命名
###硬性规则：
规则1：变量名由字母、数字和下划线构成，数字不能开头。  
规则2：大小写敏感。  
规则3：变量名不要跟Python语言的关键字（有特殊含义的单词，后面会讲到）和保留字（如函数、模块等的名字）发生重名的冲突。
###非硬性规则：
规则1：变量名通常使用小写英文字母，多个单词用下划线进行连接。  
规则2：受保护的变量用单个下划线开头。  
规则3：私有的变量用两个下划线开头。
##变量的使用
###使用变量保存数据并进行加减乘除运算。

![](http://m.qpic.cn/psc?/V10OQot10zcV7k/wSJ2S*tZT7v.5zxXfWcfXSE6XGButqxPQ.SE6boroaagFi1wExmetERscFt1YGyb0MqnxfoRwm.mvrGD59lc9RF*WqhZ7te.rlihlvJp6tc!/b&bo=sAKmAAAAAAADFyY!&rf=viewer_4)
###在Python中可以使用type函数对变量的类型进行检查。

![](https://m.qpic.cn/psc?/V10OQot13yzIbG/ZUXCJANDCZJwMw9eDcmXBZzi*kqldLKplWoj2xozdU*ms7fh9auCWjBbKrGzOqR4h8vfiZcfMFDqftCU3*uFgg!!/b&bo=8AHSAAAAAAARBxM!&rf=viewer_4)
###不同类型的变量可以相互转换，这一点可以通过Python的内置函数来实现。
int()：将一个数值或字符串转换成整数，可以指定进制。  
float()：将一个字符串转换成浮点数。  
str()：将指定的对象转换成字符串形式，可以指定编码。  
chr()：将整数转换成该编码对应的字符串（一个字符）。  
ord()：将字符串（一个字符）转换成对应的编码（整数）。

![](https://m.qpic.cn/psc?/V10OQot10zcV7k/ZUXCJANDCZJwMw9eDcmXBYmvrJCiujMjgCNHeiy4cJK8tdTxdmu.Du9pu39dZijnTBJVHFrr12wAVGxp5PYc7w!!/b&bo=pgGhAAAAAAADByQ!&rf=viewer_4)
##运算符
###算术运算符
加减乘除、整除运算符、求模（求余数）运算符和求幂运算符。

![](https://m.qpic.cn/psc?/V10OQot10zcV7k/ZUXCJANDCZJwMw9eDcmXBdgTq0fySmeRqLnNfaTNA26wXos1L77bKwKbcA6vZXfr2zt1otKIOPJb2seqhq6KUg!!/b&bo=zQHSAAAAAAADBzw!&rf=viewer_4)
###赋值运算符
它的作用是将右边的值赋给左边的变量。

![](https://m.qpic.cn/psc?/V10OQot10zcV7k/ZUXCJANDCZJwMw9eDcmXBQM0EjswB9Ipezmrnd1DTPAvjn8MNCb3Nt.1k7HDrWXwMsrGSsobY2Iu*oWYI37axQ!!/b&bo=0QFfAAAAAAADB60!&rf=viewer_4)
###比较运算符和逻辑运算符

![](https://m.qpic.cn/psc?/V10OQot10zcV7k/ZUXCJANDCZJwMw9eDcmXBZLlwx8Vo8D7cVucFnBDNGnBKbFR*PgmPEW3*FRVtoKtMVQ8XJR27Echf7w0vt8Zqg!!/b&bo=uwEkAQAAAAADB70!&rf=viewer_4)
##总结
我学会用变量来保存数据，了解变量有不同的类型，变量可以做运算，可以通过内置函数来转换变量类型。学会使用运算符以及由运算符构成的表达式，可以帮助我解决很多实际的问题。




















