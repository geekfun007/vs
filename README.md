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

## 📋 List / Tuple / Set 集合类型

### 集合类型概览

| 类型 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 动态数组 | `Array<T>` | `list` | `[]T` (slice) | `Vec<T>` |
| 固定数组 | `readonly T[]` | `tuple` | `[N]T` | `[T; N]` |
| 集合 | `Set<T>` | `set` | `map[T]struct{}` | `HashSet<T>` |

### TypeScript 数组操作

```typescript
// ==================== 创建 ====================
const arr: number[] = [1, 2, 3];
const arr2: Array<number> = [1, 2, 3];
const empty: string[] = [];
const filled = new Array(5).fill(0);  // [0, 0, 0, 0, 0]
const range = Array.from({ length: 5 }, (_, i) => i);  // [0, 1, 2, 3, 4]

// ==================== 增 (Create) ====================
arr.push(4);              // 末尾添加 [1,2,3,4]
arr.unshift(0);           // 开头添加 [0,1,2,3,4]
arr.splice(2, 0, 1.5);    // 指定位置插入
const newArr = [...arr, 5];  // 不可变添加

// ==================== 删 (Delete) ====================
arr.pop();                // 删除末尾
arr.shift();              // 删除开头
arr.splice(1, 1);         // 删除索引 1 的元素
const filtered = arr.filter(x => x !== 2);  // 不可变删除

// 删除多个元素
const toRemove = new Set([1, 3, 5]);
const result = arr.filter(x => !toRemove.has(x));

// 按索引删除多个
const indicesToRemove = [0, 2, 4];
const result2 = arr.filter((_, i) => !indicesToRemove.includes(i));

// 清空
arr.length = 0;
arr.splice(0, arr.length);

// ==================== 改 (Update) ====================
arr[0] = 10;              // 直接修改
arr.splice(1, 1, 20);     // 替换
arr.fill(0, 1, 3);        // 填充索引 1-2

// ==================== 查 (Read) ====================
arr[0];                   // 索引访问
arr.at(-1);               // 负索引 (ES2022)
arr.indexOf(2);           // 查找索引
arr.includes(2);          // 是否包含
arr.find(x => x > 2);     // 查找元素
arr.findIndex(x => x > 2);// 查找索引
arr.slice(1, 3);          // 切片 [1, 3)

// ==================== IndexError 处理 ====================
// JavaScript 返回 undefined，不抛异常
arr[100];                 // undefined
arr[-1];                  // undefined (用 at(-1))

// 安全访问
function safeGet<T>(arr: T[], index: number): T | undefined {
  if (index < 0) index = arr.length + index;
  if (index < 0 || index >= arr.length) return undefined;
  return arr[index];
}

// 断言存在
function assertGet<T>(arr: T[], index: number): T {
  if (index < 0 || index >= arr.length) {
    throw new RangeError(`Index ${index} out of bounds`);
  }
  return arr[index];
}

// ==================== Set ====================
const set = new Set<number>([1, 2, 3]);
set.add(4);
set.delete(2);
set.has(1);               // true
set.size;                 // 3
set.clear();

// Set 操作
const a = new Set([1, 2, 3]);
const b = new Set([2, 3, 4]);
const union = new Set([...a, ...b]);           // 并集
const intersection = new Set([...a].filter(x => b.has(x)));  // 交集
const difference = new Set([...a].filter(x => !b.has(x)));   // 差集

// ==================== Tuple (只读数组) ====================
const tuple: readonly [string, number] = ["hello", 42];
// tuple[0] = "world";  // 错误：只读
const [str, num] = tuple;  // 解构
```

### Python 列表操作

```python
# ==================== 创建 ====================
arr = [1, 2, 3]
empty = []
filled = [0] * 5          # [0, 0, 0, 0, 0]
range_list = list(range(5))  # [0, 1, 2, 3, 4]
comprehension = [x * 2 for x in range(5)]

# ==================== 增 (Create) ====================
arr.append(4)             # 末尾添加
arr.insert(0, 0)          # 指定位置插入
arr.extend([5, 6])        # 扩展多个
arr += [7, 8]             # 合并
new_arr = [*arr, 9]       # 不可变添加

# ==================== 删 (Delete) ====================
arr.pop()                 # 删除末尾并返回
arr.pop(0)                # 删除指定索引
arr.remove(2)             # 删除第一个值为 2 的元素
del arr[1]                # 删除索引 1
del arr[1:3]              # 删除切片

# 删除多个元素 (按值)
to_remove = {1, 3, 5}
arr = [x for x in arr if x not in to_remove]

# 删除多个元素 (按索引)
indices = {0, 2, 4}
arr = [x for i, x in enumerate(arr) if i not in indices]

# 清空
arr.clear()
arr = []
del arr[:]

# ==================== 改 (Update) ====================
arr[0] = 10               # 直接修改
arr[1:3] = [20, 30]       # 切片替换

# ==================== 查 (Read) ====================
arr[0]                    # 索引访问
arr[-1]                   # 负索引 (最后一个)
arr[1:3]                  # 切片 [1, 3)
arr[::2]                  # 步长切片
arr[::-1]                 # 反转
arr.index(2)              # 查找索引 (不存在抛 ValueError)
2 in arr                  # 是否包含
arr.count(2)              # 计数

# ==================== IndexError 处理 ====================
try:
    value = arr[100]
except IndexError:
    print("索引越界")

# 安全访问
def safe_get(arr, index, default=None):
    try:
        return arr[index]
    except IndexError:
        return default

# 或使用切片 (不抛异常)
arr[100:101]              # [] (空列表)

# ==================== set ====================
s = {1, 2, 3}
s.add(4)
s.remove(2)               # 不存在抛 KeyError
s.discard(2)              # 不存在不抛异常
s.pop()                   # 删除任意元素
2 in s                    # True
len(s)                    # 3
s.clear()

# 集合运算
a = {1, 2, 3}
b = {2, 3, 4}
a | b                     # 并集 {1, 2, 3, 4}
a & b                     # 交集 {2, 3}
a - b                     # 差集 {1}
a ^ b                     # 对称差集 {1, 4}

# ==================== tuple (不可变) ====================
t = (1, 2, 3)
t = 1, 2, 3               # 括号可省略
single = (1,)             # 单元素元组
a, b, c = t               # 解构
a, *rest = t              # 剩余元素
```

### Go 切片操作

```go
// ==================== 创建 ====================
arr := []int{1, 2, 3}
empty := []string{}
filled := make([]int, 5)       // [0, 0, 0, 0, 0]
withCap := make([]int, 0, 10)  // 长度 0，容量 10

// ==================== 增 (Create) ====================
arr = append(arr, 4)           // 末尾添加
arr = append(arr, 5, 6, 7)     // 添加多个
arr = append([]int{0}, arr...) // 开头添加
// 中间插入
arr = append(arr[:2], append([]int{10}, arr[2:]...)...)

// ==================== 删 (Delete) ====================
arr = arr[:len(arr)-1]         // 删除末尾
arr = arr[1:]                  // 删除开头
// 删除索引 i
arr = append(arr[:i], arr[i+1:]...)

// 删除多个元素 (按值)
toRemove := map[int]bool{1: true, 3: true, 5: true}
result := arr[:0]
for _, x := range arr {
    if !toRemove[x] {
        result = append(result, x)
    }
}

// 删除多个元素 (按索引)
indices := map[int]bool{0: true, 2: true, 4: true}
result := arr[:0]
for i, x := range arr {
    if !indices[i] {
        result = append(result, x)
    }
}

// 清空 (保留容量)
arr = arr[:0]

// ==================== 改 (Update) ====================
arr[0] = 10                    // 直接修改
copy(arr[1:], []int{20, 30})   // 批量替换

// ==================== 查 (Read) ====================
arr[0]                         // 索引访问
arr[1:3]                       // 切片 [1, 3)
arr[1:]                        // 从索引 1 到末尾
arr[:3]                        // 从开头到索引 3

// 查找元素
func indexOf[T comparable](slice []T, target T) int {
    for i, v := range slice {
        if v == target {
            return i
        }
    }
    return -1
}

func contains[T comparable](slice []T, target T) bool {
    return indexOf(slice, target) >= 0
}

// ==================== IndexError 处理 ====================
// Go 会 panic
defer func() {
    if r := recover(); r != nil {
        fmt.Println("索引越界:", r)
    }
}()
_ = arr[100]  // panic: index out of range

// 安全访问
func safeGet[T any](slice []T, index int) (T, bool) {
    if index < 0 || index >= len(slice) {
        var zero T
        return zero, false
    }
    return slice[index], true
}

// ==================== 固定数组 ====================
var fixedArr [5]int           // [0, 0, 0, 0, 0]
fixedArr := [5]int{1, 2, 3}   // [1, 2, 3, 0, 0]
fixedArr := [...]int{1, 2, 3} // 自动推断长度

// ==================== Set (用 map 模拟) ====================
set := make(map[int]struct{})
set[1] = struct{}{}           // 添加
delete(set, 1)                // 删除
_, exists := set[1]           // 检查存在
```

### Rust 向量操作

```rust
// ==================== 创建 ====================
let arr: Vec<i32> = vec![1, 2, 3];
let empty: Vec<String> = Vec::new();
let filled = vec![0; 5];       // [0, 0, 0, 0, 0]
let range: Vec<i32> = (0..5).collect();

// ==================== 增 (Create) ====================
let mut arr = vec![1, 2, 3];
arr.push(4);                   // 末尾添加
arr.insert(0, 0);              // 指定位置插入
arr.extend([5, 6]);            // 扩展多个
arr.append(&mut other_vec);    // 合并

// ==================== 删 (Delete) ====================
arr.pop();                     // 删除末尾并返回 Option
arr.remove(0);                 // 删除指定索引 (会 panic)
arr.swap_remove(0);            // 删除并用最后一个填充 (O(1))
arr.retain(|x| *x != 2);       // 保留满足条件的

// 删除多个元素 (按值)
let to_remove: HashSet<i32> = [1, 3, 5].into_iter().collect();
arr.retain(|x| !to_remove.contains(x));

// 删除多个元素 (按索引，从后往前删)
let mut indices = vec![0, 2, 4];
indices.sort_by(|a, b| b.cmp(a));  // 倒序
for i in indices {
    arr.remove(i);
}

// 清空
arr.clear();

// ==================== 改 (Update) ====================
arr[0] = 10;                   // 直接修改
if let Some(elem) = arr.get_mut(0) {
    *elem = 20;
}

// ==================== 查 (Read) ====================
arr[0];                        // 索引访问 (可能 panic)
arr.get(0);                    // Option<&T>
arr.first();                   // 第一个 Option<&T>
arr.last();                    // 最后一个 Option<&T>
&arr[1..3];                    // 切片 [1, 3)
arr.iter().position(|x| *x == 2);  // 查找索引
arr.contains(&2);              // 是否包含
arr.iter().find(|x| **x > 2);  // 查找元素

// ==================== IndexError 处理 ====================
// 方式 1: get 返回 Option
match arr.get(100) {
    Some(value) => println!("{}", value),
    None => println!("索引越界"),
}

// 方式 2: get_or (不存在返回默认)
let value = arr.get(100).unwrap_or(&0);

// 方式 3: 直接索引 (会 panic)
// arr[100];  // panic: index out of bounds

// ==================== HashSet ====================
use std::collections::HashSet;

let mut set: HashSet<i32> = HashSet::from([1, 2, 3]);
set.insert(4);
set.remove(&2);
set.contains(&1);             // true
set.len();                    // 3

// 集合运算
let a: HashSet<i32> = [1, 2, 3].into_iter().collect();
let b: HashSet<i32> = [2, 3, 4].into_iter().collect();
let union: HashSet<_> = a.union(&b).collect();
let intersection: HashSet<_> = a.intersection(&b).collect();
let difference: HashSet<_> = a.difference(&b).collect();

// ==================== 固定数组 / Tuple ====================
let fixed: [i32; 5] = [1, 2, 3, 4, 5];
let filled: [i32; 5] = [0; 5];
let tuple: (i32, String, bool) = (1, String::from("hello"), true);
let (a, b, c) = tuple;        // 解构
```

---

## 🔪 Slice 切片操作

### 切片概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 语法 | `arr.slice(start, end)` | `list[start:end]` | `slice[start:end]` | `&slice[start..end]` |
| 负索引 | ❌ (需 at()) | ✅ | ❌ | ❌ |
| 步长 | ❌ | ✅ `[::step]` | ❌ | ❌ (需迭代器) |
| 原地修改 | `splice()` | `list[a:b] = x` | 赋值 | `copy_from_slice` |
| 返回类型 | 新数组 | 新列表 | 切片(引用) | 切片(引用) |

### TypeScript 切片操作

```typescript
// ==================== slice() 基本语法 ====================
const arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

// slice(start, end) - 不包含 end
arr.slice(2, 5);     // [2, 3, 4]
arr.slice(3);        // [3, 4, 5, 6, 7, 8, 9] - 从索引3到末尾
arr.slice();         // [0, 1, ..., 9] - 浅拷贝整个数组
arr.slice(0, -1);    // [0, 1, ..., 8] - 负数从末尾计算
arr.slice(-3);       // [7, 8, 9] - 最后3个元素
arr.slice(-5, -2);   // [5, 6, 7] - 倒数第5到倒数第2(不含)

// ==================== 字符串切片 ====================
const str = "Hello, World!";
str.slice(0, 5);     // "Hello"
str.slice(7);        // "World!"
str.slice(-6, -1);   // "World"

// substring vs slice
str.substring(0, 5); // "Hello" - 负数视为0，参数可交换
str.slice(0, 5);     // "Hello" - 支持负数，参数不交换

// ==================== splice() 原地修改 ====================
const nums = [1, 2, 3, 4, 5];

// splice(start, deleteCount, ...items)
nums.splice(2, 1);           // 删除索引2的元素，返回 [3]，nums = [1, 2, 4, 5]
nums.splice(1, 0, 10, 11);   // 在索引1插入，nums = [1, 10, 11, 2, 4, 5]
nums.splice(2, 2, 20);       // 替换，nums = [1, 10, 20, 4, 5]

// ==================== 访问单个元素 ====================
const items = ['a', 'b', 'c', 'd', 'e'];

items[0];            // 'a'
items[items.length - 1];  // 'e' - 最后一个
items.at(0);         // 'a' - ES2022
items.at(-1);        // 'e' - 支持负索引
items.at(-2);        // 'd'

// ==================== 分块切片 ====================
function chunk<T>(arr: T[], size: number): T[][] {
    const result: T[][] = [];
    for (let i = 0; i < arr.length; i += size) {
        result.push(arr.slice(i, i + size));
    }
    return result;
}

chunk([1, 2, 3, 4, 5], 2);  // [[1, 2], [3, 4], [5]]

// ==================== 滑动窗口 ====================
function* slidingWindow<T>(arr: T[], size: number): Generator<T[]> {
    for (let i = 0; i <= arr.length - size; i++) {
        yield arr.slice(i, i + size);
    }
}

[...slidingWindow([1, 2, 3, 4, 5], 3)];  // [[1,2,3], [2,3,4], [3,4,5]]

// ==================== 头尾操作 ====================
const data = [1, 2, 3, 4, 5];

// 获取头部/尾部
const [first, ...rest] = data;     // first=1, rest=[2,3,4,5]
const [head, second] = data;       // head=1, second=2
const last = data.at(-1);          // 5

// 去除头部/尾部
data.slice(1);       // [2, 3, 4, 5] - 去除第一个
data.slice(0, -1);   // [1, 2, 3, 4] - 去除最后一个
data.slice(1, -1);   // [2, 3, 4] - 去除首尾

// ==================== TypedArray 切片 ====================
const buffer = new ArrayBuffer(16);
const int32View = new Int32Array(buffer);
int32View.set([1, 2, 3, 4]);

// subarray 返回视图(共享内存)
const sub = int32View.subarray(1, 3);  // Int32Array [2, 3]
sub[0] = 100;  // 修改会影响原数组

// slice 返回拷贝
const copy = int32View.slice(1, 3);    // Int32Array [2, 3]
copy[0] = 100;  // 不影响原数组
```

### Python 切片操作

```python
# ==================== 基本切片语法 ====================
lst = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9]

# list[start:stop:step]
lst[2:5]        # [2, 3, 4] - 不包含 stop
lst[3:]         # [3, 4, 5, 6, 7, 8, 9] - 从索引3到末尾
lst[:5]         # [0, 1, 2, 3, 4] - 从开头到索引5(不含)
lst[:]          # [0, 1, ..., 9] - 浅拷贝
lst[::2]        # [0, 2, 4, 6, 8] - 步长为2
lst[1::2]       # [1, 3, 5, 7, 9] - 从索引1开始，步长为2

# ==================== 负索引 ====================
lst[-1]         # 9 - 最后一个
lst[-3:]        # [7, 8, 9] - 最后3个
lst[:-3]        # [0, 1, 2, 3, 4, 5, 6] - 除了最后3个
lst[-5:-2]      # [5, 6, 7] - 倒数第5到倒数第2(不含)
lst[::-1]       # [9, 8, ..., 0] - 反转列表
lst[::-2]       # [9, 7, 5, 3, 1] - 反向步长2

# ==================== 切片对象 ====================
s = slice(2, 7, 2)
lst[s]          # [2, 4, 6] - 等同于 lst[2:7:2]

# 获取切片的实际索引
s.indices(len(lst))  # (2, 7, 2) - (start, stop, step)

# ==================== 字符串切片 ====================
text = "Hello, World!"
text[0:5]       # "Hello"
text[7:]        # "World!"
text[::-1]      # "!dlroW ,olleH" - 反转
text[::2]       # "Hlo ol!"

# ==================== 切片赋值 (原地修改) ====================
nums = [1, 2, 3, 4, 5]

# 替换切片
nums[1:4] = [20, 30]    # [1, 20, 30, 5] - 可以不等长
nums[1:1] = [10, 11]    # [1, 10, 11, 20, 30, 5] - 插入

# 删除切片
nums[2:4] = []          # [1, 10, 30, 5]
del nums[1:3]           # [1, 5]

# 步长切片赋值 (必须等长)
lst = [0, 1, 2, 3, 4, 5]
lst[::2] = [10, 20, 30]  # [10, 1, 20, 3, 30, 5]

# ==================== 元组切片 ====================
t = (0, 1, 2, 3, 4)
t[1:4]          # (1, 2, 3) - 返回新元组
t[::-1]         # (4, 3, 2, 1, 0)
# t[1:3] = (10, 20)  # 错误! 元组不可变

# ==================== numpy 数组切片 ====================
import numpy as np

arr = np.array([[1, 2, 3], [4, 5, 6], [7, 8, 9]])

arr[0]          # array([1, 2, 3]) - 第一行
arr[:, 0]       # array([1, 4, 7]) - 第一列
arr[0:2, 1:3]   # array([[2, 3], [5, 6]]) - 子矩阵
arr[::2, ::2]   # array([[1, 3], [7, 9]]) - 间隔取值

# numpy 切片是视图
view = arr[0:2, 0:2]
view[0, 0] = 100  # 修改会影响原数组

# 拷贝
copy = arr[0:2, 0:2].copy()

# ==================== 高级切片技巧 ====================
# 分块
def chunk(lst, size):
    return [lst[i:i+size] for i in range(0, len(lst), size)]

chunk([1, 2, 3, 4, 5], 2)  # [[1, 2], [3, 4], [5]]

# 滑动窗口
def sliding_window(lst, size):
    return [lst[i:i+size] for i in range(len(lst) - size + 1)]

sliding_window([1, 2, 3, 4, 5], 3)  # [[1,2,3], [2,3,4], [3,4,5]]

# 旋转列表
def rotate(lst, n):
    n = n % len(lst)
    return lst[n:] + lst[:n]

rotate([1, 2, 3, 4, 5], 2)  # [3, 4, 5, 1, 2]

# ==================== memoryview 切片 ====================
data = bytearray(b'Hello World')
view = memoryview(data)

# 切片共享内存
sub = view[0:5]
sub[0] = ord('h')  # data 变为 b'hello World'

# 转换为 bytes
bytes(view[6:11])  # b'World'
```

### Go 切片操作

```go
// ==================== 切片基础 ====================
// 切片是对底层数组的引用视图
arr := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

// 从数组创建切片
slice := arr[2:5]    // [2 3 4] - 不包含索引5
slice2 := arr[3:]    // [3 4 5 6 7 8 9]
slice3 := arr[:5]    // [0 1 2 3 4]
slice4 := arr[:]     // [0 1 2 3 4 5 6 7 8 9]

// 直接创建切片
nums := []int{1, 2, 3, 4, 5}
empty := []int{}
withCap := make([]int, 5, 10)  // len=5, cap=10

// ==================== 切片属性 ====================
s := []int{1, 2, 3, 4, 5}
len(s)           // 5 - 长度
cap(s)           // 5 - 容量 (到底层数组末尾的长度)

// 切片的切片
sub := s[1:4]    // [2 3 4]
len(sub)         // 3
cap(sub)         // 4 - 从索引1到原切片末尾

// ==================== 切片是引用 ====================
original := []int{1, 2, 3, 4, 5}
slice := original[1:4]
slice[0] = 100   // original 变为 [1 100 3 4 5]

// 创建独立副本
copySlice := make([]int, len(original))
copy(copySlice, original)

// 或使用 append
copySlice2 := append([]int{}, original...)

// ==================== append 操作 ====================
s := []int{1, 2, 3}

// 追加元素
s = append(s, 4)           // [1 2 3 4]
s = append(s, 5, 6, 7)     // [1 2 3 4 5 6 7]

// 追加切片
other := []int{8, 9}
s = append(s, other...)    // [1 2 3 4 5 6 7 8 9]

// 头部插入
s = append([]int{0}, s...) // [0 1 2 3 ...]

// 中间插入
idx := 3
s = append(s[:idx], append([]int{100}, s[idx:]...)...)

// ==================== 删除元素 ====================
s := []int{1, 2, 3, 4, 5}

// 删除索引 i
i := 2
s = append(s[:i], s[i+1:]...)  // [1 2 4 5]

// 删除范围 [i, j)
s = append(s[:i], s[j:]...)

// 保持顺序删除最后一个
s = s[:len(s)-1]

// 不保持顺序删除 (O(1))
s[i] = s[len(s)-1]
s = s[:len(s)-1]

// ==================== 切片技巧 ====================
// 完整切片表达式 slice[low:high:max]
arr := [5]int{1, 2, 3, 4, 5}
s := arr[1:3:4]   // [2 3], cap=3 (4-1)

// 限制容量防止意外修改
original := []int{1, 2, 3, 4, 5}
limited := original[1:3:3]  // cap=2，append 会创建新数组

// ==================== 二维切片 ====================
// 创建 3x4 矩阵
matrix := make([][]int, 3)
for i := range matrix {
    matrix[i] = make([]int, 4)
}

// 初始化
matrix := [][]int{
    {1, 2, 3},
    {4, 5, 6},
    {7, 8, 9},
}

// 访问
matrix[0]        // [1 2 3]
matrix[0][1]     // 2

// 行切片
row := matrix[1][:]   // [4 5 6]

// 列切片 (需要循环)
col := make([]int, len(matrix))
for i := range matrix {
    col[i] = matrix[i][1]
}

// ==================== 字符串切片 ====================
s := "Hello, 世界"

// 字节切片
s[0:5]           // "Hello"
s[7:]            // "世界"

// 注意: 中文字符占3个字节
s[7:10]          // "世" - 正确
// s[7:8]        // 乱码 - 错误

// 安全的字符切片
runes := []rune(s)
string(runes[7:9])  // "世界"

// ==================== 分块与滑动窗口 ====================
func chunk[T any](slice []T, size int) [][]T {
    var result [][]T
    for i := 0; i < len(slice); i += size {
        end := i + size
        if end > len(slice) {
            end = len(slice)
        }
        result = append(result, slice[i:end])
    }
    return result
}

func slidingWindow[T any](slice []T, size int) [][]T {
    if len(slice) < size {
        return nil
    }
    result := make([][]T, 0, len(slice)-size+1)
    for i := 0; i <= len(slice)-size; i++ {
        result = append(result, slice[i:i+size])
    }
    return result
}

// ==================== 预分配优化 ====================
// 已知大小时预分配
result := make([]int, 0, expectedSize)
for _, v := range data {
    result = append(result, process(v))
}

// 清空但保留容量
s = s[:0]
```

### Rust 切片操作

```rust
// ==================== 切片基础 ====================
// 切片是对连续序列的引用视图
let arr = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9];

// 创建切片 (引用)
let slice: &[i32] = &arr[2..5];    // [2, 3, 4]
let slice2 = &arr[3..];           // [3, 4, 5, 6, 7, 8, 9]
let slice3 = &arr[..5];           // [0, 1, 2, 3, 4]
let slice4 = &arr[..];            // 全部

// 包含结束索引
let inclusive = &arr[2..=5];      // [2, 3, 4, 5]

// Vec 切片
let vec = vec![1, 2, 3, 4, 5];
let vec_slice = &vec[1..4];       // [2, 3, 4]

// ==================== 可变切片 ====================
let mut arr = [1, 2, 3, 4, 5];
let slice = &mut arr[1..4];
slice[0] = 100;                   // arr 变为 [1, 100, 3, 4, 5]

// ==================== 切片方法 ====================
let s = &[1, 2, 3, 4, 5][..];

// 长度
s.len();                          // 5
s.is_empty();                     // false

// 访问元素
s[0];                             // 1 - 可能 panic
s.get(0);                         // Some(&1)
s.get(10);                        // None
s.first();                        // Some(&1)
s.last();                         // Some(&5)

// 安全访问
if let Some(val) = s.get(2) {
    println!("{}", val);
}

// ==================== 分割切片 ====================
let s = &[1, 2, 3, 4, 5][..];

// split_at
let (left, right) = s.split_at(2);  // [1, 2], [3, 4, 5]

// split_first / split_last
let (first, rest) = s.split_first().unwrap();  // 1, [2, 3, 4, 5]
let (last, init) = s.split_last().unwrap();    // 5, [1, 2, 3, 4]

// split 按条件
let parts: Vec<_> = s.split(|&x| x == 3).collect();
// [[1, 2], [4, 5]]

// splitn 限制数量
let parts: Vec<_> = s.splitn(2, |&x| x == 3).collect();

// ==================== 窗口与分块 ====================
let s = &[1, 2, 3, 4, 5][..];

// windows - 滑动窗口
for window in s.windows(3) {
    println!("{:?}", window);     // [1,2,3], [2,3,4], [3,4,5]
}

// chunks - 分块
for chunk in s.chunks(2) {
    println!("{:?}", chunk);      // [1,2], [3,4], [5]
}

// chunks_exact - 精确分块 (忽略不足)
for chunk in s.chunks_exact(2) {
    println!("{:?}", chunk);      // [1,2], [3,4]
}
let remainder = s.chunks_exact(2).remainder();  // [5]

// rchunks - 从右侧分块
for chunk in s.rchunks(2) {
    println!("{:?}", chunk);      // [4,5], [2,3], [1]
}

// ==================== 可变切片操作 ====================
let mut arr = [5, 2, 8, 1, 9, 3];
let s = &mut arr[..];

// 排序
s.sort();                         // [1, 2, 3, 5, 8, 9]
s.sort_by(|a, b| b.cmp(a));       // 降序
s.sort_by_key(|x| -x);            // 按key

// 反转
s.reverse();                      // [9, 8, 5, 3, 2, 1]

// 旋转
s.rotate_left(2);                 // [5, 3, 2, 1, 9, 8]
s.rotate_right(2);                // [9, 8, 5, 3, 2, 1]

// 交换
s.swap(0, 5);                     // 交换索引0和5

// 填充
s.fill(0);                        // 全部填充为0

// ==================== 拷贝操作 ====================
let src = [1, 2, 3];
let mut dst = [0; 5];

// copy_from_slice (长度必须相等)
dst[..3].copy_from_slice(&src);   // [1, 2, 3, 0, 0]

// clone_from_slice (Clone 类型)
dst[..3].clone_from_slice(&src);

// 部分拷贝
let mut vec = vec![1, 2, 3, 4, 5];
vec.copy_within(1..4, 0);         // [2, 3, 4, 4, 5]

// ==================== 搜索 ====================
let s = &[1, 2, 3, 4, 5, 3][..];

// 查找
s.contains(&3);                   // true
s.starts_with(&[1, 2]);           // true
s.ends_with(&[5, 3]);             // true

// 位置
s.iter().position(|&x| x == 3);   // Some(2)
s.iter().rposition(|&x| x == 3);  // Some(5)

// 二分查找 (已排序)
let sorted = &[1, 2, 3, 4, 5][..];
sorted.binary_search(&3);         // Ok(2)
sorted.binary_search(&6);         // Err(5) - 插入位置

// ==================== 字符串切片 ====================
let s = "Hello, 世界";

// 字节切片
let bytes: &[u8] = s.as_bytes();

// 字符串切片 (必须是有效 UTF-8 边界)
let hello = &s[0..5];             // "Hello"
// let bad = &s[0..8];            // panic! 无效边界

// 安全切片
let slice = s.get(0..5);          // Some("Hello")
let invalid = s.get(0..8);        // None

// 字符迭代
for (i, c) in s.char_indices() {
    println!("{}: {}", i, c);
}

// ==================== 迭代器转切片 ====================
let vec: Vec<i32> = (1..=5).collect();
let slice: &[i32] = &vec;

// 切片转 Vec
let vec2: Vec<i32> = slice.to_vec();

// 数组转切片
let arr = [1, 2, 3, 4, 5];
let slice: &[i32] = &arr;

// Box<[T]> - 堆上固定大小
let boxed: Box<[i32]> = vec![1, 2, 3].into_boxed_slice();
```

### 切片操作对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 操作            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 基本切片        │ arr.slice(1, 4)      │ lst[1:4]             │ s[1:4]               │ &s[1..4]             │
│ 从开头          │ arr.slice(0, n)      │ lst[:n]              │ s[:n]                │ &s[..n]              │
│ 到末尾          │ arr.slice(n)         │ lst[n:]              │ s[n:]                │ &s[n..]              │
│ 负索引          │ arr.at(-1)           │ lst[-1]              │ s[len(s)-1]          │ s.last()             │
│ 步长            │ ❌                   │ lst[::2]             │ ❌                   │ iter().step_by(2)    │
│ 反转            │ arr.reverse()        │ lst[::-1]            │ slices.Reverse()     │ s.reverse()          │
│ 复制            │ [...arr]             │ lst[:]               │ copy(dst, src)       │ s.to_vec()           │
│ 修改是否影响原  │ ❌ (新数组)          │ ❌ (新列表)          │ ✅ (视图)            │ ✅ (引用)            │
│ 分块            │ 手动实现             │ 手动实现             │ 手动实现             │ chunks()             │
│ 滑动窗口        │ 手动实现             │ 手动实现             │ 手动实现             │ windows()            │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 常见切片模式

| 模式 | 描述 | 示例 |
|------|------|------|
| **浅拷贝** | 创建独立副本 | `[...arr]` / `lst[:]` / `copy()` / `.to_vec()` |
| **头尾分离** | 获取首/尾元素和剩余 | 解构 / `split_first` |
| **分块处理** | 固定大小分组 | `chunks()` / 手动循环 |
| **滑动窗口** | 连续子序列 | `windows()` / 手动循环 |
| **限制容量** | 防止意外扩展 | Go `s[a:b:c]` |
| **安全访问** | 避免越界 | `.get()` / `try` |

---

## 🗺️ Map / Struct / Interface

### 映射类型概览

| 类型 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 字典/映射 | `Map<K,V>` / `{}` | `dict` | `map[K]V` | `HashMap<K,V>` |
| 结构体 | `interface` / `class` | `class` / `dataclass` | `struct` | `struct` |
| 接口 | `interface` | `Protocol` | `interface` | `trait` |

### TypeScript Map/Object 操作

```typescript
// ==================== Object ====================
const obj: Record<string, number> = { a: 1, b: 2 };

// 增
obj.c = 3;
obj["d"] = 4;

// 删
delete obj.a;
const { b, ...rest } = obj;  // 不可变删除

// 改
obj.b = 20;

// 查
obj.a;                     // 可能 undefined
obj["a"];
"a" in obj;                // 检查 key
Object.keys(obj);
Object.values(obj);
Object.entries(obj);

// ==================== Map (推荐) ====================
const map = new Map<string, number>();

// 增
map.set("a", 1);
map.set("b", 2);

// 删
map.delete("a");
map.clear();

// 改
map.set("b", 20);

// 查
map.get("a");              // undefined if not exists
map.has("a");              // boolean
map.size;
map.keys();
map.values();
map.entries();

// ==================== KeyError 处理 ====================
// Map.get 返回 undefined，不抛异常
const value = map.get("nonexistent");  // undefined

// 带默认值
function getOrDefault<K, V>(map: Map<K, V>, key: K, defaultValue: V): V {
  return map.has(key) ? map.get(key)! : defaultValue;
}

// 带初始化
function getOrSet<K, V>(map: Map<K, V>, key: K, factory: () => V): V {
  if (!map.has(key)) {
    map.set(key, factory());
  }
  return map.get(key)!;
}

// ==================== 删除多个 key ====================
const keysToRemove = ["a", "c", "e"];
keysToRemove.forEach(key => map.delete(key));

// 不可变删除
const newMap = new Map([...map].filter(([k]) => !keysToRemove.includes(k)));

// ==================== interface / class ====================
interface User {
  id: number;
  name: string;
  email?: string;
}

class UserImpl implements User {
  constructor(
    public id: number,
    public name: string,
    public email?: string
  ) {}
}

// 类型守卫
function isUser(obj: any): obj is User {
  return typeof obj.id === "number" && typeof obj.name === "string";
}
```

### Python 字典操作

```python
# ==================== dict ====================
d = {"a": 1, "b": 2}
d = dict(a=1, b=2)

# 增
d["c"] = 3
d.update({"d": 4, "e": 5})
d |= {"f": 6}              # Python 3.9+

# 删
del d["a"]
d.pop("b")                 # 删除并返回
d.pop("x", None)           # 不存在返回默认值
d.popitem()                # 删除最后插入的

# 删除多个 key
keys_to_remove = ["a", "c", "e"]
for key in keys_to_remove:
    d.pop(key, None)

# 不可变删除
new_d = {k: v for k, v in d.items() if k not in keys_to_remove}

# 清空
d.clear()

# 改
d["a"] = 10

# 查
d["a"]                     # KeyError if not exists
d.get("a")                 # None if not exists
d.get("a", 0)              # 默认值
"a" in d                   # 检查 key
d.keys()
d.values()
d.items()

# ==================== KeyError 处理 ====================
try:
    value = d["nonexistent"]
except KeyError:
    print("Key 不存在")

# 使用 get
value = d.get("nonexistent", "default")

# setdefault (获取或设置默认值)
value = d.setdefault("key", [])
value.append(1)            # d["key"] 现在是 [1]

# defaultdict
from collections import defaultdict
dd = defaultdict(list)
dd["key"].append(1)        # 自动初始化为 []

# ==================== class / dataclass ====================
from dataclasses import dataclass
from typing import Optional

@dataclass
class User:
    id: int
    name: str
    email: Optional[str] = None

user = User(id=1, name="John")
user.name                  # 访问属性

# 普通 class
class User:
    def __init__(self, id: int, name: str):
        self.id = id
        self.name = name

# Protocol (结构化类型)
from typing import Protocol

class HasName(Protocol):
    name: str

def greet(obj: HasName) -> str:
    return f"Hello, {obj.name}"
```

### Go Map/Struct 操作

```go
// ==================== map ====================
m := map[string]int{"a": 1, "b": 2}
m := make(map[string]int)

// 增
m["c"] = 3

// 删
delete(m, "a")

// 删除多个 key
keysToRemove := []string{"a", "c", "e"}
for _, key := range keysToRemove {
    delete(m, key)
}

// 清空 (重新创建)
m = make(map[string]int)

// 改
m["a"] = 10

// 查
value := m["a"]            // 不存在返回零值
value, ok := m["a"]        // comma ok 模式
if ok {
    fmt.Println(value)
}

// 遍历
for key, value := range m {
    fmt.Println(key, value)
}

// ==================== KeyError 处理 ====================
// Go 不抛异常，返回零值
value := m["nonexistent"]  // 0 (int 的零值)

// 检查是否存在
value, exists := m["nonexistent"]
if !exists {
    fmt.Println("Key 不存在")
}

// 带默认值
func getOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
    if value, ok := m[key]; ok {
        return value
    }
    return defaultValue
}

// ==================== struct ====================
type User struct {
    ID    int
    Name  string
    Email string  // 可选用指针 *string
}

// 创建
user := User{ID: 1, Name: "John"}
user := User{
    ID:   1,
    Name: "John",
}
userPtr := &User{ID: 1, Name: "John"}

// 访问
user.Name
userPtr.Name               // 自动解引用

// 修改
user.Name = "Jane"

// ==================== interface ====================
type HasName interface {
    GetName() string
}

func (u User) GetName() string {
    return u.Name
}

func greet(obj HasName) string {
    return "Hello, " + obj.GetName()
}

// 类型断言
if user, ok := obj.(User); ok {
    fmt.Println(user.ID)
}

// 类型 switch
switch v := obj.(type) {
case User:
    fmt.Println(v.ID)
case *User:
    fmt.Println(v.ID)
default:
    fmt.Println("Unknown type")
}
```

