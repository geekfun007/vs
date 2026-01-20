# TypeScript vs Python vs Go vs Rust 详解

> 四种现代编程语言的全面对比分析

## 📊 概览对比表

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **类型系统** | 静态类型（可选） | 动态类型 | 静态类型 | 静态类型 |
| **内存管理** | GC (V8) | GC | GC | 所有权系统 |
| **并发模型** | 事件循环/异步 | 多线程/协程 | Goroutines | 线程/async |
| **编译/解释** | 编译到 JS | 解释执行 | 编译型 | 编译型 |
| **学习曲线** | 中等 | 低 | 中等 | 高 |
| **性能** | 中等 | 较低 | 高 | 极高 |
| **生态成熟度** | 成熟 | 非常成熟 | 成熟 | 快速成长 |

---

## 🔷 TypeScript

### 简介
TypeScript 是 JavaScript 的超集，由微软开发，添加了可选的静态类型系统。

### 核心特点

```typescript
// 类型注解
interface User {
  id: number;
  name: string;
  email?: string; // 可选属性
}

// 泛型
function identity<T>(arg: T): T {
  return arg;
}

// 联合类型与类型守卫
type Result = Success | Error;
```

### 优势
- ✅ **渐进式类型**: 可从 JavaScript 逐步迁移
- ✅ **强大的 IDE 支持**: 智能补全、重构
- ✅ **前后端通用**: Node.js + 浏览器
- ✅ **类型推断**: 减少类型注解

### 劣势
- ❌ 需要编译步骤
- ❌ 类型系统复杂度高
- ❌ 运行时无类型检查
- ❌ 依赖 JavaScript 生态的问题

### 适用场景
- Web 前端开发
- Node.js 后端服务
- 大型团队协作项目
- 需要类型安全的 JavaScript 项目

---

## 🐍 Python

### 简介
Python 是一种高级解释型语言，以其简洁的语法和强大的生态系统著称。

### 核心特点

```python
# 简洁的语法
def greet(name: str) -> str:
    return f"Hello, {name}!"

# 列表推导
squares = [x**2 for x in range(10)]

# 装饰器
@lru_cache(maxsize=128)
def fibonacci(n):
    if n < 2:
        return n
    return fibonacci(n-1) + fibonacci(n-2)

# 上下文管理器
with open('file.txt') as f:
    content = f.read()
```

### 优势
- ✅ **语法简洁**: 易于学习和阅读
- ✅ **丰富生态**: PyPI 拥有海量包
- ✅ **科学计算**: NumPy、Pandas、TensorFlow
- ✅ **快速原型**: 开发速度快

### 劣势
- ❌ 运行速度相对较慢
- ❌ GIL 限制多线程性能
- ❌ 动态类型可能导致运行时错误
- ❌ 移动端支持有限

### 适用场景
- 数据科学与机器学习
- Web 后端 (Django/Flask/FastAPI)
- 自动化脚本
- 科学计算与研究

---

## 🔵 Go (Golang)

### 简介
Go 是 Google 开发的编译型语言，专为简单性、高效并发和可靠性设计。

### 核心特点

```go
// 简洁的语法
func greet(name string) string {
    return fmt.Sprintf("Hello, %s!", name)
}

// Goroutines - 轻量级并发
func main() {
    go func() {
        fmt.Println("并发执行")
    }()
}

// Channels - 通信机制
ch := make(chan int)
go func() { ch <- 42 }()
value := <-ch

// 接口隐式实现
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

### 优势
- ✅ **编译速度快**: 大型项目秒级编译
- ✅ **并发原语**: Goroutines 和 Channels
- ✅ **单一二进制**: 无依赖部署
- ✅ **标准库完善**: 内置 HTTP、JSON 等

### 劣势
- ❌ 缺少泛型（Go 1.18 已添加，但使用有限）
- ❌ 错误处理冗长
- ❌ 缺少高级抽象（如枚举、模式匹配）
- ❌ 包管理历史问题（已改善）

### 适用场景
- 云原生与微服务
- DevOps 工具 (Docker, K8s)
- 网络服务与 API
- CLI 工具开发

---

## 🦀 Rust

### 简介
Rust 是 Mozilla 开发的系统编程语言，专注于安全性、并发性和性能。

### 核心特点

```rust
// 所有权系统
fn main() {
    let s1 = String::from("hello");
    let s2 = s1; // s1 所有权转移给 s2
    // println!("{}", s1); // 编译错误！
}

