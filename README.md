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

## 🔍 正则表达式

### 正则表达式概览

| 语言 | 正则引擎 | 语法风格 | 特点 |
|------|----------|----------|------|
| TypeScript | V8 RegExp | Perl 风格 | 内置、支持 /flag |
| Python | re 模块 | Perl 风格 | 功能完整、verbose 模式 |
| Go | regexp (RE2) | RE2 | 保证线性时间、无回溯 |
| Rust | regex crate | RE2 类似 | 高性能、无回溯 |

### TypeScript 正则表达式

```typescript
// ==================== 创建正则 ====================

// 字面量
const re1 = /hello/i;
const re2 = /\d{3}-\d{4}/g;

// 构造函数
const re3 = new RegExp('hello', 'i');
const re4 = new RegExp(`user_${userId}`, 'g');  // 动态模式

// ==================== 标志 (flags) ====================

/pattern/g    // global - 全局匹配
/pattern/i    // ignoreCase - 忽略大小写
/pattern/m    // multiline - 多行模式 (^$ 匹配行首行尾)
/pattern/s    // dotAll - . 匹配换行符
/pattern/u    // unicode - Unicode 模式
/pattern/y    // sticky - 粘性匹配
/pattern/d    // hasIndices - 返回匹配索引 (ES2022)

// 组合使用
/pattern/gim

// ==================== 匹配方法 ====================

const text = "Hello World, hello universe";
const re = /hello/gi;

// test - 返回布尔值
re.test(text);                    // true

// exec - 返回匹配详情 (带 g 时可迭代)
let match;
while ((match = re.exec(text)) !== null) {
  console.log(match[0], match.index);
}
// "Hello" 0
// "hello" 13

// match - 字符串方法
text.match(/hello/i);             // ["Hello", index: 0, ...]
text.match(/hello/gi);            // ["Hello", "hello"]

// matchAll - 返回迭代器 (ES2020)
for (const m of text.matchAll(/hello/gi)) {
  console.log(m[0], m.index);
}

// search - 返回索引
text.search(/world/i);            // 6

// ==================== 替换 ====================

// replace
text.replace(/hello/i, 'Hi');     // "Hi World, hello universe"
text.replace(/hello/gi, 'Hi');    // "Hi World, Hi universe"

// 使用函数
text.replace(/hello/gi, (match, offset) => {
  return match.toUpperCase();
});

// 替换模式
"John Smith".replace(/(\w+) (\w+)/, '$2, $1');  // "Smith, John"

// replaceAll (ES2021)
text.replaceAll('hello', 'Hi');   // 字符串替换
text.replaceAll(/hello/gi, 'Hi'); // 必须有 g 标志

// ==================== 分割 ====================

"a, b,  c".split(/,\s*/);         // ["a", "b", "c"]

// 保留分隔符
"a1b2c3".split(/(\d)/);           // ["a", "1", "b", "2", "c", "3"]

// ==================== 捕获组 ====================

const dateRe = /(\d{4})-(\d{2})-(\d{2})/;
const dateMatch = "2024-03-15".match(dateRe);
dateMatch[0];  // "2024-03-15" (完整匹配)
dateMatch[1];  // "2024" (第一个组)
dateMatch[2];  // "03"
dateMatch[3];  // "15"

// 命名捕获组 (ES2018)
const namedRe = /(?<year>\d{4})-(?<month>\d{2})-(?<day>\d{2})/;
const namedMatch = "2024-03-15".match(namedRe);
namedMatch.groups.year;   // "2024"
namedMatch.groups.month;  // "03"
namedMatch.groups.day;    // "15"

// 非捕获组
/(?:hello|hi) world/;     // 不创建捕获组

// 替换中使用命名组
"2024-03-15".replace(namedRe, '$<day>/$<month>/$<year>');
// "15/03/2024"

// ==================== 断言 ====================

// 前瞻断言
/foo(?=bar)/;    // 正向前瞻: foo 后面是 bar
/foo(?!bar)/;    // 负向前瞻: foo 后面不是 bar

// 后顾断言 (ES2018)
/(?<=@)\w+/;     // 正向后顾: @ 后面的单词
/(?<!@)\w+/;     // 负向后顾: 前面不是 @ 的单词

// 示例
"foobar foobaz".match(/foo(?=bar)/g);  // ["foo"]
"$100 €200".match(/(?<=\$)\d+/);       // ["100"]

// ==================== 常用模式 ====================

// 邮箱
const emailRe = /^[\w.-]+@[\w.-]+\.\w{2,}$/;

// URL
const urlRe = /^https?:\/\/[\w.-]+(?:\/[\w.-]*)*\/?$/;

// 手机号 (中国)
const phoneRe = /^1[3-9]\d{9}$/;

// IP 地址
const ipRe = /^(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?)$/;

// 转义特殊字符
function escapeRegex(str: string): string {
  return str.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
}
```

### Python 正则表达式

```python
import re

# ==================== 创建正则 ====================

# 原始字符串 (推荐)
pattern = r'\d{3}-\d{4}'

# 编译正则 (重复使用时推荐)
compiled = re.compile(r'\d{3}-\d{4}')
compiled = re.compile(r'''
    \d{3}   # 区号
    -       # 分隔符
    \d{4}   # 号码
''', re.VERBOSE)  # 允许注释和空白

# ==================== 标志 (flags) ====================

re.IGNORECASE  # re.I - 忽略大小写
re.MULTILINE   # re.M - 多行模式
re.DOTALL      # re.S - . 匹配换行
re.VERBOSE     # re.X - 允许注释
re.ASCII       # re.A - ASCII 匹配
re.UNICODE     # re.U - Unicode 匹配 (默认)

# 组合
re.compile(r'pattern', re.I | re.M)

# 内联标志
re.search(r'(?i)hello', text)  # 等同于 re.I

# ==================== 匹配方法 ====================

text = "Hello World, hello universe"

# search - 搜索第一个匹配
match = re.search(r'hello', text, re.I)
if match:
    match.group()     # "Hello"
    match.start()     # 0
    match.end()       # 5
    match.span()      # (0, 5)

# match - 从开头匹配
re.match(r'Hello', text)      # 匹配
re.match(r'World', text)      # None (不是从开头)

# fullmatch - 完整匹配
re.fullmatch(r'\d+', '123')   # 匹配
re.fullmatch(r'\d+', '123a')  # None

# findall - 返回所有匹配的列表
re.findall(r'hello', text, re.I)  # ['Hello', 'hello']

# 带组时返回元组
re.findall(r'(\w+)@(\w+)', "a@b c@d")  # [('a', 'b'), ('c', 'd')]

# finditer - 返回迭代器
for m in re.finditer(r'hello', text, re.I):
    print(m.group(), m.start())

# ==================== 替换 ====================

# sub - 替换
re.sub(r'hello', 'Hi', text, flags=re.I)  # "Hi World, Hi universe"

# 限制次数
re.sub(r'hello', 'Hi', text, count=1, flags=re.I)

# 使用函数
def upper_repl(match):
    return match.group().upper()
re.sub(r'hello', upper_repl, text, flags=re.I)

# 使用组引用
re.sub(r'(\w+) (\w+)', r'\2, \1', "John Smith")  # "Smith, John"

# subn - 返回 (新字符串, 替换次数)
re.subn(r'hello', 'Hi', text, flags=re.I)  # ("Hi World, Hi universe", 2)

# ==================== 分割 ====================

re.split(r',\s*', "a, b,  c")     # ['a', 'b', 'c']
re.split(r'(\d)', "a1b2c3")       # ['a', '1', 'b', '2', 'c', '3', '']
re.split(r',', "a,b,c", maxsplit=1)  # ['a', 'b,c']

# ==================== 捕获组 ====================

# 基本组
match = re.search(r'(\d{4})-(\d{2})-(\d{2})', "2024-03-15")
match.group(0)    # "2024-03-15" (完整匹配)
match.group(1)    # "2024"
match.group(2)    # "03"
match.groups()    # ('2024', '03', '15')

# 命名组
match = re.search(r'(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})', "2024-03-15")
match.group('year')   # "2024"
match.groupdict()     # {'year': '2024', 'month': '03', 'day': '15'}

# 替换中使用命名组
re.sub(r'(?P<y>\d{4})-(?P<m>\d{2})-(?P<d>\d{2})', r'\g<d>/\g<m>/\g<y>', "2024-03-15")
# "15/03/2024"

# 非捕获组
re.search(r'(?:hello|hi) world', "hello world")

# ==================== 断言 ====================

# 前瞻
re.search(r'foo(?=bar)', "foobar")    # 匹配 "foo"
re.search(r'foo(?!bar)', "foobaz")    # 匹配 "foo"

# 后顾
re.search(r'(?<=@)\w+', "user@domain")   # 匹配 "domain"
re.search(r'(?<!@)\w+', "hello")         # 匹配 "hello"

# ==================== 高级特性 ====================

# 条件匹配
# (?(id)yes|no) - 如果组 id 匹配，则使用 yes，否则 no
re.search(r'(<)?(\w+)(?(1)>|)', "<tag>")   # 匹配有/无括号的标签

# 原子组 (Python 3.11+)
re.search(r'(?>a+)b', "aaab")  # 原子组，不回溯

# 转义
re.escape('hello.world?')  # 'hello\\.world\\?'
```

### Go 正则表达式

```go
package main

import (
    "fmt"
    "regexp"
)

// ==================== 创建正则 ====================

func main() {
    // 编译正则
    re, err := regexp.Compile(`\d{3}-\d{4}`)
    if err != nil {
        panic(err)
    }
    
    // MustCompile - 编译失败时 panic
    re = regexp.MustCompile(`\d{3}-\d{4}`)
    
    // POSIX 语法
    rePosix := regexp.MustCompilePOSIX(`[[:digit:]]+`)
}

// ==================== 标志 (内联) ====================

/*
Go 使用内联标志:
(?i)  - 忽略大小写
(?m)  - 多行模式
(?s)  - . 匹配换行
(?U)  - 非贪婪模式
*/

re := regexp.MustCompile(`(?i)hello`)  // 忽略大小写
re = regexp.MustCompile(`(?m)^line`)   // 多行模式

// ==================== 匹配方法 ====================

text := "Hello World, hello universe"
re := regexp.MustCompile(`(?i)hello`)

// MatchString - 是否匹配
re.MatchString(text)              // true

// FindString - 返回第一个匹配
re.FindString(text)               // "Hello"

// FindAllString - 返回所有匹配
re.FindAllString(text, -1)        // ["Hello", "hello"]
re.FindAllString(text, 1)         // ["Hello"] (限制数量)

// FindStringIndex - 返回索引
re.FindStringIndex(text)          // [0, 5]

// FindAllStringIndex - 所有索引
re.FindAllStringIndex(text, -1)   // [[0, 5], [13, 18]]

// ==================== 捕获组 ====================

dateRe := regexp.MustCompile(`(\d{4})-(\d{2})-(\d{2})`)

// FindStringSubmatch - 返回组
match := dateRe.FindStringSubmatch("2024-03-15")
// ["2024-03-15", "2024", "03", "15"]

// FindAllStringSubmatch
matches := dateRe.FindAllStringSubmatch("2024-03-15 and 2025-01-01", -1)
// [["2024-03-15", "2024", "03", "15"], ["2025-01-01", "2025", "01", "01"]]

// 命名组
namedRe := regexp.MustCompile(`(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})`)
match = namedRe.FindStringSubmatch("2024-03-15")
names := namedRe.SubexpNames()  // ["", "year", "month", "day"]

// 获取命名组值
for i, name := range names {
    if name != "" {
        fmt.Printf("%s: %s\n", name, match[i])
    }
}

// ==================== 替换 ====================

// ReplaceAllString
re := regexp.MustCompile(`(?i)hello`)
re.ReplaceAllString(text, "Hi")   // "Hi World, Hi universe"

// 使用组引用
dateRe.ReplaceAllString("2024-03-15", "$3/$2/$1")  // "15/03/2024"

// ReplaceAllStringFunc - 使用函数
re.ReplaceAllStringFunc(text, strings.ToUpper)  // "HELLO World, HELLO universe"

// ReplaceAllLiteralString - 字面替换 (不解析 $)
re.ReplaceAllLiteralString(text, "$1")  // "$1 World, $1 universe"

// ==================== 分割 ====================

re := regexp.MustCompile(`,\s*`)
re.Split("a, b,  c", -1)          // ["a", "b", "c"]
re.Split("a, b,  c", 2)           // ["a", "b,  c"]

// ==================== 字节操作 ====================

// 所有方法都有字节版本
re := regexp.MustCompile(`\d+`)
re.Match([]byte("hello123"))              // true
re.Find([]byte("hello123"))               // []byte("123")
re.ReplaceAll([]byte("a1b2"), []byte("X")) // []byte("aXbX")

// ==================== 性能优化 ====================

// 预编译正则
var (
    emailRe = regexp.MustCompile(`^[\w.-]+@[\w.-]+\.\w{2,}$`)
    phoneRe = regexp.MustCompile(`^1[3-9]\d{9}$`)
)

// Longest - 最长匹配
re := regexp.MustCompile(`a+`)
re.Longest()
re.FindString("aaa")  // 总是返回最长匹配

// ==================== RE2 限制 ====================

/*
Go 使用 RE2 引擎，不支持:
- 反向引用 (\1, \2)
- 前瞻/后顾断言 (?=, (?!, (?<=, (?<!）
- 原子组 (?>)
- 占有量词 (++, *+, ?+)

优点: 保证 O(n) 时间复杂度，不会被恶意输入攻击
*/
```

### Rust 正则表达式

```rust
use regex::{Regex, RegexBuilder, Captures};

// ==================== 创建正则 ====================

fn main() {
    // 基本创建
    let re = Regex::new(r"\d{3}-\d{4}").unwrap();
    
    // 使用 RegexBuilder
    let re = RegexBuilder::new(r"hello")
        .case_insensitive(true)
        .multi_line(true)
        .build()
        .unwrap();
    
    // 编译时检查 (regex-macro)
    // let re = regex!(r"\d+");  // 编译时验证
}

// ==================== 标志 ====================

// 内联标志
let re = Regex::new(r"(?i)hello").unwrap();      // 忽略大小写
let re = Regex::new(r"(?m)^line").unwrap();      // 多行模式
let re = Regex::new(r"(?s).+").unwrap();         // . 匹配换行
let re = Regex::new(r"(?x)
    \d{4}   # year
    -
    \d{2}   # month
").unwrap();  // 允许注释和空白

// 组合
let re = Regex::new(r"(?im)hello").unwrap();

// ==================== 匹配方法 ====================

let text = "Hello World, hello universe";
let re = Regex::new(r"(?i)hello").unwrap();

// is_match - 是否匹配
re.is_match(text);                    // true

// find - 返回第一个匹配
if let Some(m) = re.find(text) {
    m.as_str();   // "Hello"
    m.start();    // 0
    m.end();      // 5
}

// find_iter - 迭代所有匹配
for m in re.find_iter(text) {
    println!("{} at {}", m.as_str(), m.start());
}

// ==================== 捕获组 ====================

let date_re = Regex::new(r"(\d{4})-(\d{2})-(\d{2})").unwrap();

// captures - 捕获组
if let Some(caps) = date_re.captures("2024-03-15") {
    caps.get(0).unwrap().as_str();  // "2024-03-15"
    caps.get(1).unwrap().as_str();  // "2024"
    caps.get(2).unwrap().as_str();  // "03"
    caps.get(3).unwrap().as_str();  // "15"
    
    // 使用索引
    &caps[0];  // "2024-03-15"
    &caps[1];  // "2024"
}

// captures_iter - 迭代所有捕获
for caps in date_re.captures_iter("2024-03-15 and 2025-01-01") {
    println!("{}", &caps[0]);
}

// 命名捕获组
let named_re = Regex::new(r"(?P<year>\d{4})-(?P<month>\d{2})-(?P<day>\d{2})").unwrap();
if let Some(caps) = named_re.captures("2024-03-15") {
    &caps["year"];    // "2024"
    &caps["month"];   // "03"
    &caps["day"];     // "15"
    
    caps.name("year").unwrap().as_str();  // "2024"
}

// ==================== 替换 ====================

let re = Regex::new(r"(?i)hello").unwrap();

// replace - 替换第一个
re.replace(text, "Hi");                   // "Hi World, hello universe"

// replace_all - 替换所有
re.replace_all(text, "Hi");               // "Hi World, Hi universe"

// 使用捕获组
let date_re = Regex::new(r"(\d{4})-(\d{2})-(\d{2})").unwrap();
date_re.replace("2024-03-15", "$3/$2/$1");  // "15/03/2024"

// 命名组
let named_re = Regex::new(r"(?P<y>\d{4})-(?P<m>\d{2})-(?P<d>\d{2})").unwrap();
named_re.replace("2024-03-15", "$d/$m/$y");  // "15/03/2024"

// 使用闭包
re.replace_all(text, |caps: &Captures| {
    caps[0].to_uppercase()
});  // "HELLO World, HELLO universe"

// replacen - 替换指定次数
re.replacen(text, 1, "Hi");               // "Hi World, hello universe"

// ==================== 分割 ====================

let re = Regex::new(r",\s*").unwrap();
let parts: Vec<&str> = re.split("a, b,  c").collect();
// ["a", "b", "c"]

let parts: Vec<&str> = re.splitn("a, b,  c", 2).collect();
// ["a", "b,  c"]

// ==================== RegexSet ====================

use regex::RegexSet;

// 同时匹配多个模式
let set = RegexSet::new(&[
    r"\d+",
    r"\w+",
    r"hello",
]).unwrap();

let matches: Vec<_> = set.matches("hello123").into_iter().collect();
// [0, 1, 2] - 匹配的模式索引

set.is_match("hello");  // true

// ==================== 性能优化 ====================

use once_cell::sync::Lazy;

// 全局预编译 (推荐)
static EMAIL_RE: Lazy<Regex> = Lazy::new(|| {
    Regex::new(r"^[\w.-]+@[\w.-]+\.\w{2,}$").unwrap()
});

static PHONE_RE: Lazy<Regex> = Lazy::new(|| {
    Regex::new(r"^1[3-9]\d{9}$").unwrap()
});

// 使用
fn is_valid_email(email: &str) -> bool {
    EMAIL_RE.is_match(email)
}

// ==================== bytes 版本 ====================

use regex::bytes::Regex as BytesRegex;

let re = BytesRegex::new(r"\d+").unwrap();
re.is_match(b"hello123");
re.find(b"hello123");

// ==================== regex 限制 ====================

/*
Rust regex 类似 RE2，不支持:
- 反向引用
- 前瞻/后顾断言

使用 fancy-regex crate 获取完整功能:
use fancy_regex::Regex;
let re = Regex::new(r"(?<=@)\w+").unwrap();
*/
```

### 正则表达式对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 引擎            │ V8 (回溯)      │ re (回溯)      │ RE2 (DFA)      │ regex (DFA)    │
│ 前瞻断言        │ ✅             │ ✅             │ ❌             │ ❌ (fancy-regex)│
│ 后顾断言        │ ✅             │ ✅             │ ❌             │ ❌ (fancy-regex)│
│ 反向引用        │ ✅             │ ✅             │ ❌             │ ❌ (fancy-regex)│
│ 命名组          │ ✅ (?<name>)   │ ✅ (?P<name>)  │ ✅ (?P<name>)  │ ✅ (?P<name>)  │
│ Unicode        │ ✅ /u          │ ✅ 默认        │ ✅ 默认        │ ✅ 默认        │
│ 预编译          │ new RegExp     │ re.compile     │ MustCompile    │ Regex::new     │
│ 时间复杂度      │ O(2^n) 最坏    │ O(2^n) 最坏    │ O(n) 保证      │ O(n) 保证      │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 常用正则模式

```
┌────────────────────┬──────────────────────────────────────────────────┐
│ 用途               │ 模式                                              │
├────────────────────┼──────────────────────────────────────────────────┤
│ 邮箱               │ ^[\w.-]+@[\w.-]+\.\w{2,}$                        │
│ URL                │ ^https?://[\w.-]+(?:/[\w.-]*)*/?$                │
│ 手机号 (中国)      │ ^1[3-9]\d{9}$                                    │
│ IP 地址            │ ^(?:\d{1,3}\.){3}\d{1,3}$                        │
│ 日期 (YYYY-MM-DD)  │ ^\d{4}-(?:0[1-9]|1[0-2])-(?:0[1-9]|[12]\d|3[01])$│
│ 时间 (HH:MM:SS)    │ ^(?:[01]\d|2[0-3]):[0-5]\d:[0-5]\d$              │
│ 十六进制颜色       │ ^#(?:[0-9a-fA-F]{3}){1,2}$                       │
│ 中文字符           │ [\u4e00-\u9fa5]+                                 │
│ 空白行             │ ^\s*$                                             │
│ HTML 标签          │ <([a-z]+)[^>]*>.*?</\1>                          │
└────────────────────┴──────────────────────────────────────────────────┘
```

