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

## 📚 总结

| 语言 | 一句话总结 |
|------|-----------|
| **TypeScript** | JavaScript 的类型安全升级版，Web 开发首选 |
| **Python** | 简单优雅，数据科学无敌，适合快速开发 |
| **Go** | 简单高效，云原生时代的工程语言 |
| **Rust** | 性能与安全的极致追求，系统编程新标杆 |

> **没有最好的语言，只有最适合的场景。** 选择语言时，应综合考虑项目需求、团队技能、生态系统和长期维护成本。
