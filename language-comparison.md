# TypeScript vs Rust vs Python vs Go 差异详解

## 目录
1. [语法 (Syntax)](#1-语法-syntax)
2. [类型与方法 (Types & Methods)](#2-类型与方法-types--methods)
3. [逻辑与函数 (Logic & Functions)](#3-逻辑与函数-logic--functions)
4. [并发与异步 (Concurrency & Async)](#4-并发与异步-concurrency--async)
5. [IO与HTTP (IO & HTTP)](#5-io与http-io--http)
6. [迭代器 (Iterator)](#6-迭代器-iterator)

---

## 1. 语法 (Syntax)

### 1.1 变量声明

| 特性 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| 可变变量 | `let` / `var` | `let mut` | 直接赋值 | `var` / `:=` |
| 不可变变量 | `const` | `let` | 无(约定大写) | `const` |
| 类型推断 | ✅ | ✅ | ✅ (动态) | ✅ |

```typescript
// TypeScript
let name: string = "Alice";        // 可变
const age: number = 25;            // 不可变
let inferred = "hello";            // 类型推断
```

```rust
// Rust
let mut name: String = String::from("Alice");  // 可变
let age: i32 = 25;                             // 不可变
let inferred = "hello";                        // 类型推断
```

```python
# Python
name = "Alice"          # 可变（动态类型）
AGE = 25               # 约定不可变（实际可变）
name: str = "Alice"    # 类型注解（仅提示）
```

```go
// Go
var name string = "Alice"   // 显式声明
age := 25                   // 短声明（类型推断）
const PI = 3.14            // 常量
```

### 1.2 代码块与作用域

```typescript
// TypeScript - 花括号 + 分号（可选）
if (condition) {
    console.log("yes");
}
```

```rust
// Rust - 花括号，分号表示语句，无分号表示表达式返回值
if condition {
    println!("yes");
}
let x = {
    let a = 1;
    a + 2  // 无分号，作为返回值
};  // x = 3
```

```python
# Python - 缩进定义作用域，冒号开始代码块
if condition:
    print("yes")
```

```go
// Go - 花括号，左花括号必须在同一行
if condition {
    fmt.Println("yes")
}
```

### 1.3 注释

```typescript
// TypeScript
// 单行注释
/* 多行注释 */
/** JSDoc 文档注释 */
```

```rust
// Rust
// 单行注释
/* 多行注释 */
/// 文档注释（用于生成文档）
//! 模块级文档注释
```

```python
# Python
# 单行注释
"""
多行字符串/文档字符串
"""
```

```go
// Go
// 单行注释
/* 多行注释 */
// GoDoc 使用普通注释（函数/类型上方）
```

### 1.4 字符串

```typescript
// TypeScript
const s1 = "hello";
const s2 = 'hello';
const s3 = `hello ${name}`;  // 模板字符串
const multiline = `
    多行
    字符串
`;
```

```rust
// Rust
let s1 = "hello";                    // &str (字符串切片)
let s2 = String::from("hello");      // String (堆分配)
let s3 = format!("hello {}", name);  // 格式化
let raw = r#"原始字符串 \n 不转义"#;
let multiline = "多行\n字符串";
```

```python
# Python
s1 = "hello"
s2 = 'hello'
s3 = f"hello {name}"  # f-string
raw = r"原始字符串 \n 不转义"
multiline = """
多行
字符串
"""
```

```go
// Go
s1 := "hello"
s2 := fmt.Sprintf("hello %s", name)  // 格式化
raw := `原始字符串 \n 不转义
也支持多行`
```

---

## 2. 类型与方法 (Types & Methods)

### 2.1 基本类型

| 类型 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| 整数 | `number` | `i8/i16/i32/i64/i128/isize` | `int` | `int/int8/int16/int32/int64` |
| 无符号整数 | `number` | `u8/u16/u32/u64/u128/usize` | - | `uint/uint8...uint64` |
| 浮点数 | `number` | `f32/f64` | `float` | `float32/float64` |
| 布尔 | `boolean` | `bool` | `bool` | `bool` |
| 字符 | `string` | `char` | `str` | `rune` |
| 字符串 | `string` | `String/&str` | `str` | `string` |
| 空值 | `null/undefined` | `()` | `None` | `nil` |

### 2.2 复合类型 - 数组/列表

```typescript
// TypeScript
const arr: number[] = [1, 2, 3];
const arr2: Array<number> = [1, 2, 3];
const tuple: [string, number] = ["age", 25];

// 操作
arr.push(4);
arr.map(x => x * 2);
arr.filter(x => x > 1);
```

```rust
// Rust
let arr: [i32; 3] = [1, 2, 3];           // 固定大小数组
let vec: Vec<i32> = vec![1, 2, 3];       // 动态数组
let tuple: (&str, i32) = ("age", 25);

// 操作
let mut vec = vec![1, 2, 3];
vec.push(4);
vec.iter().map(|x| x * 2).collect::<Vec<_>>();
vec.iter().filter(|x| **x > 1).collect::<Vec<_>>();
```

```python
# Python
arr = [1, 2, 3]                    # list
arr_typed: list[int] = [1, 2, 3]  # 类型注解
tup = ("age", 25)                  # tuple（不可变）

# 操作
arr.append(4)
list(map(lambda x: x * 2, arr))
[x * 2 for x in arr]  # 列表推导式
[x for x in arr if x > 1]
```

```go
// Go
var arr [3]int = [3]int{1, 2, 3}   // 固定大小数组
slice := []int{1, 2, 3}            // 切片（动态）

// 操作
slice = append(slice, 4)
// Go 没有内置 map/filter，需手动循环或使用泛型库
result := make([]int, 0)
for _, x := range slice {
    if x > 1 {
        result = append(result, x*2)
    }
}
```

### 2.3 复合类型 - Map/Dict/HashMap

```typescript
// TypeScript
const map: Map<string, number> = new Map();
map.set("age", 25);
map.get("age");

// 对象字面量
const obj: { [key: string]: number } = { age: 25 };
const record: Record<string, number> = { age: 25 };
```

```rust
// Rust
use std::collections::HashMap;

let mut map: HashMap<String, i32> = HashMap::new();
map.insert(String::from("age"), 25);
map.get("age");  // 返回 Option<&i32>
```

```python
# Python
d = {"age": 25}
d: dict[str, int] = {"age": 25}

d["age"] = 25
d.get("age")      # 返回 None 如果不存在
d.get("age", 0)   # 默认值
```

```go
// Go
m := make(map[string]int)
m["age"] = 25
value, exists := m["age"]  // 双返回值检查是否存在
```

### 2.4 结构体/类/接口

```typescript
// TypeScript - Class
class Person {
    private name: string;
    public age: number;
    
    constructor(name: string, age: number) {
        this.name = name;
        this.age = age;
    }
    
    greet(): string {
        return `Hello, ${this.name}`;
    }
}

// Interface
interface Animal {
    name: string;
    speak(): void;
}
```

```rust
// Rust - Struct + impl
struct Person {
    name: String,  // 默认私有
    pub age: i32,  // 公开
}

impl Person {
    // 关联函数（类似静态方法）
    fn new(name: String, age: i32) -> Self {
        Person { name, age }
    }
    
    // 方法（&self 不可变借用）
    fn greet(&self) -> String {
        format!("Hello, {}", self.name)
    }
    
    // 可变方法（&mut self）
    fn birthday(&mut self) {
        self.age += 1;
    }
}

// Trait（类似接口）
trait Animal {
    fn speak(&self);
}
```

```python
# Python - Class
class Person:
    def __init__(self, name: str, age: int):
        self._name = name   # 约定私有（单下划线）
        self.age = age
    
    def greet(self) -> str:
        return f"Hello, {self._name}"

# Protocol（类似接口，Python 3.8+）
from typing import Protocol

class Animal(Protocol):
    name: str
    def speak(self) -> None: ...
```

```go
// Go - Struct
type Person struct {
    name string  // 小写私有
    Age  int     // 大写公开
}

// 方法（值接收者）
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, %s", p.name)
}

// 方法（指针接收者，可修改）
func (p *Person) Birthday() {
    p.Age++
}

// 构造函数（约定）
func NewPerson(name string, age int) *Person {
    return &Person{name: name, Age: age}
}

// Interface（隐式实现）
type Animal interface {
    Speak()
}
```

### 2.5 枚举

```typescript
// TypeScript
enum Color {
    Red,
    Green,
    Blue,
}

enum Status {
    Active = "active",
    Inactive = "inactive",
}

// 联合类型（更灵活）
type Result<T, E> = 
    | { ok: true; value: T }
    | { ok: false; error: E };
```

```rust
// Rust - 代数数据类型（最强大）
enum Color {
    Red,
    Green,
    Blue,
}

// 带数据的枚举
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// 标准库中的 Result 和 Option
enum Option<T> {
    Some(T),
    None,
}

enum Result<T, E> {
    Ok(T),
    Err(E),
}
```

```python
# Python
from enum import Enum, auto

class Color(Enum):
    RED = auto()
    GREEN = auto()
    BLUE = auto()

class Status(Enum):
    ACTIVE = "active"
    INACTIVE = "inactive"
```

```go
// Go - 使用 const + iota
type Color int

const (
    Red Color = iota
    Green
    Blue
)

// 没有原生枚举，通常用类型别名 + 常量
type Status string

const (
    StatusActive   Status = "active"
    StatusInactive Status = "inactive"
)
```

### 2.6 泛型

```typescript
// TypeScript
function identity<T>(arg: T): T {
    return arg;
}

interface Container<T> {
    value: T;
}

class Box<T> {
    constructor(public value: T) {}
}
```

```rust
// Rust
fn identity<T>(arg: T) -> T {
    arg
}

// 带 trait 约束
fn print_item<T: std::fmt::Display>(item: T) {
    println!("{}", item);
}

struct Container<T> {
    value: T,
}

impl<T> Container<T> {
    fn new(value: T) -> Self {
        Container { value }
    }
}
```

```python
# Python (3.12+)
from typing import TypeVar, Generic

T = TypeVar('T')

def identity(arg: T) -> T:
    return arg

class Container(Generic[T]):
    def __init__(self, value: T):
        self.value = value

# Python 3.12+ 新语法
def identity[T](arg: T) -> T:
    return arg
```

```go
// Go (1.18+)
func Identity[T any](arg T) T {
    return arg
}

// 带约束
func Print[T fmt.Stringer](item T) {
    fmt.Println(item.String())
}

type Container[T any] struct {
    Value T
}

func NewContainer[T any](value T) Container[T] {
    return Container[T]{Value: value}
}
```

---

## 3. 逻辑与函数 (Logic & Functions)

### 3.1 函数定义

```typescript
// TypeScript
function add(a: number, b: number): number {
    return a + b;
}

// 箭头函数
const add = (a: number, b: number): number => a + b;

// 可选参数与默认值
function greet(name: string, greeting: string = "Hello"): string {
    return `${greeting}, ${name}`;
}

// 剩余参数
function sum(...numbers: number[]): number {
    return numbers.reduce((a, b) => a + b, 0);
}
```

```rust
// Rust
fn add(a: i32, b: i32) -> i32 {
    a + b  // 无分号 = 返回值
}

// 闭包
let add = |a: i32, b: i32| -> i32 { a + b };
let add_short = |a, b| a + b;  // 类型推断

// 默认参数（Rust 不支持，用 Option 或 builder 模式）
fn greet(name: &str, greeting: Option<&str>) -> String {
    let g = greeting.unwrap_or("Hello");
    format!("{}, {}", g, name)
}

// 可变参数（使用宏或切片）
fn sum(numbers: &[i32]) -> i32 {
    numbers.iter().sum()
}
```

```python
# Python
def add(a: int, b: int) -> int:
    return a + b

# lambda
add = lambda a, b: a + b

# 默认参数
def greet(name: str, greeting: str = "Hello") -> str:
    return f"{greeting}, {name}"

# 可变参数
def sum_all(*numbers: int) -> int:
    return sum(numbers)

# 关键字参数
def configure(**kwargs):
    for key, value in kwargs.items():
        print(f"{key}: {value}")
```

```go
// Go
func add(a, b int) int {
    return a + b
}

// 多返回值
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 命名返回值
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return  // 裸返回
}

// 可变参数
func sumAll(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}
```

### 3.2 高阶函数与闭包

```typescript
// TypeScript
function apply(fn: (x: number) => number, value: number): number {
    return fn(value);
}

// 返回函数
function multiplier(factor: number): (x: number) => number {
    return (x) => x * factor;  // 闭包捕获 factor
}

const double = multiplier(2);
console.log(double(5));  // 10
```

```rust
// Rust - 闭包类型：Fn, FnMut, FnOnce
fn apply<F>(f: F, value: i32) -> i32
where
    F: Fn(i32) -> i32,
{
    f(value)
}

// 返回闭包（需要 Box 或 impl Trait）
fn multiplier(factor: i32) -> impl Fn(i32) -> i32 {
    move |x| x * factor  // move 转移所有权
}

let double = multiplier(2);
println!("{}", double(5));  // 10
```

```python
# Python
def apply(fn, value):
    return fn(value)

# 返回函数
def multiplier(factor):
    def inner(x):
        return x * factor  # 闭包
    return inner

double = multiplier(2)
print(double(5))  # 10

# 装饰器
def logged(fn):
    def wrapper(*args, **kwargs):
        print(f"Calling {fn.__name__}")
        return fn(*args, **kwargs)
    return wrapper

@logged
def greet(name):
    return f"Hello, {name}"
```

```go
// Go
func apply(fn func(int) int, value int) int {
    return fn(value)
}

// 返回函数
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor  // 闭包
    }
}

double := multiplier(2)
fmt.Println(double(5))  // 10
```

### 3.3 条件语句

```typescript
// TypeScript
if (x > 0) {
    console.log("positive");
} else if (x < 0) {
    console.log("negative");
} else {
    console.log("zero");
}

// 三元运算符
const result = x > 0 ? "positive" : "non-positive";

// switch
switch (color) {
    case "red":
        console.log("Red");
        break;
    case "blue":
    case "green":
        console.log("Not red");
        break;
    default:
        console.log("Unknown");
}
```

```rust
// Rust - if 是表达式
let result = if x > 0 {
    "positive"
} else if x < 0 {
    "negative"
} else {
    "zero"
};

// match（模式匹配，必须穷尽）
match color {
    Color::Red => println!("Red"),
    Color::Blue | Color::Green => println!("Not red"),
    _ => println!("Unknown"),
}

// match 带守卫
match x {
    n if n < 0 => println!("negative"),
    0 => println!("zero"),
    n if n > 0 => println!("positive"),
    _ => unreachable!(),
}

// if let（简化的模式匹配）
if let Some(value) = optional {
    println!("Got: {}", value);
}
```

```python
# Python
if x > 0:
    print("positive")
elif x < 0:
    print("negative")
else:
    print("zero")

# 条件表达式
result = "positive" if x > 0 else "non-positive"

# match (Python 3.10+)
match color:
    case "red":
        print("Red")
    case "blue" | "green":
        print("Not red")
    case _:
        print("Unknown")
```

```go
// Go
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// switch（无需 break）
switch color {
case "red":
    fmt.Println("Red")
case "blue", "green":
    fmt.Println("Not red")
default:
    fmt.Println("Unknown")
}

// 类型 switch
switch v := value.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown type\n")
}
```

### 3.4 循环

```typescript
// TypeScript
// for 循环
for (let i = 0; i < 10; i++) {
    console.log(i);
}

// for...of（迭代值）
for (const item of array) {
    console.log(item);
}

// for...in（迭代键）
for (const key in object) {
    console.log(key);
}

// while
while (condition) {
    // ...
}

// 数组方法
array.forEach((item, index) => console.log(item));
```

```rust
// Rust
// loop（无限循环，可返回值）
let result = loop {
    if condition {
        break value;
    }
};

// while
while condition {
    // ...
}

// for（迭代器）
for i in 0..10 {  // Range: 0 到 9
    println!("{}", i);
}

for i in 0..=10 {  // 包含 10
    println!("{}", i);
}

for item in &vec {  // 借用
    println!("{}", item);
}

for (index, item) in vec.iter().enumerate() {
    println!("{}: {}", index, item);
}
```

```python
# Python
# for（迭代）
for i in range(10):
    print(i)

for item in array:
    print(item)

for index, item in enumerate(array):
    print(index, item)

for key, value in dictionary.items():
    print(key, value)

# while
while condition:
    pass

# 列表推导式
squares = [x**2 for x in range(10)]
evens = [x for x in range(10) if x % 2 == 0]
```

```go
// Go - 只有 for
// 传统 for
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// while 风格
for condition {
    // ...
}

// 无限循环
for {
    if condition {
        break
    }
}

// range 迭代
for index, value := range slice {
    fmt.Println(index, value)
}

for key, value := range m {
    fmt.Println(key, value)
}
```

### 3.5 错误处理

```typescript
// TypeScript - try/catch
try {
    const result = riskyOperation();
} catch (error) {
    if (error instanceof Error) {
        console.error(error.message);
    }
} finally {
    cleanup();
}

// 或使用 Result 模式
type Result<T, E> = { ok: true; value: T } | { ok: false; error: E };
```

```rust
// Rust - Result 和 Option（无异常）
fn divide(a: i32, b: i32) -> Result<i32, String> {
    if b == 0 {
        Err(String::from("Division by zero"))
    } else {
        Ok(a / b)
    }
}

// 使用
match divide(10, 2) {
    Ok(result) => println!("Result: {}", result),
    Err(e) => println!("Error: {}", e),
}

// ? 操作符（错误传播）
fn process() -> Result<i32, String> {
    let x = divide(10, 2)?;  // 错误时提前返回
    let y = divide(x, 3)?;
    Ok(y)
}

// unwrap 和 expect（panic 如果是 Err/None）
let value = some_option.unwrap();
let value = some_option.expect("Should have value");

// 组合方法
let result = some_option
    .map(|x| x * 2)
    .unwrap_or(0);
```

```python
# Python - try/except
try:
    result = risky_operation()
except ValueError as e:
    print(f"Value error: {e}")
except Exception as e:
    print(f"Error: {e}")
else:
    print("Success")
finally:
    cleanup()

# 抛出异常
def divide(a, b):
    if b == 0:
        raise ValueError("Division by zero")
    return a / b
```

```go
// Go - 多返回值（无异常）
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 使用
result, err := divide(10, 2)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println("Result:", result)

// 自定义错误
type MyError struct {
    Code    int
    Message string
}

func (e *MyError) Error() string {
    return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

// panic/recover（少用）
func riskyOperation() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered:", r)
        }
    }()
    panic("something went wrong")
}
```

---

## 4. 并发与异步 (Concurrency & Async)

### 4.1 并发模型对比

| 特性 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| 模型 | 事件循环 | 系统线程 + async | GIL + async | Goroutines |
| 并行 | ❌ (单线程) | ✅ | ❌ (GIL限制) | ✅ |
| 异步关键字 | async/await | async/await | async/await | - (内置) |
| 通信机制 | Promise | Channel/Mutex | Queue/Lock | Channel |

### 4.2 异步函数

```typescript
// TypeScript
async function fetchData(url: string): Promise<string> {
    const response = await fetch(url);
    const data = await response.text();
    return data;
}

// Promise 组合
const results = await Promise.all([
    fetchData(url1),
    fetchData(url2),
]);

// Promise.race
const fastest = await Promise.race([promise1, promise2]);
```

```rust
// Rust（需要运行时如 tokio）
async fn fetch_data(url: &str) -> Result<String, reqwest::Error> {
    let response = reqwest::get(url).await?;
    let data = response.text().await?;
    Ok(data)
}

// 并发执行
use tokio::join;
let (result1, result2) = join!(
    fetch_data(url1),
    fetch_data(url2)
);

// select（竞争）
use tokio::select;
select! {
    result = future1 => { /* ... */ }
    result = future2 => { /* ... */ }
}
```

```python
# Python
import asyncio

async def fetch_data(url: str) -> str:
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.text()

# 并发执行
results = await asyncio.gather(
    fetch_data(url1),
    fetch_data(url2)
)

# 竞争
done, pending = await asyncio.wait(
    [task1, task2],
    return_when=asyncio.FIRST_COMPLETED
)
```

```go
// Go - 无需 async/await
func fetchData(url string) (string, error) {
    resp, err := http.Get(url)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()
    body, err := io.ReadAll(resp.Body)
    return string(body), err
}

// 启动 goroutine
go func() {
    result, err := fetchData(url)
    // ...
}()
```

### 4.3 线程/协程创建

```typescript
// TypeScript (Node.js Worker Threads)
const { Worker } = require('worker_threads');

const worker = new Worker('./worker.js', {
    workerData: { /* ... */ }
});

worker.on('message', (result) => {
    console.log('Result:', result);
});
```

```rust
// Rust - 系统线程
use std::thread;

let handle = thread::spawn(|| {
    println!("Hello from thread!");
    42  // 返回值
});

let result = handle.join().unwrap();  // 等待完成

// tokio spawn（异步任务）
let handle = tokio::spawn(async {
    fetch_data("http://example.com").await
});
let result = handle.await?;
```

```python
# Python - 线程
from threading import Thread

def worker():
    print("Hello from thread!")

thread = Thread(target=worker)
thread.start()
thread.join()

# 异步任务
async def main():
    task = asyncio.create_task(async_worker())
    await task

# 进程池（绕过 GIL）
from concurrent.futures import ProcessPoolExecutor

with ProcessPoolExecutor() as executor:
    results = executor.map(cpu_bound_fn, data)
```

```go
// Go - Goroutine（轻量级）
go worker()  // 启动 goroutine

// 等待完成需要 sync.WaitGroup
var wg sync.WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    worker()
}()
wg.Wait()
```

### 4.4 通道/消息传递

```typescript
// TypeScript - 无内置 channel，使用库或自定义
// 示例：简单的 async queue
class Channel<T> {
    private queue: T[] = [];
    private resolvers: ((value: T) => void)[] = [];

    send(value: T) {
        const resolver = this.resolvers.shift();
        if (resolver) {
            resolver(value);
        } else {
            this.queue.push(value);
        }
    }

    async receive(): Promise<T> {
        const value = this.queue.shift();
        if (value !== undefined) {
            return value;
        }
        return new Promise(resolve => this.resolvers.push(resolve));
    }
}
```

```rust
// Rust - std::sync::mpsc (多生产者单消费者)
use std::sync::mpsc;
use std::thread;

let (tx, rx) = mpsc::channel();

thread::spawn(move || {
    tx.send("Hello").unwrap();
});

let received = rx.recv().unwrap();
println!("Got: {}", received);

// tokio channel（异步）
use tokio::sync::mpsc;

let (tx, mut rx) = mpsc::channel(32);

tokio::spawn(async move {
    tx.send("Hello").await.unwrap();
});

while let Some(msg) = rx.recv().await {
    println!("Got: {}", msg);
}
```

```python
# Python - asyncio.Queue
import asyncio

async def producer(queue: asyncio.Queue):
    await queue.put("Hello")

async def consumer(queue: asyncio.Queue):
    msg = await queue.get()
    print(f"Got: {msg}")

queue = asyncio.Queue()
await asyncio.gather(
    producer(queue),
    consumer(queue)
)

# 线程间：queue.Queue
from queue import Queue

q = Queue()
q.put("Hello")
msg = q.get()
```

```go
// Go - Channel（内置）
ch := make(chan string)  // 无缓冲
ch := make(chan string, 10)  // 带缓冲

// 发送和接收
go func() {
    ch <- "Hello"  // 发送
}()

msg := <-ch  // 接收
fmt.Println("Got:", msg)

// select（多路复用）
select {
case msg := <-ch1:
    fmt.Println("From ch1:", msg)
case msg := <-ch2:
    fmt.Println("From ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
default:
    fmt.Println("No message")
}

// 关闭 channel
close(ch)
for msg := range ch {  // 迭代直到关闭
    fmt.Println(msg)
}
```

### 4.5 互斥锁与共享状态

```typescript
// TypeScript - 单线程，通常不需要锁
// Worker 间通过 SharedArrayBuffer 和 Atomics

const sab = new SharedArrayBuffer(1024);
const int32 = new Int32Array(sab);
Atomics.add(int32, 0, 1);
```

```rust
// Rust - Mutex 和 Arc
use std::sync::{Arc, Mutex};
use std::thread;

let counter = Arc::new(Mutex::new(0));
let mut handles = vec![];

for _ in 0..10 {
    let counter = Arc::clone(&counter);
    let handle = thread::spawn(move || {
        let mut num = counter.lock().unwrap();
        *num += 1;
    });
    handles.push(handle);
}

for handle in handles {
    handle.join().unwrap();
}

println!("Result: {}", *counter.lock().unwrap());

// RwLock（读写锁）
use std::sync::RwLock;

let lock = RwLock::new(5);
{
    let r = lock.read().unwrap();  // 多个读取者
    println!("{}", *r);
}
{
    let mut w = lock.write().unwrap();  // 独占写入
    *w += 1;
}
```

```python
# Python
from threading import Lock, RLock

lock = Lock()

def safe_increment():
    with lock:  # 自动获取和释放
        global counter
        counter += 1

# asyncio.Lock
lock = asyncio.Lock()

async def safe_operation():
    async with lock:
        await do_something()
```

```go
// Go - sync.Mutex
import "sync"

var (
    counter int
    mu      sync.Mutex
)

func increment() {
    mu.Lock()
    defer mu.Unlock()
    counter++
}

// RWMutex
var rw sync.RWMutex

func read() int {
    rw.RLock()
    defer rw.RUnlock()
    return counter
}

func write(val int) {
    rw.Lock()
    defer rw.Unlock()
    counter = val
}

// sync.Once（单次执行）
var once sync.Once
once.Do(func() {
    // 只执行一次
})
```

---

## 5. IO与HTTP (IO & HTTP)

### 5.1 文件读写

```typescript
// TypeScript (Node.js)
import { readFile, writeFile } from 'fs/promises';
import { createReadStream, createWriteStream } from 'fs';

// 读取整个文件
const content = await readFile('file.txt', 'utf-8');

// 写入文件
await writeFile('output.txt', 'Hello, World!');

// 流式读取
const stream = createReadStream('large-file.txt');
for await (const chunk of stream) {
    console.log(chunk.toString());
}

// 流式写入
const writeStream = createWriteStream('output.txt');
writeStream.write('Hello');
writeStream.end();
```

```rust
// Rust
use std::fs;
use std::io::{BufRead, BufReader, Write};

// 读取整个文件
let content = fs::read_to_string("file.txt")?;

// 写入文件
fs::write("output.txt", "Hello, World!")?;

// 逐行读取
let file = fs::File::open("file.txt")?;
let reader = BufReader::new(file);
for line in reader.lines() {
    println!("{}", line?);
}

// 异步（tokio）
use tokio::fs;
use tokio::io::{AsyncBufReadExt, AsyncWriteExt, BufReader};

let content = fs::read_to_string("file.txt").await?;
fs::write("output.txt", "Hello").await?;
```

```python
# Python
# 读取整个文件
with open('file.txt', 'r') as f:
    content = f.read()

# 写入文件
with open('output.txt', 'w') as f:
    f.write('Hello, World!')

# 逐行读取
with open('file.txt', 'r') as f:
    for line in f:
        print(line.strip())

# 异步（aiofiles）
import aiofiles

async with aiofiles.open('file.txt', 'r') as f:
    content = await f.read()
```

```go
// Go
import (
    "bufio"
    "io"
    "os"
)

// 读取整个文件
content, err := os.ReadFile("file.txt")

// 写入文件
err := os.WriteFile("output.txt", []byte("Hello, World!"), 0644)

// 逐行读取
file, _ := os.Open("file.txt")
defer file.Close()
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

// 流式复制
src, _ := os.Open("source.txt")
dst, _ := os.Create("dest.txt")
io.Copy(dst, src)
```

### 5.2 HTTP 客户端

```typescript
// TypeScript
// fetch (内置于 Node.js 18+)
const response = await fetch('https://api.example.com/data', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name: 'Alice' }),
});
const data = await response.json();

// axios（流行库）
import axios from 'axios';

const response = await axios.get('https://api.example.com/data');
const data = response.data;

await axios.post('https://api.example.com/data', {
    name: 'Alice'
});
```

```rust
// Rust (reqwest)
use reqwest;

// GET 请求
let body = reqwest::get("https://api.example.com/data")
    .await?
    .text()
    .await?;

// POST JSON
let client = reqwest::Client::new();
let response = client
    .post("https://api.example.com/data")
    .json(&serde_json::json!({"name": "Alice"}))
    .send()
    .await?;

let data: serde_json::Value = response.json().await?;

// 带超时和 header
let client = reqwest::Client::builder()
    .timeout(std::time::Duration::from_secs(10))
    .build()?;

let response = client
    .get("https://api.example.com/data")
    .header("Authorization", "Bearer token")
    .send()
    .await?;
```

```python
# Python
import requests  # 同步
import aiohttp   # 异步

# requests (同步)
response = requests.get('https://api.example.com/data')
data = response.json()

response = requests.post(
    'https://api.example.com/data',
    json={'name': 'Alice'},
    headers={'Authorization': 'Bearer token'}
)

# aiohttp (异步)
async with aiohttp.ClientSession() as session:
    async with session.get('https://api.example.com/data') as resp:
        data = await resp.json()
    
    async with session.post(
        'https://api.example.com/data',
        json={'name': 'Alice'}
    ) as resp:
        result = await resp.json()
```

```go
// Go (net/http)
import (
    "bytes"
    "encoding/json"
    "net/http"
    "io"
)

// GET 请求
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    return err
}
defer resp.Body.Close()
body, _ := io.ReadAll(resp.Body)

// POST JSON
data := map[string]string{"name": "Alice"}
jsonData, _ := json.Marshal(data)

resp, err := http.Post(
    "https://api.example.com/data",
    "application/json",
    bytes.NewBuffer(jsonData),
)

// 自定义请求
client := &http.Client{
    Timeout: 10 * time.Second,
}

req, _ := http.NewRequest("GET", "https://api.example.com/data", nil)
req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)
```

### 5.3 HTTP 服务端

```typescript
// TypeScript (Express)
import express from 'express';

const app = express();
app.use(express.json());

app.get('/api/users/:id', (req, res) => {
    const { id } = req.params;
    res.json({ id, name: 'Alice' });
});

app.post('/api/users', (req, res) => {
    const { name } = req.body;
    res.status(201).json({ id: 1, name });
});

app.listen(3000, () => {
    console.log('Server running on port 3000');
});
```

```rust
// Rust (Axum)
use axum::{
    routing::{get, post},
    Router, Json, extract::Path,
};
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
struct User {
    id: u32,
    name: String,
}

#[derive(Deserialize)]
struct CreateUser {
    name: String,
}

async fn get_user(Path(id): Path<u32>) -> Json<User> {
    Json(User { id, name: "Alice".into() })
}

async fn create_user(Json(payload): Json<CreateUser>) -> Json<User> {
    Json(User { id: 1, name: payload.name })
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/api/users/:id", get(get_user))
        .route("/api/users", post(create_user));

    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
```

```python
# Python (FastAPI)
from fastapi import FastAPI
from pydantic import BaseModel

app = FastAPI()

class User(BaseModel):
    id: int
    name: str

class CreateUser(BaseModel):
    name: str

@app.get("/api/users/{user_id}")
async def get_user(user_id: int) -> User:
    return User(id=user_id, name="Alice")

@app.post("/api/users", status_code=201)
async def create_user(user: CreateUser) -> User:
    return User(id=1, name=user.name)

# 运行: uvicorn main:app --reload
```

```go
// Go (标准库 net/http)
package main

import (
    "encoding/json"
    "net/http"
)

type User struct {
    ID   int    `json:"id"`
    Name string `json:"name"`
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.PathValue("id")  // Go 1.22+
    user := User{ID: 1, Name: "Alice"}
    
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
    var input struct {
        Name string `json:"name"`
    }
    json.NewDecoder(r.Body).Decode(&input)
    
    user := User{ID: 1, Name: input.Name}
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("GET /api/users/{id}", getUser)
    mux.HandleFunc("POST /api/users", createUser)
    
    http.ListenAndServe(":3000", mux)
}
```

### 5.4 JSON 处理

```typescript
// TypeScript
interface User {
    id: number;
    name: string;
}

// 解析
const user: User = JSON.parse('{"id": 1, "name": "Alice"}');

// 序列化
const json = JSON.stringify({ id: 1, name: "Alice" });
const pretty = JSON.stringify({ id: 1 }, null, 2);
```

```rust
// Rust (serde)
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize)]
struct User {
    id: i32,
    name: String,
    #[serde(skip_serializing_if = "Option::is_none")]
    email: Option<String>,
}

// 解析
let user: User = serde_json::from_str(r#"{"id": 1, "name": "Alice"}"#)?;

// 序列化
let json = serde_json::to_string(&user)?;
let pretty = serde_json::to_string_pretty(&user)?;

// 动态 JSON
let value: serde_json::Value = serde_json::json!({
    "id": 1,
    "name": "Alice"
});
```

```python
# Python
import json
from dataclasses import dataclass
from pydantic import BaseModel

# 内置 json
data = json.loads('{"id": 1, "name": "Alice"}')
json_str = json.dumps({"id": 1, "name": "Alice"})
pretty = json.dumps({"id": 1}, indent=2)

# Pydantic（类型安全）
class User(BaseModel):
    id: int
    name: str
    email: str | None = None

user = User.model_validate_json('{"id": 1, "name": "Alice"}')
json_str = user.model_dump_json()
```

```go
// Go
import "encoding/json"

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email,omitempty"`
}

// 解析
var user User
json.Unmarshal([]byte(`{"id": 1, "name": "Alice"}`), &user)

// 或从 Reader
json.NewDecoder(reader).Decode(&user)

// 序列化
jsonBytes, _ := json.Marshal(user)
prettyBytes, _ := json.MarshalIndent(user, "", "  ")

// 动态 JSON
var data map[string]interface{}
json.Unmarshal(jsonBytes, &data)
```

---

## 总结对比表

| 特性 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| **类型系统** | 静态（可选） | 静态（强） | 动态 | 静态（简单） |
| **内存管理** | GC | 所有权 | GC | GC |
| **空安全** | `strictNullChecks` | `Option<T>` | `None` | `nil` |
| **错误处理** | try/catch | `Result<T, E>` | try/except | 多返回值 |
| **并发模型** | 事件循环 | 线程/async | GIL/async | Goroutine |
| **学习曲线** | 中 | 高 | 低 | 低-中 |
| **编译速度** | 快 | 慢 | 无（解释） | 快 |
| **运行性能** | 中 | 高 | 低 | 高 |
| **生态系统** | npm | crates.io | PyPI | 标准库为主 |
| **适用场景** | Web前后端 | 系统/性能关键 | 脚本/ML/Web | 服务端/云原生 |

### 选择建议

- **TypeScript**: 全栈 Web 开发，前端为主
- **Rust**: 系统编程、性能敏感、安全关键
- **Python**: 快速原型、数据科学、机器学习、脚本
- **Go**: 微服务、云原生、网络服务、DevOps 工具

---

## 6. 迭代器 (Iterator)

### 6.1 迭代器协议对比

| 语言 | 协议/Trait | 核心方法 | 返回值 |
|------|-----------|---------|--------|
| TypeScript | `Iterable`/`Iterator` | `next()` | `{value, done}` |
| Rust | `Iterator` trait | `next()` | `Option<Item>` |
| Python | `__iter__`/`__next__` | `__next__()` | 值或 `StopIteration` |
| Go | 无内置协议 | `range` 关键字 | - |

### 6.2 基本迭代

```typescript
// TypeScript
const arr = [1, 2, 3];

// for...of（推荐）
for (const item of arr) {
    console.log(item);
}

// 手动迭代
const iter = arr[Symbol.iterator]();
let result = iter.next();
while (!result.done) {
    console.log(result.value);
    result = iter.next();
}
```

```rust
// Rust
let arr = vec![1, 2, 3];

// for 循环（自动调用 into_iter）
for item in &arr {
    println!("{}", item);
}

// 手动迭代
let mut iter = arr.iter();
while let Some(item) = iter.next() {
    println!("{}", item);
}

// 三种迭代方式
for item in arr.iter() {}      // &T（借用）
for item in arr.iter_mut() {}  // &mut T（可变借用）
for item in arr.into_iter() {} // T（获取所有权）
```

```python
# Python
arr = [1, 2, 3]

# for 循环
for item in arr:
    print(item)

# 手动迭代
it = iter(arr)
while True:
    try:
        item = next(it)
        print(item)
    except StopIteration:
        break

# next 带默认值
next(it, None)  # 迭代完返回 None 而非异常
```

```go
// Go
arr := []int{1, 2, 3}

// range（唯一方式）
for index, value := range arr {
    fmt.Println(index, value)
}

// 忽略索引
for _, value := range arr {
    fmt.Println(value)
}

// 只要索引
for index := range arr {
    fmt.Println(index)
}
```

### 6.3 创建迭代器

```typescript
// TypeScript
const arr = [1, 2, 3];
const set = new Set([1, 2, 3]);
const map = new Map([['a', 1], ['b', 2]]);

// 获取迭代器
arr[Symbol.iterator]();
arr.values();          // 值
arr.keys();            // 索引
arr.entries();         // [index, value]

set.values();
map.keys();
map.values();
map.entries();

// Object 迭代
Object.keys(obj);      // string[]
Object.values(obj);    // any[]
Object.entries(obj);   // [string, any][]
```

```rust
// Rust
let vec = vec![1, 2, 3];
let set: HashSet<i32> = [1, 2, 3].into();
let map: HashMap<&str, i32> = [("a", 1), ("b", 2)].into();

// Vec 迭代器
vec.iter();            // &T
vec.iter_mut();        // &mut T
vec.into_iter();       // T（消耗 vec）

// 集合迭代
set.iter();
map.keys();
map.values();
map.iter();            // (&K, &V)

// 范围迭代器
0..10                  // Range: 0-9
0..=10                 // RangeInclusive: 0-10
(0..).take(10)         // 无限范围取前10个
```

```python
# Python
arr = [1, 2, 3]
s = {1, 2, 3}
d = {'a': 1, 'b': 2}

# 列表
iter(arr)
enumerate(arr)         # (index, value)
reversed(arr)          # 反向

# 字典
d.keys()
d.values()
d.items()              # (key, value)
iter(d)                # 等同于 d.keys()

# 范围
range(10)              # 0-9
range(1, 10)           # 1-9
range(0, 10, 2)        # 0, 2, 4, 6, 8
```

```go
// Go - 没有独立迭代器，只有 range

arr := []int{1, 2, 3}
m := map[string]int{"a": 1, "b": 2}

// 切片
for i, v := range arr {}

// Map
for k, v := range m {}

// 字符串（rune 迭代）
for i, r := range "hello" {}

// Channel
for v := range ch {}  // 直到 channel 关闭
```

### 6.4 Map/Filter/Reduce

```typescript
// TypeScript
const nums = [1, 2, 3, 4, 5];

// map（立即求值）
const doubled = nums.map(x => x * 2);  // [2, 4, 6, 8, 10]

// filter
const evens = nums.filter(x => x % 2 === 0);  // [2, 4]

// reduce
const sum = nums.reduce((acc, x) => acc + x, 0);  // 15

// 链式
nums.filter(x => x % 2 === 0).map(x => x * 2);  // [4, 8]
```

```rust
// Rust（惰性求值）
let nums = vec![1, 2, 3, 4, 5];

// map
let doubled: Vec<i32> = nums.iter().map(|x| x * 2).collect();

// filter
let evens: Vec<i32> = nums.iter().filter(|x| *x % 2 == 0).copied().collect();

// fold（reduce）
let sum: i32 = nums.iter().fold(0, |acc, x| acc + x);
let sum: i32 = nums.iter().sum();  // 特化方法

// 链式（惰性，只在 collect 时执行）
let result: Vec<i32> = nums.iter()
    .filter(|x| *x % 2 == 0)
    .map(|x| x * 2)
    .collect();  // [4, 8]
```

```python
# Python
nums = [1, 2, 3, 4, 5]

# map（惰性）
doubled = list(map(lambda x: x * 2, nums))

# filter（惰性）
evens = list(filter(lambda x: x % 2 == 0, nums))

# reduce
from functools import reduce
total = reduce(lambda acc, x: acc + x, nums, 0)
total = sum(nums)  # 内置

# 列表推导式（更 Pythonic）
doubled = [x * 2 for x in nums]
evens = [x for x in nums if x % 2 == 0]
result = [x * 2 for x in nums if x % 2 == 0]  # [4, 8]
```

```go
// Go - 无内置，手动或用泛型

nums := []int{1, 2, 3, 4, 5}

// 手动 map
doubled := make([]int, len(nums))
for i, v := range nums {
    doubled[i] = v * 2
}

// 手动 filter
var evens []int
for _, v := range nums {
    if v%2 == 0 {
        evens = append(evens, v)
    }
}

// 手动 reduce
sum := 0
for _, v := range nums {
    sum += v
}

// 泛型函数
func Map[T, U any](s []T, f func(T) U) []U {
    r := make([]U, len(s))
    for i, v := range s {
        r[i] = f(v)
    }
    return r
}
```

### 6.5 其他常用方法

```typescript
// TypeScript
nums.find(x => x > 3);           // 4
nums.findIndex(x => x > 3);      // 3
nums.includes(3);                // true
nums.some(x => x > 3);           // true
nums.every(x => x > 0);          // true
nums.slice(0, 3);                // take 3
[[1, 2], [3, 4]].flat();         // [1, 2, 3, 4]
[...new Set(nums)];              // 去重
```

```rust
// Rust
nums.iter().find(|&&x| x > 3);         // Some(&4)
nums.iter().position(|&x| x > 3);      // Some(3)
nums.contains(&3);                      // true
nums.iter().any(|&x| x > 3);           // true
nums.iter().all(|&x| x > 0);           // true
nums.iter().take(3);                   // 前3个（惰性）
nums.iter().skip(2);                   // 跳过前2个
nested.into_iter().flatten();          // 展平
nums.iter().max();                     // 最大值
nums.iter().count();                   // 计数
```

```python
# Python
next((x for x in nums if x > 3), None)  # 4
3 in nums                               # True
any(x > 3 for x in nums)               # True
all(x > 0 for x in nums)               # True
nums[:3]                               # take 3
nums[2:]                               # skip 2
from itertools import chain
list(chain.from_iterable(nested))      # 展平
max(nums)                              # 最大值
len(nums)                              # 计数
```

```go
// Go (slices 包, Go 1.21+)
import "slices"

slices.Contains(nums, 3)               // true
slices.Index(nums, 3)                  // 2
slices.Max(nums)                       // 最大值
slices.Min(nums)                       // 最小值
slices.Sort(nums)                      // 排序
slices.Reverse(nums)                   // 反转
```

### 6.6 惰性求值对比

| 语言 | 惰性支持 | 说明 |
|------|---------|------|
| **Rust** | ✅ 默认惰性 | 必须 `.collect()` 触发 |
| **Python** | ✅ map/filter/生成器 | 列表推导式是立即的 |
| **TypeScript** | ❌ 数组方法立即 | 需用生成器实现惰性 |
| **Go** | ❌ | 可用 channel 模拟 |

### 6.7 自定义迭代器

```typescript
// TypeScript - 实现 Iterable 接口
class Range {
    constructor(private start: number, private end: number) {}

    [Symbol.iterator](): Iterator<number> {
        let current = this.start;
        const end = this.end;
        return {
            next(): IteratorResult<number> {
                if (current < end) {
                    return { value: current++, done: false };
                }
                return { value: undefined, done: true };
            }
        };
    }
}

for (const n of new Range(0, 5)) {
    console.log(n);  // 0, 1, 2, 3, 4
}
```

```rust
// Rust - 实现 Iterator trait
struct Counter {
    current: u32,
    max: u32,
}

impl Iterator for Counter {
    type Item = u32;

    fn next(&mut self) -> Option<Self::Item> {
        if self.current < self.max {
            let result = self.current;
            self.current += 1;
            Some(result)
        } else {
            None
        }
    }
}

// 自动获得所有迭代器方法！
let sum: u32 = Counter { current: 0, max: 10 }
    .filter(|&x| x % 2 == 0)
    .sum();  // 20
```

```python
# Python - 实现 __iter__ 和 __next__
class Counter:
    def __init__(self, max_val: int):
        self.current = 0
        self.max = max_val
    
    def __iter__(self):
        return self
    
    def __next__(self):
        if self.current < self.max:
            result = self.current
            self.current += 1
            return result
        raise StopIteration

# 更简单：使用生成器
def counter(max_val: int):
    for i in range(max_val):
        yield i
```

```go
// Go - 无协议，常用 channel 模式
func Counter(max int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < max; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

for n := range Counter(5) {
    fmt.Println(n)
}
```

### 6.8 生成器

```typescript
// TypeScript
function* range(start: number, end: number): Generator<number> {
    for (let i = start; i < end; i++) {
        yield i;
    }
}

for (const n of range(0, 5)) {
    console.log(n);
}
```

```rust
// Rust - 使用 std::iter::from_fn
fn naturals() -> impl Iterator<Item = i32> {
    let mut n = 0;
    std::iter::from_fn(move || {
        let result = n;
        n += 1;
        Some(result)
    })
}
```

```python
# Python
def range_gen(start: int, end: int):
    i = start
    while i < end:
        yield i
        i += 1

# 异步生成器
async def async_range(start: int, end: int):
    for i in range(start, end):
        await asyncio.sleep(0.1)
        yield i
```

```go
// Go - 用 goroutine + channel 模拟
func Range(start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        for i := start; i < end; i++ {
            ch <- i
        }
    }()
    return ch
}
```

### 6.9 并行迭代

```typescript
// TypeScript
const results = await Promise.all(urls.map(url => fetch(url)));
```

```rust
// Rust (rayon)
use rayon::prelude::*;

let doubled: Vec<i32> = nums.par_iter().map(|x| x * 2).collect();
let sum: i32 = nums.par_iter().sum();
```

```python
# Python
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# 线程池（IO 密集型）
with ThreadPoolExecutor(max_workers=10) as executor:
    results = list(executor.map(fetch, urls))

# 进程池（CPU 密集型）
with ProcessPoolExecutor() as executor:
    results = list(executor.map(cpu_task, data))

# asyncio
results = await asyncio.gather(*[fetch(url) for url in urls])
```

```go
// Go
func parallelMap(items []int, fn func(int) int) []int {
    results := make([]int, len(items))
    var wg sync.WaitGroup
    
    for i, item := range items {
        wg.Add(1)
        go func(i int, item int) {
            defer wg.Done()
            results[i] = fn(item)
        }(i, item)
    }
    
    wg.Wait()
    return results
}
```

### 迭代器总结

| 特性 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| **迭代器协议** | `Symbol.iterator` | `Iterator` trait | `__iter__/__next__` | 无（用 range） |
| **惰性求值** | ❌ (数组方法) | ✅ 默认惰性 | ✅ map/filter | ❌ |
| **链式调用** | ✅ 数组方法 | ✅ 丰富 | ⚠️ 需要库 | ⚠️ 需要库 |
| **生成器** | ✅ `function*` | ⚠️ nightly | ✅ `yield` | ❌ (用 channel) |
| **并行迭代** | Promise.all | rayon | multiprocessing | goroutine |
| **零成本抽象** | ❌ | ✅ | ❌ | ⚠️ |