### Rust HashMap/Struct 操作

```rust
use std::collections::HashMap;

// ==================== HashMap ====================
let mut map: HashMap<String, i32> = HashMap::new();
let map: HashMap<_, _> = [("a", 1), ("b", 2)].into_iter().collect();

// 增
map.insert("c".to_string(), 3);

// 删
map.remove("a");

// 删除多个 key
let keys_to_remove = vec!["a", "c", "e"];
for key in keys_to_remove {
    map.remove(key);
}
// 或使用 retain
map.retain(|k, _| !keys_to_remove.contains(&k.as_str()));

// 清空
map.clear();

// 改
map.insert("a".to_string(), 10);
if let Some(value) = map.get_mut("a") {
    *value = 20;
}

// 查
map.get("a");              // Option<&V>
map.get("a").copied();     // Option<V> (if V: Copy)
map.contains_key("a");
map.keys();
map.values();
map.iter();

// ==================== KeyError 处理 ====================
// get 返回 Option
match map.get("nonexistent") {
    Some(value) => println!("{}", value),
    None => println!("Key 不存在"),
}

// unwrap_or
let value = map.get("nonexistent").unwrap_or(&0);

// entry API (获取或插入)
let value = map.entry("key".to_string()).or_insert(0);
*value += 1;

// entry 带闭包
map.entry("key".to_string()).or_insert_with(|| expensive_computation());

// ==================== struct ====================
struct User {
    id: u32,
    name: String,
    email: Option<String>,
}

// 创建
let user = User {
    id: 1,
    name: String::from("John"),
    email: None,
};

// 使用 Default
#[derive(Default)]
struct User {
    id: u32,
    name: String,
    email: Option<String>,
}

let user = User {
    name: String::from("John"),
    ..Default::default()
};

// 访问
user.name;

// 修改 (需要 mut)
let mut user = User { /* ... */ };
user.name = String::from("Jane");

// ==================== trait (接口) ====================
trait HasName {
    fn get_name(&self) -> &str;
}

impl HasName for User {
    fn get_name(&self) -> &str {
        &self.name
    }
}

fn greet(obj: &impl HasName) -> String {
    format!("Hello, {}", obj.get_name())
}

// 或使用 dyn trait
fn greet_dyn(obj: &dyn HasName) -> String {
    format!("Hello, {}", obj.get_name())
}

// ==================== 解构 ====================
let User { id, name, .. } = user;

// match 解构
match user {
    User { id: 0, .. } => println!("Guest"),
    User { name, .. } => println!("User: {}", name),
}
```

### 集合操作对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 操作            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 数组追加        │ push()         │ append()       │ append()       │ push()         │
│ 数组删除        │ splice()       │ remove()/pop() │ append切片     │ remove()       │
│ 数组查找        │ indexOf()      │ index()        │ 手动循环       │ position()     │
│ 数组包含        │ includes()     │ in             │ 手动循环       │ contains()     │
│ 索引越界        │ undefined      │ IndexError     │ panic          │ panic/Option   │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ Map 设置        │ set()          │ d[k]=v         │ m[k]=v         │ insert()       │
│ Map 获取        │ get()          │ get()/d[k]     │ m[k]/comma-ok  │ get()          │
│ Map 删除        │ delete()       │ del/pop()      │ delete()       │ remove()       │
│ Key 不存在      │ undefined      │ KeyError/None  │ 零值/comma-ok  │ Option         │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

---

## 🔗 指针与引用

### 指针/引用概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 指针类型 | ❌ | ❌ | `*T` | `*const T` / `*mut T` |
| 引用类型 | 对象自动引用 | 对象自动引用 | `&T` (取地址) | `&T` / `&mut T` |
| 空指针 | `null`/`undefined` | `None` | `nil` | `Option<&T>` |
| 解引用 | 自动 | 自动 | `*ptr` | `*ptr` |
| 智能指针 | ❌ | ❌ | ❌ | `Box`/`Rc`/`Arc` |
| 裸指针 | ❌ | `ctypes` | `unsafe.Pointer` | `*const T`/`*mut T` |

### 指针与引用图解

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                         指针与引用的区别                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  引用 (Reference)                    指针 (Pointer)                         │
│  ┌─────────┐                        ┌─────────┐                             │
│  │ 变量 x  │                        │ 变量 x  │                             │
│  │  42     │                        │  42     │                             │
│  └────▲────┘                        └────▲────┘                             │
│       │                                  │                                  │
│       │ 别名                             │ 地址: 0x1234                      │
│       │                                  │                                  │
│  ┌────┴────┐                        ┌────┴────┐                             │
│  │ 引用 r  │                        │ 指针 p  │                             │
│  │  &x     │                        │ 0x1234  │ ← 存储的是地址              │
│  └─────────┘                        └─────────┘                             │
│  • 必须有效                          • 可以为 null/nil                       │
│  • 自动解引用                        • 需要显式解引用 *p                     │
│  • 编译器保证安全                    • 可能悬垂(dangling)                   │
│                                                                             │
│  智能指针 (Rust)                                                            │
│  ┌─────────────────────────────────────────────────┐                       │
│  │ Box<T>     - 堆上分配，独占所有权               │                       │
│  │ Rc<T>      - 引用计数，单线程共享               │                       │
│  │ Arc<T>     - 原子引用计数，多线程共享           │                       │
│  │ RefCell<T> - 运行时借用检查                     │                       │
│  └─────────────────────────────────────────────────┘                       │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### TypeScript 引用

```typescript
// ==================== 对象引用 ====================
// TypeScript 中对象、数组、函数都是引用类型
const obj1 = { value: 42 };
const obj2 = obj1;  // obj2 引用同一个对象

obj2.value = 100;
console.log(obj1.value);  // 100 - 同一对象

// 检查是否同一引用
console.log(obj1 === obj2);  // true

// ==================== 基本类型是值 ====================
let a = 42;
let b = a;  // 复制值
b = 100;
console.log(a);  // 42 - 不受影响

// ==================== 创建新引用 (浅拷贝) ====================
const original = { a: 1, nested: { b: 2 } };

// 浅拷贝方法
const copy1 = { ...original };
const copy2 = Object.assign({}, original);

copy1.a = 100;
console.log(original.a);  // 1 - 不受影响

copy1.nested.b = 200;
console.log(original.nested.b);  // 200 - 嵌套对象仍共享！

// ==================== 深拷贝 ====================
const deep1 = JSON.parse(JSON.stringify(original));
const deep2 = structuredClone(original);  // 现代浏览器

// ==================== WeakRef 弱引用 ====================
let target = { data: "important" };
const weakRef = new WeakRef(target);

// 获取引用（可能已被 GC）
const deref = weakRef.deref();
if (deref) {
    console.log(deref.data);
}

// ==================== 引用相等 vs 值相等 ====================
const arr1 = [1, 2, 3];
const arr2 = [1, 2, 3];
const arr3 = arr1;

arr1 === arr2;  // false - 不同对象
arr1 === arr3;  // true - 同一引用

// 值比较
JSON.stringify(arr1) === JSON.stringify(arr2);  // true

// ==================== 冻结对象防止修改 ====================
const frozen = Object.freeze({ value: 42 });
// frozen.value = 100;  // 严格模式下报错

// 深冻结
function deepFreeze<T extends object>(obj: T): Readonly<T> {
    Object.keys(obj).forEach(key => {
        const value = (obj as any)[key];
        if (typeof value === 'object' && value !== null) {
            deepFreeze(value);
        }
    });
    return Object.freeze(obj);
}
```

### Python 引用

```python
# ==================== 对象引用 ====================
# Python 中一切皆对象，变量是对象的引用/标签
list1 = [1, 2, 3]
list2 = list1  # 指向同一个列表

list2.append(4)
print(list1)  # [1, 2, 3, 4] - 同一对象

# 检查是否同一对象
print(list1 is list2)  # True
print(id(list1) == id(list2))  # True

# ==================== 不可变对象 ====================
# int, str, tuple 是不可变的
a = 42
b = a
b = 100  # 创建新对象
print(a)  # 42 - 不受影响

# 小整数缓存 (-5 到 256)
x = 100
y = 100
print(x is y)  # True - 缓存的同一对象

# 字符串驻留
s1 = "hello"
s2 = "hello"
print(s1 is s2)  # True - 驻留的同一对象

# ==================== 浅拷贝 ====================
import copy

original = [1, 2, [3, 4]]

# 浅拷贝方法
copy1 = original.copy()
copy2 = list(original)
copy3 = original[:]
copy4 = copy.copy(original)

copy1[0] = 100
print(original[0])  # 1 - 不受影响

copy1[2][0] = 300
print(original[2][0])  # 300 - 嵌套对象共享！

# ==================== 深拷贝 ====================
deep = copy.deepcopy(original)
deep[2][0] = 999
print(original[2][0])  # 300 - 完全独立

# ==================== 弱引用 ====================
import weakref

class MyClass:
    pass

obj = MyClass()
weak = weakref.ref(obj)

print(weak())  # <MyClass object>
del obj
print(weak())  # None - 对象已被回收

# WeakValueDictionary
cache = weakref.WeakValueDictionary()
obj = MyClass()
cache['key'] = obj

# ==================== 引用计数 ====================
import sys

a = [1, 2, 3]
print(sys.getrefcount(a))  # 2 (a + getrefcount 参数)

b = a
print(sys.getrefcount(a))  # 3

del b
print(sys.getrefcount(a))  # 2

# ==================== 可变参数陷阱 ====================
# 错误！默认列表在所有调用间共享
def bad(items=[]):
    items.append(1)
    return items

print(bad())  # [1]
print(bad())  # [1, 1] - 不是 [1]！

# 正确做法
def good(items=None):
    if items is None:
        items = []
    items.append(1)
    return items

# ==================== __slots__ 优化内存 ====================
class WithSlots:
    __slots__ = ['x', 'y']  # 不使用 __dict__
    
    def __init__(self, x, y):
        self.x = x
        self.y = y

# 比普通类省内存，但不能动态添加属性
```

### Go 指针

```go
// ==================== 指针基础 ====================
var x int = 42
var p *int = &x  // p 是指向 x 的指针

fmt.Println(p)   // 0xc0000b4008 (地址)
fmt.Println(*p)  // 42 (解引用)

*p = 100         // 通过指针修改值
fmt.Println(x)   // 100

// ==================== 零值是 nil ====================
var ptr *int     // nil
if ptr == nil {
    fmt.Println("ptr is nil")
}

// 解引用 nil 会 panic
// fmt.Println(*ptr)  // panic!

// ==================== new 和 & ====================
// new 返回指针，值为零值
p1 := new(int)    // *int, 值为 0
*p1 = 42

// & 取地址
val := 42
p2 := &val

// 字面量取地址
p3 := &struct{ x int }{x: 42}

// ==================== 指针作为参数 ====================
func double(x *int) {
    *x *= 2
}

num := 10
double(&num)
fmt.Println(num)  // 20

// ==================== 结构体指针 ====================
type User struct {
    Name string
    Age  int
}

// 自动解引用
func (u *User) Birthday() {
    u.Age++  // 等同于 (*u).Age++
}

user := &User{Name: "Alice", Age: 30}
user.Birthday()
fmt.Println(user.Age)  // 31

// ==================== 指针数组 vs 数组指针 ====================
// 指针数组: 元素是指针
var ptrArr [3]*int

// 数组指针: 指向数组的指针
arr := [3]int{1, 2, 3}
var arrPtr *[3]int = &arr
arrPtr[0] = 100  // 自动解引用

// ==================== 不能获取的地址 ====================
// 常量没有地址
// const c = 42
// p := &c  // 编译错误

// map 的值没有地址
m := map[string]int{"a": 1}
// p := &m["a"]  // 编译错误

// ==================== unsafe.Pointer ====================
import "unsafe"

// 任意指针类型转换
var i int64 = 42
ptr := unsafe.Pointer(&i)
floatPtr := (*float64)(ptr)

// 指针运算
arr := [3]int{10, 20, 30}
p := unsafe.Pointer(&arr[0])
p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(arr[0]))
fmt.Println(*(*int)(p))  // 20

// ==================== 返回局部变量指针 (安全) ====================
func createUser() *User {
    u := User{Name: "Bob"}  // 逃逸到堆
    return &u  // 安全！Go 会处理
}

// ==================== 指针接收者 vs 值接收者 ====================
type Counter struct {
    count int
}

// 值接收者 - 复制
func (c Counter) ValueMethod() {
    c.count++  // 修改副本
}

// 指针接收者 - 原值
func (c *Counter) PointerMethod() {
    c.count++  // 修改原值
}

c := Counter{count: 0}
c.ValueMethod()
fmt.Println(c.count)  // 0

c.PointerMethod()
fmt.Println(c.count)  // 1

// ==================== 何时使用指针 ====================
/*
使用指针:
1. 需要修改参数值
2. 大结构体避免复制
3. 方法需要修改接收者
4. 表示可选值 (nil)

使用值:
1. 小数据类型 (int, bool)
2. 不需要修改
3. 需要复制语义
4. 并发安全的不可变数据
*/
```

### Rust 引用与指针

```rust
// ==================== 引用 (安全) ====================
// 不可变引用 &T
let x = 42;
let r: &i32 = &x;
println!("{}", *r);  // 42

// 可变引用 &mut T
let mut y = 42;
let r_mut: &mut i32 = &mut y;
*r_mut = 100;
println!("{}", y);  // 100

// ==================== 借用规则 ====================
let mut s = String::from("hello");

// 规则1: 多个不可变引用 OK
let r1 = &s;
let r2 = &s;
println!("{} {}", r1, r2);

// 规则2: 一个可变引用，无其他引用
let r3 = &mut s;
// let r4 = &s;      // 编译错误！
// let r5 = &mut s;  // 编译错误！
r3.push_str(" world");

// ==================== 悬垂引用 (编译器阻止) ====================
// fn dangling() -> &String {
//     let s = String::from("hello");
//     &s  // 编译错误！s 将被释放
// }

// 正确: 返回所有权
fn not_dangling() -> String {
    let s = String::from("hello");
    s
}

// ==================== 裸指针 (unsafe) ====================
let x = 42;

// 创建裸指针 (安全)
let r1 = &x as *const i32;  // 不可变裸指针
let mut y = 42;
let r2 = &mut y as *mut i32;  // 可变裸指针

// 解引用裸指针 (unsafe)
unsafe {
    println!("{}", *r1);  // 42
    *r2 = 100;
    println!("{}", *r2);  // 100
}

// 空指针
let null_ptr: *const i32 = std::ptr::null();
let null_mut: *mut i32 = std::ptr::null_mut();

// ==================== Box<T> - 堆分配 ====================
// 独占所有权的堆分配
let boxed: Box<i32> = Box::new(42);
println!("{}", *boxed);

// 用于递归类型
enum List {
    Cons(i32, Box<List>),
    Nil,
}

let list = List::Cons(1, Box::new(List::Cons(2, Box::new(List::Nil))));

// ==================== Rc<T> - 引用计数 ====================
use std::rc::Rc;

let a = Rc::new(5);
let b = Rc::clone(&a);  // 增加引用计数
let c = Rc::clone(&a);

println!("count: {}", Rc::strong_count(&a));  // 3

// ==================== Arc<T> - 原子引用计数 (线程安全) ====================
use std::sync::Arc;
use std::thread;

let data = Arc::new(vec![1, 2, 3]);

let handles: Vec<_> = (0..3).map(|_| {
    let data = Arc::clone(&data);
    thread::spawn(move || {
        println!("{:?}", data);
    })
}).collect();

for handle in handles {
    handle.join().unwrap();
}

// ==================== RefCell<T> - 内部可变性 ====================
use std::cell::RefCell;

let cell = RefCell::new(5);

// 运行时借用检查
*cell.borrow_mut() += 1;
println!("{}", *cell.borrow());  // 6

// 多次借用会 panic
// let r1 = cell.borrow_mut();
// let r2 = cell.borrow_mut();  // panic!

// ==================== Cell<T> - 简单内部可变性 ====================
use std::cell::Cell;

let cell = Cell::new(5);
cell.set(10);
println!("{}", cell.get());  // 10

// ==================== Rc<RefCell<T>> 组合 ====================
use std::rc::Rc;
use std::cell::RefCell;

let shared = Rc::new(RefCell::new(vec![1, 2, 3]));

let a = Rc::clone(&shared);
let b = Rc::clone(&shared);

a.borrow_mut().push(4);
println!("{:?}", b.borrow());  // [1, 2, 3, 4]

// ==================== Weak<T> - 弱引用 ====================
use std::rc::{Rc, Weak};

let strong = Rc::new(5);
let weak: Weak<i32> = Rc::downgrade(&strong);

// 升级为强引用
if let Some(val) = weak.upgrade() {
    println!("{}", val);
}

drop(strong);
assert!(weak.upgrade().is_none());  // 已被释放

// ==================== Cow<T> - 写时复制 ====================
use std::borrow::Cow;

fn process(s: Cow<str>) -> Cow<str> {
    if s.contains("bad") {
        // 需要修改时才分配
        Cow::Owned(s.replace("bad", "good"))
    } else {
        s  // 不修改则保持借用
    }
}

let borrowed: Cow<str> = Cow::Borrowed("hello");
let owned: Cow<str> = Cow::Owned(String::from("world"));

// ==================== Pin<T> - 固定内存位置 ====================
use std::pin::Pin;
use std::marker::PhantomPinned;

struct Unmovable {
    data: String,
    _pin: PhantomPinned,
}

impl Unmovable {
    fn new(data: String) -> Pin<Box<Self>> {
        Box::pin(Unmovable {
            data,
            _pin: PhantomPinned,
        })
    }
}

// ==================== 指针比较 ====================
let x = 5;
let y = 5;
let rx = &x;
let ry = &y;

// 值比较
assert_eq!(*rx, *ry);

// 指针比较
assert!(!std::ptr::eq(rx, ry));

let rz = rx;
assert!(std::ptr::eq(rx, rz));
```

### 指针/引用对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 特性            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 引用类型        │ 对象/数组自动        │ 所有对象自动         │ 显式 &/*/指针        │ 显式 &/&mut          │
│ 获取地址        │ ❌                   │ id()                 │ &x                   │ &x / &mut x          │
│ 解引用          │ 自动                 │ 自动                 │ *p                   │ *p                   │
│ 空值            │ null/undefined       │ None                 │ nil                  │ Option<&T>           │
│ 可变性控制      │ readonly             │ 约定                 │ 值/指针              │ &T / &mut T          │
│ 悬垂指针        │ GC 防止              │ GC 防止              │ 可能                 │ 编译器阻止           │
│ 线程安全        │ 隔离                 │ GIL                  │ 手动                 │ Send/Sync            │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 智能指针        │ ❌                   │ ❌                   │ ❌                   │ Box/Rc/Arc           │
│ 引用计数        │ GC                   │ GC                   │ ❌                   │ Rc/Arc               │
│ 弱引用          │ WeakRef              │ weakref              │ ❌                   │ Weak                 │
│ 内部可变        │ ❌                   │ 默认可变             │ ❌                   │ Cell/RefCell         │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 使用场景指南

| 场景 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **共享只读数据** | 直接传递 | 直接传递 | 传值/指针 | `&T` |
| **共享可变数据** | 对象引用 | 对象引用 | 传指针 `*T` | `&mut T` |
| **多所有者共享** | 默认行为 | 默认行为 | 手动管理 | `Rc<T>`/`Arc<T>` |
| **堆分配** | 自动 | 自动 | `new`/`make` | `Box<T>` |
| **可选值** | `\| undefined` | `\| None` | `*T` (nil) | `Option<T>` |
| **避免循环引用** | WeakRef | weakref | 手动 | `Weak<T>` |
| **线程共享** | Worker隔离 | Queue | `sync` 包 | `Arc<Mutex<T>>` |

### 常见陷阱

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ Go: 在循环中取地址                                                          │
│   for _, v := range items {                                                 │
│       pointers = append(pointers, &v)  // 错误！都指向同一个 v              │
│   }                                                                         │
│   // 正确做法:                                                              │
│   for i := range items {                                                    │
│       pointers = append(pointers, &items[i])                                │
│   }                                                                         │
├─────────────────────────────────────────────────────────────────────────────┤
│ Rust: 同时持有可变和不可变引用                                               │
│   let mut v = vec![1, 2, 3];                                                │
│   let first = &v[0];                                                        │
│   v.push(4);  // 编译错误！first 还在使用                                   │
│   println!("{}", first);                                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│ Python: 以为赋值会复制                                                       │
│   a = [1, 2, 3]                                                             │
│   b = a  # b 是 a 的引用，不是复制！                                        │
│   b.append(4)                                                               │
│   print(a)  # [1, 2, 3, 4]                                                  │
├─────────────────────────────────────────────────────────────────────────────┤
│ TypeScript: 以为 const 能防止修改内容                                        │
│   const arr = [1, 2, 3];                                                    │
│   arr.push(4);  // OK! const 只防止重新赋值                                 │
│   // arr = [5, 6];  // 这才会报错                                           │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 🔄 枚举与状态机

### 枚举类型概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 枚举关键字 | `enum` | `Enum` | `const/iota` | `enum` |
| 关联数据 | ❌ | ❌ | ❌ | ✅ |
| 模式匹配 | ❌ | ✅ match | switch | ✅ match |
| 穷尽检查 | ❌ | ❌ | ❌ | ✅ |
| 字符串枚举 | ✅ | ✅ | 手动 | 需派生 |
| 位标志 | 手动 | `Flag` | `iota` | `bitflags` |

### TypeScript 枚举

```typescript
// ==================== 数字枚举 ====================
enum Direction {
    Up,      // 0
    Down,    // 1
    Left,    // 2
    Right,   // 3
}

// 指定值
enum Status {
    Pending = 1,
    Active = 2,
    Inactive = 4,
    Deleted = 8,
}

// 使用
const dir: Direction = Direction.Up;
const name = Direction[0];  // "Up" (反向映射)

// ==================== 字符串枚举 ====================
enum Color {
    Red = "RED",
    Green = "GREEN",
    Blue = "BLUE",
}

// 无反向映射
const color: Color = Color.Red;

// ==================== const 枚举 (编译时内联) ====================
const enum HttpStatus {
    OK = 200,
    NotFound = 404,
    ServerError = 500,
}
// 编译后直接内联为数字，无运行时对象

// ==================== 联合类型替代枚举 (推荐) ====================
type Direction = "up" | "down" | "left" | "right";
type Status = "pending" | "active" | "inactive";

// 更好的类型推断
function move(dir: Direction) {
    switch (dir) {
        case "up": return { y: -1 };
        case "down": return { y: 1 };
        case "left": return { x: -1 };
        case "right": return { x: 1 };
    }
}

// ==================== 位标志 ====================
enum Permission {
    None = 0,
    Read = 1 << 0,    // 1
    Write = 1 << 1,   // 2
    Execute = 1 << 2, // 4
    All = Read | Write | Execute,
}

const perms = Permission.Read | Permission.Write;
const canRead = (perms & Permission.Read) !== 0;

// ==================== 状态机 ====================
type OrderStatus = "pending" | "confirmed" | "shipped" | "delivered" | "cancelled";

interface Order {
    id: string;
    status: OrderStatus;
}

// 状态转换规则
const transitions: Record<OrderStatus, OrderStatus[]> = {
    pending: ["confirmed", "cancelled"],
    confirmed: ["shipped", "cancelled"],
    shipped: ["delivered"],
    delivered: [],
    cancelled: [],
};

function canTransition(from: OrderStatus, to: OrderStatus): boolean {
    return transitions[from].includes(to);
}

function transition(order: Order, newStatus: OrderStatus): Order {
    if (!canTransition(order.status, newStatus)) {
        throw new Error(`Cannot transition from ${order.status} to ${newStatus}`);
    }
    return { ...order, status: newStatus };
}

// ==================== 类型安全的状态机 ====================
type State = 
    | { type: "idle" }
    | { type: "loading" }
    | { type: "success"; data: string }
    | { type: "error"; message: string };

type Action = 
    | { type: "FETCH" }
    | { type: "SUCCESS"; data: string }
    | { type: "ERROR"; message: string }
    | { type: "RESET" };

function reducer(state: State, action: Action): State {
    switch (state.type) {
        case "idle":
            if (action.type === "FETCH") return { type: "loading" };
            break;
        case "loading":
            if (action.type === "SUCCESS") return { type: "success", data: action.data };
            if (action.type === "ERROR") return { type: "error", message: action.message };
            break;
        case "success":
        case "error":
            if (action.type === "RESET") return { type: "idle" };
            break;
    }
    return state;
}
```

### Python 枚举

```python
from enum import Enum, IntEnum, Flag, auto, unique

# ==================== 基本枚举 ====================
class Color(Enum):
    RED = 1
    GREEN = 2
    BLUE = 3

# 使用
color = Color.RED
print(color.name)   # "RED"
print(color.value)  # 1

# 迭代
for c in Color:
    print(c)

# 比较
Color.RED == Color.RED   # True
Color.RED is Color.RED   # True

# ==================== 自动值 ====================
class Direction(Enum):
    UP = auto()      # 1
    DOWN = auto()    # 2
    LEFT = auto()    # 3
    RIGHT = auto()   # 4

# ==================== 字符串枚举 ====================
class Status(str, Enum):
    PENDING = "pending"
    ACTIVE = "active"
    INACTIVE = "inactive"

# 可直接当字符串用
print(f"Status: {Status.PENDING}")  # "Status: pending"

# ==================== 整数枚举 ====================
class HttpStatus(IntEnum):
    OK = 200
    NOT_FOUND = 404
    SERVER_ERROR = 500

# 可直接比较数字
HttpStatus.OK == 200  # True

# ==================== 唯一值约束 ====================
@unique
class Unique(Enum):
    A = 1
    B = 2
    # C = 1  # 报错！值重复

# ==================== 位标志 ====================
class Permission(Flag):
    NONE = 0
    READ = auto()     # 1
    WRITE = auto()    # 2
    EXECUTE = auto()  # 4
    ALL = READ | WRITE | EXECUTE

perms = Permission.READ | Permission.WRITE
Permission.READ in perms  # True

# ==================== 模式匹配 (3.10+) ====================
def describe_color(color: Color) -> str:
    match color:
        case Color.RED:
            return "Hot color"
        case Color.BLUE:
            return "Cool color"
        case Color.GREEN:
            return "Nature color"

# ==================== 状态机 ====================
from enum import Enum
from typing import Dict, Set

class OrderStatus(Enum):
    PENDING = "pending"
    CONFIRMED = "confirmed"
    SHIPPED = "shipped"
    DELIVERED = "delivered"
    CANCELLED = "cancelled"

# 状态转换规则
TRANSITIONS: Dict[OrderStatus, Set[OrderStatus]] = {
    OrderStatus.PENDING: {OrderStatus.CONFIRMED, OrderStatus.CANCELLED},
    OrderStatus.CONFIRMED: {OrderStatus.SHIPPED, OrderStatus.CANCELLED},
    OrderStatus.SHIPPED: {OrderStatus.DELIVERED},
    OrderStatus.DELIVERED: set(),
    OrderStatus.CANCELLED: set(),
}

class Order:
    def __init__(self, id: str):
        self.id = id
        self._status = OrderStatus.PENDING
    
    @property
    def status(self) -> OrderStatus:
        return self._status
    
    def transition(self, new_status: OrderStatus) -> None:
        if new_status not in TRANSITIONS[self._status]:
            raise ValueError(
                f"Cannot transition from {self._status.value} to {new_status.value}"
            )
        self._status = new_status

# ==================== 带方法的枚举 ====================
class Planet(Enum):
    MERCURY = (3.303e+23, 2.4397e6)
    VENUS = (4.869e+24, 6.0518e6)
    EARTH = (5.976e+24, 6.37814e6)
    
    def __init__(self, mass: float, radius: float):
        self.mass = mass
        self.radius = radius
    
    @property
    def surface_gravity(self) -> float:
        G = 6.67430e-11
        return G * self.mass / (self.radius ** 2)

print(Planet.EARTH.surface_gravity)
```

### Go 枚举与状态机

```go
// ==================== iota 枚举 ====================
type Direction int

const (
    Up Direction = iota  // 0
    Down                 // 1
    Left                 // 2
    Right                // 3
)

// 方法
func (d Direction) String() string {
    return [...]string{"Up", "Down", "Left", "Right"}[d]
}

// ==================== 字符串枚举 ====================
type Status string

const (
    StatusPending  Status = "pending"
    StatusActive   Status = "active"
    StatusInactive Status = "inactive"
)

// ==================== 位标志 ====================
type Permission int

const (
    PermNone    Permission = 0
    PermRead    Permission = 1 << iota  // 1
    PermWrite                           // 2
    PermExecute                         // 4
    PermAll     = PermRead | PermWrite | PermExecute
)

func (p Permission) Has(flag Permission) bool {
    return p&flag != 0
}

perms := PermRead | PermWrite
perms.Has(PermRead)  // true

// ==================== 状态机 ====================
type OrderStatus int

const (
    OrderPending OrderStatus = iota
    OrderConfirmed
    OrderShipped
    OrderDelivered
    OrderCancelled
)

// 转换规则
var transitions = map[OrderStatus][]OrderStatus{
    OrderPending:   {OrderConfirmed, OrderCancelled},
    OrderConfirmed: {OrderShipped, OrderCancelled},
    OrderShipped:   {OrderDelivered},
    OrderDelivered: {},
    OrderCancelled: {},
}

type Order struct {
    ID     string
    Status OrderStatus
}

func (o *Order) CanTransition(to OrderStatus) bool {
    allowed := transitions[o.Status]
    for _, s := range allowed {
        if s == to {
            return true
        }
    }
    return false
}

func (o *Order) Transition(to OrderStatus) error {
    if !o.CanTransition(to) {
        return fmt.Errorf("cannot transition from %d to %d", o.Status, to)
    }
    o.Status = to
    return nil
}

// ==================== 接口状态机 ====================
type State interface {
    Handle(event Event) State
    Name() string
}

type Event string

const (
    EventStart   Event = "start"
    EventSuccess Event = "success"
    EventError   Event = "error"
    EventReset   Event = "reset"
)

// 具体状态
type IdleState struct{}

func (s IdleState) Handle(e Event) State {
    if e == EventStart {
        return LoadingState{}
    }
    return s
}

func (s IdleState) Name() string { return "idle" }

type LoadingState struct{}

func (s LoadingState) Handle(e Event) State {
    switch e {
    case EventSuccess:
        return SuccessState{}
    case EventError:
        return ErrorState{}
    }
    return s
}

func (s LoadingState) Name() string { return "loading" }

type SuccessState struct{}
type ErrorState struct{}

// 状态机
type StateMachine struct {
    current State
}

func (sm *StateMachine) Send(e Event) {
    sm.current = sm.current.Handle(e)
}

// ==================== 函数式状态机 ====================
type StateFunc func(Event) StateFunc

func idleState(e Event) StateFunc {
    if e == EventStart {
        return loadingState
    }
    return idleState
}

func loadingState(e Event) StateFunc {
    switch e {
    case EventSuccess:
        return successState
    case EventError:
        return errorState
    }
    return loadingState
}

func successState(e Event) StateFunc {
    if e == EventReset {
        return idleState
    }
    return successState
}

func errorState(e Event) StateFunc {
    if e == EventReset {
        return idleState
    }
    return errorState
}
```

### Rust 枚举与状态机

```rust
// ==================== 基本枚举 ====================
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

// 使用
let dir = Direction::Up;

// 模式匹配 (必须穷尽)
match dir {
    Direction::Up => println!("Going up"),
    Direction::Down => println!("Going down"),
    Direction::Left => println!("Going left"),
    Direction::Right => println!("Going right"),
}

// ==================== 带关联数据的枚举 (代数数据类型) ====================
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

let msg = Message::Move { x: 10, y: 20 };

match msg {
    Message::Quit => println!("Quit"),
    Message::Move { x, y } => println!("Move to ({}, {})", x, y),
    Message::Write(text) => println!("Text: {}", text),
    Message::ChangeColor(r, g, b) => println!("Color: ({}, {}, {})", r, g, b),
}

// ==================== Option 和 Result ====================
// 内置的带数据枚举
enum Option<T> {
    Some(T),
    None,
}

enum Result<T, E> {
    Ok(T),
    Err(E),
}

// if let 简化匹配
if let Some(value) = maybe_value {
    println!("Got: {}", value);
}

// ==================== 数值枚举 ====================
#[derive(Debug, Clone, Copy, PartialEq)]
#[repr(u8)]
enum HttpStatus {
    Ok = 200,
    NotFound = 404,
    ServerError = 500,
}

// ==================== 字符串转换 ====================
use strum_macros::{Display, EnumString};

#[derive(Debug, Display, EnumString)]
enum Status {
    #[strum(serialize = "pending")]
    Pending,
    #[strum(serialize = "active")]
    Active,
    #[strum(serialize = "inactive")]
    Inactive,
}

let s: Status = "pending".parse().unwrap();
println!("{}", Status::Active);  // "active"

// ==================== 位标志 ====================
use bitflags::bitflags;

bitflags! {
    struct Permission: u32 {
        const NONE = 0;
        const READ = 1 << 0;
        const WRITE = 1 << 1;
        const EXECUTE = 1 << 2;
        const ALL = Self::READ.bits | Self::WRITE.bits | Self::EXECUTE.bits;
    }
}

let perms = Permission::READ | Permission::WRITE;
perms.contains(Permission::READ);  // true

// ==================== 类型安全状态机 ====================
// 使用类型状态模式
struct Order<S> {
    id: String,
    state: S,
}

// 状态类型
struct Pending;
struct Confirmed;
struct Shipped;
struct Delivered;
struct Cancelled;

// 仅允许特定转换
impl Order<Pending> {
    fn new(id: String) -> Self {
        Order { id, state: Pending }
    }
    
    fn confirm(self) -> Order<Confirmed> {
        Order { id: self.id, state: Confirmed }
    }
    
    fn cancel(self) -> Order<Cancelled> {
        Order { id: self.id, state: Cancelled }
    }
}

impl Order<Confirmed> {
    fn ship(self) -> Order<Shipped> {
        Order { id: self.id, state: Shipped }
    }
    
    fn cancel(self) -> Order<Cancelled> {
        Order { id: self.id, state: Cancelled }
    }
}

impl Order<Shipped> {
    fn deliver(self) -> Order<Delivered> {
        Order { id: self.id, state: Delivered }
    }
}

// 使用 - 编译时检查状态转换
let order = Order::new("123".to_string());
let order = order.confirm();  // OK
let order = order.ship();     // OK
let order = order.deliver();  // OK
// order.ship();  // 编译错误！Delivered 状态没有 ship 方法

// ==================== 枚举状态机 ====================
#[derive(Debug, Clone)]
enum OrderState {
    Pending,
    Confirmed { confirmed_at: DateTime },
    Shipped { shipped_at: DateTime, tracking: String },
    Delivered { delivered_at: DateTime },
    Cancelled { reason: String },
}

enum OrderEvent {
    Confirm,
    Ship { tracking: String },
    Deliver,
    Cancel { reason: String },
}

impl OrderState {
    fn transition(self, event: OrderEvent) -> Result<OrderState, &'static str> {
        match (self, event) {
            (OrderState::Pending, OrderEvent::Confirm) => {
                Ok(OrderState::Confirmed { confirmed_at: now() })
            }
            (OrderState::Pending, OrderEvent::Cancel { reason }) => {
                Ok(OrderState::Cancelled { reason })
            }
            (OrderState::Confirmed { .. }, OrderEvent::Ship { tracking }) => {
                Ok(OrderState::Shipped { shipped_at: now(), tracking })
            }
            (OrderState::Confirmed { .. }, OrderEvent::Cancel { reason }) => {
                Ok(OrderState::Cancelled { reason })
            }
            (OrderState::Shipped { .. }, OrderEvent::Deliver) => {
                Ok(OrderState::Delivered { delivered_at: now() })
            }
            _ => Err("Invalid state transition"),
        }
    }
}

// ==================== 带方法的枚举 ====================
enum Shape {
    Circle { radius: f64 },
    Rectangle { width: f64, height: f64 },
    Triangle { base: f64, height: f64 },
}

impl Shape {
    fn area(&self) -> f64 {
        match self {
            Shape::Circle { radius } => std::f64::consts::PI * radius * radius,
            Shape::Rectangle { width, height } => width * height,
            Shape::Triangle { base, height } => 0.5 * base * height,
        }
    }
}
```

### 枚举与状态机对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 特性            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 枚举定义        │ enum E { A, B }      │ class E(Enum)        │ const + iota         │ enum E { A, B }      │
│ 关联数据        │ ❌ (用联合类型)      │ ❌                   │ ❌                   │ ✅ E::A(data)        │
│ 字符串枚举      │ enum E { A="a" }     │ class E(str,Enum)    │ type E string        │ strum 派生           │
│ 位标志          │ 手动位运算           │ Flag 类              │ iota + 位运算        │ bitflags! 宏         │
│ 穷尽检查        │ ❌                   │ ❌                   │ ❌                   │ ✅ 编译时            │
│ 模式匹配        │ switch (有限)        │ match (3.10+)        │ switch               │ match (强大)         │
│ 状态机类型安全  │ 联合类型             │ 运行时检查           │ 运行时检查           │ 类型状态模式         │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 状态机设计模式

