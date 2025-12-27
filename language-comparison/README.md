# TypeScript vs Python vs Go vs Rust 全面对比指南

> 四种现代编程语言的深度对比与实战示例

## 📊 语言概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **类型系统** | 静态 (可选) | 动态 (可选类型提示) | 静态 | 静态 |
| **编译/解释** | 编译到JS | 解释执行 | 编译 | 编译 |
| **内存管理** | GC (V8) | GC | GC | 所有权系统 |
| **并发模型** | 事件循环 | GIL/多进程 | Goroutine | async/线程 |
| **范式** | 多范式 | 多范式 | 过程式+并发 | 多范式 |
| **性能** | 中等 | 较慢 | 快 | 最快 |
| **学习曲线** | 中等 | 低 | 低-中 | 高 |

---

## 1️⃣ 基础语法对比

### 1.1 变量声明

```typescript
// TypeScript
let name: string = "Alice";      // 可变变量
const age: number = 25;          // 常量
let score = 100;                 // 类型推断
```

```python
# Python
name: str = "Alice"              # 类型提示 (可选)
age = 25                         # 动态类型
NAME: str = "Alice"              # 常量 (约定大写，实际可变)
```

```go
// Go
var name string = "Alice"        // 显式声明
age := 25                        // 短变量声明 (类型推断)
const PI = 3.14159               // 常量
```

```rust
// Rust
let name: &str = "Alice";        // 不可变绑定 (默认)
let mut age: i32 = 25;           // 可变绑定
const PI: f64 = 3.14159;         // 常量
```

### 1.2 条件语句

```typescript
// TypeScript
if (score >= 90) {
    console.log("A");
} else if (score >= 60) {
    console.log("Pass");
} else {
    console.log("Fail");
}

// 三元运算
const result = score >= 60 ? "Pass" : "Fail";
```

```python
# Python
if score >= 90:
    print("A")
elif score >= 60:
    print("Pass")
else:
    print("Fail")

# 三元表达式
result = "Pass" if score >= 60 else "Fail"
```

```go
// Go
if score >= 90 {
    fmt.Println("A")
} else if score >= 60 {
    fmt.Println("Pass")
} else {
    fmt.Println("Fail")
}
// Go 没有三元运算符
```

```rust
// Rust - if 是表达式
let result = if score >= 90 {
    "A"
} else if score >= 60 {
    "Pass"
} else {
    "Fail"
};
```

### 1.3 循环

```typescript
// TypeScript
// for 循环
for (let i = 0; i < 10; i++) {
    console.log(i);
}

// for...of 遍历
for (const item of [1, 2, 3]) {
    console.log(item);
}

// while
while (condition) { /* ... */ }
```

```python
# Python
# for 循环 (range)
for i in range(10):
    print(i)

# for 遍历
for item in [1, 2, 3]:
    print(item)

# while
while condition:
    pass

# 列表推导式
squares = [x**2 for x in range(10)]
```

```go
// Go
// for 循环 (唯一的循环结构)
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// for range 遍历
for index, value := range []int{1, 2, 3} {
    fmt.Println(index, value)
}

// while 风格
for condition {
    // ...
}
```

```rust
// Rust
// for 循环 (迭代器)
for i in 0..10 {
    println!("{}", i);
}

// for 遍历
for item in [1, 2, 3] {
    println!("{}", item);
}

// loop 无限循环
loop {
    break;
}

// while
while condition {
    // ...
}
```

---

## 2️⃣ 数据类型对比

### 2.1 基本类型

| 类型 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **整数** | `number` | `int` | `int`, `int8-64`, `uint8-64` | `i8-128`, `u8-128`, `isize`, `usize` |
| **浮点** | `number` | `float` | `float32`, `float64` | `f32`, `f64` |
| **布尔** | `boolean` | `bool` | `bool` | `bool` |
| **字符串** | `string` | `str` | `string` | `String`, `&str` |
| **字符** | - | - | `rune` | `char` |
| **空值** | `null`, `undefined` | `None` | `nil` | `Option<T>` |

### 2.2 复合类型

#### Array / List / Vec

```typescript
// TypeScript - Array
const numbers: number[] = [1, 2, 3];
const mixed: Array<number | string> = [1, "two", 3];

// 方法
numbers.push(4);                    // 添加
numbers.pop();                      // 删除末尾
numbers.slice(0, 2);                // 切片
numbers.map(x => x * 2);            // 映射
numbers.filter(x => x > 1);         // 过滤
numbers.reduce((a, b) => a + b, 0); // 归约
numbers.find(x => x > 1);           // 查找
numbers.includes(2);                // 包含
```

```python
# Python - List
numbers: list[int] = [1, 2, 3]
mixed: list = [1, "two", 3]

# 方法
numbers.append(4)                   # 添加
numbers.pop()                       # 删除末尾
numbers[0:2]                        # 切片
list(map(lambda x: x * 2, numbers)) # 映射
[x * 2 for x in numbers]            # 列表推导式
list(filter(lambda x: x > 1, numbers))  # 过滤
from functools import reduce
reduce(lambda a, b: a + b, numbers, 0)  # 归约
2 in numbers                        # 包含
```

```go
// Go - Slice
numbers := []int{1, 2, 3}

// 操作
numbers = append(numbers, 4)        // 添加
numbers = numbers[:len(numbers)-1]  // 删除末尾
numbers[0:2]                        // 切片

// Go 没有内置 map/filter，需要循环或泛型
// 使用 Go 1.21+ 的 slices 包
import "slices"
slices.Contains(numbers, 2)         // 包含
```

```rust
// Rust - Vec
let mut numbers: Vec<i32> = vec![1, 2, 3];

// 方法
numbers.push(4);                        // 添加
numbers.pop();                          // 删除末尾
&numbers[0..2];                         // 切片
numbers.iter().map(|x| x * 2).collect::<Vec<_>>();  // 映射
numbers.iter().filter(|&x| *x > 1).collect::<Vec<_>>();  // 过滤
numbers.iter().fold(0, |a, b| a + b);   // 归约
numbers.iter().find(|&&x| x > 1);       // 查找
numbers.contains(&2);                   // 包含
```

#### Tuple

```typescript
// TypeScript - Tuple
const point: [number, number] = [10, 20];
const mixed: [string, number, boolean] = ["hello", 42, true];

// 解构
const [x, y] = point;
```