---

## 🔄 迭代器与生成器

### 迭代器概览

| 语言 | 迭代器协议 | 生成器语法 | 惰性求值 | 特点 |
|------|------------|------------|----------|------|
| TypeScript | Symbol.iterator | function* | ✅ | 与 JS 一致 |
| Python | \_\_iter\_\_/\_\_next\_\_ | yield | ✅ | 最灵活 |
| Go | 无内置 (range) | 无 | 部分 | channel 模拟 |
| Rust | Iterator trait | 无 (迭代器适配器) | ✅ | 零成本抽象 |

### TypeScript 迭代器与生成器

```typescript
// ==================== 可迭代协议 ====================

// 实现 Iterable 接口
class Range implements Iterable<number> {
  constructor(
    private start: number,
    private end: number
  ) {}

  [Symbol.iterator](): Iterator<number> {
    let current = this.start;
    const end = this.end;
    
    return {
      next(): IteratorResult<number> {
        if (current <= end) {
          return { value: current++, done: false };
        }
        return { value: undefined, done: true };
      }
    };
  }
}

// 使用
for (const n of new Range(1, 5)) {
  console.log(n);  // 1, 2, 3, 4, 5
}

[...new Range(1, 3)];  // [1, 2, 3]

// ==================== 生成器函数 ====================

function* range(start: number, end: number): Generator<number> {
  for (let i = start; i <= end; i++) {
    yield i;
  }
}

// 使用
for (const n of range(1, 5)) {
  console.log(n);
}

// 生成器对象方法
const gen = range(1, 3);
gen.next();  // { value: 1, done: false }
gen.next();  // { value: 2, done: false }
gen.next();  // { value: 3, done: false }
gen.next();  // { value: undefined, done: true }

// ==================== 无限生成器 ====================

function* naturals(): Generator<number> {
  let n = 0;
  while (true) {
    yield n++;
  }
}

function* fibonacci(): Generator<number> {
  let [a, b] = [0, 1];
  while (true) {
    yield a;
    [a, b] = [b, a + b];
  }
}

// ==================== yield* 委托 ====================

function* concat<T>(...iterables: Iterable<T>[]): Generator<T> {
  for (const iterable of iterables) {
    yield* iterable;  // 委托给另一个可迭代对象
  }
}

[...concat([1, 2], [3, 4])];  // [1, 2, 3, 4]

// ==================== 双向通信 ====================

function* accumulator(): Generator<number, void, number> {
  let total = 0;
  while (true) {
    const value = yield total;
    if (value === undefined) break;
    total += value;
  }
}

const acc = accumulator();
acc.next();       // { value: 0 }
acc.next(10);     // { value: 10 }
acc.next(20);     // { value: 30 }
acc.return(0);    // 终止生成器

// ==================== 异步生成器 ====================

async function* fetchPages(urls: string[]): AsyncGenerator<Response> {
  for (const url of urls) {
    yield await fetch(url);
  }
}

async function* paginate<T>(
  fetchFn: (page: number) => Promise<T[]>
): AsyncGenerator<T> {
  let page = 1;
  while (true) {
    const items = await fetchFn(page++);
    if (items.length === 0) break;
    yield* items;
  }
}

// 使用
for await (const item of paginate(fetchPage)) {
  console.log(item);
}

// ==================== 迭代器工具函数 ====================

// take - 取前 n 个
function* take<T>(iterable: Iterable<T>, n: number): Generator<T> {
  let i = 0;
  for (const item of iterable) {
    if (i++ >= n) break;
    yield item;
  }
}

// map
function* map<T, U>(
  iterable: Iterable<T>,
  fn: (item: T) => U
): Generator<U> {
  for (const item of iterable) {
    yield fn(item);
  }
}

// filter
function* filter<T>(
  iterable: Iterable<T>,
  predicate: (item: T) => boolean
): Generator<T> {
  for (const item of iterable) {
    if (predicate(item)) yield item;
  }
}

// flatMap
function* flatMap<T, U>(
  iterable: Iterable<T>,
  fn: (item: T) => Iterable<U>
): Generator<U> {
  for (const item of iterable) {
    yield* fn(item);
  }
}

// zip
function* zip<T, U>(
  iter1: Iterable<T>,
  iter2: Iterable<U>
): Generator<[T, U]> {
  const it1 = iter1[Symbol.iterator]();
  const it2 = iter2[Symbol.iterator]();
  while (true) {
    const r1 = it1.next();
    const r2 = it2.next();
    if (r1.done || r2.done) break;
    yield [r1.value, r2.value];
  }
}

// 组合使用
const result = [...take(
  filter(
    map(naturals(), x => x * 2),
    x => x % 3 === 0
  ),
  5
)];  // [0, 6, 12, 18, 24]
```

### Python 迭代器与生成器

```python
from typing import Iterator, Generator, Iterable
from collections.abc import Iterator as ABCIterator

# ==================== 迭代器协议 ====================

class Range:
    def __init__(self, start: int, end: int):
        self.start = start
        self.end = end
    
    def __iter__(self) -> Iterator[int]:
        return RangeIterator(self.start, self.end)

class RangeIterator:
    def __init__(self, start: int, end: int):
        self.current = start
        self.end = end
    
    def __iter__(self) -> 'RangeIterator':
        return self
    
    def __next__(self) -> int:
        if self.current > self.end:
            raise StopIteration
        result = self.current
        self.current += 1
        return result

# 使用
for n in Range(1, 5):
    print(n)  # 1, 2, 3, 4, 5

list(Range(1, 3))  # [1, 2, 3]

# ==================== 生成器函数 ====================

def range_gen(start: int, end: int) -> Generator[int, None, None]:
    current = start
    while current <= end:
        yield current
        current += 1

# 使用
for n in range_gen(1, 5):
    print(n)

# 生成器对象
gen = range_gen(1, 3)
next(gen)  # 1
next(gen)  # 2
next(gen)  # 3
next(gen)  # StopIteration

# ==================== 生成器表达式 ====================

# 类似列表推导，但惰性求值
squares = (x**2 for x in range(10))
list(squares)  # [0, 1, 4, 9, 16, 25, 36, 49, 64, 81]

# 内存高效
sum(x**2 for x in range(1000000))  # 不创建列表

# 条件过滤
evens = (x for x in range(10) if x % 2 == 0)

# ==================== 无限生成器 ====================

def naturals() -> Generator[int, None, None]:
    n = 0
    while True:
        yield n
        n += 1

def fibonacci() -> Generator[int, None, None]:
    a, b = 0, 1
    while True:
        yield a
        a, b = b, a + b

# ==================== yield from ====================

def concat(*iterables):
    for iterable in iterables:
        yield from iterable

list(concat([1, 2], [3, 4]))  # [1, 2, 3, 4]

def flatten(nested):
    for item in nested:
        if isinstance(item, (list, tuple)):
            yield from flatten(item)
        else:
            yield item

list(flatten([1, [2, [3, 4]], 5]))  # [1, 2, 3, 4, 5]

# ==================== 双向通信 ====================

def accumulator() -> Generator[int, int, str]:
    total = 0
    while True:
        value = yield total
        if value is None:
            break
        total += value
    return f"Final: {total}"

acc = accumulator()
next(acc)           # 0 (启动生成器)
acc.send(10)        # 10
acc.send(20)        # 30
try:
    acc.send(None)  # 触发 return
except StopIteration as e:
    print(e.value)  # "Final: 30"

# throw 发送异常
gen.throw(ValueError("error"))

# close 关闭生成器
gen.close()

# ==================== 异步生成器 ====================

async def fetch_pages(urls: list[str]):
    for url in urls:
        response = await fetch(url)
        yield response

async def paginate(fetch_fn):
    page = 1
    while True:
        items = await fetch_fn(page)
        if not items:
            break
        for item in items:
            yield item
        page += 1

# 使用
async for item in paginate(fetch_page):
    print(item)

# 异步生成器表达式
async_gen = (await process(x) async for x in async_source)

# ==================== itertools ====================

import itertools

# count - 无限计数
itertools.count(10, 2)        # 10, 12, 14, 16, ...

# cycle - 无限循环
itertools.cycle([1, 2, 3])    # 1, 2, 3, 1, 2, 3, ...

# repeat - 重复
itertools.repeat('A', 3)      # A, A, A

# chain - 连接
itertools.chain([1, 2], [3, 4])  # 1, 2, 3, 4

# islice - 切片
itertools.islice(naturals(), 5)  # 0, 1, 2, 3, 4
itertools.islice(naturals(), 2, 5)  # 2, 3, 4

# takewhile / dropwhile
itertools.takewhile(lambda x: x < 5, range(10))  # 0, 1, 2, 3, 4
itertools.dropwhile(lambda x: x < 5, range(10))  # 5, 6, 7, 8, 9

# filterfalse
itertools.filterfalse(lambda x: x % 2, range(10))  # 0, 2, 4, 6, 8

# accumulate
itertools.accumulate([1, 2, 3, 4])  # 1, 3, 6, 10

# groupby
data = [('a', 1), ('a', 2), ('b', 3)]
for key, group in itertools.groupby(data, key=lambda x: x[0]):
    print(key, list(group))

# product - 笛卡尔积
itertools.product([1, 2], ['a', 'b'])  # (1,'a'), (1,'b'), (2,'a'), (2,'b')

# permutations / combinations
itertools.permutations([1, 2, 3], 2)   # 排列
itertools.combinations([1, 2, 3], 2)   # 组合

# zip_longest
itertools.zip_longest([1, 2], [3], fillvalue=0)  # (1, 3), (2, 0)

# starmap
itertools.starmap(pow, [(2, 3), (3, 2)])  # 8, 9

# ==================== more-itertools ====================

# pip install more-itertools
import more_itertools

more_itertools.chunked([1,2,3,4,5], 2)    # [[1,2], [3,4], [5]]
more_itertools.flatten([[1,2], [3,4]])     # 1, 2, 3, 4
more_itertools.first([1, 2, 3])            # 1
more_itertools.one([42])                   # 42 (单元素)
more_itertools.unique_everseen([1,2,1,3])  # 1, 2, 3
```

### Go 迭代器模式

```go
package main

import "iter"  // Go 1.23+

// ==================== 传统 for-range ====================

// 切片
for i, v := range []int{1, 2, 3} {
    fmt.Println(i, v)
}

// map
for k, v := range map[string]int{"a": 1} {
    fmt.Println(k, v)
}

// channel
ch := make(chan int)
for v := range ch {
    fmt.Println(v)
}

// 字符串 (遍历 rune)
for i, r := range "hello" {
    fmt.Println(i, string(r))
}

// ==================== Go 1.23+ 迭代器 ====================

// iter.Seq[V] - 单值序列
func Range(start, end int) iter.Seq[int] {
    return func(yield func(int) bool) {
        for i := start; i <= end; i++ {
            if !yield(i) {
                return
            }
        }
    }
}

// 使用
for n := range Range(1, 5) {
    fmt.Println(n)  // 1, 2, 3, 4, 5
}

// iter.Seq2[K, V] - 键值对序列
func Enumerate[T any](s []T) iter.Seq2[int, T] {
    return func(yield func(int, T) bool) {
        for i, v := range s {
            if !yield(i, v) {
                return
            }
        }
    }
}

for i, v := range Enumerate([]string{"a", "b", "c"}) {
    fmt.Println(i, v)
}

// ==================== 迭代器组合 ====================

// Map
func Map[T, U any](seq iter.Seq[T], fn func(T) U) iter.Seq[U] {
    return func(yield func(U) bool) {
        for v := range seq {
            if !yield(fn(v)) {
                return
            }
        }
    }
}

// Filter
func Filter[T any](seq iter.Seq[T], pred func(T) bool) iter.Seq[T] {
    return func(yield func(T) bool) {
        for v := range seq {
            if pred(v) {
                if !yield(v) {
                    return
                }
            }
        }
    }
}

// Take
func Take[T any](seq iter.Seq[T], n int) iter.Seq[T] {
    return func(yield func(T) bool) {
        i := 0
        for v := range seq {
            if i >= n {
                return
            }
            if !yield(v) {
                return
            }
            i++
        }
    }
}

// 组合使用
for v := range Take(Filter(Map(Range(1, 100), 
    func(x int) int { return x * 2 }),
    func(x int) bool { return x%3 == 0 }),
    5) {
    fmt.Println(v)  // 6, 12, 18, 24, 30
}

// ==================== slices/maps 包 ====================

import "slices"
import "maps"

// slices.All - 转为迭代器
for i, v := range slices.All([]int{1, 2, 3}) {
    fmt.Println(i, v)
}

// slices.Values - 仅值
for v := range slices.Values([]int{1, 2, 3}) {
    fmt.Println(v)
}

// slices.Collect - 收集为切片
result := slices.Collect(Range(1, 5))  // []int{1, 2, 3, 4, 5}

// maps.All
for k, v := range maps.All(myMap) {
    fmt.Println(k, v)
}

// ==================== Channel 模拟生成器 ====================

func Fibonacci(n int) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        a, b := 0, 1
        for i := 0; i < n; i++ {
            ch <- a
            a, b = b, a+b
        }
    }()
    return ch
}

for f := range Fibonacci(10) {
    fmt.Println(f)
}

// 带取消
func FibonacciWithCancel(ctx context.Context) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        a, b := 0, 1
        for {
            select {
            case <-ctx.Done():
                return
            case ch <- a:
                a, b = b, a+b
            }
        }
    }()
    return ch
}

// ==================== 传统回调模式 ====================

// 在 Go 1.23 之前的常用模式
type IntIterator struct {
    current int
    end     int
}

func NewIntIterator(start, end int) *IntIterator {
    return &IntIterator{current: start, end: end}
}

func (it *IntIterator) HasNext() bool {
    return it.current <= it.end
}

func (it *IntIterator) Next() int {
    v := it.current
    it.current++
    return v
}

// 使用
it := NewIntIterator(1, 5)
for it.HasNext() {
    fmt.Println(it.Next())
}
```

### Rust 迭代器

```rust
// ==================== Iterator trait ====================

// 自定义迭代器
struct Range {
    current: i32,
    end: i32,
}

impl Range {
    fn new(start: i32, end: i32) -> Self {
        Range { current: start, end }
    }
}

impl Iterator for Range {
    type Item = i32;
    
    fn next(&mut self) -> Option<Self::Item> {
        if self.current <= self.end {
            let result = self.current;
            self.current += 1;
            Some(result)
        } else {
            None
        }
    }
}

// 使用
for n in Range::new(1, 5) {
    println!("{}", n);
}

// ==================== IntoIterator trait ====================

// 实现 IntoIterator 以支持 for 循环
struct MyCollection {
    items: Vec<i32>,
}

impl IntoIterator for MyCollection {
    type Item = i32;
    type IntoIter = std::vec::IntoIter<i32>;
    
    fn into_iter(self) -> Self::IntoIter {
        self.items.into_iter()
    }
}

// 三种迭代方式
let v = vec![1, 2, 3];
for x in v.iter() {}      // &T (借用)
for x in v.iter_mut() {}  // &mut T (可变借用)
for x in v.into_iter() {} // T (获取所有权)

// ==================== 迭代器适配器 ====================

let v = vec![1, 2, 3, 4, 5];

// map - 转换
v.iter().map(|x| x * 2);

// filter - 过滤
v.iter().filter(|x| **x % 2 == 0);

// filter_map - 过滤并转换
v.iter().filter_map(|x| if *x > 2 { Some(x * 2) } else { None });

// take / skip
v.iter().take(3);         // 前 3 个
v.iter().skip(2);         // 跳过前 2 个
v.iter().take_while(|x| **x < 4);
v.iter().skip_while(|x| **x < 3);

// enumerate
for (i, x) in v.iter().enumerate() {
    println!("{}: {}", i, x);
}

// zip
let a = [1, 2, 3];
let b = ["a", "b", "c"];
for (x, y) in a.iter().zip(b.iter()) {
    println!("{} {}", x, y);
}

// chain - 连接
let c = a.iter().chain(b.iter());

// flatten
let nested = vec![vec![1, 2], vec![3, 4]];
nested.into_iter().flatten();  // 1, 2, 3, 4

// flat_map
v.iter().flat_map(|x| vec![*x, x * 10]);

// rev - 反转
v.iter().rev();

// cycle - 无限循环
v.iter().cycle().take(10);

// inspect - 调试
v.iter()
    .inspect(|x| println!("before: {}", x))
    .map(|x| x * 2)
    .inspect(|x| println!("after: {}", x));

// peekable - 预览
let mut iter = v.iter().peekable();
if iter.peek() == Some(&&1) {
    // ...
}

// ==================== 消费者 ====================

let v = vec![1, 2, 3, 4, 5];

// collect - 收集
let doubled: Vec<i32> = v.iter().map(|x| x * 2).collect();
let set: HashSet<_> = v.iter().collect();
let s: String = ['h', 'e', 'l', 'l', 'o'].iter().collect();

// sum / product
let sum: i32 = v.iter().sum();
let product: i32 = v.iter().product();

// fold - 折叠
let sum = v.iter().fold(0, |acc, x| acc + x);

// reduce
let sum = v.iter().copied().reduce(|a, b| a + b);

// count / min / max
v.iter().count();
v.iter().min();
v.iter().max();
v.iter().min_by_key(|x| x.abs());

// find / position
v.iter().find(|x| **x > 3);       // Some(&4)
v.iter().position(|x| *x > 3);    // Some(3)

// any / all
v.iter().any(|x| *x > 3);         // true
v.iter().all(|x| *x > 0);         // true

// for_each
v.iter().for_each(|x| println!("{}", x));

// partition
let (evens, odds): (Vec<_>, Vec<_>) = v.iter().partition(|x| **x % 2 == 0);

// ==================== 链式调用示例 ====================

let result: Vec<i32> = (0..100)
    .filter(|x| x % 2 == 0)    // 偶数
    .map(|x| x * x)            // 平方
    .filter(|x| x % 3 == 0)    // 能被 3 整除
    .take(5)                    // 前 5 个
    .collect();
// [0, 36, 144, 324, 576]

// ==================== 自定义适配器 ====================

// 使用 std::iter::from_fn
fn fibonacci() -> impl Iterator<Item = u64> {
    let mut state = (0, 1);
    std::iter::from_fn(move || {
        let result = state.0;
        state = (state.1, state.0 + state.1);
        Some(result)
    })
}

// 使用 std::iter::successors
let powers_of_2 = std::iter::successors(Some(1), |&n| Some(n * 2));

// 使用 std::iter::repeat_with
let randoms = std::iter::repeat_with(|| rand::random::<u32>());

// ==================== 并行迭代器 (rayon) ====================

use rayon::prelude::*;

let sum: i32 = (0..1000000)
    .into_par_iter()          // 并行迭代器
    .filter(|x| x % 2 == 0)
    .map(|x| x * 2)
    .sum();

// 并行排序
let mut data = vec![3, 1, 4, 1, 5];
data.par_sort();

// ==================== 异步迭代器 (Stream) ====================

use futures::stream::{self, StreamExt};

async fn example() {
    let stream = stream::iter(vec![1, 2, 3]);
    
    // map / filter
    let doubled = stream.map(|x| x * 2);
    
    // collect
    let v: Vec<_> = stream.collect().await;
    
    // for_each
    stream::iter(vec![1, 2, 3])
        .for_each(|x| async move {
            println!("{}", x);
        })
        .await;
}
```

### 迭代器对比总结

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 迭代器协议      │ Symbol.iterator│ __iter__       │ iter.Seq (1.23)│ Iterator trait │
│ 生成器语法      │ function*      │ yield          │ 无 (channel)   │ 无             │
│ 惰性求值        │ ✅             │ ✅             │ ✅ (1.23+)     │ ✅             │
│ 异步迭代        │ async function*│ async for      │ channel        │ Stream trait   │
│ 标准库工具      │ 少             │ itertools      │ slices/maps    │ 丰富的适配器   │
│ 并行迭代        │ 无内置         │ multiprocessing│ goroutine      │ rayon          │
│ 零成本抽象      │ ❌             │ ❌             │ ❌             │ ✅             │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 常用操作速查