| 模式 | 描述 | 推荐语言 |
|------|------|----------|
| **转换表** | Map 存储允许的转换 | 全部 |
| **状态接口** | 每个状态实现接口 | Go, TypeScript |
| **类型状态** | 编译时状态检查 | Rust |
| **代数数据类型** | 枚举带关联数据 | Rust |
| **联合类型** | 类型安全的状态表示 | TypeScript |
| **函数式** | 状态函数返回下一状态 | Go, Rust |

### 状态机最佳实践

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ 1. 明确定义所有状态和事件                                                    │
│ 2. 使用转换表/矩阵明确允许的转换                                            │
│ 3. 尽可能利用类型系统在编译时检查                                           │
│ 4. 考虑状态携带的数据 (进入时间、关联信息等)                                │
│ 5. 处理无效转换 (返回错误 vs 忽略 vs panic)                                 │
│ 6. 考虑是否需要状态变更通知/回调                                            │
│ 7. 测试所有状态转换路径                                                     │
└─────────────────────────────────────────────────────────────────────────────┘
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

## 🔄 Round 取整函数对比

### 取整函数概览

| 函数 | 行为 | 示例 |
|------|------|------|
| `round` | 四舍五入 | round(2.5) → 2 或 3 |
| `floor` | 向下取整（向负无穷） | floor(-2.3) → -3 |
| `ceil` | 向上取整（向正无穷） | ceil(2.1) → 3 |
| `trunc` | 向零取整 | trunc(-2.7) → -2 |

### 行为对比图示

```
数轴:  -4    -3    -2    -1     0     1     2     3     4
        |     |     |     |     |     |     |     |     |
        
输入: -2.7
├─ floor(-2.7) = -3  ←───────●
├─ ceil(-2.7)  = -2          ●───────→
├─ trunc(-2.7) = -2          ●───────→  (向零)
└─ round(-2.7) = -3  ←───────●         (四舍五入)

输入: 2.3
├─ floor(2.3)  = 2   ←───────●
├─ ceil(2.3)   = 3           ●───────→
├─ trunc(2.3)  = 2   ←───────●         (向零)
└─ round(2.3)  = 2   ←───────●         (四舍五入)
```

### TypeScript 取整函数

```typescript
const n = 2.7;
const neg = -2.7;

// ==================== round (四舍五入) ====================
Math.round(2.4);    // 2
Math.round(2.5);    // 3  (0.5 向上)
Math.round(2.6);    // 3
Math.round(-2.4);   // -2
Math.round(-2.5);   // -2 (0.5 向上，即向正无穷)
Math.round(-2.6);   // -3

// ==================== floor (向下取整) ====================
Math.floor(2.7);    // 2
Math.floor(2.1);    // 2
Math.floor(-2.1);   // -3 (向负无穷)
Math.floor(-2.7);   // -3

// ==================== ceil (向上取整) ====================
Math.ceil(2.1);     // 3
Math.ceil(2.9);     // 3
Math.ceil(-2.1);    // -2 (向正无穷)
Math.ceil(-2.9);    // -2

// ==================== trunc (向零取整) ====================
Math.trunc(2.7);    // 2
Math.trunc(-2.7);   // -2

// ==================== 精度问题处理 ====================
// 银行家舍入法 (四舍六入五成双)
function bankersRound(n: number, decimals: number = 0): number {
  const factor = Math.pow(10, decimals);
  const shifted = n * factor;
  const truncated = Math.trunc(shifted);
  const decimal = shifted - truncated;
  
  if (Math.abs(decimal) === 0.5) {
    // 0.5 时取最近的偶数
    return (truncated % 2 === 0 ? truncated : truncated + Math.sign(n)) / factor;
  }
  return Math.round(shifted) / factor;
}

bankersRound(2.5);   // 2 (取偶)
bankersRound(3.5);   // 4 (取偶)
bankersRound(2.6);   // 3

// ==================== 保留小数位 ====================
// 四舍五入到 2 位小数
Math.round(3.14159 * 100) / 100;  // 3.14

// 通用函数
function roundTo(n: number, decimals: number): number {
  const factor = Math.pow(10, decimals);
  return Math.round(n * factor) / factor;
}

roundTo(3.14159, 2);   // 3.14
roundTo(3.14159, 3);   // 3.142
```

### Python 取整函数

```python
import math
from decimal import Decimal, ROUND_HALF_UP, ROUND_HALF_EVEN

n = 2.7
neg = -2.7

# ==================== round (银行家舍入) ====================
round(2.4)      # 2
round(2.5)      # 2  ⚠️ Python 3 使用银行家舍入！
round(2.6)      # 3
round(3.5)      # 4  (取偶数)
round(-2.5)     # -2 (取偶数)

# 指定小数位
round(3.14159, 2)   # 3.14
round(3.145, 2)     # 3.14 ⚠️ (银行家舍入)
round(3.155, 2)     # 3.15

# ==================== floor (向下取整) ====================
math.floor(2.7)     # 2
math.floor(2.1)     # 2
math.floor(-2.1)    # -3
math.floor(-2.7)    # -3

# 整数除法也是 floor
7 // 3              # 2
-7 // 3             # -3 (向下)

# ==================== ceil (向上取整) ====================
math.ceil(2.1)      # 3
math.ceil(2.9)      # 3
math.ceil(-2.1)     # -2
math.ceil(-2.9)     # -2

# ==================== trunc (向零取整) ====================
math.trunc(2.7)     # 2
math.trunc(-2.7)    # -2
int(2.7)            # 2 (等效)
int(-2.7)           # -2

# ==================== Decimal 精确舍入 ====================
# 传统四舍五入
Decimal('2.5').quantize(Decimal('1'), rounding=ROUND_HALF_UP)  # 3
Decimal('3.5').quantize(Decimal('1'), rounding=ROUND_HALF_UP)  # 4

# 银行家舍入
Decimal('2.5').quantize(Decimal('1'), rounding=ROUND_HALF_EVEN)  # 2
Decimal('3.5').quantize(Decimal('1'), rounding=ROUND_HALF_EVEN)  # 4

# 保留小数位
Decimal('3.14159').quantize(Decimal('0.01'), rounding=ROUND_HALF_UP)  # 3.14

# ==================== 自定义四舍五入 ====================
def round_half_up(n, decimals=0):
    """传统四舍五入 (非银行家舍入)"""
    multiplier = 10 ** decimals
    return math.floor(n * multiplier + 0.5) / multiplier

round_half_up(2.5)    # 3
round_half_up(3.5)    # 4
round_half_up(-2.5)   # -2
```

### Go 取整函数

```go
import "math"

n := 2.7
neg := -2.7

// ==================== Round (四舍五入) ====================
math.Round(2.4)     // 2
math.Round(2.5)     // 3  (Go 使用 "round half away from zero")
math.Round(2.6)     // 3
math.Round(-2.4)    // -2
math.Round(-2.5)    // -3 (远离零)
math.Round(-2.6)    // -3

// ==================== Floor (向下取整) ====================
math.Floor(2.7)     // 2
math.Floor(2.1)     // 2
math.Floor(-2.1)    // -3
math.Floor(-2.7)    // -3

// ==================== Ceil (向上取整) ====================
math.Ceil(2.1)      // 3
math.Ceil(2.9)      // 3
math.Ceil(-2.1)     // -2
math.Ceil(-2.9)     // -2

// ==================== Trunc (向零取整) ====================
math.Trunc(2.7)     // 2
math.Trunc(-2.7)    // -2

// ==================== RoundToEven (银行家舍入, Go 1.10+) ====================
math.RoundToEven(2.5)   // 2
math.RoundToEven(3.5)   // 4
math.RoundToEven(-2.5)  // -2
math.RoundToEven(-3.5)  // -4

// ==================== 保留小数位 ====================
func roundTo(n float64, decimals int) float64 {
    factor := math.Pow(10, float64(decimals))
    return math.Round(n*factor) / factor
}

roundTo(3.14159, 2)  // 3.14
roundTo(3.14159, 3)  // 3.142

// 转为整数
int(math.Round(2.7))  // 3
int(math.Floor(2.7))  // 2
```

### Rust 取整函数

```rust
let n: f64 = 2.7;
let neg: f64 = -2.7;

// ==================== round (四舍五入) ====================
(2.4_f64).round()   // 2.0
(2.5_f64).round()   // 3.0 (远离零)
(2.6_f64).round()   // 3.0
(-2.4_f64).round()  // -2.0
(-2.5_f64).round()  // -3.0 (远离零)
(-2.6_f64).round()  // -3.0

// ==================== floor (向下取整) ====================
(2.7_f64).floor()   // 2.0
(2.1_f64).floor()   // 2.0
(-2.1_f64).floor()  // -3.0
(-2.7_f64).floor()  // -3.0

// ==================== ceil (向上取整) ====================
(2.1_f64).ceil()    // 3.0
(2.9_f64).ceil()    // 3.0
(-2.1_f64).ceil()   // -2.0
(-2.9_f64).ceil()   // -2.0

// ==================== trunc (向零取整) ====================
(2.7_f64).trunc()   // 2.0
(-2.7_f64).trunc()  // -2.0

// ==================== 转为整数 ====================
(2.7_f64).round() as i32    // 3
(2.7_f64).floor() as i32    // 2
(-2.7_f64).round() as i32   // -3

// ==================== 保留小数位 ====================
fn round_to(n: f64, decimals: u32) -> f64 {
    let factor = 10_f64.powi(decimals as i32);
    (n * factor).round() / factor
}

round_to(3.14159, 2)  // 3.14
round_to(3.14159, 3)  // 3.142

// ==================== 银行家舍入 (需要手动实现或使用 crate) ====================
fn round_half_even(n: f64) -> f64 {
    let floor = n.floor();
    let frac = n - floor;
    
    if (frac - 0.5).abs() < f64::EPSILON {
        // 正好是 0.5，取偶数
        if floor as i64 % 2 == 0 {
            floor
        } else {
            floor + 1.0
        }
    } else {
        n.round()
    }
}

round_half_even(2.5)  // 2.0
round_half_even(3.5)  // 4.0
```

### 取整函数对比总结

```
┌──────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 输入         │ round          │ floor          │ ceil           │ trunc          │
├──────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│  2.3         │  2             │  2             │  3             │  2             │
│  2.5         │  3 (或 2*)     │  2             │  3             │  2             │
│  2.7         │  3             │  2             │  3             │  2             │
│ -2.3         │ -2             │ -3             │ -2             │ -2             │
│ -2.5         │ -3 (或 -2*)    │ -3             │ -2             │ -2             │
│ -2.7         │ -3             │ -3             │ -2             │ -2             │
└──────────────┴────────────────┴────────────────┴────────────────┴────────────────┘

* round(2.5) 结果因语言而异:
  - TypeScript/Go/Rust: 3 (远离零)
  - Python: 2 (银行家舍入，取偶数)
```

### 0.5 舍入规则对比

| 语言 | round(2.5) | round(3.5) | round(-2.5) | 规则名称 |
|------|------------|------------|-------------|----------|
| TypeScript | 3 | 4 | -2 | Round half up |
| Python | 2 | 4 | -2 | Bankers (half even) |
| Go `Round` | 3 | 4 | -3 | Round half away from zero |
| Go `RoundToEven` | 2 | 4 | -2 | Bankers (half even) |
| Rust | 3 | 4 | -3 | Round half away from zero |

### 何时使用哪种取整

| 场景 | 推荐函数 | 原因 |
|------|----------|------|
| 金融计算 | 银行家舍入 | 统计上无偏差 |
| 像素/坐标 | floor/ceil | 确定方向 |
| 数组索引 | floor 或 trunc | 向下取整 |
| 四舍五入显示 | round | 用户直觉 |
| 向零截断 | trunc | 去掉小数 |

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

## 📍 变量作用域与代码块

### 作用域概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 块级作用域 | ✅ `let/const` | ❌ | ✅ | ✅ |
| 函数作用域 | ✅ `var` | ✅ | ✅ | ✅ |
| 全局作用域 | ✅ | ✅ | ✅ (包级) | ✅ (crate级) |
| 变量提升 | ✅ `var` | ❌ | ❌ | ❌ |
| 暂时性死区 | ✅ `let/const` | ❌ | ❌ | ❌ |
| 变量遮蔽 | ✅ | ✅ (嵌套) | ✅ `:=` | ✅ |
| 块表达式 | ❌ | ❌ | ❌ | ✅ |

### 作用域图解

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                              作用域层级                                      │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  ┌─────────────────────────────────────────────────────────────────────┐   │
│  │ 全局/模块作用域                                                      │   │
│  │                                                                     │   │
│  │  ┌─────────────────────────────────────────────────────────────┐   │   │
│  │  │ 函数作用域                                                   │   │   │
│  │  │                                                             │   │   │
│  │  │  ┌─────────────────────────────────────────────────────┐   │   │   │
│  │  │  │ 块作用域 (if/for/while/{})                          │   │   │   │
│  │  │  │                                                     │   │   │   │
│  │  │  │  ┌─────────────────────────────────────────────┐   │   │   │   │
│  │  │  │  │ 嵌套块作用域                                 │   │   │   │   │
│  │  │  │  │  内层可访问外层变量                          │   │   │   │   │
│  │  │  │  │  外层不可访问内层变量                        │   │   │   │   │
│  │  │  │  └─────────────────────────────────────────────┘   │   │   │   │
│  │  │  └─────────────────────────────────────────────────────┘   │   │   │
│  │  └─────────────────────────────────────────────────────────────┘   │   │
│  └─────────────────────────────────────────────────────────────────────┘   │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### TypeScript 作用域

```typescript
// ==================== var vs let vs const ====================
// var - 函数作用域，有变量提升
function varExample() {
    console.log(x);  // undefined (提升但未初始化)
    var x = 10;
    
    if (true) {
        var x = 20;  // 同一个 x
    }
    console.log(x);  // 20
}

// let - 块级作用域，暂时性死区
function letExample() {
    // console.log(x);  // ReferenceError: 暂时性死区
    let x = 10;
    
    if (true) {
        let x = 20;  // 新的 x
        console.log(x);  // 20
    }
    console.log(x);  // 10
}

// const - 块级作用域，不可重新赋值
function constExample() {
    const x = 10;
    // x = 20;  // TypeError
    
    const obj = { a: 1 };
    obj.a = 2;  // OK，可以修改属性
    // obj = {};  // TypeError
}

// ==================== 块级作用域 {} ====================
{
    let blockScoped = "只在块内可见";
    const alsoBlockScoped = "同样只在块内";
    var notBlockScoped = "函数作用域";
}
// console.log(blockScoped);  // ReferenceError
console.log(notBlockScoped);  // OK

// ==================== 循环中的作用域 ====================
// var 的经典陷阱
for (var i = 0; i < 3; i++) {
    setTimeout(() => console.log(i), 100);
}
// 输出: 3, 3, 3 (都是同一个 i)

// let 解决问题
for (let i = 0; i < 3; i++) {
    setTimeout(() => console.log(i), 100);
}
// 输出: 0, 1, 2 (每次迭代新的 i)

// ==================== 变量遮蔽 (Shadowing) ====================
let value = 10;

function outer() {
    let value = 20;  // 遮蔽外层
    
    function inner() {
        let value = 30;  // 再次遮蔽
        console.log(value);  // 30
    }
    
    inner();
    console.log(value);  // 20
}

outer();
console.log(value);  // 10

// ==================== 闭包与作用域 ====================
function createCounter() {
    let count = 0;  // 被闭包捕获
    
    return {
        increment: () => ++count,
        decrement: () => --count,
        getCount: () => count,
    };
}

const counter = createCounter();
counter.increment();
counter.increment();
console.log(counter.getCount());  // 2

// ==================== 立即执行函数 (IIFE) ====================
// 创建独立作用域
(function() {
    var privateVar = "不污染全局";
    // ...
})();

// 现代替代：块 + let/const
{
    let privateVar = "不污染全局";
    // ...
}

// ==================== 全局作用域 ====================
// 浏览器: window
// Node.js: global
// 通用: globalThis

globalThis.myGlobal = "全局变量";

// 避免污染全局
// 使用模块系统代替全局变量
```

### Python 作用域

```python
# ==================== LEGB 规则 ====================
# Local -> Enclosing -> Global -> Built-in

# Built-in (内置)
print  # 内置函数

# Global (全局/模块级)
global_var = "全局变量"

def outer():
    # Enclosing (闭包)
    enclosing_var = "闭包变量"
    
    def inner():
        # Local (局部)
        local_var = "局部变量"
        
        print(local_var)      # Local
        print(enclosing_var)  # Enclosing
        print(global_var)     # Global
        print(len)            # Built-in

# ==================== Python 没有块级作用域！ ====================
if True:
    x = 10  # 不是块级变量！

print(x)  # 10 - 可以访问

for i in range(3):
    y = i

print(i, y)  # 2, 2 - 循环变量泄漏

# 列表推导式有自己的作用域 (Python 3)
[z for z in range(3)]
# print(z)  # NameError in Python 3 (Python 2 会泄漏)

# ==================== global 关键字 ====================
counter = 0

def increment():
    global counter  # 声明使用全局变量
    counter += 1

increment()
print(counter)  # 1

# 不使用 global
def bad_increment():
    # counter += 1  # UnboundLocalError
    # 赋值会创建局部变量，但右侧引用了未定义的局部变量
    pass

# ==================== nonlocal 关键字 ====================
def outer():
    count = 0
    
    def inner():
        nonlocal count  # 声明使用闭包变量
        count += 1
        return count
    
    return inner

counter = outer()
print(counter())  # 1
print(counter())  # 2

# ==================== 闭包陷阱 ====================
# 经典错误
functions = []
for i in range(3):
    functions.append(lambda: i)

print([f() for f in functions])  # [2, 2, 2] - 都是最后的 i

# 解决方案 1: 默认参数捕获
functions = []
for i in range(3):
    functions.append(lambda i=i: i)  # i=i 捕获当前值

print([f() for f in functions])  # [0, 1, 2]

# 解决方案 2: 使用 functools.partial
from functools import partial

functions = []
for i in range(3):
    functions.append(partial(lambda x: x, i))

# ==================== 模拟块级作用域 ====================
# 方法 1: 函数
def block():
    x = "块内变量"
    return x

# 方法 2: 删除变量
if True:
    temp = expensive_computation()
    result = process(temp)
    del temp  # 手动清理

# ==================== 类作用域 ====================
class MyClass:
    class_var = "类变量"  # 类级别
    
    def __init__(self):
        self.instance_var = "实例变量"  # 实例级别
    
    def method(self):
        local_var = "局部变量"  # 方法内局部
        print(self.class_var)     # 通过 self 访问
        print(MyClass.class_var)  # 通过类名访问
```

### Go 作用域

```go
// ==================== 块级作用域 ====================
func blockScope() {
    x := 10
    
    {
        y := 20  // 只在块内可见
        x := 30  // 遮蔽外层 x
        fmt.Println(x, y)  // 30, 20
    }
    
    // fmt.Println(y)  // 编译错误: undefined
    fmt.Println(x)  // 10
}

// ==================== if/for/switch 作用域 ====================
func controlScope() {
    // if 初始化语句的变量只在 if 块内可见
    if x := compute(); x > 0 {
        fmt.Println(x)
    } else {
        fmt.Println(-x)
    }
    // fmt.Println(x)  // 编译错误
    
    // for 的变量只在循环内可见
    for i := 0; i < 3; i++ {
        fmt.Println(i)
    }
    // fmt.Println(i)  // 编译错误
    
    // switch 初始化
    switch x := getValue(); x {
    case 1:
        fmt.Println("one")
    default:
        fmt.Println(x)
    }
}

// ==================== 变量遮蔽 (Shadowing) ====================
var x = 10  // 包级变量

func shadowExample() {
    fmt.Println(x)  // 10 - 包级
    
    x := 20  // 遮蔽包级变量
    fmt.Println(x)  // 20
    
    {
        x := 30  // 再次遮蔽
        fmt.Println(x)  // 30
    }
    
    fmt.Println(x)  // 20
}

// 常见陷阱: 短声明遮蔽
func shadowTrap() error {
    var err error
    
    if true {
        result, err := doSomething()  // 新的 err，遮蔽外层！
        if err != nil {
            return err
        }
        _ = result
    }
    
    return err  // 始终为 nil！
}

// 正确做法
func shadowFixed() error {
    var err error
    var result int
    
    if true {
        result, err = doSomething()  // = 不是 :=
        if err != nil {
            return err
        }
    }
    
    _ = result
    return err
}

// ==================== 包级作用域 ====================
package mypackage

var PackageVar = "包内所有文件可见"    // 大写: 导出
var privateVar = "仅包内可见"          // 小写: 私有

const PackageConst = 100

func init() {
    // 包初始化时执行
    // 可以访问包级变量
}

// ==================== 闭包与作用域 ====================
func closureExample() {
    count := 0
    
    increment := func() int {
        count++  // 捕获外层变量
        return count
    }
    
    fmt.Println(increment())  // 1
    fmt.Println(increment())  // 2
}

// 循环闭包陷阱
func loopTrap() {
    funcs := make([]func(), 3)
    
    for i := 0; i < 3; i++ {
        funcs[i] = func() {
            fmt.Println(i)  // 都引用同一个 i
        }
    }
    
    for _, f := range funcs {
        f()  // 3, 3, 3
    }
}

// 解决方案 1: 参数传递
func loopFixed1() {
    funcs := make([]func(), 3)
    
    for i := 0; i < 3; i++ {
        funcs[i] = func(n int) func() {
            return func() { fmt.Println(n) }
        }(i)
    }
}

// 解决方案 2: 局部变量
func loopFixed2() {
    funcs := make([]func(), 3)
    
    for i := 0; i < 3; i++ {
        i := i  // 创建新的局部变量
        funcs[i] = func() {
            fmt.Println(i)
        }
    }
}

// Go 1.22+ 循环变量语义改变
// for i := 0; i < 3; i++ 的 i 在每次迭代都是新变量
```

### Rust 作用域

```rust
// ==================== 块级作用域 ====================
fn block_scope() {
    let x = 10;
    
    {
        let y = 20;  // 只在块内可见
        let x = 30;  // 遮蔽外层 x
        println!("{} {}", x, y);  // 30 20
    }  // y 在这里被 drop
    
    // println!("{}", y);  // 编译错误
    println!("{}", x);  // 10
}

// ==================== 变量遮蔽 (Shadowing) ====================
fn shadowing() {
    let x = 5;
    let x = x + 1;  // 遮蔽，可以改变类型
    
    {
        let x = x * 2;
        println!("{}", x);  // 12
    }
    
    println!("{}", x);  // 6
    
    // 遮蔽可以改变类型
    let spaces = "   ";
    let spaces = spaces.len();  // 从 &str 变成 usize
}

// 与 mut 的区别
fn shadowing_vs_mut() {
    // 遮蔽: 创建新变量
    let x = 5;
    let x = "hello";  // OK，新类型
    
    // mut: 可变绑定
    let mut y = 5;
    y = 10;  // OK，相同类型
    // y = "hello";  // 编译错误，类型不匹配
}

// ==================== 块表达式 ====================
fn block_expression() {
    // 块可以返回值
    let x = {
        let a = 1;
        let b = 2;
        a + b  // 无分号 = 返回值
    };
    println!("{}", x);  // 3
    
    // if 是表达式
    let y = if true { 1 } else { 2 };
    
    // match 是表达式
    let z = match y {
        1 => "one",
        _ => "other",
    };
    
    // loop 可以返回值
    let result = loop {
        break 42;
    };
}

// ==================== 所有权与作用域 ====================
fn ownership_scope() {
    let s1 = String::from("hello");
    
    {
        let s2 = s1;  // s1 移动到 s2
        println!("{}", s2);
    }  // s2 被 drop，内存释放
    
    // println!("{}", s1);  // 编译错误，s1 已移动
}

// 引用的生命周期与作用域
fn reference_scope() {
    let r;
    
    {
        let x = 5;
        r = &x;
        println!("{}", r);  // OK
    }  // x 被 drop
    
    // println!("{}", r);  // 编译错误，x 已不存在
}

// ==================== 非词法生命周期 (NLL) ====================
fn nll_example() {
    let mut data = vec![1, 2, 3];
    
    let first = &data[0];  // 不可变借用
    println!("{}", first);
    // first 的生命周期在这里结束 (NLL)
    
    data.push(4);  // 可变借用，NLL 使这成为可能
}

// ==================== 静态与常量 ====================
// 全局静态变量
static GLOBAL: &str = "全局静态";
static mut MUTABLE_GLOBAL: i32 = 0;  // 需要 unsafe 访问

// 常量 (编译时求值)
const MAX_SIZE: usize = 100;

fn static_example() {
    println!("{}", GLOBAL);
    
    unsafe {
        MUTABLE_GLOBAL += 1;  // 必须 unsafe
    }
}

// ==================== 闭包捕获 ====================
fn closure_capture() {
    let x = 10;
    let y = String::from("hello");
    
    // 不可变借用
    let borrow = || println!("{} {}", x, y);
    borrow();
    println!("{}", y);  // y 仍可用
    
    // 可变借用
    let mut count = 0;
    let mut increment = || count += 1;
    increment();
    increment();
    println!("{}", count);  // 2
    
    // move: 获取所有权
    let owned = move || println!("{}", y);
    owned();
    // println!("{}", y);  // 编译错误，y 已移动
}

// ==================== Drop 顺序 ====================
struct Droppable(i32);

impl Drop for Droppable {
    fn drop(&mut self) {
        println!("Dropping {}", self.0);
    }
}

fn drop_order() {
    let a = Droppable(1);
    let b = Droppable(2);
    let c = Droppable(3);
    
    {
        let d = Droppable(4);
        let e = Droppable(5);
    }  // 输出: Dropping 5, Dropping 4 (后声明先 drop)
    
    println!("End of function");
}  // 输出: Dropping 3, Dropping 2, Dropping 1
```

### 作用域对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 特性            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 块作用域        │ let/const            │ ❌                   │ ✅                   │ ✅                   │
│ 函数作用域      │ var                  │ ✅                   │ ✅                   │ ✅                   │
│ 变量提升        │ var                  │ ❌                   │ ❌                   │ ❌                   │
│ 循环变量        │ let 每次新建         │ 共享/泄漏            │ 共享 (1.22 前)       │ 每次新建             │
│ 遮蔽改类型      │ ❌ (需要断言)        │ ✅                   │ ❌                   │ ✅                   │
│ 块返回值        │ ❌                   │ ❌                   │ ❌                   │ ✅                   │
│ 全局声明        │ 顶层/window          │ 模块顶层             │ 包级 var             │ static/const         │
│ 修改外层变量    │ 直接修改             │ global/nonlocal      │ 直接修改             │ &mut                 │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 常见陷阱与最佳实践

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ TypeScript                                                                  │
│ • 始终使用 let/const，避免 var                                              │
│ • 注意 for 循环中 var 的闭包问题                                            │
│ • const 不能防止对象属性修改                                                │
├─────────────────────────────────────────────────────────────────────────────┤
│ Python                                                                      │
│ • 记住没有块级作用域，if/for 内的变量会泄漏                                 │
│ • 闭包捕获变量引用，循环中用默认参数捕获值                                  │
│ • 修改全局变量需要 global，闭包变量需要 nonlocal                            │
├─────────────────────────────────────────────────────────────────────────────┤
│ Go                                                                          │
│ • 注意 := 短声明可能意外遮蔽变量                                            │
│ • Go 1.22 前循环变量是共享的，闭包需要复制                                  │
│ • 包级变量用于共享状态，但要注意并发安全                                    │
├─────────────────────────────────────────────────────────────────────────────┤
│ Rust                                                                        │
│ • 遮蔽是惯用法，常用于类型转换                                              │
│ • 理解所有权与作用域的关系                                                  │
│ • 块表达式是强大的工具，善用返回值                                          │
└─────────────────────────────────────────────────────────────────────────────┘
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

## 📨 函数参数传递规则

### 参数传递概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 基本类型 | 值传递 | 值传递(不可变) | 值传递 | 值传递(Move/Copy) |
| 对象/结构体 | 引用传递 | 引用传递 | 值传递(复制) | Move(默认)/借用 |
| 数组 | 引用传递 | 引用传递 | 值传递(复制) | Move/借用 |
| 显式引用 | ❌ | ❌ | `*T` 指针 | `&T` / `&mut T` |
| 修改原值 | 可以(对象) | 可以(可变对象) | 需要指针 | 需要 `&mut` |

### 传递方式图解

```
┌─────────────────────────────────────────────────────────────────────────────┐
│                           参数传递方式                                        │
├─────────────────────────────────────────────────────────────────────────────┤
│                                                                             │
│  值传递 (Copy)                     引用传递 (Reference)                      │
│  ┌─────────┐                      ┌─────────┐                               │
│  │ 原始值   │                      │ 原始值   │                               │
│  │   42    │                      │ {a: 1}  │◄─────────┐                    │
│  └────┬────┘                      └─────────┘          │                    │
│       │ 复制                                           │ 指向                │
│       ▼                                                │                    │
│  ┌─────────┐                      ┌─────────┐          │                    │
│  │ 函数参数 │                      │ 函数参数 │──────────┘                    │
│  │   42    │                      │   ref   │                               │
│  └─────────┘                      └─────────┘                               │
│  修改不影响原值                     修改会影响原值                             │
│                                                                             │
│  Move (Rust)                      Borrow (Rust)                             │
│  ┌─────────┐                      ┌─────────┐                               │
│  │ 原始值   │ ──转移──►            │ 原始值   │                               │
│  │ (失效)  │                      │ String  │◄─────────┐                    │
│  └─────────┘                      └─────────┘          │ 借用               │
│                                                        │                    │
│  ┌─────────┐                      ┌─────────┐          │                    │
│  │ 函数参数 │                      │ 函数参数 │──────────┘                    │
│  │ String  │ (拥有所有权)          │   &str  │ (只读借用)                     │
│  └─────────┘                      └─────────┘                               │
│                                                                             │
└─────────────────────────────────────────────────────────────────────────────┘
```

### TypeScript 参数传递

```typescript
// ==================== 基本类型 - 值传递 ====================
function modifyNumber(x: number): void {
    x = 100;  // 不影响原值
}

let num = 42;
modifyNumber(num);
console.log(num);  // 42 - 未改变

// 同样适用于 string, boolean, bigint, symbol
function modifyString(s: string): void {
    s = "modified";  // 不影响原值
}

let str = "original";
modifyString(str);
console.log(str);  // "original"

// ==================== 对象 - 引用传递 ====================
function modifyObject(obj: { value: number }): void {
    obj.value = 100;  // 修改原对象
}

const myObj = { value: 42 };
modifyObject(myObj);
console.log(myObj.value);  // 100 - 已改变

// 重新赋值不影响原引用
function reassignObject(obj: { value: number }): void {
    obj = { value: 999 };  // 创建新对象，不影响原引用
}

const myObj2 = { value: 42 };
reassignObject(myObj2);
console.log(myObj2.value);  // 42 - 未改变

// ==================== 数组 - 引用传递 ====================
function modifyArray(arr: number[]): void {
    arr.push(4);      // 修改原数组
    arr[0] = 100;     // 修改原数组
}

const myArr = [1, 2, 3];
modifyArray(myArr);
console.log(myArr);  // [100, 2, 3, 4]

// 重新赋值不影响原引用
function reassignArray(arr: number[]): void {
    arr = [7, 8, 9];  // 不影响原数组
}

const myArr2 = [1, 2, 3];
reassignArray(myArr2);
console.log(myArr2);  // [1, 2, 3]

// ==================== 防止修改 - 浅拷贝 ====================
function safeModify(obj: { value: number }): { value: number } {
    const copy = { ...obj };  // 浅拷贝
    copy.value = 100;
    return copy;
}

// 深拷贝
function deepCopy<T>(obj: T): T {
    return JSON.parse(JSON.stringify(obj));
}

function safeDeepModify<T>(obj: T): T {
    const copy = deepCopy(obj);
    // 修改 copy...
    return copy;
}

// ==================== readonly 防止修改 ====================
function readOnly(arr: readonly number[]): void {
    // arr.push(4);    // 编译错误
    // arr[0] = 100;   // 编译错误
    console.log(arr[0]);  // 只能读取
}

interface ReadonlyUser {
    readonly id: number;
    readonly name: string;
}

function processUser(user: ReadonlyUser): void {
    // user.name = "new";  // 编译错误
}

// Readonly<T> 工具类型
function immutableProcess(data: Readonly<{ x: number; y: number }>): void {
    // data.x = 10;  // 编译错误
}

// ==================== 类实例 - 引用传递 ====================
class Counter {
    count = 0;
    increment() { this.count++; }
}

function modifyCounter(c: Counter): void {
    c.increment();  // 修改原实例
}

const counter = new Counter();
modifyCounter(counter);
console.log(counter.count);  // 1
```

### Python 参数传递

```python
# ==================== 不可变类型 - 值传递语义 ====================
# int, float, str, tuple, frozenset 是不可变的
def modify_number(x: int) -> None:
    x = 100  # 创建新对象，不影响原值

num = 42
modify_number(num)
print(num)  # 42 - 未改变

def modify_string(s: str) -> None:
    s = "modified"  # 创建新对象

text = "original"
modify_string(text)
print(text)  # "original"

# ==================== 可变类型 - 引用传递 ====================
# list, dict, set, 自定义类 是可变的
def modify_list(lst: list) -> None:
    lst.append(4)    # 修改原列表
    lst[0] = 100     # 修改原列表

my_list = [1, 2, 3]
modify_list(my_list)
print(my_list)  # [100, 2, 3, 4]

def modify_dict(d: dict) -> None:
    d['new_key'] = 'value'  # 修改原字典

my_dict = {'a': 1}
modify_dict(my_dict)
print(my_dict)  # {'a': 1, 'new_key': 'value'}

# 重新赋值不影响原引用
def reassign_list(lst: list) -> None:
    lst = [7, 8, 9]  # 局部变量指向新对象

my_list2 = [1, 2, 3]
reassign_list(my_list2)
print(my_list2)  # [1, 2, 3] - 未改变

# ==================== 可变默认参数陷阱 ====================
# 错误示例！
def bad_append(item, lst=[]):  # 默认列表在所有调用间共享
    lst.append(item)
    return lst

bad_append(1)  # [1]
bad_append(2)  # [1, 2] - 不是 [2]！

# 正确做法
def good_append(item, lst=None):
    if lst is None:
        lst = []
    lst.append(item)
    return lst

# ==================== 防止修改 - 拷贝 ====================
import copy

def safe_modify(lst: list) -> list:
    copy_lst = lst.copy()  # 浅拷贝
    # 或 copy_lst = lst[:]
    # 或 copy_lst = list(lst)
    copy_lst.append(4)
    return copy_lst

def deep_safe_modify(data: dict) -> dict:
    copy_data = copy.deepcopy(data)  # 深拷贝
    # 修改 copy_data...
    return copy_data

# ==================== 类型提示表达意图 ====================
from typing import List, Sequence, MutableSequence

# Sequence - 暗示不修改
def read_only_process(items: Sequence[int]) -> int:
    return sum(items)

# MutableSequence - 明确会修改
def will_modify(items: MutableSequence[int]) -> None:
    items.append(0)

# ==================== 类实例 ====================
class User:
    def __init__(self, name: str):
        self.name = name

def modify_user(user: User) -> None:
    user.name = "Modified"  # 修改原实例

u = User("Original")
modify_user(u)
print(u.name)  # "Modified"

# ==================== dataclass 与不可变 ====================
from dataclasses import dataclass

@dataclass
class Point:
    x: int
    y: int

@dataclass(frozen=True)  # 不可变
class ImmutablePoint:
    x: int
    y: int

def try_modify(p: ImmutablePoint) -> None:
    # p.x = 10  # 运行时错误: FrozenInstanceError
    pass

# ==================== id() 查看对象标识 ====================
def show_id(x):
    print(f"Inside function: id = {id(x)}")

num = 42
print(f"Before: id = {id(num)}")
show_id(num)  # 相同 id（小整数缓存）

lst = [1, 2, 3]
print(f"Before: id = {id(lst)}")
show_id(lst)  # 相同 id - 同一个对象
```

### Go 参数传递