```python
# Python - Tuple (不可变)
point: tuple[int, int] = (10, 20)
mixed: tuple[str, int, bool] = ("hello", 42, True)

# 解构
x, y = point

# 命名元组
from collections import namedtuple
Point = namedtuple('Point', ['x', 'y'])
p = Point(10, 20)
print(p.x, p.y)
```

```go
// Go - 没有内置元组，使用 struct 或多返回值
type Point struct {
    X, Y int
}
point := Point{10, 20}

// 多返回值模拟元组
func getPoint() (int, int) {
    return 10, 20
}
x, y := getPoint()
```

```rust
// Rust - Tuple
let point: (i32, i32) = (10, 20);
let mixed: (&str, i32, bool) = ("hello", 42, true);

// 解构
let (x, y) = point;

// 访问
let x = point.0;
let y = point.1;
```

#### Set

```typescript
// TypeScript - Set
const set = new Set<number>([1, 2, 3]);

set.add(4);           // 添加
set.delete(1);        // 删除
set.has(2);           // 包含
set.size;             // 大小

// 集合运算 (需要手动实现)
const union = new Set([...set1, ...set2]);
const intersection = new Set([...set1].filter(x => set2.has(x)));
```

```python
# Python - Set
numbers: set[int] = {1, 2, 3}

numbers.add(4)        # 添加
numbers.remove(1)     # 删除 (不存在报错)
numbers.discard(1)    # 删除 (不存在不报错)
2 in numbers          # 包含
len(numbers)          # 大小

# 集合运算
union = set1 | set2               # 并集
intersection = set1 & set2         # 交集
difference = set1 - set2           # 差集
symmetric_diff = set1 ^ set2       # 对称差集
```

```go
// Go - 没有内置 Set，使用 map[T]struct{} 模拟
set := make(map[int]struct{})
set[1] = struct{}{}   // 添加
set[2] = struct{}{}
delete(set, 1)        // 删除
_, exists := set[2]   // 包含
len(set)              // 大小
```

```rust
// Rust - HashSet
use std::collections::HashSet;

let mut set: HashSet<i32> = HashSet::from([1, 2, 3]);

set.insert(4);        // 添加
set.remove(&1);       // 删除
set.contains(&2);     // 包含
set.len();            // 大小

// 集合运算
let union: HashSet<_> = set1.union(&set2).collect();
let intersection: HashSet<_> = set1.intersection(&set2).collect();
let difference: HashSet<_> = set1.difference(&set2).collect();
```

#### Map / Dict / HashMap

```typescript
// TypeScript - Map & Object
const map = new Map<string, number>();
map.set("a", 1);
map.get("a");         // 1
map.has("a");         // true
map.delete("a");
map.size;

// Object 作为字典
const obj: Record<string, number> = { a: 1, b: 2 };
obj["c"] = 3;
Object.keys(obj);     // ["a", "b", "c"]
Object.values(obj);   // [1, 2, 3]
Object.entries(obj);  // [["a", 1], ["b", 2], ["c", 3]]
```

```python
# Python - Dict
d: dict[str, int] = {"a": 1, "b": 2}

d["c"] = 3            # 设置
d.get("a")            # 获取 (不存在返回 None)
d["a"]                # 获取 (不存在报错)
"a" in d              # 包含
del d["a"]            # 删除
len(d)                # 大小

d.keys()              # 所有键
d.values()            # 所有值
d.items()             # 所有键值对

# 字典推导式
squared = {k: v**2 for k, v in d.items()}
```

```go
// Go - Map
m := make(map[string]int)
m["a"] = 1            // 设置
val := m["a"]         // 获取
val, ok := m["a"]     // 获取 + 检查存在
delete(m, "a")        // 删除
len(m)                // 大小

// 遍历
for key, value := range m {
    fmt.Println(key, value)
}
```

```rust
// Rust - HashMap
use std::collections::HashMap;

let mut map: HashMap<&str, i32> = HashMap::new();

map.insert("a", 1);   // 设置
map.get("a");         // 获取 Option<&V>
map.contains_key("a"); // 包含
map.remove("a");      // 删除
map.len();            // 大小

// Entry API
map.entry("b").or_insert(2);

// 遍历
for (key, value) in &map {
    println!("{}: {}", key, value);
}
```

### 2.3 字符串操作

```typescript
// TypeScript
const s = "Hello, World!";

s.length;                    // 长度
s.toUpperCase();             // 大写
s.toLowerCase();             // 小写
s.trim();                    // 去空白
s.split(", ");               // 分割
s.includes("World");         // 包含
s.startsWith("Hello");       // 前缀
s.endsWith("!");             // 后缀
s.replace("World", "Rust");  // 替换
s.slice(0, 5);               // 切片
`Hello, ${name}!`;           // 模板字符串
```

```python
# Python
s = "Hello, World!"

len(s)                       # 长度
s.upper()                    # 大写
s.lower()                    # 小写
s.strip()                    # 去空白
s.split(", ")                # 分割
"World" in s                 # 包含
s.startswith("Hello")        # 前缀
s.endswith("!")              # 后缀
s.replace("World", "Rust")   # 替换
s[0:5]                       # 切片
f"Hello, {name}!"            # f-string
```

```go
// Go
import "strings"

s := "Hello, World!"

len(s)                         // 字节长度
strings.ToUpper(s)             // 大写
strings.ToLower(s)             // 小写
strings.TrimSpace(s)           // 去空白
strings.Split(s, ", ")         // 分割
strings.Contains(s, "World")   // 包含
strings.HasPrefix(s, "Hello")  // 前缀
strings.HasSuffix(s, "!")      // 后缀
strings.Replace(s, "World", "Go", 1)  // 替换
s[0:5]                         // 切片 (字节)
fmt.Sprintf("Hello, %s!", name) // 格式化
```

```rust
// Rust
let s = String::from("Hello, World!");

s.len();                       // 字节长度
s.chars().count();             // 字符数
s.to_uppercase();              // 大写
s.to_lowercase();              // 小写
s.trim();                      // 去空白
s.split(", ");                 // 分割 (迭代器)
s.contains("World");           // 包含
s.starts_with("Hello");        // 前缀
s.ends_with("!");              // 后缀
s.replace("World", "Rust");    // 替换
&s[0..5];                      // 切片 (字节)
format!("Hello, {}!", name);   // 格式化
```

---

## 3️⃣ 类与结构体

### 3.1 Class / Struct 定义

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
        return `Hello, I'm ${this.name}`;
    }
    
    // Getter/Setter
    get info(): string {
        return `${this.name}, ${this.age}`;
    }
    
    // 静态方法
    static create(name: string): Person {
        return new Person(name, 0);
    }
}