// 借用与生命周期
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

// 模式匹配
match value {
    Some(x) => println!("Got: {}", x),
    None => println!("Nothing"),
}

// 零成本抽象
let v: Vec<i32> = (0..10).filter(|x| x % 2 == 0).collect();
```

### 优势
- ✅ **内存安全**: 无 GC，无数据竞争
- ✅ **零成本抽象**: 高级特性无性能损失
- ✅ **强大的类型系统**: 枚举、模式匹配、trait
- ✅ **优秀的工具链**: Cargo、rustfmt、clippy

### 劣势
- ❌ 学习曲线陡峭
- ❌ 编译时间较长
- ❌ 借用检查器有时过于严格
- ❌ 生态系统相对年轻

### 适用场景
- 系统编程与嵌入式
- WebAssembly
- 高性能网络服务
- 命令行工具
- 游戏引擎

---

## 📈 性能对比

```
计算密集型任务 (相对性能，Rust = 100)
┌─────────────┬──────────────────────────────────┐
│ Rust        │ ████████████████████████████ 100 │
│ Go          │ ██████████████████████████   90  │
│ TypeScript  │ █████████████                45  │
│ Python      │ ████                         15  │
└─────────────┴──────────────────────────────────┘

并发处理能力 (相对效率)
┌─────────────┬──────────────────────────────────┐
│ Rust        │ ████████████████████████████ 100 │
│ Go          │ ███████████████████████████  95  │
│ TypeScript  │ ████████████████             55  │
│ Python      │ ████████                     30  │
└─────────────┴──────────────────────────────────┘
```

---

## 🎯 选择建议

### 按项目类型选择

| 项目类型 | 推荐语言 | 原因 |
|---------|---------|------|
| Web 前端 | TypeScript | 唯一选择，类型安全 |
| Web 后端 (通用) | Go / TypeScript | 开发效率与性能平衡 |
| 数据科学/ML | Python | 生态无可替代 |
| 系统编程 | Rust | 性能与安全 |
| 微服务 | Go | 部署简单，并发优秀 |
| CLI 工具 | Go / Rust | 单二进制分发 |
| 嵌入式 | Rust | 无 GC，内存控制 |
| 快速原型 | Python | 开发速度最快 |

### 按团队情况选择

- **新手团队**: Python → Go → TypeScript → Rust
- **追求性能**: Rust > Go > TypeScript > Python
- **追求开发速度**: Python > TypeScript > Go > Rust
- **大型项目**: TypeScript / Go / Rust (强类型优势)

---

## 🔄 互操作性

```
┌─────────────────────────────────────────────────────┐
│                    互操作生态                        │
├─────────────────────────────────────────────────────┤
│  TypeScript ←→ JavaScript (原生)                    │
│  Python     ←→ C/C++ (ctypes, Cython)              │
│  Go         ←→ C (cgo)                             │
│  Rust       ←→ C (FFI), Python (PyO3), JS (wasm)   │
└─────────────────────────────────────────────────────┘
```

---

## 📝 代码风格对比

### 相同功能的实现：HTTP 服务器

**TypeScript (Express)**
```typescript
import express from 'express';
const app = express();
app.get('/', (req, res) => res.send('Hello'));
app.listen(3000);
```

**Python (FastAPI)**
```python
from fastapi import FastAPI
app = FastAPI()

@app.get("/")
def read_root():
    return {"Hello": "World"}
```

**Go (标准库)**
```go
package main
import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello"))
    })
    http.ListenAndServe(":3000", nil)
}
```

**Rust (Axum)**
```rust
use axum::{routing::get, Router};