```go
// ==================== 基本类型 - 值传递 ====================
func modifyInt(x int) {
    x = 100  // 不影响原值
}

func main() {
    num := 42
    modifyInt(num)
    fmt.Println(num)  // 42 - 未改变
}

// ==================== 结构体 - 值传递(复制) ====================
type User struct {
    Name string
    Age  int
}

func modifyUser(u User) {
    u.Name = "Modified"  // 修改的是副本
}

func main() {
    user := User{Name: "Original", Age: 30}
    modifyUser(user)
    fmt.Println(user.Name)  // "Original" - 未改变
}

// ==================== 指针 - 引用传递 ====================
func modifyUserPtr(u *User) {
    u.Name = "Modified"  // 修改原结构体
}

func main() {
    user := User{Name: "Original", Age: 30}
    modifyUserPtr(&user)
    fmt.Println(user.Name)  // "Modified" - 已改变
}

// ==================== 数组 - 值传递(复制) ====================
func modifyArray(arr [3]int) {
    arr[0] = 100  // 修改的是副本
}

func main() {
    arr := [3]int{1, 2, 3}
    modifyArray(arr)
    fmt.Println(arr)  // [1 2 3] - 未改变
}

// ==================== 切片 - 引用语义 ====================
// 切片本身是值传递，但底层数组是共享的
func modifySlice(s []int) {
    s[0] = 100        // 修改原数组
    s = append(s, 4)  // 可能创建新数组，不影响原切片
}

func main() {
    slice := []int{1, 2, 3}
    modifySlice(slice)
    fmt.Println(slice)  // [100 2 3] - 元素已改变，但长度不变
}

// 要修改切片本身（长度、容量），需要指针
func appendSlice(s *[]int, val int) {
    *s = append(*s, val)
}

func main() {
    slice := []int{1, 2, 3}
    appendSlice(&slice, 4)
    fmt.Println(slice)  // [1 2 3 4]
}

// ==================== Map - 引用语义 ====================
func modifyMap(m map[string]int) {
    m["new"] = 100  // 修改原 map
}

func main() {
    m := map[string]int{"a": 1}
    modifyMap(m)
    fmt.Println(m)  // map[a:1 new:100]
}

// ==================== Channel - 引用语义 ====================
func sendToChannel(ch chan int) {
    ch <- 42  // 发送到原 channel
}

// ==================== 接口 - 取决于底层类型 ====================
func modifyInterface(i interface{}) {
    // 需要类型断言才能修改
    if ptr, ok := i.(*User); ok {
        ptr.Name = "Modified"
    }
}

// ==================== 防止修改 - 返回新值 ====================
func safeModifyUser(u User) User {
    u.Name = "Modified"
    return u  // 返回修改后的副本
}

// 深拷贝结构体
func deepCopyUser(u *User) User {
    return User{
        Name: u.Name,
        Age:  u.Age,
    }
}

// ==================== 何时使用指针 ====================
// 1. 需要修改原值
func (u *User) SetName(name string) {
    u.Name = name
}

// 2. 大结构体避免复制开销
type LargeStruct struct {
    Data [1000000]int
}

func processLarge(ls *LargeStruct) {
    // 只传递指针，不复制整个数组
}

// 3. 表示可选值 (nil)
func findUser(id int) *User {
    if id == 0 {
        return nil
    }
    return &User{Name: "Found"}
}

// ==================== 接收者类型选择 ====================
type Counter struct {
    count int
}

// 值接收者 - 不修改原值
func (c Counter) Value() int {
    return c.count
}

// 指针接收者 - 修改原值
func (c *Counter) Increment() {
    c.count++
}
```

### Rust 参数传递

```rust
// ==================== 所有权转移 (Move) ====================
fn take_ownership(s: String) {
    println!("{}", s);
}  // s 在这里被 drop

fn main() {
    let s = String::from("hello");
    take_ownership(s);
    // println!("{}", s);  // 编译错误！s 已被移动
}

// ==================== Copy 类型 - 自动复制 ====================
// 实现了 Copy trait 的类型会自动复制
fn use_number(x: i32) {
    println!("{}", x);
}

fn main() {
    let num = 42;
    use_number(num);
    println!("{}", num);  // OK - i32 实现了 Copy
}

// Copy 类型包括: i32, f64, bool, char, 元组(如果元素都是Copy)
// 非 Copy: String, Vec, Box, 自定义结构体(默认)

// ==================== 借用 (Borrow) - 不可变引用 ====================
fn borrow_string(s: &String) {
    println!("{}", s);
}  // 借用结束，不会 drop

fn main() {
    let s = String::from("hello");
    borrow_string(&s);
    println!("{}", s);  // OK - 只是借用
}

// 更好的做法：使用切片
fn borrow_str(s: &str) {
    println!("{}", s);
}

fn main() {
    let s = String::from("hello");
    borrow_str(&s);      // String 可以自动解引用为 &str
    borrow_str("world"); // 字符串字面量也可以
}

// ==================== 可变借用 (Mutable Borrow) ====================
fn modify_string(s: &mut String) {
    s.push_str(" world");
}

fn main() {
    let mut s = String::from("hello");
    modify_string(&mut s);
    println!("{}", s);  // "hello world"
}

// ==================== 借用规则 ====================
fn main() {
    let mut s = String::from("hello");
    
    // 规则1: 多个不可变借用 OK
    let r1 = &s;
    let r2 = &s;
    println!("{} {}", r1, r2);
    
    // 规则2: 一个可变借用，不能有其他借用
    let r3 = &mut s;
    // let r4 = &s;       // 编译错误！
    // let r5 = &mut s;   // 编译错误！
    println!("{}", r3);
}

// ==================== 结构体所有权 ====================
struct User {
    name: String,
    age: u32,
}

// 获取所有权
fn take_user(user: User) {
    println!("{}", user.name);
}  // user 被 drop

// 借用
fn borrow_user(user: &User) {
    println!("{}", user.name);
}

// 可变借用
fn modify_user(user: &mut User) {
    user.name = String::from("Modified");
}

// ==================== Clone - 显式复制 ====================
fn main() {
    let s1 = String::from("hello");
    let s2 = s1.clone();  // 显式深拷贝
    
    take_ownership(s1);
    println!("{}", s2);  // OK - s2 是独立的副本
}

// ==================== 实现 Copy trait ====================
#[derive(Copy, Clone)]
struct Point {
    x: i32,
    y: i32,
}

fn use_point(p: Point) {
    println!("({}, {})", p.x, p.y);
}

fn main() {
    let p = Point { x: 1, y: 2 };
    use_point(p);
    println!("{}", p.x);  // OK - Point 实现了 Copy
}

// 注意：包含非 Copy 字段的结构体不能实现 Copy
// struct Invalid {
//     name: String,  // String 没有实现 Copy
// }
// #[derive(Copy)]  // 编译错误！

// ==================== 方法中的 self ====================
struct Counter {
    count: i32,
}

impl Counter {
    // 获取所有权
    fn consume(self) -> i32 {
        self.count
    }  // self 被 drop
    
    // 不可变借用
    fn get(&self) -> i32 {
        self.count
    }
    
    // 可变借用
    fn increment(&mut self) {
        self.count += 1;
    }
}

fn main() {
    let mut c = Counter { count: 0 };
    println!("{}", c.get());       // 借用
    c.increment();                  // 可变借用
    println!("{}", c.get());       // 借用
    let final_count = c.consume(); // 移动
    // c.get();  // 编译错误！c 已被移动
}

// ==================== 切片借用 ====================
fn sum(slice: &[i32]) -> i32 {
    slice.iter().sum()
}

fn main() {
    let arr = [1, 2, 3, 4, 5];
    let vec = vec![1, 2, 3, 4, 5];
    
    println!("{}", sum(&arr));     // 数组切片
    println!("{}", sum(&vec));     // Vec 切片
    println!("{}", sum(&arr[1..4])); // 部分切片
}

// ==================== 返回引用需要生命周期 ====================
fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
    if x.len() > y.len() { x } else { y }
}

fn main() {
    let s1 = String::from("hello");
    let s2 = String::from("world!");
    let result = longest(&s1, &s2);
    println!("{}", result);
}

// ==================== Cow - 写时复制 ====================
use std::borrow::Cow;

fn process(input: &str) -> Cow<str> {
    if input.contains("bad") {
        // 需要修改时才分配
        Cow::Owned(input.replace("bad", "good"))
    } else {
        // 不需要修改，返回借用
        Cow::Borrowed(input)
    }
}
```

### 参数传递规则对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 类型            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 数字            │ 值传递               │ 不可变(值语义)       │ 值传递               │ Copy                 │
│ 字符串          │ 值传递               │ 不可变(值语义)       │ 值传递               │ Move / &str          │
│ 布尔            │ 值传递               │ 不可变(值语义)       │ 值传递               │ Copy                 │
│ 数组/列表       │ 引用传递             │ 引用传递(可变)       │ 值传递(复制)         │ Move / &[T]          │
│ 对象/结构体     │ 引用传递             │ 引用传递(可变)       │ 值传递(复制)         │ Move / &T            │
│ Map/Dict        │ 引用传递             │ 引用传递(可变)       │ 引用传递             │ Move / &HashMap      │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 修改原值        │ 对象字段可改         │ 可变对象可改         │ 需要指针 *T          │ 需要 &mut T          │
│ 防止修改        │ readonly / 拷贝      │ 拷贝 / frozen        │ 不传指针             │ 不传 &mut            │
│ 显式复制        │ {...} / 深拷贝       │ copy.deepcopy        │ 手动赋值             │ .clone()             │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 最佳实践

| 场景 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| **只读访问** | 直接传递 | 直接传递 | 值传递 | `&T` 借用 |
| **需要修改** | 传对象 | 传可变对象 | 传指针 `*T` | `&mut T` |
| **避免大对象复制** | 默认引用 | 默认引用 | 传指针 | 传引用 |
| **防止意外修改** | `readonly` / 拷贝 | `frozen` / 拷贝 | 值传递 | 只传 `&T` |
| **转移所有权** | N/A | N/A | N/A | 直接传递(move) |
| **可选参数** | `?` / `undefined` | `= None` | `*T` (nil) | `Option<T>` |

### 常见陷阱

```
┌─────────────────────────────────────────────────────────────────────────────┐
│ TypeScript: 以为修改了原值，实际重新赋值了引用                                │
│   function f(arr) { arr = [1,2,3]; }  // 不影响原数组                       │
│                                                                             │
│ Python: 可变默认参数在所有调用间共享                                         │
│   def f(lst=[]):  // 错误！应该用 lst=None                                  │
│                                                                             │
│ Go: 以为切片会自动扩容影响原切片                                             │
│   func f(s []int) { s = append(s, 1) }  // 可能不影响原切片                 │
│                                                                             │
│ Rust: 忘记所有权已转移                                                       │
│   let s = String::from("hi");                                               │
│   take(s);                                                                  │
│   println!("{}", s);  // 编译错误！                                         │
└─────────────────────────────────────────────────────────────────────────────┘
```

---

## 📝 函数默认参数

### 默认参数支持概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 默认参数 | ✅ | ✅ | ❌ | ❌ |
| 命名参数 | ❌ (对象模拟) | ✅ | ❌ | ❌ |
| 可选参数 | ✅ `?` | ✅ `= None` | ✅ 指针/变参 | ✅ `Option<T>` |
| 变长参数 | ✅ `...` | ✅ `*args` | ✅ `...` | ❌ (宏实现) |
| 关键字参数 | ❌ | ✅ `**kwargs` | ❌ | ❌ |

### TypeScript 默认参数

```typescript
// ==================== 基础默认参数 ====================
function greet(name: string, greeting: string = "Hello"): string {
  return `${greeting}, ${name}!`;
}

greet("John");              // "Hello, John!"
greet("John", "Hi");        // "Hi, John!"

// ==================== 默认参数可以是表达式 ====================
function createId(prefix: string = "id", timestamp: number = Date.now()): string {
  return `${prefix}_${timestamp}`;
}

// ==================== 可选参数 (与默认参数不同) ====================
function log(message: string, level?: string): void {
  // level 类型是 string | undefined
  console.log(`[${level ?? "INFO"}] ${message}`);
}

// ==================== 默认参数必须在必选参数之后 ====================
// function bad(a = 1, b: number) {}  // ❌ 错误

// 但可以用 undefined 触发默认值
function example(a: string = "default", b: number): void {
  console.log(a, b);
}
example(undefined, 42);     // "default" 42

// ==================== 解构默认值 ====================
function config({ host = "localhost", port = 3000 } = {}): void {
  console.log(`${host}:${port}`);
}

config();                   // localhost:3000
config({ port: 8080 });     // localhost:8080

// ==================== 模拟命名参数 (推荐) ====================
interface Options {
  name: string;
  age?: number;
  email?: string;
}

function createUser({ name, age = 18, email = "" }: Options): void {
  console.log(name, age, email);
}

createUser({ name: "John" });
createUser({ name: "John", age: 25, email: "john@example.com" });

// ==================== 剩余参数 ====================
function sum(...numbers: number[]): number {
  return numbers.reduce((a, b) => a + b, 0);
}

sum(1, 2, 3);               // 6

// ==================== 函数重载模拟不同参数 ====================
function format(value: string): string;
function format(value: number, decimals?: number): string;
function format(value: string | number, decimals?: number): string {
  if (typeof value === "string") {
    return value.toUpperCase();
  }
  return value.toFixed(decimals ?? 2);
}
```

### Python 默认参数

```python
# ==================== 基础默认参数 ====================
def greet(name: str, greeting: str = "Hello") -> str:
    return f"{greeting}, {name}!"

greet("John")               # "Hello, John!"
greet("John", "Hi")         # "Hi, John!"
greet("John", greeting="Hi")  # 命名参数调用

# ==================== ⚠️ 可变默认参数陷阱 ====================
# 错误示范 - 默认值在函数定义时创建，所有调用共享
def bad_append(item, lst=[]):  # ❌ 危险！
    lst.append(item)
    return lst

bad_append(1)               # [1]
bad_append(2)               # [1, 2] 不是 [2]！

# 正确做法 - 使用 None 作为哨兵
def good_append(item, lst=None):
    if lst is None:
        lst = []
    lst.append(item)
    return lst

# ==================== 仅位置参数 (Python 3.8+) ====================
def func(pos_only, /, standard, *, kw_only):
    pass

# func(1, 2, 3)             # ❌ 错误：kw_only 必须用关键字
func(1, 2, kw_only=3)       # ✅
func(1, standard=2, kw_only=3)  # ✅

# ==================== *args 和 **kwargs ====================
def flexible(*args, **kwargs):
    print(f"args: {args}")
    print(f"kwargs: {kwargs}")

flexible(1, 2, 3, name="John", age=25)
# args: (1, 2, 3)
# kwargs: {'name': 'John', 'age': 25}

# ==================== 解包传参 ====================
def add(a, b, c):
    return a + b + c

args = [1, 2, 3]
add(*args)                  # 6

kwargs = {"a": 1, "b": 2, "c": 3}
add(**kwargs)               # 6

# ==================== 类型提示的可选参数 ====================
from typing import Optional

def process(data: str, timeout: Optional[int] = None) -> str:
    if timeout is None:
        timeout = 30  # 默认超时
    return f"Processing {data} with timeout {timeout}"

# ==================== dataclass 默认值 ====================
from dataclasses import dataclass, field

@dataclass
class Config:
    host: str = "localhost"
    port: int = 8080
    tags: list = field(default_factory=list)  # 可变类型用 field

# ==================== 函数重载 (类型提示) ====================
from typing import overload

@overload
def parse(value: str) -> str: ...
@overload
def parse(value: int) -> int: ...

def parse(value):
    if isinstance(value, str):
        return value.upper()
    return value * 2
```

### Go 模拟默认参数

```go
// Go 不支持默认参数，需要使用其他模式

// ==================== 方式 1: 函数重载模拟 (多个函数) ====================
func Greet(name string) string {
    return GreetWith(name, "Hello")
}

func GreetWith(name, greeting string) string {
    return greeting + ", " + name + "!"
}

// ==================== 方式 2: 可变参数 ====================
func Connect(host string, ports ...int) {
    port := 3306  // 默认值
    if len(ports) > 0 {
        port = ports[0]
    }
    fmt.Printf("Connecting to %s:%d\n", host, port)
}

Connect("localhost")        // 使用默认端口
Connect("localhost", 5432)  // 指定端口

// ==================== 方式 3: 选项结构体 (推荐) ====================
type Options struct {
    Host    string
    Port    int
    Timeout time.Duration
    Retry   int
}

// 默认值
func defaultOptions() Options {
    return Options{
        Host:    "localhost",
        Port:    8080,
        Timeout: 30 * time.Second,
        Retry:   3,
    }
}

func NewClient(opts ...Options) *Client {
    // 使用默认值
    opt := defaultOptions()
    if len(opts) > 0 {
        // 覆盖提供的值
        if opts[0].Host != "" {
            opt.Host = opts[0].Host
        }
        if opts[0].Port != 0 {
            opt.Port = opts[0].Port
        }
        // ... 其他字段
    }
    return &Client{options: opt}
}

// 使用
NewClient()                             // 全部默认
NewClient(Options{Port: 9000})          // 部分覆盖

// ==================== 方式 4: 函数式选项 (最灵活) ====================
type ClientOption func(*Client)

func WithHost(host string) ClientOption {
    return func(c *Client) {
        c.host = host
    }
}

func WithPort(port int) ClientOption {
    return func(c *Client) {
        c.port = port
    }
}

func WithTimeout(d time.Duration) ClientOption {
    return func(c *Client) {
        c.timeout = d
    }
}

func NewClient(options ...ClientOption) *Client {
    // 默认值
    c := &Client{
        host:    "localhost",
        port:    8080,
        timeout: 30 * time.Second,
    }
    // 应用选项
    for _, opt := range options {
        opt(c)
    }
    return c
}

// 使用 - 非常清晰
NewClient()
NewClient(WithHost("example.com"))
NewClient(WithHost("example.com"), WithPort(9000), WithTimeout(time.Minute))

// ==================== 方式 5: 指针表示可选 ====================
func Process(required string, optional *int) {
    value := 100  // 默认值
    if optional != nil {
        value = *optional
    }
    fmt.Println(required, value)
}

// 辅助函数
func IntPtr(v int) *int { return &v }

Process("data", nil)        // 使用默认值
Process("data", IntPtr(50)) // 指定值
```

### Rust 模拟默认参数

```rust
// Rust 不支持默认参数，需要使用其他模式

// ==================== 方式 1: Option 类型 ====================
fn greet(name: &str, greeting: Option<&str>) -> String {
    let g = greeting.unwrap_or("Hello");
    format!("{}, {}!", g, name)
}

greet("John", None);            // "Hello, John!"
greet("John", Some("Hi"));      // "Hi, John!"

// ==================== 方式 2: Default trait ====================
#[derive(Default)]
struct Config {
    host: String,
    port: u16,
    timeout: u64,
}

impl Default for Config {
    fn default() -> Self {
        Config {
            host: String::from("localhost"),
            port: 8080,
            timeout: 30,
        }
    }
}

// 使用 struct update 语法
let config = Config {
    port: 9000,
    ..Default::default()
};

// ==================== 方式 3: Builder 模式 (推荐) ====================
struct Client {
    host: String,
    port: u16,
    timeout: u64,
}

struct ClientBuilder {
    host: String,
    port: u16,
    timeout: u64,
}

impl ClientBuilder {
    fn new() -> Self {
        ClientBuilder {
            host: String::from("localhost"),
            port: 8080,
            timeout: 30,
        }
    }
    
    fn host(mut self, host: &str) -> Self {
        self.host = host.to_string();
        self
    }
    
    fn port(mut self, port: u16) -> Self {
        self.port = port;
        self
    }
    
    fn timeout(mut self, timeout: u64) -> Self {
        self.timeout = timeout;
        self
    }
    
    fn build(self) -> Client {
        Client {
            host: self.host,
            port: self.port,
            timeout: self.timeout,
        }
    }
}

// 使用 - 链式调用
let client = ClientBuilder::new()
    .host("example.com")
    .port(9000)
    .build();

// ==================== 方式 4: 使用 derive_builder crate ====================
// Cargo.toml: derive_builder = "0.12"

use derive_builder::Builder;

#[derive(Builder, Default)]
#[builder(setter(into))]
struct Request {
    #[builder(default = "\"localhost\".to_string()")]
    host: String,
    #[builder(default = "8080")]
    port: u16,
    #[builder(default)]
    headers: Vec<String>,
}

let req = RequestBuilder::default()
    .host("example.com")
    .build()
    .unwrap();

// ==================== 方式 5: 函数重载 (多个函数) ====================
fn connect(host: &str) -> Connection {
    connect_with_port(host, 8080)
}

fn connect_with_port(host: &str, port: u16) -> Connection {
    connect_with_options(host, port, 30)
}

fn connect_with_options(host: &str, port: u16, timeout: u64) -> Connection {
    // 实际连接逻辑
    Connection { host: host.to_string(), port, timeout }
}

// ==================== 方式 6: 宏实现可变参数 ====================
macro_rules! greet {
    ($name:expr) => {
        greet!($name, "Hello")
    };
    ($name:expr, $greeting:expr) => {
        format!("{}, {}!", $greeting, $name)
    };
}

greet!("John");             // "Hello, John!"
greet!("John", "Hi");       // "Hi, John!"

// ==================== 方式 7: Into trait 灵活参数 ====================
fn process<S: Into<String>>(data: S) {
    let s: String = data.into();
    println!("{}", s);
}

process("literal");         // &str
process(String::from("owned"));  // String
```

### 默认参数模式对比

```
┌─────────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 模式                │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 原生默认参数        │ ✅ param = val │ ✅ param = val │ ❌             │ ❌             │
│ 可选参数            │ param?         │ param = None   │ *type / ...    │ Option<T>      │
│ 命名参数            │ 对象解构       │ ✅ 原生支持    │ 结构体         │ 结构体/Builder │
│ 变长参数            │ ...rest        │ *args          │ ...type        │ 宏             │
│ 关键字参数          │ ❌             │ **kwargs       │ ❌             │ ❌             │
├─────────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 推荐的复杂参数模式  │ Options 对象   │ dataclass      │ Functional Opt │ Builder        │
│ 可变默认值陷阱      │ ❌ 无          │ ⚠️ 有！        │ ❌ 无          │ ❌ 无          │
└─────────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 最佳实践总结

| 语言 | 推荐方式 | 说明 |
|------|----------|------|
| TypeScript | 默认参数 + 解构对象 | 对于简单情况用默认参数，复杂配置用 Options 接口 |
| Python | 默认参数 + 命名参数 | 避免可变默认值，使用 `None` + 类型提示 |
| Go | Functional Options | 最灵活、清晰，适合公共 API |
| Rust | Builder Pattern | 类型安全、链式调用、编译期检查 |

---

## 🧬 泛型函数

### 泛型支持概览

| 语言 | 泛型支持 | 语法 | 约束方式 | 特点 |
|------|----------|------|----------|------|
| TypeScript | ✅ | `<T>` | `extends` | 类型擦除 |
| Python | ✅ (3.12+) | `[T]` | `TypeVar` bound | 仅类型检查 |
| Go | ✅ (1.18+) | `[T]` | 接口约束 | 编译时单态化 |
| Rust | ✅ | `<T>` | trait bound | 编译时单态化 |

### TypeScript 泛型

```typescript
// ==================== 基础泛型函数 ====================

function identity<T>(value: T): T {
  return value;
}

// 调用
identity<string>("hello");  // 显式指定
identity("hello");          // 类型推断
identity(42);               // number

// ==================== 多类型参数 ====================

function pair<T, U>(first: T, second: U): [T, U] {
  return [first, second];
}

function map<T, U>(arr: T[], fn: (item: T) => U): U[] {
  return arr.map(fn);
}

// ==================== 类型约束 (extends) ====================

// 必须有 length 属性
function logLength<T extends { length: number }>(value: T): T {
  console.log(value.length);
  return value;
}

logLength("hello");     // OK
logLength([1, 2, 3]);   // OK
// logLength(123);      // 错误: number 没有 length

// 约束为特定类型
interface HasId {
  id: number;
}

function findById<T extends HasId>(items: T[], id: number): T | undefined {
  return items.find(item => item.id === id);
}

// ==================== keyof 约束 ====================

function getProperty<T, K extends keyof T>(obj: T, key: K): T[K] {
  return obj[key];
}

const user = { name: "John", age: 30 };
getProperty(user, "name");  // string
getProperty(user, "age");   // number
// getProperty(user, "foo"); // 错误

// ==================== 条件类型 ====================

type Flatten<T> = T extends Array<infer U> ? U : T;

type A = Flatten<string[]>;  // string
type B = Flatten<number>;    // number

// ==================== 泛型类 ====================

class Container<T> {
  private value: T;
  
  constructor(value: T) {
    this.value = value;
  }
  
  getValue(): T {
    return this.value;
  }
  
  map<U>(fn: (value: T) => U): Container<U> {
    return new Container(fn(this.value));
  }
}

// ==================== 泛型接口 ====================

interface Repository<T> {
  find(id: number): T | undefined;
  findAll(): T[];
  save(entity: T): void;
  delete(id: number): void;
}

class UserRepository implements Repository<User> {
  find(id: number): User | undefined { /* ... */ }
  findAll(): User[] { /* ... */ }
  save(entity: User): void { /* ... */ }
  delete(id: number): void { /* ... */ }
}

// ==================== 默认类型参数 ====================

interface Response<T = unknown> {
  data: T;
  status: number;
}

function fetch<T = unknown>(url: string): Promise<Response<T>> {
  // ...
}
```

### Python 泛型

```python
from typing import TypeVar, Generic, Callable, Sequence
from collections.abc import Iterable

# ==================== 基础泛型函数 ====================

T = TypeVar('T')

def identity(value: T) -> T:
    return value

# Python 3.12+ 新语法
def identity[T](value: T) -> T:
    return value

# ==================== 多类型参数 ====================

T = TypeVar('T')
U = TypeVar('U')

def pair(first: T, second: U) -> tuple[T, U]:
    return (first, second)

def map_list[T, U](items: list[T], fn: Callable[[T], U]) -> list[U]:
    return [fn(item) for item in items]

# ==================== 类型约束 (bound) ====================

from typing import Protocol

class HasLength(Protocol):
    def __len__(self) -> int: ...

T = TypeVar('T', bound=HasLength)

def log_length(value: T) -> T:
    print(len(value))
    return value

# Python 3.12+
def log_length[T: HasLength](value: T) -> T:
    print(len(value))
    return value

# 限制为特定类型
T = TypeVar('T', int, float)  # 只能是 int 或 float

def add(a: T, b: T) -> T:
    return a + b

# ==================== Protocol (结构化类型) ====================

from typing import Protocol

class Comparable(Protocol):
    def __lt__(self, other: 'Comparable') -> bool: ...

def min_value[T: Comparable](a: T, b: T) -> T:
    return a if a < b else b

# ==================== 泛型类 ====================

class Container(Generic[T]):
    def __init__(self, value: T):
        self.value = value
    
    def get_value(self) -> T:
        return self.value
    
    def map[U](self, fn: Callable[[T], U]) -> 'Container[U]':
        return Container(fn(self.value))

# Python 3.12+
class Container[T]:
    def __init__(self, value: T):
        self.value = value
    
    def get_value(self) -> T:
        return self.value

# ==================== 泛型协议 ====================

class Repository(Protocol[T]):
    def find(self, id: int) -> T | None: ...
    def find_all(self) -> list[T]: ...
    def save(self, entity: T) -> None: ...

# ==================== 协变与逆变 ====================

from typing import TypeVar

T_co = TypeVar('T_co', covariant=True)      # 协变
T_contra = TypeVar('T_contra', contravariant=True)  # 逆变

class Reader(Generic[T_co]):
    def read(self) -> T_co: ...

class Writer(Generic[T_contra]):
    def write(self, value: T_contra) -> None: ...

# ==================== ParamSpec (函数签名) ====================

from typing import ParamSpec, Concatenate

P = ParamSpec('P')

def with_logging[T, **P](fn: Callable[P, T]) -> Callable[P, T]:
    def wrapper(*args: P.args, **kwargs: P.kwargs) -> T:
        print(f"Calling {fn.__name__}")
        return fn(*args, **kwargs)
    return wrapper
```

### Go 泛型

```go
// ==================== 基础泛型函数 ====================

func Identity[T any](value T) T {
    return value
}

// 调用
Identity[string]("hello")  // 显式指定
Identity("hello")          // 类型推断

// ==================== 多类型参数 ====================

func Pair[T, U any](first T, second U) (T, U) {
    return first, second
}

func Map[T, U any](slice []T, fn func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = fn(v)
    }
    return result
}

// ==================== 类型约束 ====================

// 使用接口约束
type Stringer interface {
    String() string
}

func ToString[T Stringer](value T) string {
    return value.String()
}

// 内置约束
import "golang.org/x/exp/constraints"

func Min[T constraints.Ordered](a, b T) T {
    if a < b {
        return a
    }
    return b
}

// ==================== 自定义约束 ====================

type Number interface {
    int | int8 | int16 | int32 | int64 |
    uint | uint8 | uint16 | uint32 | uint64 |
    float32 | float64
}

func Sum[T Number](numbers []T) T {
    var sum T
    for _, n := range numbers {
        sum += n
    }
    return sum
}

// 近似约束 (~)
type Integer interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64
}

type MyInt int
// MyInt 满足 Integer 约束，因为底层类型是 int

// ==================== 泛型结构体 ====================

type Container[T any] struct {
    value T
}

func NewContainer[T any](value T) *Container[T] {
    return &Container[T]{value: value}
}

func (c *Container[T]) Get() T {
    return c.value
}

func (c *Container[T]) Set(value T) {
    c.value = value
}

// ==================== 泛型接口 ====================

type Repository[T any] interface {
    Find(id int) (T, error)
    FindAll() ([]T, error)
    Save(entity T) error
    Delete(id int) error
}

type UserRepository struct{}

func (r *UserRepository) Find(id int) (User, error) { /* ... */ }
func (r *UserRepository) FindAll() ([]User, error) { /* ... */ }
func (r *UserRepository) Save(entity User) error { /* ... */ }
func (r *UserRepository) Delete(id int) error { /* ... */ }

// ==================== 类型推断 ====================

// 部分推断
func MakePair[T, U any](first T) func(U) (T, U) {
    return func(second U) (T, U) {
        return first, second
    }
}

// 使用
makePairWithString := MakePair[string, int]("hello")
pair := makePairWithString(42)  // ("hello", 42)

// ==================== 常用泛型函数 ====================

// Filter
func Filter[T any](slice []T, predicate func(T) bool) []T {
    var result []T
    for _, v := range slice {
        if predicate(v) {
            result = append(result, v)
        }
    }
    return result
}

// Reduce
func Reduce[T, U any](slice []T, initial U, fn func(U, T) U) U {
    result := initial
    for _, v := range slice {
        result = fn(result, v)
    }
    return result
}

// Contains
func Contains[T comparable](slice []T, target T) bool {
    for _, v := range slice {
        if v == target {
            return true
        }
    }
    return false
}
```

### Rust 泛型

```rust
// ==================== 基础泛型函数 ====================

fn identity<T>(value: T) -> T {
    value
}

// 调用
identity::<String>(String::from("hello"));  // 显式指定
identity("hello");                          // 类型推断

// ==================== 多类型参数 ====================

fn pair<T, U>(first: T, second: U) -> (T, U) {
    (first, second)
}

fn map<T, U, F>(vec: Vec<T>, f: F) -> Vec<U>
where
    F: Fn(T) -> U,
{
    vec.into_iter().map(f).collect()
}

// ==================== trait 约束 ====================

use std::fmt::Display;

fn print_value<T: Display>(value: T) {
    println!("{}", value);
}

// 多个约束
fn compare_and_print<T: PartialOrd + Display>(a: T, b: T) {
    if a > b {
        println!("{} > {}", a, b);
    }
}

// where 子句 (更清晰)
fn process<T, U>(t: T, u: U) -> String
where
    T: Display + Clone,
    U: Display + Default,
{
    format!("{} - {}", t, u)
}

// ==================== 常用 trait 约束 ====================

// Clone: 可克隆
fn duplicate<T: Clone>(value: T) -> (T, T) {
    (value.clone(), value)
}

// Default: 有默认值
fn with_default<T: Default>() -> T {
    T::default()
}

// PartialEq: 可比较相等
fn contains<T: PartialEq>(slice: &[T], target: &T) -> bool {
    slice.iter().any(|x| x == target)
}

// Ord: 可排序
fn min<T: Ord>(a: T, b: T) -> T {
    if a < b { a } else { b }
}

// ==================== 泛型结构体 ====================

struct Container<T> {
    value: T,
}

impl<T> Container<T> {
    fn new(value: T) -> Self {
        Container { value }
    }
    
    fn get(&self) -> &T {
        &self.value
    }
    
    fn into_inner(self) -> T {
        self.value
    }
}

// 特定类型的实现
impl Container<String> {
    fn len(&self) -> usize {
        self.value.len()
    }
}

// 带约束的实现
impl<T: Display> Container<T> {
    fn print(&self) {
        println!("{}", self.value);
    }
}

// ==================== 泛型 trait ====================

trait Repository<T> {
    fn find(&self, id: u32) -> Option<T>;
    fn find_all(&self) -> Vec<T>;
    fn save(&mut self, entity: T);
}

struct UserRepository {
    users: Vec<User>,
}

impl Repository<User> for UserRepository {
    fn find(&self, id: u32) -> Option<User> {
        self.users.iter().find(|u| u.id == id).cloned()
    }
    fn find_all(&self) -> Vec<User> {
        self.users.clone()
    }
    fn save(&mut self, entity: User) {
        self.users.push(entity);
    }
}

// ==================== 关联类型 vs 泛型 ====================

// 泛型 trait: 一个类型可以实现多次
trait From<T> {
    fn from(value: T) -> Self;
}

// 关联类型: 一个类型只能实现一次
trait Iterator {
    type Item;
    fn next(&mut self) -> Option<Self::Item>;
}

// ==================== impl Trait ====================

// 返回实现某 trait 的类型 (不暴露具体类型)
fn make_iter() -> impl Iterator<Item = i32> {
    (0..10).filter(|x| x % 2 == 0)
}

// 参数中使用
fn process(iter: impl Iterator<Item = i32>) {
    for item in iter {
        println!("{}", item);
    }
}

// ==================== 生命周期与泛型 ====================

fn longest<'a, T>(x: &'a T, y: &'a T) -> &'a T
where
    T: PartialOrd,
{
    if x > y { x } else { y }
}

struct Ref<'a, T> {
    value: &'a T,
}

impl<'a, T> Ref<'a, T> {
    fn new(value: &'a T) -> Self {
        Ref { value }
    }
}

// ==================== const 泛型 ====================

fn create_array<T: Default + Copy, const N: usize>() -> [T; N] {
    [T::default(); N]
}

let arr: [i32; 5] = create_array();

struct Matrix<T, const ROWS: usize, const COLS: usize> {
    data: [[T; COLS]; ROWS],
}
```

### 泛型对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 语法            │ <T>            │ [T] (3.12+)    │ [T]            │ <T>            │
│ 约束语法        │ extends        │ bound/Protocol │ interface      │ trait bound    │
│ 类型推断        │ ✅             │ ✅             │ ✅             │ ✅             │
│ 默认类型        │ ✅             │ ✅             │ ❌             │ ✅             │
│ 协变/逆变       │ 自动           │ 显式声明       │ ❌             │ 自动           │
│ const 泛型      │ ❌             │ ❌             │ ❌             │ ✅             │
│ 特化            │ 条件类型       │ @overload      │ ❌             │ 部分支持       │
│ 运行时类型      │ 擦除           │ 擦除           │ 保留           │ 单态化         │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 常用泛型模式

```
┌────────────────────────────────────────────────────────────────────────┐
│                          常用泛型模式                                   │
├────────────────────────────────────────────────────────────────────────┤
│                                                                        │
│  1. 容器类型                                                           │
│     Container<T>, Option<T>, Result<T, E>, Vec<T>                     │
│                                                                        │
│  2. 函数式操作                                                         │
│     map<T, U>(T) -> U                                                 │
│     filter<T>(predicate: T -> bool)                                   │
│     reduce<T, U>(initial: U, fn: (U, T) -> U)                        │
│                                                                        │
│  3. 仓储模式                                                           │
│     Repository<T> { find, findAll, save, delete }                     │
│                                                                        │
│  4. 工厂模式                                                           │
│     Factory<T> { create() -> T }                                      │
│                                                                        │
│  5. 比较与排序                                                         │
│     min<T: Ord>(a, b) -> T                                            │
│     sort<T: Ord>(slice: []T)                                          │
│                                                                        │
└────────────────────────────────────────────────────────────────────────┘
```

---

## 🔙 函数返回空值处理

### 返回空值模式概览

| 语言 | 空值返回 | 推荐模式 | 特点 |
|------|----------|----------|------|
| TypeScript | `null \| undefined` | `T \| null` 或 `T \| undefined` | 需手动检查 |
| Python | `None` | `Optional[T]` | 需手动检查 |
| Go | `nil` + `error` | 多返回值 `(T, error)` | 惯用模式 |
| Rust | 无 null | `Option<T>` / `Result<T, E>` | 编译器强制 |

### TypeScript 函数返回空值

```typescript
// ==================== 返回 null/undefined ====================

// 可能返回 null
function findUser(id: number): User | null {
  const user = db.find(u => u.id === id);
  return user ?? null;
}

// 可能返回 undefined
function getValue(key: string): string | undefined {
  return cache.get(key);
}

// ==================== 调用方处理 ====================

// 方式 1: if 检查
const user = findUser(1);
if (user !== null) {
  console.log(user.name);  // 类型收窄为 User
}

// 方式 2: 可选链
const name = findUser(1)?.name;

// 方式 3: 空值合并
const name = findUser(1)?.name ?? "Anonymous";

// 方式 4: 断言 (确定不为空时)
const user = findUser(1)!;  // 危险！

// ==================== 早期返回模式 ====================

function processUser(id: number): string {
  const user = findUser(id);
  if (!user) {
    return "User not found";
  }
  // 此处 user 类型为 User
  return `Hello, ${user.name}`;
}

// ==================== 抛出异常 vs 返回 null ====================

// 返回 null (调用方决定如何处理)
function findUserOrNull(id: number): User | null {
  return db.find(u => u.id === id) ?? null;
}