// 继承
class Student extends Person {
    school: string;
    
    constructor(name: string, age: number, school: string) {
        super(name, age);
        this.school = school;
    }
    
    // 重写
    greet(): string {
        return `${super.greet()}, I study at ${this.school}`;
    }
}
```

```python
# Python - Class
from dataclasses import dataclass
from typing import ClassVar

class Person:
    def __init__(self, name: str, age: int):
        self._name = name  # 约定私有
        self.age = age
    
    def greet(self) -> str:
        return f"Hello, I'm {self._name}"
    
    # Property
    @property
    def info(self) -> str:
        return f"{self._name}, {self.age}"
    
    # 静态方法
    @staticmethod
    def create(name: str) -> "Person":
        return Person(name, 0)
    
    # 类方法
    @classmethod
    def from_dict(cls, data: dict) -> "Person":
        return cls(data["name"], data["age"])

# Dataclass (简化版)
@dataclass
class Point:
    x: int
    y: int
    
# 继承
class Student(Person):
    def __init__(self, name: str, age: int, school: str):
        super().__init__(name, age)
        self.school = school
    
    def greet(self) -> str:
        return f"{super().greet()}, I study at {self.school}"
```

```go
// Go - Struct (没有 class，只有 struct + 方法)
type Person struct {
    name string  // 小写 = 私有
    Age  int     // 大写 = 公开
}

// 方法 (值接收者)
func (p Person) Greet() string {
    return fmt.Sprintf("Hello, I'm %s", p.name)
}

// 方法 (指针接收者 - 可修改)
func (p *Person) SetAge(age int) {
    p.Age = age
}

// 构造函数 (约定)
func NewPerson(name string, age int) *Person {
    return &Person{name: name, Age: age}
}

// "继承" - 嵌入
type Student struct {
    Person        // 嵌入
    School string
}

func (s Student) Greet() string {
    return fmt.Sprintf("%s, I study at %s", s.Person.Greet(), s.School)
}
```

```rust
// Rust - Struct + impl
pub struct Person {
    name: String,      // 默认私有
    pub age: i32,      // 公开
}

impl Person {
    // 关联函数 (类似静态方法)
    pub fn new(name: String, age: i32) -> Self {
        Person { name, age }
    }
    
    // 方法 (&self = 不可变借用)
    pub fn greet(&self) -> String {
        format!("Hello, I'm {}", self.name)
    }
    
    // 方法 (&mut self = 可变借用)
    pub fn set_age(&mut self, age: i32) {
        self.age = age;
    }
}

// "继承" - Rust 使用组合 + trait
pub struct Student {
    person: Person,
    pub school: String,
}

impl Student {
    pub fn new(name: String, age: i32, school: String) -> Self {
        Student {
            person: Person::new(name, age),
            school,
        }
    }
    
    pub fn greet(&self) -> String {
        format!("{}, I study at {}", self.person.greet(), self.school)
    }
}
```

### 3.2 接口 / Trait / Protocol

```typescript
// TypeScript - Interface
interface Greeter {
    greet(): string;
}

interface Named {
    name: string;
}

// 实现接口 (隐式)
class Person implements Greeter, Named {
    name: string;
    constructor(name: string) {
        this.name = name;
    }
    greet(): string {
        return `Hello, ${this.name}`;
    }
}

// 泛型接口
interface Repository<T> {
    find(id: string): T | null;
    save(item: T): void;
}
```

```python
# Python - Protocol (typing) / ABC
from abc import ABC, abstractmethod
from typing import Protocol, TypeVar

# Protocol (结构化子类型 - 鸭子类型)
class Greeter(Protocol):
    def greet(self) -> str: ...

# ABC (抽象基类 - 名义子类型)
class Animal(ABC):
    @abstractmethod
    def speak(self) -> str:
        pass

class Dog(Animal):
    def speak(self) -> str:
        return "Woof!"

# 泛型
T = TypeVar('T')

class Repository(Protocol[T]):
    def find(self, id: str) -> T | None: ...
    def save(self, item: T) -> None: ...
```

```go
// Go - Interface (隐式实现)
type Greeter interface {
    Greet() string
}

type Named interface {
    Name() string
}

// 组合接口
type NamedGreeter interface {
    Greeter
    Named
}

// 任何实现了 Greet() 方法的类型都自动实现 Greeter
type Person struct {
    name string
}

func (p Person) Greet() string {
    return "Hello, " + p.name
}

// 空接口 (任意类型)
func printAny(v interface{}) {  // 或 any
    fmt.Println(v)
}
```

```rust
// Rust - Trait
pub trait Greeter {
    fn greet(&self) -> String;
    
    // 默认实现
    fn greet_loud(&self) -> String {
        self.greet().to_uppercase()
    }
}

// 实现 trait
impl Greeter for Person {
    fn greet(&self) -> String {
        format!("Hello, {}", self.name)
    }
}

// Trait 约束
fn print_greeting<T: Greeter>(item: &T) {
    println!("{}", item.greet());
}

// 多重约束
fn complex<T: Greeter + Clone + Send>(item: T) { }

// Trait 对象 (动态分发)
fn dynamic_greeting(item: &dyn Greeter) {
    println!("{}", item.greet());
}
```

---

## 4️⃣ 枚举 (Enum)

```typescript
// TypeScript - Enum
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

enum HttpStatus {
    Ok = 200,
    NotFound = 404,
    InternalError = 500,
}

// 字符串枚举
enum Color {
    Red = "RED",
    Green = "GREEN",
    Blue = "BLUE",
}

// 使用
const dir = Direction.Up;
```

```python
# Python - Enum
from enum import Enum, auto

class Direction(Enum):
    UP = auto()
    DOWN = auto()
    LEFT = auto()
    RIGHT = auto()

class HttpStatus(Enum):
    OK = 200
    NOT_FOUND = 404
    INTERNAL_ERROR = 500

# 使用
dir = Direction.UP
status = HttpStatus.OK.value  # 200
```

```go
// Go - iota (没有真正的 enum)
type Direction int

const (
    Up Direction = iota  // 0
    Down                 // 1
    Left                 // 2
    Right                // 3
)

type HttpStatus int

const (
    StatusOK            HttpStatus = 200
    StatusNotFound      HttpStatus = 404
    StatusInternalError HttpStatus = 500
)
```

```rust
// Rust - Enum (代数数据类型 - 最强大)
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

// 带值枚举
enum HttpStatus {
    Ok = 200,
    NotFound = 404,
    InternalError = 500,
}

