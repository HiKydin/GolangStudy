# GolangStudy
学习go语言的代码和笔记

## 一、变量

### 1.1输出变量类型的方法  

~~~go
var a int =10
fmt.Printf("type of a=%T",a)
~~~

### 1.2局部变量的声明

~~~go
//方法一：声明一个变量，默认值为0
var a int
//方法二：声明一个变量，初始化一个值
var a int = 10
//方法三：初始化的时候，省去数据类型，通过值自动匹配
var a = 10
//方法四：使用冒等（常用）
a := 10
//注意：方法四不支持声明 全局变量
~~~

### 1.3多变量的声明

~~~go
//单行写法
var xx,yy int = 100,200
var kk,ll = 100,"Aceld"
//多行写法
var(
    vv int = 100
    jj bool = true 
)
~~~

### 1.4常量与iota

使用const定义常量，常量是只读的，不允许修改

~~~go
const a int =10
const(
    a = 10
    b = 20
)
~~~

const可以用来定义枚举

~~~go
const{
    BEIJING = 0
    SHANGHAI = 1
    NANJING = 2
}
~~~

const还可以和**iota**一起使用来定义**有规则地枚举**

~~~go
const(
    //第一行地iota默认值是0 逐行++
    BEIJING = iota   //iota = 0
    SHANGHAI         //iota = 1
    NANJING          //iota = 2
)
~~~

### 1.5Golang的变量

Golang支持的变量有：int、int8、int16、int32、float32、float64、string、byte

## 二、循环语句

### 1.for循环

Golang中的for循环有三种写法

~~~go
for init;condition;post{}
for condition{}
for {}
~~~

且看例子

~~~go
num := [6]int{1,2,3}
//方法一
for i := 0; i < len(num); i++ {
    fmt.Println(num[i])
} 
//方法二
i := 0
for i < len(num) {
    fmt.Println(num[i])
    i++
}
//方法三
for i,x := range num {
    fmt.Printf("index：%d,value：%d\n",i,x)
}
//无限循环
for {
    fmt.Println("endless...")
}
~~~