// 抛出异常 (表示程序错误)
function findUserOrThrow(id: number): User {
  const user = db.find(u => u.id === id);
  if (!user) {
    throw new Error(`User ${id} not found`);
  }
  return user;
}

// ==================== Result 模式 (推荐) ====================

type Result<T, E = Error> = 
  | { ok: true; value: T }
  | { ok: false; error: E };

function findUser(id: number): Result<User, string> {
  const user = db.find(u => u.id === id);
  if (!user) {
    return { ok: false, error: "User not found" };
  }
  return { ok: true, value: user };
}

// 使用
const result = findUser(1);
if (result.ok) {
  console.log(result.value.name);
} else {
  console.error(result.error);
}

// ==================== neverthrow 库 ====================
import { ok, err, Result } from 'neverthrow';

function divide(a: number, b: number): Result<number, string> {
  if (b === 0) return err("Division by zero");
  return ok(a / b);
}

divide(10, 2)
  .map(n => n * 2)
  .mapErr(e => `Error: ${e}`)
  .match(
    value => console.log(value),
    error => console.error(error)
  );
```

### Python 函数返回空值

```python
from typing import Optional, Union
from dataclasses import dataclass

# ==================== 返回 None ====================

def find_user(user_id: int) -> Optional[User]:
    """返回用户或 None"""
    user = db.get(user_id)
    return user  # 可能是 None

def get_value(key: str) -> str | None:  # Python 3.10+
    return cache.get(key)

# ==================== 调用方处理 ====================

# 方式 1: if 检查
user = find_user(1)
if user is not None:
    print(user.name)

# 方式 2: or 默认值 (注意 falsy 值问题)
name = find_user(1) or default_user

# 方式 3: 条件表达式
name = user.name if (user := find_user(1)) else "Anonymous"

# 方式 4: getattr 带默认值
name = getattr(find_user(1), 'name', 'Anonymous')

# ==================== 早期返回模式 ====================

def process_user(user_id: int) -> str:
    user = find_user(user_id)
    if user is None:
        return "User not found"
    return f"Hello, {user.name}"

# ==================== 抛出异常 vs 返回 None ====================

# 返回 None
def find_user_or_none(user_id: int) -> Optional[User]:
    return db.get(user_id)

# 抛出异常
def find_user_or_raise(user_id: int) -> User:
    user = db.get(user_id)
    if user is None:
        raise ValueError(f"User {user_id} not found")
    return user

# ==================== Result 模式 ====================

@dataclass
class Ok[T]:
    value: T

@dataclass
class Err[E]:
    error: E

Result = Ok[T] | Err[E]

def find_user(user_id: int) -> Result[User, str]:
    user = db.get(user_id)
    if user is None:
        return Err("User not found")
    return Ok(user)

# 使用
match find_user(1):
    case Ok(user):
        print(user.name)
    case Err(error):
        print(f"Error: {error}")

# ==================== returns 库 ====================
# pip install returns
from returns.result import Result, Success, Failure
from returns.maybe import Maybe, Some, Nothing

def find_user(user_id: int) -> Result[User, str]:
    user = db.get(user_id)
    if user is None:
        return Failure("User not found")
    return Success(user)

# Maybe 类型
def get_name(user_id: int) -> Maybe[str]:
    user = db.get(user_id)
    if user is None:
        return Nothing
    return Some(user.name)
```

### Go 函数返回空值

```go
// ==================== 返回 nil + error (惯用模式) ====================

func FindUser(id int) (*User, error) {
    user := db.Get(id)
    if user == nil {
        return nil, fmt.Errorf("user %d not found", id)
    }
    return user, nil
}

// ==================== 调用方处理 ====================

// 方式 1: 检查 error
user, err := FindUser(1)
if err != nil {
    log.Printf("Error: %v", err)
    return
}
fmt.Println(user.Name)

// 方式 2: 忽略错误 (不推荐)
user, _ := FindUser(1)
if user != nil {
    fmt.Println(user.Name)
}

// ==================== 多种返回模式 ====================

// 模式 1: 返回 (value, error)
func GetConfig(key string) (string, error) {
    value, exists := config[key]
    if !exists {
        return "", fmt.Errorf("config %s not found", key)
    }
    return value, nil
}

// 模式 2: 返回 (value, bool)
func GetCache(key string) (string, bool) {
    value, ok := cache[key]
    return value, ok
}

// 使用
if value, ok := GetCache("key"); ok {
    fmt.Println(value)
}

// 模式 3: 只返回 error
func SaveUser(user *User) error {
    if user == nil {
        return errors.New("user cannot be nil")
    }
    return db.Save(user)
}

// ==================== 哨兵错误 ====================

var (
    ErrNotFound = errors.New("not found")
    ErrInvalid  = errors.New("invalid")
)

func FindUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrInvalid
    }
    user := db.Get(id)
    if user == nil {
        return nil, ErrNotFound
    }
    return user, nil
}

// 检查特定错误
user, err := FindUser(1)
if errors.Is(err, ErrNotFound) {
    // 用户不存在
} else if err != nil {
    // 其他错误
}

// ==================== 错误包装 ====================

func GetUserProfile(id int) (*Profile, error) {
    user, err := FindUser(id)
    if err != nil {
        return nil, fmt.Errorf("get user profile: %w", err)
    }
    return user.Profile, nil
}

// ==================== 泛型 Optional (Go 1.18+) ====================

type Optional[T any] struct {
    value   T
    present bool
}

func Some[T any](value T) Optional[T] {
    return Optional[T]{value: value, present: true}
}

func None[T any]() Optional[T] {
    return Optional[T]{}
}

func (o Optional[T]) IsPresent() bool {
    return o.present
}

func (o Optional[T]) Get() (T, bool) {
    return o.value, o.present
}

func (o Optional[T]) OrElse(defaultValue T) T {
    if o.present {
        return o.value
    }
    return defaultValue
}

// 使用
func FindUser(id int) Optional[User] {
    user := db.Get(id)
    if user == nil {
        return None[User]()
    }
    return Some(*user)
}

result := FindUser(1)
if result.IsPresent() {
    user, _ := result.Get()
    fmt.Println(user.Name)
}
```

### Rust 函数返回空值

```rust
// ==================== Option<T> ====================

fn find_user(id: u32) -> Option<User> {
    db.get(&id).cloned()
}

// ==================== 调用方处理 ====================

// 方式 1: match
match find_user(1) {
    Some(user) => println!("{}", user.name),
    None => println!("User not found"),
}

// 方式 2: if let
if let Some(user) = find_user(1) {
    println!("{}", user.name);
}

// 方式 3: let else
let Some(user) = find_user(1) else {
    println!("User not found");
    return;
};
println!("{}", user.name);

// 方式 4: 方法链
let name = find_user(1)
    .map(|u| u.name.clone())
    .unwrap_or_else(|| "Anonymous".to_string());

// ==================== Option 方法 ====================

let user = find_user(1);

// 获取值
user.unwrap();              // 有值返回，None 则 panic
user.expect("No user");     // 带消息的 unwrap
user.unwrap_or(default);    // 默认值
user.unwrap_or_else(|| compute_default());  // 惰性默认值
user.unwrap_or_default();   // 类型默认值

// 转换
user.map(|u| u.name);       // Option<String>
user.and_then(|u| u.email); // 链式 Option
user.filter(|u| u.age > 18);
user.ok_or("Not found")?;   // 转为 Result

// 检查
user.is_some();
user.is_none();

// ==================== Result<T, E> ====================

fn find_user(id: u32) -> Result<User, UserError> {
    if id == 0 {
        return Err(UserError::InvalidId);
    }
    db.get(&id)
        .cloned()
        .ok_or(UserError::NotFound(id))
}

// ==================== ? 操作符 ====================

fn get_user_email(id: u32) -> Result<String, UserError> {
    let user = find_user(id)?;  // 错误自动返回
    let email = user.email.ok_or(UserError::NoEmail)?;
    Ok(email)
}

// ==================== 自定义错误 ====================

#[derive(Debug, thiserror::Error)]
enum UserError {
    #[error("Invalid user ID")]
    InvalidId,
    #[error("User {0} not found")]
    NotFound(u32),
    #[error("User has no email")]
    NoEmail,
}

// ==================== 组合多个 Option/Result ====================

// Option 组合
fn get_full_name(id: u32) -> Option<String> {
    let user = find_user(id)?;
    let first = user.first_name.as_ref()?;
    let last = user.last_name.as_ref()?;
    Some(format!("{} {}", first, last))
}

// Result 组合
fn process_user(id: u32) -> Result<String, UserError> {
    let user = find_user(id)?;
    let profile = get_profile(user.id)?;
    let settings = get_settings(user.id)?;
    Ok(format!("{}: {:?}", profile.name, settings))
}

// ==================== 转换 Option <-> Result ====================

// Option -> Result
let result: Result<User, &str> = find_user(1).ok_or("Not found");

// Result -> Option
let option: Option<User> = find_user_result(1).ok();

// ==================== anyhow (应用代码) ====================

use anyhow::{Context, Result, bail};

fn find_user(id: u32) -> Result<User> {
    db.get(&id)
        .cloned()
        .context(format!("User {} not found", id))
}

fn process(id: u32) -> Result<()> {
    let user = find_user(id)?;
    if user.banned {
        bail!("User is banned");
    }
    Ok(())
}
```

### 函数返回空值对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 空值类型        │ null/undefined │ None           │ nil            │ 无             │
│ 推荐返回类型    │ T | null       │ Optional[T]    │ (*T, error)    │ Option<T>      │
│ 错误返回        │ throw / Result │ raise / Result │ error          │ Result<T, E>   │
│ 强制检查        │ ❌ (strictNull)│ ❌             │ ❌             │ ✅             │
│ 链式处理        │ ?.             │ 有限           │ ❌             │ map/and_then   │
│ 提前返回        │ if + return    │ if + return    │ if err != nil  │ ?              │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 最佳实践

```
┌─────────────────────────────────────────────────────────────────────────┐
│                        函数返回空值最佳实践                              │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  1. 明确语义                                                            │
│     • 返回 null/None: 表示"没有结果"是正常情况                          │
│     • 抛出异常/返回 error: 表示"发生了错误"                             │
│                                                                         │
│  2. 类型签名要诚实                                                      │
│     • TypeScript: 使用 strictNullChecks                                │
│     • Python: 使用 Optional[T] 类型提示                                │
│     • Go: 始终返回 error 作为第二个值                                  │
│     • Rust: 使用 Option/Result 而非 panic                              │
│                                                                         │
│  3. 调用方处理                                                          │
│     • 总是检查空值再使用                                               │
│     • 使用早期返回减少嵌套                                             │
│     • 提供合理的默认值                                                 │
│                                                                         │
│  4. 文档说明                                                            │
│     • 说明何时返回空值                                                 │
│     • 说明调用方应如何处理                                             │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

---

## 📁 文件系统与 IO

### IO 操作概览

| 特性 | TypeScript (Node) | Python | Go | Rust |
|------|-------------------|--------|-----|------|
| 同步读写 | `fs.readFileSync` | `open().read()` | `os.ReadFile` | `std::fs::read` |
| 异步读写 | `fs/promises` | `aiofiles` | goroutine | `tokio::fs` |
| 流式读写 | `createReadStream` | 迭代器 | `bufio` | `BufReader` |
| 路径处理 | `path` | `pathlib` | `filepath` | `std::path` |
| 目录操作 | `fs.mkdir/readdir` | `os/pathlib` | `os` | `std::fs` |

### TypeScript 文件操作

```typescript
import * as fs from 'fs';
import * as fsp from 'fs/promises';
import * as path from 'path';

// ==================== 同步操作 ====================
const content = fs.readFileSync('file.txt', 'utf-8');
fs.writeFileSync('output.txt', 'Hello World');
fs.appendFileSync('log.txt', 'New line\n');

// ==================== 异步操作 (Promise) ====================
const content = await fsp.readFile('file.txt', 'utf-8');
await fsp.writeFile('output.txt', 'Hello World');

// ==================== 流式操作 ====================
import { createReadStream, createWriteStream } from 'fs';
import { pipeline } from 'stream/promises';

const readStream = createReadStream('large-file.txt');
for await (const chunk of readStream) {
    process(chunk);
}

await pipeline(
    createReadStream('source.txt'),
    createWriteStream('dest.txt')
);

// ==================== 目录操作 ====================
await fsp.mkdir('new-dir', { recursive: true });
const entries = await fsp.readdir('.', { withFileTypes: true });
await fsp.rm('dir', { recursive: true });

// ==================== 路径处理 ====================
path.join('dir', 'subdir', 'file.txt');
path.resolve('relative');
path.dirname('/a/b/c.txt');   // /a/b
path.basename('/a/b/c.txt');  // c.txt
path.extname('file.txt');     // .txt
```

### Python 文件操作

```python
from pathlib import Path
import aiofiles

# ==================== 同步操作 ====================
content = Path('file.txt').read_text(encoding='utf-8')
Path('output.txt').write_text('Hello World')

with open('file.txt', 'r') as f:
    for line in f:
        print(line.strip())

# ==================== 异步操作 ====================
async with aiofiles.open('file.txt', 'r') as f:
    content = await f.read()

# ==================== pathlib (推荐) ====================
p = Path('dir/subdir/file.txt')
p.parent          # dir/subdir
p.name            # file.txt
p.suffix          # .txt
new_path = Path('dir') / 'subdir' / 'file.txt'

# ==================== 目录操作 ====================
Path('new-dir').mkdir(parents=True, exist_ok=True)
list(Path('.').rglob('*.py'))  # 递归查找
import shutil
shutil.rmtree('dir')
```

### Go 文件操作

```go
import (
    "bufio"
    "io"
    "os"
    "path/filepath"
)

// ==================== 同步操作 ====================
content, _ := os.ReadFile("file.txt")
os.WriteFile("output.txt", []byte("Hello"), 0644)

// ==================== 缓冲读写 ====================
f, _ := os.Open("file.txt")
defer f.Close()
scanner := bufio.NewScanner(f)
for scanner.Scan() {
    fmt.Println(scanner.Text())
}

// ==================== 目录操作 ====================
os.MkdirAll("path/to/dir", 0755)
entries, _ := os.ReadDir(".")
filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
})

// ==================== 路径处理 ====================
filepath.Join("dir", "subdir", "file.txt")
filepath.Dir("/a/b/c.txt")   // /a/b
filepath.Base("/a/b/c.txt")  // c.txt
```

### Rust 文件操作

```rust
use std::fs;
use std::io::{BufRead, BufReader};
use std::path::Path;

// ==================== 同步操作 ====================
let content = fs::read_to_string("file.txt")?;
fs::write("output.txt", "Hello World")?;

// ==================== 缓冲读取 ====================
let file = fs::File::open("file.txt")?;
for line in BufReader::new(file).lines() {
    println!("{}", line?);
}

// ==================== 异步操作 (tokio) ====================
let content = tokio::fs::read_to_string("file.txt").await?;
tokio::fs::write("output.txt", "Hello").await?;

// ==================== 目录操作 ====================
fs::create_dir_all("path/to/dir")?;
for entry in fs::read_dir(".")? {
    println!("{:?}", entry?.path());
}

// ==================== 路径处理 ====================
let path = Path::new("dir").join("file.txt");
path.parent();      // Some("dir")
path.file_name();   // Some("file.txt")
path.extension();   // Some("txt")
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

### Web 框架详细对比 (Koa vs FastAPI vs Gin vs Axum)

#### 路由参数 (Path Parameters)

```typescript
// ==================== Koa (koa-router) ====================
import Router from '@koa/router';
const router = new Router();

router.get('/users/:id', async (ctx) => {
    const id = ctx.params.id;
    ctx.body = { id };
});

router.get('/posts/:year/:month', async (ctx) => {
    const { year, month } = ctx.params;
    ctx.body = { year, month };
});
```

```python
# ==================== FastAPI ====================
from fastapi import FastAPI, Path

app = FastAPI()

@app.get("/users/{user_id}")
async def get_user(user_id: int):  # 自动类型转换
    return {"user_id": user_id}

@app.get("/posts/{year}/{month}")
async def get_posts(
    year: int = Path(..., ge=2000, le=2100),  # 验证
    month: int = Path(..., ge=1, le=12)
):
    return {"year": year, "month": month}
```

```go
// ==================== Gin ====================
r := gin.Default()

r.GET("/users/:id", func(c *gin.Context) {
    id := c.Param("id")
    c.JSON(200, gin.H{"id": id})
})

r.GET("/posts/:year/:month", func(c *gin.Context) {
    year := c.Param("year")
    month := c.Param("month")
    c.JSON(200, gin.H{"year": year, "month": month})
})
```

```rust
// ==================== Axum ====================
use axum::{extract::Path, routing::get, Router};

async fn get_user(Path(id): Path<u32>) -> String {
    format!("User {}", id)
}

async fn get_posts(Path((year, month)): Path<(u32, u32)>) -> String {
    format!("{}/{}", year, month)
}

let app = Router::new()
    .route("/users/:id", get(get_user))
    .route("/posts/:year/:month", get(get_posts));
```

#### 查询参数 (Query Parameters)

```typescript
// ==================== Koa ====================
router.get('/search', async (ctx) => {
    const { q, page = '1', limit = '10' } = ctx.query;
    ctx.body = { q, page: parseInt(page), limit: parseInt(limit) };
});
```

```python
# ==================== FastAPI ====================
from fastapi import Query
from typing import Optional

@app.get("/search")
async def search(
    q: str,                                    # 必填
    page: int = 1,                            # 默认值
    limit: int = Query(10, ge=1, le=100),     # 带验证
    tags: Optional[list[str]] = Query(None)   # 可选列表
):
    return {"q": q, "page": page, "limit": limit, "tags": tags}
```

```go
// ==================== Gin ====================
type SearchQuery struct {
    Q     string   `form:"q" binding:"required"`
    Page  int      `form:"page,default=1"`
    Limit int      `form:"limit,default=10"`
    Tags  []string `form:"tags"`
}

r.GET("/search", func(c *gin.Context) {
    var query SearchQuery
    if err := c.ShouldBindQuery(&query); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, query)
})
```

```rust
// ==================== Axum ====================
use axum::extract::Query;
use serde::Deserialize;

#[derive(Deserialize)]
struct SearchQuery {
    q: String,
    #[serde(default = "default_page")]
    page: u32,
    #[serde(default = "default_limit")]
    limit: u32,
}

fn default_page() -> u32 { 1 }
fn default_limit() -> u32 { 10 }

async fn search(Query(query): Query<SearchQuery>) -> String {
    format!("Search: {} page {}", query.q, query.page)
}
```

#### 请求头 (Headers)

```typescript
// ==================== Koa ====================
router.get('/api', async (ctx) => {
    const auth = ctx.get('Authorization');
    const contentType = ctx.get('Content-Type');
    const userAgent = ctx.request.headers['user-agent'];
    
    ctx.set('X-Custom-Header', 'value');
    ctx.body = { auth };
});
```

```python
# ==================== FastAPI ====================
from fastapi import Header

@app.get("/api")
async def api(
    authorization: str = Header(...),
    user_agent: str = Header(None, alias="User-Agent"),
    x_token: list[str] = Header(None)
):
    return {"auth": authorization}

# 响应头
from fastapi import Response

@app.get("/download")
async def download(response: Response):
    response.headers["X-Custom"] = "value"
    return {"file": "data"}
```

```go
// ==================== Gin ====================
r.GET("/api", func(c *gin.Context) {
    auth := c.GetHeader("Authorization")
    userAgent := c.Request.Header.Get("User-Agent")
    
    c.Header("X-Custom-Header", "value")
    c.JSON(200, gin.H{"auth": auth})
})

// 绑定到结构体
type Headers struct {
    Authorization string `header:"Authorization" binding:"required"`
    UserAgent     string `header:"User-Agent"`
}

r.GET("/api", func(c *gin.Context) {
    var h Headers
    c.ShouldBindHeader(&h)
})
```

```rust
// ==================== Axum ====================
use axum::http::{HeaderMap, header};
use axum_extra::TypedHeader;
use headers::Authorization;

async fn api(
    TypedHeader(auth): TypedHeader<Authorization<headers::authorization::Bearer>>,
    headers: HeaderMap,
) -> impl IntoResponse {
    let user_agent = headers.get(header::USER_AGENT);
    
    (
        [(header::CONTENT_TYPE, "application/json")],
        format!("Token: {}", auth.token())
    )
}
```

#### 请求体 (Body)

```typescript
// ==================== Koa (koa-bodyparser) ====================
import bodyParser from 'koa-bodyparser';
app.use(bodyParser());

router.post('/users', async (ctx) => {
    const body = ctx.request.body;  // JSON 自动解析
    ctx.body = body;
});

// 文件上传 (koa-multer)
import multer from '@koa/multer';
const upload = multer({ dest: 'uploads/' });

router.post('/upload', upload.single('file'), async (ctx) => {
    const file = ctx.request.file;
    ctx.body = { filename: file.filename };
});
```

```python
# ==================== FastAPI ====================
from pydantic import BaseModel, Field

class CreateUser(BaseModel):
    name: str = Field(..., min_length=1, max_length=100)
    email: str
    age: int = Field(None, ge=0, le=150)

@app.post("/users")
async def create_user(user: CreateUser):  # 自动验证
    return user

# 原始 body
from fastapi import Body

@app.post("/raw")
async def raw_body(data: dict = Body(...)):
    return data

# 文件上传
from fastapi import File, UploadFile

@app.post("/upload")
async def upload(file: UploadFile = File(...)):
    content = await file.read()
    return {"filename": file.filename, "size": len(content)}
```

```go
// ==================== Gin ====================
type CreateUser struct {
    Name  string `json:"name" binding:"required,min=1,max=100"`
    Email string `json:"email" binding:"required,email"`
    Age   int    `json:"age" binding:"gte=0,lte=150"`
}

r.POST("/users", func(c *gin.Context) {
    var user CreateUser
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(400, gin.H{"error": err.Error()})
        return
    }
    c.JSON(201, user)
})

// 文件上传
r.POST("/upload", func(c *gin.Context) {
    file, _ := c.FormFile("file")
    c.SaveUploadedFile(file, "uploads/"+file.Filename)
    c.JSON(200, gin.H{"filename": file.Filename})
})

// 多文件
r.POST("/uploads", func(c *gin.Context) {
    form, _ := c.MultipartForm()
    files := form.File["files"]
    for _, file := range files {
        c.SaveUploadedFile(file, "uploads/"+file.Filename)
    }
})
```

```rust
// ==================== Axum ====================
use axum::{Json, extract::Multipart};
use serde::{Deserialize, Serialize};

#[derive(Deserialize)]
struct CreateUser {
    name: String,
    email: String,
    age: Option<u32>,
}

async fn create_user(Json(user): Json<CreateUser>) -> Json<CreateUser> {
    Json(user)
}

// 文件上传
async fn upload(mut multipart: Multipart) -> String {
    while let Some(field) = multipart.next_field().await.unwrap() {
        let name = field.name().unwrap().to_string();
        let data = field.bytes().await.unwrap();
        println!("Field: {} Size: {}", name, data.len());
    }
    "OK".to_string()
}
```

#### 应用状态 (State)

```typescript
// ==================== Koa ====================
// 方式1: app.context
app.context.db = database;

router.get('/users', async (ctx) => {
    const users = await ctx.db.getUsers();
    ctx.body = users;
});

// 方式2: 中间件注入
app.use(async (ctx, next) => {
    ctx.state.db = database;
    ctx.state.config = config;
    await next();
});

router.get('/users', async (ctx) => {
    const users = await ctx.state.db.getUsers();
});
```

```python
# ==================== FastAPI ====================
from fastapi import Depends

# 依赖注入
def get_db():
    db = Database()
    try:
        yield db
    finally:
        db.close()

@app.get("/users")
async def get_users(db: Database = Depends(get_db)):
    return await db.get_users()

# app.state
app.state.config = Config()

@app.get("/config")
async def get_config(request: Request):
    return request.app.state.config
```

```go
// ==================== Gin ====================
// 方式1: 中间件
func DatabaseMiddleware(db *Database) gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Set("db", db)
        c.Next()
    }
}

r.Use(DatabaseMiddleware(db))

r.GET("/users", func(c *gin.Context) {
    db := c.MustGet("db").(*Database)
    users := db.GetUsers()
    c.JSON(200, users)
})

// 方式2: 闭包
func SetupRoutes(db *Database) *gin.Engine {
    r := gin.Default()
    
    r.GET("/users", func(c *gin.Context) {
        users := db.GetUsers()  // 直接使用
        c.JSON(200, users)
    })
    
    return r
}
```

```rust
// ==================== Axum ====================
use axum::extract::State;
use std::sync::Arc;

struct AppState {
    db: Database,
    config: Config,
}

async fn get_users(State(state): State<Arc<AppState>>) -> Json<Vec<User>> {
    let users = state.db.get_users().await;
    Json(users)
}

let state = Arc::new(AppState { db, config });
let app = Router::new()
    .route("/users", get(get_users))
    .with_state(state);
```

#### 路由分组 (Router Groups)

```typescript
// ==================== Koa ====================
const apiRouter = new Router({ prefix: '/api/v1' });
const adminRouter = new Router({ prefix: '/admin' });

apiRouter.get('/users', getUsers);
apiRouter.post('/users', createUser);

adminRouter.get('/stats', getStats);
adminRouter.use(authMiddleware);  // 组级中间件

app.use(apiRouter.routes());
app.use(adminRouter.routes());
```

```python
# ==================== FastAPI ====================
from fastapi import APIRouter

# 创建路由组
users_router = APIRouter(prefix="/users", tags=["users"])
admin_router = APIRouter(prefix="/admin", tags=["admin"])

@users_router.get("/")
async def list_users():
    return []

@users_router.get("/{id}")
async def get_user(id: int):
    return {"id": id}

@admin_router.get("/stats")
async def stats():
    return {"count": 100}

# 注册路由组
app.include_router(users_router, prefix="/api/v1")
app.include_router(admin_router, dependencies=[Depends(auth)])
```

```go
// ==================== Gin ====================
r := gin.Default()

// API v1 组
v1 := r.Group("/api/v1")
{
    v1.GET("/users", listUsers)
    v1.POST("/users", createUser)
    
    // 嵌套组
    users := v1.Group("/users")
    {
        users.GET("/:id", getUser)
        users.PUT("/:id", updateUser)
        users.DELETE("/:id", deleteUser)
    }
}

// Admin 组 (带中间件)
admin := r.Group("/admin")
admin.Use(AuthMiddleware())
{
    admin.GET("/stats", getStats)
    admin.GET("/users", adminListUsers)
}
```

```rust
// ==================== Axum ====================
use axum::{routing::{get, post}, Router};

// 子路由
fn users_routes() -> Router<AppState> {
    Router::new()
        .route("/", get(list_users).post(create_user))
        .route("/:id", get(get_user).put(update_user).delete(delete_user))
}

fn admin_routes() -> Router<AppState> {
    Router::new()
        .route("/stats", get(stats))
        .layer(middleware::from_fn(auth_middleware))
}

let app = Router::new()
    .nest("/api/v1/users", users_routes())
    .nest("/admin", admin_routes())
    .with_state(state);
```

#### 中间件 (Middleware)

```typescript
// ==================== Koa ====================
// 日志中间件
app.use(async (ctx, next) => {
    const start = Date.now();
    await next();
    const ms = Date.now() - start;
    console.log(`${ctx.method} ${ctx.url} - ${ms}ms`);
});

// 错误处理
app.use(async (ctx, next) => {
    try {
        await next();
    } catch (err) {
        ctx.status = err.status || 500;
        ctx.body = { error: err.message };
    }
});

// 认证中间件
const auth = async (ctx, next) => {
    const token = ctx.get('Authorization')?.replace('Bearer ', '');
    if (!token) {
        ctx.throw(401, 'Unauthorized');
    }
    ctx.state.user = await verifyToken(token);
    await next();
};

router.get('/protected', auth, async (ctx) => {
    ctx.body = { user: ctx.state.user };
});
```

```python
# ==================== FastAPI ====================
from fastapi import Request
from starlette.middleware.base import BaseHTTPMiddleware
import time

# 自定义中间件
class TimingMiddleware(BaseHTTPMiddleware):
    async def dispatch(self, request: Request, call_next):
        start = time.time()
        response = await call_next(request)
        duration = time.time() - start
        response.headers["X-Process-Time"] = str(duration)
        return response

app.add_middleware(TimingMiddleware)

# CORS 中间件
from fastapi.middleware.cors import CORSMiddleware
app.add_middleware(
    CORSMiddleware,
    allow_origins=["*"],
    allow_methods=["*"],
    allow_headers=["*"],
)

# 依赖作为中间件
async def verify_token(authorization: str = Header(...)):
    if not authorization.startswith("Bearer "):
        raise HTTPException(401, "Invalid token")
    return decode_token(authorization[7:])

@app.get("/protected")
async def protected(user: User = Depends(verify_token)):
    return {"user": user}
```

```go
// ==================== Gin ====================
// 日志中间件
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        
        c.Next()  // 处理请求
        
        latency := time.Since(start)
        status := c.Writer.Status()
        log.Printf("%s %s %d %v", c.Request.Method, c.Request.URL, status, latency)
    }
}

// 错误恢复
func Recovery() gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                c.JSON(500, gin.H{"error": "Internal Server Error"})
                c.Abort()
            }
        }()
        c.Next()
    }
}

// 认证中间件
func Auth() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(401, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }
        user, err := verifyToken(token)
        if err != nil {
            c.JSON(401, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }
        c.Set("user", user)
        c.Next()
    }
}

r.Use(Logger(), Recovery())
r.GET("/protected", Auth(), protectedHandler)
```

```rust
// ==================== Axum ====================
use axum::{
    middleware::{self, Next},
    response::Response,
    http::Request,
};
use std::time::Instant;

// 日志中间件
async fn logging<B>(req: Request<B>, next: Next<B>) -> Response {
    let start = Instant::now();
    let method = req.method().clone();
    let uri = req.uri().clone();
    
    let response = next.run(req).await;
    
    let duration = start.elapsed();
    println!("{} {} - {:?}", method, uri, duration);
    
    response
}

// 认证中间件
async fn auth<B>(
    mut req: Request<B>,
    next: Next<B>,
) -> Result<Response, StatusCode> {
    let token = req.headers()
        .get("Authorization")
        .and_then(|h| h.to_str().ok())
        .ok_or(StatusCode::UNAUTHORIZED)?;
    
    let user = verify_token(token)
        .map_err(|_| StatusCode::UNAUTHORIZED)?;
    
    req.extensions_mut().insert(user);
    Ok(next.run(req).await)
}

let app = Router::new()
    .route("/protected", get(protected))
    .layer(middleware::from_fn(auth))
    .layer(middleware::from_fn(logging));
```

### Web 框架对比表

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 功能            │ Koa                  │ FastAPI              │ Gin                  │ Axum                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 路由参数        │ ctx.params           │ 函数参数             │ c.Param              │ Path extractor       │
│ 查询参数        │ ctx.query            │ Query                │ c.ShouldBindQuery    │ Query extractor      │
│ 请求头          │ ctx.get()            │ Header               │ c.GetHeader          │ TypedHeader          │
│ 请求体          │ ctx.request.body     │ BaseModel            │ c.ShouldBindJSON     │ Json extractor       │
│ 应用状态        │ ctx.state            │ Depends / app.state  │ c.Set / c.Get        │ State extractor      │
│ 路由分组        │ Router prefix        │ APIRouter            │ r.Group              │ Router::nest         │
│ 中间件          │ app.use(fn)          │ add_middleware       │ r.Use                │ .layer               │
│ 验证            │ 第三方库             │ Pydantic             │ binding tag          │ validator crate      │
│ 自动文档        │ ❌                   │ ✅ OpenAPI           │ swag                 │ utoipa               │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
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

### 并发控制模式 (All/Race/Timeout/Semaphore/Bucket-Limit)

#### All - 等待所有完成

**TypeScript**
```typescript
// ==================== Promise.all ====================
// 等待所有 Promise 完成，任一失败则整体失败
const results = await Promise.all([
    fetch('/api/users'),
    fetch('/api/posts'),
    fetch('/api/comments')
]);

// Promise.allSettled - 等待所有完成，不管成功失败
const settled = await Promise.allSettled([
    fetch('/api/users'),
    fetch('/api/might-fail'),
    fetch('/api/posts')
]);
// 结果: [{ status: 'fulfilled', value }, { status: 'rejected', reason }, ...]

// 带类型的 Promise.all
async function fetchAll<T>(urls: string[]): Promise<T[]> {
    return Promise.all(urls.map(url => fetch(url).then(r => r.json())));
}

// 并发限制的 all
async function allWithLimit<T>(
    tasks: (() => Promise<T>)[],
    limit: number
): Promise<T[]> {
    const results: T[] = [];
    const executing: Promise<void>[] = [];
    
    for (const task of tasks) {
        const p = task().then(result => {
            results.push(result);
        });
        executing.push(p);
        
        if (executing.length >= limit) {
            await Promise.race(executing);
            executing.splice(executing.findIndex(e => e === p), 1);
        }
    }
    await Promise.all(executing);
    return results;
}
```

**Python**
```python
import asyncio
from typing import List, Any, Coroutine

# ==================== asyncio.gather ====================
# 等待所有协程完成
async def fetch_all():
    results = await asyncio.gather(
        fetch_users(),
        fetch_posts(),
        fetch_comments()
    )
    return results

# return_exceptions=True 类似 allSettled
async def fetch_all_settled():
    results = await asyncio.gather(
        fetch_users(),
        might_fail(),
        fetch_posts(),
        return_exceptions=True  # 异常作为结果返回，不抛出
    )
    for r in results:
        if isinstance(r, Exception):
            print(f"Failed: {r}")
        else:
            print(f"Success: {r}")

# 使用 TaskGroup (Python 3.11+)
async def fetch_with_taskgroup():
    async with asyncio.TaskGroup() as tg:
        task1 = tg.create_task(fetch_users())
        task2 = tg.create_task(fetch_posts())
        task3 = tg.create_task(fetch_comments())
    # 所有任务完成后继续
    return task1.result(), task2.result(), task3.result()

# 带并发限制
async def gather_with_limit(coros: List[Coroutine], limit: int):
    semaphore = asyncio.Semaphore(limit)
    
    async def limited_coro(coro):
        async with semaphore:
            return await coro
    
    return await asyncio.gather(*[limited_coro(c) for c in coros])
```

**Go**
```go
// ==================== WaitGroup - 等待所有完成 ====================
import (
    "sync"
    "golang.org/x/sync/errgroup"
)

// 基础 WaitGroup
func fetchAll() {
    var wg sync.WaitGroup
    results := make([]string, 3)
    
    urls := []string{"/api/users", "/api/posts", "/api/comments"}
    
    for i, url := range urls {
        wg.Add(1)
        go func(idx int, u string) {
            defer wg.Done()
            results[idx] = fetch(u)
        }(i, url)
    }
    
    wg.Wait()  // 等待所有完成
    fmt.Println(results)
}

// errgroup - 带错误处理的等待组
func fetchAllWithError() error {
    g, ctx := errgroup.WithContext(context.Background())
    
    var users, posts, comments string
    
    g.Go(func() error {
        var err error
        users, err = fetchWithCtx(ctx, "/api/users")
        return err
    })
    
    g.Go(func() error {
        var err error
        posts, err = fetchWithCtx(ctx, "/api/posts")
        return err
    })
    
    g.Go(func() error {
        var err error
        comments, err = fetchWithCtx(ctx, "/api/comments")
        return err
    })
    
    if err := g.Wait(); err != nil {
        return err  // 任一失败返回错误
    }
    
    fmt.Println(users, posts, comments)
    return nil
}

// 带并发限制的 errgroup
func fetchAllLimited() error {
    g, ctx := errgroup.WithContext(context.Background())
    g.SetLimit(3)  // 最多 3 个并发
    
    urls := []string{...}
    results := make([]string, len(urls))
    
    for i, url := range urls {
        i, url := i, url
        g.Go(func() error {
            result, err := fetchWithCtx(ctx, url)
            if err != nil {
                return err
            }
            results[i] = result
            return nil
        })
    }
    
    return g.Wait()
}
```

**Rust**
```rust
use futures::future::{join_all, try_join_all};
use tokio::task::JoinSet;

// ==================== join_all - 等待所有完成 ====================
async fn fetch_all() -> Vec<String> {
    let futures = vec![
        fetch_users(),
        fetch_posts(),
        fetch_comments(),
    ];
    
    // join_all 等待所有完成
    join_all(futures).await
}

// try_join_all - 任一失败则返回错误
async fn fetch_all_or_fail() -> Result<Vec<String>, Error> {
    let futures = vec![
        fetch_users(),
        fetch_posts(),
        fetch_comments(),
    ];
    
    try_join_all(futures).await
}

// tokio::join! 宏 - 固定数量的 futures
async fn fetch_multiple() {
    let (users, posts, comments) = tokio::join!(
        fetch_users(),
        fetch_posts(),
        fetch_comments()
    );
}

// try_join! 宏 - 带错误处理
async fn fetch_multiple_or_fail() -> Result<(), Error> {
    let (users, posts, comments) = tokio::try_join!(
        fetch_users(),
        fetch_posts(),
        fetch_comments()
    )?;
    Ok(())
}