// 带数据枚举 (ADT)
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

// Option 和 Result (标准库)
enum Option<T> {
    Some(T),
    None,
}

enum Result<T, E> {
    Ok(T),
    Err(E),
}

// 模式匹配
fn handle_message(msg: Message) {
    match msg {
        Message::Quit => println!("Quit"),
        Message::Move { x, y } => println!("Move to ({}, {})", x, y),
        Message::Write(text) => println!("Text: {}", text),
        Message::ChangeColor(r, g, b) => println!("Color: {} {} {}", r, g, b),
    }
}
```

---

## 5️⃣ 错误处理

```typescript
// TypeScript - try/catch + throw
function divide(a: number, b: number): number {
    if (b === 0) {
        throw new Error("Division by zero");
    }
    return a / b;
}

try {
    const result = divide(10, 0);
} catch (error) {
    if (error instanceof Error) {
        console.error(error.message);
    }
} finally {
    console.log("Cleanup");
}

// 自定义错误
class ValidationError extends Error {
    constructor(public field: string, message: string) {
        super(message);
        this.name = "ValidationError";
    }
}
```

```python
# Python - try/except + raise
def divide(a: float, b: float) -> float:
    if b == 0:
        raise ValueError("Division by zero")
    return a / b

try:
    result = divide(10, 0)
except ValueError as e:
    print(f"Error: {e}")
except Exception as e:
    print(f"Unexpected: {e}")
else:
    print("Success")
finally:
    print("Cleanup")

# 自定义异常
class ValidationError(Exception):
    def __init__(self, field: str, message: str):
        self.field = field
        super().__init__(message)
```

```go
// Go - 显式错误返回 (没有异常)
import (
    "errors"
    "fmt"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// 使用
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(result)

// 自定义错误
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// 错误包装 (Go 1.13+)
if err != nil {
    return fmt.Errorf("operation failed: %w", err)
}

// 检查特定错误
if errors.Is(err, ErrNotFound) { }
var valErr *ValidationError
if errors.As(err, &valErr) { }
```

```rust
// Rust - Result<T, E> (没有异常)
use std::error::Error;
use std::fmt;

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("Division by zero".to_string())
    } else {
        Ok(a / b)
    }
}

// 使用
match divide(10.0, 0.0) {
    Ok(result) => println!("Result: {}", result),
    Err(e) => println!("Error: {}", e),
}

// ? 操作符 (错误传播)
fn calculate() -> Result<f64, String> {
    let x = divide(10.0, 2.0)?;  // 错误自动返回
    let y = divide(x, 3.0)?;
    Ok(y)
}

// if let
if let Ok(result) = divide(10.0, 2.0) {
    println!("{}", result);
}

// 自定义错误
#[derive(Debug)]
struct ValidationError {
    field: String,
    message: String,
}

impl fmt::Display for ValidationError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}: {}", self.field, self.message)
    }
}

impl Error for ValidationError {}

// thiserror crate (简化)
use thiserror::Error;

#[derive(Error, Debug)]
enum AppError {
    #[error("IO error: {0}")]
    Io(#[from] std::io::Error),
    #[error("Parse error: {0}")]
    Parse(String),
}
```

---

## 6️⃣ 日期时间

```typescript
// TypeScript - Date (内置) / dayjs (推荐)
const now = new Date();
const specific = new Date("2024-01-15T10:30:00Z");

now.getFullYear();    // 年
now.getMonth();       // 月 (0-11)
now.getDate();        // 日
now.getHours();       // 时
now.getTime();        // 时间戳 (毫秒)

// 格式化 (推荐使用 dayjs/date-fns)
import dayjs from 'dayjs';
dayjs().format('YYYY-MM-DD HH:mm:ss');
dayjs('2024-01-15').add(1, 'day');
dayjs('2024-01-15').diff('2024-01-10', 'day');
```

```python
# Python - datetime
from datetime import datetime, date, time, timedelta
from zoneinfo import ZoneInfo

now = datetime.now()
utc_now = datetime.now(ZoneInfo("UTC"))
specific = datetime(2024, 1, 15, 10, 30, 0)

now.year        # 年
now.month       # 月
now.day         # 日
now.hour        # 时
now.timestamp() # 时间戳 (秒)

# 格式化
now.strftime("%Y-%m-%d %H:%M:%S")
datetime.strptime("2024-01-15", "%Y-%m-%d")

# 计算
tomorrow = now + timedelta(days=1)
diff = (datetime(2024, 1, 15) - datetime(2024, 1, 10)).days
```

```go
// Go - time
import "time"

now := time.Now()
specific := time.Date(2024, 1, 15, 10, 30, 0, 0, time.UTC)

now.Year()        // 年
now.Month()       // 月
now.Day()         // 日
now.Hour()        // 时
now.Unix()        // 时间戳 (秒)

// 格式化 (Go 特有: 使用参考时间 2006-01-02 15:04:05)
now.Format("2006-01-02 15:04:05")
time.Parse("2006-01-02", "2024-01-15")

// 计算
tomorrow := now.Add(24 * time.Hour)
diff := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC).Sub(
    time.Date(2024, 1, 10, 0, 0, 0, 0, time.UTC),
)
```

```rust
// Rust - chrono crate
use chrono::{DateTime, Utc, Local, NaiveDate, Duration, TimeZone};

let now: DateTime<Local> = Local::now();
let utc_now: DateTime<Utc> = Utc::now();
let specific = NaiveDate::from_ymd_opt(2024, 1, 15)
    .unwrap()
    .and_hms_opt(10, 30, 0)
    .unwrap();

now.year();       // 年
now.month();      // 月
now.day();        // 日
now.hour();       // 时
now.timestamp();  // 时间戳 (秒)

// 格式化
now.format("%Y-%m-%d %H:%M:%S").to_string();
NaiveDate::parse_from_str("2024-01-15", "%Y-%m-%d");

// 计算
let tomorrow = now + Duration::days(1);
let diff = NaiveDate::from_ymd_opt(2024, 1, 15).unwrap()
    .signed_duration_since(NaiveDate::from_ymd_opt(2024, 1, 10).unwrap());
```

---

## 7️⃣ 正则表达式

```typescript
// TypeScript - RegExp
const pattern = /\d{3}-\d{4}/g;
const pattern2 = new RegExp("\\d{3}-\\d{4}", "g");

const text = "Call 123-4567 or 987-6543";

// 测试
pattern.test(text);                    // true

// 匹配
text.match(pattern);                   // ["123-4567", "987-6543"]