```
┌────────────────┬──────────────────┬──────────────────┬──────────────────┬──────────────────┐
│ 操作           │ TypeScript       │ Python           │ Go (1.23+)       │ Rust             │
├────────────────┼──────────────────┼──────────────────┼──────────────────┼──────────────────┤
│ 创建范围       │ range(1,5)       │ range(1,6)       │ Range(1,5)       │ 1..=5            │
│ 映射           │ map(fn)          │ map(fn, iter)    │ Map(seq, fn)     │ .map(fn)         │
│ 过滤           │ filter(fn)       │ filter(fn, iter) │ Filter(seq, fn)  │ .filter(fn)      │
│ 取前N个        │ take(n)          │ islice(iter,n)   │ Take(seq, n)     │ .take(n)         │
│ 跳过N个        │ skip(n)          │ islice(iter,n,..)│ 手动             │ .skip(n)         │
│ 折叠           │ reduce(fn)       │ reduce(fn, iter) │ 手动             │ .fold(init,fn)   │
│ 收集           │ [...iter]        │ list(iter)       │ slices.Collect   │ .collect()       │
│ 枚举           │ entries()        │ enumerate(iter)  │ Enumerate        │ .enumerate()     │
│ 压缩           │ zip(a, b)        │ zip(a, b)        │ 手动             │ .zip(other)      │
│ 连接           │ concat(a, b)     │ chain(a, b)      │ 手动             │ .chain(other)    │
│ 扁平化         │ flat()           │ chain.from_iter  │ 手动             │ .flatten()       │
│ 任一满足       │ some(fn)         │ any(fn, iter)    │ 手动             │ .any(fn)         │
│ 全部满足       │ every(fn)        │ all(fn, iter)    │ 手动             │ .all(fn)         │
└────────────────┴──────────────────┴──────────────────┴──────────────────┴──────────────────┘
```

---

## 🚫 空值处理

### 空值类型概览

| 语言 | 空值类型 | 安全机制 | 特点 |
|------|----------|----------|------|
| TypeScript | `null`, `undefined` | 可选链、空值合并 | 两种空值 |
| Python | `None` | or、if 检查 | 单一空值 |
| Go | `nil` | if 检查 | 零值概念 |
| Rust | 无 null | `Option<T>` | 编译时保证 |

### TypeScript 空值处理

```typescript
// ==================== null vs undefined ====================
let a: null = null;           // 显式空
let b: undefined = undefined; // 未定义
let c: string | null = null;  // 可空类型

// 区别
typeof null;       // "object" (历史遗留)
typeof undefined;  // "undefined"

// ==================== 可选链 (?.) ====================
const name = user?.profile?.name;
const first = arr?.[0];
const result = obj?.method?.();

// ==================== 空值合并 (??) ====================
const value = input ?? "default";  // 仅 null/undefined
const name = user.name ?? "Anonymous";

// 与 || 区别
0 || "default";    // "default" (0 是 falsy)
0 ?? "default";    // 0 (仅 null/undefined 触发)

// ==================== 非空断言 (!) ====================
const name = user!.name;  // 告诉编译器不为空 (危险)

// ==================== 类型守卫 ====================
function process(value: string | null) {
  if (value === null) return;
  console.log(value.toUpperCase());  // 类型收窄
}

function isNotNull<T>(value: T | null): value is T {
  return value !== null;
}

const items = [1, null, 2].filter(isNotNull);  // number[]

// ==================== 可选参数与属性 ====================
interface User {
  name: string;
  email?: string;  // string | undefined
}

function greet(name: string, title?: string) {
  return title ? `${title} ${name}` : name;
}
```

### Python 空值处理

```python
# ==================== None 基础 ====================
value = None
value is None      # True (推荐用 is)
value == None      # True (但不推荐)

# ==================== 类型注解 ====================
from typing import Optional

def find_user(id: int) -> Optional[str]:  # str | None
    return None if id < 0 else "User"

def process(name: str | None) -> str:  # Python 3.10+
    return name or "default"

# ==================== or 默认值 ====================
name = user_name or "Anonymous"

# 注意: 空字符串、0、[] 也会触发
"" or "default"    # "default"
0 or 42            # 42

# ==================== 条件表达式 ====================
name = user_name if user_name is not None else "default"

# ==================== walrus 运算符 ====================
if (match := re.search(pattern, text)) is not None:
    print(match.group())

# ==================== getattr / get ====================
value = getattr(obj, 'attr', 'default')
value = dict.get('key', 'default')

# ==================== 链式访问 (需要库) ====================
# pip install glom
from glom import glom
name = glom(user, 'profile.name', default=None)
```

### Go 空值处理

```go
// ==================== nil 与零值 ====================
var s string    // "" (零值)
var n int       // 0
var b bool      // false
var p *int      // nil
var sl []int    // nil
var m map[string]int  // nil
var ch chan int // nil
var fn func()   // nil
var i interface{} // nil

// ==================== nil 检查 ====================
if p != nil {
    fmt.Println(*p)
}

if m == nil {
    m = make(map[string]int)
}

// ==================== 指针与可选值 ====================
type User struct {
    Name  string
    Email *string  // 可选字段
}

func NewUser(name string, email *string) User {
    return User{Name: name, Email: email}
}

// 使用
email := "test@example.com"
user := NewUser("John", &email)
user2 := NewUser("Jane", nil)

// ==================== comma ok 模式 ====================
value, ok := m["key"]
if !ok {
    // 不存在
}

if v, ok := m["key"]; ok {
    fmt.Println(v)
}

// 类型断言
if str, ok := i.(string); ok {
    fmt.Println(str)
}

// ==================== 默认值函数 ====================
func OrDefault[T any](ptr *T, def T) T {
    if ptr == nil {
        return def
    }
    return *ptr
}

name := OrDefault(user.Nickname, "Anonymous")
```

### Rust 空值处理

```rust
// ==================== Option<T> ====================
let some_value: Option<i32> = Some(42);
let no_value: Option<i32> = None;

// ==================== 模式匹配 ====================
match some_value {
    Some(v) => println!("Value: {}", v),
    None => println!("No value"),
}

// if let
if let Some(v) = some_value {
    println!("Value: {}", v);
}

// let else
let Some(v) = some_value else {
    println!("No value");
    return;
};

// ==================== 方法链 ====================
// unwrap - 有值返回，None 则 panic
some_value.unwrap();

// expect - 带错误消息的 unwrap
some_value.expect("Should have value");

// unwrap_or - 默认值
no_value.unwrap_or(0);  // 0

// unwrap_or_else - 惰性默认值
no_value.unwrap_or_else(|| compute_default());

// unwrap_or_default - 类型默认值
no_value.unwrap_or_default();  // 0

// ==================== 转换方法 ====================
// map - 转换内部值
some_value.map(|v| v * 2);  // Some(84)

// and_then - 链式 Option
some_value.and_then(|v| if v > 0 { Some(v) } else { None });

// filter
some_value.filter(|v| *v > 0);

// ok_or - 转为 Result
some_value.ok_or("No value")?;

// ==================== ? 操作符 ====================
fn get_name(user: Option<User>) -> Option<String> {
    let user = user?;  // None 则提前返回
    let profile = user.profile?;
    Some(profile.name)
}

// ==================== Option 组合 ====================
let a: Option<i32> = Some(1);
let b: Option<i32> = Some(2);

// zip
a.zip(b);  // Some((1, 2))

// or
None.or(Some(1));  // Some(1)

// and
Some(1).and(Some(2));  // Some(2)
```

---

## 🔧 函数

### 函数特性概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 一等公民 | ✅ | ✅ | ✅ | ✅ |
| 闭包 | ✅ | ✅ | ✅ | ✅ |
| 默认参数 | ✅ | ✅ | ❌ | ❌ |
| 命名参数 | ❌ (对象模拟) | ✅ | ❌ | ❌ |
| 可变参数 | ✅ | ✅ | ✅ | ❌ (宏) |
| 多返回值 | ❌ (元组模拟) | ✅ | ✅ | ❌ (元组) |
| 重载 | ✅ (类型) | ❌ | ❌ | ✅ (trait) |

### TypeScript 函数

```typescript
// ==================== 函数声明 ====================
function add(a: number, b: number): number {
  return a + b;
}

// 函数表达式
const multiply = function(a: number, b: number): number {
  return a * b;
};

// 箭头函数
const divide = (a: number, b: number): number => a / b;

// ==================== 参数 ====================
// 默认参数
function greet(name: string, greeting = "Hello"): string {
  return `${greeting}, ${name}!`;
}

// 可选参数
function log(message: string, level?: string): void {
  console.log(level ? `[${level}] ${message}` : message);
}

// 剩余参数
function sum(...numbers: number[]): number {
  return numbers.reduce((a, b) => a + b, 0);
}

// 解构参数
function createUser({ name, age = 18 }: { name: string; age?: number }) {
  return { name, age };
}

// ==================== 函数类型 ====================
type MathFn = (a: number, b: number) => number;
type Callback = (error: Error | null, result?: string) => void;

interface Calculator {
  add: MathFn;
  subtract: MathFn;
}

// ==================== 闭包 ====================
function counter() {
  let count = 0;
  return {
    increment: () => ++count,
    decrement: () => --count,
    get: () => count,
  };
}

// ==================== 重载 ====================
function process(x: string): string;
function process(x: number): number;
function process(x: string | number): string | number {
  return typeof x === "string" ? x.toUpperCase() : x * 2;
}

// ==================== 泛型函数 ====================
function identity<T>(value: T): T {
  return value;
}

function map<T, U>(arr: T[], fn: (item: T) => U): U[] {
  return arr.map(fn);
}

// ==================== 高阶函数 ====================
const compose = <T>(...fns: ((x: T) => T)[]) =>
  (x: T) => fns.reduceRight((acc, fn) => fn(acc), x);

const pipe = <T>(...fns: ((x: T) => T)[]) =>
  (x: T) => fns.reduce((acc, fn) => fn(acc), x);
```

### Python 函数

```python
from typing import Callable, TypeVar, ParamSpec

# ==================== 函数定义 ====================
def add(a: int, b: int) -> int:
    return a + b

# lambda
multiply = lambda a, b: a * b

# ==================== 参数 ====================
# 默认参数
def greet(name: str, greeting: str = "Hello") -> str:
    return f"{greeting}, {name}!"

# 可变位置参数
def sum_all(*numbers: int) -> int:
    return sum(numbers)

# 可变关键字参数
def create_user(**kwargs) -> dict:
    return kwargs

# 混合参数
def func(pos1, pos2, /, pos_or_kw, *, kw_only):
    pass
# pos1, pos2: 仅位置参数
# pos_or_kw: 位置或关键字
# kw_only: 仅关键字参数

# 解包调用
args = [1, 2, 3]
func(*args)
kwargs = {"name": "John", "age": 30}
func(**kwargs)

# ==================== 类型注解 ====================
Callback = Callable[[str, int], bool]

T = TypeVar('T')
P = ParamSpec('P')

def decorator(fn: Callable[P, T]) -> Callable[P, T]:
    def wrapper(*args: P.args, **kwargs: P.kwargs) -> T:
        return fn(*args, **kwargs)
    return wrapper

# ==================== 闭包 ====================
def counter():
    count = 0
    def increment():
        nonlocal count  # 访问外部变量
        count += 1
        return count
    return increment

# ==================== 装饰器 ====================
def log_calls(func):
    @functools.wraps(func)
    def wrapper(*args, **kwargs):
        print(f"Calling {func.__name__}")
        return func(*args, **kwargs)
    return wrapper

@log_calls
def my_function():
    pass

# ==================== 高阶函数 ====================
from functools import reduce, partial

# map, filter, reduce
list(map(lambda x: x * 2, [1, 2, 3]))
list(filter(lambda x: x > 1, [1, 2, 3]))
reduce(lambda a, b: a + b, [1, 2, 3])

# partial
def power(base, exp):
    return base ** exp
square = partial(power, exp=2)
```

### Go 函数

```go
// ==================== 函数定义 ====================
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

// ==================== 可变参数 ====================
func sum(numbers ...int) int {
    total := 0
    for _, n := range numbers {
        total += n
    }
    return total
}

// 展开切片
nums := []int{1, 2, 3}
sum(nums...)

// ==================== 函数类型 ====================
type MathFunc func(int, int) int
type Handler func(w http.ResponseWriter, r *http.Request)

// ==================== 闭包 ====================
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

// ==================== 匿名函数 ====================
func() {
    fmt.Println("Immediately invoked")
}()

sort.Slice(items, func(i, j int) bool {
    return items[i] < items[j]
})

// ==================== 方法 ====================
type Rectangle struct {
    Width, Height float64
}

// 值接收者
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 指针接收者
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}

// ==================== 泛型函数 (Go 1.18+) ====================
func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

func Filter[T any](slice []T, pred func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if pred(v) {
            result = append(result, v)
        }
    }
    return result
}
```

### Rust 函数

```rust
// ==================== 函数定义 ====================
fn add(a: i32, b: i32) -> i32 {
    a + b  // 无分号 = 返回值
}

// 无返回值
fn greet(name: &str) {
    println!("Hello, {}!", name);
}

// 返回 Result
fn divide(a: i32, b: i32) -> Result<i32, &'static str> {
    if b == 0 {
        Err("division by zero")
    } else {
        Ok(a / b)
    }
}

// ==================== 闭包 ====================
let add = |a, b| a + b;
let add_typed = |a: i32, b: i32| -> i32 { a + b };

// 捕获环境
let x = 10;
let add_x = |n| n + x;     // 借用 x
let take_x = move |n| n + x; // 移动 x

// ==================== 闭包 trait ====================
// Fn: 借用捕获 (&self)
// FnMut: 可变借用 (&mut self)
// FnOnce: 获取所有权 (self)

fn apply<F>(f: F) where F: Fn(i32) -> i32 {
    println!("{}", f(10));
}

fn apply_mut<F>(mut f: F) where F: FnMut(i32) -> i32 {
    println!("{}", f(10));
}

fn apply_once<F>(f: F) where F: FnOnce(i32) -> i32 {
    println!("{}", f(10));
}

// ==================== 函数指针 ====================
fn apply_fn(f: fn(i32) -> i32, x: i32) -> i32 {
    f(x)
}

// ==================== 泛型函数 ====================
fn identity<T>(value: T) -> T {
    value
}

fn map<T, U, F>(vec: Vec<T>, f: F) -> Vec<U>
where
    F: Fn(T) -> U,
{
    vec.into_iter().map(f).collect()
}

// ==================== impl Trait ====================
fn make_adder(x: i32) -> impl Fn(i32) -> i32 {
    move |y| x + y
}

fn iter_numbers() -> impl Iterator<Item = i32> {
    (0..10).filter(|x| x % 2 == 0)
}

// ==================== 方法 ====================
struct Rectangle {
    width: f64,
    height: f64,
}

impl Rectangle {
    // 关联函数 (构造器)
    fn new(width: f64, height: f64) -> Self {
        Self { width, height }
    }
    
    // 方法
    fn area(&self) -> f64 {
        self.width * self.height
    }
    
    fn scale(&mut self, factor: f64) {
        self.width *= factor;
        self.height *= factor;
    }
}
```

---

## 🌐 HTTP 客户端 (Fetch)

### HTTP 客户端概览

| 语言 | 内置/推荐库 | 异步支持 | 特点 |
|------|-------------|----------|------|
| TypeScript | fetch / axios | ✅ | 原生 Promise |
| Python | requests / httpx | ✅ (httpx) | 简洁 API |
| Go | net/http | ✅ (goroutine) | 标准库完善 |
| Rust | reqwest | ✅ | 类型安全 |

### TypeScript HTTP 请求

```typescript
// ==================== fetch API ====================
// GET
const response = await fetch('https://api.example.com/users');
const users = await response.json();

// POST
const newUser = await fetch('https://api.example.com/users', {
  method: 'POST',
  headers: { 'Content-Type': 'application/json' },
  body: JSON.stringify({ name: 'John', age: 30 }),
});

// 完整示例
async function fetchData<T>(url: string): Promise<T> {
  const response = await fetch(url);
  if (!response.ok) {
    throw new Error(`HTTP ${response.status}`);
  }
  return response.json();
}

// ==================== 超时与取消 ====================
const controller = new AbortController();
setTimeout(() => controller.abort(), 5000);

const response = await fetch(url, {
  signal: controller.signal,
});

// ==================== axios ====================
import axios from 'axios';

// 创建实例
const api = axios.create({
  baseURL: 'https://api.example.com',
  timeout: 5000,
  headers: { 'Authorization': 'Bearer token' },
});

// 请求
const { data } = await api.get('/users');
await api.post('/users', { name: 'John' });

// 拦截器
api.interceptors.request.use(config => {
  config.headers.Authorization = `Bearer ${getToken()}`;
  return config;
});

api.interceptors.response.use(
  response => response,
  error => {
    if (error.response?.status === 401) {
      // 处理认证错误
    }
    return Promise.reject(error);
  }
);
```

### Python HTTP 请求

```python
import requests
import httpx

# ==================== requests (同步) ====================
# GET
response = requests.get('https://api.example.com/users')
users = response.json()

# POST
response = requests.post(
    'https://api.example.com/users',
    json={'name': 'John', 'age': 30},
    headers={'Authorization': 'Bearer token'},
    timeout=5
)

# Session (连接复用)
with requests.Session() as session:
    session.headers.update({'Authorization': 'Bearer token'})
    response = session.get('/users')

# ==================== httpx (同步+异步) ====================
# 同步
response = httpx.get('https://api.example.com/users')

# 异步
async with httpx.AsyncClient() as client:
    response = await client.get('https://api.example.com/users')
    users = response.json()

# 异步并发
async with httpx.AsyncClient() as client:
    tasks = [client.get(url) for url in urls]
    responses = await asyncio.gather(*tasks)

# 客户端配置
client = httpx.AsyncClient(
    base_url='https://api.example.com',
    timeout=5.0,
    headers={'Authorization': 'Bearer token'},
)
```

### Go HTTP 请求

```go
import (
    "net/http"
    "encoding/json"
    "bytes"
    "time"
)

// ==================== 基本请求 ====================
// GET
resp, err := http.Get("https://api.example.com/users")
if err != nil {
    return err
}
defer resp.Body.Close()

var users []User
json.NewDecoder(resp.Body).Decode(&users)

// POST
data := map[string]interface{}{"name": "John", "age": 30}
jsonData, _ := json.Marshal(data)

resp, err := http.Post(
    "https://api.example.com/users",
    "application/json",
    bytes.NewBuffer(jsonData),
)

// ==================== 自定义客户端 ====================
client := &http.Client{
    Timeout: 5 * time.Second,
    Transport: &http.Transport{
        MaxIdleConns:        100,
        MaxIdleConnsPerHost: 10,
    },
}

// 自定义请求
req, _ := http.NewRequest("GET", url, nil)
req.Header.Set("Authorization", "Bearer token")

resp, err := client.Do(req)

// ==================== 带 Context ====================
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
defer cancel()

req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
resp, err := client.Do(req)
```

### Rust HTTP 请求

```rust
use reqwest;

// ==================== 基本请求 ====================
// GET
let users: Vec<User> = reqwest::get("https://api.example.com/users")
    .await?
    .json()
    .await?;

// POST
let client = reqwest::Client::new();
let new_user = client
    .post("https://api.example.com/users")
    .json(&serde_json::json!({"name": "John", "age": 30}))
    .send()
    .await?;

// ==================== 客户端配置 ====================
let client = reqwest::Client::builder()
    .timeout(Duration::from_secs(5))
    .default_headers({
        let mut headers = reqwest::header::HeaderMap::new();
        headers.insert("Authorization", "Bearer token".parse().unwrap());
        headers
    })
    .build()?;

// ==================== 错误处理 ====================
let response = client.get(url).send().await?;

if response.status().is_success() {
    let data: Data = response.json().await?;
} else {
    eprintln!("Error: {}", response.status());
}
```

---

## 🖥️ HTTP & RPC 服务器

### 框架概览

| 语言 | HTTP 框架 | RPC 框架 |
|------|-----------|----------|
| TypeScript | Express, Fastify, Hono | tRPC, gRPC |
| Python | FastAPI, Flask | gRPC, Thrift |
| Go | net/http, Gin, Echo | gRPC, Twirp |
| Rust | Axum, Actix-web | tonic (gRPC) |

### TypeScript HTTP 服务器

```typescript
// ==================== Express ====================
import express from 'express';
const app = express();
app.use(express.json());

app.get('/users', async (req, res) => {
  const users = await db.getUsers();
  res.json(users);
});

app.post('/users', async (req, res) => {
  const user = await db.createUser(req.body);
  res.status(201).json(user);
});

app.listen(3000);

// ==================== Fastify ====================
import Fastify from 'fastify';
const fastify = Fastify({ logger: true });

fastify.get('/users', async (request, reply) => {
  return await db.getUsers();
});

await fastify.listen({ port: 3000 });

// ==================== Hono (边缘计算友好) ====================
import { Hono } from 'hono';
const app = new Hono();

app.get('/users', (c) => c.json(users));
app.post('/users', async (c) => {
  const body = await c.req.json();
  return c.json(body, 201);
});

export default app;
```

### Python HTTP 服务器

