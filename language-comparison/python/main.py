#!/usr/bin/env python3
"""
Python 语言特性全面示例
"""
from __future__ import annotations

import asyncio
import re
from collections import namedtuple
from dataclasses import dataclass
from datetime import datetime, timedelta
from enum import Enum, auto
from functools import reduce
from itertools import islice, chain, cycle
from pathlib import Path
from typing import (
    TypeVar, Generic, Protocol, Iterator, 
    Callable, Any
)

# ============================================
# 1. 基本类型与变量
# ============================================
print("=== 1. 基本类型与变量 ===")

# 变量声明 (动态类型 + 类型提示)
name: str = "Alice"
age: int = 25
height: float = 1.75
is_active: bool = True
nothing: None = None

# 类型推断
score = 100  # int
message = "Hello"  # str

print(f"姓名: {name}, 年龄: {age}, 身高: {height}, 活跃: {is_active}")

# ============================================
# 2. 列表 (List)
# ============================================
print("\n=== 2. 列表 (List) ===")

numbers: list[int] = [1, 2, 3, 4, 5]
mixed: list = [1, "two", 3.0, True]

# 列表方法
print("原列表:", numbers)
numbers_copy = numbers.copy()

numbers_copy.append(6)
print("append(6):", numbers_copy)

numbers_copy.pop()
print("pop():", numbers_copy)

print("slice [1:3]:", numbers[1:3])
print("map (x2):", list(map(lambda x: x * 2, numbers)))
print("filter (>2):", list(filter(lambda x: x > 2, numbers)))
print("reduce (sum):", reduce(lambda a, b: a + b, numbers, 0))
print("3 in numbers:", 3 in numbers)

# 列表推导式
squares = [x**2 for x in range(1, 6)]
evens = [x for x in numbers if x % 2 == 0]
print("列表推导式 (平方):", squares)
print("列表推导式 (偶数):", evens)

# ============================================
# 3. 元组 (Tuple)
# ============================================
print("\n=== 3. 元组 (Tuple) ===")

point: tuple[int, int] = (10, 20)
record: tuple[str, int, bool] = ("Alice", 25, True)

# 解构
x, y = point
print(f"Point: ({x}, {y})")

# 命名元组
Point = namedtuple('Point', ['x', 'y'])
p = Point(10, 20)
print(f"命名元组: x={p.x}, y={p.y}")

# ============================================
# 4. 集合 (Set)
# ============================================
print("\n=== 4. 集合 (Set) ===")

numbers_set: set[int] = {1, 2, 3, 3, 4}  # 重复自动去除
print("Set:", numbers_set)

numbers_set.add(5)
numbers_set.discard(1)
print("add(5), discard(1):", numbers_set)
print("2 in set:", 2 in numbers_set)

# 集合运算
set1 = {1, 2, 3}
set2 = {2, 3, 4}
print("并集 |:", set1 | set2)
print("交集 &:", set1 & set2)
print("差集 -:", set1 - set2)
print("对称差集 ^:", set1 ^ set2)

# ============================================
# 5. 字典 (Dict)
# ============================================
print("\n=== 5. 字典 (Dict) ===")

person: dict[str, Any] = {
    "name": "Alice",
    "age": 25,
    "email": "alice@example.com"
}

print("字典:", person)
print("get('name'):", person.get("name"))
print("get('phone', 'N/A'):", person.get("phone", "N/A"))
print("keys():", list(person.keys()))
print("values():", list(person.values()))
print("items():", list(person.items()))

# 字典推导式
squared_dict = {x: x**2 for x in range(1, 6)}
print("字典推导式:", squared_dict)

# ============================================
# 6. 字符串操作
# ============================================
print("\n=== 6. 字符串操作 ===")

s = "  Hello, World!  "
print("原字符串:", repr(s))
print("len():", len(s))
print("strip():", repr(s.strip()))
print("upper():", s.upper())
print("lower():", s.lower())
print("split(','):", s.strip().split(", "))
print("'World' in s:", "World" in s)
print("startswith('  H'):", s.startswith("  H"))
print("replace():", s.replace("World", "Python"))
print("slice [2:7]:", s[2:7])

# f-string
greeting = f"你好, {name}! 你今年 {age} 岁。"
print("f-string:", greeting)

# ============================================
# 7. 类与继承
# ============================================
print("\n=== 7. 类与继承 ===")


class Person:
    """基础 Person 类"""
    
    def __init__(self, name: str, age: int):
        self._name = name  # 约定私有
        self.age = age
    
    def greet(self) -> str:
        return f"Hello, I'm {self._name}"
    
    @property
    def name(self) -> str:
        return self._name
    
    @property
    def info(self) -> str:
        return f"{self._name}, {self.age} years old"
    
    @staticmethod
    def create(name: str) -> Person:
        return Person(name, 0)
    
    @classmethod
    def from_dict(cls, data: dict) -> Person:
        return cls(data["name"], data["age"])