// 替换
text.replace(pattern, "XXX-XXXX");

// 捕获组
const emailPattern = /(\w+)@(\w+\.\w+)/;
const match = "test@example.com".match(emailPattern);
// match[0] = "test@example.com", match[1] = "test", match[2] = "example.com"

// 命名捕获组
const namedPattern = /(?<user>\w+)@(?<domain>\w+\.\w+)/;
const groups = "test@example.com".match(namedPattern)?.groups;
// groups.user = "test", groups.domain = "example.com"
```

```python
# Python - re
import re

pattern = re.compile(r"\d{3}-\d{4}")
text = "Call 123-4567 or 987-6543"

# 测试/搜索
re.search(pattern, text)               # Match object or None
re.match(r"Call", text)                # 从开头匹配

# 查找所有
re.findall(r"\d{3}-\d{4}", text)       # ["123-4567", "987-6543"]

# 替换
re.sub(r"\d{3}-\d{4}", "XXX-XXXX", text)

# 捕获组
match = re.search(r"(\w+)@(\w+\.\w+)", "test@example.com")
match.group(0)  # "test@example.com"
match.group(1)  # "test"
match.group(2)  # "example.com"

# 命名捕获组
match = re.search(r"(?P<user>\w+)@(?P<domain>\w+\.\w+)", "test@example.com")
match.group("user")    # "test"
match.group("domain")  # "example.com"
```

```go
// Go - regexp
import "regexp"

pattern := regexp.MustCompile(`\d{3}-\d{4}`)
text := "Call 123-4567 or 987-6543"

// 测试
pattern.MatchString(text)              // true

// 查找
pattern.FindString(text)               // "123-4567"
pattern.FindAllString(text, -1)        // ["123-4567", "987-6543"]

// 替换
pattern.ReplaceAllString(text, "XXX-XXXX")

// 捕获组
emailPattern := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
match := emailPattern.FindStringSubmatch("test@example.com")
// match[0] = "test@example.com", match[1] = "test", match[2] = "example.com"

// 命名捕获组
namedPattern := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+\.\w+)`)
match = namedPattern.FindStringSubmatch("test@example.com")
names := namedPattern.SubexpNames()
```

```rust
// Rust - regex crate
use regex::Regex;

let pattern = Regex::new(r"\d{3}-\d{4}").unwrap();
let text = "Call 123-4567 or 987-6543";

// 测试
pattern.is_match(text);                // true

// 查找
pattern.find(text);                    // Some(Match)
pattern.find_iter(text).collect::<Vec<_>>();

// 替换
pattern.replace_all(text, "XXX-XXXX");

// 捕获组
let email_pattern = Regex::new(r"(\w+)@(\w+\.\w+)").unwrap();
if let Some(caps) = email_pattern.captures("test@example.com") {
    caps.get(0);  // "test@example.com"
    caps.get(1);  // "test"
    caps.get(2);  // "example.com"
}

// 命名捕获组
let named_pattern = Regex::new(r"(?P<user>\w+)@(?P<domain>\w+\.\w+)").unwrap();
if let Some(caps) = named_pattern.captures("test@example.com") {
    &caps["user"];    // "test"
    &caps["domain"];  // "example.com"
}
```

---

## 8️⃣ 迭代器与函数式编程

```typescript
// TypeScript - 迭代器和高阶函数
const numbers = [1, 2, 3, 4, 5];

// 链式调用
const result = numbers
    .filter(x => x % 2 === 0)     // [2, 4]
    .map(x => x * 2)               // [4, 8]
    .reduce((a, b) => a + b, 0);   // 12

// Generator
function* range(start: number, end: number) {
    for (let i = start; i < end; i++) {
        yield i;
    }
}

for (const n of range(0, 5)) {
    console.log(n);
}

// 自定义迭代器
const iterable = {
    [Symbol.iterator]() {
        let i = 0;
        return {
            next() {
                return i < 3
                    ? { value: i++, done: false }
                    : { done: true };
            }
        };
    }
};
```

```python
# Python - 迭代器和高阶函数
from functools import reduce
from itertools import islice, chain, cycle, repeat

numbers = [1, 2, 3, 4, 5]

# 链式调用 (使用生成器表达式)
result = reduce(
    lambda a, b: a + b,
    (x * 2 for x in numbers if x % 2 == 0),
    0
)  # 12

# 更 Pythonic 的方式
evens = [x for x in numbers if x % 2 == 0]
doubled = [x * 2 for x in evens]
total = sum(doubled)

# Generator
def range_gen(start: int, end: int):
    i = start
    while i < end:
        yield i
        i += 1

# 自定义迭代器
class Counter:
    def __init__(self, max_count: int):
        self.max = max_count
        self.count = 0
    
    def __iter__(self):
        return self
    
    def __next__(self):
        if self.count >= self.max:
            raise StopIteration
        self.count += 1
        return self.count - 1

# itertools
list(islice(cycle([1, 2, 3]), 10))  # 取前10个
list(chain([1, 2], [3, 4]))          # 连接
```

```go
// Go - 没有内置迭代器，Go 1.23 引入迭代器
// 通常使用 for 循环 + channel

numbers := []int{1, 2, 3, 4, 5}

// 手动实现 map/filter/reduce
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func Filter[T any](slice []T, fn func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if fn(v) {
            result = append(result, v)
        }
    }
    return result
}

func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    acc := initial
    for _, v := range slice {
        acc = fn(acc, v)
    }
    return acc
}

// Channel 作为迭代器
func Range(start, end int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := start; i < end; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

for n := range Range(0, 5) {
    fmt.Println(n)
}
```

```rust
// Rust - 迭代器 (最强大)
let numbers = vec![1, 2, 3, 4, 5];

// 链式调用 (零成本抽象)
let result: i32 = numbers.iter()
    .filter(|&x| x % 2 == 0)      // [2, 4]
    .map(|x| x * 2)                // [4, 8]
    .sum();                        // 12

// 更多迭代器方法
numbers.iter().take(3);            // 取前3个
numbers.iter().skip(2);            // 跳过前2个
numbers.iter().enumerate();        // 带索引
numbers.iter().zip([10, 20, 30]);  // 配对
numbers.iter().all(|&x| x > 0);    // 全部满足
numbers.iter().any(|&x| x > 3);    // 任意满足
numbers.iter().max();              // 最大值
numbers.iter().min();              // 最小值
numbers.iter().count();            // 计数
numbers.iter().collect::<Vec<_>>(); // 收集

// 自定义迭代器
struct Counter {
    count: u32,
    max: u32,
}

impl Iterator for Counter {
    type Item = u32;
    
    fn next(&mut self) -> Option<Self::Item> {
        if self.count < self.max {
            self.count += 1;
            Some(self.count - 1)
        } else {
            None
        }
    }
}

// 使用 impl Iterator 返回复杂迭代器
fn even_squares(n: i32) -> impl Iterator<Item = i32> {
    (0..n).filter(|x| x % 2 == 0).map(|x| x * x)
}
```