#[tokio::main]
async fn main() {
    let app = Router::new().route("/", get(|| async { "Hello" }));
    axum::Server::bind(&"0.0.0.0:3000".parse().unwrap())
        .serve(app.into_make_service()).await.unwrap();
}
```

---

## 🔢 Math.trunc 截断函数对比

`trunc` 函数用于截断数字的小数部分，只保留整数部分（向零取整）。

### 行为说明

```
trunc(3.7)  →  3   (正数：向下取整)
trunc(-3.7) → -3   (负数：向上取整)
trunc(3.2)  →  3
trunc(-3.2) → -3
```

> **注意**: `trunc` 与 `floor` 不同！`floor(-3.7) = -4`，而 `trunc(-3.7) = -3`

### 四种语言实现

**TypeScript / JavaScript**
```typescript
// 使用 Math.trunc()
Math.trunc(3.7);    // 3
Math.trunc(-3.7);   // -3
Math.trunc(3.2);    // 3
Math.trunc(-3.2);   // -3

// 其他方式（不完全等价）
Math.floor(3.7);    // 3  (正数时等价)
Math.floor(-3.7);   // -4 (负数时不同！)
Math.ceil(-3.7);    // -3 (负数时等价)

// 位运算截断（仅适用于 32 位整数范围）
(3.7 | 0);          // 3
~~3.7;              // 3
```

**Python**
```python
import math

# 使用 math.trunc()
math.trunc(3.7)     # 3
math.trunc(-3.7)    # -3

# int() 函数效果相同
int(3.7)            # 3
int(-3.7)           # -3

# 对比 floor
math.floor(3.7)     # 3
math.floor(-3.7)    # -4  (不同！)

# 特殊方法 __trunc__
class MyNumber:
    def __init__(self, value):
        self.value = value
    def __trunc__(self):
        return int(self.value)

math.trunc(MyNumber(3.7))  # 3
```

**Go**
```go
import "math"

// 使用 math.Trunc()
math.Trunc(3.7)     // 3.0 (返回 float64)
math.Trunc(-3.7)    // -3.0

// 转换为整数
int(math.Trunc(3.7))  // 3

// 直接类型转换（效果相同）
int(3.7)            // 3
int(-3.7)           // -3

// 对比 Floor
math.Floor(3.7)     // 3.0
math.Floor(-3.7)    // -4.0 (不同！)
```

**Rust**
```rust
// 使用 trunc() 方法
(3.7_f64).trunc()      // 3.0
(-3.7_f64).trunc()     // -3.0

// 转换为整数
3.7_f64.trunc() as i32   // 3
(-3.7_f64).trunc() as i32 // -3

// 直接转换（效果相同）
3.7 as i32             // 3
-3.7 as i32            // -3

// 对比 floor
(3.7_f64).floor()      // 3.0
(-3.7_f64).floor()     // -4.0 (不同！)

// 对比 round（四舍五入）
(3.7_f64).round()      // 4.0
(-3.7_f64).round()     // -4.0
```

### 对比总结

| 语言 | 截断函数 | 返回类型 | 备注 |
|------|----------|----------|------|
| TypeScript | `Math.trunc()` | number | ES6+，位运算可替代但有限制 |
| Python | `math.trunc()` / `int()` | int | 支持自定义 `__trunc__` |
| Go | `math.Trunc()` | float64 | 需手动转 int |
| Rust | `.trunc()` | 原浮点类型 | 方法调用，类型安全 |

### 边界情况处理

```
┌──────────────┬────────────┬────────────┬────────────┬────────────┐
│ 输入         │ TypeScript │ Python     │ Go         │ Rust       │
├──────────────┼────────────┼────────────┼────────────┼────────────┤
│ Infinity     │ Infinity   │ Error      │ +Inf       │ inf        │
│ -Infinity    │ -Infinity  │ Error      │ -Inf       │ -inf       │
│ NaN          │ NaN        │ Error      │ NaN        │ NaN        │
│ 0.0          │ 0          │ 0          │ 0.0        │ 0.0        │
│ -0.0         │ -0         │ 0          │ -0.0       │ -0.0       │
└──────────────┴────────────┴────────────┴────────────┴────────────┘
```

---

## 🎭 装饰器模式对比

装饰器是一种在不修改原始代码的情况下增强函数/类行为的模式。

### 概念对比

| 语言 | 原生支持 | 实现方式 | 语法糖 |
|------|---------|----------|--------|
| TypeScript | ✅ (实验性) | 装饰器 | `@decorator` |
| Python | ✅ | 装饰器 | `@decorator` |
| Go | ❌ | 高阶函数/中间件 | 无 |
| Rust | ❌ | 过程宏/trait | `#[attribute]` |