class Student(Person):
    """Student 继承自 Person"""
    
    def __init__(self, name: str, age: int, school: str):
        super().__init__(name, age)
        self.school = school
    
    def greet(self) -> str:
        return f"{super().greet()}, I study at {self.school}"


# Dataclass
@dataclass
class Point3D:
    x: float
    y: float
    z: float = 0.0


person = Person("Alice", 25)
student = Student("Bob", 20, "MIT")
point3d = Point3D(1.0, 2.0, 3.0)

print("Person:", person.greet())
print("Student:", student.greet())
print("Info:", person.info)
print("Dataclass:", point3d)

# ============================================
# 8. Protocol (接口)
# ============================================
print("\n=== 8. Protocol (接口) ===")


class Greeter(Protocol):
    """Greeter Protocol"""
    def greet(self) -> str: ...


def print_greeting(g: Greeter) -> None:
    print(f"Greeting: {g.greet()}")


# Person 和 Student 都实现了 greet()，自动满足 Protocol
print_greeting(person)
print_greeting(student)

# ============================================
# 9. 枚举 (Enum)
# ============================================
print("\n=== 9. 枚举 (Enum) ===")


class Direction(Enum):
    UP = auto()
    DOWN = auto()
    LEFT = auto()
    RIGHT = auto()


class HttpStatus(Enum):
    OK = 200
    NOT_FOUND = 404
    INTERNAL_ERROR = 500


print("Direction.UP:", Direction.UP)
print("Direction.UP.value:", Direction.UP.value)
print("HttpStatus.OK:", HttpStatus.OK)
print("HttpStatus.OK.value:", HttpStatus.OK.value)

# ============================================
# 10. 泛型
# ============================================
print("\n=== 10. 泛型 ===")

T = TypeVar('T')


class Container(Generic[T]):
    def __init__(self, value: T):
        self._value = value
    
    def get(self) -> T:
        return self._value
    
    def set(self, value: T) -> None:
        self._value = value


int_container: Container[int] = Container(42)
str_container: Container[str] = Container("Hello")

print("int_container:", int_container.get())
print("str_container:", str_container.get())

# ============================================
# 11. 错误处理
# ============================================
print("\n=== 11. 错误处理 ===")


class ValidationError(Exception):
    def __init__(self, field: str, message: str):
        self.field = field
        super().__init__(message)


def divide(a: float, b: float) -> float:
    if b == 0:
        raise ValueError("Division by zero")
    return a / b


def validate_age(age: int) -> None:
    if age < 0:
        raise ValidationError("age", "Age cannot be negative")


try:
    print("10 / 2 =", divide(10, 2))
    print("10 / 0 =", divide(10, 0))
except ValueError as e:
    print("ValueError:", e)

try:
    validate_age(-5)
except ValidationError as e:
    print(f"ValidationError [{e.field}]:", e)

# ============================================
# 12. 日期时间
# ============================================
print("\n=== 12. 日期时间 ===")

now = datetime.now()
print("当前时间:", now.isoformat())
print("年:", now.year)
print("月:", now.month)
print("日:", now.day)
print("时间戳:", now.timestamp())

# 格式化
print("格式化:", now.strftime("%Y-%m-%d %H:%M:%S"))

# 解析
parsed = datetime.strptime("2024-01-15", "%Y-%m-%d")
print("解析:", parsed)

# 计算
tomorrow = now + timedelta(days=1)
print("明天:", tomorrow.strftime("%Y-%m-%d"))

diff = datetime(2024, 12, 31) - datetime(2024, 1, 1)
print("日期差:", diff.days, "天")

# ============================================
# 13. 正则表达式
# ============================================
print("\n=== 13. 正则表达式 ===")

text = "Call 123-4567 or 987-6543"
pattern = r"\d{3}-\d{4}"

print("文本:", text)
print("search():", re.search(pattern, text))
print("findall():", re.findall(pattern, text))
print("sub():", re.sub(pattern, "XXX-XXXX", text))

# 捕获组
email_text = "Contact: test@example.com"
email_pattern = r"(?P<user>\w+)@(?P<domain>\w+\.\w+)"
match = re.search(email_pattern, email_text)
if match:
    print("邮箱捕获组:", match.groupdict())

# ============================================
# 14. 迭代器与生成器
# ============================================
print("\n=== 14. 迭代器与生成器 ===")


# 生成器函数
def range_gen(start: int, end: int) -> Iterator[int]:
    i = start
    while i < end:
        yield i
        i += 1


print("生成器 range_gen(0, 5):", list(range_gen(0, 5)))


# 自定义迭代器
class Counter:
    def __init__(self, max_count: int):
        self.max = max_count
        self.count = 0
    
    def __iter__(self) -> Iterator[int]:
        return self
    
    def __next__(self) -> int:
        if self.count >= self.max:
            raise StopIteration
        self.count += 1
        return self.count - 1


print("自定义迭代器:", list(Counter(5)))