```python
# ==================== FastAPI ====================
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel

app = FastAPI()

class User(BaseModel):
    name: str
    age: int

@app.get("/users")
async def get_users():
    return await db.get_users()

@app.post("/users", status_code=201)
async def create_user(user: User):
    return await db.create_user(user)

@app.get("/users/{user_id}")
async def get_user(user_id: int):
    user = await db.get_user(user_id)
    if not user:
        raise HTTPException(status_code=404, detail="Not found")
    return user

# 运行: uvicorn main:app --reload

# ==================== Flask ====================
from flask import Flask, request, jsonify

app = Flask(__name__)

@app.route('/users', methods=['GET'])
def get_users():
    return jsonify(db.get_users())

@app.route('/users', methods=['POST'])
def create_user():
    return jsonify(db.create_user(request.json)), 201
```

### Go HTTP 服务器

```go
// ==================== net/http ====================
func main() {
    http.HandleFunc("/users", handleUsers)
    http.ListenAndServe(":3000", nil)
}

func handleUsers(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case "GET":
        users := db.GetUsers()
        json.NewEncoder(w).Encode(users)
    case "POST":
        var user User
        json.NewDecoder(r.Body).Decode(&user)
        db.CreateUser(user)
        w.WriteHeader(http.StatusCreated)
    }
}

// ==================== Gin ====================
import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default()
    
    r.GET("/users", func(c *gin.Context) {
        c.JSON(200, db.GetUsers())
    })
    
    r.POST("/users", func(c *gin.Context) {
        var user User
        c.BindJSON(&user)
        db.CreateUser(user)
        c.JSON(201, user)
    })
    
    r.GET("/users/:id", func(c *gin.Context) {
        id := c.Param("id")
        c.JSON(200, db.GetUser(id))
    })
    
    r.Run(":3000")
}
```

### Rust HTTP 服务器

```rust
// ==================== Axum ====================
use axum::{routing::{get, post}, Router, Json, extract::Path};

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/users", get(get_users).post(create_user))
        .route("/users/:id", get(get_user));
    
    let listener = tokio::net::TcpListener::bind("0.0.0.0:3000").await.unwrap();
    axum::serve(listener, app).await.unwrap();
}

async fn get_users() -> Json<Vec<User>> {
    Json(db::get_users().await)
}

async fn create_user(Json(user): Json<CreateUser>) -> (StatusCode, Json<User>) {
    let user = db::create_user(user).await;
    (StatusCode::CREATED, Json(user))
}

async fn get_user(Path(id): Path<u32>) -> Result<Json<User>, StatusCode> {
    db::get_user(id).await
        .map(Json)
        .ok_or(StatusCode::NOT_FOUND)
}
```

### gRPC 服务器对比

```protobuf
// user.proto
syntax = "proto3";
package user;

service UserService {
  rpc GetUser(GetUserRequest) returns (User);
  rpc CreateUser(CreateUserRequest) returns (User);
}

message User {
  int32 id = 1;
  string name = 2;
}
```

**Go gRPC:**
```go
type server struct {
    pb.UnimplementedUserServiceServer
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.User, error) {
    return &pb.User{Id: req.Id, Name: "John"}, nil
}

func main() {
    lis, _ := net.Listen("tcp", ":50051")
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &server{})
    s.Serve(lis)
}
```

**Python gRPC:**
```python
class UserServicer(user_pb2_grpc.UserServiceServicer):
    def GetUser(self, request, context):
        return user_pb2.User(id=request.id, name="John")

server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
user_pb2_grpc.add_UserServiceServicer_to_server(UserServicer(), server)
server.add_insecure_port('[::]:50051')
server.start()
```

**Rust gRPC (tonic):**
```rust
#[tonic::async_trait]
impl UserService for MyUserService {
    async fn get_user(&self, request: Request<GetUserRequest>) 
        -> Result<Response<User>, Status> {
        Ok(Response::new(User {
            id: request.into_inner().id,
            name: "John".to_string(),
        }))
    }
}
```

---

## 🔢 数字类型

### 数字类型概览

| 语言 | 整数 | 浮点数 | 大整数 | 特点 |
|------|------|--------|--------|------|
| TypeScript | `number` | `number` | `bigint` | 统一 number |
| Python | `int` | `float` | `int` (无限) | 自动大整数 |
| Go | `int8`~`int64` | `float32/64` | `math/big` | 精确控制 |
| Rust | `i8`~`i128` | `f32/f64` | 外部 crate | 显式溢出 |

### TypeScript 数字

```typescript
// ==================== number (64位浮点) ====================
const int = 42;
const float = 3.14;
const negative = -100;
const scientific = 1e10;
const hex = 0xff;
const binary = 0b1010;
const octal = 0o755;

// 精度问题
0.1 + 0.2;           // 0.30000000000000004
0.1 + 0.2 === 0.3;   // false

// 特殊值
Infinity;            // 正无穷
-Infinity;           // 负无穷
NaN;                 // 非数字
Number.MAX_VALUE;    // 最大值 ~1.8e308
Number.MIN_VALUE;    // 最小正值 ~5e-324
Number.MAX_SAFE_INTEGER;  // 2^53 - 1

// 检查
Number.isNaN(NaN);        // true
Number.isFinite(100);     // true
Number.isInteger(42);     // true
Number.isSafeInteger(42); // true

// ==================== bigint ====================
const big = 9007199254740993n;
const bigHex = 0xffffffffffn;
BigInt("12345678901234567890");

// 运算 (不能与 number 混用)
big + 1n;            // OK
// big + 1;          // 错误!
big * 2n;
big ** 10n;

// ==================== 转换 ====================
parseInt("42");           // 42
parseInt("ff", 16);       // 255
parseFloat("3.14");       // 3.14
Number("42");             // 42
+"42";                    // 42

(42).toString();          // "42"
(255).toString(16);       // "ff"
(3.14159).toFixed(2);     // "3.14"
(1234.5).toExponential(); // "1.2345e+3"

// ==================== Math 对象 ====================
Math.abs(-5);        // 5
Math.round(3.5);     // 4
Math.floor(3.9);     // 3
Math.ceil(3.1);      // 4
Math.trunc(3.9);     // 3
Math.max(1, 2, 3);   // 3
Math.min(1, 2, 3);   // 1
Math.pow(2, 10);     // 1024
Math.sqrt(16);       // 4
Math.random();       // 0~1 随机数
```

### Python 数字

```python
# ==================== int (任意精度) ====================
x = 42
big = 123456789012345678901234567890  # 自动大整数
hex_num = 0xff
binary = 0b1010
octal = 0o755

# 无溢出
2 ** 1000  # 正常计算

# ==================== float (64位) ====================
pi = 3.14159
scientific = 1e10
inf = float('inf')
neg_inf = float('-inf')
nan = float('nan')

# 精度问题
0.1 + 0.2  # 0.30000000000000004

# ==================== decimal (精确小数) ====================
from decimal import Decimal, ROUND_HALF_UP

price = Decimal('19.99')
tax = Decimal('0.0825')
total = price * (1 + tax)
total.quantize(Decimal('0.01'), rounding=ROUND_HALF_UP)

# ==================== fractions (分数) ====================
from fractions import Fraction

f = Fraction(1, 3)
f + Fraction(1, 6)  # Fraction(1, 2)
Fraction('0.25')    # Fraction(1, 4)

# ==================== complex (复数) ====================
c = 3 + 4j
c.real    # 3.0
c.imag    # 4.0
abs(c)    # 5.0 (模)

# ==================== 转换 ====================
int("42")
int("ff", 16)       # 255
float("3.14")
str(42)
hex(255)            # "0xff"
bin(10)             # "0b1010"
oct(8)              # "0o10"

# ==================== 检查 ====================
import math

math.isnan(float('nan'))
math.isinf(float('inf'))
math.isfinite(42)
isinstance(42, int)

# ==================== math 模块 ====================
import math

math.floor(3.9)     # 3
math.ceil(3.1)      # 4
math.trunc(3.9)     # 3
math.sqrt(16)       # 4.0
math.pow(2, 10)     # 1024.0
math.log(100, 10)   # 2.0
math.sin(math.pi/2) # 1.0
math.factorial(5)   # 120
math.gcd(12, 18)    # 6
```

### Go 数字

```go
// ==================== 整数类型 ====================
var i8 int8    // -128 ~ 127
var i16 int16  // -32768 ~ 32767
var i32 int32  // -2^31 ~ 2^31-1
var i64 int64  // -2^63 ~ 2^63-1
var i int      // 32 或 64 位 (平台相关)

var u8 uint8   // 0 ~ 255 (byte)
var u16 uint16
var u32 uint32
var u64 uint64
var u uint

// 字面量
x := 42
hex := 0xff
binary := 0b1010
octal := 0o755
withSep := 1_000_000  // Go 1.13+

// ==================== 浮点类型 ====================
var f32 float32  // ~6位有效数字
var f64 float64  // ~15位有效数字

pi := 3.14159
scientific := 1e10

// 特殊值
inf := math.Inf(1)
negInf := math.Inf(-1)
nan := math.NaN()

// ==================== 转换 ====================
import "strconv"

strconv.Atoi("42")            // int, error
strconv.ParseInt("ff", 16, 64) // int64, error
strconv.ParseFloat("3.14", 64) // float64, error

strconv.Itoa(42)              // "42"
strconv.FormatInt(255, 16)    // "ff"
strconv.FormatFloat(3.14, 'f', 2, 64) // "3.14"

// 类型转换
float64(42)
int(3.14)  // 3 (截断)

// ==================== 大整数 ====================
import "math/big"

a := big.NewInt(1)
b := big.NewInt(1)
for i := 2; i <= 100; i++ {
    a.Mul(a, b.SetInt64(int64(i)))
}
// a = 100!

// ==================== math 包 ====================
import "math"

math.Abs(-5)
math.Floor(3.9)
math.Ceil(3.1)
math.Trunc(3.9)
math.Round(3.5)
math.Max(1, 2)
math.Min(1, 2)
math.Pow(2, 10)
math.Sqrt(16)
math.IsNaN(nan)
math.IsInf(inf, 1)

// ==================== 溢出处理 ====================
import "math/bits"

sum, carry := bits.Add64(a, b, 0)
diff, borrow := bits.Sub64(a, b, 0)
hi, lo := bits.Mul64(a, b)
```

### Rust 数字

```rust
// ==================== 整数类型 ====================
let i8: i8 = -128;      // -128 ~ 127
let i16: i16 = -32768;
let i32: i32 = -2_147_483_648;
let i64: i64 = 0;
let i128: i128 = 0;
let isize: isize = 0;   // 指针大小

let u8: u8 = 255;
let u16: u16 = 0;
let u32: u32 = 0;
let u64: u64 = 0;
let u128: u128 = 0;
let usize: usize = 0;

// 字面量
let decimal = 98_222;
let hex = 0xff;
let octal = 0o77;
let binary = 0b1111_0000;
let byte = b'A';  // u8

// 类型后缀
let x = 42i32;
let y = 3.14f64;

// ==================== 浮点类型 ====================
let f32: f32 = 3.14;  // 单精度
let f64: f64 = 3.14;  // 双精度 (默认)

let inf = f64::INFINITY;
let neg_inf = f64::NEG_INFINITY;
let nan = f64::NAN;

// ==================== 溢出处理 ====================
// Debug 模式: panic
// Release 模式: 环绕

// 显式处理
let (result, overflowed) = 255u8.overflowing_add(1);
let result = 255u8.wrapping_add(1);      // 环绕: 0
let result = 255u8.saturating_add(1);    // 饱和: 255
let result = 255u8.checked_add(1);       // Option: None

// ==================== 转换 ====================
// 解析
let n: i32 = "42".parse().unwrap();
let n = "42".parse::<i32>().unwrap();
let n = i32::from_str_radix("ff", 16).unwrap();

// 格式化
42.to_string();
format!("{:x}", 255);     // "ff"
format!("{:b}", 10);      // "1010"
format!("{:.2}", 3.14159); // "3.14"

// 类型转换
42i32 as f64;
3.14f64 as i32;  // 3 (截断)

// 安全转换
let n: u8 = 256i32.try_into().unwrap_or(255);

// ==================== 方法 ====================
(-5i32).abs();
3.14f64.floor();
3.14f64.ceil();
3.14f64.trunc();
3.14f64.round();
16f64.sqrt();
2f64.powi(10);  // 整数幂
2f64.powf(0.5); // 浮点幂
x.max(y);
x.min(y);
x.clamp(min, max);

f64::NAN.is_nan();
f64::INFINITY.is_infinite();
42f64.is_finite();

// ==================== num crate ====================
// 大整数、有理数等
use num_bigint::BigInt;
use num_rational::Ratio;

let big: BigInt = "123456789012345678901234567890".parse().unwrap();
let ratio = Ratio::new(1, 3);
```

### 数字类型对比

```
┌──────────────────┬────────────┬────────────┬────────────┬────────────┐
│ 特性             │ TypeScript │ Python     │ Go         │ Rust       │
├──────────────────┼────────────┼────────────┼────────────┼────────────┤
│ 默认整数         │ number     │ int        │ int        │ i32        │
│ 默认浮点         │ number     │ float      │ float64    │ f64        │
│ 整数精度         │ 53位       │ 无限       │ 64位       │ 128位      │
│ 大整数           │ bigint     │ 内置       │ math/big   │ num crate  │
│ 溢出处理         │ 无         │ 自动扩展   │ 环绕       │ 可选       │
│ 精确小数         │ 外部库     │ Decimal    │ 外部库     │ 外部库     │
│ 复数             │ ❌         │ 内置       │ complex128 │ num crate  │
└──────────────────┴────────────┴────────────┴────────────┴────────────┘
```

---

## 🔀 逻辑运算

### 布尔类型与运算符

| 语言 | 布尔类型 | 与 | 或 | 非 | 异或 |
|------|----------|-----|-----|-----|------|
| TypeScript | `boolean` | `&&` | `\|\|` | `!` | `^` (位) |
| Python | `bool` | `and` | `or` | `not` | `^` |
| Go | `bool` | `&&` | `\|\|` | `!` | `^` (位) |
| Rust | `bool` | `&&` | `\|\|` | `!` | `^` |

### TypeScript 逻辑运算

```typescript
// ==================== 布尔值 ====================
const t: boolean = true;
const f: boolean = false;

// ==================== 逻辑运算符 ====================
true && false;   // false (短路与)
true || false;   // true  (短路或)
!true;           // false

// ==================== 比较运算符 ====================
1 === 1;         // true  (严格相等)
1 == "1";        // true  (宽松相等，避免使用)
1 !== 2;         // true
1 < 2;           // true
1 <= 2;          // true
2 > 1;           // true
2 >= 1;          // true

// ==================== 短路求值 ====================
const name = user && user.name;
const value = input || "default";
const value2 = input ?? "default";  // 仅 null/undefined

// ==================== 条件表达式 ====================
const result = condition ? "yes" : "no";

// ==================== Falsy 值 ====================
// false, 0, -0, 0n, "", null, undefined, NaN
if (!0) console.log("0 is falsy");
if (!"") console.log('"" is falsy');

// ==================== 类型守卫 ====================
if (typeof x === "string") {
  x.toUpperCase();
}

if (x instanceof Date) {
  x.getTime();
}

if ("name" in obj) {
  obj.name;
}

if (Array.isArray(x)) {
  x.length;
}

// ==================== 位运算 ====================
0b1010 & 0b1100;  // 0b1000 (AND)
0b1010 | 0b1100;  // 0b1110 (OR)
0b1010 ^ 0b1100;  // 0b0110 (XOR)
~0b1010;          // 取反
0b1 << 4;         // 0b10000 (左移)
0b10000 >> 4;     // 0b1 (右移)
-1 >>> 1;         // 无符号右移
```

### Python 逻辑运算

```python
# ==================== 布尔值 ====================
t = True
f = False

# bool 是 int 子类
True + True   # 2
False * 10    # 0

# ==================== 逻辑运算符 ====================
True and False   # False
True or False    # True
not True         # False

# ==================== 比较运算符 ====================
1 == 1           # True
1 != 2           # True
1 < 2            # True
1 <= 2           # True
2 > 1            # True
2 >= 1           # True

# 链式比较
1 < x < 10       # 等价于 1 < x and x < 10
a == b == c      # 三者相等

# is vs ==
a = [1, 2]
b = [1, 2]
a == b           # True (值相等)
a is b           # False (不同对象)
a is not b       # True

# ==================== 短路求值 ====================
name = user and user.name
value = input or "default"

# 返回决定结果的值 (不一定是 bool)
1 and 2          # 2
0 and 2          # 0
1 or 2           # 1
0 or 2           # 2

# ==================== 条件表达式 ====================
result = "yes" if condition else "no"

# ==================== Falsy 值 ====================
# False, 0, 0.0, 0j, "", [], {}, set(), None
bool(0)          # False
bool([])         # False
bool("")         # False

# 自定义 Falsy
class MyClass:
    def __bool__(self):
        return False

# ==================== all / any ====================
all([True, True, False])   # False
any([True, False, False])  # True

all(x > 0 for x in [1, 2, 3])  # True
any(x < 0 for x in [1, 2, 3])  # False

# ==================== 位运算 ====================
0b1010 & 0b1100   # 0b1000
0b1010 | 0b1100   # 0b1110
0b1010 ^ 0b1100   # 0b0110
~0b1010           # -0b1011 (取反)
0b1 << 4          # 0b10000
0b10000 >> 4      # 0b1
```

### Go 逻辑运算

```go
// ==================== 布尔值 ====================
var t bool = true
var f bool = false

// ==================== 逻辑运算符 ====================
true && false    // false
true || false    // true
!true            // false

// ==================== 比较运算符 ====================
1 == 1           // true
1 != 2           // true
1 < 2            // true
1 <= 2           // true
2 > 1            // true
2 >= 1           // true

// 结构体比较 (可比较类型)
type Point struct { X, Y int }
p1 := Point{1, 2}
p2 := Point{1, 2}
p1 == p2         // true

// ==================== 短路求值 ====================
if user != nil && user.Name != "" {
    // 安全访问
}

// ==================== 条件语句 ====================
// Go 没有三元运算符
var result string
if condition {
    result = "yes"
} else {
    result = "no"
}

// if 初始化语句
if err := doSomething(); err != nil {
    return err
}

// ==================== switch ====================
switch value {
case 1:
    fmt.Println("one")
case 2, 3:
    fmt.Println("two or three")
default:
    fmt.Println("other")
}

// 无条件 switch (替代 if-else 链)
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}

// ==================== 位运算 ====================
0b1010 & 0b1100   // 0b1000
0b1010 | 0b1100   // 0b1110
0b1010 ^ 0b1100   // 0b0110
^0b1010           // 取反 (^x = -x-1)
0b1 << 4          // 0b10000
0b10000 >> 4      // 0b1
0b1010 &^ 0b1100  // 0b0010 (AND NOT)
```

### Rust 逻辑运算

```rust
// ==================== 布尔值 ====================
let t: bool = true;
let f: bool = false;

// ==================== 逻辑运算符 ====================
true && false    // false
true || false    // true
!true            // false

// ==================== 比较运算符 ====================
1 == 1           // true
1 != 2           // true
1 < 2            // true
1 <= 2           // true
2 > 1            // true
2 >= 1           // true

// 需要实现 PartialEq / Ord trait
#[derive(PartialEq, Eq, PartialOrd, Ord)]
struct Point { x: i32, y: i32 }

// ==================== 短路求值 ====================
let name = user.is_some() && user.unwrap().name.is_some();

// ==================== 条件表达式 ====================
// if 是表达式，可返回值
let result = if condition { "yes" } else { "no" };

let grade = if score >= 90 {
    'A'
} else if score >= 80 {
    'B'
} else {
    'C'
};

// ==================== match (模式匹配) ====================
let result = match value {
    1 => "one",
    2 | 3 => "two or three",
    4..=10 => "four to ten",
    n if n < 0 => "negative",
    _ => "other",
};

// 解构匹配
match point {
    Point { x: 0, y: 0 } => "origin",
    Point { x, y: 0 } => format!("on x-axis at {}", x),
    Point { x: 0, y } => format!("on y-axis at {}", y),
    Point { x, y } => format!("({}, {})", x, y),
}

// ==================== if let / while let ====================
if let Some(value) = option {
    println!("{}", value);
}

while let Some(item) = iter.next() {
    println!("{}", item);
}

// let else
let Some(value) = option else {
    return;
};

// ==================== 位运算 ====================
0b1010 & 0b1100   // 0b1000
0b1010 | 0b1100   // 0b1110
0b1010 ^ 0b1100   // 0b0110
!0b1010u8         // 0b11110101 (取反)
0b1 << 4          // 0b10000
0b10000 >> 4      // 0b1
```

---

## 📤 包发布

### 发布流程概览