// JoinSet - 动态任务管理
async fn fetch_dynamic(urls: Vec<String>) -> Vec<String> {
    let mut set = JoinSet::new();
    
    for url in urls {
        set.spawn(async move {
            fetch(&url).await
        });
    }
    
    let mut results = Vec::new();
    while let Some(res) = set.join_next().await {
        if let Ok(value) = res {
            results.push(value);
        }
    }
    results
}
```

#### Race - 返回最快的结果

**TypeScript**
```typescript
// ==================== Promise.race ====================
// 返回第一个完成的（成功或失败）
const fastest = await Promise.race([
    fetch('/api/server1/data'),
    fetch('/api/server2/data'),
    fetch('/api/server3/data')
]);

// Promise.any - 返回第一个成功的（忽略失败）
const firstSuccess = await Promise.any([
    fetch('/api/primary'),
    fetch('/api/backup1'),
    fetch('/api/backup2')
]);
// 全部失败才抛出 AggregateError

// 实现带超时的 race
function raceWithTimeout<T>(promise: Promise<T>, ms: number): Promise<T> {
    const timeout = new Promise<never>((_, reject) => {
        setTimeout(() => reject(new Error('Timeout')), ms);
    });
    return Promise.race([promise, timeout]);
}

// 多服务器竞速获取
async function fetchFromFastestServer<T>(urls: string[]): Promise<T> {
    const controller = new AbortController();
    
    try {
        const result = await Promise.any(
            urls.map(url => 
                fetch(url, { signal: controller.signal })
                    .then(r => r.json())
            )
        );
        return result;
    } finally {
        controller.abort();  // 取消其他请求
    }
}
```

**Python**
```python
import asyncio

# ==================== asyncio.wait - FIRST_COMPLETED ====================
async def race_tasks():
    tasks = [
        asyncio.create_task(fetch_from_server1()),
        asyncio.create_task(fetch_from_server2()),
        asyncio.create_task(fetch_from_server3()),
    ]
    
    # 等待第一个完成
    done, pending = await asyncio.wait(
        tasks,
        return_when=asyncio.FIRST_COMPLETED
    )
    
    # 取消其他任务
    for task in pending:
        task.cancel()
    
    # 获取结果
    return done.pop().result()

# 类似 Promise.any - 返回第一个成功的
async def first_success(coros):
    tasks = [asyncio.create_task(c) for c in coros]
    
    while tasks:
        done, tasks = await asyncio.wait(
            tasks,
            return_when=asyncio.FIRST_COMPLETED
        )
        
        for task in done:
            if not task.exception():
                # 取消剩余任务
                for t in tasks:
                    t.cancel()
                return task.result()
    
    raise Exception("All tasks failed")

# 使用 asyncio.wait_for 实现 race with timeout
async def race_with_timeout(coros, timeout):
    tasks = [asyncio.create_task(c) for c in coros]
    try:
        done, pending = await asyncio.wait(
            tasks,
            timeout=timeout,
            return_when=asyncio.FIRST_COMPLETED
        )
        if done:
            return done.pop().result()
        raise asyncio.TimeoutError()
    finally:
        for task in pending:
            task.cancel()
```

**Go**
```go
// ==================== select - 竞速选择 ====================
func raceRequests(ctx context.Context) (string, error) {
    ch := make(chan string, 3)
    errCh := make(chan error, 3)
    
    // 启动多个请求
    go func() {
        result, err := fetchFromServer1(ctx)
        if err != nil {
            errCh <- err
            return
        }
        ch <- result
    }()
    
    go func() {
        result, err := fetchFromServer2(ctx)
        if err != nil {
            errCh <- err
            return
        }
        ch <- result
    }()
    
    // 等待第一个结果
    select {
    case result := <-ch:
        return result, nil
    case err := <-errCh:
        return "", err
    case <-ctx.Done():
        return "", ctx.Err()
    }
}

// 使用 context 取消剩余请求
func raceFetch(urls []string) (string, error) {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()  // 第一个完成后取消其他
    
    ch := make(chan string, len(urls))
    
    for _, url := range urls {
        url := url
        go func() {
            if result, err := fetchWithCtx(ctx, url); err == nil {
                select {
                case ch <- result:
                default:
                }
            }
        }()
    }
    
    select {
    case result := <-ch:
        return result, nil
    case <-time.After(10 * time.Second):
        return "", errors.New("all requests timed out")
    }
}
```

**Rust**
```rust
use tokio::select;
use futures::future::select_all;

// ==================== select! 宏 - 竞速 ====================
async fn race_requests() -> Result<String, Error> {
    select! {
        result = fetch_from_server1() => result,
        result = fetch_from_server2() => result,
        result = fetch_from_server3() => result,
    }
}

// select_all - 动态数量的 futures 竞速
async fn race_dynamic(urls: Vec<String>) -> String {
    let futures: Vec<_> = urls
        .into_iter()
        .map(|url| Box::pin(fetch(&url)))
        .collect();
    
    let (result, _index, _remaining) = select_all(futures).await;
    result
}

// 带取消的竞速
async fn race_with_cancel(urls: Vec<String>) -> Result<String, Error> {
    let token = CancellationToken::new();
    let mut set = JoinSet::new();
    
    for url in urls {
        let token = token.clone();
        set.spawn(async move {
            select! {
                result = fetch(&url) => result,
                _ = token.cancelled() => Err(Error::Cancelled),
            }
        });
    }
    
    if let Some(Ok(Ok(result))) = set.join_next().await {
        token.cancel();  // 取消其他任务
        return Ok(result);
    }
    
    Err(Error::AllFailed)
}
```

#### Timeout - 超时控制

**TypeScript**
```typescript
// ==================== 超时控制 ====================
// 基础超时封装
function withTimeout<T>(promise: Promise<T>, ms: number): Promise<T> {
    return new Promise((resolve, reject) => {
        const timer = setTimeout(() => {
            reject(new Error(`Timeout after ${ms}ms`));
        }, ms);
        
        promise
            .then(resolve)
            .catch(reject)
            .finally(() => clearTimeout(timer));
    });
}

// 使用 AbortController 实现可取消超时
async function fetchWithTimeout(url: string, ms: number): Promise<Response> {
    const controller = new AbortController();
    const timeoutId = setTimeout(() => controller.abort(), ms);
    
    try {
        return await fetch(url, { signal: controller.signal });
    } finally {
        clearTimeout(timeoutId);
    }
}

// AbortSignal.timeout (现代浏览器)
async function fetchModern(url: string): Promise<Response> {
    return fetch(url, { signal: AbortSignal.timeout(5000) });
}

// 重试 + 超时
async function fetchWithRetry(
    url: string,
    options: { timeout: number; retries: number }
): Promise<Response> {
    for (let i = 0; i < options.retries; i++) {
        try {
            return await fetchWithTimeout(url, options.timeout);
        } catch (err) {
            if (i === options.retries - 1) throw err;
            await new Promise(r => setTimeout(r, 1000 * (i + 1)));  // 退避
        }
    }
    throw new Error('Unreachable');
}
```

**Python**
```python
import asyncio
from contextlib import asynccontextmanager

# ==================== asyncio.timeout (Python 3.11+) ====================
async def fetch_with_timeout():
    async with asyncio.timeout(5.0):  # 5秒超时
        return await fetch_data()

# asyncio.wait_for (兼容旧版本)
async def fetch_with_wait_for():
    try:
        result = await asyncio.wait_for(fetch_data(), timeout=5.0)
        return result
    except asyncio.TimeoutError:
        print("Request timed out")
        raise

# 自定义超时上下文管理器
@asynccontextmanager
async def timeout_context(seconds: float):
    task = asyncio.current_task()
    loop = asyncio.get_running_loop()
    
    def cancel_task():
        task.cancel()
    
    handle = loop.call_later(seconds, cancel_task)
    try:
        yield
    finally:
        handle.cancel()

# 使用
async def example():
    async with timeout_context(5.0):
        await long_running_operation()

# 带重试的超时
async def fetch_with_retry(url: str, timeout: float, retries: int):
    for i in range(retries):
        try:
            async with asyncio.timeout(timeout):
                return await fetch(url)
        except asyncio.TimeoutError:
            if i == retries - 1:
                raise
            await asyncio.sleep(1.0 * (i + 1))  # 指数退避
```

**Go**
```go
import (
    "context"
    "time"
)

// ==================== context.WithTimeout ====================
func fetchWithTimeout(url string) (string, error) {
    // 创建带超时的 context
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()
    
    return fetchWithCtx(ctx, url)
}

// 在函数中检查超时
func fetchWithCtx(ctx context.Context, url string) (string, error) {
    // 创建请求
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return "", err
    }
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return "", err  // 超时会返回 context deadline exceeded
    }
    defer resp.Body.Close()
    
    body, _ := io.ReadAll(resp.Body)
    return string(body), nil
}

// select 实现超时
func fetchWithSelectTimeout(url string) (string, error) {
    ch := make(chan string, 1)
    errCh := make(chan error, 1)
    
    go func() {
        result, err := fetch(url)
        if err != nil {
            errCh <- err
            return
        }
        ch <- result
    }()
    
    select {
    case result := <-ch:
        return result, nil
    case err := <-errCh:
        return "", err
    case <-time.After(5 * time.Second):
        return "", errors.New("timeout")
    }
}

// 带重试的超时
func fetchWithRetry(url string, timeout time.Duration, retries int) (string, error) {
    var lastErr error
    
    for i := 0; i < retries; i++ {
        ctx, cancel := context.WithTimeout(context.Background(), timeout)
        result, err := fetchWithCtx(ctx, url)
        cancel()
        
        if err == nil {
            return result, nil
        }
        
        lastErr = err
        time.Sleep(time.Duration(i+1) * time.Second)  // 退避
    }
    
    return "", lastErr
}
```

**Rust**
```rust
use tokio::time::{timeout, Duration};
use std::time::Instant;

// ==================== tokio::time::timeout ====================
async fn fetch_with_timeout(url: &str) -> Result<String, Error> {
    match timeout(Duration::from_secs(5), fetch(url)).await {
        Ok(result) => result,
        Err(_) => Err(Error::Timeout),
    }
}

// 使用 select! 实现超时
async fn fetch_with_select_timeout(url: &str) -> Result<String, Error> {
    select! {
        result = fetch(url) => result,
        _ = tokio::time::sleep(Duration::from_secs(5)) => {
            Err(Error::Timeout)
        }
    }
}

// 带取消令牌的超时
async fn fetch_cancellable(url: &str, token: CancellationToken) -> Result<String, Error> {
    select! {
        result = fetch(url) => result,
        _ = token.cancelled() => Err(Error::Cancelled),
        _ = tokio::time::sleep(Duration::from_secs(5)) => {
            Err(Error::Timeout)
        }
    }
}

// 带重试的超时
async fn fetch_with_retry(
    url: &str,
    timeout_duration: Duration,
    retries: u32,
) -> Result<String, Error> {
    let mut last_err = Error::Unknown;
    
    for i in 0..retries {
        match timeout(timeout_duration, fetch(url)).await {
            Ok(Ok(result)) => return Ok(result),
            Ok(Err(e)) => last_err = e,
            Err(_) => last_err = Error::Timeout,
        }
        
        tokio::time::sleep(Duration::from_secs((i + 1) as u64)).await;
    }
    
    Err(last_err)
}

// deadline 而非 duration
async fn fetch_with_deadline(url: &str) -> Result<String, Error> {
    let deadline = Instant::now() + Duration::from_secs(5);
    tokio::time::timeout_at(deadline.into(), fetch(url)).await?
}
```

#### Cancel - 取消操作

**TypeScript**
```typescript
// ==================== AbortController ====================
// 基础取消
const controller = new AbortController();
const { signal } = controller;

// 发起可取消的请求
fetch('/api/data', { signal })
    .then(response => response.json())
    .catch(err => {
        if (err.name === 'AbortError') {
            console.log('Request was cancelled');
        }
    });

// 取消请求
controller.abort();

// 带原因的取消
controller.abort(new Error('User cancelled'));

// ==================== 多请求共享取消 ====================
class CancellableRequestManager {
    private controller: AbortController | null = null;
    
    async fetch(url: string): Promise<Response> {
        // 取消之前的请求
        this.controller?.abort();
        this.controller = new AbortController();
        
        return fetch(url, { signal: this.controller.signal });
    }
    
    cancel(): void {
        this.controller?.abort();
        this.controller = null;
    }
}

// ==================== 可取消的 Promise ====================
function cancellablePromise<T>(
    executor: (signal: AbortSignal) => Promise<T>
): { promise: Promise<T>; cancel: () => void } {
    const controller = new AbortController();
    
    const promise = executor(controller.signal);
    
    return {
        promise,
        cancel: () => controller.abort()
    };
}

// 使用
const { promise, cancel } = cancellablePromise(async (signal) => {
    const response = await fetch('/api/data', { signal });
    return response.json();
});

// 需要时取消
cancel();

// ==================== 链接多个 AbortSignal ====================
function mergeSignals(...signals: AbortSignal[]): AbortSignal {
    const controller = new AbortController();
    
    for (const signal of signals) {
        if (signal.aborted) {
            controller.abort(signal.reason);
            break;
        }
        signal.addEventListener('abort', () => {
            controller.abort(signal.reason);
        }, { once: true });
    }
    
    return controller.signal;
}

// AbortSignal.any (现代浏览器)
const combined = AbortSignal.any([signal1, signal2, signal3]);

// ==================== 可取消的异步迭代 ====================
async function* cancellableFetch(
    urls: string[],
    signal: AbortSignal
): AsyncGenerator<Response> {
    for (const url of urls) {
        if (signal.aborted) {
            throw new Error('Cancelled');
        }
        yield await fetch(url, { signal });
    }
}

// 使用
const controller = new AbortController();
for await (const response of cancellableFetch(urls, controller.signal)) {
    // 处理响应
    if (shouldStop) {
        controller.abort();
        break;
    }
}
```

**Python**
```python
import asyncio
from contextlib import asynccontextmanager
from typing import Optional

# ==================== asyncio.Task.cancel ====================
async def cancellable_operation():
    task = asyncio.create_task(long_running_operation())
    
    # 稍后取消
    await asyncio.sleep(1)
    task.cancel()
    
    try:
        await task
    except asyncio.CancelledError:
        print("Task was cancelled")

# ==================== 处理取消 ====================
async def graceful_cancel():
    try:
        await long_running_operation()
    except asyncio.CancelledError:
        # 清理资源
        await cleanup()
        raise  # 重新抛出让调用者知道被取消了

# 屏蔽取消（谨慎使用）
async def unshieldable_operation():
    try:
        # 这部分不能被取消
        await asyncio.shield(critical_operation())
    except asyncio.CancelledError:
        print("Outer cancelled, but inner completed")
        raise

# ==================== 取消令牌模式 ====================
class CancellationToken:
    def __init__(self):
        self._cancelled = False
        self._event = asyncio.Event()
    
    def cancel(self):
        self._cancelled = True
        self._event.set()
    
    @property
    def is_cancelled(self) -> bool:
        return self._cancelled
    
    async def wait(self):
        """等待取消"""
        await self._event.wait()
    
    def check(self):
        """检查是否取消，是则抛出异常"""
        if self._cancelled:
            raise asyncio.CancelledError("Operation cancelled")

# 使用取消令牌
async def long_operation(token: CancellationToken):
    for i in range(100):
        token.check()  # 检查点
        await asyncio.sleep(0.1)
        # 执行工作...

# ==================== 取消多个任务 ====================
async def cancel_all_tasks():
    tasks = [
        asyncio.create_task(fetch(url))
        for url in urls
    ]
    
    # 等待第一个完成
    done, pending = await asyncio.wait(
        tasks,
        return_when=asyncio.FIRST_COMPLETED
    )
    
    # 取消其余任务
    for task in pending:
        task.cancel()
    
    # 等待取消完成
    await asyncio.gather(*pending, return_exceptions=True)
    
    return done.pop().result()

# ==================== TaskGroup 取消 (Python 3.11+) ====================
async def taskgroup_cancel():
    try:
        async with asyncio.TaskGroup() as tg:
            task1 = tg.create_task(operation1())
            task2 = tg.create_task(operation2())
            task3 = tg.create_task(operation3())
            # 如果任一任务失败，其他都会被取消
    except* ValueError as eg:
        print(f"Some tasks failed: {eg.exceptions}")

# ==================== 超时自动取消 ====================
async def auto_cancel_on_timeout():
    async with asyncio.timeout(5.0) as cm:
        await long_operation()
    
    if cm.expired():
        print("Operation was cancelled due to timeout")
```

**Go**
```go
import (
    "context"
    "errors"
)

// ==================== context.WithCancel ====================
func cancellableOperation() {
    ctx, cancel := context.WithCancel(context.Background())
    
    go func() {
        // 模拟某个条件触发取消
        time.Sleep(2 * time.Second)
        cancel()
    }()
    
    // 执行可取消的操作
    result, err := fetchWithCtx(ctx, "/api/data")
    if errors.Is(err, context.Canceled) {
        fmt.Println("Operation was cancelled")
    }
}

// ==================== 检查取消状态 ====================
func longOperation(ctx context.Context) error {
    for i := 0; i < 100; i++ {
        select {
        case <-ctx.Done():
            return ctx.Err()  // context.Canceled 或 context.DeadlineExceeded
        default:
            // 继续工作
            time.Sleep(100 * time.Millisecond)
            doWork(i)
        }
    }
    return nil
}

// ==================== 传播取消 ====================
func parentOperation(ctx context.Context) error {
    // 子操作会继承父的取消信号
    g, ctx := errgroup.WithContext(ctx)
    
    g.Go(func() error {
        return childOperation1(ctx)
    })
    
    g.Go(func() error {
        return childOperation2(ctx)
    })
    
    // 任一子操作失败或被取消，其他都会收到取消信号
    return g.Wait()
}

// ==================== 带原因的取消 ====================
func cancelWithCause() {
    ctx, cancel := context.WithCancelCause(context.Background())
    
    go func() {
        cancel(errors.New("user requested cancellation"))
    }()
    
    <-ctx.Done()
    
    // 获取取消原因
    cause := context.Cause(ctx)
    fmt.Printf("Cancelled because: %v\n", cause)
}

// ==================== 优雅关闭模式 ====================
type Worker struct {
    ctx    context.Context
    cancel context.CancelFunc
    done   chan struct{}
}

func NewWorker() *Worker {
    ctx, cancel := context.WithCancel(context.Background())
    w := &Worker{
        ctx:    ctx,
        cancel: cancel,
        done:   make(chan struct{}),
    }
    go w.run()
    return w
}

func (w *Worker) run() {
    defer close(w.done)
    
    for {
        select {
        case <-w.ctx.Done():
            // 清理资源
            w.cleanup()
            return
        default:
            w.doWork()
        }
    }
}

func (w *Worker) Stop() {
    w.cancel()
    <-w.done  // 等待完全停止
}

func (w *Worker) cleanup() {
    fmt.Println("Cleaning up...")
}

func (w *Worker) doWork() {
    // 工作逻辑
}

// ==================== HTTP 请求取消 ====================
func cancellableHTTPRequest(ctx context.Context, url string) ([]byte, error) {
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return nil, err
    }
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        if errors.Is(err, context.Canceled) {
            return nil, fmt.Errorf("request cancelled: %w", err)
        }
        return nil, err
    }
    defer resp.Body.Close()
    
    return io.ReadAll(resp.Body)
}

// ==================== Channel 方式取消 ====================
func workerWithStopChannel(stop <-chan struct{}) {
    for {
        select {
        case <-stop:
            fmt.Println("Worker stopped")
            return
        default:
            doWork()
        }
    }
}

// 使用
func main() {
    stop := make(chan struct{})
    go workerWithStopChannel(stop)
    
    time.Sleep(5 * time.Second)
    close(stop)  // 发送停止信号
}
```

**Rust**
```rust
use tokio::select;
use tokio_util::sync::CancellationToken;
use std::sync::Arc;

// ==================== CancellationToken ====================
async fn cancellable_operation(token: CancellationToken) -> Result<String, Error> {
    select! {
        result = fetch_data() => result,
        _ = token.cancelled() => Err(Error::Cancelled),
    }
}

// 使用
async fn main() {
    let token = CancellationToken::new();
    let token_clone = token.clone();
    
    // 启动可取消的任务
    let handle = tokio::spawn(async move {
        cancellable_operation(token_clone).await
    });
    
    // 稍后取消
    tokio::time::sleep(Duration::from_secs(2)).await;
    token.cancel();
    
    match handle.await {
        Ok(Ok(result)) => println!("Got result: {}", result),
        Ok(Err(Error::Cancelled)) => println!("Operation was cancelled"),
        _ => println!("Task failed"),
    }
}

// ==================== 子令牌（层级取消） ====================
async fn hierarchical_cancel() {
    let parent_token = CancellationToken::new();
    let child_token = parent_token.child_token();
    
    // 取消父令牌会同时取消子令牌
    tokio::spawn({
        let token = child_token.clone();
        async move {
            token.cancelled().await;
            println!("Child task cancelled");
        }
    });
    
    parent_token.cancel();  // 子任务也会被取消
}

// ==================== JoinSet 取消 ====================
async fn cancel_joinset() {
    let mut set = tokio::task::JoinSet::new();
    
    for i in 0..10 {
        set.spawn(async move {
            tokio::time::sleep(Duration::from_secs(i)).await;
            i
        });
    }
    
    // 等待第一个完成
    if let Some(result) = set.join_next().await {
        println!("First result: {:?}", result);
    }
    
    // 取消剩余所有任务
    set.abort_all();
    
    // 等待所有取消完成
    while set.join_next().await.is_some() {}
}

// ==================== tokio::task::abort ====================
async fn abort_task() {
    let handle = tokio::spawn(async {
        loop {
            tokio::time::sleep(Duration::from_secs(1)).await;
            println!("Working...");
        }
    });
    
    tokio::time::sleep(Duration::from_secs(3)).await;
    handle.abort();  // 取消任务
    
    match handle.await {
        Ok(_) => println!("Task completed"),
        Err(e) if e.is_cancelled() => println!("Task was cancelled"),
        Err(e) => println!("Task failed: {}", e),
    }
}

// ==================== 检查点取消 ====================
async fn long_operation(token: &CancellationToken) -> Result<(), Error> {
    for i in 0..100 {
        // 检查点
        if token.is_cancelled() {
            return Err(Error::Cancelled);
        }
        
        // 或者使用 select 在异步点检查
        select! {
            _ = async_work(i) => {},
            _ = token.cancelled() => return Err(Error::Cancelled),
        }
    }
    Ok(())
}

// ==================== Drop 时自动取消 ====================
struct CancellableTask {
    token: CancellationToken,
    handle: tokio::task::JoinHandle<()>,
}

impl CancellableTask {
    fn new<F>(future: F) -> Self
    where
        F: std::future::Future<Output = ()> + Send + 'static,
    {
        let token = CancellationToken::new();
        let token_clone = token.clone();
        
        let handle = tokio::spawn(async move {
            select! {
                _ = future => {},
                _ = token_clone.cancelled() => {},
            }
        });
        
        Self { token, handle }
    }
    
    fn cancel(&self) {
        self.token.cancel();
    }
}

impl Drop for CancellableTask {
    fn drop(&mut self) {
        self.token.cancel();
    }
}

// ==================== 优雅关闭 ====================
async fn graceful_shutdown(token: CancellationToken) {
    // 等待取消信号
    token.cancelled().await;
    
    println!("Shutdown signal received, cleaning up...");
    
    // 执行清理，但设置最大等待时间
    let cleanup_result = tokio::time::timeout(
        Duration::from_secs(30),
        cleanup_resources()
    ).await;
    
    match cleanup_result {
        Ok(_) => println!("Cleanup completed"),
        Err(_) => println!("Cleanup timed out, forcing shutdown"),
    }
}

// 信号处理
async fn run_with_shutdown() {
    let token = CancellationToken::new();
    
    // 监听 Ctrl+C
    let shutdown_token = token.clone();
    tokio::spawn(async move {
        tokio::signal::ctrl_c().await.expect("Failed to listen for Ctrl+C");
        shutdown_token.cancel();
    });
    
    // 运行主逻辑
    select! {
        _ = main_loop() => {},
        _ = token.cancelled() => {
            println!("Shutting down...");
        }
    }
}
```

#### Semaphore - 信号量（并发数控制）

**TypeScript**
```typescript
// ==================== Semaphore 实现 ====================
class Semaphore {
    private permits: number;
    private queue: (() => void)[] = [];
    
    constructor(permits: number) {
        this.permits = permits;
    }
    
    async acquire(): Promise<void> {
        if (this.permits > 0) {
            this.permits--;
            return;
        }
        
        return new Promise(resolve => {
            this.queue.push(resolve);
        });
    }
    
    release(): void {
        const next = this.queue.shift();
        if (next) {
            next();
        } else {
            this.permits++;
        }
    }
    
    async withPermit<T>(fn: () => Promise<T>): Promise<T> {
        await this.acquire();
        try {
            return await fn();
        } finally {
            this.release();
        }
    }
}

// 使用 Semaphore 控制并发
async function fetchAllLimited(urls: string[], limit: number): Promise<string[]> {
    const semaphore = new Semaphore(limit);
    
    return Promise.all(
        urls.map(url => 
            semaphore.withPermit(() => fetch(url).then(r => r.text()))
        )
    );
}

// p-limit 库的使用（推荐）
import pLimit from 'p-limit';

const limit = pLimit(5);  // 最多 5 个并发
const results = await Promise.all(
    urls.map(url => limit(() => fetch(url)))
);
```

**Python**
```python
import asyncio
from contextlib import asynccontextmanager

# ==================== asyncio.Semaphore ====================
async def fetch_with_semaphore(urls: list[str], limit: int):
    semaphore = asyncio.Semaphore(limit)
    
    async def fetch_one(url: str):
        async with semaphore:  # 自动获取和释放
            return await fetch(url)
    
    return await asyncio.gather(*[fetch_one(url) for url in urls])

# BoundedSemaphore - 防止多次释放
async def safe_semaphore_example():
    sem = asyncio.BoundedSemaphore(3)
    
    async with sem:
        await do_work()
    
    # sem.release()  # 额外释放会抛出 ValueError

# 手动获取和释放
async def manual_semaphore():
    sem = asyncio.Semaphore(3)
    
    await sem.acquire()
    try:
        await do_work()
    finally:
        sem.release()

# 线程版本
import threading

def threaded_work(urls: list[str], limit: int):
    semaphore = threading.Semaphore(limit)
    results = []
    threads = []
    
    def worker(url: str):
        with semaphore:
            result = fetch_sync(url)
            results.append(result)
    
    for url in urls:
        t = threading.Thread(target=worker, args=(url,))
        threads.append(t)
        t.start()
    
    for t in threads:
        t.join()
    
    return results
```

**Go**
```go
import (
    "context"
    "golang.org/x/sync/semaphore"
)

// ==================== 使用 channel 实现信号量 ====================
func fetchWithSemaphore(urls []string, limit int) []string {
    sem := make(chan struct{}, limit)
    results := make([]string, len(urls))
    var wg sync.WaitGroup
    
    for i, url := range urls {
        wg.Add(1)
        go func(idx int, u string) {
            defer wg.Done()
            
            sem <- struct{}{}        // 获取许可
            defer func() { <-sem }() // 释放许可
            
            results[idx] = fetch(u)
        }(i, url)
    }
    
    wg.Wait()
    return results
}

// 使用 golang.org/x/sync/semaphore 包
func fetchWithWeightedSemaphore(urls []string, limit int64) ([]string, error) {
    sem := semaphore.NewWeighted(limit)
    ctx := context.Background()
    results := make([]string, len(urls))
    var wg sync.WaitGroup
    
    for i, url := range urls {
        wg.Add(1)
        go func(idx int, u string) {
            defer wg.Done()
            
            // 获取 1 个许可
            if err := sem.Acquire(ctx, 1); err != nil {
                return
            }
            defer sem.Release(1)
            
            results[idx] = fetch(u)
        }(i, url)
    }
    
    wg.Wait()
    return results, nil
}

// 加权信号量 - 不同任务消耗不同许可
func weightedTasks(tasks []Task) error {
    sem := semaphore.NewWeighted(100)  // 总容量 100
    ctx := context.Background()
    
    var wg sync.WaitGroup
    for _, task := range tasks {
        wg.Add(1)
        go func(t Task) {
            defer wg.Done()
            
            // 大任务消耗更多许可
            weight := int64(t.Weight)
            if err := sem.Acquire(ctx, weight); err != nil {
                return
            }
            defer sem.Release(weight)
            
            t.Execute()
        }(task)
    }
    
    wg.Wait()
    return nil
}

// 带超时的信号量获取
func acquireWithTimeout(sem *semaphore.Weighted, timeout time.Duration) error {
    ctx, cancel := context.WithTimeout(context.Background(), timeout)
    defer cancel()
    
    return sem.Acquire(ctx, 1)
}
```

**Rust**
```rust
use tokio::sync::{Semaphore, SemaphorePermit, OwnedSemaphorePermit};
use std::sync::Arc;

// ==================== tokio::sync::Semaphore ====================
async fn fetch_with_semaphore(urls: Vec<String>, limit: usize) -> Vec<String> {
    let semaphore = Arc::new(Semaphore::new(limit));
    let mut handles = vec![];
    
    for url in urls {
        let sem = semaphore.clone();
        handles.push(tokio::spawn(async move {
            let _permit = sem.acquire().await.unwrap();
            fetch(&url).await
        }));
    }
    
    let mut results = vec![];
    for handle in handles {
        if let Ok(result) = handle.await {
            results.push(result);
        }
    }
    results
}

// OwnedSemaphorePermit - 所有权转移
async fn fetch_owned_permit(
    url: String,
    semaphore: Arc<Semaphore>,
) -> (String, OwnedSemaphorePermit) {
    let permit = semaphore.clone().acquire_owned().await.unwrap();
    let result = fetch(&url).await;
    (result, permit)  // permit 可以被移动
}

// try_acquire - 非阻塞获取
async fn try_fetch(url: &str, semaphore: &Semaphore) -> Option<String> {
    match semaphore.try_acquire() {
        Ok(_permit) => Some(fetch(url).await),
        Err(_) => None,  // 没有可用许可
    }
}

// 带超时的获取
async fn acquire_with_timeout(
    semaphore: &Semaphore,
    timeout: Duration,
) -> Result<SemaphorePermit, Error> {
    match tokio::time::timeout(timeout, semaphore.acquire()).await {
        Ok(Ok(permit)) => Ok(permit),
        Ok(Err(_)) => Err(Error::SemaphoreClosed),
        Err(_) => Err(Error::Timeout),
    }
}

// 使用 std 的 Semaphore（同步）
use std::sync::Semaphore as StdSemaphore;

