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

## 📅 日期操作对比

### 日期类型概览

| 语言 | 主要类型 | 时区支持 | 精度 |
|------|----------|----------|------|
| TypeScript | `Date` | 有限 | 毫秒 |
| Python | `datetime` | 完善 (pytz/zoneinfo) | 微秒 |
| Go | `time.Time` | 内置完善 | 纳秒 |
| Rust | `chrono` crate | 完善 | 纳秒 |

### 创建日期

**TypeScript**
```typescript
// 当前时间
const now = new Date();

// 指定日期
const date1 = new Date('2024-03-15');
const date2 = new Date(2024, 2, 15);  // 月份从 0 开始！
const date3 = new Date(2024, 2, 15, 10, 30, 0);

// 时间戳
const fromTimestamp = new Date(1710489600000);

// ISO 字符串
const iso = new Date('2024-03-15T10:30:00Z');
```

**Python**
```python
from datetime import datetime, date, timedelta
from zoneinfo import ZoneInfo  # Python 3.9+

# 当前时间
now = datetime.now()
utc_now = datetime.utcnow()  # 已弃用
utc_now = datetime.now(ZoneInfo('UTC'))  # 推荐

# 指定日期
date1 = datetime(2024, 3, 15)
date2 = datetime(2024, 3, 15, 10, 30, 0)
date_only = date(2024, 3, 15)

# 时间戳
from_timestamp = datetime.fromtimestamp(1710489600)

# 字符串解析
parsed = datetime.strptime('2024-03-15', '%Y-%m-%d')
iso_parsed = datetime.fromisoformat('2024-03-15T10:30:00')
```

**Go**
```go
import "time"

// 当前时间
now := time.Now()
utcNow := time.Now().UTC()

// 指定日期 (年, 月, 日, 时, 分, 秒, 纳秒, 时区)
date1 := time.Date(2024, time.March, 15, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2024, 3, 15, 10, 30, 0, 0, time.Local)

// 时间戳
fromUnix := time.Unix(1710489600, 0)
fromUnixMilli := time.UnixMilli(1710489600000)

// 字符串解析 (Go 使用参考时间: 2006-01-02 15:04:05)
parsed, _ := time.Parse("2006-01-02", "2024-03-15")
parsed2, _ := time.Parse(time.RFC3339, "2024-03-15T10:30:00Z")
```

**Rust**
```rust
use chrono::{DateTime, Utc, Local, NaiveDate, NaiveDateTime, TimeZone};

// 当前时间
let now: DateTime<Utc> = Utc::now();
let local_now: DateTime<Local> = Local::now();

// 指定日期
let date1 = NaiveDate::from_ymd_opt(2024, 3, 15).unwrap();
let datetime1 = NaiveDateTime::new(
    NaiveDate::from_ymd_opt(2024, 3, 15).unwrap(),
    chrono::NaiveTime::from_hms_opt(10, 30, 0).unwrap()
);

// 带时区
let utc_date = Utc.with_ymd_and_hms(2024, 3, 15, 10, 30, 0).unwrap();

// 时间戳
let from_timestamp = DateTime::from_timestamp(1710489600, 0).unwrap();

// 字符串解析
let parsed = NaiveDate::parse_from_str("2024-03-15", "%Y-%m-%d").unwrap();
let parsed_dt = DateTime::parse_from_rfc3339("2024-03-15T10:30:00Z").unwrap();
```

### 日期格式化

**TypeScript**
```typescript
const date = new Date('2024-03-15T10:30:00');

// 内置方法
date.toISOString();       // "2024-03-15T10:30:00.000Z"
date.toLocaleDateString(); // "3/15/2024" (依赖 locale)
date.toLocaleString('zh-CN'); // "2024/3/15 10:30:00"

// Intl API (推荐)
const formatter = new Intl.DateTimeFormat('zh-CN', {
  year: 'numeric',
  month: '2-digit',
  day: '2-digit',
  hour: '2-digit',
  minute: '2-digit'
});
formatter.format(date);  // "2024/03/15 10:30"

// 手动格式化
const pad = (n: number) => n.toString().padStart(2, '0');
`${date.getFullYear()}-${pad(date.getMonth()+1)}-${pad(date.getDate())}`;
```

**Python**
```python
from datetime import datetime

dt = datetime(2024, 3, 15, 10, 30, 0)

# strftime 格式化
dt.strftime('%Y-%m-%d')           # "2024-03-15"
dt.strftime('%Y年%m月%d日')        # "2024年03月15日"
dt.strftime('%Y-%m-%d %H:%M:%S')  # "2024-03-15 10:30:00"
dt.strftime('%A, %B %d, %Y')      # "Friday, March 15, 2024"

# ISO 格式
dt.isoformat()                    # "2024-03-15T10:30:00"

# f-string
f"{dt:%Y/%m/%d}"                  # "2024/03/15"
```

**Go**
```go
t := time.Date(2024, 3, 15, 10, 30, 0, 0, time.UTC)

// Go 使用参考时间格式: Mon Jan 2 15:04:05 MST 2006
t.Format("2006-01-02")           // "2024-03-15"
t.Format("2006年01月02日")        // "2024年03月15日"
t.Format("2006-01-02 15:04:05")  // "2024-03-15 10:30:00"
t.Format(time.RFC3339)           // "2024-03-15T10:30:00Z"
t.Format("Monday, January 2, 2006") // "Friday, March 15, 2024"

// 预定义格式
t.Format(time.Kitchen)           // "10:30AM"
t.Format(time.RFC822)            // "15 Mar 24 10:30 UTC"
```

**Rust**
```rust
use chrono::{Utc, TimeZone};

let dt = Utc.with_ymd_and_hms(2024, 3, 15, 10, 30, 0).unwrap();

// format! 宏
dt.format("%Y-%m-%d").to_string()           // "2024-03-15"
dt.format("%Y年%m月%d日").to_string()        // "2024年03月15日"
dt.format("%Y-%m-%d %H:%M:%S").to_string()  // "2024-03-15 10:30:00"
dt.format("%A, %B %d, %Y").to_string()      // "Friday, March 15, 2024"

// RFC 格式
dt.to_rfc3339()                             // "2024-03-15T10:30:00+00:00"
dt.to_rfc2822()                             // "Fri, 15 Mar 2024 10:30:00 +0000"
```

### 日期差计算

**TypeScript**
```typescript
const date1 = new Date('2024-03-15');
const date2 = new Date('2024-03-20');

// 毫秒差
const diffMs = date2.getTime() - date1.getTime();

// 转换为各单位
const diffSeconds = diffMs / 1000;
const diffMinutes = diffMs / (1000 * 60);
const diffHours = diffMs / (1000 * 60 * 60);
const diffDays = diffMs / (1000 * 60 * 60 * 24);  // 5

// 添加天数
const addDays = (date: Date, days: number): Date => {
  const result = new Date(date);
  result.setDate(result.getDate() + days);
  return result;
};

// 添加月份（注意月末边界）
const addMonths = (date: Date, months: number): Date => {
  const result = new Date(date);
  result.setMonth(result.getMonth() + months);
  return result;
};
```

**Python**
```python
from datetime import datetime, timedelta
from dateutil.relativedelta import relativedelta  # pip install python-dateutil

date1 = datetime(2024, 3, 15)
date2 = datetime(2024, 3, 20)

# timedelta 差值
diff = date2 - date1
diff.days          # 5
diff.total_seconds()  # 432000.0

# 添加时间
date1 + timedelta(days=5)        # 2024-03-20
date1 + timedelta(hours=48)      # 2024-03-17
date1 + timedelta(weeks=2)       # 2024-03-29

# relativedelta (处理月/年更准确)
date1 + relativedelta(months=1)  # 2024-04-15
date1 + relativedelta(years=1)   # 2025-03-15
date1 + relativedelta(months=1, days=5)  # 2024-04-20

# 计算精确差值
rd = relativedelta(date2, date1)
rd.days  # 5
```

**Go**
```go
date1 := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)

// Duration 差值
diff := date2.Sub(date1)
diff.Hours()    // 120
diff.Minutes()  // 7200
diff.Seconds()  // 432000

// 天数
days := int(diff.Hours() / 24)  // 5

// 添加时间
date1.Add(5 * 24 * time.Hour)      // 添加 5 天
date1.Add(time.Hour * 48)          // 添加 48 小时
date1.AddDate(1, 2, 3)             // 添加 1年2月3天

// 使用 AddDate
date1.AddDate(0, 1, 0)  // 添加 1 个月 -> 2024-04-15
date1.AddDate(1, 0, 0)  // 添加 1 年 -> 2025-03-15
```

**Rust**
```rust
use chrono::{Duration, Utc, TimeZone, Datelike};

let date1 = Utc.with_ymd_and_hms(2024, 3, 15, 0, 0, 0).unwrap();
let date2 = Utc.with_ymd_and_hms(2024, 3, 20, 0, 0, 0).unwrap();

// Duration 差值
let diff = date2.signed_duration_since(date1);
diff.num_days()      // 5
diff.num_hours()     // 120
diff.num_seconds()   // 432000

// 添加时间
date1 + Duration::days(5)       // 添加 5 天
date1 + Duration::hours(48)     // 添加 48 小时
date1 + Duration::weeks(2)      // 添加 2 周

// 添加月份 (chrono 需要手动处理)
use chrono::Months;
date1.checked_add_months(Months::new(1)).unwrap()  // 2024-04-15

// 年份操作
date1.with_year(2025).unwrap()  // 2025-03-15
```

### 日期比较

**TypeScript**
```typescript
const date1 = new Date('2024-03-15');
const date2 = new Date('2024-03-20');
const date3 = new Date('2024-03-15');

// 比较时间戳
date1.getTime() < date2.getTime()   // true
date1.getTime() === date3.getTime() // true

// 直接比较 (仅 < > 可用)
date1 < date2   // true
date1 > date2   // false
// date1 === date3  // false! (对象引用比较)

// 判断是否同一天
const isSameDay = (d1: Date, d2: Date): boolean =>
  d1.getFullYear() === d2.getFullYear() &&
  d1.getMonth() === d2.getMonth() &&
  d1.getDate() === d2.getDate();

// 范围判断
const isInRange = (date: Date, start: Date, end: Date): boolean =>
  date >= start && date <= end;
```

**Python**
```python
from datetime import datetime

date1 = datetime(2024, 3, 15)
date2 = datetime(2024, 3, 20)
date3 = datetime(2024, 3, 15)

# 直接比较
date1 < date2   # True
date1 > date2   # False
date1 == date3  # True
date1 != date2  # True
date1 <= date2  # True

# 范围判断
start = datetime(2024, 3, 1)
end = datetime(2024, 3, 31)
start <= date1 <= end  # True (Python 支持链式比较)

# 判断是否同一天
date1.date() == date3.date()  # True

# 最早/最晚
min(date1, date2, date3)  # date1
max(date1, date2, date3)  # date2

# 排序
dates = [date2, date1, date3]
sorted(dates)  # [date1, date3, date2]
```