| 语言 | 仓库 | 配置文件 | 发布命令 |
|------|------|----------|----------|
| TypeScript | npmjs.com | package.json | `npm publish` |
| Python | pypi.org | pyproject.toml | `poetry publish` |
| Go | pkg.go.dev | go.mod | git tag |
| Rust | crates.io | Cargo.toml | `cargo publish` |

### TypeScript 包发布

```json
// package.json
{
  "name": "@myorg/mypackage",
  "version": "1.0.0",
  "description": "My awesome package",
  "main": "dist/index.js",
  "module": "dist/index.mjs",
  "types": "dist/index.d.ts",
  "exports": {
    ".": {
      "import": "./dist/index.mjs",
      "require": "./dist/index.js",
      "types": "./dist/index.d.ts"
    }
  },
  "files": ["dist"],
  "scripts": {
    "build": "tsup src/index.ts --format cjs,esm --dts",
    "prepublishOnly": "npm run build"
  },
  "keywords": ["typescript", "utility"],
  "author": "Your Name <email@example.com>",
  "license": "MIT",
  "repository": {
    "type": "git",
    "url": "https://github.com/user/repo"
  },
  "publishConfig": {
    "access": "public"
  }
}
```

```bash
# 登录
npm login

# 发布
npm publish                    # 公开包
npm publish --access public    # scope 包

# 版本管理
npm version patch              # 1.0.0 -> 1.0.1
npm version minor              # 1.0.0 -> 1.1.0
npm version major              # 1.0.0 -> 2.0.0

# 预发布
npm publish --tag beta
npm version prerelease --preid=beta  # 1.0.0 -> 1.0.1-beta.0

# 废弃版本
npm deprecate mypackage@1.0.0 "Use v2 instead"

# 撤销 (72小时内)
npm unpublish mypackage@1.0.0
```

### Python 包发布

```toml
# pyproject.toml
[project]
name = "mypackage"
version = "1.0.0"
description = "My awesome package"
readme = "README.md"
license = {text = "MIT"}
authors = [{name = "Your Name", email = "email@example.com"}]
keywords = ["python", "utility"]
classifiers = [
    "Development Status :: 4 - Beta",
    "Intended Audience :: Developers",
    "License :: OSI Approved :: MIT License",
    "Programming Language :: Python :: 3.10",
]
requires-python = ">=3.10"
dependencies = ["requests>=2.28.0"]

[project.optional-dependencies]
dev = ["pytest>=7.0.0", "mypy>=1.0.0"]

[project.urls]
Homepage = "https://github.com/user/repo"
Documentation = "https://mypackage.readthedocs.io"

[project.scripts]
mycli = "mypackage.cli:main"

[build-system]
requires = ["hatchling"]
build-backend = "hatchling.build"
```

```bash
# Poetry 发布
poetry config pypi-token.pypi your-api-token
poetry build
poetry publish

# 或使用 twine
pip install build twine
python -m build
twine upload dist/*

# 测试发布 (TestPyPI)
twine upload --repository testpypi dist/*

# 版本管理
poetry version patch           # 1.0.0 -> 1.0.1
poetry version minor           # 1.0.0 -> 1.1.0
poetry version major           # 1.0.0 -> 2.0.0
```

### Go 包发布

```go
// go.mod
module github.com/user/mypackage

go 1.21

// 无需 require 用于库
```

```bash
# Go 使用 Git 标签发布，无需上传到中央仓库

# 1. 确保代码可导入
# github.com/user/mypackage 必须是有效的导入路径

# 2. 版本标签
git tag v1.0.0
git push origin v1.0.0

# 语义化版本
git tag v1.0.1    # 补丁
git tag v1.1.0    # 次版本
git tag v2.0.0    # 主版本 (需要更新 module 路径)

# 主版本 > 1 需要更新 go.mod
module github.com/user/mypackage/v2

# 3. pkg.go.dev 自动索引
# 访问 https://pkg.go.dev/github.com/user/mypackage

# 预发布
git tag v1.0.0-beta.1

# 撤回版本 (在 go.mod 中)
retract v1.0.0  // 有严重 bug

# 私有模块
go env -w GOPRIVATE=github.com/myorg/*
```

### Rust 包发布

```toml
# Cargo.toml
[package]
name = "mypackage"
version = "1.0.0"
edition = "2021"
authors = ["Your Name <email@example.com>"]
description = "My awesome package"
documentation = "https://docs.rs/mypackage"
readme = "README.md"
homepage = "https://github.com/user/repo"
repository = "https://github.com/user/repo"
license = "MIT"
keywords = ["rust", "utility"]
categories = ["development-tools"]
exclude = ["tests/*", "benches/*"]

[dependencies]
serde = { version = "1.0", features = ["derive"] }

[dev-dependencies]
tokio-test = "0.4"

[features]
default = ["std"]
std = []
full = ["std", "extra"]
```

```bash
# 登录
cargo login your-api-token

# 检查
cargo publish --dry-run

# 发布
cargo publish

# 版本管理 (手动修改 Cargo.toml 或使用 cargo-release)
cargo install cargo-release
cargo release patch            # 1.0.0 -> 1.0.1
cargo release minor            # 1.0.0 -> 1.1.0
cargo release major            # 1.0.0 -> 2.0.0

# 撤回版本 (yank)
cargo yank --version 1.0.0
cargo yank --version 1.0.0 --undo

# 发布工作空间中的包
cargo publish -p mypackage
```

### 发布检查清单

```
┌─────────────────────────────────────────────────────────────────┐
│                        发布前检查清单                            │
├─────────────────────────────────────────────────────────────────┤
│ □ README.md 完整清晰                                            │
│ □ CHANGELOG.md 更新                                             │
│ □ LICENSE 文件存在                                              │
│ □ 版本号符合语义化版本                                          │
│ □ 依赖版本固定或有合理范围                                      │
│ □ 测试全部通过                                                  │
│ □ 文档完整 (API 文档、示例)                                     │
│ □ 无敏感信息 (token、密钥)                                      │
│ □ .gitignore / .npmignore 配置正确                             │
│ □ CI/CD 检查通过                                                │
└─────────────────────────────────────────────────────────────────┘
```

### 版本号规范 (SemVer)

```
MAJOR.MINOR.PATCH[-PRERELEASE][+BUILD]

1.0.0        稳定版本
1.0.1        补丁版本 (bug 修复)
1.1.0        次版本 (新功能，向后兼容)
2.0.0        主版本 (破坏性变更)
1.0.0-alpha  预发布版本
1.0.0-beta.1 预发布迭代
1.0.0+build  构建元数据

版本范围:
^1.2.3    >=1.2.3 <2.0.0   (兼容)
~1.2.3    >=1.2.3 <1.3.0   (补丁)
>=1.2.3   >=1.2.3          (最小版本)
1.2.*     >=1.2.0 <1.3.0   (通配符)
```

---

## 🧠 内存管理

### 内存管理模型概览

| 语言 | 管理方式 | GC 类型 | 特点 |
|------|----------|---------|------|
| TypeScript | 自动 GC | 标记-清除 (V8) | 完全自动，开发者无感 |
| Python | 自动 GC | 引用计数 + 分代 GC | 引用计数为主 |
| Go | 自动 GC | 三色标记并发 GC | 低延迟，可调优 |
| Rust | 所有权系统 | 无 GC | 编译时内存安全 |

### TypeScript/JavaScript 内存管理

```typescript
// ==================== 内存分配 ====================
// 自动分配，无需手动管理
const obj = { name: "John" };     // 堆分配
const arr = [1, 2, 3];            // 堆分配
const str = "hello";              // 字符串池/堆

// ==================== 垃圾回收 (V8) ====================
/*
V8 使用分代 GC:
- 新生代 (Young Generation): 小对象，频繁 GC
- 老生代 (Old Generation): 存活久的对象，较少 GC

GC 算法:
- Scavenge: 新生代，复制算法
- Mark-Sweep: 老生代，标记-清除
- Mark-Compact: 老生代，标记-整理
*/

// ==================== 内存泄漏常见原因 ====================

// 1. 全局变量
function leak() {
  globalVar = "leaked";  // 忘记 let/const，变成全局变量
}

// 2. 闭包引用
function createClosure() {
  const largeData = new Array(1000000);
  return function() {
    console.log(largeData.length);  // largeData 无法被回收
  };
}

// 3. 事件监听器未移除
element.addEventListener('click', handler);
// 忘记: element.removeEventListener('click', handler);

// 4. 定时器未清除
const timer = setInterval(() => {}, 1000);
// 忘记: clearInterval(timer);

// 5. DOM 引用
const elements = {
  button: document.getElementById('button')
};
document.body.removeChild(elements.button);
// elements.button 仍持有引用

// ==================== 最佳实践 ====================

// WeakMap/WeakSet (弱引用，不阻止 GC)
const cache = new WeakMap();
cache.set(obj, computedValue);
// obj 被回收时，缓存自动清理

// 手动解除引用
let data = fetchLargeData();
processData(data);
data = null;  // 帮助 GC

// 使用对象池
class ObjectPool<T> {
  private pool: T[] = [];
  
  acquire(factory: () => T): T {
    return this.pool.pop() ?? factory();
  }
  
  release(obj: T): void {
    this.pool.push(obj);
  }
}

// ==================== 内存分析 ====================
// Chrome DevTools -> Memory
// - Heap Snapshot: 堆快照
// - Allocation Timeline: 分配时间线
// - Allocation Sampling: 分配采样

// Node.js
// node --inspect app.js
// process.memoryUsage()
```

### Python 内存管理

```python
import sys
import gc

# ==================== 引用计数 ====================
a = [1, 2, 3]
sys.getrefcount(a)  # 引用计数 (会多 1，因为参数传递)

b = a               # 引用计数 +1
del b               # 引用计数 -1
a = None            # 引用计数 -1，可能被回收

# ==================== 循环引用 ====================
class Node:
    def __init__(self):
        self.ref = None

a = Node()
b = Node()
a.ref = b
b.ref = a  # 循环引用！引用计数永不为 0

# 分代 GC 处理循环引用
gc.collect()  # 手动触发 GC

# ==================== 垃圾回收控制 ====================
gc.disable()           # 禁用 GC
gc.enable()            # 启用 GC
gc.collect()           # 手动 GC
gc.get_count()         # 各代计数
gc.get_threshold()     # GC 阈值

# 设置阈值 (generation 0, 1, 2)
gc.set_threshold(700, 10, 10)

# 获取无法回收的对象
gc.garbage

# ==================== 弱引用 ====================
import weakref

class Data:
    pass

data = Data()
weak_ref = weakref.ref(data)

weak_ref()  # 返回对象或 None
del data
weak_ref()  # None

# WeakValueDictionary
cache = weakref.WeakValueDictionary()
cache['key'] = Data()  # 对象可被 GC

# ==================== 内存优化 ====================

# __slots__ 减少内存
class Point:
    __slots__ = ['x', 'y']  # 不使用 __dict__
    def __init__(self, x, y):
        self.x = x
        self.y = y

# 生成器代替列表
def numbers():
    for i in range(1000000):
        yield i
# 而不是: [i for i in range(1000000)]

# array 代替 list (数值)
from array import array
arr = array('i', [1, 2, 3])  # 更紧凑

# ==================== 内存分析 ====================
import tracemalloc

tracemalloc.start()
# ... 代码 ...
snapshot = tracemalloc.take_snapshot()
top_stats = snapshot.statistics('lineno')

for stat in top_stats[:10]:
    print(stat)

# 内存使用
import resource
resource.getrusage(resource.RUSAGE_SELF).ru_maxrss

# 对象大小
sys.getsizeof(obj)

# memory_profiler
# pip install memory-profiler
# @profile
# def my_func():
#     ...
# python -m memory_profiler script.py
```

### Go 内存管理

```go
import (
    "runtime"
    "runtime/debug"
)

// ==================== 内存分配 ====================

// 栈分配 (小对象，编译器自动决定)
func stackAlloc() {
    x := 42           // 栈
    arr := [10]int{}  // 小数组，栈
}

// 堆分配 (逃逸分析决定)
func heapAlloc() *int {
    x := 42
    return &x  // x 逃逸到堆
}

// new 和 make
ptr := new(int)           // 分配并返回指针
slice := make([]int, 10)  // 分配切片
m := make(map[string]int) // 分配 map
ch := make(chan int, 10)  // 分配 channel

// ==================== 垃圾回收 ====================
/*
Go GC 特点:
- 三色标记并发 GC
- 写屏障 (Write Barrier)
- 目标: 低延迟 (< 1ms STW)
*/

// 手动触发 GC
runtime.GC()

// GC 统计
var stats runtime.MemStats
runtime.ReadMemStats(&stats)
stats.Alloc      // 当前堆分配
stats.TotalAlloc // 累计分配
stats.Sys        // 系统内存
stats.NumGC      // GC 次数

// ==================== GC 调优 ====================

// GOGC 环境变量 (默认 100)
// GOGC=200  内存翻倍时触发 GC
// GOGC=50   内存增加 50% 触发 GC
// GOGC=off  禁用 GC

// 程序中设置
debug.SetGCPercent(200)

// 内存限制 (Go 1.19+)
debug.SetMemoryLimit(1 << 30)  // 1GB

// ==================== 内存池 ====================
import "sync"

var bufferPool = sync.Pool{
    New: func() interface{} {
        return make([]byte, 1024)
    },
}

func process() {
    buf := bufferPool.Get().([]byte)
    defer bufferPool.Put(buf)
    // 使用 buf
}

// ==================== 逃逸分析 ====================
// go build -gcflags="-m" main.go

// 不逃逸
func noEscape() int {
    x := 42
    return x
}

// 逃逸到堆
func escape() *int {
    x := 42
    return &x  // "moved to heap: x"
}

// 接口导致逃逸
func interfaceEscape() {
    x := 42
    fmt.Println(x)  // x 逃逸 (interface{} 参数)
}

// ==================== 最佳实践 ====================

// 预分配切片
slice := make([]int, 0, expectedSize)

// 复用对象
// 使用 sync.Pool

// 避免不必要的指针
type Good struct {
    x, y int  // 值类型
}

type Bad struct {
    x, y *int  // 指针增加 GC 压力
}

// ==================== 内存分析 ====================
import _ "net/http/pprof"

// go tool pprof http://localhost:6060/debug/pprof/heap
// go tool pprof http://localhost:6060/debug/pprof/allocs

// 生成 profile
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
// go tool pprof mem.prof
```

### Rust 内存管理

```rust
// ==================== 所有权系统 ====================
fn main() {
    // 每个值有唯一所有者
    let s1 = String::from("hello");
    
    // 移动 (Move): 所有权转移
    let s2 = s1;
    // println!("{}", s1);  // 错误！s1 已无效
    
    // 克隆 (Clone): 深拷贝
    let s3 = s2.clone();
    println!("{} {}", s2, s3);  // OK
    
    // Copy trait: 栈上数据自动复制
    let x = 5;
    let y = x;
    println!("{} {}", x, y);  // OK，整数实现了 Copy
}

// ==================== 借用 (Borrowing) ====================

// 不可变借用 (&T)
fn print_len(s: &String) {
    println!("{}", s.len());
}

// 可变借用 (&mut T)
fn push_str(s: &mut String) {
    s.push_str(" world");
}

// 借用规则:
// 1. 任意数量的不可变借用，或
// 2. 一个可变借用
// 不能同时存在

let mut s = String::from("hello");
let r1 = &s;
let r2 = &s;      // OK: 多个不可变借用
// let r3 = &mut s;  // 错误！已有不可变借用

// ==================== 生命周期 ====================

// 显式生命周期标注
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

// 结构体中的引用
struct Excerpt<'a> {
    part: &'a str,
}

// 静态生命周期
let s: &'static str = "hello";  // 整个程序期间有效

// ==================== 智能指针 ====================

// Box<T> - 堆分配
let b = Box::new(5);

// Rc<T> - 引用计数 (单线程)
use std::rc::Rc;
let a = Rc::new(5);
let b = Rc::clone(&a);  // 引用计数 +1
Rc::strong_count(&a);   // 2

// Arc<T> - 原子引用计数 (多线程)
use std::sync::Arc;
let a = Arc::new(5);
let b = Arc::clone(&a);

// RefCell<T> - 内部可变性 (运行时借用检查)
use std::cell::RefCell;
let data = RefCell::new(5);
*data.borrow_mut() += 1;

// Mutex<T> - 线程安全的内部可变性
use std::sync::Mutex;
let m = Mutex::new(5);
let mut num = m.lock().unwrap();
*num = 6;

// ==================== Drop trait ====================

struct CustomDrop {
    data: String,
}

impl Drop for CustomDrop {
    fn drop(&mut self) {
        println!("Dropping: {}", self.data);
    }
}

// 手动释放
let c = CustomDrop { data: String::from("test") };
drop(c);  // 立即调用 drop

// ==================== 内存布局 ====================

use std::mem;

mem::size_of::<i32>();      // 4
mem::size_of::<String>();   // 24 (ptr + len + cap)
mem::align_of::<i32>();     // 4

// 零大小类型
struct Empty;
mem::size_of::<Empty>();    // 0

// ==================== unsafe 内存操作 ====================

// 裸指针
let x = 5;
let raw_ptr = &x as *const i32;

unsafe {
    println!("{}", *raw_ptr);
}

// 手动内存管理
use std::alloc::{alloc, dealloc, Layout};

unsafe {
    let layout = Layout::new::<i32>();
    let ptr = alloc(layout) as *mut i32;
    *ptr = 42;
    dealloc(ptr as *mut u8, layout);
}

// ==================== 内存泄漏 ====================

// 故意泄漏 (有时有用)
let leaked: &'static str = Box::leak(String::from("hello").into_boxed_str());

// Rc 循环引用导致泄漏
use std::rc::{Rc, Weak};
use std::cell::RefCell;

struct Node {
    next: Option<Rc<RefCell<Node>>>,
    prev: Option<Weak<RefCell<Node>>>,  // 用 Weak 打破循环
}

// ==================== 内存分析 ====================

// Valgrind
// valgrind --tool=memcheck ./target/debug/myapp

// Heaptrack
// heaptrack ./target/release/myapp
// heaptrack_gui heaptrack.myapp.*.gz

// 自定义分配器
#[global_allocator]
static GLOBAL: jemallocator::Jemalloc = jemallocator::Jemalloc;
```

### 内存管理对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 管理方式        │ 自动 GC        │ 自动 GC        │ 自动 GC        │ 所有权系统     │
│ GC 算法         │ 分代标记清除   │ 引用计数+分代  │ 三色并发标记   │ 无             │
│ GC 暂停         │ 有             │ 有             │ <1ms           │ 无             │
│ 循环引用        │ 自动处理       │ 分代 GC 处理   │ 自动处理       │ 编译时防止     │
│ 手动控制        │ ❌             │ gc 模块        │ runtime 包     │ 完全控制       │
│ 内存池          │ 外部实现       │ 外部实现       │ sync.Pool      │ 自定义分配器   │
│ 弱引用          │ WeakMap/Set    │ weakref        │ ❌             │ Weak<T>        │
│ 逃逸分析        │ V8 内部        │ ❌             │ ✅             │ 编译时确定     │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 内存优化建议

| 场景 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 对象复用 | 对象池 | `__slots__` | `sync.Pool` | 自定义 |
| 减少分配 | 预分配数组 | 生成器 | 预分配切片 | 栈分配 |
| 大数据 | 流处理 | 迭代器 | 流式处理 | 零拷贝 |
| 缓存 | WeakMap | WeakValueDict | 带 TTL | `Weak<T>` |
| 分析工具 | Chrome DevTools | tracemalloc | pprof | Valgrind |

---

## 🔄 并发模型

### 并发模型概览

| 语言 | 并发模型 | 并行能力 | 线程安全 | 适用场景 |
|------|----------|----------|----------|----------|
| TypeScript | 事件循环 | Worker | 隔离内存 | I/O 密集 |
| Python | GIL + 多进程 | 多进程 | GIL 保护 | I/O/多进程 |
| Go | CSP (Goroutine) | 原生并行 | Channel | 通用 |
| Rust | 线程 + async | 原生并行 | 类型系统 | 系统级 |

### 并发 vs 并行

```
并发 (Concurrency): 同时处理多个任务（交替执行）
并行 (Parallelism): 同时执行多个任务（真正同时）

┌─────────────────────────────────────────────────────────┐
│ 并发 (单核)                                             │
│ Task A: ████░░░░████░░░░████                            │
│ Task B: ░░░░████░░░░████░░░░████                        │
├─────────────────────────────────────────────────────────┤
│ 并行 (多核)                                             │
│ Core 1: ████████████████                               │
│ Core 2: ████████████████                               │
└─────────────────────────────────────────────────────────┘
```

### TypeScript 并发模型