fn sync_fetch_limited(urls: Vec<String>, limit: usize) -> Vec<String> {
    let sem = Arc::new(StdSemaphore::new(limit));
    let results = Arc::new(Mutex::new(vec![]));
    let mut handles = vec![];
    
    for url in urls {
        let sem = sem.clone();
        let results = results.clone();
        handles.push(std::thread::spawn(move || {
            let _permit = sem.acquire().unwrap();
            let result = fetch_sync(&url);
            results.lock().unwrap().push(result);
        }));
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    Arc::try_unwrap(results).unwrap().into_inner().unwrap()
}
```

#### Bucket-Limit - 令牌桶/漏桶限流

**TypeScript**
```typescript
// ==================== 令牌桶 (Token Bucket) ====================
class TokenBucket {
    private tokens: number;
    private lastRefill: number;
    
    constructor(
        private capacity: number,      // 桶容量
        private refillRate: number,    // 每秒填充速率
    ) {
        this.tokens = capacity;
        this.lastRefill = Date.now();
    }
    
    private refill(): void {
        const now = Date.now();
        const elapsed = (now - this.lastRefill) / 1000;
        this.tokens = Math.min(this.capacity, this.tokens + elapsed * this.refillRate);
        this.lastRefill = now;
    }
    
    tryAcquire(tokens: number = 1): boolean {
        this.refill();
        if (this.tokens >= tokens) {
            this.tokens -= tokens;
            return true;
        }
        return false;
    }
    
    async acquire(tokens: number = 1): Promise<void> {
        while (!this.tryAcquire(tokens)) {
            await new Promise(r => setTimeout(r, 100));
        }
    }
}

// 使用令牌桶限流 API 调用
const bucket = new TokenBucket(10, 2);  // 容量10，每秒补充2个

async function rateLimitedFetch(url: string): Promise<Response> {
    await bucket.acquire();
    return fetch(url);
}

// ==================== 漏桶 (Leaky Bucket) ====================
class LeakyBucket {
    private queue: (() => void)[] = [];
    private processing = false;
    
    constructor(
        private ratePerSecond: number  // 每秒处理数量
    ) {}
    
    async add<T>(task: () => Promise<T>): Promise<T> {
        return new Promise((resolve, reject) => {
            this.queue.push(async () => {
                try {
                    resolve(await task());
                } catch (e) {
                    reject(e);
                }
            });
            this.process();
        });
    }
    
    private async process(): Promise<void> {
        if (this.processing) return;
        this.processing = true;
        
        while (this.queue.length > 0) {
            const task = this.queue.shift()!;
            await task();
            await new Promise(r => 
                setTimeout(r, 1000 / this.ratePerSecond)
            );
        }
        
        this.processing = false;
    }
}

// 滑动窗口限流
class SlidingWindowRateLimiter {
    private requests: number[] = [];
    
    constructor(
        private windowMs: number,
        private maxRequests: number
    ) {}
    
    tryAcquire(): boolean {
        const now = Date.now();
        this.requests = this.requests.filter(t => now - t < this.windowMs);
        
        if (this.requests.length < this.maxRequests) {
            this.requests.push(now);
            return true;
        }
        return false;
    }
}
```

**Python**
```python
import asyncio
import time
from collections import deque
from dataclasses import dataclass

# ==================== 令牌桶 (Token Bucket) ====================
class TokenBucket:
    def __init__(self, capacity: int, refill_rate: float):
        self.capacity = capacity
        self.refill_rate = refill_rate  # tokens per second
        self.tokens = capacity
        self.last_refill = time.monotonic()
        self._lock = asyncio.Lock()
    
    def _refill(self):
        now = time.monotonic()
        elapsed = now - self.last_refill
        self.tokens = min(self.capacity, self.tokens + elapsed * self.refill_rate)
        self.last_refill = now
    
    async def acquire(self, tokens: int = 1):
        async with self._lock:
            while True:
                self._refill()
                if self.tokens >= tokens:
                    self.tokens -= tokens
                    return
                # 等待足够的 token
                wait_time = (tokens - self.tokens) / self.refill_rate
                await asyncio.sleep(wait_time)

# 使用令牌桶
bucket = TokenBucket(capacity=10, refill_rate=2.0)

async def rate_limited_fetch(url: str):
    await bucket.acquire()
    return await fetch(url)

# ==================== 漏桶 (Leaky Bucket) ====================
class LeakyBucket:
    def __init__(self, rate_per_second: float):
        self.rate = rate_per_second
        self.queue: asyncio.Queue = asyncio.Queue()
        self._processing = False
    
    async def add(self, coro):
        future = asyncio.Future()
        await self.queue.put((coro, future))
        asyncio.create_task(self._process())
        return await future
    
    async def _process(self):
        if self._processing:
            return
        self._processing = True
        
        while not self.queue.empty():
            coro, future = await self.queue.get()
            try:
                result = await coro
                future.set_result(result)
            except Exception as e:
                future.set_exception(e)
            await asyncio.sleep(1.0 / self.rate)
        
        self._processing = False

# ==================== 滑动窗口限流 ====================
class SlidingWindowLimiter:
    def __init__(self, window_seconds: float, max_requests: int):
        self.window = window_seconds
        self.max_requests = max_requests
        self.requests: deque = deque()
        self._lock = asyncio.Lock()
    
    async def acquire(self) -> bool:
        async with self._lock:
            now = time.monotonic()
            # 移除窗口外的请求
            while self.requests and now - self.requests[0] > self.window:
                self.requests.popleft()
            
            if len(self.requests) < self.max_requests:
                self.requests.append(now)
                return True
            return False
    
    async def wait_and_acquire(self):
        while not await self.acquire():
            await asyncio.sleep(0.1)
```

**Go**
```go
import (
    "context"
    "sync"
    "time"
    "golang.org/x/time/rate"
)

// ==================== 使用 golang.org/x/time/rate ====================
func rateLimitedFetch(urls []string) []string {
    // 每秒 10 个请求，突发最多 5 个
    limiter := rate.NewLimiter(rate.Limit(10), 5)
    
    results := make([]string, len(urls))
    var wg sync.WaitGroup
    
    for i, url := range urls {
        wg.Add(1)
        go func(idx int, u string) {
            defer wg.Done()
            
            // 等待获取令牌
            if err := limiter.Wait(context.Background()); err != nil {
                return
            }
            
            results[idx] = fetch(u)
        }(i, url)
    }
    
    wg.Wait()
    return results
}

// 预留令牌
func reserveExample(limiter *rate.Limiter) {
    r := limiter.Reserve()
    if !r.OK() {
        return  // 无法预留
    }
    
    delay := r.Delay()
    time.Sleep(delay)  // 等待
    
    // 执行操作
}

// 尝试获取（非阻塞）
func tryAcquire(limiter *rate.Limiter) bool {
    return limiter.Allow()
}

// ==================== 自定义令牌桶 ====================
type TokenBucket struct {
    capacity   int64
    tokens     int64
    refillRate float64  // per second
    lastRefill time.Time
    mu         sync.Mutex
}

func NewTokenBucket(capacity int64, refillRate float64) *TokenBucket {
    return &TokenBucket{
        capacity:   capacity,
        tokens:     capacity,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (b *TokenBucket) refill() {
    now := time.Now()
    elapsed := now.Sub(b.lastRefill).Seconds()
    b.tokens = min(b.capacity, b.tokens+int64(elapsed*b.refillRate))
    b.lastRefill = now
}

func (b *TokenBucket) Acquire(n int64) {
    b.mu.Lock()
    defer b.mu.Unlock()
    
    for {
        b.refill()
        if b.tokens >= n {
            b.tokens -= n
            return
        }
        
        // 计算等待时间
        needed := float64(n - b.tokens)
        waitTime := time.Duration(needed/b.refillRate) * time.Second
        b.mu.Unlock()
        time.Sleep(waitTime)
        b.mu.Lock()
    }
}

// ==================== 滑动窗口限流 ====================
type SlidingWindowLimiter struct {
    window      time.Duration
    maxRequests int
    requests    []time.Time
    mu          sync.Mutex
}

func NewSlidingWindowLimiter(window time.Duration, max int) *SlidingWindowLimiter {
    return &SlidingWindowLimiter{
        window:      window,
        maxRequests: max,
        requests:    make([]time.Time, 0),
    }
}

func (l *SlidingWindowLimiter) Allow() bool {
    l.mu.Lock()
    defer l.mu.Unlock()
    
    now := time.Now()
    cutoff := now.Add(-l.window)
    
    // 移除过期请求
    valid := l.requests[:0]
    for _, t := range l.requests {
        if t.After(cutoff) {
            valid = append(valid, t)
        }
    }
    l.requests = valid
    
    if len(l.requests) < l.maxRequests {
        l.requests = append(l.requests, now)
        return true
    }
    return false
}
```

**Rust**
```rust
use std::sync::Arc;
use tokio::sync::Mutex;
use std::time::{Duration, Instant};
use std::collections::VecDeque;

// ==================== 令牌桶 (Token Bucket) ====================
pub struct TokenBucket {
    capacity: f64,
    tokens: f64,
    refill_rate: f64,  // tokens per second
    last_refill: Instant,
}

impl TokenBucket {
    pub fn new(capacity: f64, refill_rate: f64) -> Self {
        Self {
            capacity,
            tokens: capacity,
            refill_rate,
            last_refill: Instant::now(),
        }
    }
    
    fn refill(&mut self) {
        let now = Instant::now();
        let elapsed = now.duration_since(self.last_refill).as_secs_f64();
        self.tokens = (self.tokens + elapsed * self.refill_rate).min(self.capacity);
        self.last_refill = now;
    }
    
    pub fn try_acquire(&mut self, tokens: f64) -> bool {
        self.refill();
        if self.tokens >= tokens {
            self.tokens -= tokens;
            true
        } else {
            false
        }
    }
    
    pub async fn acquire(&mut self, tokens: f64) {
        loop {
            self.refill();
            if self.tokens >= tokens {
                self.tokens -= tokens;
                return;
            }
            
            let wait_time = (tokens - self.tokens) / self.refill_rate;
            tokio::time::sleep(Duration::from_secs_f64(wait_time)).await;
        }
    }
}

// 线程安全版本
pub struct AsyncTokenBucket {
    inner: Arc<Mutex<TokenBucket>>,
}

impl AsyncTokenBucket {
    pub fn new(capacity: f64, refill_rate: f64) -> Self {
        Self {
            inner: Arc::new(Mutex::new(TokenBucket::new(capacity, refill_rate))),
        }
    }
    
    pub async fn acquire(&self, tokens: f64) {
        loop {
            {
                let mut bucket = self.inner.lock().await;
                if bucket.try_acquire(tokens) {
                    return;
                }
            }
            tokio::time::sleep(Duration::from_millis(100)).await;
        }
    }
}

// ==================== 使用 governor crate ====================
use governor::{Quota, RateLimiter};
use std::num::NonZeroU32;

async fn rate_limited_fetch(urls: Vec<String>) -> Vec<String> {
    // 每秒 10 个请求
    let limiter = RateLimiter::direct(Quota::per_second(NonZeroU32::new(10).unwrap()));
    
    let mut results = Vec::new();
    for url in urls {
        limiter.until_ready().await;
        results.push(fetch(&url).await);
    }
    results
}

// ==================== 滑动窗口限流 ====================
pub struct SlidingWindowLimiter {
    window: Duration,
    max_requests: usize,
    requests: VecDeque<Instant>,
}

impl SlidingWindowLimiter {
    pub fn new(window: Duration, max_requests: usize) -> Self {
        Self {
            window,
            max_requests,
            requests: VecDeque::new(),
        }
    }
    
    pub fn allow(&mut self) -> bool {
        let now = Instant::now();
        let cutoff = now - self.window;
        
        // 移除过期请求
        while let Some(front) = self.requests.front() {
            if *front < cutoff {
                self.requests.pop_front();
            } else {
                break;
            }
        }
        
        if self.requests.len() < self.max_requests {
            self.requests.push_back(now);
            true
        } else {
            false
        }
    }
}

// 异步版本
pub struct AsyncSlidingWindowLimiter {
    inner: Arc<Mutex<SlidingWindowLimiter>>,
}

impl AsyncSlidingWindowLimiter {
    pub async fn wait_and_acquire(&self) {
        loop {
            if self.inner.lock().await.allow() {
                return;
            }
            tokio::time::sleep(Duration::from_millis(100)).await;
        }
    }
}
```

### 并发控制模式对比

| 模式 | 用途 | TypeScript | Python | Go | Rust |
|------|------|------------|--------|-----|------|
| **All** | 等待所有任务完成 | `Promise.all` | `asyncio.gather` | `sync.WaitGroup` | `join_all` |
| **AllSettled** | 等待所有（含失败） | `Promise.allSettled` | `gather(return_exceptions=True)` | `errgroup` | `JoinSet` |
| **Race** | 返回最快结果 | `Promise.race` | `asyncio.wait(FIRST_COMPLETED)` | `select` | `select!` |
| **Any** | 返回首个成功 | `Promise.any` | 手动实现 | 手动实现 | 手动实现 |
| **Timeout** | 超时控制 | `AbortSignal.timeout` | `asyncio.timeout` | `context.WithTimeout` | `tokio::time::timeout` |
| **Cancel** | 取消操作 | `AbortController` | `Task.cancel()` | `context.WithCancel` | `CancellationToken` |
| **Semaphore** | 并发数控制 | 手动/p-limit | `asyncio.Semaphore` | `channel`/`semaphore` | `tokio::sync::Semaphore` |
| **Token Bucket** | 令牌桶限流 | 手动实现 | 手动实现 | `x/time/rate` | `governor` |
| **Leaky Bucket** | 漏桶限流 | 手动实现 | 手动实现 | 手动实现 | 手动实现 |
| **Sliding Window** | 滑动窗口限流 | 手动实现 | 手动实现 | 手动实现 | 手动实现 |

### 并发数据传递

#### 数据传递方式概览

| 方式 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 消息传递 | `postMessage` | `Queue` | `chan` | `mpsc/broadcast` |
| 共享内存 | `SharedArrayBuffer` | `multiprocessing.Value` | `sync.Mutex` | `Arc<Mutex<T>>` |
| 原子操作 | `Atomics` | `threading.Lock` | `atomic` | `std::sync::atomic` |
| 返回值 | `Promise` | `Future` | 返回值 | `JoinHandle` |
| 回调 | 回调函数 | 回调函数 | 函数参数 | 闭包 |

#### TypeScript 并发数据传递

```typescript
// ==================== Promise 链传递 ====================
async function pipeline() {
    const data = await fetchData();
    const processed = await processData(data);
    const result = await saveData(processed);
    return result;
}

// Promise 结果收集
const results = await Promise.all([
    fetchUsers(),
    fetchPosts(),
    fetchComments()
]);
const [users, posts, comments] = results;

// ==================== 回调传递 ====================
function fetchWithCallback(
    url: string,
    onSuccess: (data: any) => void,
    onError: (err: Error) => void
) {
    fetch(url)
        .then(res => res.json())
        .then(onSuccess)
        .catch(onError);
}

// ==================== EventEmitter 模式 ====================
import { EventEmitter } from 'events';

class DataProcessor extends EventEmitter {
    process(data: any) {
        this.emit('start', data);
        const result = transform(data);
        this.emit('complete', result);
        return result;
    }
}

const processor = new DataProcessor();
processor.on('complete', (result) => console.log(result));

// ==================== Worker 线程通信 ====================
// main.ts
const worker = new Worker('./worker.js');

// 发送数据到 Worker
worker.postMessage({ type: 'process', data: largeArray });

// 接收 Worker 结果
worker.onmessage = (event) => {
    const { type, result } = event.data;
    if (type === 'result') {
        console.log('Processed:', result);
    }
};

// worker.ts
self.onmessage = (event) => {
    const { type, data } = event.data;
    if (type === 'process') {
        const result = heavyComputation(data);
        self.postMessage({ type: 'result', result });
    }
};

// ==================== SharedArrayBuffer 共享内存 ====================
// 创建共享缓冲区
const sharedBuffer = new SharedArrayBuffer(1024);
const sharedArray = new Int32Array(sharedBuffer);

// 主线程
worker.postMessage({ buffer: sharedBuffer });
sharedArray[0] = 42;  // 直接修改共享内存

// Worker 线程
self.onmessage = (event) => {
    const view = new Int32Array(event.data.buffer);
    console.log(view[0]);  // 42 - 读取共享内存
    Atomics.add(view, 0, 1);  // 原子操作
};

// ==================== Atomics 原子操作 ====================
const sab = new SharedArrayBuffer(4);
const arr = new Int32Array(sab);

// 原子读写
Atomics.store(arr, 0, 123);
Atomics.load(arr, 0);  // 123

// 原子加减
Atomics.add(arr, 0, 10);      // 返回旧值，arr[0] += 10
Atomics.sub(arr, 0, 5);       // 返回旧值，arr[0] -= 5

// 比较并交换
Atomics.compareExchange(arr, 0, 128, 200);  // 如果是128则改为200

// 等待/通知 (线程同步)
// Worker 1: 等待
Atomics.wait(arr, 0, 0);  // 阻塞直到 arr[0] != 0

// Worker 2: 通知
Atomics.store(arr, 0, 1);
Atomics.notify(arr, 0, 1);  // 唤醒一个等待者

// ==================== MessageChannel ====================
const channel = new MessageChannel();
const port1 = channel.port1;
const port2 = channel.port2;

// 发送到另一个上下文
worker.postMessage({ port: port2 }, [port2]);

// 通过 port 通信
port1.onmessage = (e) => console.log(e.data);
port1.postMessage('Hello from main');

// ==================== BroadcastChannel ====================
// 跨标签页/Worker 广播
const broadcast = new BroadcastChannel('app-channel');

// 发送
broadcast.postMessage({ type: 'update', data: newData });

// 接收 (所有订阅者)
broadcast.onmessage = (event) => {
    console.log('Received:', event.data);
};

// ==================== 流式数据传递 ====================
async function* streamData(): AsyncGenerator<number> {
    for (let i = 0; i < 100; i++) {
        yield await fetchChunk(i);
    }
}

// 消费流
for await (const chunk of streamData()) {
    process(chunk);
}

// ReadableStream
const stream = new ReadableStream({
    async start(controller) {
        for (let i = 0; i < 10; i++) {
            controller.enqueue(i);
            await delay(100);
        }
        controller.close();
    }
});

const reader = stream.getReader();
while (true) {
    const { done, value } = await reader.read();
    if (done) break;
    console.log(value);
}
```

#### Python 并发数据传递

```python
import asyncio
import queue
import threading
import multiprocessing
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

# ==================== asyncio.Queue ====================
async def producer(q: asyncio.Queue):
    for i in range(10):
        await q.put(i)
        print(f"Produced: {i}")
        await asyncio.sleep(0.1)
    await q.put(None)  # 结束信号

async def consumer(q: asyncio.Queue):
    while True:
        item = await q.get()
        if item is None:
            break
        print(f"Consumed: {item}")
        q.task_done()

async def main():
    q = asyncio.Queue(maxsize=5)  # 有界队列
    await asyncio.gather(
        producer(q),
        consumer(q)
    )

# ==================== threading.Queue ====================
def threaded_producer(q: queue.Queue):
    for i in range(10):
        q.put(i)
    q.put(None)

def threaded_consumer(q: queue.Queue):
    while True:
        item = q.get()
        if item is None:
            break
        print(f"Got: {item}")
        q.task_done()

q = queue.Queue()
producer_thread = threading.Thread(target=threaded_producer, args=(q,))
consumer_thread = threading.Thread(target=threaded_consumer, args=(q,))

# 优先队列
pq = queue.PriorityQueue()
pq.put((1, "high priority"))
pq.put((10, "low priority"))

# ==================== multiprocessing 进程间通信 ====================
# Queue
def mp_producer(q: multiprocessing.Queue):
    for i in range(10):
        q.put(i)
    q.put(None)

def mp_consumer(q: multiprocessing.Queue):
    while True:
        item = q.get()
        if item is None:
            break
        print(f"Process got: {item}")

if __name__ == '__main__':
    q = multiprocessing.Queue()
    p1 = multiprocessing.Process(target=mp_producer, args=(q,))
    p2 = multiprocessing.Process(target=mp_consumer, args=(q,))
    p1.start()
    p2.start()

# Pipe - 双向通信
def pipe_sender(conn):
    conn.send("Hello")
    conn.send([1, 2, 3])
    conn.close()

def pipe_receiver(conn):
    print(conn.recv())  # "Hello"
    print(conn.recv())  # [1, 2, 3]

parent_conn, child_conn = multiprocessing.Pipe()

# ==================== 共享内存 ====================
# Value - 单个值
counter = multiprocessing.Value('i', 0)  # 'i' = int

def increment(counter):
    for _ in range(1000):
        with counter.get_lock():
            counter.value += 1

# Array - 数组
shared_array = multiprocessing.Array('d', [0.0] * 10)  # 'd' = double

# shared_memory (Python 3.8+)
from multiprocessing import shared_memory

# 创建共享内存
shm = shared_memory.SharedMemory(create=True, size=1024)
buffer = shm.buf

# 写入
buffer[0:5] = b'Hello'

# 另一个进程访问
shm2 = shared_memory.SharedMemory(name=shm.name)
print(bytes(shm2.buf[0:5]))  # b'Hello'

# 清理
shm.close()
shm.unlink()

# ==================== Manager - 共享复杂对象 ====================
manager = multiprocessing.Manager()
shared_dict = manager.dict()
shared_list = manager.list()

def worker(d, l, key, value):
    d[key] = value
    l.append(value)

# ==================== Future 结果获取 ====================
with ThreadPoolExecutor(max_workers=4) as executor:
    # submit 返回 Future
    future = executor.submit(heavy_task, arg1, arg2)
    
    # 获取结果 (阻塞)
    result = future.result(timeout=10)
    
    # 检查状态
    future.done()       # 是否完成
    future.cancelled()  # 是否取消
    future.exception()  # 获取异常

# map 批量获取结果
with ProcessPoolExecutor() as executor:
    results = list(executor.map(process, items))

# as_completed 按完成顺序获取
from concurrent.futures import as_completed

futures = [executor.submit(task, i) for i in range(10)]
for future in as_completed(futures):
    result = future.result()
    print(result)

# ==================== 回调函数 ====================
def on_complete(future):
    print(f"Result: {future.result()}")

future = executor.submit(task)
future.add_done_callback(on_complete)

# ==================== asyncio 事件 ====================
event = asyncio.Event()

async def waiter():
    print("Waiting...")
    await event.wait()
    print("Event fired!")

async def setter():
    await asyncio.sleep(1)
    event.set()

# Condition
condition = asyncio.Condition()

async def consumer():
    async with condition:
        await condition.wait()
        # 处理数据

async def producer():
    async with condition:
        # 准备数据
        condition.notify_all()
```

#### Go 并发数据传递

```go
// ==================== Channel 基础 ====================
// 无缓冲 channel (同步)
ch := make(chan int)

// 有缓冲 channel (异步)
buffered := make(chan int, 10)

// 发送和接收
go func() {
    ch <- 42  // 发送
}()
value := <-ch  // 接收

// 关闭 channel
close(ch)

// 检查是否关闭
value, ok := <-ch
if !ok {
    fmt.Println("Channel closed")
}

// ==================== Channel 方向 ====================
// 只发送
func producer(out chan<- int) {
    for i := 0; i < 10; i++ {
        out <- i
    }
    close(out)
}

// 只接收
func consumer(in <-chan int) {
    for value := range in {
        fmt.Println(value)
    }
}

// ==================== 生产者-消费者模式 ====================
func main() {
    ch := make(chan int, 5)
    
    // 生产者
    go func() {
        for i := 0; i < 10; i++ {
            ch <- i
            fmt.Printf("Produced: %d\n", i)
        }
        close(ch)
    }()
    
    // 消费者
    for value := range ch {
        fmt.Printf("Consumed: %d\n", value)
    }
}

// ==================== select 多路复用 ====================
func main() {
    ch1 := make(chan string)
    ch2 := make(chan string)
    
    go func() {
        time.Sleep(1 * time.Second)
        ch1 <- "one"
    }()
    
    go func() {
        time.Sleep(2 * time.Second)
        ch2 <- "two"
    }()
    
    for i := 0; i < 2; i++ {
        select {
        case msg1 := <-ch1:
            fmt.Println("Received", msg1)
        case msg2 := <-ch2:
            fmt.Println("Received", msg2)
        case <-time.After(3 * time.Second):
            fmt.Println("Timeout")
        }
    }
}

// 非阻塞 select
select {
case msg := <-ch:
    fmt.Println(msg)
default:
    fmt.Println("No message")
}

// ==================== Fan-out / Fan-in ====================
// Fan-out: 一个 channel 分发给多个 worker
func fanOut(input <-chan int, workers int) []<-chan int {
    outputs := make([]<-chan int, workers)
    for i := 0; i < workers; i++ {
        outputs[i] = worker(input)
    }
    return outputs
}

func worker(input <-chan int) <-chan int {
    output := make(chan int)
    go func() {
        defer close(output)
        for n := range input {
            output <- process(n)
        }
    }()
    return output
}

// Fan-in: 多个 channel 合并为一个
func fanIn(inputs ...<-chan int) <-chan int {
    output := make(chan int)
    var wg sync.WaitGroup
    
    for _, ch := range inputs {
        wg.Add(1)
        go func(c <-chan int) {
            defer wg.Done()
            for n := range c {
                output <- n
            }
        }(ch)
    }
    
    go func() {
        wg.Wait()
        close(output)
    }()
    
    return output
}

// ==================== Pipeline 模式 ====================
func gen(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}

func double(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * 2
        }
        close(out)
    }()
    return out
}

// 使用
func main() {
    // gen -> square -> double
    for n := range double(square(gen(1, 2, 3, 4))) {
        fmt.Println(n)  // 2, 8, 18, 32
    }
}

// ==================== 共享内存 (Mutex) ====================
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
type SafeMap struct {
    mu sync.RWMutex
    m  map[string]int
}

func (sm *SafeMap) Get(key string) int {
    sm.mu.RLock()  // 读锁
    defer sm.mu.RUnlock()
    return sm.m[key]
}

func (sm *SafeMap) Set(key string, value int) {
    sm.mu.Lock()   // 写锁
    defer sm.mu.Unlock()
    sm.m[key] = value
}

// ==================== sync.Map ====================
var m sync.Map

// 存储
m.Store("key", "value")

// 读取
value, ok := m.Load("key")

// 读取或存储
actual, loaded := m.LoadOrStore("key", "default")

// 删除
m.Delete("key")

// 遍历
m.Range(func(key, value any) bool {
    fmt.Printf("%v: %v\n", key, value)
    return true  // 继续遍历
})

// ==================== atomic 原子操作 ====================
var counter int64

// 原子加
atomic.AddInt64(&counter, 1)

// 原子读写
atomic.StoreInt64(&counter, 100)
value := atomic.LoadInt64(&counter)

// 比较并交换
swapped := atomic.CompareAndSwapInt64(&counter, 100, 200)

// atomic.Value 存储任意类型
var config atomic.Value
config.Store(map[string]string{"key": "value"})
cfg := config.Load().(map[string]string)

// ==================== sync.Pool 对象池 ====================
var bufferPool = sync.Pool{
    New: func() any {
        return make([]byte, 1024)
    },
}

// 获取
buf := bufferPool.Get().([]byte)

// 使用后归还
bufferPool.Put(buf)

// ==================== Context 传递数据 ====================
type key string

func main() {
    ctx := context.Background()
    ctx = context.WithValue(ctx, key("userID"), "12345")
    
    processRequest(ctx)
}

func processRequest(ctx context.Context) {
    userID := ctx.Value(key("userID")).(string)
    fmt.Println("User:", userID)
}
```

#### Rust 并发数据传递

```rust
use std::sync::{Arc, Mutex, RwLock, mpsc, atomic::{AtomicUsize, Ordering}};
use std::thread;
use tokio::sync::{mpsc as tokio_mpsc, broadcast, oneshot, watch};

// ==================== std::sync::mpsc (多生产者单消费者) ====================
fn mpsc_example() {
    let (tx, rx) = mpsc::channel();
    
    // 多个生产者
    for i in 0..3 {
        let tx_clone = tx.clone();
        thread::spawn(move || {
            tx_clone.send(format!("Message from {}", i)).unwrap();
        });
    }
    drop(tx);  // 关闭原始发送端
    
    // 单个消费者
    for received in rx {
        println!("Got: {}", received);
    }
}

// 同步 channel (有界)
fn sync_channel_example() {
    let (tx, rx) = mpsc::sync_channel(2);  // 缓冲区大小 2
    
    thread::spawn(move || {
        tx.send(1).unwrap();
        tx.send(2).unwrap();
        tx.send(3).unwrap();  // 阻塞，直到有空间
    });
    
    for val in rx {
        println!("{}", val);
    }
}

// ==================== tokio::sync::mpsc (异步) ====================
async fn tokio_mpsc_example() {
    let (tx, mut rx) = tokio_mpsc::channel(100);
    
    // 生产者
    tokio::spawn(async move {
        for i in 0..10 {
            tx.send(i).await.unwrap();
        }
    });
    
    // 消费者
    while let Some(value) = rx.recv().await {
        println!("Received: {}", value);
    }
}

// ==================== broadcast (多生产者多消费者) ====================
async fn broadcast_example() {
    let (tx, mut rx1) = broadcast::channel(16);
    let mut rx2 = tx.subscribe();
    
    tokio::spawn(async move {
        while let Ok(value) = rx1.recv().await {
            println!("Receiver 1: {}", value);
        }
    });
    
    tokio::spawn(async move {
        while let Ok(value) = rx2.recv().await {
            println!("Receiver 2: {}", value);
        }
    });
    
    tx.send("Hello").unwrap();
    tx.send("World").unwrap();
}

// ==================== oneshot (单次传递) ====================
async fn oneshot_example() {
    let (tx, rx) = oneshot::channel();
    
    tokio::spawn(async move {
        // 执行计算
        let result = expensive_computation().await;
        tx.send(result).unwrap();
    });
    
    // 等待结果
    let result = rx.await.unwrap();
    println!("Result: {}", result);
}

// ==================== watch (单生产者多消费者，最新值) ====================
async fn watch_example() {
    let (tx, mut rx) = watch::channel("initial");
    
    tokio::spawn(async move {
        loop {
            // 等待值改变
            rx.changed().await.unwrap();
            println!("Value changed to: {}", *rx.borrow());
        }
    });
    
    tx.send("updated").unwrap();
    tx.send("final").unwrap();
}

// ==================== Arc<Mutex<T>> 共享可变状态 ====================
fn shared_state_example() {
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
}

// 异步版本
async fn async_shared_state() {
    let counter = Arc::new(tokio::sync::Mutex::new(0));
    
    let mut handles = vec![];
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        handles.push(tokio::spawn(async move {
            let mut num = counter.lock().await;
            *num += 1;
        }));
    }
    
    for handle in handles {
        handle.await.unwrap();
    }
}

// ==================== RwLock 读写锁 ====================
fn rwlock_example() {
    let data = Arc::new(RwLock::new(vec![1, 2, 3]));
    
    // 多个读者
    let data_clone = Arc::clone(&data);
    thread::spawn(move || {
        let reader = data_clone.read().unwrap();
        println!("Read: {:?}", *reader);
    });
    
    // 写者
    {
        let mut writer = data.write().unwrap();
        writer.push(4);
    }
}

// ==================== Atomic 原子类型 ====================
fn atomic_example() {
    let counter = Arc::new(AtomicUsize::new(0));
    let mut handles = vec![];
    
    for _ in 0..10 {
        let counter = Arc::clone(&counter);
        handles.push(thread::spawn(move || {
            for _ in 0..1000 {
                counter.fetch_add(1, Ordering::SeqCst);
            }
        }));
    }
    
    for handle in handles {
        handle.join().unwrap();
    }
    
    println!("Counter: {}", counter.load(Ordering::SeqCst));
}

// ==================== JoinHandle 获取返回值 ====================
fn join_handle_example() {
    let handle = thread::spawn(|| {
        // 计算
        42
    });
    
    let result = handle.join().unwrap();
    println!("Thread returned: {}", result);
}

// tokio 版本
async fn tokio_join_example() {
    let handle = tokio::spawn(async {
        expensive_computation().await
    });
    
    let result = handle.await.unwrap();
}

// ==================== crossbeam channel (高性能) ====================
use crossbeam_channel::{bounded, unbounded, select};

fn crossbeam_example() {
    let (s, r) = bounded(10);  // 有界
    // let (s, r) = unbounded(); // 无界
    
    thread::spawn(move || {
        s.send("Hello").unwrap();
    });
    
    println!("{}", r.recv().unwrap());
}

// select 宏
fn crossbeam_select() {
    let (s1, r1) = unbounded();
    let (s2, r2) = unbounded();
    
    thread::spawn(move || s1.send(1).unwrap());
    thread::spawn(move || s2.send(2).unwrap());
    
    select! {
        recv(r1) -> msg => println!("r1: {:?}", msg),
        recv(r2) -> msg => println!("r2: {:?}", msg),
    }
}

// ==================== flume (高性能 mpmc) ====================
use flume;

async fn flume_example() {
    let (tx, rx) = flume::bounded(100);
    
    // 同步发送
    tx.send(1).unwrap();
    
    // 异步发送
    tx.send_async(2).await.unwrap();
    
    // 同步接收
    let val = rx.recv().unwrap();
    
    // 异步接收
    let val = rx.recv_async().await.unwrap();
}

// ==================== 并发数据结构 ====================
use dashmap::DashMap;

fn dashmap_example() {
    let map = DashMap::new();
    
    // 并发插入
    map.insert("key", "value");
    
    // 并发读取
    if let Some(val) = map.get("key") {
        println!("{}", *val);
    }
    
    // 并发修改
    map.alter("key", |_, v| format!("{}_modified", v));
}
```

### 并发数据传递对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 模式            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 基础 Channel    │ postMessage          │ queue.Queue          │ chan                 │ mpsc::channel        │
│ 异步 Channel    │ MessageChannel       │ asyncio.Queue        │ chan                 │ tokio::sync::mpsc    │
│ 广播            │ BroadcastChannel     │ 手动实现             │ 手动实现             │ broadcast::channel   │
│ 单次传递        │ Promise              │ asyncio.Future       │ chan (cap=1)         │ oneshot::channel     │
│ 共享状态        │ SharedArrayBuffer    │ Manager/Value        │ sync.Mutex           │ Arc<Mutex<T>>        │
│ 读写锁          │ ❌                   │ threading.RWLock     │ sync.RWMutex         │ RwLock               │
│ 原子操作        │ Atomics              │ ❌                   │ atomic               │ std::sync::atomic    │
│ 对象池          │ ❌                   │ ❌                   │ sync.Pool            │ ❌                   │
│ 并发 Map        │ ❌                   │ Manager.dict()       │ sync.Map             │ DashMap              │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 选择数据传递方式

| 场景 | 推荐方式 | 原因 |
|------|----------|------|
| **任务结果返回** | Promise/Future/JoinHandle | 简单直接 |
| **生产者-消费者** | Channel/Queue | 解耦，背压控制 |
| **配置热更新** | watch/atomic.Value | 最新值广播 |
| **高频读低频写** | RwLock | 读不阻塞 |
| **计数器/标志** | Atomic | 无锁高性能 |
| **复杂共享状态** | Mutex + Arc | 灵活但需注意死锁 |
| **跨进程** | 共享内存/IPC | 高性能大数据 |

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

## 📦 模块导入导出

### 模块系统概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 模块单位 | 文件 | 文件/包 | 包(目录) | 文件/crate |
| 导入关键字 | `import` | `import` | `import` | `use` |
| 导出关键字 | `export` | 无(默认公开) | 大写首字母 | `pub` |
| 默认导出 | `export default` | 无 | 无 | 无 |
| 重导出 | `export { } from` | `__all__` | 无(需包装) | `pub use` |
| 循环依赖 | 支持(需注意) | 支持(需注意) | 禁止 | 禁止 |
| 动态导入 | `import()` | `importlib` | 插件机制 | 不支持 |

### TypeScript 模块导入导出

```typescript
// ==================== 命名导出 (Named Export) ====================
// math.ts
export const PI = 3.14159;
export const E = 2.71828;

export function add(a: number, b: number): number {
    return a + b;
}

export function multiply(a: number, b: number): number {
    return a * b;
}

export interface Point {
    x: number;
    y: number;
}

export class Calculator {
    add(a: number, b: number): number {
        return a + b;
    }
}

// ==================== 命名导入 ====================
// 导入特定项
import { add, multiply, PI } from './math';
import { Point, Calculator } from './math';

// 重命名导入
import { add as addNumbers, multiply as mult } from './math';

// 导入全部为命名空间
import * as MathUtils from './math';
MathUtils.add(1, 2);
MathUtils.PI;

// ==================== 默认导出 (Default Export) ====================
// logger.ts
export default class Logger {
    log(message: string): void {
        console.log(`[LOG] ${message}`);
    }
}

// 或函数
export default function createLogger(prefix: string) {
    return (msg: string) => console.log(`[${prefix}] ${msg}`);
}

// ==================== 默认导入 ====================
import Logger from './logger';
import createLogger from './logger';

// 可以用任意名称
import MyLogger from './logger';

// 同时导入默认和命名
import Logger, { LogLevel, formatMessage } from './logger';

// ==================== 重导出 (Re-export) ====================
// index.ts - 桶文件 (Barrel file)
export { add, multiply } from './math';
export { default as Logger } from './logger';
export * from './utils';  // 导出所有命名导出
export * as MathUtils from './math';  // 作为命名空间重导出

// 重命名后重导出
export { add as addition } from './math';

// ==================== 类型导入导出 ====================
// types.ts
export type UserId = string;
export interface User {
    id: UserId;
    name: string;
}

// 仅类型导入 (不会编译到 JS)
import type { User, UserId } from './types';

// 内联类型导入
import { type User, createUser } from './user';

// 类型重导出
export type { User, UserId } from './types';

// ==================== 动态导入 ====================
// 懒加载模块
async function loadModule() {
    const { add } = await import('./math');
    return add(1, 2);
}

// 条件导入
async function loadLocale(lang: string) {
    const locale = await import(`./locales/${lang}.json`);
    return locale.default;
}

// 与 React.lazy 结合
const LazyComponent = React.lazy(() => import('./HeavyComponent'));

// ==================== 模块解析 ====================
// 相对路径
import { foo } from './utils';        // 同目录
import { bar } from '../helpers';     // 上级目录
import { baz } from './sub/module';   // 子目录

// 绝对路径 (通过 tsconfig paths)
import { api } from '@/services/api';
import { Button } from '@components/Button';

// Node 模块
import express from 'express';
import { readFile } from 'fs/promises';

// ==================== CommonJS 互操作 ====================
// 导入 CommonJS 模块
import lodash from 'lodash';  // default import
import * as _ from 'lodash';  // namespace import

// 导出为 CommonJS (当 module: commonjs)
module.exports = { add, multiply };
exports.PI = 3.14;

// ==================== tsconfig.json 配置 ====================
{
    "compilerOptions": {
        "module": "ESNext",           // 模块系统
        "moduleResolution": "bundler", // 解析策略
        "baseUrl": "./src",           // 基础路径
        "paths": {                    // 路径别名
            "@/*": ["./*"],
            "@components/*": ["components/*"]
        },
        "esModuleInterop": true,      // CJS/ESM 互操作
        "allowSyntheticDefaultImports": true
    }
}
```

### Python 模块导入导出

```python
# ==================== 基本导入 ====================
# 导入整个模块
import math
import os.path

# 使用
math.sqrt(16)
os.path.join('a', 'b')

# 导入并重命名
import numpy as np
import pandas as pd

# ==================== 从模块导入 ====================
# 导入特定项
from math import sqrt, pi, ceil
from os.path import join, exists

# 导入并重命名
from math import sqrt as square_root
from collections import defaultdict as dd

# 导入所有 (不推荐)
from math import *

# ==================== 包结构 ====================
"""
mypackage/
├── __init__.py          # 包初始化文件
├── module1.py
├── module2.py
└── subpackage/
    ├── __init__.py
    └── module3.py
"""

# ==================== __init__.py 控制导出 ====================
# mypackage/__init__.py
from .module1 import func1, Class1
from .module2 import func2
from .subpackage import module3

# 定义公开 API
__all__ = ['func1', 'Class1', 'func2']

# 版本信息
__version__ = '1.0.0'

# 使用
from mypackage import func1, Class1

# ==================== 相对导入 ====================
# 在包内部使用
# mypackage/module2.py
from . import module1           # 同级模块
from .module1 import func1      # 同级模块的函数
from .. import other_package    # 上级包
from ..sibling import helper    # 兄弟包

# ==================== 绝对导入 ====================
# 推荐方式
from mypackage.module1 import func1
from mypackage.subpackage.module3 import something

# ==================== 条件导入 ====================
import sys

if sys.version_info >= (3, 11):
    from tomllib import load
else:
    from tomli import load

# 可选依赖
try:
    import numpy as np
    HAS_NUMPY = True
except ImportError:
    HAS_NUMPY = False

# 类型检查时导入 (避免循环导入)
from typing import TYPE_CHECKING

if TYPE_CHECKING:
    from .heavy_module import HeavyClass

def process(obj: 'HeavyClass') -> None:
    pass

# ==================== 动态导入 ====================
import importlib

# 按名称导入模块
module = importlib.import_module('math')
module.sqrt(16)

# 动态导入子模块
def load_plugin(name: str):
    return importlib.import_module(f'plugins.{name}')

# 重新加载模块
importlib.reload(module)

# ==================== 导入钩子 ====================
import sys
from importlib.abc import MetaPathFinder, Loader
from importlib.machinery import ModuleSpec

class CustomFinder(MetaPathFinder):
    def find_spec(self, fullname, path, target=None):
        if fullname.startswith('custom.'):
            return ModuleSpec(fullname, CustomLoader())
        return None

class CustomLoader(Loader):
    def create_module(self, spec):
        return None
    
    def exec_module(self, module):
        module.custom_attr = 'value'

sys.meta_path.insert(0, CustomFinder())

# ==================== 私有约定 ====================
# module.py
public_var = 'accessible'
_private_var = 'internal use'  # 单下划线: 约定私有
__very_private = 'name mangled'  # 双下划线: 名称改写

def public_function():
    pass

def _private_function():  # 不会被 from module import * 导入
    pass

# ==================== 命名空间包 (PEP 420) ====================
"""
无需 __init__.py 的包:
namespace_pkg/
    └── sub1/
        └── module.py

other_location/namespace_pkg/
    └── sub2/
        └── module.py

两者合并为同一命名空间
"""

# ==================== __all__ 详解 ====================
# utils.py
__all__ = ['public_func', 'PublicClass']

def public_func():
    """会被 from utils import * 导入"""
    pass

def _helper():
    """不在 __all__ 中，* 导入不包含"""
    pass

class PublicClass:
    pass

class _InternalClass:
    pass
```

### Go 模块导入导出

```go
// ==================== 包声明 ====================
// 每个 Go 文件必须声明包名
// main.go
package main  // 可执行程序的入口包

// utils/helper.go
package utils  // 库包，目录名通常与包名一致

// ==================== 导入语法 ====================
package main

import (
    // 标准库
    "fmt"
    "os"
    "net/http"
    
    // 第三方包
    "github.com/gin-gonic/gin"
    "github.com/spf13/cobra"
    
    // 本地包 (模块路径 + 相对路径)
    "myproject/internal/utils"
    "myproject/pkg/models"
)

// 单个导入
import "fmt"

// ==================== 导入别名 ====================
import (
    "fmt"
    
    // 别名
    myfmt "myproject/pkg/fmt"
    
    // 点导入 (不推荐，污染命名空间)
    . "math"
    
    // 空白导入 (仅执行 init，不使用)
    _ "github.com/lib/pq"
)