**Go**
```go
date1 := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)
date2 := time.Date(2024, 3, 20, 0, 0, 0, 0, time.UTC)
date3 := time.Date(2024, 3, 15, 0, 0, 0, 0, time.UTC)

// 比较方法
date1.Before(date2)  // true
date1.After(date2)   // false
date1.Equal(date3)   // true

// 范围判断
start := time.Date(2024, 3, 1, 0, 0, 0, 0, time.UTC)
end := time.Date(2024, 3, 31, 0, 0, 0, 0, time.UTC)
inRange := (date1.After(start) || date1.Equal(start)) && 
           (date1.Before(end) || date1.Equal(end))

// 判断是否为零值
date1.IsZero()  // false

// 同一天判断
func isSameDay(t1, t2 time.Time) bool {
    y1, m1, d1 := t1.Date()
    y2, m2, d2 := t2.Date()
    return y1 == y2 && m1 == m2 && d1 == d2
}
```

**Rust**
```rust
use chrono::{Utc, TimeZone};

let date1 = Utc.with_ymd_and_hms(2024, 3, 15, 0, 0, 0).unwrap();
let date2 = Utc.with_ymd_and_hms(2024, 3, 20, 0, 0, 0).unwrap();
let date3 = Utc.with_ymd_and_hms(2024, 3, 15, 0, 0, 0).unwrap();

// 直接比较 (实现了 Ord trait)
date1 < date2   // true
date1 > date2   // false
date1 == date3  // true
date1 != date2  // true

// 范围判断
let start = Utc.with_ymd_and_hms(2024, 3, 1, 0, 0, 0).unwrap();
let end = Utc.with_ymd_and_hms(2024, 3, 31, 0, 0, 0).unwrap();
let in_range = date1 >= start && date1 <= end;  // true

// 最早/最晚
let earliest = date1.min(date2);
let latest = date1.max(date2);

// 排序
let mut dates = vec![date2, date1, date3];
dates.sort();  // [date1, date3, date2]

// 同一天判断
date1.date_naive() == date3.date_naive()  // true
```

### 日期格式化速查表

```
┌──────────┬─────────────┬─────────────┬─────────────┬─────────────┐
│ 含义     │ TypeScript  │ Python      │ Go          │ Rust        │
├──────────┼─────────────┼─────────────┼─────────────┼─────────────┤
│ 四位年份  │ getFullYear │ %Y          │ 2006        │ %Y          │
│ 两位年份  │ -           │ %y          │ 06          │ %y          │
│ 月(01-12)│ getMonth+1  │ %m          │ 01          │ %m          │
│ 日(01-31)│ getDate     │ %d          │ 02          │ %d          │
│ 时(00-23)│ getHours    │ %H          │ 15          │ %H          │
│ 分(00-59)│ getMinutes  │ %M          │ 04          │ %M          │
│ 秒(00-59)│ getSeconds  │ %S          │ 05          │ %S          │
│ 星期名称  │ -           │ %A          │ Monday      │ %A          │
│ 月份名称  │ -           │ %B          │ January     │ %B          │
└──────────┴─────────────┴─────────────┴─────────────┴─────────────┘
```

---

## 🔤 字符串操作对比

### 字符串类型概览

| 语言 | 类型 | 可变性 | 编码 | 索引单位 |
|------|------|--------|------|----------|
| TypeScript | `string` | 不可变 | UTF-16 | 代码单元 |
| Python | `str` | 不可变 | UTF-8 | Unicode 码点 |
| Go | `string` | 不可变 | UTF-8 | 字节 |
| Rust | `String` / `&str` | String 可变 | UTF-8 | 字节 |

### 基本操作

**TypeScript**
```typescript
const s = "Hello, World!";

// 长度
s.length;           // 13 (UTF-16 代码单元数)

// 访问字符
s[0];               // "H"
s.charAt(0);        // "H"
s.at(-1);           // "!" (ES2022)

// 切片
s.slice(0, 5);      // "Hello"
s.slice(-6);        // "World!"
s.substring(0, 5);  // "Hello"

// 查找
s.indexOf("o");     // 4
s.lastIndexOf("o"); // 8
s.includes("World"); // true
s.startsWith("Hello"); // true
s.endsWith("!");    // true

// 拼接
"Hello" + " " + "World";
`Hello ${name}`;    // 模板字符串
["Hello", "World"].join(" ");
```

**Python**
```python
s = "Hello, World!"

# 长度
len(s)              # 13

# 访问字符
s[0]                # "H"
s[-1]               # "!"

# 切片
s[0:5]              # "Hello"
s[-6:]              # "World!"
s[::2]              # "Hlo ol!" (步长为2)
s[::-1]             # "!dlroW ,olleH" (反转)

# 查找
s.index("o")        # 4 (找不到抛异常)
s.find("o")         # 4 (找不到返回 -1)
s.rfind("o")        # 8
"World" in s        # True
s.startswith("Hello")  # True
s.endswith("!")     # True
s.count("o")        # 2

# 拼接
"Hello" + " " + "World"
f"Hello {name}"     # f-string
" ".join(["Hello", "World"])
```

**Go**
```go
s := "Hello, World!"

// 长度
len(s)              // 13 (字节数)
utf8.RuneCountInString(s)  // 13 (字符数)

// 访问 (返回字节)
s[0]                // 72 (byte 'H')
string(s[0])        // "H"

// 切片
s[0:5]              // "Hello"
s[7:]               // "World!"

// 查找
strings.Index(s, "o")      // 4
strings.LastIndex(s, "o")  // 8
strings.Contains(s, "World") // true
strings.HasPrefix(s, "Hello") // true
strings.HasSuffix(s, "!")  // true
strings.Count(s, "o")      // 2

// 拼接
"Hello" + " " + "World"
fmt.Sprintf("Hello %s", name)
strings.Join([]string{"Hello", "World"}, " ")

// 构建器 (高效拼接)
var builder strings.Builder
builder.WriteString("Hello")
builder.WriteString(" World")
builder.String()
```

**Rust**
```rust
let s = String::from("Hello, World!");
let s_slice: &str = "Hello, World!";

// 长度
s.len()             // 13 (字节数)
s.chars().count()   // 13 (字符数)

// 访问字符 (返回迭代器)
s.chars().nth(0)    // Some('H')
s.chars().last()    // Some('!')

// 切片 (必须是有效 UTF-8 边界)
&s[0..5]            // "Hello"
&s[7..]             // "World!"

// 查找
s.find("o")         // Some(4)
s.rfind("o")        // Some(8)
s.contains("World") // true
s.starts_with("Hello") // true
s.ends_with("!")    // true
s.matches("o").count() // 2

// 拼接
format!("{} {}", "Hello", "World")
let mut owned = String::from("Hello");
owned.push_str(" World");
["Hello", "World"].join(" ")
```

### 字符串转换

**TypeScript**
```typescript
const s = "  Hello, World!  ";

// 大小写
s.toUpperCase();    // "  HELLO, WORLD!  "
s.toLowerCase();    // "  hello, world!  "

// 去空白
s.trim();           // "Hello, World!"
s.trimStart();      // "Hello, World!  "
s.trimEnd();        // "  Hello, World!"

// 替换
s.replace("World", "TypeScript");      // 替换第一个
s.replaceAll("l", "L");                // 替换所有
s.replace(/o/g, "0");                  // 正则替换

// 分割
"a,b,c".split(",");         // ["a", "b", "c"]
"a  b  c".split(/\s+/);     // ["a", "b", "c"]

// 填充
"5".padStart(3, "0");       // "005"
"5".padEnd(3, "0");         // "500"

// 重复
"ab".repeat(3);             // "ababab"
```

**Python**
```python
s = "  Hello, World!  "

# 大小写
s.upper()           # "  HELLO, WORLD!  "
s.lower()           # "  hello, world!  "
s.capitalize()      # "  hello, world!  "
s.title()           # "  Hello, World!  "
s.swapcase()        # "  hELLO, wORLD!  "

# 去空白
s.strip()           # "Hello, World!"
s.lstrip()          # "Hello, World!  "
s.rstrip()          # "  Hello, World!"
s.strip(" !")       # "Hello, World"

# 替换
s.replace("World", "Python")           # 替换所有
s.replace("l", "L", 1)                 # 只替换第一个

# 分割
"a,b,c".split(",")          # ["a", "b", "c"]
"a  b  c".split()           # ["a", "b", "c"] (默认按空白)
"a,b,c".split(",", 1)       # ["a", "b,c"] (限制分割次数)
"a\nb\nc".splitlines()      # ["a", "b", "c"]

# 填充
"5".zfill(3)                # "005"
"5".rjust(3, "0")           # "005"
"5".ljust(3, "0")           # "500"
"5".center(5, "-")          # "--5--"

# 重复
"ab" * 3                    # "ababab"
```

**Go**
```go
import "strings"

s := "  Hello, World!  "

// 大小写
strings.ToUpper(s)          // "  HELLO, WORLD!  "
strings.ToLower(s)          // "  hello, world!  "
strings.Title(s)            // "  Hello, World!  " (已弃用)
cases.Title(language.English).String(s)  // golang.org/x/text

// 去空白
strings.TrimSpace(s)        // "Hello, World!"
strings.TrimLeft(s, " ")    // "Hello, World!  "
strings.TrimRight(s, " ")   // "  Hello, World!"
strings.Trim(s, " !")       // "Hello, World"
strings.TrimPrefix(s, "  ") // "Hello, World!  "
strings.TrimSuffix(s, "  ") // "  Hello, World!"

// 替换
strings.Replace(s, "World", "Go", 1)   // 替换第一个
strings.ReplaceAll(s, "l", "L")        // 替换所有

// 分割
strings.Split("a,b,c", ",")           // ["a", "b", "c"]
strings.Fields("a  b  c")             // ["a", "b", "c"]
strings.SplitN("a,b,c", ",", 2)       // ["a", "b,c"]

// 填充 (需手动实现)
fmt.Sprintf("%03s", "5")              // "  5" (不完全等价)
fmt.Sprintf("%03d", 5)                // "005"

// 重复
strings.Repeat("ab", 3)               // "ababab"
```

**Rust**
```rust
let s = "  Hello, World!  ";

// 大小写
s.to_uppercase()            // "  HELLO, WORLD!  "
s.to_lowercase()            // "  hello, world!  "

// 去空白
s.trim()                    // "Hello, World!"
s.trim_start()              // "Hello, World!  "
s.trim_end()                // "  Hello, World!"
s.trim_matches(|c| c == ' ' || c == '!')  // "Hello, World"
s.strip_prefix("  ")        // Some("Hello, World!  ")
s.strip_suffix("  ")        // Some("  Hello, World!")

// 替换
s.replace("World", "Rust")             // 替换所有
s.replacen("l", "L", 1)                // 只替换第一个

// 分割
"a,b,c".split(',').collect::<Vec<_>>()  // ["a", "b", "c"]
"a  b  c".split_whitespace().collect::<Vec<_>>()  // ["a", "b", "c"]
"a,b,c".splitn(2, ',').collect::<Vec<_>>()  // ["a", "b,c"]

// 填充
format!("{:0>3}", "5")      // "005"
format!("{:0<3}", "5")      // "500"
format!("{:-^5}", "5")      // "--5--"

// 重复
"ab".repeat(3)              // "ababab"
```

### 字符串解析