```typescript
// ==================== 事件循环 (Event Loop) ====================
/*
┌───────────────────────────────────────┐
│           Call Stack                  │
├───────────────────────────────────────┤
│                                       │
│  ┌─────────────────────────────────┐  │
│  │         Event Loop              │  │
│  │  ┌─────────────────────────┐    │  │
│  │  │ 1. Microtask Queue      │    │  │
│  │  │    (Promise, queueMicro)│    │  │
│  │  ├─────────────────────────┤    │  │
│  │  │ 2. Macrotask Queue      │    │  │
│  │  │    (setTimeout, I/O)    │    │  │
│  │  └─────────────────────────┘    │  │
│  └─────────────────────────────────┘  │
│                                       │
└───────────────────────────────────────┘

执行顺序:
1. 执行同步代码
2. 清空微任务队列
3. 执行一个宏任务
4. 重复 2-3
*/

console.log('1');
setTimeout(() => console.log('2'), 0);
Promise.resolve().then(() => console.log('3'));
console.log('4');
// 输出: 1, 4, 3, 2

// ==================== 单线程非阻塞 ====================
// 所有 I/O 操作都是非阻塞的
const response = await fetch(url);  // 不阻塞主线程
const data = await response.json();

// ==================== Web Workers (浏览器) ====================
// main.js
const worker = new Worker('worker.js');

worker.postMessage({ type: 'compute', data: [1, 2, 3] });

worker.onmessage = (e) => {
  console.log('Result:', e.data);
};

worker.onerror = (e) => {
  console.error('Worker error:', e);
};

// worker.js
self.onmessage = (e) => {
  const result = heavyComputation(e.data);
  self.postMessage(result);
};

// ==================== Worker Threads (Node.js) ====================
import { Worker, isMainThread, parentPort, workerData } from 'worker_threads';

if (isMainThread) {
  const worker = new Worker(__filename, {
    workerData: { nums: [1, 2, 3, 4, 5] }
  });
  
  worker.on('message', (result) => {
    console.log('Sum:', result);
  });
} else {
  const sum = workerData.nums.reduce((a, b) => a + b, 0);
  parentPort.postMessage(sum);
}

// ==================== SharedArrayBuffer (共享内存) ====================
const sab = new SharedArrayBuffer(1024);
const arr = new Int32Array(sab);

// Atomics 原子操作
Atomics.add(arr, 0, 5);
Atomics.load(arr, 0);
Atomics.store(arr, 0, 10);
Atomics.compareExchange(arr, 0, 10, 20);

// 等待/通知
Atomics.wait(arr, 0, 0);   // 等待值变化
Atomics.notify(arr, 0, 1); // 唤醒等待者
```

### Python 并发模型

```python
import threading
import multiprocessing
import asyncio
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# ==================== GIL (Global Interpreter Lock) ====================
"""
GIL 特点:
- 同一时刻只有一个线程执行 Python 字节码
- I/O 操作会释放 GIL
- CPU 密集型任务无法利用多核

解决方案:
- I/O 密集型: 多线程/asyncio
- CPU 密集型: 多进程
"""

# ==================== 多线程 (I/O 密集型) ====================
def fetch_url(url):
    response = requests.get(url)
    return response.text

# 方式 1: Thread
threads = []
for url in urls:
    t = threading.Thread(target=fetch_url, args=(url,))
    t.start()
    threads.append(t)

for t in threads:
    t.join()

# 方式 2: ThreadPoolExecutor
with ThreadPoolExecutor(max_workers=10) as executor:
    results = list(executor.map(fetch_url, urls))
    
    # 或使用 submit
    futures = [executor.submit(fetch_url, url) for url in urls]
    results = [f.result() for f in futures]

# ==================== 多进程 (CPU 密集型) ====================
def cpu_bound(n):
    return sum(i * i for i in range(n))

# 方式 1: Process
processes = []
for i in range(4):
    p = multiprocessing.Process(target=cpu_bound, args=(10**7,))
    p.start()
    processes.append(p)

for p in processes:
    p.join()

# 方式 2: ProcessPoolExecutor
with ProcessPoolExecutor(max_workers=4) as executor:
    results = list(executor.map(cpu_bound, [10**7] * 4))

# 方式 3: Pool
with multiprocessing.Pool(4) as pool:
    results = pool.map(cpu_bound, [10**7] * 4)

# ==================== 进程间通信 ====================
# Queue
queue = multiprocessing.Queue()
queue.put(item)
item = queue.get()

# Pipe
parent_conn, child_conn = multiprocessing.Pipe()
parent_conn.send(data)
data = child_conn.recv()

# 共享内存
shared_value = multiprocessing.Value('i', 0)
shared_array = multiprocessing.Array('d', [0.0] * 10)

with shared_value.get_lock():
    shared_value.value += 1

# Manager (更灵活但更慢)
manager = multiprocessing.Manager()
shared_dict = manager.dict()
shared_list = manager.list()

# ==================== 线程同步 ====================
lock = threading.Lock()
rlock = threading.RLock()        # 可重入锁
semaphore = threading.Semaphore(5)
event = threading.Event()
condition = threading.Condition()
barrier = threading.Barrier(3)

# 使用锁
with lock:
    # 临界区
    shared_resource += 1

# ==================== asyncio (协程) ====================
async def fetch(url):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            return await response.text()

async def main():
    tasks = [fetch(url) for url in urls]
    results = await asyncio.gather(*tasks)

asyncio.run(main())
```

### Go 并发模型

```go
// ==================== CSP 模型 ====================
/*
CSP (Communicating Sequential Processes):
"不要通过共享内存来通信，而要通过通信来共享内存"

Goroutine: 轻量级线程 (~2KB 栈)
Channel: Goroutine 间通信的管道
*/

// ==================== Goroutine ====================
func main() {
    // 启动 goroutine
    go func() {
        fmt.Println("Hello from goroutine")
    }()
    
    // 启动多个
    for i := 0; i < 100; i++ {
        go worker(i)
    }
    
    time.Sleep(time.Second)  // 等待 (实际应用用 WaitGroup)
}

// ==================== Channel ====================
// 无缓冲 channel (同步)
ch := make(chan int)

go func() {
    ch <- 42  // 阻塞直到被接收
}()

value := <-ch  // 阻塞直到有值

// 有缓冲 channel (异步)
ch := make(chan int, 100)
ch <- 1  // 不阻塞 (直到满)

// 关闭 channel
close(ch)

// 遍历 channel
for v := range ch {
    fmt.Println(v)
}

// 检查是否关闭
v, ok := <-ch
if !ok {
    // channel 已关闭
}

// ==================== Select ====================
select {
case v := <-ch1:
    fmt.Println("from ch1:", v)
case v := <-ch2:
    fmt.Println("from ch2:", v)
case ch3 <- value:
    fmt.Println("sent to ch3")
case <-time.After(time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("no communication")
}

// ==================== 并发模式 ====================

// Fan-out: 一个输入，多个处理者
func fanOut(in <-chan int, workers int) []<-chan int {
    outs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        outs[i] = worker(in)
    }
    return outs
}

// Fan-in: 多个输入，合并为一个
func fanIn(channels ...<-chan int) <-chan int {
    out := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range channels {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for v := range c {
                out <- v
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(out)
    }()
    
    return out
}

// Pipeline: 链式处理
func pipeline() {
    nums := gen(1, 2, 3, 4, 5)
    squared := square(nums)
    for v := range squared {
        fmt.Println(v)
    }
}

func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for _, n := range nums {
            out <- n
        }
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        defer close(out)
        for n := range in {
            out <- n * n
        }
    }()
    return out
}

// Worker Pool
func workerPool(jobs <-chan Job, results chan<- Result, workers int) {
    var wg sync.WaitGroup
    for i := 0; i < workers; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            for job := range jobs {
                results <- process(job)
            }
        }()
    }
    wg.Wait()
    close(results)
}

// ==================== 同步原语 ====================
var (
    mu      sync.Mutex
    rwMu    sync.RWMutex
    once    sync.Once
    wg      sync.WaitGroup
    cond    = sync.NewCond(&sync.Mutex{})
    pool    = sync.Pool{New: func() interface{} { return new(Buffer) }}
)

// Mutex
mu.Lock()
defer mu.Unlock()
// 临界区

// RWMutex
rwMu.RLock()   // 读锁
rwMu.RUnlock()
rwMu.Lock()    // 写锁
rwMu.Unlock()

// WaitGroup
wg.Add(1)
go func() {
    defer wg.Done()
    // 工作
}()
wg.Wait()

// Once
once.Do(func() {
    // 只执行一次
})

// ==================== 原子操作 ====================
import "sync/atomic"

var counter int64

atomic.AddInt64(&counter, 1)
atomic.LoadInt64(&counter)
atomic.StoreInt64(&counter, 0)
atomic.CompareAndSwapInt64(&counter, old, new)

// atomic.Value
var config atomic.Value
config.Store(newConfig)
c := config.Load().(Config)
```

### Rust 并发模型

```rust
use std::thread;
use std::sync::{Arc, Mutex, RwLock, mpsc, Barrier};
use std::sync::atomic::{AtomicUsize, Ordering};

// ==================== 线程安全保证 ====================
/*
Rust 编译时保证线程安全:
- Send: 可以安全地在线程间转移所有权
- Sync: 可以安全地在线程间共享引用

大多数类型自动实现 Send + Sync
Rc<T> 不是 Send (用 Arc<T> 代替)
RefCell<T> 不是 Sync (用 Mutex<T> 代替)
*/

// ==================== 线程 ====================
// 创建线程
let handle = thread::spawn(|| {
    println!("Hello from thread");
    42
});

let result = handle.join().unwrap();

// 移动所有权到线程
let data = vec![1, 2, 3];
let handle = thread::spawn(move || {
    println!("{:?}", data);
});

// 作用域线程 (可借用数据)
let data = vec![1, 2, 3];
thread::scope(|s| {
    s.spawn(|| {
        println!("{:?}", data);  // 借用 data
    });
    s.spawn(|| {
        println!("len: {}", data.len());
    });
});
// 作用域结束时自动 join

// ==================== 共享状态 ====================
// Arc + Mutex
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

// RwLock (多读单写)
let data = Arc::new(RwLock::new(vec![1, 2, 3]));

// 读
let read_guard = data.read().unwrap();

// 写
let mut write_guard = data.write().unwrap();
write_guard.push(4);

// ==================== Channel ====================
// mpsc: 多生产者单消费者
let (tx, rx) = mpsc::channel();

thread::spawn(move || {
    tx.send(42).unwrap();
});

let received = rx.recv().unwrap();

// 多生产者
let tx2 = tx.clone();

// 同步 channel (有界)
let (tx, rx) = mpsc::sync_channel(10);

// 迭代接收
for received in rx {
    println!("{}", received);
}

// ==================== 原子操作 ====================
let counter = AtomicUsize::new(0);

counter.fetch_add(1, Ordering::SeqCst);
counter.load(Ordering::SeqCst);
counter.store(0, Ordering::SeqCst);
counter.compare_exchange(0, 1, Ordering::SeqCst, Ordering::SeqCst);

// Ordering:
// - Relaxed: 最弱，只保证原子性
// - Acquire/Release: 同步点
// - SeqCst: 最强，顺序一致

// ==================== async/await ====================
use tokio;

#[tokio::main]
async fn main() {
    let handle = tokio::spawn(async {
        // 异步任务
        42
    });
    
    let result = handle.await.unwrap();
}

// 并发执行
async fn concurrent() {
    let (a, b) = tokio::join!(
        async_task_1(),
        async_task_2()
    );
}

// 选择第一个完成的
async fn select_first() {
    tokio::select! {
        v = async_task_1() => println!("task 1: {}", v),
        v = async_task_2() => println!("task 2: {}", v),
    }
}

// ==================== 并行迭代 (rayon) ====================
use rayon::prelude::*;

// 并行 map
let results: Vec<_> = data.par_iter()
    .map(|x| expensive_computation(x))
    .collect();

// 并行 filter
let filtered: Vec<_> = data.par_iter()
    .filter(|x| predicate(x))
    .collect();

// 并行 reduce
let sum: i32 = data.par_iter().sum();

// 并行排序
data.par_sort();

// ==================== crossbeam (高级并发) ====================
use crossbeam::channel;
use crossbeam::scope;

// 多生产者多消费者 channel
let (s, r) = channel::unbounded();
let (s, r) = channel::bounded(10);

// select
crossbeam::select! {
    recv(r1) -> msg => println!("r1: {:?}", msg),
    recv(r2) -> msg => println!("r2: {:?}", msg),
    send(s, value) -> res => println!("sent"),
    default => println!("no operation"),
}
```

### 并发模型对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 并发模型        │ 事件循环       │ GIL + 多进程   │ CSP            │ 线程 + async   │
│ 并发单元        │ Promise        │ Thread/Process │ Goroutine      │ Thread/Task    │
│ 通信方式        │ 消息传递       │ Queue/Pipe     │ Channel        │ Channel        │
│ 共享状态        │ SharedArray    │ multiprocessing│ Mutex/Channel  │ Arc<Mutex>     │
│ CPU 并行        │ Worker         │ 多进程         │ 原生           │ 原生           │
│ 内存开销        │ ~1MB/Worker    │ ~10MB/Process  │ ~2KB/Goroutine │ ~8KB/Thread    │
│ 数据竞争        │ 隔离           │ GIL 保护       │ 运行时检测     │ 编译时防止     │
│ 死锁检测        │ ❌             │ ❌             │ ❌             │ 部分编译时     │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 何时使用何种并发

```
┌─────────────────────────────────────────────────────────────────────────┐
│                           选择指南                                       │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  I/O 密集型 (网络/磁盘)                                                 │
│  ├─ TypeScript: async/await (首选)                                      │
│  ├─ Python: asyncio 或 多线程                                          │
│  ├─ Go: Goroutine + Channel                                            │
│  └─ Rust: tokio async                                                  │
│                                                                         │
│  CPU 密集型 (计算)                                                      │
│  ├─ TypeScript: Worker Threads                                         │
│  ├─ Python: multiprocessing (必须)                                     │
│  ├─ Go: Goroutine (自动利用多核)                                       │
│  └─ Rust: rayon / std::thread                                          │
│                                                                         │
│  高并发连接 (10K+)                                                      │
│  ├─ TypeScript: 事件循环 (单线程高效)                                  │
│  ├─ Python: asyncio                                                    │
│  ├─ Go: Goroutine (百万级)                                             │
│  └─ Rust: tokio (百万级)                                               │
│                                                                         │
│  共享状态                                                               │
│  ├─ TypeScript: 避免 (Worker 隔离)                                     │
│  ├─ Python: Manager / 共享内存                                         │
│  ├─ Go: Channel (推荐) / Mutex                                         │
│  └─ Rust: Arc<Mutex> / Channel                                         │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### 常见并发模式

| 模式 | 描述 | 最佳语言 |
|------|------|----------|
| 生产者-消费者 | 队列解耦生产和消费 | Go (Channel) |
| Worker Pool | 固定数量 worker 处理任务 | 全部支持 |
| Fan-out/Fan-in | 分发任务，汇总结果 | Go |
| Pipeline | 链式数据处理 | Go, Rust |
| Pub/Sub | 发布订阅消息 | 全部支持 |
| Actor | 独立状态的消息处理单元 | Rust (actix) |
| CSP | 通过通信共享内存 | Go (原生) |

---

## ❓ 三元表达式

### 语法概览

| 语言 | 语法 | 特点 |
|------|------|------|
| TypeScript | `cond ? a : b` | 标准三元 |
| Python | `a if cond else b` | 可读性优先 |
| Go | 无 | 必须用 if |
| Rust | `if cond { a } else { b }` | if 是表达式 |

### TypeScript 三元表达式

```typescript
// ==================== 基本语法 ====================
const result = condition ? "yes" : "no";

const max = a > b ? a : b;

const status = score >= 90 ? "A" 
             : score >= 80 ? "B" 
             : score >= 70 ? "C" 
             : "D";

// ==================== 类型推断 ====================
const value = condition ? 42 : "hello";  // number | string

// 强制类型
const num = condition ? 42 : 0;  // number

// ==================== 短路替代 ====================
// 与运算短路
const name = user && user.name;

// 或运算短路
const value = input || "default";

// 空值合并
const value = input ?? "default";  // 仅 null/undefined

// ==================== 嵌套 (避免过度嵌套) ====================
// 不推荐
const result = a ? b ? c : d : e ? f : g;

// 推荐：拆分或使用函数
function getResult() {
  if (a) return b ? c : d;
  return e ? f : g;
}

// ==================== 常见用法 ====================
// 条件渲染 (React)
{isLoggedIn ? <Dashboard /> : <Login />}

// 条件样式
const className = `btn ${isActive ? "active" : ""}`;

// 条件属性
const props = {
  disabled: isLoading ? true : undefined,
};
```

### Python 条件表达式

```python
# ==================== 基本语法 ====================
result = "yes" if condition else "no"

max_val = a if a > b else b

status = ("A" if score >= 90 else
          "B" if score >= 80 else
          "C" if score >= 70 else "D")

# ==================== 与三元不同的顺序 ====================
# 其他语言: condition ? true_value : false_value
# Python:   true_value if condition else false_value

# ==================== 短路替代 ====================
# or 短路 (注意: 空字符串、0、[] 也会触发)
name = user_name or "Anonymous"

# and 短路
value = data and data[0]

# ==================== 海象运算符 (Python 3.8+) ====================
if (n := len(data)) > 10:
    print(f"List is too long ({n} elements)")

# 在条件表达式中
result = y if (y := f(x)) else default

# ==================== 常见用法 ====================
# 列表推导中
[x if x > 0 else 0 for x in numbers]

# 字典推导中
{k: v if v else "N/A" for k, v in data.items()}

# 函数默认值
def greet(name=None):
    name = name if name is not None else "World"
    return f"Hello, {name}!"

# 类型注解
from typing import Optional
def process(value: Optional[int]) -> int:
    return value if value is not None else 0
```

### Go 条件赋值

```go
// ==================== Go 没有三元运算符 ====================
// 必须使用 if-else

// 基本形式
var result string
if condition {
    result = "yes"
} else {
    result = "no"
}

// 单行 (不推荐，可读性差)
var result string; if condition { result = "yes" } else { result = "no" }

// ==================== 泛型辅助函数 (Go 1.18+) ====================
func Ternary[T any](cond bool, a, b T) T {
    if cond {
        return a
    }
    return b
}

result := Ternary(condition, "yes", "no")
max := Ternary(a > b, a, b)

// 注意：两个分支都会被求值！
// 不适合有副作用的表达式

// ==================== 惰性求值版本 ====================
func TernaryLazy[T any](cond bool, a, b func() T) T {
    if cond {
        return a()
    }
    return b()
}

result := TernaryLazy(condition, 
    func() string { return expensiveA() },
    func() string { return expensiveB() },
)

// ==================== 常见模式 ====================
// max/min
max := a
if b > a {
    max = b
}

// 使用 math 包
max := math.Max(float64(a), float64(b))

// 条件初始化
var config Config
if env == "prod" {
    config = prodConfig()
} else {
    config = devConfig()
}

// 立即执行函数 (IIFE)
result := func() string {
    if condition {
        return "yes"
    }
    return "no"
}()
```

### Rust 条件表达式

```rust
// ==================== if 是表达式 ====================
let result = if condition { "yes" } else { "no" };

let max = if a > b { a } else { b };

let status = if score >= 90 { "A" }
             else if score >= 80 { "B" }
             else if score >= 70 { "C" }
             else { "D" };

// ==================== 必须有 else (除非返回 ()) ====================
// 错误：缺少 else
// let x = if condition { 42 };

// 正确
let x = if condition { 42 } else { 0 };

// 返回 () 时可省略
if condition {
    println!("yes");
}

// ==================== 类型必须一致 ====================
// 错误：类型不匹配
// let x = if condition { 42 } else { "hello" };

// 正确：使用枚举或 trait object
enum Value {
    Int(i32),
    Str(String),
}
let x = if condition { Value::Int(42) } else { Value::Str("hello".into()) };

// ==================== match 替代复杂条件 ====================
let result = match score {
    90..=100 => "A",
    80..=89 => "B",
    70..=79 => "C",
    _ => "D",
};

// ==================== Option/Result 的条件方法 ====================
// unwrap_or
let value = option.unwrap_or(default);

// unwrap_or_else (惰性)
let value = option.unwrap_or_else(|| compute_default());

// map_or
let len = option.map_or(0, |s| s.len());

// ok_or
let result = option.ok_or("error message")?;

// ==================== 常见用法 ====================
// 条件赋值
let config = if is_prod { prod_config() } else { dev_config() };

// 块表达式
let value = {
    let temp = compute();
    if temp > 0 { temp } else { 0 }
};