### Python 装饰器（最完善）

```python
from functools import wraps
import time

# 基础装饰器
def log_calls(func):
    @wraps(func)  # 保留原函数元信息
    def wrapper(*args, **kwargs):
        print(f"调用 {func.__name__}")
        return func(*args, **kwargs)
    return wrapper

@log_calls
def greet(name):
    return f"Hello, {name}!"

# 带参数的装饰器
def retry(times=3, delay=1):
    def decorator(func):
        @wraps(func)
        def wrapper(*args, **kwargs):
            for i in range(times):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    if i == times - 1:
                        raise
                    time.sleep(delay)
        return wrapper
    return decorator

@retry(times=5, delay=2)
def fetch_data(url):
    # 可能失败的操作
    pass

# 类装饰器
def singleton(cls):
    instances = {}
    @wraps(cls)
    def get_instance(*args, **kwargs):
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]
    return get_instance

@singleton
class Database:
    pass

# 多个装饰器（从下往上执行）
@log_calls
@retry(times=3)
def risky_operation():
    pass
```

### TypeScript 装饰器

```typescript
// 需要在 tsconfig.json 中启用: "experimentalDecorators": true

// 方法装饰器
function log(
  target: any,
  propertyKey: string,
  descriptor: PropertyDescriptor
) {
  const original = descriptor.value;
  descriptor.value = function (...args: any[]) {
    console.log(`调用 ${propertyKey}`);
    return original.apply(this, args);
  };
  return descriptor;
}

class UserService {
  @log
  getUser(id: number) {
    return { id, name: "John" };
  }
}

// 类装饰器
function sealed(constructor: Function) {
  Object.seal(constructor);
  Object.seal(constructor.prototype);
}

@sealed
class BankAccount {
  balance: number = 0;
}

// 属性装饰器
function readonly(target: any, key: string) {
  Object.defineProperty(target, key, {
    writable: false,
  });
}

class Config {
  @readonly
  apiKey = "secret";
}

// 参数装饰器
function required(
  target: any,
  propertyKey: string,
  parameterIndex: number
) {
  // 验证逻辑
}

class Validator {
  validate(@required name: string) {}
}

// 装饰器工厂（带参数）
function debounce(ms: number) {
  return function (
    target: any,
    key: string,
    descriptor: PropertyDescriptor
  ) {
    let timeout: NodeJS.Timeout;
    const original = descriptor.value;
    descriptor.value = function (...args: any[]) {
      clearTimeout(timeout);
      timeout = setTimeout(() => original.apply(this, args), ms);
    };
    return descriptor;
  };
}

class SearchBox {
  @debounce(300)
  search(query: string) {
    console.log("Searching:", query);
  }
}
```

### Go 高阶函数模式

```go
package main

import (
    "fmt"
    "log"
    "net/http"
    "time"
)

// 函数类型定义
type Handler func(w http.ResponseWriter, r *http.Request)

// 日志中间件
func withLogging(h Handler) Handler {
    return func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        h(w, r)
        log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
    }
}

// 认证中间件
func withAuth(h Handler) Handler {
    return func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", 401)
            return
        }
        h(w, r)
    }
}

// 重试装饰器
func withRetry(times int, fn func() error) func() error {
    return func() error {
        var err error
        for i := 0; i < times; i++ {
            if err = fn(); err == nil {
                return nil
            }
            time.Sleep(time.Second * time.Duration(i+1))
        }
        return err
    }
}

// 计时装饰器
func measure[T any](name string, fn func() T) T {
    start := time.Now()
    result := fn()
    fmt.Printf("%s took %v\n", name, time.Since(start))
    return result
}

// 使用
func main() {
    handler := withLogging(withAuth(func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello!"))
    }))
    
    http.HandleFunc("/", handler)
    
    // 泛型版本使用
    result := measure("computation", func() int {
        // 耗时操作
        return 42
    })
}
```

### Rust 过程宏与 Trait