**TypeScript**
```typescript
// 数字解析
parseInt("42");           // 42
parseInt("42.5");         // 42
parseInt("1010", 2);      // 10 (二进制)
parseInt("ff", 16);       // 255 (十六进制)
parseFloat("3.14");       // 3.14
Number("42");             // 42
+"42";                    // 42 (一元加号)

// JSON 解析
JSON.parse('{"name":"John","age":30}');
JSON.stringify({ name: "John", age: 30 });

// URL 解析
const url = new URL("https://example.com/path?name=john");
url.hostname;             // "example.com"
url.pathname;             // "/path"
url.searchParams.get("name");  // "john"

// 正则提取
const match = "age: 25".match(/age: (\d+)/);
match?.[1];               // "25"

// 模板解析
const template = "Hello, {name}!";
template.replace("{name}", "World");
```

**Python**
```python
# 数字解析
int("42")                 # 42
int("1010", 2)            # 10 (二进制)
int("ff", 16)             # 255 (十六进制)
float("3.14")             # 3.14

# JSON 解析
import json
json.loads('{"name":"John","age":30}')
json.dumps({"name": "John", "age": 30})

# URL 解析
from urllib.parse import urlparse, parse_qs
url = urlparse("https://example.com/path?name=john")
url.hostname              # "example.com"
url.path                  # "/path"
parse_qs(url.query)       # {"name": ["john"]}

# 正则提取
import re
match = re.search(r'age: (\d+)', "age: 25")
match.group(1)            # "25"
re.findall(r'\d+', "a1b2c3")  # ["1", "2", "3"]

# 格式化字符串解析
from string import Template
t = Template("Hello, $name!")
t.substitute(name="World")

# 结构化解析
name, age = "John,30".split(",")
```

**Go**
```go
import (
    "encoding/json"
    "fmt"
    "net/url"
    "regexp"
    "strconv"
)

// 数字解析
strconv.Atoi("42")                    // 42, nil
strconv.ParseInt("1010", 2, 64)       // 10 (二进制)
strconv.ParseInt("ff", 16, 64)        // 255 (十六进制)
strconv.ParseFloat("3.14", 64)        // 3.14

// JSON 解析
type Person struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}
var p Person
json.Unmarshal([]byte(`{"name":"John","age":30}`), &p)
data, _ := json.Marshal(p)

// URL 解析
u, _ := url.Parse("https://example.com/path?name=john")
u.Hostname()              // "example.com"
u.Path                    // "/path"
u.Query().Get("name")     // "john"

// 正则提取
re := regexp.MustCompile(`age: (\d+)`)
match := re.FindStringSubmatch("age: 25")
match[1]                  // "25"
re.FindAllString("a1b2c3", -1)  // 配合模式

// 格式化解析
fmt.Sscanf("John 30", "%s %d", &name, &age)
```

**Rust**
```rust
use std::str::FromStr;

// 数字解析
"42".parse::<i32>().unwrap()              // 42
i32::from_str_radix("1010", 2).unwrap()   // 10 (二进制)
i32::from_str_radix("ff", 16).unwrap()    // 255 (十六进制)
"3.14".parse::<f64>().unwrap()            // 3.14

// JSON 解析 (serde_json)
use serde::{Deserialize, Serialize};
use serde_json;

#[derive(Serialize, Deserialize)]
struct Person {
    name: String,
    age: u32,
}
let p: Person = serde_json::from_str(r#"{"name":"John","age":30}"#).unwrap();
let json = serde_json::to_string(&p).unwrap();

// URL 解析 (url crate)
use url::Url;
let u = Url::parse("https://example.com/path?name=john").unwrap();
u.host_str()              // Some("example.com")
u.path()                  // "/path"
u.query_pairs().find(|(k, _)| k == "name")  // ("name", "john")

// 正则提取 (regex crate)
use regex::Regex;
let re = Regex::new(r"age: (\d+)").unwrap();
let caps = re.captures("age: 25").unwrap();
&caps[1]                  // "25"
re.find_iter("a1b2c3").map(|m| m.as_str()).collect::<Vec<_>>()

// 模式匹配解析
let parts: Vec<&str> = "John,30".split(',').collect();
let (name, age) = (parts[0], parts[1].parse::<i32>().unwrap());
```

### 字符串方法速查表

```
┌────────────────┬───────────────────┬─────────────────┬─────────────────┬─────────────────┐
│ 操作           │ TypeScript        │ Python          │ Go              │ Rust            │
├────────────────┼───────────────────┼─────────────────┼─────────────────┼─────────────────┤
│ 长度           │ .length           │ len()           │ len()           │ .len()          │
│ 大写           │ .toUpperCase()    │ .upper()        │ strings.ToUpper │ .to_uppercase() │
│ 小写           │ .toLowerCase()    │ .lower()        │ strings.ToLower │ .to_lowercase() │
│ 去空白         │ .trim()           │ .strip()        │ strings.TrimSpace│ .trim()        │
│ 分割           │ .split()          │ .split()        │ strings.Split   │ .split()        │
│ 连接           │ .join()           │ .join()         │ strings.Join    │ .join()         │
│ 替换           │ .replace()        │ .replace()      │ strings.Replace │ .replace()      │
│ 包含           │ .includes()       │ in              │ strings.Contains│ .contains()     │
│ 开始于         │ .startsWith()     │ .startswith()   │ strings.HasPrefix│.starts_with()  │
│ 结束于         │ .endsWith()       │ .endswith()     │ strings.HasSuffix│.ends_with()    │
│ 查找           │ .indexOf()        │ .find()         │ strings.Index   │ .find()         │
│ 重复           │ .repeat()         │ * 运算符        │ strings.Repeat  │ .repeat()       │
│ 填充           │ .padStart()       │ .zfill()/.rjust │ fmt.Sprintf     │ format!         │
└────────────────┴───────────────────┴─────────────────┴─────────────────┴─────────────────┘
```

### Unicode 处理注意事项

```typescript
// TypeScript: UTF-16 代码单元
"😀".length;              // 2 (代理对)
[..."😀"].length;         // 1 (展开为码点)
"café".normalize("NFD");  // 规范化
```

```python
# Python: Unicode 码点
len("😀")                 # 1
"😀".encode('utf-8')      # b'\xf0\x9f\x98\x80' (4字节)
import unicodedata
unicodedata.normalize('NFD', 'café')
```

```go
// Go: UTF-8 字节
len("😀")                           // 4 (字节)
utf8.RuneCountInString("😀")        // 1 (符文)
[]rune("😀")                        // [128512]
for _, r := range "Hello😀" { }     // 遍历符文
```

```rust
// Rust: UTF-8 字节
"😀".len()                // 4 (字节)
"😀".chars().count()      // 1 (字符)
"😀".as_bytes()           // [240, 159, 152, 128]
for c in "Hello😀".chars() { }  // 遍历字符
```

---

## ⚡ 异步与并发

### 并发模型概览

| 语言 | 并发模型 | 异步关键字 | 运行时 | 特点 |
|------|----------|------------|--------|------|
| TypeScript | 事件循环 | `async/await` | V8/Node | 单线程非阻塞 |
| Python | 协程 + 多进程 | `async/await` | asyncio | GIL 限制多线程 |
| Go | CSP (Goroutines) | 无 (原生) | Go Runtime | M:N 调度 |
| Rust | 零成本异步 | `async/await` | tokio/async-std | 无运行时开销 |

### TypeScript 异步编程

```typescript
// ==================== 基础 async/await ====================

async function fetchUser(id: number): Promise<User> {
  const response = await fetch(`/api/users/${id}`);
  return response.json();
}

// 调用
const user = await fetchUser(1);

// ==================== Promise 基础 ====================

// 创建 Promise
const promise = new Promise<string>((resolve, reject) => {
  setTimeout(() => resolve("done"), 1000);
});

// Promise 链式调用
fetch('/api/data')
  .then(res => res.json())
  .then(data => console.log(data))
  .catch(err => console.error(err))
  .finally(() => console.log('完成'));

// ==================== 并行执行 ====================

// Promise.all - 全部成功才成功
const [users, posts] = await Promise.all([
  fetchUsers(),
  fetchPosts()
]);

// Promise.allSettled - 等待全部完成（不管成功失败）
const results = await Promise.allSettled([
  fetchUser(1),
  fetchUser(999)  // 可能失败
]);
results.forEach(r => {
  if (r.status === 'fulfilled') console.log(r.value);
  else console.log(r.reason);
});

// Promise.race - 第一个完成的
const fastest = await Promise.race([
  fetchFromServer1(),
  fetchFromServer2()
]);

// Promise.any - 第一个成功的 (ES2021)
const firstSuccess = await Promise.any([
  fetchFromServer1(),
  fetchFromServer2()
]);

// ==================== 并发控制 ====================

// 限制并发数
async function parallelLimit<T>(
  tasks: (() => Promise<T>)[],
  limit: number
): Promise<T[]> {
  const results: T[] = [];
  const executing: Promise<void>[] = [];

  for (const task of tasks) {
    const p = task().then(r => { results.push(r); });
    executing.push(p);

    if (executing.length >= limit) {
      await Promise.race(executing);
      executing.splice(executing.findIndex(e => e === p), 1);
    }
  }
  await Promise.all(executing);
  return results;
}

// ==================== 超时控制 ====================

function withTimeout<T>(promise: Promise<T>, ms: number): Promise<T> {
  const timeout = new Promise<never>((_, reject) =>
    setTimeout(() => reject(new Error('Timeout')), ms)
  );
  return Promise.race([promise, timeout]);
}

const data = await withTimeout(fetchData(), 5000);

// ==================== Web Worker (真正的多线程) ====================

// main.ts
const worker = new Worker('worker.js');
worker.postMessage({ type: 'compute', data: bigData });
worker.onmessage = (e) => console.log(e.data);

// worker.js
self.onmessage = (e) => {
  const result = heavyComputation(e.data);
  self.postMessage(result);
};
```

### Python 异步编程

```python
import asyncio
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# ==================== 基础 async/await ====================

async def fetch_user(user_id: int) -> dict:
    await asyncio.sleep(1)  # 模拟 I/O
    return {"id": user_id, "name": "John"}

# 运行协程
user = asyncio.run(fetch_user(1))

# ==================== 并行执行 ====================

async def main():
    # asyncio.gather - 并行执行多个协程
    users = await asyncio.gather(
        fetch_user(1),
        fetch_user(2),
        fetch_user(3)
    )
    
    # 允许部分失败
    results = await asyncio.gather(
        fetch_user(1),
        fetch_user(999),
        return_exceptions=True  # 失败不抛异常，返回 Exception 对象
    )
    
    # asyncio.wait - 更灵活的控制
    tasks = [asyncio.create_task(fetch_user(i)) for i in range(10)]
    done, pending = await asyncio.wait(
        tasks,
        return_when=asyncio.FIRST_COMPLETED  # 或 ALL_COMPLETED
    )

asyncio.run(main())

# ==================== 超时控制 ====================

async def fetch_with_timeout():
    try:
        result = await asyncio.wait_for(
            fetch_user(1),
            timeout=5.0
        )
    except asyncio.TimeoutError:
        print("超时了")

# ==================== 异步迭代器 ====================

async def fetch_pages():
    for page in range(1, 10):
        data = await fetch_page(page)
        yield data

async def process_pages():
    async for page in fetch_pages():
        print(page)

# ==================== 异步上下文管理器 ====================

class AsyncConnection:
    async def __aenter__(self):
        await self.connect()
        return self
    
    async def __aexit__(self, *args):
        await self.disconnect()

async def use_connection():
    async with AsyncConnection() as conn:
        await conn.execute("SELECT 1")

# ==================== 信号量控制并发 ====================

async def fetch_with_limit(urls: list[str], limit: int = 10):
    semaphore = asyncio.Semaphore(limit)
    
    async def fetch_one(url):
        async with semaphore:
            return await fetch(url)
    
    return await asyncio.gather(*[fetch_one(url) for url in urls])

# ==================== 多线程/多进程 (CPU 密集型) ====================

# 线程池 (I/O 密集型，绕过 GIL 限制)
def blocking_io(path):
    with open(path) as f:
        return f.read()

async def read_files(paths):
    loop = asyncio.get_event_loop()
    with ThreadPoolExecutor() as pool:
        tasks = [loop.run_in_executor(pool, blocking_io, p) for p in paths]
        return await asyncio.gather(*tasks)

# 进程池 (CPU 密集型)
def cpu_bound(n):
    return sum(i * i for i in range(n))

async def parallel_compute():
    loop = asyncio.get_event_loop()
    with ProcessPoolExecutor() as pool:
        results = await asyncio.gather(
            loop.run_in_executor(pool, cpu_bound, 10**7),
            loop.run_in_executor(pool, cpu_bound, 10**7)
        )

# ==================== 队列 ====================

async def producer(queue: asyncio.Queue):
    for i in range(10):
        await queue.put(i)
    await queue.put(None)  # 结束信号

async def consumer(queue: asyncio.Queue):
    while True:
        item = await queue.get()
        if item is None:
            break
        print(f"处理: {item}")

async def main():
    queue = asyncio.Queue(maxsize=5)
    await asyncio.gather(producer(queue), consumer(queue))
```