// let else (Rust 1.65+)
let Some(value) = option else {
    return Err("no value");
};
```

---

## 🔀 Switch 语句

### 语法概览

| 语言 | 关键字 | 穿透 | 模式匹配 |
|------|--------|------|----------|
| TypeScript | `switch` | 需 break | 有限 |
| Python | `match` (3.10+) | 无 | 完整 |
| Go | `switch` | 默认不穿透 | 有限 |
| Rust | `match` | 无 | 完整 |

### TypeScript switch

```typescript
// ==================== 基本语法 ====================
switch (value) {
  case 1:
    console.log("one");
    break;
  case 2:
    console.log("two");
    break;
  default:
    console.log("other");
}

// ==================== 多值匹配 ====================
switch (day) {
  case "Saturday":
  case "Sunday":
    console.log("Weekend");
    break;
  default:
    console.log("Weekday");
}

// ==================== 穿透 (fallthrough) ====================
switch (grade) {
  case "A":
    console.log("Excellent");
    // 没有 break，穿透到下一个
  case "B":
    console.log("Good");
    break;
}

// ==================== 表达式作为 case ====================
switch (true) {
  case score >= 90:
    grade = "A";
    break;
  case score >= 80:
    grade = "B";
    break;
  default:
    grade = "C";
}

// ==================== 类型收窄 ====================
type Shape = 
  | { kind: "circle"; radius: number }
  | { kind: "rectangle"; width: number; height: number };

function area(shape: Shape): number {
  switch (shape.kind) {
    case "circle":
      return Math.PI * shape.radius ** 2;
    case "rectangle":
      return shape.width * shape.height;
  }
}

// ==================== 穷尽检查 ====================
function exhaustiveCheck(value: never): never {
  throw new Error(`Unhandled value: ${value}`);
}

switch (shape.kind) {
  case "circle":
    // ...
    break;
  case "rectangle":
    // ...
    break;
  default:
    exhaustiveCheck(shape);  // 如果漏掉 case，编译错误
}

// ==================== 对象映射替代 ====================
const handlers: Record<string, () => void> = {
  start: () => console.log("Starting"),
  stop: () => console.log("Stopping"),
  pause: () => console.log("Pausing"),
};

handlers[action]?.() ?? console.log("Unknown action");
```

### Python match (3.10+)

```python
# ==================== 基本语法 ====================
match value:
    case 1:
        print("one")
    case 2:
        print("two")
    case _:
        print("other")

# ==================== 多值匹配 ====================
match day:
    case "Saturday" | "Sunday":
        print("Weekend")
    case _:
        print("Weekday")

# ==================== 模式匹配 ====================
# 序列解构
match point:
    case (0, 0):
        print("Origin")
    case (0, y):
        print(f"On Y axis at {y}")
    case (x, 0):
        print(f"On X axis at {x}")
    case (x, y):
        print(f"Point at ({x}, {y})")

# 字典解构
match config:
    case {"debug": True, "verbose": True}:
        print("Debug verbose mode")
    case {"debug": True}:
        print("Debug mode")
    case _:
        print("Normal mode")

# ==================== 类匹配 ====================
from dataclasses import dataclass

@dataclass
class Point:
    x: int
    y: int

@dataclass
class Circle:
    center: Point
    radius: float

match shape:
    case Circle(center=Point(0, 0), radius=r):
        print(f"Circle at origin with radius {r}")
    case Circle(center=c, radius=r):
        print(f"Circle at {c} with radius {r}")

# ==================== 守卫条件 (Guard) ====================
match point:
    case (x, y) if x == y:
        print(f"On diagonal at {x}")
    case (x, y) if x > y:
        print("Above diagonal")
    case (x, y):
        print("Below diagonal")

# ==================== 捕获匹配值 ====================
match command:
    case ["quit"]:
        quit()
    case ["load", filename]:
        load(filename)
    case ["save", filename]:
        save(filename)
    case ["move", *coords] if len(coords) == 2:
        move(*coords)
    case _:
        print("Unknown command")

# ==================== 类型匹配 ====================
match value:
    case int():
        print("Integer")
    case str():
        print("String")
    case list():
        print("List")
    case _:
        print("Other")

# ==================== 旧版 Python (< 3.10) ====================
# 使用 if-elif-else 或字典映射
handlers = {
    "start": lambda: print("Starting"),
    "stop": lambda: print("Stopping"),
}
handlers.get(action, lambda: print("Unknown"))()
```

### Go switch

```go
// ==================== 基本语法 ====================
switch value {
case 1:
    fmt.Println("one")
case 2:
    fmt.Println("two")
default:
    fmt.Println("other")
}

// 无需 break，默认不穿透

// ==================== 多值匹配 ====================
switch day {
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}

// ==================== 显式穿透 (fallthrough) ====================
switch value {
case 1:
    fmt.Println("one")
    fallthrough  // 继续执行下一个 case
case 2:
    fmt.Println("one or two")
}

// ==================== 无条件 switch ====================
switch {
case score >= 90:
    grade = "A"
case score >= 80:
    grade = "B"
case score >= 70:
    grade = "C"
default:
    grade = "D"
}

// ==================== 初始化语句 ====================
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("macOS")
case "linux":
    fmt.Println("Linux")
default:
    fmt.Printf("%s\n", os)
}

// ==================== 类型 switch ====================
func describe(i interface{}) {
    switch v := i.(type) {
    case int:
        fmt.Printf("Integer: %d\n", v)
    case string:
        fmt.Printf("String: %s\n", v)
    case bool:
        fmt.Printf("Boolean: %t\n", v)
    default:
        fmt.Printf("Unknown type: %T\n", v)
    }
}

// ==================== 泛型约束 (Go 1.18+) ====================
type Number interface {
    int | int64 | float64
}

func sum[T Number](a, b T) T {
    return a + b
}

// ==================== 标签跳转 ====================
OuterLoop:
    for i := 0; i < 10; i++ {
        switch i {
        case 5:
            break OuterLoop  // 跳出外层循环
        }
    }
```

### Rust match

```rust
// ==================== 基本语法 ====================
match value {
    1 => println!("one"),
    2 => println!("two"),
    _ => println!("other"),
}

// match 是表达式
let result = match value {
    1 => "one",
    2 => "two",
    _ => "other",
};

// ==================== 多值匹配 ====================
match day {
    "Saturday" | "Sunday" => println!("Weekend"),
    _ => println!("Weekday"),
}

// ==================== 范围匹配 ====================
match score {
    90..=100 => "A",
    80..=89 => "B",
    70..=79 => "C",
    0..=69 => "D",
    _ => "Invalid",
}

match c {
    'a'..='z' => println!("lowercase"),
    'A'..='Z' => println!("uppercase"),
    '0'..='9' => println!("digit"),
    _ => println!("other"),
}

// ==================== 解构匹配 ====================
// 元组
match point {
    (0, 0) => println!("Origin"),
    (0, y) => println!("On Y axis at {}", y),
    (x, 0) => println!("On X axis at {}", x),
    (x, y) => println!("Point at ({}, {})", x, y),
}

// 结构体
struct Point { x: i32, y: i32 }

match point {
    Point { x: 0, y: 0 } => println!("Origin"),
    Point { x, y: 0 } => println!("On X axis at {}", x),
    Point { x: 0, y } => println!("On Y axis at {}", y),
    Point { x, y } => println!("Point at ({}, {})", x, y),
}

// 枚举
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

match msg {
    Message::Quit => println!("Quit"),
    Message::Move { x, y } => println!("Move to ({}, {})", x, y),
    Message::Write(text) => println!("Write: {}", text),
    Message::ChangeColor(r, g, b) => println!("Color: {}, {}, {}", r, g, b),
}

// ==================== 守卫条件 (Guard) ====================
match num {
    x if x < 0 => println!("Negative"),
    x if x > 0 => println!("Positive"),
    _ => println!("Zero"),
}

match point {
    (x, y) if x == y => println!("On diagonal"),
    (x, y) if x > y => println!("Above diagonal"),
    (x, y) => println!("Below diagonal"),
}

// ==================== 绑定 (@) ====================
match age {
    n @ 0..=12 => println!("Child aged {}", n),
    n @ 13..=19 => println!("Teen aged {}", n),
    n => println!("Adult aged {}", n),
}

// ==================== Option/Result ====================
match option {
    Some(value) => println!("Got: {}", value),
    None => println!("Nothing"),
}

match result {
    Ok(value) => println!("Success: {}", value),
    Err(e) => println!("Error: {}", e),
}

// ==================== 穷尽性检查 ====================
// 必须处理所有可能的情况，否则编译错误
enum Color { Red, Green, Blue }

match color {
    Color::Red => println!("Red"),
    Color::Green => println!("Green"),
    // 编译错误：未处理 Color::Blue
}

// ==================== if let / while let ====================
// 简化单分支 match
if let Some(value) = option {
    println!("Got: {}", value);
}

while let Some(item) = iter.next() {
    println!("{}", item);
}

// let else
let Some(value) = option else {
    return Err("No value");
};

// ==================== matches! 宏 ====================
let is_letter = matches!(c, 'a'..='z' | 'A'..='Z');
let is_some_and_positive = matches!(option, Some(x) if x > 0);
```

### Switch/Match 对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 关键字          │ switch         │ match          │ switch         │ match          │
│ 表达式          │ ❌ (语句)      │ ❌ (语句)      │ ❌ (语句)      │ ✅             │
│ 穿透            │ 默认穿透       │ 无             │ 需 fallthrough │ 无             │
│ 模式匹配        │ 有限           │ 完整           │ 类型 switch    │ 完整           │
│ 守卫条件        │ ❌             │ if             │ ❌             │ if             │
│ 穷尽检查        │ 手动           │ ❌             │ ❌             │ ✅ 编译时      │
│ 解构            │ ❌             │ ✅             │ ❌             │ ✅             │
│ 范围匹配        │ ❌             │ ❌             │ ❌             │ ✅             │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

---

## 💻 CLI 开发

### CLI 开发概览

| 语言 | 参数解析库 | 推荐框架 | 分发方式 |
|------|------------|----------|----------|
| TypeScript | commander, yargs | oclif | npm, pkg |
| Python | argparse, click | typer | pip, PyInstaller |
| Go | flag, cobra | cobra | 单二进制 |
| Rust | clap, structopt | clap | 单二进制 |

### TypeScript CLI 开发

```typescript
// ==================== 原生参数读取 ====================
// process.argv: ['node', 'script.js', 'arg1', 'arg2']
const args = process.argv.slice(2);
console.log(args);

// ==================== commander ====================
import { Command } from 'commander';

const program = new Command();

program
  .name('mycli')
  .description('My awesome CLI tool')
  .version('1.0.0');

program
  .command('greet <name>')
  .description('Greet someone')
  .option('-l, --loud', 'Say it loudly')
  .option('-t, --times <n>', 'Repeat times', '1')
  .action((name, options) => {
    const greeting = `Hello, ${name}!`;
    const times = parseInt(options.times);
    for (let i = 0; i < times; i++) {
      console.log(options.loud ? greeting.toUpperCase() : greeting);
    }
  });

program
  .command('config')
  .description('Manage configuration')
  .option('--get <key>', 'Get config value')
  .option('--set <key=value>', 'Set config value')
  .action((options) => {
    if (options.get) {
      console.log(getConfig(options.get));
    }
    if (options.set) {
      const [key, value] = options.set.split('=');
      setConfig(key, value);
    }
  });

program.parse();

// ==================== yargs ====================
import yargs from 'yargs';
import { hideBin } from 'yargs/helpers';

yargs(hideBin(process.argv))
  .command('greet <name>', 'Greet someone', (yargs) => {
    return yargs.positional('name', {
      describe: 'Name to greet',
      type: 'string',
    });
  }, (argv) => {
    console.log(`Hello, ${argv.name}!`);
  })
  .option('verbose', {
    alias: 'v',
    type: 'boolean',
    description: 'Run with verbose logging',
  })
  .demandCommand(1)
  .help()
  .parse();

// ==================== 错误退出 ====================
// 退出码约定:
// 0: 成功
// 1: 一般错误
// 2: 参数错误

function main() {
  try {
    const result = doSomething();
    console.log(result);
    process.exit(0);
  } catch (error) {
    console.error('Error:', error.message);
    process.exit(1);
  }
}

// 未捕获异常
process.on('uncaughtException', (error) => {
  console.error('Uncaught:', error);
  process.exit(1);
});

// 信号处理
process.on('SIGINT', () => {
  console.log('\nInterrupted');
  process.exit(130);  // 128 + 2 (SIGINT)
});

// ==================== 用户交互 ====================
import inquirer from 'inquirer';

const answers = await inquirer.prompt([
  {
    type: 'input',
    name: 'name',
    message: 'What is your name?',
  },
  {
    type: 'list',
    name: 'color',
    message: 'Pick a color:',
    choices: ['Red', 'Green', 'Blue'],
  },
  {
    type: 'confirm',
    name: 'confirm',
    message: 'Are you sure?',
  },
]);

// ==================== 输出美化 ====================
import chalk from 'chalk';
import ora from 'ora';

console.log(chalk.green('Success!'));
console.log(chalk.red.bold('Error!'));
console.log(chalk.yellow('Warning'));

const spinner = ora('Loading...').start();
await doAsyncWork();
spinner.succeed('Done!');
// spinner.fail('Failed!');
```

### Python CLI 开发

```python
import sys
import argparse

# ==================== 原生参数读取 ====================
# sys.argv: ['script.py', 'arg1', 'arg2']
args = sys.argv[1:]
print(args)

# ==================== argparse (标准库) ====================
parser = argparse.ArgumentParser(
    prog='mycli',
    description='My awesome CLI tool'
)

parser.add_argument('name', help='Name to greet')
parser.add_argument('-l', '--loud', action='store_true', help='Say it loudly')
parser.add_argument('-t', '--times', type=int, default=1, help='Repeat times')
parser.add_argument('--version', action='version', version='%(prog)s 1.0.0')

args = parser.parse_args()

greeting = f"Hello, {args.name}!"
for _ in range(args.times):
    print(greeting.upper() if args.loud else greeting)

# 子命令
parser = argparse.ArgumentParser()
subparsers = parser.add_subparsers(dest='command')

greet_parser = subparsers.add_parser('greet', help='Greet someone')
greet_parser.add_argument('name')

config_parser = subparsers.add_parser('config', help='Manage config')
config_parser.add_argument('--get', metavar='KEY')
config_parser.add_argument('--set', metavar='KEY=VALUE')

# ==================== click ====================
import click

@click.group()
@click.version_option('1.0.0')
def cli():
    """My awesome CLI tool"""
    pass

@cli.command()
@click.argument('name')
@click.option('-l', '--loud', is_flag=True, help='Say it loudly')
@click.option('-t', '--times', default=1, help='Repeat times')
def greet(name, loud, times):
    """Greet someone"""
    greeting = f"Hello, {name}!"
    for _ in range(times):
        click.echo(greeting.upper() if loud else greeting)

@cli.command()
@click.option('--get', 'key', help='Get config value')
@click.option('--set', 'keyvalue', help='Set config value')
def config(key, keyvalue):
    """Manage configuration"""
    if key:
        click.echo(get_config(key))
    if keyvalue:
        k, v = keyvalue.split('=')
        set_config(k, v)

if __name__ == '__main__':
    cli()

# ==================== typer (推荐，类型提示) ====================
import typer

app = typer.Typer(help="My awesome CLI tool")

@app.command()
def greet(
    name: str = typer.Argument(..., help="Name to greet"),
    loud: bool = typer.Option(False, "--loud", "-l", help="Say it loudly"),
    times: int = typer.Option(1, "--times", "-t", help="Repeat times"),
):
    """Greet someone"""
    greeting = f"Hello, {name}!"
    for _ in range(times):
        typer.echo(greeting.upper() if loud else greeting)

@app.command()
def config(
    get: str = typer.Option(None, "--get", help="Get config value"),
    set_: str = typer.Option(None, "--set", help="Set config value"),
):
    """Manage configuration"""
    if get:
        typer.echo(get_config(get))
    if set_:
        k, v = set_.split('=')
        set_config(k, v)

if __name__ == "__main__":
    app()

# ==================== 错误退出 ====================
import sys

def main():
    try:
        result = do_something()
        print(result)
        sys.exit(0)
    except ValueError as e:
        print(f"Invalid input: {e}", file=sys.stderr)
        sys.exit(2)  # 参数错误
    except Exception as e:
        print(f"Error: {e}", file=sys.stderr)
        sys.exit(1)

# click/typer 方式
@app.command()
def cmd():
    if error_condition:
        raise typer.Exit(code=1)
    
    # 或带消息
    raise typer.Abort()  # 退出码 1

# ==================== 用户交互 ====================
# 基础输入
name = input("What is your name? ")

# rich 库 (推荐)
from rich.prompt import Prompt, Confirm
from rich.console import Console

console = Console()

name = Prompt.ask("What is your name?")
color = Prompt.ask("Pick a color", choices=["red", "green", "blue"])
confirmed = Confirm.ask("Are you sure?")

# 输出美化
console.print("[green]Success![/green]")
console.print("[red bold]Error![/red bold]")

# 进度条
from rich.progress import track
for item in track(items, description="Processing..."):
    process(item)
```

### Go CLI 开发

```go
package main

import (
    "flag"
    "fmt"
    "os"
)

// ==================== 原生参数读取 ====================
func main() {
    // os.Args: ["./program", "arg1", "arg2"]
    args := os.Args[1:]
    fmt.Println(args)
}

// ==================== flag (标准库) ====================
func main() {
    // 定义 flag
    name := flag.String("name", "World", "Name to greet")
    loud := flag.Bool("loud", false, "Say it loudly")
    times := flag.Int("times", 1, "Repeat times")
    
    // 解析
    flag.Parse()
    
    // 位置参数
    args := flag.Args()
    
    greeting := fmt.Sprintf("Hello, %s!", *name)
    for i := 0; i < *times; i++ {
        if *loud {
            fmt.Println(strings.ToUpper(greeting))
        } else {
            fmt.Println(greeting)
        }
    }
}

// ==================== cobra (推荐) ====================
import "github.com/spf13/cobra"

var (
    loud  bool
    times int
)

var rootCmd = &cobra.Command{
    Use:     "mycli",
    Short:   "My awesome CLI tool",
    Version: "1.0.0",
}

var greetCmd = &cobra.Command{
    Use:   "greet <name>",
    Short: "Greet someone",
    Args:  cobra.ExactArgs(1),
    Run: func(cmd *cobra.Command, args []string) {
        name := args[0]
        greeting := fmt.Sprintf("Hello, %s!", name)
        for i := 0; i < times; i++ {
            if loud {
                fmt.Println(strings.ToUpper(greeting))
            } else {
                fmt.Println(greeting)
            }
        }
    },
}

var configCmd = &cobra.Command{
    Use:   "config",
    Short: "Manage configuration",
    Run: func(cmd *cobra.Command, args []string) {
        if key, _ := cmd.Flags().GetString("get"); key != "" {
            fmt.Println(getConfig(key))
        }
        if kv, _ := cmd.Flags().GetString("set"); kv != "" {
            parts := strings.SplitN(kv, "=", 2)
            setConfig(parts[0], parts[1])
        }
    },
}

func init() {
    greetCmd.Flags().BoolVarP(&loud, "loud", "l", false, "Say it loudly")
    greetCmd.Flags().IntVarP(&times, "times", "t", 1, "Repeat times")
    
    configCmd.Flags().String("get", "", "Get config value")
    configCmd.Flags().String("set", "", "Set config value")
    
    rootCmd.AddCommand(greetCmd)
    rootCmd.AddCommand(configCmd)
}

func main() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

// ==================== 错误退出 ====================
func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "Error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    // 业务逻辑
    if badInput {
        return fmt.Errorf("invalid input: %s", input)
    }
    return nil
}

// 自定义退出码
const (
    ExitSuccess     = 0
    ExitError       = 1
    ExitUsageError  = 2
)

// ==================== 用户交互 ====================
import "github.com/AlecAivazis/survey/v2"

var name string
prompt := &survey.Input{
    Message: "What is your name?",
}
survey.AskOne(prompt, &name)

var color string
survey.AskOne(&survey.Select{
    Message: "Pick a color:",
    Options: []string{"Red", "Green", "Blue"},
}, &color)

var confirmed bool
survey.AskOne(&survey.Confirm{
    Message: "Are you sure?",
}, &confirmed)

// ==================== 输出美化 ====================
import "github.com/fatih/color"

color.Green("Success!")
color.Red("Error!")
color.Yellow("Warning")

// 或使用模板
red := color.New(color.FgRed, color.Bold)
red.Println("Bold Red!")

// 进度条
import "github.com/schollz/progressbar/v3"

bar := progressbar.Default(100)
for i := 0; i < 100; i++ {
    bar.Add(1)
    time.Sleep(10 * time.Millisecond)
}
```

### Rust CLI 开发

```rust
use std::env;
use std::process;