# itertools
print("islice(cycle, 10):", list(islice(cycle([1, 2, 3]), 10)))
print("chain:", list(chain([1, 2], [3, 4])))

# 链式操作
result = reduce(
    lambda a, b: a + b,
    map(lambda x: x * 2, filter(lambda x: x % 2 == 0, range(1, 11))),
    0
)
print("链式操作 (偶数 * 2 求和):", result)

# ============================================
# 15. 异步编程
# ============================================
print("\n=== 15. 异步编程 ===")


async def fetch_data(id: int) -> dict:
    await asyncio.sleep(0.1)
    return {"id": id, "data": f"Data for {id}"}


async def async_demo():
    print("开始异步操作...")
    
    # 顺序执行
    result1 = await fetch_data(1)
    print("顺序执行:", result1)
    
    # 并行执行
    results = await asyncio.gather(
        fetch_data(1),
        fetch_data(2),
        fetch_data(3),
    )
    print("并行执行:", results)
    
    # 超时
    try:
        result = await asyncio.wait_for(fetch_data(1), timeout=5.0)
        print("超时测试:", result)
    except asyncio.TimeoutError:
        print("超时!")


asyncio.run(async_demo())

# ============================================
# 16. 文件 IO
# ============================================
print("\n=== 16. 文件 IO ===")

test_dir = Path(__file__).parent / "test_output"
test_dir.mkdir(exist_ok=True)

# 写入
test_file = test_dir / "test.txt"
test_file.write_text("Hello, Python!\n第二行内容\n第三行", encoding="utf-8")

# 读取
content = test_file.read_text(encoding="utf-8")
print("文件内容:")
print(content)

# 逐行读取
print("\n逐行读取:")
with open(test_file, "r", encoding="utf-8") as f:
    for i, line in enumerate(f, 1):
        print(f"  行 {i}: {line.strip()}")

# 目录遍历
print("\n目录遍历:")
for p in test_dir.glob("*.txt"):
    print(f"  {p.name}")

# pathlib 操作
print("\nPathlib 操作:")
print("  exists():", test_file.exists())
print("  is_file():", test_file.is_file())
print("  name:", test_file.name)
print("  suffix:", test_file.suffix)
print("  parent:", test_file.parent)

# ============================================
# 17. HTTP Client
# ============================================
print("\n=== 17. HTTP Client ===")

# 使用标准库 urllib (asyncio 版本使用 aiohttp)
from urllib.request import urlopen, Request
from urllib.error import URLError
import json


def http_client_demo():
    try:
        # GET 请求
        with urlopen("https://httpbin.org/get", timeout=5) as response:
            data = json.loads(response.read().decode())
            print("GET 响应状态:", response.status)
            print("GET 响应 URL:", data.get("url"))
        
        # POST 请求
        post_data = json.dumps({"name": "Alice", "age": 25}).encode()
        req = Request(
            "https://httpbin.org/post",
            data=post_data,
            headers={"Content-Type": "application/json"},
            method="POST"
        )
        with urlopen(req, timeout=5) as response:
            data = json.loads(response.read().decode())
            print("POST 响应:", data.get("json"))
    except URLError as e:
        print("HTTP 请求失败 (可能无网络):", e)
    except Exception as e:
        print("HTTP 请求失败:", e)


http_client_demo()

# ============================================
# 装饰器示例
# ============================================
print("\n=== 装饰器示例 ===")


def timer(func: Callable) -> Callable:
    """计时装饰器"""
    import time
    
    def wrapper(*args, **kwargs):
        start = time.time()
        result = func(*args, **kwargs)
        end = time.time()
        print(f"  {func.__name__} 耗时: {end - start:.4f}s")
        return result
    return wrapper


def retry(times: int = 3):
    """重试装饰器"""
    def decorator(func: Callable) -> Callable:
        def wrapper(*args, **kwargs):
            for i in range(times):
                try:
                    return func(*args, **kwargs)
                except Exception as e:
                    print(f"  重试 {i + 1}/{times}: {e}")
            raise Exception(f"重试 {times} 次后失败")
        return wrapper
    return decorator


@timer
def slow_function():
    import time
    time.sleep(0.1)
    return "done"


@retry(times=3)
def may_fail(success_on: int, current: list):
    current[0] += 1
    if current[0] < success_on:
        raise ValueError(f"尝试 {current[0]}")
    return "success"


print("计时装饰器:")
slow_function()

print("重试装饰器:")
counter = [0]
result = may_fail(3, counter)
print(f"  结果: {result}")

# ============================================
# Context Manager
# ============================================
print("\n=== Context Manager ===")

from contextlib import contextmanager


@contextmanager
def managed_resource(name: str):
    print(f"  获取资源: {name}")
    try:
        yield f"Resource<{name}>"
    finally:
        print(f"  释放资源: {name}")


with managed_resource("Database") as resource:
    print(f"  使用资源: {resource}")

# ============================================
# 完成
# ============================================
print("\n✅ 所有示例执行完成!")