// 使用
func main() {
    fmt.Println("standard")
    myfmt.Custom()
    
    // 点导入后可直接使用
    result := Sqrt(16)  // 而非 math.Sqrt
}

// ==================== 导出规则 (大小写) ====================
// models/user.go
package models

// 公开 (大写开头) - 其他包可访问
type User struct {
    ID   int    // 公开字段
    Name string // 公开字段
    age  int    // 私有字段 (小写)
}

// 公开函数
func NewUser(name string) *User {
    return &User{Name: name}
}

// 公开方法
func (u *User) GetName() string {
    return u.Name
}

// 私有函数 (小写开头) - 仅包内可访问
func validateUser(u *User) bool {
    return u.Name != ""
}

// 公开常量和变量
const MaxUsers = 100
var DefaultUser = &User{Name: "guest"}

// 私有常量和变量
const maxRetries = 3
var internalCache = make(map[string]string)

// ==================== internal 包 ====================
/*
项目结构:
myproject/
├── cmd/
│   └── app/
│       └── main.go
├── internal/          # 内部包，外部项目不可导入
│   ├── auth/
│   │   └── auth.go
│   └── database/
│       └── db.go
├── pkg/               # 公开包，可被外部导入
│   └── models/
│       └── user.go
└── go.mod
*/

// internal 包只能被同一模块内的代码导入
// 外部项目导入 internal 会报错

// ==================== init 函数 ====================
package database

import "database/sql"

var db *sql.DB

// init 在包被导入时自动执行
// 一个包可以有多个 init，按文件名顺序执行
func init() {
    var err error
    db, err = sql.Open("postgres", "...")
    if err != nil {
        panic(err)
    }
}

// 另一个文件的 init
func init() {
    // 执行迁移等
}

// ==================== 包组织最佳实践 ====================
/*
myproject/
├── cmd/                    # 可执行文件入口
│   ├── server/
│   │   └── main.go
│   └── cli/
│       └── main.go
├── internal/               # 私有代码
│   ├── handler/
│   ├── service/
│   └── repository/
├── pkg/                    # 可导出的库代码
│   ├── api/
│   └── models/
├── api/                    # API 定义 (OpenAPI, protobuf)
├── configs/                # 配置文件
├── scripts/                # 脚本
├── go.mod
└── go.sum
*/

// ==================== go.mod 模块文件 ====================
/*
module github.com/username/myproject

go 1.21

require (
    github.com/gin-gonic/gin v1.9.1
    github.com/spf13/cobra v1.7.0
)

require (
    // 间接依赖
    github.com/inconshreveable/mousetrap v1.1.0 // indirect
)

replace (
    // 本地开发替换
    github.com/some/pkg => ../local/pkg
)
*/

// ==================== 子包导入 ====================
// 同一模块内的包导入
// myproject/internal/service/user.go
package service

import (
    "myproject/internal/repository"
    "myproject/pkg/models"
)

type UserService struct {
    repo *repository.UserRepository
}

func (s *UserService) GetUser(id int) (*models.User, error) {
    return s.repo.FindByID(id)
}

// ==================== 接口与实现分离 ====================
// pkg/storage/interface.go
package storage

type Storage interface {
    Get(key string) ([]byte, error)
    Set(key string, value []byte) error
    Delete(key string) error
}

// internal/storage/redis/redis.go
package redis

import "myproject/pkg/storage"

type RedisStorage struct {
    // ...
}

// 确保实现接口
var _ storage.Storage = (*RedisStorage)(nil)

func (r *RedisStorage) Get(key string) ([]byte, error) {
    // 实现
}
```

### Rust 模块导入导出

```rust
// ==================== 模块声明 ====================
// src/lib.rs 或 src/main.rs

// 声明模块 (从文件或目录加载)
mod utils;      // 加载 src/utils.rs 或 src/utils/mod.rs
mod models;     // 加载 src/models.rs 或 src/models/mod.rs

// 内联模块
mod inline_module {
    pub fn helper() {}
}

// ==================== 文件结构 ====================
/*
my_crate/
├── Cargo.toml
└── src/
    ├── lib.rs          # 库 crate 根
    ├── main.rs         # 二进制 crate 根
    ├── utils.rs        # utils 模块
    └── models/         # models 模块 (目录形式)
        ├── mod.rs      # 模块入口
        ├── user.rs     # 子模块
        └── post.rs     # 子模块
*/

// ==================== 可见性 (pub) ====================
// src/lib.rs
mod internal {
    // 私有 - 仅当前模块可见
    fn private_fn() {}
    
    // 公开 - 外部可见
    pub fn public_fn() {}
    
    // pub(crate) - 仅当前 crate 可见
    pub(crate) fn crate_fn() {}
    
    // pub(super) - 仅父模块可见
    pub(super) fn parent_fn() {}
    
    // pub(in path) - 指定路径可见
    pub(in crate::internal) fn specific_fn() {}
}

// 公开结构体
pub struct User {
    pub name: String,      // 公开字段
    pub(crate) email: String,  // crate 内可见
    password: String,      // 私有字段
}

impl User {
    // 公开关联函数
    pub fn new(name: String) -> Self {
        Self {
            name,
            email: String::new(),
            password: String::new(),
        }
    }
    
    // 私有方法
    fn validate(&self) -> bool {
        !self.name.is_empty()
    }
}

// ==================== use 导入 ====================
// 导入标准库
use std::collections::HashMap;
use std::io::{self, Read, Write};  // self 导入 io 本身

// 导入外部 crate
use serde::{Serialize, Deserialize};
use tokio::sync::mpsc;

// 导入当前 crate
use crate::models::User;
use crate::utils::helper;

// 导入父模块
use super::parent_function;

// 导入同级模块
use self::sibling_module::something;

// ==================== 重命名与通配符 ====================
// 重命名
use std::collections::HashMap as Map;
use std::io::Result as IoResult;

// 通配符 (不推荐，除非 prelude)
use std::collections::*;

// 嵌套导入
use std::{
    collections::{HashMap, HashSet},
    io::{self, Read, Write},
    sync::{Arc, Mutex},
};

// ==================== 重导出 (pub use) ====================
// src/lib.rs
mod internal_impl;

// 重导出为公开 API
pub use internal_impl::ImportantStruct;
pub use internal_impl::important_function;

// 重命名后重导出
pub use internal_impl::OldName as NewName;

// 重导出外部 crate
pub use serde_json::Value as JsonValue;

// ==================== prelude 模式 ====================
// src/prelude.rs
pub use crate::models::{User, Post, Comment};
pub use crate::traits::{Validate, Serialize};
pub use crate::error::{Error, Result};

// 使用方在一行导入常用项
use my_crate::prelude::*;

// ==================== 模块文件组织 ====================
// src/models/mod.rs
mod user;
mod post;
mod comment;

// 公开子模块内容
pub use user::User;
pub use post::Post;
pub use comment::Comment;

// 或公开整个子模块
pub mod user;
pub mod post;

// src/models/user.rs
#[derive(Debug, Clone)]
pub struct User {
    pub id: u64,
    pub name: String,
}

impl User {
    pub fn new(name: String) -> Self {
        Self { id: 0, name }
    }
}

// ==================== 条件编译与导入 ====================
// 平台特定
#[cfg(target_os = "windows")]
mod windows_impl;

#[cfg(target_os = "linux")]
mod linux_impl;

#[cfg(target_os = "windows")]
pub use windows_impl::*;

#[cfg(target_os = "linux")]
pub use linux_impl::*;

// feature 特定
#[cfg(feature = "async")]
pub mod async_api;

#[cfg(feature = "async")]
pub use async_api::*;

// ==================== 外部 crate 导入 ====================
// Cargo.toml
/*
[dependencies]
serde = { version = "1.0", features = ["derive"] }
tokio = { version = "1.0", features = ["full"] }
anyhow = "1.0"

[dev-dependencies]
mockall = "0.11"

[build-dependencies]
cc = "1.0"
*/

// 使用
use serde::{Serialize, Deserialize};
use tokio::runtime::Runtime;
use anyhow::{Result, Context};

// ==================== 路径类型 ====================
// 绝对路径 (从 crate 根开始)
use crate::models::User;
use crate::utils::helper;

// 相对路径
use self::submodule::Item;      // 当前模块的子模块
use super::sibling::Other;      // 父模块的兄弟模块
use super::super::ancestor::X;  // 祖父模块

// ==================== workspace 模块 ====================
// workspace Cargo.toml
/*
[workspace]
members = [
    "crates/core",
    "crates/api",
    "crates/cli",
]
*/

// crates/api/Cargo.toml
/*
[dependencies]
core = { path = "../core" }
*/

// crates/api/src/lib.rs
use core::models::User;  // 使用 workspace 内其他 crate

// ==================== 测试模块 ====================
// src/lib.rs
pub fn add(a: i32, b: i32) -> i32 {
    a + b
}

#[cfg(test)]
mod tests {
    use super::*;  // 导入父模块所有公开项
    
    #[test]
    fn test_add() {
        assert_eq!(add(2, 3), 5);
    }
}

// 集成测试 (tests/integration_test.rs)
use my_crate::add;  // 只能访问公开 API

#[test]
fn integration_test() {
    assert_eq!(add(1, 2), 3);
}
```

### 模块导入导出对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 操作            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 导入模块        │ import * as M        │ import module        │ import "pkg"         │ use crate::module    │
│ 导入特定项      │ import { a, b }      │ from m import a, b   │ (自动全部)           │ use mod::{a, b}      │
│ 重命名导入      │ import { a as x }    │ from m import a as x │ import pkg "path"    │ use mod::a as x      │
│ 默认导出        │ export default       │ ❌                   │ ❌                   │ ❌                   │
│ 命名导出        │ export { a, b }      │ __all__ = [...]      │ 大写首字母           │ pub                  │
│ 重导出          │ export { } from      │ from m import *      │ 包装函数             │ pub use              │
│ 私有项          │ 不导出               │ _前缀 (约定)         │ 小写首字母           │ 不加 pub             │
│ 动态导入        │ import()             │ importlib            │ 插件机制             │ ❌                   │
│ 循环依赖        │ ⚠️ 运行时处理        │ ⚠️ 需注意顺序        │ ❌ 编译错误          │ ❌ 编译错误          │
│ 路径别名        │ tsconfig paths       │ setuptools           │ go.mod replace       │ Cargo.toml           │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 常见模块模式

| 模式 | 描述 | 推荐语言 |
|------|------|----------|
| **Barrel 文件** | index 文件重导出子模块 | TypeScript |
| **Prelude** | 常用项集中导出 | Rust |
| **\_\_all\_\_** | 控制 * 导入范围 | Python |
| **internal 包** | 内部实现不可外部导入 | Go |
| **pub(crate)** | crate 内可见，外部不可见 | Rust |
| **条件导入** | 按环境/特性选择模块 | 全部支持 |
| **懒加载** | 按需动态导入模块 | TypeScript/Python |

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

## ⚡ 性能测试

### 性能测试概览

| 语言 | 基准测试工具 | 性能分析 | 内置支持 |
|------|--------------|----------|----------|
| TypeScript | Benchmark.js, vitest | Chrome DevTools | ❌ |
| Python | timeit, pytest-benchmark | cProfile, py-spy | ✅ timeit |
| Go | testing.B | pprof | ✅ 内置 |
| Rust | criterion, cargo bench | perf, flamegraph | ✅ 内置 |

### TypeScript 性能测试

```typescript
// ==================== 简单计时 ====================

console.time('operation');
doSomething();
console.timeEnd('operation');  // operation: 123.456ms

// performance API
const start = performance.now();
doSomething();
const end = performance.now();
console.log(`耗时: ${end - start}ms`);

// ==================== Benchmark.js ====================
import Benchmark from 'benchmark';

const suite = new Benchmark.Suite();

suite
  .add('Array.push', function() {
    const arr = [];
    for (let i = 0; i < 1000; i++) {
      arr.push(i);
    }
  })
  .add('Array spread', function() {
    let arr = [];
    for (let i = 0; i < 1000; i++) {
      arr = [...arr, i];
    }
  })
  .on('cycle', function(event: Benchmark.Event) {
    console.log(String(event.target));
  })
  .on('complete', function(this: Benchmark.Suite) {
    console.log('Fastest is ' + this.filter('fastest').map('name'));
  })
  .run({ async: true });

// 输出:
// Array.push x 45,678 ops/sec ±1.23% (89 runs sampled)
// Array spread x 1,234 ops/sec ±2.34% (78 runs sampled)
// Fastest is Array.push

// ==================== Vitest benchmark ====================
import { bench, describe } from 'vitest';

describe('array operations', () => {
  bench('Array.push', () => {
    const arr = [];
    for (let i = 0; i < 1000; i++) {
      arr.push(i);
    }
  });

  bench('Array spread', () => {
    let arr = [];
    for (let i = 0; i < 1000; i++) {
      arr = [...arr, i];
    }
  });
});

// 运行: vitest bench

// ==================== 内存分析 ====================
// Node.js
const used = process.memoryUsage();
console.log(`heapUsed: ${Math.round(used.heapUsed / 1024 / 1024)} MB`);

// Chrome DevTools
// 1. Performance 面板录制
// 2. Memory 面板堆快照
// 3. Allocation Timeline

// ==================== 性能优化示例 ====================
// 避免重复计算
const memoize = <T extends (...args: any[]) => any>(fn: T): T => {
  const cache = new Map();
  return ((...args: any[]) => {
    const key = JSON.stringify(args);
    if (cache.has(key)) return cache.get(key);
    const result = fn(...args);
    cache.set(key, result);
    return result;
  }) as T;
};

// 使用 Web Worker 避免阻塞
const worker = new Worker('heavy-task.js');
```

### Python 性能测试

```python
import time
import timeit
import cProfile
import pstats

# ==================== 简单计时 ====================

start = time.time()
do_something()
end = time.time()
print(f"耗时: {end - start:.3f}s")

# time.perf_counter (更精确)
start = time.perf_counter()
do_something()
print(f"耗时: {time.perf_counter() - start:.6f}s")

# ==================== timeit (推荐) ====================

# 命令行
# python -m timeit "sum(range(1000))"

# 代码中
result = timeit.timeit(
    'sum(range(1000))',
    number=10000
)
print(f"平均耗时: {result / 10000 * 1000:.3f}ms")

# 比较多个实现
def method1():
    return sum(range(1000))

def method2():
    total = 0
    for i in range(1000):
        total += i
    return total

print(timeit.timeit(method1, number=10000))
print(timeit.timeit(method2, number=10000))

# ==================== pytest-benchmark ====================

# pip install pytest-benchmark

def test_my_function(benchmark):
    result = benchmark(my_function, arg1, arg2)
    assert result == expected

# 运行: pytest --benchmark-only
# 输出:
# Name                 Min      Max     Mean    StdDev   Rounds
# test_my_function    1.23ms   1.45ms  1.30ms   0.05ms   1000

# 比较基准
def test_compare(benchmark):
    benchmark.pedantic(
        my_function,
        args=(arg1,),
        iterations=100,
        rounds=10
    )

# ==================== cProfile 性能分析 ====================

# 命令行
# python -m cProfile -s cumtime my_script.py

# 代码中
cProfile.run('my_function()', 'output.prof')

# 分析结果
stats = pstats.Stats('output.prof')
stats.sort_stats('cumulative')
stats.print_stats(20)  # 前 20 行

# 装饰器形式
def profile(func):
    def wrapper(*args, **kwargs):
        profiler = cProfile.Profile()
        result = profiler.runcall(func, *args, **kwargs)
        profiler.print_stats(sort='cumulative')
        return result
    return wrapper

@profile
def my_function():
    pass

# ==================== line_profiler 逐行分析 ====================

# pip install line_profiler

# @profile  # 装饰需要分析的函数
# def slow_function():
#     ...

# 运行: kernprof -l -v my_script.py

# ==================== memory_profiler ====================

# pip install memory-profiler

from memory_profiler import profile

@profile
def my_function():
    a = [1] * 1000000
    b = [2] * 2000000
    del b
    return a

# 运行: python -m memory_profiler my_script.py

# ==================== py-spy 采样分析 ====================

# pip install py-spy

# 实时查看
# py-spy top --pid 12345

# 生成火焰图
# py-spy record -o profile.svg --pid 12345
# py-spy record -o profile.svg -- python my_script.py

# ==================== scalene (综合分析) ====================

# pip install scalene

# scalene my_script.py
# 同时分析 CPU、内存、GPU
```

### Go 性能测试

```go
import (
    "testing"
    "time"
)

// ==================== 简单计时 ====================

func main() {
    start := time.Now()
    doSomething()
    elapsed := time.Since(start)
    fmt.Printf("耗时: %v\n", elapsed)
}

// ==================== testing.B 基准测试 ====================

// benchmark_test.go
func BenchmarkArrayPush(b *testing.B) {
    for i := 0; i < b.N; i++ {
        arr := make([]int, 0)
        for j := 0; j < 1000; j++ {
            arr = append(arr, j)
        }
    }
}

func BenchmarkPrealloc(b *testing.B) {
    for i := 0; i < b.N; i++ {
        arr := make([]int, 0, 1000)
        for j := 0; j < 1000; j++ {
            arr = append(arr, j)
        }
    }
}

// 运行: go test -bench=. -benchmem
// 输出:
// BenchmarkArrayPush-8    50000    25000 ns/op    40000 B/op    20 allocs/op
// BenchmarkPrealloc-8    200000     8000 ns/op     8192 B/op     1 allocs/op

// ==================== 子基准测试 ====================

func BenchmarkSort(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
            data := generateData(size)
            b.ResetTimer()  // 重置计时器
            for i := 0; i < b.N; i++ {
                sort.Ints(data)
            }
        })
    }
}

// ==================== 并行基准测试 ====================

func BenchmarkParallel(b *testing.B) {
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            doSomething()
        }
    })
}

// ==================== pprof 性能分析 ====================

import (
    "net/http"
    _ "net/http/pprof"
    "runtime/pprof"
)

// HTTP 方式
func main() {
    go func() {
        http.ListenAndServe("localhost:6060", nil)
    }()
    // 访问 http://localhost:6060/debug/pprof/
}

// CPU profile
// go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30

// 内存 profile
// go tool pprof http://localhost:6060/debug/pprof/heap

// 文件方式
func main() {
    f, _ := os.Create("cpu.prof")
    pprof.StartCPUProfile(f)
    defer pprof.StopCPUProfile()
    
    // 业务代码
    doWork()
    
    // 内存 profile
    f2, _ := os.Create("mem.prof")
    pprof.WriteHeapProfile(f2)
}

// 分析
// go tool pprof cpu.prof
// (pprof) top10
// (pprof) web      # 生成 SVG
// (pprof) list funcName

// ==================== trace 追踪 ====================

import "runtime/trace"

func main() {
    f, _ := os.Create("trace.out")
    trace.Start(f)
    defer trace.Stop()
    
    // 业务代码
}

// 查看: go tool trace trace.out

// ==================== 基准测试比较 ====================

// 保存基准结果
// go test -bench=. -count=10 > old.txt
// (修改代码后)
// go test -bench=. -count=10 > new.txt

// 比较
// go install golang.org/x/perf/cmd/benchstat@latest
// benchstat old.txt new.txt
```

### Rust 性能测试

```rust
use std::time::Instant;

// ==================== 简单计时 ====================

fn main() {
    let start = Instant::now();
    do_something();
    let elapsed = start.elapsed();
    println!("耗时: {:?}", elapsed);
}

// ==================== 内置基准测试 (nightly) ====================

#![feature(test)]
extern crate test;

#[cfg(test)]
mod tests {
    use super::*;
    use test::Bencher;

    #[bench]
    fn bench_push(b: &mut Bencher) {
        b.iter(|| {
            let mut arr = Vec::new();
            for i in 0..1000 {
                arr.push(i);
            }
        });
    }
    
    #[bench]
    fn bench_prealloc(b: &mut Bencher) {
        b.iter(|| {
            let mut arr = Vec::with_capacity(1000);
            for i in 0..1000 {
                arr.push(i);
            }
        });
    }
}

// 运行: cargo +nightly bench

// ==================== criterion (推荐) ====================

// Cargo.toml
// [dev-dependencies]
// criterion = "0.5"
// [[bench]]
// name = "my_benchmark"
// harness = false

// benches/my_benchmark.rs
use criterion::{black_box, criterion_group, criterion_main, Criterion};

fn fibonacci(n: u64) -> u64 {
    match n {
        0 => 1,
        1 => 1,
        n => fibonacci(n - 1) + fibonacci(n - 2),
    }
}

fn criterion_benchmark(c: &mut Criterion) {
    c.bench_function("fib 20", |b| b.iter(|| fibonacci(black_box(20))));
}

// 比较多个实现
fn benchmark_compare(c: &mut Criterion) {
    let mut group = c.benchmark_group("Fibonacci");
    
    for i in [10, 15, 20].iter() {
        group.bench_with_input(
            BenchmarkId::new("recursive", i),
            i,
            |b, i| b.iter(|| fibonacci_recursive(*i))
        );
        group.bench_with_input(
            BenchmarkId::new("iterative", i),
            i,
            |b, i| b.iter(|| fibonacci_iterative(*i))
        );
    }
    group.finish();
}

criterion_group!(benches, criterion_benchmark, benchmark_compare);
criterion_main!(benches);

// 运行: cargo bench
// 输出包含统计信息、HTML 报告

// ==================== 防止优化 ====================

use std::hint::black_box;

// 防止编译器优化掉测试代码
let result = black_box(compute_something(black_box(input)));

// ==================== 火焰图 ====================

// 安装
// cargo install flamegraph

// 生成
// cargo flamegraph --bench my_benchmark
// 或
// cargo flamegraph -- my_args

// ==================== perf (Linux) ====================

// 构建 release with debug info
// [profile.release]
// debug = true

// perf record -g ./target/release/myapp
// perf report

// ==================== 内存分析 ====================

// DHAT (堆分析)
// cargo install dhat
// 添加到代码:
#[global_allocator]
static ALLOC: dhat::Alloc = dhat::Alloc;

fn main() {
    let _profiler = dhat::Profiler::new_heap();
    // ... 代码
}

// Valgrind (Linux)
// valgrind --tool=massif ./target/release/myapp
// ms_print massif.out.*

// ==================== 编译时间分析 ====================

// cargo build --timings
// 生成 cargo-timing.html

// 详细编译信息
// RUSTFLAGS="-Z time-passes" cargo +nightly build
```

### 性能测试对比

```
┌─────────────────┬────────────────┬────────────────┬────────────────┬────────────────┐
│ 特性            │ TypeScript     │ Python         │ Go             │ Rust           │
├─────────────────┼────────────────┼────────────────┼────────────────┼────────────────┤
│ 基准测试工具    │ Benchmark.js   │ pytest-benchmark│ testing.B     │ criterion      │
│ 内置支持        │ ❌             │ timeit         │ ✅             │ ✅ (nightly)   │
│ CPU 分析        │ Chrome DevTools│ cProfile       │ pprof          │ perf/flamegraph│
│ 内存分析        │ Chrome Memory  │ memory_profiler│ pprof heap     │ DHAT/Valgrind  │
│ 火焰图          │ Chrome         │ py-spy         │ pprof/go-torch │ flamegraph     │
│ 统计分析        │ ✅             │ ✅             │ benchstat      │ criterion      │
│ HTML 报告       │ ❌             │ ❌             │ ❌             │ ✅             │
└─────────────────┴────────────────┴────────────────┴────────────────┴────────────────┘
```

### 性能测试最佳实践

```
┌─────────────────────────────────────────────────────────────────────────┐
│                        性能测试最佳实践                                  │
├─────────────────────────────────────────────────────────────────────────┤
│                                                                         │
│  1. 测试环境                                                            │
│     □ 关闭其他程序，减少干扰                                           │
│     □ 使用固定的硬件环境                                               │
│     □ 多次运行取平均值                                                 │
│     □ 记录环境信息（CPU、内存、OS）                                    │
│                                                                         │
│  2. 测试方法                                                            │
│     □ 预热 (warmup) 避免冷启动影响                                     │
│     □ 使用 black_box 防止编译器优化                                    │
│     □ 测试真实数据，避免缓存命中                                       │
│     □ 分离 setup 和测试代码                                            │
│                                                                         │
│  3. 结果分析                                                            │
│     □ 关注 p50/p99 而非平均值                                          │
│     □ 检查标准差，确保结果稳定                                         │
│     □ 比较相对性能而非绝对数值                                         │
│     □ 结合火焰图定位热点                                               │
│                                                                         │
│  4. 持续监控                                                            │
│     □ CI 中运行基准测试                                                │
│     □ 对比历史数据                                                     │
│     □ 设置性能回退告警                                                 │
│                                                                         │
└─────────────────────────────────────────────────────────────────────────┘
```

### 常用命令速查

```bash
# TypeScript
npx vitest bench                    # Vitest 基准测试

# Python
python -m timeit "sum(range(1000))" # 快速计时
python -m cProfile -s time script.py # CPU 分析
py-spy record -o out.svg -- python script.py # 火焰图

# Go
go test -bench=. -benchmem          # 运行基准测试
go test -bench=. -cpuprofile=cpu.prof # CPU profile
go tool pprof -http=:8080 cpu.prof  # 可视化分析
benchstat old.txt new.txt           # 比较结果

# Rust
cargo bench                         # 运行基准测试
cargo flamegraph                    # 生成火焰图
cargo build --timings               # 编译时间分析
```

---

## 🎯 语言特性与特殊语法

### 特性支持概览

| 特性 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 解构赋值 | ✅ | ✅ | ❌ | ✅ |
| 展开运算符 | ✅ `...` | ✅ `*`/`**` | ✅ `...` | ❌ |
| 可选链 | ✅ `?.` | ❌ | ❌ | ❌ |
| 空值合并 | ✅ `??` | ❌ | ❌ | ❌ |
| 模式匹配 | ❌ | ✅ `match` | ❌ | ✅ `match` |
| 装饰器 | ✅ | ✅ | ❌ | ✅ (属性宏) |
| 运算符重载 | ❌ | ✅ | ❌ | ✅ |
| 宏系统 | ❌ | ❌ | ❌ | ✅ |
| 标签语句 | ✅ | ❌ | ✅ | ✅ |
| defer/finally | finally | finally | defer | Drop |

### TypeScript 特殊语法

```typescript
// ==================== 解构赋值 ====================
const { name, age } = user;                    // 对象解构
const { name: userName } = user;               // 重命名
const { name, ...rest } = user;                // 剩余属性
const { name, country = "USA" } = user;        // 默认值
const [first, second, ...rest] = [1, 2, 3, 4]; // 数组解构
const [a, , b] = [1, 2, 3];                    // 跳过元素
let a = 1, b = 2; [a, b] = [b, a];            // 交换变量

// ==================== 展开运算符 ====================
const arr2 = [...arr1, 4, 5];          // 数组展开
const obj2 = { ...obj1, c: 3 };        // 对象展开
Math.max(...args);                      // 函数调用展开
function sum(...nums: number[]) {}     // 剩余参数

// ==================== 可选链 (?.) ====================
user?.profile?.address?.city;          // 安全访问
user.getName?.();                       // 方法调用
arr?.[0];                               // 数组索引

// ==================== 空值合并 (??) ====================
const value = null ?? "default";       // "default"
const value2 = 0 ?? "default";         // 0 (只检查 null/undefined)
x ??= 10;                               // 赋值运算符

// ==================== 类型断言 ====================
const input = el as HTMLInputElement;   // as 断言
const colors = ["red", "blue"] as const; // const 断言
const config = {} satisfies Config;     // satisfies (TS 4.9+)

// ==================== 非空断言 (!) ====================
const element = document.getElementById("app")!;

// ==================== 标签语句 ====================
outer: for (let i = 0; i < 3; i++) {
    for (let j = 0; j < 3; j++) {
        if (condition) break outer;     // 跳出外层循环
    }
}

// ==================== 模板字面量类型 ====================
type EventName = `on${Capitalize<string>}`;

// ==================== 装饰器 ====================
@log
class MyClass {
    @validate
    method() {}
}
```

### Python 特殊语法

```python
# ==================== 解构赋值 ====================
a, b, c = (1, 2, 3)                    # 元组解包
a, b = b, a                            # 交换
first, *rest = [1, 2, 3, 4]            # 剩余元素
first, *middle, last = [1, 2, 3, 4, 5]

# ==================== 展开运算符 ====================
arr2 = [*arr1, 4, 5]                   # 列表展开
dict2 = {**dict1, "c": 3}              # 字典展开
print(*args, **kwargs)                 # 函数调用展开
def func(*args, **kwargs): pass        # 接收任意参数
def func(a, *, b, c): pass             # 仅关键字参数
def func(a, b, /, c): pass             # 仅位置参数 (3.8+)

# ==================== 模式匹配 (3.10+) ====================
match command:
    case ["load", filename]:
        load(filename)
    case {"action": "click", "x": x, "y": y}:
        click(x, y)
    case Point(x=0, y=y):
        print(f"On Y-axis at {y}")
    case _:
        print("Unknown")

# ==================== 海象运算符 (:=) 3.8+ ====================
if (n := len(data)) > 10:
    print(f"Too long: {n}")
while (line := file.readline()):
    process(line)

# ==================== f-string ====================
f"{name=}"                             # 调试: name='value'
f"{value:>10}"                         # 格式化
f"{value:.2f}"                         # 小数位

# ==================== 装饰器 ====================
@decorator
def func(): pass

@decorator_with_args(arg)
def func(): pass

# ==================== 上下文管理器 ====================
with open("file") as f, lock:
    content = f.read()

# ==================== 运算符重载 ====================
class Vector:
    def __add__(self, other): ...
    def __mul__(self, scalar): ...
    def __eq__(self, other): ...

# ==================== 生成器表达式 ====================
gen = (x**2 for x in range(1000))      # 惰性求值
lst = [x**2 for x in range(1000)]      # 立即求值
```

### Go 特殊语法

```go
// ==================== 短变量声明 ====================
x := 42                                // 类型推断
name, age := "Alice", 30               // 多变量
_, err := doSomething()                // 忽略值

// ==================== defer ====================
defer f.Close()                        // 函数返回前执行
// 多个 defer - LIFO 顺序执行

// ==================== 多返回值 ====================
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
result, err := divide(10, 2)

// 命名返回值
func f() (result int, err error) {
    result = 42
    return  // 裸 return
}

// ==================== 类型断言 ====================
s := i.(string)                        // 可能 panic
s, ok := i.(string)                    // 安全

// type switch
switch v := i.(type) {
case int:    fmt.Println("int", v)
case string: fmt.Println("string", v)
default:     fmt.Println("unknown")
}

// ==================== 标签与跳转 ====================
OuterLoop:
    for i := 0; i < 3; i++ {
        for j := 0; j < 3; j++ {
            if condition {
                break OuterLoop
            }
        }
    }

// ==================== iota 枚举 ====================
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
const (
    KB = 1 << (10 * iota)  // 1024
    MB                      // 1048576
    GB                      // 1073741824
)

// ==================== 展开运算符 ====================
sum(nums...)                           // 切片展开
combined := append(s1, s2...)          // append 展开

// ==================== 嵌入 (组合) ====================
type Dog struct {
    Animal                             // 嵌入
    Breed string
}
dog.Name                               // 访问嵌入字段
dog.Speak()                            // 调用嵌入方法

// ==================== 空白标识符 ====================
_, err := f()                          // 忽略返回值
import _ "pkg"                         // 仅导入副作用
for _, v := range slice {}             // 忽略索引

// ==================== init 函数 ====================
func init() {
    // 包初始化时自动执行
}

// ==================== 方法表达式 ====================
f := n.Double                          // 方法值
g := MyInt.Double                      // 方法表达式
```

### Rust 特殊语法

```rust
// ==================== 模式匹配 ====================
match n {
    0 => "zero",
    1 | 2 | 3 => "small",
    4..=9 => "medium",
    x if x < 0 => "negative",
    x @ 10..=20 => format!("{} in range", x),
    _ => "other",
}

// ==================== if let / while let ====================
if let Some(v) = maybe { use(v); }
while let Some(v) = iter.next() { use(v); }
let Some(v) = maybe else { return; };  // let else (1.65+)

// ==================== ? 运算符 ====================
let content = fs::read_to_string(path)?;  // 错误传播
let value = option?;                       // Option 也可用

// ==================== 解构赋值 ====================
let (x, y) = (1, 2);                       // 元组
let Point { x, y } = point;                // 结构体
let [first, rest @ ..] = arr;              // 数组

// ==================== 闭包 ====================
let add = |a, b| a + b;                    // 简化形式
let closure = move || println!("{}", s);  // move 获取所有权

// ==================== 迭代器链 ====================
numbers.iter()
    .filter(|&&x| x % 2 == 0)
    .map(|&x| x * 2)
    .collect::<Vec<_>>();

// ==================== 宏 ====================
println!("Hello, {}!", name);
vec![1, 2, 3];
assert_eq!(a, b);
dbg!(expression);                          // 调试打印
todo!();                                   // 标记未实现

// 声明式宏
macro_rules! say_hello {
    () => { println!("Hello!"); };
    ($name:expr) => { println!("Hello, {}!", $name); };
}

// ==================== 属性宏 ====================
#[derive(Debug, Clone, PartialEq)]
struct Point { x: i32, y: i32 }

#[cfg(test)]
mod tests {
    #[test]
    fn test_it() { assert!(true); }
}

#[inline]
#[allow(dead_code)]
#[deprecated(note = "use new_fn")]

// ==================== 运算符重载 ====================
impl Add for Vector {
    type Output = Self;
    fn add(self, other: Self) -> Self { ... }
}

// ==================== 标签循环 ====================
'outer: for i in 0..3 {
    for j in 0..3 {
        if cond { break 'outer; }
    }
}
let result = 'search: loop {
    if found { break 'search value; }
};

// ==================== turbofish ::<> ====================
let parsed = "42".parse::<i32>().unwrap();
let vec = iter.collect::<Vec<_>>();

// ==================== unsafe ====================
unsafe {
    let ptr = &x as *const i32;
    println!("{}", *ptr);
}
```

### 特殊语法对比

```
┌─────────────────┬──────────────────────┬──────────────────────┬──────────────────────┬──────────────────────┐
│ 特性            │ TypeScript           │ Python               │ Go                   │ Rust                 │
├─────────────────┼──────────────────────┼──────────────────────┼──────────────────────┼──────────────────────┤
│ 解构赋值        │ const {a} = obj      │ a, b = tuple         │ ❌                   │ let (a,b) = t        │
│ 展开            │ ...arr               │ *args, **kw          │ slice...             │ ❌                   │
│ 可选链          │ obj?.prop            │ ❌                   │ ❌                   │ ❌                   │
│ 空值合并        │ a ?? b               │ ❌                   │ ❌                   │ .unwrap_or(b)        │
│ 错误传播        │ throw                │ raise                │ return err           │ ?                    │
│ 模式匹配        │ ❌                   │ match                │ switch.(type)        │ match                │
│ 延迟执行        │ finally              │ finally              │ defer                │ Drop                 │
│ 装饰器          │ @decorator           │ @decorator           │ ❌                   │ #[attr]              │
│ 运算符重载      │ ❌                   │ __add__              │ ❌                   │ impl Add             │
│ 宏              │ ❌                   │ ❌                   │ ❌                   │ macro_rules!         │
│ 标签循环        │ label: for           │ ❌                   │ Label: for           │ 'label: for          │
└─────────────────┴──────────────────────┴──────────────────────┴──────────────────────┴──────────────────────┘
```

### 语法糖对照表

| 功能 | TypeScript | Python | Go | Rust |
|------|------------|--------|-----|------|
| 短路求值 | `a && b` | `a and b` | `a && b` | `a && b` |
| 三元运算 | `a ? b : c` | `b if a else c` | 无(用if) | `if a {b} else {c}` |
| 字符串插值 | `` `${x}` `` | `f"{x}"` | `fmt.Sprintf` | `format!("{}", x)` |
| 范围 | `for(;;)` | `range(n)` | `for i:=0;i<n;i++` | `0..n` |
| 包含范围 | 无 | 无 | 无 | `0..=n` |
| 类型推断 | `const x = 1` | `x = 1` | `x := 1` | `let x = 1` |
| 匿名函数 | `() => {}` | `lambda: x` | `func() {}` | `\|\| {}` |

### 独特语法特性

| 语言 | 独特特性 |
|------|----------|
| **TypeScript** | 可选链 `?.`、空值合并 `??`、类型守卫、模板字面量类型、satisfies |
| **Python** | 海象运算符 `:=`、f-string 调试 `{x=}`、模式匹配 `match`、切片步长 `[::2]` |
| **Go** | `defer`、多返回值、`iota` 枚举、类型嵌入、`init()` 函数、`:=` 短声明 |
| **Rust** | `?` 错误传播、`match` 模式匹配、生命周期 `'a`、宏系统、`unsafe`、turbofish `::<>` |

---

## 📚 总结

| 语言 | 一句话总结 |
|------|-----------|
| **TypeScript** | JavaScript 的类型安全升级版，Web 开发首选 |
| **Python** | 简单优雅，数据科学无敌，适合快速开发 |
| **Go** | 简单高效，云原生时代的工程语言 |
| **Rust** | 性能与安全的极致追求，系统编程新标杆 |

> **没有最好的语言，只有最适合的场景。** 选择语言时，应综合考虑项目需求、团队技能、生态系统和长期维护成本。