// ==================== 原生参数读取 ====================
fn main() {
    // std::env::args(): ["./program", "arg1", "arg2"]
    let args: Vec<String> = env::args().collect();
    println!("{:?}", &args[1..]);
}

// ==================== clap (推荐) ====================
use clap::{Parser, Subcommand, Args};

#[derive(Parser)]
#[command(name = "mycli")]
#[command(about = "My awesome CLI tool", version)]
struct Cli {
    #[command(subcommand)]
    command: Commands,
}

#[derive(Subcommand)]
enum Commands {
    /// Greet someone
    Greet(GreetArgs),
    /// Manage configuration
    Config(ConfigArgs),
}

#[derive(Args)]
struct GreetArgs {
    /// Name to greet
    name: String,
    
    /// Say it loudly
    #[arg(short, long)]
    loud: bool,
    
    /// Repeat times
    #[arg(short, long, default_value_t = 1)]
    times: u32,
}

#[derive(Args)]
struct ConfigArgs {
    /// Get config value
    #[arg(long)]
    get: Option<String>,
    
    /// Set config value (KEY=VALUE)
    #[arg(long)]
    set: Option<String>,
}

fn main() {
    let cli = Cli::parse();
    
    match cli.command {
        Commands::Greet(args) => {
            let greeting = format!("Hello, {}!", args.name);
            for _ in 0..args.times {
                if args.loud {
                    println!("{}", greeting.to_uppercase());
                } else {
                    println!("{}", greeting);
                }
            }
        }
        Commands::Config(args) => {
            if let Some(key) = args.get {
                println!("{}", get_config(&key));
            }
            if let Some(kv) = args.set {
                let parts: Vec<&str> = kv.splitn(2, '=').collect();
                set_config(parts[0], parts[1]);
            }
        }
    }
}

// ==================== 错误退出 ====================
use std::process::ExitCode;

fn main() -> ExitCode {
    match run() {
        Ok(()) => ExitCode::SUCCESS,
        Err(e) => {
            eprintln!("Error: {}", e);
            ExitCode::FAILURE
        }
    }
}

fn run() -> Result<(), Box<dyn std::error::Error>> {
    // 业务逻辑
    if bad_input {
        return Err("invalid input".into());
    }
    Ok(())
}

// 或使用 anyhow
use anyhow::{Context, Result, bail};

fn main() -> Result<()> {
    let config = load_config()
        .context("Failed to load config")?;
    
    if config.invalid {
        bail!("Invalid configuration");
    }
    
    Ok(())
}

// 自定义退出码
fn main() {
    let code = match run() {
        Ok(()) => 0,
        Err(e) if e.is_usage_error() => 2,
        Err(_) => 1,
    };
    process::exit(code);
}

// ==================== 用户交互 ====================
use dialoguer::{Input, Select, Confirm};

let name: String = Input::new()
    .with_prompt("What is your name?")
    .interact_text()?;

let colors = vec!["Red", "Green", "Blue"];
let selection = Select::new()
    .with_prompt("Pick a color")
    .items(&colors)
    .interact()?;

let confirmed = Confirm::new()
    .with_prompt("Are you sure?")
    .interact()?;

// ==================== 输出美化 ====================
use colored::*;

println!("{}", "Success!".green());
println!("{}", "Error!".red().bold());
println!("{}", "Warning".yellow());

// 进度条
use indicatif::{ProgressBar, ProgressStyle};

let pb = ProgressBar::new(100);
pb.set_style(ProgressStyle::default_bar()
    .template("{spinner:.green} [{bar:40.cyan/blue}] {pos}/{len} ({eta})")?);

for _ in 0..100 {
    pb.inc(1);
    std::thread::sleep(Duration::from_millis(10));
}
pb.finish_with_message("Done!");
```

### CLI 开发对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 参数库          │ commander      │ typer/click    │ cobra          │ clap           │
│ 标准库支持      │ process.argv   │ argparse       │ flag           │ std::env       │
│ 类型安全        │ ✅             │ ✅ (typer)     │ ✅             │ ✅             │
│ 子命令          │ ✅             │ ✅             │ ✅             │ ✅             │
│ 自动补全        │ 插件           │ ✅             │ ✅             │ ✅             │
│ 交互式          │ inquirer       │ rich           │ survey         │ dialoguer      │
│ 颜色输出        │ chalk          │ rich           │ color          │ colored        │
│ 分发            │ npm/pkg        │ pip/PyInstaller│ 单二进制       │ 单二进制       │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 退出码约定

```
┌──────────┬────────────────────────────────────────┐
│ 退出码   │ 含义                                   │
├──────────┼────────────────────────────────────────┤
│ 0        │ 成功                                   │
│ 1        │ 一般错误                               │
│ 2        │ 命令行参数错误                         │
│ 126      │ 命令不可执行                           │
│ 127      │ 命令未找到                             │
│ 128+N    │ 被信号 N 终止 (如 130 = SIGINT)        │
│ 255      │ 退出码超出范围                         │
└──────────┴────────────────────────────────────────┘
```

### CLI 项目结构示例

```
mycli/
├── src/
│   ├── main.ts/py/go/rs    # 入口
│   ├── commands/           # 子命令
│   │   ├── greet.ts
│   │   └── config.ts
│   ├── utils/              # 工具函数
│   └── config/             # 配置管理
├── tests/                  # 测试
├── package.json / pyproject.toml / go.mod / Cargo.toml
└── README.md
```

---

## 📝 格式化

### 格式化概览

| 类型 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 字符串插值 | 模板字符串 | f-string | fmt.Sprintf | format! |
| 代码格式化 | Prettier | Black/Ruff | gofmt | rustfmt |
| 数字格式 | Intl.NumberFormat | format() | fmt | format! |

### TypeScript 格式化

```typescript
// ==================== 字符串格式化 ====================

// 模板字符串 (推荐)
const name = "World";
const age = 25;
console.log(`Hello, ${name}! You are ${age} years old.`);

// 表达式
console.log(`2 + 2 = ${2 + 2}`);
console.log(`Upper: ${name.toUpperCase()}`);

// 多行
const html = `
  <div>
    <h1>${title}</h1>
    <p>${content}</p>
  </div>
`;

// 标签模板
function highlight(strings: TemplateStringsArray, ...values: any[]) {
  return strings.reduce((acc, str, i) => 
    acc + str + (values[i] ? `<mark>${values[i]}</mark>` : ''), '');
}
highlight`Hello ${name}, you are ${age} years old.`;

// ==================== 数字格式化 ====================

const num = 1234567.89;

// toLocaleString
num.toLocaleString('en-US');           // "1,234,567.89"
num.toLocaleString('de-DE');           // "1.234.567,89"
num.toLocaleString('zh-CN');           // "1,234,567.89"

// Intl.NumberFormat
new Intl.NumberFormat('en-US').format(num);  // "1,234,567.89"

// 货币
new Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
}).format(num);  // "$1,234,567.89"

new Intl.NumberFormat('zh-CN', {
  style: 'currency',
  currency: 'CNY',
}).format(num);  // "¥1,234,567.89"

// 百分比
new Intl.NumberFormat('en-US', {
  style: 'percent',
  minimumFractionDigits: 2,
}).format(0.1234);  // "12.34%"

// 紧凑表示
new Intl.NumberFormat('en-US', {
  notation: 'compact',
}).format(1234567);  // "1.2M"

// 固定小数位
num.toFixed(2);           // "1234567.89"
num.toPrecision(4);       // "1.235e+6"
num.toExponential(2);     // "1.23e+6"

// 进制
(255).toString(16);       // "ff"
(255).toString(2);        // "11111111"
(255).toString(8);        // "377"

// 填充
String(5).padStart(3, '0');  // "005"
String(5).padEnd(3, '0');    // "500"

// ==================== 日期格式化 ====================

const date = new Date();

date.toLocaleDateString('zh-CN');  // "2024/3/15"
date.toLocaleTimeString('zh-CN');  // "10:30:00"
date.toLocaleString('zh-CN');      // "2024/3/15 10:30:00"

new Intl.DateTimeFormat('zh-CN', {
  year: 'numeric',
  month: 'long',
  day: 'numeric',
  weekday: 'long',
}).format(date);  // "2024年3月15日星期五"

// ==================== 代码格式化工具 ====================
// Prettier: npx prettier --write .
// ESLint:   npx eslint --fix .
// 配置: .prettierrc
/*
{
  "semi": true,
  "singleQuote": true,
  "tabWidth": 2,
  "trailingComma": "es5"
}
*/
```

### Python 格式化

```python
# ==================== 字符串格式化 ====================

name = "World"
age = 25

# f-string (推荐, Python 3.6+)
f"Hello, {name}! You are {age} years old."
f"2 + 2 = {2 + 2}"
f"Upper: {name.upper()}"

# 格式规范
f"{age:03d}"           # "025" (补零)
f"{3.14159:.2f}"       # "3.14" (小数位)
f"{1000000:,}"         # "1,000,000" (千分位)
f"{0.25:.1%}"          # "25.0%" (百分比)
f"{255:x}"             # "ff" (十六进制)
f"{255:b}"             # "11111111" (二进制)
f"{name:>10}"          # "     World" (右对齐)
f"{name:<10}"          # "World     " (左对齐)
f"{name:^10}"          # "  World   " (居中)
f"{name:*^10}"         # "**World***" (填充)

# 调试 (Python 3.8+)
f"{name=}"             # "name='World'"
f"{age=}"              # "age=25"

# format() 方法
"Hello, {}!".format(name)
"Hello, {0}! {0} is {1}.".format(name, age)
"Hello, {name}!".format(name=name)

# % 格式化 (旧式)
"Hello, %s! You are %d years old." % (name, age)
"Pi is %.2f" % 3.14159

# ==================== 数字格式化 ====================

num = 1234567.89

# format() 内置函数
format(num, ',')           # "1,234,567.89"
format(num, ',.2f')        # "1,234,567.89"
format(0.25, '.1%')        # "25.0%"
format(255, 'x')           # "ff"
format(255, '08b')         # "11111111"

# locale 模块
import locale
locale.setlocale(locale.LC_ALL, 'en_US.UTF-8')
locale.format_string('%d', num, grouping=True)

# Decimal 精确格式化
from decimal import Decimal
Decimal('1234.5').quantize(Decimal('0.00'))  # Decimal('1234.50')

# ==================== 日期格式化 ====================

from datetime import datetime

now = datetime.now()

now.strftime('%Y-%m-%d')          # "2024-03-15"
now.strftime('%Y年%m月%d日')       # "2024年03月15日"
now.strftime('%H:%M:%S')          # "10:30:00"
now.strftime('%Y-%m-%d %H:%M:%S') # "2024-03-15 10:30:00"
now.strftime('%A, %B %d, %Y')     # "Friday, March 15, 2024"

# f-string 中使用
f"{now:%Y-%m-%d}"                 # "2024-03-15"

# ==================== 多行/文本格式化 ====================

import textwrap

# 去除缩进
text = """
    Hello
    World
"""
textwrap.dedent(text)

# 自动换行
textwrap.fill(long_text, width=80)

# pprint 美化输出
from pprint import pprint
pprint(complex_data, width=60, depth=2)

# ==================== 代码格式化工具 ====================
# Black: black .
# Ruff:  ruff format .
# isort: isort .

# pyproject.toml
"""
[tool.black]
line-length = 88
target-version = ['py310']

[tool.ruff]
line-length = 88
select = ["E", "F", "I"]
"""
```

### Go 格式化

```go
import "fmt"

// ==================== 字符串格式化 ====================

name := "World"
age := 25

// fmt.Sprintf
fmt.Sprintf("Hello, %s! You are %d years old.", name, age)

// 常用格式动词
fmt.Sprintf("%v", value)     // 默认格式
fmt.Sprintf("%+v", struct_)  // 带字段名
fmt.Sprintf("%#v", value)    // Go 语法表示
fmt.Sprintf("%T", value)     // 类型

fmt.Sprintf("%t", true)      // "true" (布尔)
fmt.Sprintf("%d", 42)        // "42" (十进制)
fmt.Sprintf("%b", 42)        // "101010" (二进制)
fmt.Sprintf("%o", 42)        // "52" (八进制)
fmt.Sprintf("%x", 42)        // "2a" (十六进制小写)
fmt.Sprintf("%X", 42)        // "2A" (十六进制大写)

fmt.Sprintf("%f", 3.14)      // "3.140000" (浮点)
fmt.Sprintf("%.2f", 3.14159) // "3.14" (精度)
fmt.Sprintf("%e", 1234.5)    // "1.234500e+03" (科学计数)
fmt.Sprintf("%g", 1234.5)    // "1234.5" (紧凑)

fmt.Sprintf("%s", "hello")   // "hello" (字符串)
fmt.Sprintf("%q", "hello")   // "\"hello\"" (带引号)
fmt.Sprintf("%c", 65)        // "A" (字符)

fmt.Sprintf("%p", &value)    // "0xc000..." (指针)

// 宽度和精度
fmt.Sprintf("%5d", 42)       // "   42" (宽度 5)
fmt.Sprintf("%-5d", 42)      // "42   " (左对齐)
fmt.Sprintf("%05d", 42)      // "00042" (补零)
fmt.Sprintf("%8.2f", 3.14)   // "    3.14"

// ==================== 数字格式化 ====================

import "golang.org/x/text/language"
import "golang.org/x/text/message"

num := 1234567.89

// 千分位分隔
p := message.NewPrinter(language.English)
p.Sprintf("%d", 1234567)     // "1,234,567"

p = message.NewPrinter(language.German)
p.Sprintf("%d", 1234567)     // "1.234.567"

// 手动实现千分位
func formatNumber(n int) string {
    s := strconv.Itoa(n)
    // ... 添加分隔符逻辑
}

// ==================== 日期格式化 ====================

import "time"

now := time.Now()

// Go 使用参考时间: Mon Jan 2 15:04:05 MST 2006
now.Format("2006-01-02")           // "2024-03-15"
now.Format("2006年01月02日")        // "2024年03月15日"
now.Format("15:04:05")             // "10:30:00"
now.Format("2006-01-02 15:04:05")  // "2024-03-15 10:30:00"
now.Format(time.RFC3339)           // "2024-03-15T10:30:00+08:00"

// 预定义格式
now.Format(time.Kitchen)           // "10:30AM"
now.Format(time.RFC822)            // "15 Mar 24 10:30 CST"

// ==================== JSON 格式化 ====================

import "encoding/json"

// 紧凑 JSON
data, _ := json.Marshal(obj)

// 美化 JSON
data, _ := json.MarshalIndent(obj, "", "  ")

// ==================== 代码格式化工具 ====================
// gofmt:    gofmt -w .
// goimports: goimports -w .
// go fmt:   go fmt ./...

// 无需配置，强制统一风格
```

### Rust 格式化

```rust
// ==================== 字符串格式化 ====================

let name = "World";
let age = 25;

// format! 宏
format!("Hello, {}! You are {} years old.", name, age);
format!("Hello, {name}! You are {age} years old.");  // 命名参数

// 位置参数
format!("{0} {1} {0}", "Hello", "World");  // "Hello World Hello"

// 格式规范
format!("{:5}", 42);         // "   42" (宽度)
format!("{:<5}", 42);        // "42   " (左对齐)
format!("{:>5}", 42);        // "   42" (右对齐)
format!("{:^5}", 42);        // " 42  " (居中)
format!("{:0>5}", 42);       // "00042" (填充)
format!("{:*^5}", 42);       // "*42**"

format!("{:.2}", 3.14159);   // "3.14" (精度)
format!("{:8.2}", 3.14);     // "    3.14"

format!("{:b}", 42);         // "101010" (二进制)
format!("{:o}", 42);         // "52" (八进制)
format!("{:x}", 255);        // "ff" (十六进制)
format!("{:X}", 255);        // "FF"
format!("{:#x}", 255);       // "0xff" (带前缀)
format!("{:#b}", 42);        // "0b101010"

format!("{:e}", 1234.5);     // "1.2345e3" (科学计数)
format!("{:E}", 1234.5);     // "1.2345E3"

format!("{:?}", value);      // Debug 格式
format!("{:#?}", value);     // Debug 美化格式
format!("{:p}", &value);     // 指针地址

// ==================== 数字格式化 ====================

let num = 1234567.89f64;

// 千分位 (需要外部 crate 或手动实现)
// 使用 num-format crate
use num_format::{Locale, ToFormattedString};
let formatted = 1234567.to_formatted_string(&Locale::en);  // "1,234,567"

// 手动实现
fn format_number(n: i64) -> String {
    n.to_string()
        .as_bytes()
        .rchunks(3)
        .rev()
        .map(|chunk| std::str::from_utf8(chunk).unwrap())
        .collect::<Vec<_>>()
        .join(",")
}

// ==================== 日期格式化 ====================

use chrono::{DateTime, Utc, Local};

let now: DateTime<Local> = Local::now();

now.format("%Y-%m-%d").to_string();           // "2024-03-15"
now.format("%Y年%m月%d日").to_string();        // "2024年03月15日"
now.format("%H:%M:%S").to_string();           // "10:30:00"
now.format("%Y-%m-%d %H:%M:%S").to_string();  // "2024-03-15 10:30:00"
now.format("%A, %B %d, %Y").to_string();      // "Friday, March 15, 2024"

now.to_rfc3339();                             // "2024-03-15T10:30:00+08:00"
now.to_rfc2822();                             // "Fri, 15 Mar 2024 10:30:00 +0800"

// ==================== 自定义格式化 ====================

use std::fmt;

struct Point {
    x: i32,
    y: i32,
}

// Display trait (用户友好)
impl fmt::Display for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        write!(f, "({}, {})", self.x, self.y)
    }
}

// Debug trait (调试用)
impl fmt::Debug for Point {
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Point")
            .field("x", &self.x)
            .field("y", &self.y)
            .finish()
    }
}

// 或使用 derive
#[derive(Debug)]
struct Point { x: i32, y: i32 }

let p = Point { x: 10, y: 20 };
format!("{}", p);    // "(10, 20)" (Display)
format!("{:?}", p);  // "Point { x: 10, y: 20 }" (Debug)

// ==================== 输出宏 ====================

println!("Hello, World!");           // 带换行
print!("Hello");                     // 不带换行
eprintln!("Error!");                 // 标准错误
eprint!("Warning");                  // 标准错误，不带换行

dbg!(value);                         // 调试输出 (含文件行号)
// [src/main.rs:10] value = 42

// ==================== 代码格式化工具 ====================
// rustfmt: rustfmt src/*.rs
// cargo:   cargo fmt

// rustfmt.toml
/*
max_width = 100
tab_spaces = 4
edition = "2021"
use_small_heuristics = "Default"
*/
```

### 格式化对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 字符串插值      │ `${}`          │ f"{}"          │ fmt.Sprintf    │ format!        │
│ 十六进制        │ toString(16)   │ {:x}           │ %x             │ {:x}           │
│ 补零            │ padStart       │ {:05d}         │ %05d           │ {:05}          │
│ 小数精度        │ toFixed(2)     │ {:.2f}         │ %.2f           │ {:.2}          │
│ 千分位          │ toLocaleString │ {:,}           │ message.Printer│ 外部 crate     │
│ 调试输出        │ console.log    │ pprint         │ %#v            │ {:?} / dbg!    │
│ 代码格式化      │ Prettier       │ Black/Ruff     │ gofmt          │ rustfmt        │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 格式化速查表

```
┌──────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 需求         │ TypeScript     │ Python         │ Go             │ Rust           │
├──────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 整数         │ `${n}`         │ f"{n}"         │ %d             │ {}             │
│ 浮点数       │ `${n}`         │ f"{n}"         │ %f             │ {}             │
│ 2位小数      │ n.toFixed(2)   │ f"{n:.2f}"     │ %.2f           │ {:.2}          │
│ 补零5位      │ padStart(5,'0')│ f"{n:05d}"     │ %05d           │ {:05}          │
│ 左对齐10位   │ padEnd(10)     │ f"{s:<10}"     │ %-10s          │ {:<10}         │
│ 右对齐10位   │ padStart(10)   │ f"{s:>10}"     │ %10s           │ {:>10}         │
│ 居中10位     │ 手动           │ f"{s:^10}"     │ 手动           │ {:^10}         │
│ 十六进制     │ n.toString(16) │ f"{n:x}"       │ %x             │ {:x}           │
│ 二进制       │ n.toString(2)  │ f"{n:b}"       │ %b             │ {:b}           │
│ 百分比       │ Intl           │ f"{n:.1%}"     │ 手动           │ 手动           │
│ 千分位       │ toLocaleString │ f"{n:,}"       │ message        │ num-format     │
│ 科学计数     │ toExponential  │ f"{n:e}"       │ %e             │ {:e}           │
└──────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
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