---

## 9️⃣ 并发与异步

### 9.1 异步编程

```typescript
// TypeScript - Promise / async-await
// 基于事件循环的单线程异步

// Promise
function fetchData(url: string): Promise<string> {
    return new Promise((resolve, reject) => {
        setTimeout(() => resolve("data"), 1000);
    });
}

// async/await
async function getData(): Promise<string> {
    const data = await fetchData("https://api.example.com");
    return data;
}

// 并行执行
async function parallel() {
    const [a, b, c] = await Promise.all([
        fetchData("url1"),
        fetchData("url2"),
        fetchData("url3"),
    ]);
}

// 竞争
const fastest = await Promise.race([fetch1(), fetch2()]);

// 错误处理
try {
    const data = await fetchData("url");
} catch (error) {
    console.error(error);
}
```

```python
# Python - asyncio
import asyncio
from typing import Coroutine

# async 函数
async def fetch_data(url: str) -> str:
    await asyncio.sleep(1)  # 模拟 IO
    return "data"

# 运行异步代码
async def main():
    data = await fetch_data("https://api.example.com")
    print(data)

asyncio.run(main())

# 并行执行
async def parallel():
    results = await asyncio.gather(
        fetch_data("url1"),
        fetch_data("url2"),
        fetch_data("url3"),
    )
    return results

# 超时
async def with_timeout():
    try:
        result = await asyncio.wait_for(fetch_data("url"), timeout=5.0)
    except asyncio.TimeoutError:
        print("Timeout!")

# 任务
async def tasks():
    task = asyncio.create_task(fetch_data("url"))
    # ... 其他工作
    result = await task
```

```go
// Go - Goroutine + Channel
// CSP 模型 (Communicating Sequential Processes)

// Goroutine (轻量级线程)
go func() {
    fmt.Println("In goroutine")
}()

// Channel (通信)
ch := make(chan string)

go func() {
    time.Sleep(time.Second)
    ch <- "data"
}()

data := <-ch  // 阻塞等待

// 带缓冲 Channel
buffered := make(chan int, 10)

// Select (多路复用)
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

// WaitGroup (等待多个 goroutine)
import "sync"

var wg sync.WaitGroup
for i := 0; i < 3; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        fmt.Println(n)
    }(i)
}
wg.Wait()

// Mutex (互斥锁)
var mu sync.Mutex
mu.Lock()
// 临界区
mu.Unlock()
```

```rust
// Rust - async/await (需要运行时如 tokio)
use tokio;

// async 函数
async fn fetch_data(url: &str) -> String {
    tokio::time::sleep(tokio::time::Duration::from_secs(1)).await;
    "data".to_string()
}

#[tokio::main]
async fn main() {
    let data = fetch_data("https://api.example.com").await;
    println!("{}", data);
}

// 并行执行
async fn parallel() {
    let (a, b, c) = tokio::join!(
        fetch_data("url1"),
        fetch_data("url2"),
        fetch_data("url3"),
    );
}

// 生成任务
async fn spawn_tasks() {
    let handle = tokio::spawn(async {
        fetch_data("url").await
    });
    let result = handle.await.unwrap();
}

// Channel
use tokio::sync::mpsc;

async fn channel_example() {
    let (tx, mut rx) = mpsc::channel(32);
    
    tokio::spawn(async move {
        tx.send("data").await.unwrap();
    });
    
    while let Some(msg) = rx.recv().await {
        println!("{}", msg);
    }
}

// 超时
use tokio::time::{timeout, Duration};

async fn with_timeout() {
    match timeout(Duration::from_secs(5), fetch_data("url")).await {
        Ok(result) => println!("{}", result),
        Err(_) => println!("Timeout!"),
    }
}
```

### 9.2 多线程/多进程

```typescript
// TypeScript/Node.js - Worker Threads / Cluster
import { Worker, isMainThread, parentPort } from 'worker_threads';
import cluster from 'cluster';
import os from 'os';

// Worker Thread
if (isMainThread) {
    const worker = new Worker(__filename);
    worker.on('message', (msg) => console.log(msg));
    worker.postMessage('Hello');
} else {
    parentPort?.on('message', (msg) => {
        parentPort?.postMessage(`Received: ${msg}`);
    });
}

// Cluster (多进程)
if (cluster.isPrimary) {
    const numCPUs = os.cpus().length;
    for (let i = 0; i < numCPUs; i++) {
        cluster.fork();
    }
} else {
    // Worker process
}
```

```python
# Python - threading / multiprocessing
import threading
import multiprocessing
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# Threading (受 GIL 限制，适合 IO 密集)
def worker(name: str):
    print(f"Thread {name} running")

threads = []
for i in range(5):
    t = threading.Thread(target=worker, args=(f"T{i}",))
    threads.append(t)
    t.start()

for t in threads:
    t.join()

# Lock
lock = threading.Lock()
with lock:
    # 临界区
    pass

# Multiprocessing (真正并行，适合 CPU 密集)
def cpu_intensive(n: int) -> int:
    return sum(i * i for i in range(n))

if __name__ == "__main__":
    with multiprocessing.Pool(4) as pool:
        results = pool.map(cpu_intensive, [1000000] * 4)

# ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=5) as executor:
    futures = [executor.submit(worker, f"T{i}") for i in range(5)]
    for future in futures:
        result = future.result()

# ProcessPoolExecutor
with ProcessPoolExecutor(max_workers=4) as executor:
    results = list(executor.map(cpu_intensive, [1000000] * 4))
```

```go
// Go - goroutine (前面已介绍)
// Go 的 goroutine 非常轻量，可以创建数百万个

// 工作池模式
func workerPool(numWorkers int, jobs <-chan int, results chan<- int) {
    for i := 0; i < numWorkers; i++ {
        go func() {
            for job := range jobs {
                results <- job * 2
            }
        }()
    }
}

// 使用
jobs := make(chan int, 100)
results := make(chan int, 100)

workerPool(5, jobs, results)

for i := 0; i < 100; i++ {
    jobs <- i
}
close(jobs)

for i := 0; i < 100; i++ {
    <-results
}
```