### Go 并发编程

```go
package main

import (
    "context"
    "fmt"
    "sync"
    "time"
)

// ==================== Goroutine 基础 ====================

func main() {
    // 启动 goroutine
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    // 等待（生产环境用 sync 或 channel）
    time.Sleep(time.Millisecond)
}

// ==================== Channel 通信 ====================

func channelBasics() {
    // 无缓冲 channel（同步）
    ch := make(chan int)
    
    go func() {
        ch <- 42  // 发送（阻塞直到有人接收）
    }()
    
    value := <-ch  // 接收（阻塞直到有数据）
    
    // 有缓冲 channel（异步）
    buffered := make(chan int, 10)
    buffered <- 1  // 不阻塞（直到缓冲满）
    
    // 关闭 channel
    close(ch)
    
    // 检查是否关闭
    v, ok := <-ch  // ok=false 表示已关闭
    
    // 遍历 channel
    for v := range ch {
        fmt.Println(v)
    }
}

// ==================== Select 多路复用 ====================

func selectExample() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() { ch1 <- "one" }()
    go func() { ch2 <- "two" }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println(msg1)
        case msg2 := <-ch2:
            fmt.Println(msg2)
        case <-time.After(time.Second):  // 超时
            fmt.Println("timeout")
        default:  // 非阻塞
            fmt.Println("no message")
        }
    }
}

// ==================== WaitGroup 等待完成 ====================

func waitGroupExample() {
    var wg sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func(id int) {
            defer wg.Done()
            fmt.Printf("Worker %d\n", id)
        }(i)
    }
    
    wg.Wait()  // 等待所有 goroutine 完成
}

// ==================== 工作池模式 ====================

func workerPool() {
    jobs := make(chan int, 100)
    results := make(chan int, 100)
    
    // 启动 3 个 worker
    for w := 0; w < 3; w++ {
        go func(id int) {
            for job := range jobs {
                results <- job * 2  // 处理任务
            }
        }(w)
    }
    
    // 发送任务
    for j := 0; j < 10; j++ {
        jobs <- j
    }
    close(jobs)
    
    // 收集结果
    for r := 0; r < 10; r++ {
        fmt.Println(<-results)
    }
}

// ==================== Context 取消与超时 ====================

func contextExample() {
    // 带超时的 context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    // 带取消的 context
    ctx, cancel = context.WithCancel(context.Background())
    
    go func() {
        select {
        case <-ctx.Done():
            fmt.Println("cancelled:", ctx.Err())
            return
        case <-time.After(10 * time.Second):
            fmt.Println("completed")
        }
    }()
    
    time.Sleep(time.Second)
    cancel()  // 取消
}

// ==================== Mutex 互斥锁 ====================

type SafeCounter struct {
    mu    sync.Mutex
    value int
}

func (c *SafeCounter) Inc() {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.value++
}

func (c *SafeCounter) Value() int {
    c.mu.Lock()
    defer c.mu.Unlock()
    return c.value
}

// RWMutex 读写锁
type SafeCache struct {
    mu    sync.RWMutex
    items map[string]string
}

func (c *SafeCache) Get(key string) string {
    c.mu.RLock()  // 读锁（多个读可并行）
    defer c.mu.RUnlock()
    return c.items[key]
}

func (c *SafeCache) Set(key, value string) {
    c.mu.Lock()  // 写锁（独占）
    defer c.mu.Unlock()
    c.items[key] = value
}

// ==================== 原子操作 ====================

import "sync/atomic"

var counter int64

func atomicExample() {
    atomic.AddInt64(&counter, 1)
    atomic.LoadInt64(&counter)
    atomic.StoreInt64(&counter, 100)
    atomic.CompareAndSwapInt64(&counter, 100, 200)
}

// ==================== Once 只执行一次 ====================

var once sync.Once
var instance *Database

func GetInstance() *Database {
    once.Do(func() {
        instance = &Database{}
    })
    return instance
}

// ==================== errgroup 错误处理 ====================

import "golang.org/x/sync/errgroup"

func errgroupExample() error {
    g, ctx := errgroup.WithContext(context.Background())
    
    g.Go(func() error {
        return fetchData1(ctx)
    })
    
    g.Go(func() error {
        return fetchData2(ctx)
    })
    
    // 等待所有完成，返回第一个错误
    return g.Wait()
}
```

### Rust 异步编程

```rust
use tokio;
use std::sync::{Arc, Mutex, RwLock};
use std::sync::atomic::{AtomicUsize, Ordering};

// ==================== 基础 async/await ====================

async fn fetch_user(id: u32) -> Result<User, Error> {
    // 异步操作
    let response = reqwest::get(&format!("/api/users/{}", id)).await?;
    response.json().await
}

// 运行异步函数
#[tokio::main]
async fn main() {
    let user = fetch_user(1).await.unwrap();
}

// ==================== 并行执行 ====================

use futures::future::{join, join_all, try_join, try_join_all, select};

async fn parallel_tasks() {
    // join! - 等待全部完成
    let (user, posts) = tokio::join!(
        fetch_user(1),
        fetch_posts(1)
    );
    
    // try_join! - 任一失败则返回错误
    let result = tokio::try_join!(
        fetch_user(1),
        fetch_posts(1)
    );
    
    // join_all - 动态数量的 futures
    let users = join_all(
        (1..10).map(|id| fetch_user(id))
    ).await;
    
    // select! - 第一个完成的
    tokio::select! {
        user = fetch_user(1) => println!("user: {:?}", user),
        posts = fetch_posts(1) => println!("posts: {:?}", posts),
    }
}

// ==================== 超时控制 ====================

use tokio::time::{timeout, Duration};

async fn with_timeout() {
    match timeout(Duration::from_secs(5), fetch_user(1)).await {
        Ok(result) => println!("成功: {:?}", result),
        Err(_) => println!("超时"),
    }
}

// ==================== 异步 Stream ====================

use futures::stream::{self, StreamExt};
use tokio_stream::wrappers::IntervalStream;

async fn stream_example() {
    // 从迭代器创建 stream
    let mut stream = stream::iter(vec![1, 2, 3]);
    while let Some(value) = stream.next().await {
        println!("{}", value);
    }
    
    // 定时器 stream
    let interval = tokio::time::interval(Duration::from_secs(1));
    let mut stream = IntervalStream::new(interval).take(5);
    while let Some(_) = stream.next().await {
        println!("tick");
    }
    
    // 并发处理 stream
    let results: Vec<_> = stream::iter(urls)
        .map(|url| fetch(url))
        .buffer_unordered(10)  // 最多 10 个并发
        .collect()
        .await;
}

// ==================== Spawn 任务 ====================

async fn spawn_tasks() {
    // spawn 独立任务
    let handle = tokio::spawn(async {
        fetch_user(1).await
    });
    
    // 等待任务完成
    let result = handle.await.unwrap();
    
    // spawn_blocking (CPU 密集型)
    let result = tokio::task::spawn_blocking(|| {
        heavy_computation()
    }).await.unwrap();
}

// ==================== Channel ====================

use tokio::sync::{mpsc, oneshot, broadcast, watch};

async fn channel_examples() {
    // mpsc - 多生产者单消费者
    let (tx, mut rx) = mpsc::channel::<i32>(100);
    
    tokio::spawn(async move {
        tx.send(42).await.unwrap();
    });
    
    while let Some(value) = rx.recv().await {
        println!("{}", value);
    }
    
    // oneshot - 单次发送
    let (tx, rx) = oneshot::channel::<String>();
    tx.send("hello".to_string()).unwrap();
    let value = rx.await.unwrap();
    
    // broadcast - 多消费者
    let (tx, _) = broadcast::channel::<i32>(16);
    let mut rx1 = tx.subscribe();
    let mut rx2 = tx.subscribe();
    
    tx.send(1).unwrap();
    
    // watch - 最新值
    let (tx, mut rx) = watch::channel("initial");
    tx.send("updated").unwrap();
    println!("{}", *rx.borrow());
}

// ==================== Mutex & RwLock ====================

use tokio::sync::{Mutex as AsyncMutex, RwLock as AsyncRwLock};

async fn async_mutex_example() {
    let data = Arc::new(AsyncMutex::new(0));
    
    let data_clone = Arc::clone(&data);
    tokio::spawn(async move {
        let mut lock = data_clone.lock().await;
        *lock += 1;
    });
}

// 同步 Mutex (短时间持有)
fn sync_mutex_example() {
    let counter = Arc::new(Mutex::new(0));
    
    let handles: Vec<_> = (0..10).map(|_| {
        let counter = Arc::clone(&counter);
        std::thread::spawn(move || {
            let mut num = counter.lock().unwrap();
            *num += 1;
        })
    }).collect();
    
    for handle in handles {
        handle.join().unwrap();
    }
}

// ==================== Semaphore 信号量 ====================

use tokio::sync::Semaphore;

async fn semaphore_example() {
    let semaphore = Arc::new(Semaphore::new(10));  // 最多 10 并发
    
    let permit = semaphore.acquire().await.unwrap();
    // 执行受限操作
    drop(permit);  // 释放许可
}

// ==================== 原子操作 ====================

fn atomic_example() {
    let counter = AtomicUsize::new(0);
    
    counter.fetch_add(1, Ordering::SeqCst);
    counter.load(Ordering::SeqCst);
    counter.store(100, Ordering::SeqCst);
    counter.compare_exchange(100, 200, Ordering::SeqCst, Ordering::SeqCst);
}

// ==================== 线程 ====================

use std::thread;

fn thread_example() {
    let handle = thread::spawn(|| {
        println!("Hello from thread");
        42
    });
    
    let result = handle.join().unwrap();
    
    // 作用域线程 (可借用数据)
    let data = vec![1, 2, 3];
    thread::scope(|s| {
        s.spawn(|| {
            println!("{:?}", data);  // 可以借用 data
        });
    });
}
```