```rust
// 属性宏（类似装饰器，需要定义过程宏）
use proc_macro::TokenStream;

// 在 proc-macro crate 中定义
#[proc_macro_attribute]
pub fn log_calls(_attr: TokenStream, item: TokenStream) -> TokenStream {
    // 解析并转换代码
    item
}

// 使用属性宏
#[log_calls]
fn my_function() {
    println!("Hello!");
}

// ----------------------------
// 常用的高阶函数模式
// ----------------------------

use std::time::Instant;

// 计时装饰器
fn measure<F, T>(name: &str, f: F) -> T
where
    F: FnOnce() -> T,
{
    let start = Instant::now();
    let result = f();
    println!("{} took {:?}", name, start.elapsed());
    result
}

// 重试装饰器
fn with_retry<F, T, E>(times: usize, mut f: F) -> Result<T, E>
where
    F: FnMut() -> Result<T, E>,
{
    let mut last_err = None;
    for _ in 0..times {
        match f() {
            Ok(v) => return Ok(v),
            Err(e) => last_err = Some(e),
        }
    }
    Err(last_err.unwrap())
}

// 使用
fn main() {
    let result = measure("calculation", || {
        // 耗时操作
        42
    });

    let data = with_retry(3, || {
        fetch_from_api()
    });
}

// ----------------------------
// 使用 Trait 实现装饰器模式
// ----------------------------

trait Drawable {
    fn draw(&self);
}

struct Circle {
    radius: f64,
}

impl Drawable for Circle {
    fn draw(&self) {
        println!("Drawing circle with radius {}", self.radius);
    }
}

// 装饰器结构体
struct BorderDecorator<T: Drawable> {
    inner: T,
    border_width: u32,
}

impl<T: Drawable> Drawable for BorderDecorator<T> {
    fn draw(&self) {
        println!("Drawing border (width: {})", self.border_width);
        self.inner.draw();
    }
}

// 使用
fn main() {
    let circle = Circle { radius: 5.0 };
    let bordered = BorderDecorator {
        inner: circle,
        border_width: 2,
    };
    bordered.draw();
}

// ----------------------------
// 常见的内置属性宏
// ----------------------------
#[derive(Debug, Clone, PartialEq)]  // 自动派生 trait
struct Point { x: i32, y: i32 }

#[cfg(test)]  // 条件编译
mod tests {
    #[test]
    fn it_works() {}
}

#[inline]  // 内联提示
fn fast_function() {}

#[deprecated(since = "1.0.0", note = "use new_function instead")]
fn old_function() {}
```

### 装饰器模式对比总结

| 特性 | Python | TypeScript | Go | Rust |
|------|--------|------------|-----|------|
| 语法简洁度 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ |
| 类型安全 | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| 运行时灵活性 | ⭐⭐⭐⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐ | ⭐⭐ |
| 编译时优化 | ⭐ | ⭐⭐ | ⭐⭐⭐⭐ | ⭐⭐⭐⭐⭐ |
| 学习曲线 | 低 | 中等 | 中等 | 高 |

### 常见应用场景

```
┌────────────────┬─────────────────────────────────────┐
│ 应用场景        │ 最佳选择                             │
├────────────────┼─────────────────────────────────────┤
│ Web 中间件      │ Go (标准模式) / Python (Flask)      │
│ AOP 切面编程    │ Python / TypeScript                │
│ 权限验证        │ 所有语言都适用                       │
│ 日志记录        │ 所有语言都适用                       │
│ 缓存/记忆化     │ Python (@lru_cache)                │
│ 依赖注入        │ TypeScript (NestJS)                │
│ 编译时代码生成  │ Rust (过程宏)                       │
└────────────────┴─────────────────────────────────────┘
```

---

## 📚 总结

| 语言 | 一句话总结 |
|------|-----------|
| **TypeScript** | JavaScript 的类型安全升级版，Web 开发首选 |
| **Python** | 简单优雅，数据科学无敌，适合快速开发 |
| **Go** | 简单高效，云原生时代的工程语言 |
| **Rust** | 性能与安全的极致追求，系统编程新标杆 |

> **没有最好的语言，只有最适合的场景。** 选择语言时，应综合考虑项目需求、团队技能、生态系统和长期维护成本。