```rust
// Rust - std::thread / rayon / tokio
use std::thread;
use std::sync::{Arc, Mutex, mpsc};

// 基本线程
let handle = thread::spawn(|| {
    println!("In thread");
    42
});
let result = handle.join().unwrap();

// 共享状态
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

// Channel
let (tx, rx) = mpsc::channel();

thread::spawn(move || {
    tx.send("Hello").unwrap();
});

let received = rx.recv().unwrap();

// Rayon (数据并行)
use rayon::prelude::*;

let sum: i32 = (0..1000000)
    .into_par_iter()  // 并行迭代器
    .map(|x| x * 2)
    .sum();

// Scoped threads (借用检查友好)
thread::scope(|s| {
    let data = vec![1, 2, 3];
    s.spawn(|| {
        println!("{:?}", data);  // 可以借用外部数据
    });
});
```

---

## 🔟 文件 IO

```typescript
// TypeScript/Node.js - fs
import { readFileSync, writeFileSync, readFile, writeFile } from 'fs';
import { readFile as readFileAsync } from 'fs/promises';

// 同步
const content = readFileSync('file.txt', 'utf-8');
writeFileSync('output.txt', 'Hello World');

// 异步 (callback)
readFile('file.txt', 'utf-8', (err, data) => {
    if (err) throw err;
    console.log(data);
});

// 异步 (Promise)
const data = await readFileAsync('file.txt', 'utf-8');

// 流
import { createReadStream, createWriteStream } from 'fs';

const readStream = createReadStream('large-file.txt');
const writeStream = createWriteStream('output.txt');

readStream.pipe(writeStream);

// 逐行读取
import { createInterface } from 'readline';

const rl = createInterface({
    input: createReadStream('file.txt'),
});

for await (const line of rl) {
    console.log(line);
}
```

```python
# Python - 内置 open / pathlib
from pathlib import Path

# 基本读写
with open("file.txt", "r", encoding="utf-8") as f:
    content = f.read()

with open("output.txt", "w", encoding="utf-8") as f:
    f.write("Hello World")

# 逐行读取
with open("file.txt", "r") as f:
    for line in f:
        print(line.strip())

# pathlib (推荐)
path = Path("file.txt")
content = path.read_text(encoding="utf-8")
path.write_text("Hello World", encoding="utf-8")

# 目录操作
Path("dir").mkdir(parents=True, exist_ok=True)
for p in Path(".").glob("**/*.py"):
    print(p)

# 二进制
with open("image.png", "rb") as f:
    data = f.read()

# 异步 IO (aiofiles)
import aiofiles

async def read_async():
    async with aiofiles.open("file.txt", "r") as f:
        content = await f.read()
```

```go
// Go - os / io / bufio
import (
    "os"
    "io"
    "bufio"
)

// 读取整个文件
data, err := os.ReadFile("file.txt")
if err != nil {
    log.Fatal(err)
}

// 写入文件
err = os.WriteFile("output.txt", []byte("Hello World"), 0644)

// 使用 File
file, err := os.Open("file.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// 逐行读取
scanner := bufio.NewScanner(file)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

// 写入
file, err = os.Create("output.txt")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

writer := bufio.NewWriter(file)
writer.WriteString("Hello World")
writer.Flush()

// 目录操作
os.Mkdir("dir", 0755)
os.MkdirAll("path/to/dir", 0755)

entries, _ := os.ReadDir(".")
for _, entry := range entries {
    fmt.Println(entry.Name())
}
```

```rust
// Rust - std::fs / std::io
use std::fs::{self, File};
use std::io::{self, BufRead, BufReader, Write, Read};
use std::path::Path;

// 读取整个文件
let content = fs::read_to_string("file.txt")?;

// 写入文件
fs::write("output.txt", "Hello World")?;

// 使用 File
let file = File::open("file.txt")?;
let mut content = String::new();
BufReader::new(file).read_to_string(&mut content)?;

// 逐行读取
let file = File::open("file.txt")?;
let reader = BufReader::new(file);
for line in reader.lines() {
    println!("{}", line?);
}

// 写入
let mut file = File::create("output.txt")?;
file.write_all(b"Hello World")?;

// 目录操作
fs::create_dir("dir")?;
fs::create_dir_all("path/to/dir")?;

for entry in fs::read_dir(".")? {
    let entry = entry?;
    println!("{}", entry.path().display());
}

// 异步 IO (tokio)
use tokio::fs;
use tokio::io::{AsyncReadExt, AsyncWriteExt};

async fn async_io() -> io::Result<()> {
    let content = fs::read_to_string("file.txt").await?;
    fs::write("output.txt", "Hello World").await?;
    Ok(())
}
```

---

## 1️⃣1️⃣ HTTP Client

```typescript
// TypeScript - fetch (内置) / axios
// fetch (Node.js 18+)
const response = await fetch('https://api.example.com/data');
const data = await response.json();

// POST 请求
const postResponse = await fetch('https://api.example.com/data', {
    method: 'POST',
    headers: {
        'Content-Type': 'application/json',
    },
    body: JSON.stringify({ name: 'Alice' }),
});

// axios
import axios from 'axios';

const { data } = await axios.get('https://api.example.com/data');

await axios.post('https://api.example.com/data', {
    name: 'Alice',
}, {
    headers: { 'Authorization': 'Bearer token' },
});

// 错误处理
try {
    const response = await fetch(url);
    if (!response.ok) {
        throw new Error(`HTTP error! status: ${response.status}`);
    }
} catch (error) {
    console.error('Fetch failed:', error);
}
```

```python
# Python - requests / httpx / aiohttp
import requests
import httpx
import aiohttp

# requests (同步)
response = requests.get("https://api.example.com/data")
data = response.json()

response = requests.post(
    "https://api.example.com/data",
    json={"name": "Alice"},
    headers={"Authorization": "Bearer token"},
)

# httpx (同步 + 异步)
# 同步
with httpx.Client() as client:
    response = client.get("https://api.example.com/data")
    
# 异步
async with httpx.AsyncClient() as client:
    response = await client.get("https://api.example.com/data")
    data = response.json()

# aiohttp (异步)
async with aiohttp.ClientSession() as session:
    async with session.get("https://api.example.com/data") as response:
        data = await response.json()

# 错误处理
try:
    response = requests.get(url)
    response.raise_for_status()
except requests.exceptions.HTTPError as e:
    print(f"HTTP error: {e}")
except requests.exceptions.RequestException as e:
    print(f"Request failed: {e}")
```