### 并发模式对比总结

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 并发原语        │ Promise        │ coroutine      │ goroutine      │ Future         │
│ 通信机制        │ 无内置         │ Queue          │ Channel        │ Channel        │
│ 真正并行        │ Worker         │ multiprocessing│ 原生           │ 原生           │
│ 共享状态        │ 无 (Worker隔离)│ Lock           │ Mutex/Channel  │ Arc<Mutex>     │
│ 取消机制        │ AbortController│ CancelScope    │ Context        │ select!/drop   │
│ 错误传播        │ Promise.all    │ gather         │ errgroup       │ try_join!      │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

---

## ❌ 错误处理

### 错误处理模型概览

| 语言 | 主要模式 | 错误类型 | 特点 |
|------|----------|----------|------|
| TypeScript | 异常 + 可选链 | `Error` | try/catch，灵活但不强制 |
| Python | 异常 | `Exception` | try/except，EAFP 风格 |
| Go | 返回值 | `error` | 显式检查，无异常 |
| Rust | Result/Option | `Result<T, E>` | 编译器强制处理 |

### TypeScript 错误处理

```typescript
// ==================== 基础 try/catch ====================

try {
  const data = JSON.parse(invalidJson);
} catch (error) {
  if (error instanceof SyntaxError) {
    console.error("JSON 解析错误:", error.message);
  } else {
    throw error;  // 重新抛出未知错误
  }
} finally {
  cleanup();
}

// ==================== 自定义错误 ====================

class ValidationError extends Error {
  constructor(
    message: string,
    public field: string,
    public code: string
  ) {
    super(message);
    this.name = 'ValidationError';
    Error.captureStackTrace(this, ValidationError);
  }
}

throw new ValidationError("无效的邮箱", "email", "INVALID_EMAIL");

// ==================== 类型安全的错误处理 ====================

// Result 类型模式
type Result<T, E = Error> = 
  | { success: true; data: T }
  | { success: false; error: E };

function parseJSON<T>(json: string): Result<T> {
  try {
    return { success: true, data: JSON.parse(json) };
  } catch (e) {
    return { success: false, error: e as Error };
  }
}

const result = parseJSON<User>('{"name":"John"}');
if (result.success) {
  console.log(result.data.name);
} else {
  console.error(result.error.message);
}

// ==================== 异步错误处理 ====================

// async/await
async function fetchData() {
  try {
    const response = await fetch('/api/data');
    if (!response.ok) {
      throw new Error(`HTTP ${response.status}`);
    }
    return await response.json();
  } catch (error) {
    console.error("获取数据失败:", error);
    throw error;  // 或返回默认值
  }
}

// Promise 链
fetch('/api/data')
  .then(res => res.json())
  .catch(err => {
    console.error(err);
    return defaultData;
  });

// Promise.allSettled 处理部分失败
const results = await Promise.allSettled([
  fetchUser(1),
  fetchUser(999)
]);

results.forEach((result, i) => {
  if (result.status === 'fulfilled') {
    console.log(`用户 ${i}:`, result.value);
  } else {
    console.error(`用户 ${i} 失败:`, result.reason);
  }
});

// ==================== 可选链与空值合并 ====================

// 可选链 (?.)
const name = user?.profile?.name;
const first = arr?.[0];
const result = obj?.method?.();

// 空值合并 (??)
const value = input ?? defaultValue;  // 仅 null/undefined 时使用默认值
const name = user.name ?? "Anonymous";

// 组合使用
const city = user?.address?.city ?? "Unknown";

// ==================== 断言函数 ====================

function assertNonNull<T>(
  value: T | null | undefined,
  message: string
): asserts value is T {
  if (value === null || value === undefined) {
    throw new Error(message);
  }
}

assertNonNull(user, "用户不存在");
console.log(user.name);  // 类型收窄为非空

// ==================== 错误边界 (React) ====================

class ErrorBoundary extends React.Component {
  state = { hasError: false };

  static getDerivedStateFromError(error: Error) {
    return { hasError: true };
  }

  componentDidCatch(error: Error, info: React.ErrorInfo) {
    logError(error, info);
  }

  render() {
    if (this.state.hasError) {
      return <h1>出错了</h1>;
    }
    return this.props.children;
  }
}
```

### Python 错误处理

```python
# ==================== 基础 try/except ====================

try:
    result = risky_operation()
except ValueError as e:
    print(f"值错误: {e}")
except (TypeError, KeyError) as e:  # 多个异常
    print(f"类型或键错误: {e}")
except Exception as e:  # 捕获所有异常
    print(f"未知错误: {e}")
    raise  # 重新抛出
else:
    print("没有异常时执行")
finally:
    cleanup()

# ==================== 自定义异常 ====================

class ValidationError(Exception):
    def __init__(self, message: str, field: str, code: str):
        super().__init__(message)
        self.field = field
        self.code = code

class NotFoundError(Exception):
    pass

# 异常链
try:
    parse_config()
except FileNotFoundError as e:
    raise ConfigError("配置加载失败") from e

# ==================== 上下文管理器 ====================

# 自动资源清理
with open('file.txt') as f:
    content = f.read()  # 异常时也会自动关闭

# 自定义上下文管理器
from contextlib import contextmanager

@contextmanager
def managed_resource():
    resource = acquire_resource()
    try:
        yield resource
    finally:
        release_resource(resource)

with managed_resource() as r:
    r.do_something()

# suppress 忽略特定异常
from contextlib import suppress

with suppress(FileNotFoundError):
    os.remove('file.txt')  # 文件不存在也不报错

# ==================== 异常组 (Python 3.11+) ====================

# ExceptionGroup 同时抛出多个异常
errors = []
for item in items:
    try:
        process(item)
    except Exception as e:
        errors.append(e)

if errors:
    raise ExceptionGroup("批量处理失败", errors)

# 捕获异常组
try:
    process_batch()
except* ValueError as eg:  # except* 语法
    print(f"值错误: {eg.exceptions}")
except* TypeError as eg:
    print(f"类型错误: {eg.exceptions}")

# ==================== 类型安全的错误处理 ====================

from typing import TypeVar, Generic
from dataclasses import dataclass

T = TypeVar('T')
E = TypeVar('E')

@dataclass
class Ok(Generic[T]):
    value: T

@dataclass  
class Err(Generic[E]):
    error: E

Result = Ok[T] | Err[E]

def divide(a: int, b: int) -> Result[float, str]:
    if b == 0:
        return Err("除数不能为零")
    return Ok(a / b)

result = divide(10, 0)
match result:
    case Ok(value):
        print(f"结果: {value}")
    case Err(error):
        print(f"错误: {error}")

# ==================== 断言 ====================

assert condition, "条件不满足"  # 可用 -O 参数禁用

# 运行时类型检查
def process(data: list[int]) -> int:
    if not isinstance(data, list):
        raise TypeError(f"期望 list，得到 {type(data)}")
    return sum(data)

# ==================== 日志记录 ====================

import logging

logging.basicConfig(level=logging.INFO)
logger = logging.getLogger(__name__)

try:
    risky_operation()
except Exception:
    logger.exception("操作失败")  # 自动记录堆栈

# ==================== 警告 ====================

import warnings

warnings.warn("这个函数已弃用", DeprecationWarning)

# 将警告转为异常
warnings.filterwarnings('error')
```

### Go 错误处理

```go
package main

import (
    "errors"
    "fmt"
    "log"
)

// ==================== 基础错误处理 ====================

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("除数不能为零")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(result)
}

// ==================== 自定义错误 ====================

// 简单自定义错误
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

// 带错误码
type AppError struct {
    Code    int
    Message string
    Err     error  // 包装的原始错误
}

func (e *AppError) Error() string {
    if e.Err != nil {
        return fmt.Sprintf("[%d] %s: %v", e.Code, e.Message, e.Err)
    }
    return fmt.Sprintf("[%d] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
    return e.Err
}

// ==================== 错误包装 (Go 1.13+) ====================

func readConfig() error {
    data, err := os.ReadFile("config.json")
    if err != nil {
        return fmt.Errorf("读取配置失败: %w", err)  // %w 包装错误
    }
    return nil
}

// 检查错误链
if errors.Is(err, os.ErrNotExist) {
    fmt.Println("文件不存在")
}

// 提取特定错误类型
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    fmt.Println("路径:", pathErr.Path)
}

// ==================== 哨兵错误 ====================

var (
    ErrNotFound     = errors.New("not found")
    ErrUnauthorized = errors.New("unauthorized")
    ErrInternal     = errors.New("internal error")
)

func findUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrNotFound
    }
    return &User{}, nil
}

// 使用
_, err := findUser(-1)
if errors.Is(err, ErrNotFound) {
    // 处理未找到
}

// ==================== 多错误合并 (Go 1.20+) ====================

func validateUser(u *User) error {
    var errs []error
    
    if u.Name == "" {
        errs = append(errs, errors.New("名称为空"))
    }
    if u.Email == "" {
        errs = append(errs, errors.New("邮箱为空"))
    }
    
    return errors.Join(errs...)  // 合并多个错误
}

// ==================== defer 资源清理 ====================

func readFile(path string) ([]byte, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()  // 确保关闭
    
    return io.ReadAll(f)
}

// 带错误处理的 defer
func writeFile(path string, data []byte) (err error) {
    f, err := os.Create(path)
    if err != nil {
        return err
    }
    defer func() {
        closeErr := f.Close()
        if err == nil {
            err = closeErr  // 只在没有其他错误时设置
        }
    }()
    
    _, err = f.Write(data)
    return err
}

// ==================== panic 和 recover ====================

func mayPanic() {
    panic("something went wrong")
}

func safeCall() (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic recovered: %v", r)
        }
    }()
    
    mayPanic()
    return nil
}

// ==================== 错误处理模式 ====================

// 提前返回模式
func process(data []byte) error {
    if len(data) == 0 {
        return errors.New("空数据")
    }
    
    parsed, err := parse(data)
    if err != nil {
        return fmt.Errorf("解析失败: %w", err)
    }
    
    if err := validate(parsed); err != nil {
        return fmt.Errorf("验证失败: %w", err)
    }
    
    return save(parsed)
}

// 错误处理辅助函数
func must[T any](v T, err error) T {
    if err != nil {
        panic(err)
    }
    return v
}

config := must(loadConfig())  // 失败直接 panic

// ==================== 结构化日志 ====================

import "log/slog"

func handleRequest(r *Request) error {
    logger := slog.With("request_id", r.ID)
    
    user, err := getUser(r.UserID)
    if err != nil {
        logger.Error("获取用户失败",
            "user_id", r.UserID,
            "error", err,
        )
        return err
    }
    
    logger.Info("请求处理成功", "user", user.Name)
    return nil
}
```

### Rust 错误处理

