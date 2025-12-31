# Iterator 迭代器详解：TypeScript vs Rust vs Python vs Go

## 目录
1. [迭代器基础](#1-迭代器基础)
2. [创建迭代器](#2-创建迭代器)
3. [常用迭代方法](#3-常用迭代方法)
4. [惰性求值与链式调用](#4-惰性求值与链式调用)
5. [自定义迭代器](#5-自定义迭代器)
6. [生成器](#6-生成器)
7. [并行迭代](#7-并行迭代)

---

## 1. 迭代器基础

### 迭代器协议对比

| 语言 | 协议/Trait | 核心方法 | 返回值 |
|------|-----------|---------|--------|
| TypeScript | `Iterable`/`Iterator` | `next()` | `{value, done}` |
| Rust | `Iterator` trait | `next()` | `Option<Item>` |
| Python | `__iter__`/`__next__` | `__next__()` | 值或 `StopIteration` |
| Go | 无内置协议 | `range` 关键字 | - |

### 基本迭代

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

// Go 1.23+ iter 包（实验性）
// 目前 Go 主要依赖 for-range，无通用迭代器
```

---

## 2. 创建迭代器

### 从集合创建

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

---

## 3. 常用迭代方法

### 3.1 Map（映射）

```typescript
// TypeScript
const nums = [1, 2, 3];

// Array.map（立即求值）
const doubled = nums.map(x => x * 2);  // [2, 4, 6]

// 链式
nums
    .map(x => x * 2)
    .map(x => x + 1);  // [3, 5, 7]
```

```rust
// Rust
let nums = vec![1, 2, 3];

// iter().map()（惰性）
let doubled: Vec<i32> = nums.iter()
    .map(|x| x * 2)
    .collect();  // [2, 4, 6]

// 链式
let result: Vec<i32> = nums.iter()
    .map(|x| x * 2)
    .map(|x| x + 1)
    .collect();  // [3, 5, 7]

// 不 collect，保持惰性
let lazy = nums.iter().map(|x| x * 2);
for x in lazy {
    println!("{}", x);
}
```

```python
# Python
nums = [1, 2, 3]

# map()（惰性）
doubled = map(lambda x: x * 2, nums)
list(doubled)  # [2, 4, 6]

# 列表推导式（立即求值，更 Pythonic）
doubled = [x * 2 for x in nums]

# 生成器表达式（惰性）
doubled = (x * 2 for x in nums)
```

```go
// Go - 无内置 map，手动实现
nums := []int{1, 2, 3}

doubled := make([]int, len(nums))
for i, v := range nums {
    doubled[i] = v * 2
}

// 泛型函数（Go 1.18+）
func Map[T, U any](slice []T, f func(T) U) []U {
    result := make([]U, len(slice))
    for i, v := range slice {
        result[i] = f(v)
    }
    return result
}

doubled := Map(nums, func(x int) int { return x * 2 })
```

### 3.2 Filter（过滤）

```typescript
// TypeScript
const nums = [1, 2, 3, 4, 5];

const evens = nums.filter(x => x % 2 === 0);  // [2, 4]

// 链式
nums
    .filter(x => x % 2 === 0)
    .map(x => x * 2);  // [4, 8]
```

```rust
// Rust
let nums = vec![1, 2, 3, 4, 5];

let evens: Vec<i32> = nums.iter()
    .filter(|x| *x % 2 == 0)
    .copied()  // &i32 -> i32
    .collect();  // [2, 4]

// 链式
let result: Vec<i32> = nums.iter()
    .filter(|x| *x % 2 == 0)
    .map(|x| x * 2)
    .collect();  // [4, 8]

// filter_map（同时过滤和映射）
let result: Vec<i32> = strings.iter()
    .filter_map(|s| s.parse::<i32>().ok())
    .collect();
```

```python
# Python
nums = [1, 2, 3, 4, 5]

# filter()（惰性）
evens = filter(lambda x: x % 2 == 0, nums)
list(evens)  # [2, 4]

# 列表推导式
evens = [x for x in nums if x % 2 == 0]

# 生成器表达式
evens = (x for x in nums if x % 2 == 0)

# 链式（推导式更清晰）
result = [x * 2 for x in nums if x % 2 == 0]  # [4, 8]
```

```go
// Go - 手动实现
nums := []int{1, 2, 3, 4, 5}

var evens []int
for _, v := range nums {
    if v%2 == 0 {
        evens = append(evens, v)
    }
}

// 泛型函数
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

### 3.3 Reduce/Fold（归约）

```typescript
// TypeScript
const nums = [1, 2, 3, 4, 5];

// reduce
const sum = nums.reduce((acc, x) => acc + x, 0);  // 15
const product = nums.reduce((acc, x) => acc * x, 1);  // 120

// reduceRight（从右向左）
const result = nums.reduceRight((acc, x) => acc - x, 0);
```

```rust
// Rust
let nums = vec![1, 2, 3, 4, 5];

// fold（需要初始值）
let sum: i32 = nums.iter().fold(0, |acc, x| acc + x);  // 15

// reduce（使用第一个元素作为初始值）
let sum: Option<i32> = nums.iter().copied().reduce(|acc, x| acc + x);

// sum / product（特化方法）
let sum: i32 = nums.iter().sum();
let product: i32 = nums.iter().product();

// 从右向左
let result: i32 = nums.iter().rev().fold(0, |acc, x| acc - x);
```

```python
# Python
from functools import reduce

nums = [1, 2, 3, 4, 5]

# reduce
sum_val = reduce(lambda acc, x: acc + x, nums, 0)  # 15

# 内置函数
sum(nums)      # 15
max(nums)      # 5
min(nums)      # 1

# any / all
any(x > 3 for x in nums)  # True
all(x > 0 for x in nums)  # True
```

```go
// Go - 手动实现
nums := []int{1, 2, 3, 4, 5}

sum := 0
for _, v := range nums {
    sum += v
}

// 泛型 Reduce
func Reduce[T, U any](slice []T, init U, f func(U, T) U) U {
    acc := init
    for _, v := range slice {
        acc = f(acc, v)
    }
    return acc
}

sum := Reduce(nums, 0, func(acc, x int) int { return acc + x })
```

### 3.4 其他常用方法

```typescript
// TypeScript
const nums = [1, 2, 3, 4, 5];

// 查找
nums.find(x => x > 3);           // 4
nums.findIndex(x => x > 3);      // 3
nums.includes(3);                // true
nums.indexOf(3);                 // 2

// 判断
nums.some(x => x > 3);           // true（任一满足）
nums.every(x => x > 0);          // true（全部满足）

// 切片
nums.slice(1, 3);                // [2, 3]

// 展平
[[1, 2], [3, 4]].flat();         // [1, 2, 3, 4]
nums.flatMap(x => [x, x * 2]);   // [1, 2, 2, 4, ...]

// 排序（原地修改）
nums.sort((a, b) => a - b);
[...nums].sort();                // 不修改原数组

// 去重
[...new Set(nums)];

// take / skip
nums.slice(0, 3);                // take 3
nums.slice(2);                   // skip 2
```

```rust
// Rust
let nums = vec![1, 2, 3, 4, 5];

// 查找
nums.iter().find(|&&x| x > 3);         // Some(&4)
nums.iter().position(|&x| x > 3);      // Some(3)
nums.contains(&3);                      // true

// 判断
nums.iter().any(|&x| x > 3);           // true
nums.iter().all(|&x| x > 0);           // true

// take / skip（惰性）
nums.iter().take(3);                   // 前3个
nums.iter().skip(2);                   // 跳过前2个
nums.iter().take_while(|&&x| x < 4);   // 满足条件时取
nums.iter().skip_while(|&&x| x < 3);   // 满足条件时跳

// 展平
let nested = vec![vec![1, 2], vec![3, 4]];
nested.into_iter().flatten().collect::<Vec<_>>();  // [1, 2, 3, 4]

nums.iter().flat_map(|&x| vec![x, x * 2]);

// 排序（需要 mut）
let mut nums = nums;
nums.sort();
nums.sort_by(|a, b| b.cmp(a));         // 降序
nums.sort_by_key(|x| -x);              // 按键排序

// 去重（需要先排序）
nums.dedup();

// 分区
let (evens, odds): (Vec<_>, Vec<_>) = nums.iter()
    .partition(|&&x| x % 2 == 0);

// 分组计数
use std::collections::HashMap;
let counts: HashMap<_, _> = nums.iter()
    .fold(HashMap::new(), |mut acc, &x| {
        *acc.entry(x).or_insert(0) += 1;
        acc
    });

// 第 n 个
nums.iter().nth(2);                    // Some(&3)

// 最值
nums.iter().max();
nums.iter().min();
nums.iter().max_by_key(|x| *x);

// 计数
nums.iter().count();
nums.iter().filter(|&&x| x > 2).count();

// 收集为其他类型
let set: HashSet<_> = nums.iter().collect();
let s: String = chars.iter().collect();
```

```python
# Python
nums = [1, 2, 3, 4, 5]

# 查找
next((x for x in nums if x > 3), None)  # 4
next((i for i, x in enumerate(nums) if x > 3), None)  # 3
3 in nums                               # True
nums.index(3)                           # 2

# 判断
any(x > 3 for x in nums)               # True
all(x > 0 for x in nums)               # True

# 切片
nums[1:3]                              # [2, 3]
nums[:3]                               # take 3
nums[2:]                               # skip 2

# 展平
from itertools import chain
list(chain.from_iterable([[1, 2], [3, 4]]))  # [1, 2, 3, 4]

# itertools 模块
from itertools import takewhile, dropwhile, islice, groupby

list(takewhile(lambda x: x < 4, nums))  # [1, 2, 3]
list(dropwhile(lambda x: x < 3, nums))  # [3, 4, 5]
list(islice(nums, 3))                   # [1, 2, 3]

# 排序
sorted(nums)                           # 新列表
sorted(nums, reverse=True)             # 降序
sorted(nums, key=lambda x: -x)         # 按键

# 去重
list(set(nums))                        # 无序
list(dict.fromkeys(nums))              # 保持顺序

# 分组
from itertools import groupby
for key, group in groupby(sorted(data), key=lambda x: x['type']):
    print(key, list(group))

# 计数
from collections import Counter
Counter(nums)                          # {1: 1, 2: 1, ...}
```

```go
// Go - 标准库 slices 包（Go 1.21+）
import "slices"

nums := []int{1, 2, 3, 4, 5}

// 查找
slices.Contains(nums, 3)               // true
slices.Index(nums, 3)                  // 2

// 排序
slices.Sort(nums)
slices.SortFunc(nums, func(a, b int) int { return b - a })

// 最值
slices.Max(nums)
slices.Min(nums)

// 去重（需先排序）
slices.Sort(nums)
nums = slices.Compact(nums)

// 反转
slices.Reverse(nums)

// 二分查找
slices.BinarySearch(nums, 3)
```

---

## 4. 惰性求值与链式调用

### 惰性 vs 立即求值

```typescript
// TypeScript - Array 方法是立即求值的
const result = [1, 2, 3, 4, 5]
    .map(x => {
        console.log('map', x);  // 全部执行
        return x * 2;
    })
    .filter(x => x > 4)
    .slice(0, 2);

// 惰性迭代（使用生成器）
function* lazyMap<T, U>(iter: Iterable<T>, fn: (x: T) => U) {
    for (const x of iter) {
        yield fn(x);
    }
}
```

```rust
// Rust - 迭代器是惰性的！
let result: Vec<i32> = vec![1, 2, 3, 4, 5]
    .into_iter()
    .map(|x| {
        println!("map {}", x);  // 只在消费时执行
        x * 2
    })
    .filter(|&x| x > 4)
    .take(2)
    .collect();  // 触发求值

// 不 collect 就不会执行任何操作
let lazy = (0..1000000)
    .map(|x| x * 2)
    .filter(|x| x % 3 == 0);
// 此时什么都没发生

// 只取需要的
let first_three: Vec<_> = lazy.take(3).collect();
```

```python
# Python
# map/filter 是惰性的
lazy = map(lambda x: x * 2, range(1000000))  # 不执行
list(lazy)  # 触发求值

# 列表推导式是立即求值的
eager = [x * 2 for x in range(1000000)]  # 立即执行

# 生成器表达式是惰性的
lazy = (x * 2 for x in range(1000000))  # 不执行

# itertools 全是惰性的
from itertools import islice, takewhile
result = islice(map(lambda x: x * 2, range(1000000)), 3)
list(result)  # [0, 2, 4]
```

```go
// Go - 没有惰性迭代器，range 是立即迭代的
// 可以用 channel 模拟惰性

func lazyRange(n int) <-chan int {
    ch := make(chan int)
    go func() {
        for i := 0; i < n; i++ {
            ch <- i
        }
        close(ch)
    }()
    return ch
}

for v := range lazyRange(1000000) {
    if v >= 3 {
        break  // 提前终止
    }
    fmt.Println(v * 2)
}
```

### 链式调用示例

```typescript
// TypeScript
interface User {
    name: string;
    age: number;
    active: boolean;
}

const users: User[] = [...];

const result = users
    .filter(u => u.active)
    .filter(u => u.age >= 18)
    .map(u => u.name)
    .sort()
    .slice(0, 10);
```

```rust
// Rust
struct User {
    name: String,
    age: u32,
    active: bool,
}

let result: Vec<String> = users.iter()
    .filter(|u| u.active)
    .filter(|u| u.age >= 18)
    .map(|u| u.name.clone())
    .sorted()  // 需要 itertools crate
    .take(10)
    .collect();

// 使用 itertools 获得更多方法
use itertools::Itertools;

let result: Vec<_> = nums.iter()
    .unique()
    .sorted()
    .chunks(3)  // 分块
    .into_iter()
    .map(|chunk| chunk.sum::<i32>())
    .collect();
```

```python
# Python - 可以用 more-itertools 或链式生成器
from itertools import islice

# 方法链（使用库如 toolz 或 pipe）
from toolz import pipe, curry
from toolz.curried import map, filter, take

result = pipe(
    users,
    filter(lambda u: u['active']),
    filter(lambda u: u['age'] >= 18),
    map(lambda u: u['name']),
    sorted,
    take(10),
    list
)

# 或者用生成器表达式
names = (u['name'] for u in users if u['active'] and u['age'] >= 18)
result = sorted(names)[:10]
```

```go
// Go - 使用 lo 库实现链式
import "github.com/samber/lo"

result := lo.Map(
    lo.Filter(users, func(u User, _ int) bool {
        return u.Active && u.Age >= 18
    }),
    func(u User, _ int) string {
        return u.Name
    },
)
sort.Strings(result)
result = result[:min(10, len(result))]
```

---

## 5. 自定义迭代器

### TypeScript

```typescript
// 实现 Iterable 接口
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

// 使用
for (const n of new Range(0, 5)) {
    console.log(n);  // 0, 1, 2, 3, 4
}

// 展开
const arr = [...new Range(0, 5)];
```

### Rust

```rust
// 实现 Iterator trait
struct Counter {
    current: u32,
    max: u32,
}

impl Counter {
    fn new(max: u32) -> Self {
        Counter { current: 0, max }
    }
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

// 使用
for n in Counter::new(5) {
    println!("{}", n);  // 0, 1, 2, 3, 4
}

// 链式调用（自动获得所有迭代器方法！）
let sum: u32 = Counter::new(10)
    .filter(|&x| x % 2 == 0)
    .sum();  // 20

// 实现 IntoIterator 让类型可以被 for 循环
impl IntoIterator for MyCollection {
    type Item = i32;
    type IntoIter = std::vec::IntoIter<i32>;

    fn into_iter(self) -> Self::IntoIter {
        self.data.into_iter()
    }
}
```

### Python

```python
# 实现 __iter__ 和 __next__
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

# 使用
for n in Counter(5):
    print(n)  # 0, 1, 2, 3, 4

# 更简单：使用生成器函数
def counter(max_val: int):
    current = 0
    while current < max_val:
        yield current
        current += 1

# 或者 __iter__ 返回生成器
class Counter:
    def __init__(self, max_val: int):
        self.max = max_val
    
    def __iter__(self):
        for i in range(self.max):
            yield i
```

### Go

```go
// Go 没有迭代器协议，常用模式：

// 1. 返回 channel
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

// 2. 回调函数模式
func ForEach(max int, fn func(int)) {
    for i := 0; i < max; i++ {
        fn(i)
    }
}

ForEach(5, func(n int) {
    fmt.Println(n)
})

// 3. Next 方法模式
type Counter struct {
    current, max int
}

func NewCounter(max int) *Counter {
    return &Counter{max: max}
}

func (c *Counter) Next() (int, bool) {
    if c.current < c.max {
        result := c.current
        c.current++
        return result, true
    }
    return 0, false
}

// 使用
c := NewCounter(5)
for n, ok := c.Next(); ok; n, ok = c.Next() {
    fmt.Println(n)
}

// 4. Go 1.23+ iter 包（实验性）
// iter.Seq[V] 和 iter.Seq2[K, V]
```

---

## 6. 生成器

### TypeScript

```typescript
// Generator 函数
function* range(start: number, end: number): Generator<number> {
    for (let i = start; i < end; i++) {
        yield i;
    }
}

// 使用
for (const n of range(0, 5)) {
    console.log(n);
}

// 无限生成器
function* naturals(): Generator<number> {
    let n = 0;
    while (true) {
        yield n++;
    }
}

// 双向通信
function* chat(): Generator<string, void, string> {
    const name = yield "What's your name?";
    yield `Hello, ${name}!`;
}

const gen = chat();
console.log(gen.next().value);        // "What's your name?"
console.log(gen.next("Alice").value); // "Hello, Alice!"

// yield* 委托
function* concat<T>(...iters: Iterable<T>[]): Generator<T> {
    for (const iter of iters) {
        yield* iter;
    }
}
```

### Rust

```rust
// Rust 没有内置生成器语法，但可以用 Iterator 实现

// 使用闭包创建迭代器
fn range(start: i32, end: i32) -> impl Iterator<Item = i32> {
    (start..end)
}

// 使用 std::iter::from_fn
fn naturals() -> impl Iterator<Item = i32> {
    let mut n = 0;
    std::iter::from_fn(move || {
        let result = n;
        n += 1;
        Some(result)
    })
}

// 使用 std::iter::successors
let powers_of_2 = std::iter::successors(Some(1), |&n| Some(n * 2));

// 使用 gen 块（nightly，实验性）
#![feature(gen_blocks)]
fn range(start: i32, end: i32) -> impl Iterator<Item = i32> {
    gen {
        for i in start..end {
            yield i;
        }
    }
}

// 第三方库：genawaiter
use genawaiter::{sync::gen, yield_};

let generator = gen!({
    yield_!(1);
    yield_!(2);
    yield_!(3);
});
```

### Python

```python
# 生成器函数
def range_gen(start: int, end: int):
    i = start
    while i < end:
        yield i
        i += 1

# 使用
for n in range_gen(0, 5):
    print(n)

# 无限生成器
def naturals():
    n = 0
    while True:
        yield n
        n += 1

# 带 send 的双向通信
def chat():
    name = yield "What's your name?"
    yield f"Hello, {name}!"

gen = chat()
print(next(gen))           # "What's your name?"
print(gen.send("Alice"))   # "Hello, Alice!"

# yield from 委托
def concat(*iterables):
    for it in iterables:
        yield from it

# 异步生成器
async def async_range(start: int, end: int):
    for i in range(start, end):
        await asyncio.sleep(0.1)
        yield i

async for n in async_range(0, 5):
    print(n)
```

### Go

```go
// Go 没有生成器语法，用 goroutine + channel 模拟

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

// 无限生成器
func Naturals() <-chan int {
    ch := make(chan int)
    go func() {
        n := 0
        for {
            ch <- n
            n++
        }
    }()
    return ch
}

// 需要手动停止
func NaturalsWithCancel(ctx context.Context) <-chan int {
    ch := make(chan int)
    go func() {
        defer close(ch)
        n := 0
        for {
            select {
            case <-ctx.Done():
                return
            case ch <- n:
                n++
            }
        }
    }()
    return ch
}

ctx, cancel := context.WithCancel(context.Background())
for n := range NaturalsWithCancel(ctx) {
    fmt.Println(n)
    if n >= 10 {
        cancel()
        break
    }
}
```

---

## 7. 并行迭代

### TypeScript

```typescript
// Promise.all 并行处理
const urls = ['url1', 'url2', 'url3'];

const results = await Promise.all(
    urls.map(url => fetch(url))
);

// 控制并发数
async function parallelLimit<T, R>(
    items: T[],
    limit: number,
    fn: (item: T) => Promise<R>
): Promise<R[]> {
    const results: R[] = [];
    const executing: Promise<void>[] = [];

    for (const item of items) {
        const p = fn(item).then(result => {
            results.push(result);
        });
        executing.push(p);

        if (executing.length >= limit) {
            await Promise.race(executing);
        }
    }

    await Promise.all(executing);
    return results;
}
```

### Rust

```rust
// Rayon - 数据并行库
use rayon::prelude::*;

let nums: Vec<i32> = (0..1000000).collect();

// 并行 map
let doubled: Vec<i32> = nums.par_iter()
    .map(|x| x * 2)
    .collect();

// 并行 filter
let evens: Vec<i32> = nums.par_iter()
    .filter(|&&x| x % 2 == 0)
    .copied()
    .collect();

// 并行 reduce
let sum: i32 = nums.par_iter().sum();

// 链式并行操作
let result: i32 = nums.par_iter()
    .filter(|&&x| x % 2 == 0)
    .map(|x| x * 2)
    .sum();

// tokio 并发（异步）
use futures::future::join_all;

let urls = vec!["url1", "url2", "url3"];
let futures: Vec<_> = urls.iter()
    .map(|url| fetch(url))
    .collect();

let results = join_all(futures).await;

// tokio::spawn 并行任务
let handles: Vec<_> = urls.iter()
    .map(|url| {
        let url = url.to_string();
        tokio::spawn(async move { fetch(&url).await })
    })
    .collect();

let results: Vec<_> = futures::future::try_join_all(handles).await?;
```

### Python

```python
# concurrent.futures - 线程/进程池
from concurrent.futures import ThreadPoolExecutor, ProcessPoolExecutor

urls = ['url1', 'url2', 'url3']

# 线程池（IO 密集型）
with ThreadPoolExecutor(max_workers=10) as executor:
    results = list(executor.map(fetch, urls))

# 进程池（CPU 密集型，绕过 GIL）
with ProcessPoolExecutor() as executor:
    results = list(executor.map(cpu_bound_task, data))

# asyncio 并发
import asyncio

async def fetch_all(urls):
    tasks = [fetch(url) for url in urls]
    return await asyncio.gather(*tasks)

# 控制并发数
async def fetch_with_limit(urls, limit=10):
    semaphore = asyncio.Semaphore(limit)
    
    async def fetch_one(url):
        async with semaphore:
            return await fetch(url)
    
    return await asyncio.gather(*[fetch_one(url) for url in urls])

# multiprocessing 并行
from multiprocessing import Pool

with Pool(4) as p:
    results = p.map(cpu_task, data)
```

### Go

```go
// Goroutine + WaitGroup
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

// Worker Pool 模式
func workerPool(items []int, workers int, fn func(int) int) []int {
    jobs := make(chan int, len(items))
    results := make(chan int, len(items))
    
    // 启动 workers
    for w := 0; w < workers; w++ {
        go func() {
            for item := range jobs {
                results <- fn(item)
            }
        }()
    }
    
    // 发送任务
    for _, item := range items {
        jobs <- item
    }
    close(jobs)
    
    // 收集结果
    output := make([]int, len(items))
    for i := 0; i < len(items); i++ {
        output[i] = <-results
    }
    return output
}

// errgroup 处理错误
import "golang.org/x/sync/errgroup"

func fetchAll(urls []string) ([]string, error) {
    results := make([]string, len(urls))
    g, ctx := errgroup.WithContext(context.Background())
    
    for i, url := range urls {
        i, url := i, url  // 捕获变量
        g.Go(func() error {
            result, err := fetchWithContext(ctx, url)
            if err != nil {
                return err
            }
            results[i] = result
            return nil
        })
    }
    
    if err := g.Wait(); err != nil {
        return nil, err
    }
    return results, nil
}
```

---

## 总结对比

| 特性 | TypeScript | Rust | Python | Go |
|------|------------|------|--------|-----|
| **迭代器协议** | `Symbol.iterator` | `Iterator` trait | `__iter__/__next__` | 无（用 range） |
| **惰性求值** | ❌ (数组方法) | ✅ 默认惰性 | ✅ map/filter | ❌ |
| **链式调用** | ✅ 数组方法 | ✅ 丰富 | ⚠️ 需要库 | ⚠️ 需要库 |
| **生成器** | ✅ `function*` | ⚠️ nightly | ✅ `yield` | ❌ (用 channel) |
| **并行迭代** | Promise.all | rayon | multiprocessing | goroutine |
| **类型安全** | ✅ | ✅ 最强 | ⚠️ 运行时 | ✅ |
| **零成本抽象** | ❌ | ✅ | ❌ | ⚠️ |

### 最佳实践建议

- **TypeScript**: 小数据用数组方法，大数据用生成器
- **Rust**: 优先使用迭代器链，性能最优且安全
- **Python**: 大数据用生成器表达式，CPU密集用多进程
- **Go**: 用 range + goroutine，考虑使用 lo 库简化