```go
// Go - net/http
import (
    "net/http"
    "io"
    "encoding/json"
    "bytes"
)

// GET
resp, err := http.Get("https://api.example.com/data")
if err != nil {
    log.Fatal(err)
}
defer resp.Body.Close()

body, _ := io.ReadAll(resp.Body)

var data map[string]interface{}
json.Unmarshal(body, &data)

// POST
payload := map[string]string{"name": "Alice"}
jsonData, _ := json.Marshal(payload)

resp, err = http.Post(
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

resp, err = client.Do(req)
```

```rust
// Rust - reqwest
use reqwest;
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
struct CreateUser {
    name: String,
}

#[derive(Deserialize)]
struct User {
    id: i32,
    name: String,
}

// 异步 (推荐)
#[tokio::main]
async fn main() -> Result<(), reqwest::Error> {
    // GET
    let body = reqwest::get("https://api.example.com/data")
        .await?
        .text()
        .await?;
    
    // JSON
    let user: User = reqwest::get("https://api.example.com/user/1")
        .await?
        .json()
        .await?;
    
    // POST
    let client = reqwest::Client::new();
    let new_user = CreateUser { name: "Alice".to_string() };
    
    let response = client.post("https://api.example.com/users")
        .json(&new_user)
        .header("Authorization", "Bearer token")
        .send()
        .await?;
    
    Ok(())
}

// 同步 (blocking feature)
let body = reqwest::blocking::get("https://api.example.com/data")?
    .text()?;
```

---

## 1️⃣2️⃣ HTTP Server

```typescript
// TypeScript - Express / Fastify / Hono
import express from 'express';

const app = express();
app.use(express.json());

// 路由
app.get('/users', (req, res) => {
    res.json([{ id: 1, name: 'Alice' }]);
});

app.get('/users/:id', (req, res) => {
    const { id } = req.params;
    res.json({ id, name: 'Alice' });
});

app.post('/users', (req, res) => {
    const { name } = req.body;
    res.status(201).json({ id: 1, name });
});

// 中间件
app.use((req, res, next) => {
    console.log(`${req.method} ${req.path}`);
    next();
});

// 错误处理
app.use((err, req, res, next) => {
    console.error(err);
    res.status(500).json({ error: 'Internal Server Error' });
});

app.listen(3000, () => {
    console.log('Server running on port 3000');
});
```

```python
# Python - FastAPI / Flask
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI()

class User(BaseModel):
    name: str

class UserResponse(BaseModel):
    id: int
    name: str

# 路由
@app.get("/users")
async def get_users() -> list[UserResponse]:
    return [UserResponse(id=1, name="Alice")]

@app.get("/users/{user_id}")
async def get_user(user_id: int) -> UserResponse:
    return UserResponse(id=user_id, name="Alice")

@app.post("/users", status_code=201)
async def create_user(user: User) -> UserResponse:
    return UserResponse(id=1, name=user.name)

# 中间件
from fastapi import Request

@app.middleware("http")
async def log_requests(request: Request, call_next):
    print(f"{request.method} {request.url.path}")
    response = await call_next(request)
    return response

# 错误处理
@app.exception_handler(ValueError)
async def value_error_handler(request: Request, exc: ValueError):
    return JSONResponse(status_code=400, content={"error": str(exc)})

# 运行: uvicorn main:app --reload
```

```go
// Go - net/http / gin / chi
import (
    "encoding/json"
    "net/http"
)

// 标准库
func main() {
    http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
        switch r.Method {
        case "GET":
            json.NewEncoder(w).Encode([]User{{ID: 1, Name: "Alice"}})
        case "POST":
            var user User
            json.NewDecoder(r.Body).Decode(&user)
            user.ID = 1
            w.WriteHeader(http.StatusCreated)
            json.NewEncoder(w).Encode(user)
        }
    })
    
    http.ListenAndServe(":8080", nil)
}

// Gin 框架
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    r.GET("/users", func(c *gin.Context) {
        c.JSON(200, []User{{ID: 1, Name: "Alice"}})
    })
    
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(200, User{ID: 1, Name: "Alice"})
    })
    
    r.POST("/users", func(c *gin.Context) {
        var user User
        if err := c.ShouldBindJSON(&user); err != nil {
            c.JSON(400, gin.H{"error": err.Error()})
            return
        }
        user.ID = 1
        c.JSON(201, user)
    })
    
    // 中间件
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    r.Run(":8080")
}
```

```rust
// Rust - Axum / Actix-web
use axum::{
    routing::{get, post},
    Router, Json, extract::Path,
    http::StatusCode,
};
use serde::{Deserialize, Serialize};

#[derive(Serialize)]
struct User {
    id: i32,
    name: String,
}

#[derive(Deserialize)]
struct CreateUser {
    name: String,
}

// 路由处理器
async fn get_users() -> Json<Vec<User>> {
    Json(vec![User { id: 1, name: "Alice".to_string() }])
}

async fn get_user(Path(id): Path<i32>) -> Json<User> {
    Json(User { id, name: "Alice".to_string() })
}

async fn create_user(
    Json(payload): Json<CreateUser>,
) -> (StatusCode, Json<User>) {
    let user = User { id: 1, name: payload.name };
    (StatusCode::CREATED, Json(user))
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/users", get(get_users).post(create_user))
        .route("/users/:id", get(get_user));
    
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000")
        .await
        .unwrap();
    
    axum::serve(listener, app).await.unwrap();
}
```

---

## 📝 总结对比表

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **类型系统** | 渐进式 | 动态+提示 | 静态 | 静态+所有权 |
| **空安全** | 可选链 `?.` | None检查 | nil检查 | `Option<T>` |
| **错误处理** | try/catch | try/except | 返回error | `Result<T,E>` |
| **并发** | Promise/async | asyncio/多进程 | goroutine | async/线程 |
| **包管理** | npm/yarn | pip/poetry | go mod | cargo |
| **主要用途** | Web前后端 | 脚本/AI/后端 | 云原生/CLI | 系统/WebAssembly |

## 🚀 快速开始

```bash
# TypeScript
npm init -y && npm install typescript tsx
npx tsx examples/typescript/main.ts

# Python
python3 examples/python/main.py

# Go
go run examples/go/main.go

# Rust
cargo run --manifest-path examples/rust/Cargo.toml
```

---

**提示**: 每种语言的完整可运行示例在对应的 `examples/` 目录中。