```rust
use std::error::Error;
use std::fmt;
use std::fs::File;
use std::io::{self, Read};

// ==================== Result 基础 ====================

fn divide(a: f64, b: f64) -> Result<f64, String> {
    if b == 0.0 {
        Err("除数不能为零".to_string())
    } else {
        Ok(a / b)
    }
}

fn main() {
    match divide(10.0, 0.0) {
        Ok(result) => println!("结果: {}", result),
        Err(e) => println!("错误: {}", e),
    }
}

// ==================== Option 类型 ====================

fn find_user(id: u32) -> Option<User> {
    if id > 0 {
        Some(User { id, name: "John".to_string() })
    } else {
        None
    }
}

// Option 方法
let user = find_user(1);
let name = user.map(|u| u.name);           // Option<String>
let name = user.and_then(|u| u.nickname);  // 链式 Option
let name = user.unwrap_or_default();       // 默认值
let name = user.unwrap_or_else(|| get_default()); // 惰性默认值
let name = user.ok_or("用户不存在")?;      // 转为 Result

// if let / while let
if let Some(user) = find_user(1) {
    println!("{}", user.name);
}

// ==================== ? 操作符 ====================

fn read_file(path: &str) -> Result<String, io::Error> {
    let mut file = File::open(path)?;  // 错误自动返回
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;
    Ok(contents)
}

// 链式调用
fn read_config() -> Result<Config, io::Error> {
    let contents = std::fs::read_to_string("config.json")?;
    let config: Config = serde_json::from_str(&contents)?;
    Ok(config)
}

// ==================== 自定义错误 ====================

#[derive(Debug)]
enum AppError {
    NotFound(String),
    Validation { field: String, message: String },
    Database(sqlx::Error),
    Io(io::Error),
}

impl fmt::Display for AppError {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        match self {
            AppError::NotFound(msg) => write!(f, "未找到: {}", msg),
            AppError::Validation { field, message } => 
                write!(f, "{}: {}", field, message),
            AppError::Database(e) => write!(f, "数据库错误: {}", e),
            AppError::Io(e) => write!(f, "IO 错误: {}", e),
        }
    }
}

impl Error for AppError {
    fn source(&self) -> Option<&(dyn Error + 'static)> {
        match self {
            AppError::Database(e) => Some(e),
            AppError::Io(e) => Some(e),
            _ => None,
        }
    }
}

// From trait 实现自动转换
impl From<io::Error> for AppError {
    fn from(err: io::Error) -> Self {
        AppError::Io(err)
    }
}

// ==================== thiserror (推荐) ====================

use thiserror::Error;

#[derive(Error, Debug)]
enum MyError {
    #[error("未找到: {0}")]
    NotFound(String),
    
    #[error("验证错误 - {field}: {message}")]
    Validation { field: String, message: String },
    
    #[error("IO 错误")]
    Io(#[from] io::Error),
    
    #[error("JSON 错误")]
    Json(#[from] serde_json::Error),
}

// ==================== anyhow (应用程序) ====================

use anyhow::{Context, Result, bail, ensure};

fn load_config() -> Result<Config> {
    let contents = std::fs::read_to_string("config.json")
        .context("无法读取配置文件")?;
    
    let config: Config = serde_json::from_str(&contents)
        .context("配置文件格式错误")?;
    
    // bail! 宏快速返回错误
    if config.name.is_empty() {
        bail!("配置名称不能为空");
    }
    
    // ensure! 宏断言
    ensure!(config.port > 0, "端口必须大于 0");
    
    Ok(config)
}

// 使用 anyhow::Result
fn main() -> Result<()> {
    let config = load_config()?;
    println!("{:?}", config);
    Ok(())
}

// ==================== Result 方法 ====================

let result: Result<i32, &str> = Ok(42);

// 转换
result.map(|v| v * 2);           // Ok(84)
result.map_err(|e| e.to_string()); // 转换错误类型
result.and_then(|v| if v > 0 { Ok(v) } else { Err("负数") });

// 默认值
result.unwrap_or(0);
result.unwrap_or_default();
result.unwrap_or_else(|_| calculate_default());

// 危险操作 (仅在确定不会失败时使用)
result.unwrap();   // 失败时 panic
result.expect("错误消息");  // 带消息的 panic

// 检查
result.is_ok();
result.is_err();

// 转换为 Option
result.ok();   // Option<T>
result.err();  // Option<E>

// ==================== 模式匹配 ====================

fn handle_result(r: Result<i32, AppError>) {
    match r {
        Ok(v) if v > 100 => println!("大值: {}", v),
        Ok(v) => println!("值: {}", v),
        Err(AppError::NotFound(msg)) => println!("未找到: {}", msg),
        Err(e) => println!("其他错误: {}", e),
    }
}

// let else (Rust 1.65+)
fn process(input: Option<String>) {
    let Some(value) = input else {
        println!("没有输入");
        return;
    };
    println!("处理: {}", value);
}

// ==================== panic 处理 ====================

// 设置 panic 钩子
std::panic::set_hook(Box::new(|info| {
    eprintln!("Panic: {:?}", info);
}));

// 捕获 panic
let result = std::panic::catch_unwind(|| {
    panic!("出错了");
});

match result {
    Ok(_) => println!("正常"),
    Err(_) => println!("捕获到 panic"),
}

// ==================== 错误传播最佳实践 ====================

// 库代码: 使用具体错误类型
pub fn library_function() -> Result<Data, LibraryError> {
    // ...
}

// 应用代码: 使用 anyhow
fn app_function() -> anyhow::Result<()> {
    let data = library_function()
        .context("调用库函数失败")?;
    Ok(())
}
```

### 错误处理对比总结

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 主要机制        │ try/catch      │ try/except     │ 返回 error     │ Result<T,E>    │
│ 强制处理        │ ❌             │ ❌             │ ❌ (可忽略)    │ ✅ (编译器)    │
│ 错误传播        │ throw          │ raise          │ return err     │ ?              │
│ 空值处理        │ ?. / ??        │ or / None      │ nil 检查       │ Option<T>      │
│ 异常类型        │ Error 类       │ Exception 类   │ error 接口     │ Error trait    │
│ 错误链          │ cause          │ from           │ %w / Unwrap    │ source()       │
│ 资源清理        │ finally        │ finally/with   │ defer          │ Drop trait     │
│ panic 恢复      │ 无             │ 无 (异常即可)  │ recover        │ catch_unwind   │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 最佳实践建议

| 场景 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 业务错误 | 自定义 Error 类 | 自定义 Exception | 自定义 error | thiserror |
| 应用错误 | Result 模式 | 异常 | 错误包装 | anyhow |
| 验证错误 | Zod/Yup | Pydantic | 结构体验证 | validator |
| 可恢复错误 | try/catch | try/except | if err != nil | Result + ? |
| 不可恢复 | throw | raise | panic | panic! |

---

## 📦 模块化系统

### 模块系统概览

| 语言 | 模块单位 | 导入关键字 | 可见性控制 | 特点 |
|------|----------|------------|------------|------|
| TypeScript | 文件 | `import/export` | export 显式导出 | ES Modules |
| Python | 文件/目录 | `import/from` | `_` 前缀约定 | 包即目录 |
| Go | 目录 | `import` | 大小写控制 | 简单直接 |
| Rust | 文件/mod | `use` | `pub` 关键字 | 精细控制 |

### TypeScript 模块系统

```typescript
// ==================== 导出 ====================

// 命名导出 (named export)
export const PI = 3.14159;
export function add(a: number, b: number): number {
  return a + b;
}
export class Calculator {
  // ...
}
export interface Config {
  name: string;
}
export type Result<T> = { ok: true; value: T } | { ok: false; error: Error };

// 默认导出 (default export)
export default class App {
  // ...
}

// 重新导出
export { foo, bar } from './other';
export * from './utils';                    // 导出所有
export * as utils from './utils';           // 命名空间导出
export { default as Helper } from './helper';

// ==================== 导入 ====================

// 命名导入
import { add, PI } from './math';
import { add as addition } from './math';   // 重命名
import type { Config } from './config';     // 仅类型导入

// 默认导入
import App from './App';
import App, { helper } from './App';        // 混合导入

// 命名空间导入
import * as math from './math';
math.add(1, 2);

// 动态导入 (代码分割)
const module = await import('./heavy-module');
module.doSomething();

// 条件导入
if (condition) {
  const { feature } = await import('./feature');
}

// ==================== 项目结构 ====================

/*
src/
├── index.ts              # 入口文件
├── types/
│   └── index.ts          # 类型定义
├── utils/
│   ├── index.ts          # 桶文件 (barrel)
│   ├── string.ts
│   └── date.ts
├── services/
│   ├── index.ts
│   └── api.service.ts
└── components/
    ├── index.ts
    └── Button.tsx
*/

// utils/index.ts (桶文件)
export * from './string';
export * from './date';
export { default as format } from './format';

// 使用
import { formatDate, capitalize } from '@/utils';

// ==================== 路径别名 (tsconfig.json) ====================

/*
{
  "compilerOptions": {
    "baseUrl": ".",
    "paths": {
      "@/*": ["src/*"],
      "@components/*": ["src/components/*"],
      "@utils/*": ["src/utils/*"]
    }
  }
}
*/

import { Button } from '@components/Button';

// ==================== 模块声明 ====================

// 声明模块类型 (*.d.ts)
declare module '*.css' {
  const content: { [className: string]: string };
  export default content;
}

declare module '*.svg' {
  const content: React.FC<React.SVGProps<SVGSVGElement>>;
  export default content;
}

// 扩展已有模块
declare module 'express' {
  interface Request {
    user?: User;
  }
}
```

### Python 模块系统

```python
# ==================== 导出 ====================

# Python 没有显式导出，所有顶层定义默认可导入
# 使用 __all__ 控制 from module import * 的行为

# math_utils.py
__all__ = ['add', 'subtract', 'PI']  # 限制 * 导入的内容

PI = 3.14159

def add(a, b):
    return a + b

def subtract(a, b):
    return a - b

def _private_helper():  # 下划线前缀表示私有（约定）
    pass

# ==================== 导入 ====================

# 导入模块
import math
math.sqrt(4)

# 导入特定对象
from math import sqrt, pi
from math import sqrt as square_root  # 重命名

# 导入所有 (不推荐)
from math import *

# 相对导入 (包内)
from . import sibling_module
from .sibling import function
from .. import parent_module
from ..parent import function

# 条件导入
try:
    import ujson as json  # 优先使用更快的实现
except ImportError:
    import json

# 延迟导入 (避免循环依赖)
def my_function():
    from heavy_module import heavy_function
    return heavy_function()

# 类型检查时导入 (避免运行时导入)
from typing import TYPE_CHECKING
if TYPE_CHECKING:
    from expensive_module import ExpensiveClass

# ==================== 包结构 ====================

"""
mypackage/
├── __init__.py           # 包初始化（必需）
├── __main__.py           # python -m mypackage 入口
├── core/
│   ├── __init__.py
│   ├── engine.py
│   └── utils.py
├── api/
│   ├── __init__.py
│   └── endpoints.py
└── models/
    ├── __init__.py
    └── user.py
"""

# mypackage/__init__.py
from .core.engine import Engine
from .core.utils import helper
from .models.user import User

__all__ = ['Engine', 'helper', 'User']
__version__ = '1.0.0'

# 使用
from mypackage import Engine, User

# ==================== 命名空间包 (无 __init__.py) ====================

"""
Python 3.3+ 支持命名空间包，允许包分布在多个目录

site-packages/
├── mynamespace/
│   └── subpkg1/
│       └── __init__.py
└── mynamespace/          # 另一个位置
    └── subpkg2/
        └── __init__.py

两个子包合并为 mynamespace
"""

# ==================== 模块搜索路径 ====================

import sys

# 查看搜索路径
print(sys.path)

# 添加搜索路径
sys.path.insert(0, '/path/to/modules')

# 使用 PYTHONPATH 环境变量
# export PYTHONPATH=/path/to/modules:$PYTHONPATH

# ==================== 模块属性 ====================

# 模块名
__name__      # 'mymodule' 或 '__main__'
__file__      # 模块文件路径
__package__   # 包名
__doc__       # 模块文档字符串

# 主模块判断
if __name__ == '__main__':
    main()

# ==================== importlib 动态导入 ====================

import importlib

# 动态导入模块
module = importlib.import_module('mypackage.submodule')

# 重新加载模块 (开发时有用)
importlib.reload(module)

# 导入插件
def load_plugins(plugin_names):
    plugins = []
    for name in plugin_names:
        module = importlib.import_module(f'plugins.{name}')
        plugins.append(module.Plugin())
    return plugins
```

### Go 模块系统

```go
// ==================== 包声明与导入 ====================

// 每个 Go 文件必须声明包名
package main

import (
    // 标准库
    "fmt"
    "net/http"
    
    // 第三方包
    "github.com/gin-gonic/gin"
    
    // 本地包
    "myproject/internal/auth"
    "myproject/pkg/utils"
    
    // 别名导入
    log "github.com/sirupsen/logrus"
    
    // 匿名导入 (仅执行 init)
    _ "github.com/lib/pq"
    
    // 点导入 (不推荐)
    . "math"
)

// ==================== 可见性规则 ====================

// 大写开头 = 公开 (exported)
func PublicFunction() {}
type PublicStruct struct {
    PublicField  string  // 公开字段
    privateField string  // 私有字段
}
const PublicConst = 1
var PublicVar = "hello"

// 小写开头 = 私有 (unexported)
func privateFunction() {}
type privateStruct struct{}
const privateConst = 2
var privateVar = "secret"

// ==================== 项目结构 ====================

/*
myproject/
├── go.mod                # 模块定义
├── go.sum                # 依赖校验
├── main.go               # 入口点
├── cmd/                  # 可执行程序
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/             # 私有包 (不可被外部导入)
│   ├── auth/
│   │   └── auth.go
│   └── database/
│       └── db.go
├── pkg/                  # 公开包 (可被外部导入)
│   └── utils/
│       └── utils.go
├── api/                  # API 定义
│   └── handlers.go
└── configs/              # 配置文件
    └── config.yaml
*/

// ==================== init 函数 ====================

package database

import "database/sql"

var db *sql.DB

// init 在包导入时自动执行
// 可以有多个 init，按声明顺序执行
func init() {
    var err error
    db, err = sql.Open("postgres", connStr)
    if err != nil {
        panic(err)
    }
}

// 执行顺序: 导入包 → 包级变量初始化 → init() → main()

// ==================== internal 包 ====================

/*
internal 目录下的包只能被同一模块导入

myproject/
├── internal/
│   └── secret/         # 只有 myproject 内部可用
│       └── secret.go
└── pkg/
    └── public/         # 任何人都可导入
        └── public.go

外部项目无法导入:
import "myproject/internal/secret"  // 编译错误
*/

// ==================== 同一目录多文件 ====================

// 同一目录下的所有 .go 文件属于同一个包
// utils/string.go
package utils

func Capitalize(s string) string { ... }

// utils/number.go  
package utils

func Abs(n int) int { ... }

// 使用时整个包一起导入
import "myproject/pkg/utils"
utils.Capitalize("hello")
utils.Abs(-5)

// ==================== 构建标签 ====================

// 条件编译

//go:build linux
// +build linux

package main
// 仅在 Linux 上编译

//go:build !windows
// +build !windows

package main
// 除 Windows 外的所有平台

// 文件名约定
// file_linux.go    - 仅 Linux
// file_windows.go  - 仅 Windows
// file_test.go     - 仅测试

// ==================== 嵌入文件 ====================

import "embed"

//go:embed templates/*
var templates embed.FS

//go:embed config.json
var configData []byte

//go:embed version.txt
var version string
```

### Rust 模块系统

```rust
// ==================== 模块声明 ====================

// 方式 1: 内联模块
mod math {
    pub const PI: f64 = 3.14159;
    
    pub fn add(a: i32, b: i32) -> i32 {
        a + b
    }
    
    fn private_helper() {
        // 私有函数
    }
    
    // 嵌套模块
    pub mod advanced {
        pub fn complex_calc() {}
    }
}

// 方式 2: 单独文件
// 在 main.rs 或 lib.rs 中声明
mod utils;      // 加载 utils.rs 或 utils/mod.rs
mod database;   // 加载 database.rs 或 database/mod.rs

// ==================== 可见性控制 ====================

mod outer {
    pub fn public_fn() {}           // 完全公开
    fn private_fn() {}              // 私有 (默认)
    
    pub(crate) fn crate_fn() {}     // crate 内可见
    pub(super) fn parent_fn() {}    // 父模块可见
    pub(in crate::path) fn path_fn() {} // 指定路径可见
    
    pub struct Person {
        pub name: String,           // 公开字段
        age: u32,                   // 私有字段
        pub(crate) id: u64,         // crate 内可见
    }
    
    impl Person {
        // 公开构造函数 (因为有私有字段)
        pub fn new(name: String, age: u32) -> Self {
            Self { name, age, id: 0 }
        }
    }
}

// ==================== use 导入 ====================

// 绝对路径
use crate::math::add;
use crate::math::PI;

// 相对路径
use self::utils::helper;
use super::parent_module;

// 多个导入
use std::collections::{HashMap, HashSet, BTreeMap};
use std::io::{self, Read, Write};  // self = std::io

// 重命名
use std::io::Result as IoResult;

// 通配符 (不推荐)
use std::collections::*;

// 重新导出
pub use self::internal::PublicApi;
pub use crate::utils::*;

// ==================== 项目结构 ====================

/*
my_project/
├── Cargo.toml            # 项目配置
├── Cargo.lock            # 依赖锁定
├── src/
│   ├── main.rs           # 二进制入口 (bin crate)
│   ├── lib.rs            # 库入口 (lib crate)
│   ├── utils.rs          # 模块文件
│   ├── utils/            # 或模块目录
│   │   ├── mod.rs        # 模块入口
│   │   ├── string.rs
│   │   └── date.rs
│   └── models/
│       ├── mod.rs
│       └── user.rs
├── tests/                # 集成测试
│   └── integration_test.rs
├── benches/              # 基准测试
│   └── benchmark.rs
└── examples/             # 示例
    └── example.rs
*/

// ==================== mod.rs vs 文件名.rs ====================

// Rust 2018+ 推荐: 使用文件名.rs 代替 mod.rs

// 旧方式 (Rust 2015)
// utils/mod.rs

// 新方式 (Rust 2018+)
// utils.rs          - 模块声明
// utils/string.rs   - 子模块

// src/utils.rs
pub mod string;  // 加载 utils/string.rs
pub mod date;    // 加载 utils/date.rs

// ==================== lib.rs 组织 ====================

// src/lib.rs
pub mod api;
pub mod models;
pub mod utils;

// 重新导出常用类型
pub use api::Client;
pub use models::{User, Post};

// prelude 模式
pub mod prelude {
    pub use crate::api::Client;
    pub use crate::models::*;
    pub use crate::utils::*;
}

// 用户使用
use my_crate::prelude::*;

// ==================== 条件编译 ====================

#[cfg(target_os = "linux")]
mod linux_specific {
    pub fn platform_fn() {}
}

#[cfg(target_os = "windows")]
mod windows_specific {
    pub fn platform_fn() {}
}

#[cfg(feature = "advanced")]
pub mod advanced_features {
    // 仅在启用 feature 时编译
}

#[cfg(test)]
mod tests {
    use super::*;
    
    #[test]
    fn test_something() {}
}

// ==================== 工作空间 (Workspace) ====================

/*
workspace/
├── Cargo.toml           # 工作空间配置
├── crates/
│   ├── core/
│   │   ├── Cargo.toml
│   │   └── src/lib.rs
│   ├── cli/
│   │   ├── Cargo.toml
│   │   └── src/main.rs
│   └── web/
│       ├── Cargo.toml
│       └── src/main.rs
└── shared/
    ├── Cargo.toml
    └── src/lib.rs
*/

// workspace/Cargo.toml
/*
[workspace]
members = [
    "crates/core",
    "crates/cli",
    "crates/web",
    "shared",
]

[workspace.dependencies]
serde = "1.0"
tokio = "1.0"
*/

// crates/cli/Cargo.toml
/*
[dependencies]
core = { path = "../core" }
shared = { path = "../../shared" }
serde = { workspace = true }
*/
```

### 模块系统对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 模块单位        │ 文件           │ 文件/目录      │ 目录           │ 文件/mod      │
│ 导入语法        │ import/export  │ import/from    │ import         │ use            │
│ 默认可见性      │ 私有           │ 公开           │ 私有(小写)     │ 私有           │
│ 公开标记        │ export         │ 无 (__all__)   │ 大写首字母     │ pub            │
│ 循环依赖        │ 允许           │ 允许(需注意)   │ 不允许         │ 不允许         │
│ 动态导入        │ import()       │ importlib      │ 不支持         │ 不支持         │
│ 类型导入        │ import type    │ TYPE_CHECKING  │ 不适用         │ 自动           │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

---

## 📦 包管理

### 包管理器概览

| 语言 | 包管理器 | 配置文件 | 锁文件 | 仓库 |
|------|----------|----------|--------|------|
| TypeScript | npm/yarn/pnpm | package.json | package-lock.json | npmjs.com |
| Python | pip/poetry/uv | requirements.txt/pyproject.toml | poetry.lock | pypi.org |
| Go | go mod | go.mod | go.sum | proxy.golang.org |
| Rust | cargo | Cargo.toml | Cargo.lock | crates.io |

### TypeScript/JavaScript 包管理

```bash
# ==================== npm ====================

# 初始化项目
npm init -y

# 安装依赖
npm install express                    # 生产依赖
npm install -D typescript              # 开发依赖
npm install -g ts-node                 # 全局安装

# 版本控制
npm install lodash@4.17.21             # 精确版本
npm install lodash@^4.17.0             # 兼容版本
npm install lodash@~4.17.0             # 补丁版本

# 其他命令
npm update                             # 更新依赖
npm outdated                           # 检查过时包
npm audit                              # 安全审计
npm audit fix                          # 自动修复
npm run build                          # 运行脚本
npm exec ts-node                       # 执行包命令

# ==================== yarn ====================

yarn init -y
yarn add express
yarn add -D typescript
yarn global add ts-node
yarn upgrade
yarn upgrade-interactive

# ==================== pnpm (推荐) ====================

pnpm init
pnpm add express
pnpm add -D typescript
pnpm add -g ts-node
pnpm update
pnpm store prune                       # 清理存储
```

```json
// package.json
{
  "name": "my-project",
  "version": "1.0.0",
  "type": "module",
  "main": "dist/index.js",
  "types": "dist/index.d.ts",
  "scripts": {
    "dev": "ts-node src/index.ts",
    "build": "tsc",
    "test": "jest",
    "lint": "eslint src/",
    "prepare": "husky install"
  },
  "dependencies": {
    "express": "^4.18.0"
  },
  "devDependencies": {
    "typescript": "^5.0.0",
    "@types/express": "^4.17.0",
    "@types/node": "^20.0.0"
  },
  "engines": {
    "node": ">=18.0.0"
  },
  "peerDependencies": {
    "react": "^18.0.0"
  },
  "optionalDependencies": {
    "fsevents": "^2.3.0"
  }
}
```

```typescript
// ==================== Monorepo (pnpm workspace) ====================

/*
my-monorepo/
├── pnpm-workspace.yaml
├── package.json
├── packages/
│   ├── core/
│   │   └── package.json
│   ├── cli/
│   │   └── package.json
│   └── web/
│       └── package.json
└── apps/
    └── api/
        └── package.json
*/

// pnpm-workspace.yaml
/*
packages:
  - 'packages/*'
  - 'apps/*'
*/

// packages/cli/package.json
/*
{
  "dependencies": {
    "@my-org/core": "workspace:*"
  }
}
*/

// 工作空间命令
// pnpm -F @my-org/cli add lodash      # 给特定包添加依赖
// pnpm -r build                        # 所有包执行构建
```

### Python 包管理

```bash
# ==================== pip ====================

# 安装包
pip install requests
pip install requests==2.28.0           # 精确版本
pip install "requests>=2.28.0,<3.0.0"  # 版本范围
pip install -e .                       # 可编辑安装
pip install -r requirements.txt        # 从文件安装

# 导出依赖
pip freeze > requirements.txt

# 升级
pip install --upgrade requests
pip install --upgrade pip

# 卸载
pip uninstall requests

# ==================== 虚拟环境 ====================

# venv (内置)
python -m venv .venv
source .venv/bin/activate              # Linux/macOS
.venv\Scripts\activate                 # Windows
deactivate

# ==================== poetry (推荐) ====================

# 安装 poetry
curl -sSL https://install.python-poetry.org | python3 -

# 创建项目
poetry new my-project
poetry init                            # 交互式初始化

# 依赖管理
poetry add requests
poetry add --group dev pytest          # 开发依赖
poetry remove requests
poetry update

# 运行
poetry run python main.py
poetry shell                           # 进入虚拟环境

# 构建与发布
poetry build
poetry publish

# ==================== uv (新一代，超快) ====================

# 安装
curl -LsSf https://astral.sh/uv/install.sh | sh

# 项目管理
uv init my-project
uv add requests
uv add --dev pytest
uv remove requests
uv sync                                # 同步依赖
uv run python main.py
uv pip compile requirements.in -o requirements.txt
```

```toml
# pyproject.toml (现代标准)
[project]
name = "my-project"
version = "1.0.0"
description = "My awesome project"
readme = "README.md"
requires-python = ">=3.10"
license = {text = "MIT"}
authors = [
    {name = "Author", email = "author@example.com"}
]
dependencies = [
    "requests>=2.28.0",
    "pydantic>=2.0.0",
]

[project.optional-dependencies]
dev = [
    "pytest>=7.0.0",
    "mypy>=1.0.0",
    "ruff>=0.1.0",
]

[project.scripts]
my-cli = "my_project.cli:main"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"

# Poetry 特定配置
[tool.poetry]
name = "my-project"
version = "1.0.0"

[tool.poetry.dependencies]
python = "^3.10"
requests = "^2.28.0"

[tool.poetry.group.dev.dependencies]
pytest = "^7.0.0"
```

```txt
# requirements.txt (传统方式)
requests==2.28.0
pydantic>=2.0.0,<3.0.0
numpy~=1.24.0              # 兼容版本

# requirements-dev.txt
-r requirements.txt        # 包含基础依赖
pytest==7.4.0
mypy==1.5.0
```

### Go 包管理

```bash
# ==================== go mod ====================

# 初始化模块
go mod init github.com/user/myproject

# 添加依赖 (自动)
go get github.com/gin-gonic/gin
go get github.com/gin-gonic/gin@v1.9.0   # 指定版本
go get github.com/gin-gonic/gin@latest

# 整理依赖
go mod tidy                              # 添加缺失/移除未用

# 下载依赖
go mod download

# 更新依赖
go get -u ./...                          # 更新所有
go get -u github.com/gin-gonic/gin       # 更新特定包

# 依赖图
go mod graph

# 验证
go mod verify

# vendor 模式
go mod vendor
go build -mod=vendor

# ==================== 工作空间 (Go 1.18+) ====================

# 创建工作空间
go work init ./app ./lib

# 添加模块
go work use ./another-module

# 同步
go work sync
```

```go
// go.mod
module github.com/user/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/go-sql-driver/mysql v1.7.1
    golang.org/x/sync v0.3.0
)

require (
    // indirect 表示间接依赖
    github.com/bytedance/sonic v1.9.1 // indirect
    golang.org/x/net v0.10.0 // indirect
)

// 替换依赖 (本地开发/fork)
replace github.com/original/pkg => ../local/pkg
replace github.com/original/pkg => github.com/fork/pkg v1.0.0

// 排除版本
exclude github.com/broken/pkg v1.0.0

// 撤回版本 (库作者使用)
retract (
    v1.0.0 // 有严重 bug
    [v1.1.0, v1.2.0] // 范围撤回
)
```

```go
// go.work (工作空间)
go 1.21

use (
    ./app
    ./lib
    ./shared
)

replace github.com/myorg/shared => ./shared
```

```bash
# ==================== 私有模块 ====================

# 设置私有仓库
go env -w GOPRIVATE=github.com/mycompany/*
go env -w GOPRIVATE=*.internal.company.com

# Git 配置 (使用 SSH)
git config --global url."git@github.com:".insteadOf "https://github.com/"

# 使用 token
git config --global url."https://${TOKEN}@github.com/".insteadOf "https://github.com/"
```

### Rust 包管理

```bash
# ==================== cargo ====================

# 创建项目
cargo new my-project                     # 二进制项目
cargo new --lib my-library               # 库项目
cargo init                               # 在当前目录初始化

# 构建
cargo build                              # 调试构建
cargo build --release                    # 发布构建

# 运行
cargo run
cargo run --release
cargo run --example demo                 # 运行示例

# 测试
cargo test
cargo test test_name                     # 运行特定测试
cargo test --doc                         # 文档测试

# 检查
cargo check                              # 快速检查 (不生成二进制)
cargo clippy                             # lint
cargo fmt                                # 格式化

# 依赖管理
cargo add serde                          # 添加依赖
cargo add serde --features derive        # 带 feature
cargo add tokio -F full                  # 简写
cargo add --dev mockall                  # 开发依赖
cargo remove serde                       # 移除

# 更新
cargo update                             # 更新所有
cargo update serde                       # 更新特定包

# 发布
cargo login
cargo publish
cargo publish --dry-run                  # 预检

# 文档
cargo doc --open                         # 生成并打开文档
```

```toml
# Cargo.toml
[package]
name = "my-project"
version = "1.0.0"
edition = "2021"
authors = ["Author <author@example.com>"]
description = "My awesome project"
license = "MIT"
repository = "https://github.com/user/my-project"
readme = "README.md"
keywords = ["cli", "tool"]
categories = ["command-line-utilities"]

[dependencies]
serde = { version = "1.0", features = ["derive"] }
tokio = { version = "1.0", features = ["full"] }
reqwest = { version = "0.11", default-features = false, features = ["json", "rustls-tls"] }

# 可选依赖
fancy-feature = { version = "1.0", optional = true }

# 本地依赖
my-lib = { path = "../my-lib" }

# Git 依赖
new-crate = { git = "https://github.com/user/new-crate", branch = "main" }

[dev-dependencies]
mockall = "0.11"
criterion = "0.5"

[build-dependencies]
cc = "1.0"

[features]
default = ["std"]
std = []
full = ["std", "fancy-feature"]

# 条件依赖
[target.'cfg(unix)'.dependencies]
nix = "0.26"

[target.'cfg(windows)'.dependencies]
windows = "0.48"

[[bin]]
name = "my-cli"
path = "src/bin/cli.rs"

[[example]]
name = "demo"
path = "examples/demo.rs"

[[bench]]
name = "my-bench"
harness = false

[profile.release]
opt-level = 3
lto = true
codegen-units = 1
strip = true

[profile.dev]
opt-level = 0
debug = true
```

```toml
# .cargo/config.toml (项目配置)
[build]
target = "x86_64-unknown-linux-gnu"

[target.x86_64-unknown-linux-gnu]
linker = "clang"

[registries.my-registry]
index = "https://my-registry.example.com/index"

[net]
git-fetch-with-cli = true

[alias]
b = "build"
t = "test"
r = "run"

# 使用国内镜像
[source.crates-io]
replace-with = 'ustc'

[source.ustc]
registry = "sparse+https://mirrors.ustc.edu.cn/crates.io-index/"
```

### 包管理对比总结

```
┌─────────────────┬────────────────────┬────────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript         │ Python             │ Go             │ Rust           │
├─────────────────┼────────────────────┼────────────────────┼────────────────┼────────────────┤
│ 包管理器        │ npm/yarn/pnpm      │ pip/poetry/uv      │ go mod         │ cargo          │
│ 配置文件        │ package.json       │ pyproject.toml     │ go.mod         │ Cargo.toml     │
│ 锁文件          │ package-lock.json  │ poetry.lock        │ go.sum         │ Cargo.lock     │
│ 中央仓库        │ npmjs.com          │ pypi.org           │ proxy.golang.org│ crates.io     │
│ 本地依赖        │ file:../path       │ path = "../path"   │ replace        │ path = "../"   │
│ Git 依赖        │ git+url            │ git+url            │ 自动支持       │ git = "url"    │
│ 工作空间        │ npm/yarn/pnpm      │ poetry             │ go work        │ workspace      │
│ 虚拟环境        │ 不需要             │ venv/poetry        │ 不需要         │ 不需要         │
│ 速度            │ pnpm 最快          │ uv 最快            │ 快             │ 快             │
└─────────────────┴────────────────────┴────────────────────┴────────────────┴────────────────┘
```

### 常用命令速查

```
┌─────────────────┬────────────────────┬────────────────────┬────────────────┬────────────────┐
│ 操作            │ npm/pnpm           │ poetry/uv          │ go             │ cargo          │
├─────────────────┼────────────────────┼────────────────────┼────────────────┼────────────────┤
│ 初始化          │ npm init           │ poetry new         │ go mod init    │ cargo new      │
│ 添加依赖        │ npm i pkg          │ poetry add pkg     │ go get pkg     │ cargo add pkg  │
│ 移除依赖        │ npm rm pkg         │ poetry remove pkg  │ (手动删除)     │ cargo rm pkg   │
│ 安装全部        │ npm install        │ poetry install     │ go mod download│ cargo build    │
│ 更新依赖        │ npm update         │ poetry update      │ go get -u      │ cargo update   │
│ 运行脚本        │ npm run dev        │ poetry run cmd     │ go run .       │ cargo run      │
│ 构建           │ npm run build      │ poetry build       │ go build       │ cargo build    │
│ 测试           │ npm test           │ poetry run pytest  │ go test        │ cargo test     │
│ 发布           │ npm publish        │ poetry publish     │ (git tag)      │ cargo publish  │
└─────────────────┴────────────────────┴────────────────────┴────────────────┴────────────────┘
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
