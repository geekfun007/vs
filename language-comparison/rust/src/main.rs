//! Rust 语言特性全面示例

use chrono::{DateTime, Duration, Local, NaiveDate, Utc};
use regex::Regex;
use serde::{Deserialize, Serialize};
use std::collections::{HashMap, HashSet};
use std::error::Error;
use std::fmt;
use std::fs::{self, File};
use std::io::{BufRead, BufReader, Write};
use std::path::Path;
use std::sync::{Arc, Mutex};
use std::thread;
use thiserror::Error;

// ============================================
// 1. 基本类型与变量
// ============================================
fn basic_types() {
    println!("=== 1. 基本类型与变量 ===");

    // 不可变绑定 (默认)
    let name: &str = "Alice";
    let age: i32 = 25;
    let height: f64 = 1.75;
    let is_active: bool = true;

    // 可变绑定
    let mut score: i32 = 100;
    score += 10;

    // 类型推断
    let message = "Hello"; // &str
    let number = 42; // i32

    // 常量
    const PI: f64 = 3.14159;
    const MAX_SIZE: usize = 1000;

    // 静态变量
    static APP_NAME: &str = "RustDemo";

    println!(
        "姓名: {}, 年龄: {}, 身高: {}, 活跃: {}",
        name, age, height, is_active
    );
    println!("score: {}, message: {}", score, message);
    println!("PI: {}, MAX_SIZE: {}, APP: {}", PI, MAX_SIZE, APP_NAME);

    // 类型转换
    let x: i32 = 42;
    let y: f64 = x as f64;
    let z: i64 = x as i64;
    println!("类型转换: i32({}) -> f64({}) -> i64({})", x, y, z);

    // Option 类型 (空值处理)
    let some_value: Option<i32> = Some(42);
    let no_value: Option<i32> = None;
    println!(
        "Option: some={:?}, none={:?}",
        some_value.unwrap_or(0),
        no_value.unwrap_or(-1)
    );
}

// ============================================
// 2. 数组、向量与切片
// ============================================
fn arrays_and_vectors() {
    println!("\n=== 2. 数组、向量与切片 ===");

    // 数组 (固定长度)
    let arr: [i32; 5] = [1, 2, 3, 4, 5];
    println!("数组: {:?}", arr);
    println!("数组长度: {}", arr.len());

    // Vec (动态数组)
    let mut numbers: Vec<i32> = vec![1, 2, 3, 4, 5];
    println!("原向量: {:?}", numbers);

    // 添加/删除
    numbers.push(6);
    println!("push(6): {:?}", numbers);
    numbers.pop();
    println!("pop(): {:?}", numbers);

    // 切片
    let slice = &numbers[1..3];
    println!("slice [1..3]: {:?}", slice);

    // 迭代器方法
    let mapped: Vec<i32> = numbers.iter().map(|x| x * 2).collect();
    println!("map (x2): {:?}", mapped);

    let filtered: Vec<&i32> = numbers.iter().filter(|&&x| x > 2).collect();
    println!("filter (>2): {:?}", filtered);

    let sum: i32 = numbers.iter().sum();
    println!("sum: {}", sum);

    let product: i32 = numbers.iter().product();
    println!("product: {}", product);

    // 其他常用方法
    println!("contains(&3): {}", numbers.contains(&3));
    println!("first(): {:?}", numbers.first());
    println!("last(): {:?}", numbers.last());
    println!("len(): {}", numbers.len());
    println!("is_empty(): {}", numbers.is_empty());

    // 排序
    let mut to_sort = vec![5, 2, 8, 1, 9];
    to_sort.sort();
    println!("sorted: {:?}", to_sort);

    // 反转
    to_sort.reverse();
    println!("reversed: {:?}", to_sort);
}

// ============================================
// 3. 元组
// ============================================
fn tuples_demo() {
    println!("\n=== 3. 元组 ===");

    let point: (i32, i32) = (10, 20);
    let record: (&str, i32, bool) = ("Alice", 25, true);

    // 访问
    println!("point.0: {}, point.1: {}", point.0, point.1);

    // 解构
    let (x, y) = point;
    println!("解构: x={}, y={}", x, y);

    let (name, age, active) = record;
    println!("record: name={}, age={}, active={}", name, age, active);

    // 函数返回多个值
    fn min_max(numbers: &[i32]) -> (i32, i32) {
        let min = *numbers.iter().min().unwrap_or(&0);
        let max = *numbers.iter().max().unwrap_or(&0);
        (min, max)
    }

    let nums = vec![3, 1, 4, 1, 5, 9, 2, 6];
    let (min, max) = min_max(&nums);
    println!("min_max: min={}, max={}", min, max);
}

// ============================================
// 4. HashSet
// ============================================
fn hashset_demo() {
    println!("\n=== 4. HashSet ===");

    let mut set: HashSet<i32> = HashSet::from([1, 2, 3, 3, 4]);
    println!("Set: {:?}", set);

    set.insert(5);
    set.remove(&1);
    println!("insert(5), remove(1): {:?}", set);

    println!("contains(&2): {}", set.contains(&2));
    println!("len(): {}", set.len());

    // 集合运算
    let set1: HashSet<i32> = HashSet::from([1, 2, 3]);
    let set2: HashSet<i32> = HashSet::from([2, 3, 4]);

    let union: HashSet<&i32> = set1.union(&set2).collect();
    println!("并集: {:?}", union);

    let intersection: HashSet<&i32> = set1.intersection(&set2).collect();
    println!("交集: {:?}", intersection);

    let difference: HashSet<&i32> = set1.difference(&set2).collect();
    println!("差集: {:?}", difference);

    let symmetric: HashSet<&i32> = set1.symmetric_difference(&set2).collect();
    println!("对称差集: {:?}", symmetric);
}

// ============================================
// 5. HashMap
// ============================================
fn hashmap_demo() {
    println!("\n=== 5. HashMap ===");

    let mut person: HashMap<&str, &str> = HashMap::new();
    person.insert("name", "Alice");
    person.insert("email", "alice@example.com");
    println!("HashMap: {:?}", person);

    // 访问
    println!("get('name'): {:?}", person.get("name"));
    println!("get('phone'): {:?}", person.get("phone"));

    // Entry API
    person.entry("phone").or_insert("N/A");
    println!("entry('phone'): {:?}", person);

    // 遍历
    println!("遍历:");
    for (key, value) in &person {
        println!("  {}: {}", key, value);
    }

    // 检查存在
    println!("contains_key('name'): {}", person.contains_key("name"));

    // 数值 HashMap
    let mut scores: HashMap<&str, i32> = HashMap::new();
    scores.insert("Alice", 95);
    scores.insert("Bob", 87);

    // 更新值
    *scores.entry("Alice").or_insert(0) += 5;
    println!("scores: {:?}", scores);
}

// ============================================
// 6. 字符串操作
// ============================================
fn string_operations() {
    println!("\n=== 6. 字符串操作 ===");

    let s = String::from("  Hello, World!  ");
    println!("原字符串: {:?}", s);

    println!("len(): {}", s.len());
    println!("trim(): {:?}", s.trim());
    println!("to_uppercase(): {}", s.to_uppercase());
    println!("to_lowercase(): {}", s.to_lowercase());
    println!("contains('World'): {}", s.contains("World"));
    println!("starts_with('  H'): {}", s.starts_with("  H"));
    println!("ends_with('  '): {}", s.ends_with("  "));
    println!("replace(): {}", s.replace("World", "Rust"));

    // split
    let parts: Vec<&str> = s.trim().split(", ").collect();
    println!("split(', '): {:?}", parts);

    // chars
    let chars: Vec<char> = "你好".chars().collect();
    println!("chars('你好'): {:?}", chars);
    println!("chars count: {}", "你好世界".chars().count());

    // 格式化
    let name = "Alice";
    let age = 25;
    let greeting = format!("你好, {}! 你今年 {} 岁。", name, age);
    println!("format!: {}", greeting);

    // String vs &str
    let s1: &str = "static string"; // 字符串切片 (不可变)
    let s2: String = String::from("owned string"); // 拥有的字符串
    let s3: String = s1.to_string(); // 转换
    let s4: &str = &s2; // 借用
    println!("&str: {}, String: {}, to_string: {}, borrow: {}", s1, s2, s3, s4);
}

// ============================================
// 7. 结构体与方法
// ============================================
#[derive(Debug, Clone)]
struct Person {
    name: String,
    age: i32,
}

impl Person {
    // 关联函数 (构造器)
    fn new(name: String, age: i32) -> Self {
        Person { name, age }
    }

    // 方法 (&self - 不可变借用)
    fn greet(&self) -> String {
        format!("Hello, I'm {}", self.name)
    }

    // 方法 (&mut self - 可变借用)
    fn set_age(&mut self, age: i32) {
        self.age = age;
    }

    fn info(&self) -> String {
        format!("{}, {} years old", self.name, self.age)
    }
}

// "继承" - 组合
#[derive(Debug)]
struct Student {
    person: Person,
    school: String,
}

impl Student {
    fn new(name: String, age: i32, school: String) -> Self {
        Student {
            person: Person::new(name, age),
            school,
        }
    }

    fn greet(&self) -> String {
        format!("{}, I study at {}", self.person.greet(), self.school)
    }
}

fn structs_demo() {
    println!("\n=== 7. 结构体与方法 ===");

    let mut person = Person::new("Alice".to_string(), 25);
    let student = Student::new("Bob".to_string(), 20, "MIT".to_string());

    println!("Person: {}", person.greet());
    println!("Student: {}", student.greet());
    println!("Info: {}", person.info());

    person.set_age(26);
    println!("修改后年龄: {}", person.age);

    // Debug trait
    println!("Debug: {:?}", person);
}

// ============================================
// 8. Trait (接口)
// ============================================
trait Greeter {
    fn greet(&self) -> String;

    // 默认实现
    fn greet_loud(&self) -> String {
        self.greet().to_uppercase()
    }
}

trait Named {
    fn name(&self) -> &str;
}

// 为 Person 实现 Greeter
impl Greeter for Person {
    fn greet(&self) -> String {
        format!("Hello, I'm {}", self.name)
    }
}

// Trait 约束函数
fn print_greeting<T: Greeter>(item: &T) {
    println!("Greeting: {}", item.greet());
}

// Trait 对象 (动态分发)
fn print_greeting_dyn(item: &dyn Greeter) {
    println!("Dyn Greeting: {}", item.greet());
}

fn traits_demo() {
    println!("\n=== 8. Trait (接口) ===");

    let person = Person::new("Alice".to_string(), 25);

    print_greeting(&person);
    print_greeting_dyn(&person);
    println!("greet_loud: {}", person.greet_loud());
}

// ============================================
// 9. 枚举 (代数数据类型)
// ============================================
#[derive(Debug)]
enum Direction {
    Up,
    Down,
    Left,
    Right,
}

#[derive(Debug)]
enum HttpStatus {
    Ok = 200,
    NotFound = 404,
    InternalError = 500,
}

// 带数据的枚举 (ADT)
#[derive(Debug)]
enum Message {
    Quit,
    Move { x: i32, y: i32 },
    Write(String),
    ChangeColor(i32, i32, i32),
}

fn handle_message(msg: &Message) {
    match msg {
        Message::Quit => println!("  Quit"),
        Message::Move { x, y } => println!("  Move to ({}, {})", x, y),
        Message::Write(text) => println!("  Write: {}", text),
        Message::ChangeColor(r, g, b) => println!("  Color: {} {} {}", r, g, b),
    }
}

fn enums_demo() {
    println!("\n=== 9. 枚举 ===");

    println!("Direction::Up: {:?}", Direction::Up);
    println!("HttpStatus::Ok: {:?} = {}", HttpStatus::Ok, HttpStatus::Ok as i32);

    let messages = vec![
        Message::Quit,
        Message::Move { x: 10, y: 20 },
        Message::Write("Hello".to_string()),
        Message::ChangeColor(255, 128, 0),
    ];

    println!("消息处理:");
    for msg in &messages {
        handle_message(msg);
    }

    // Option 和 Result
    let some: Option<i32> = Some(42);
    let none: Option<i32> = None;

    // if let
    if let Some(value) = some {
        println!("Some value: {}", value);
    }

    // match
    match none {
        Some(v) => println!("Value: {}", v),
        None => println!("No value"),
    }
}

// ============================================
// 10. 错误处理
// ============================================
// 使用 thiserror
#[derive(Error, Debug)]
enum AppError {
    #[error("Division by zero")]
    DivisionByZero,

    #[error("Validation error [{field}]: {message}")]
    Validation { field: String, message: String },

    #[error("IO error: {0}")]
    Io(#[from] std::io::Error),
}

fn divide(a: f64, b: f64) -> Result<f64, AppError> {
    if b == 0.0 {
        Err(AppError::DivisionByZero)
    } else {
        Ok(a / b)
    }
}

fn validate_age(age: i32) -> Result<(), AppError> {
    if age < 0 {
        Err(AppError::Validation {
            field: "age".to_string(),
            message: "cannot be negative".to_string(),
        })
    } else {
        Ok(())
    }
}

fn error_handling_demo() {
    println!("\n=== 10. 错误处理 ===");

    // match 处理
    match divide(10.0, 2.0) {
        Ok(result) => println!("10 / 2 = {}", result),
        Err(e) => println!("Error: {}", e),
    }

    match divide(10.0, 0.0) {
        Ok(result) => println!("10 / 0 = {}", result),
        Err(e) => println!("Error: {}", e),
    }

    // if let
    if let Err(e) = validate_age(-5) {
        println!("Validation Error: {}", e);
    }

    // ? 操作符 (错误传播)
    fn calculate() -> Result<f64, AppError> {
        let x = divide(10.0, 2.0)?;
        let y = divide(x, 2.0)?;
        Ok(y)
    }

    println!("calculate(): {:?}", calculate());

    // unwrap_or
    let result = divide(10.0, 0.0).unwrap_or(-1.0);
    println!("unwrap_or: {}", result);
}

// ============================================
// 11. 日期时间
// ============================================
fn datetime_demo() {
    println!("\n=== 11. 日期时间 ===");

    let now: DateTime<Local> = Local::now();
    let utc_now: DateTime<Utc> = Utc::now();

    println!("当前时间 (Local): {}", now.format("%Y-%m-%d %H:%M:%S"));
    println!("当前时间 (UTC): {}", utc_now.format("%Y-%m-%d %H:%M:%S"));
    println!("年: {}", now.format("%Y"));
    println!("月: {}", now.format("%m"));
    println!("日: {}", now.format("%d"));
    println!("时间戳: {}", now.timestamp());

    // 解析
    let parsed = NaiveDate::parse_from_str("2024-01-15", "%Y-%m-%d");
    println!("解析: {:?}", parsed);

    // 计算
    let tomorrow = now + Duration::days(1);
    println!("明天: {}", tomorrow.format("%Y-%m-%d"));

    if let (Some(start), Some(end)) = (
        NaiveDate::from_ymd_opt(2024, 1, 1),
        NaiveDate::from_ymd_opt(2024, 12, 31),
    ) {
        let diff = end.signed_duration_since(start);
        println!("日期差: {} 天", diff.num_days());
    }

    // 计时
    let start = std::time::Instant::now();
    std::thread::sleep(std::time::Duration::from_millis(100));
    let elapsed = start.elapsed();
    println!("耗时: {:?}", elapsed);
}

// ============================================
// 12. 正则表达式
// ============================================
fn regex_demo() {
    println!("\n=== 12. 正则表达式 ===");

    let text = "Call 123-4567 or 987-6543";
    let pattern = Regex::new(r"\d{3}-\d{4}").unwrap();

    println!("文本: {}", text);
    println!("is_match(): {}", pattern.is_match(text));

    // find
    if let Some(m) = pattern.find(text) {
        println!("find(): {} at {}..{}", m.as_str(), m.start(), m.end());
    }

    // find_iter
    let matches: Vec<&str> = pattern.find_iter(text).map(|m| m.as_str()).collect();
    println!("find_iter(): {:?}", matches);

    // replace
    let replaced = pattern.replace_all(text, "XXX-XXXX");
    println!("replace_all(): {}", replaced);

    // 捕获组
    let email_pattern = Regex::new(r"(\w+)@(\w+\.\w+)").unwrap();
    if let Some(caps) = email_pattern.captures("test@example.com") {
        println!("捕获组:");
        println!("  完整: {}", &caps[0]);
        println!("  用户: {}", &caps[1]);
        println!("  域名: {}", &caps[2]);
    }

    // 命名捕获组
    let named_pattern = Regex::new(r"(?P<user>\w+)@(?P<domain>\w+\.\w+)").unwrap();
    if let Some(caps) = named_pattern.captures("test@example.com") {
        println!("命名捕获组:");
        println!("  user: {}", &caps["user"]);
        println!("  domain: {}", &caps["domain"]);
    }
}

// ============================================
// 13. 迭代器
// ============================================
fn iterators_demo() {
    println!("\n=== 13. 迭代器 ===");

    let numbers = vec![1, 2, 3, 4, 5];

    // 链式操作
    let result: i32 = numbers
        .iter()
        .filter(|&&x| x % 2 == 0) // 偶数
        .map(|x| x * 10) // * 10
        .sum(); // 求和
    println!("链式操作 (偶数 * 10 求和): {}", result);

    // 更多迭代器方法
    let nums = vec![1, 2, 3, 4, 5];
    println!("take(3): {:?}", nums.iter().take(3).collect::<Vec<_>>());
    println!("skip(2): {:?}", nums.iter().skip(2).collect::<Vec<_>>());
    println!(
        "enumerate: {:?}",
        nums.iter().enumerate().collect::<Vec<_>>()
    );
    println!(
        "zip: {:?}",
        nums.iter().zip(vec![10, 20, 30]).collect::<Vec<_>>()
    );
    println!("all(> 0): {}", nums.iter().all(|&x| x > 0));
    println!("any(> 3): {}", nums.iter().any(|&x| x > 3));
    println!("max: {:?}", nums.iter().max());
    println!("min: {:?}", nums.iter().min());
    println!("count: {}", nums.iter().count());

    // fold (reduce)
    let sum = nums.iter().fold(0, |acc, x| acc + x);
    println!("fold (sum): {}", sum);

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

    let counter = Counter { count: 0, max: 5 };
    println!("自定义迭代器: {:?}", counter.collect::<Vec<_>>());

    // 范围迭代器
    println!("range (0..5): {:?}", (0..5).collect::<Vec<_>>());
    println!("range inclusive (0..=5): {:?}", (0..=5).collect::<Vec<_>>());
}

// ============================================
// 14. 并发
// ============================================
fn concurrency_demo() {
    println!("\n=== 14. 并发 ===");

    // 基本线程
    println!("基本线程:");
    let handle = thread::spawn(|| {
        thread::sleep(std::time::Duration::from_millis(50));
        42
    });
    let result = handle.join().unwrap();
    println!("  线程返回: {}", result);

    // 多线程
    println!("多线程:");
    let mut handles = vec![];
    for i in 0..3 {
        let handle = thread::spawn(move || {
            thread::sleep(std::time::Duration::from_millis(50));
            println!("    线程 {} 完成", i);
        });
        handles.push(handle);
    }
    for handle in handles {
        handle.join().unwrap();
    }

    // 共享状态 (Arc + Mutex)
    println!("共享状态 (Arc + Mutex):");
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
    println!("  计数器: {}", *counter.lock().unwrap());

    // Channel
    println!("Channel:");
    use std::sync::mpsc;
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        for i in 0..3 {
            tx.send(format!("消息 {}", i)).unwrap();
            thread::sleep(std::time::Duration::from_millis(30));
        }
    });

    for received in rx {
        println!("  收到: {}", received);
    }

    // Scoped threads (借用安全)
    println!("Scoped threads:");
    let data = vec![1, 2, 3];
    thread::scope(|s| {
        s.spawn(|| {
            println!("  scoped thread: {:?}", data);
        });
    });
}

// ============================================
// 15. 文件 IO
// ============================================
fn file_io_demo() {
    println!("\n=== 15. 文件 IO ===");

    let test_dir = Path::new("./test_output");
    fs::create_dir_all(test_dir).unwrap();

    let test_file = test_dir.join("test.txt");

    // 写入
    fs::write(&test_file, "Hello, Rust!\n第二行内容\n第三行").unwrap();

    // 读取整个文件
    let content = fs::read_to_string(&test_file).unwrap();
    println!("文件内容:");
    println!("{}", content);

    // 逐行读取
    println!("\n逐行读取:");
    let file = File::open(&test_file).unwrap();
    let reader = BufReader::new(file);
    for (i, line) in reader.lines().enumerate() {
        println!("  行 {}: {}", i + 1, line.unwrap());
    }

    // 追加写入
    let mut file = fs::OpenOptions::new()
        .append(true)
        .open(&test_file)
        .unwrap();
    writeln!(file, "第四行 (追加)").unwrap();

    // 目录遍历
    println!("\n目录遍历:");
    for entry in fs::read_dir(test_dir).unwrap() {
        let entry = entry.unwrap();
        println!("  {}", entry.file_name().to_string_lossy());
    }

    // 文件元数据
    println!("\n文件信息:");
    let metadata = fs::metadata(&test_file).unwrap();
    println!("  大小: {} bytes", metadata.len());
    println!("  是文件: {}", metadata.is_file());
    println!("  是目录: {}", metadata.is_dir());
}

// ============================================
// 16. 所有权与借用
// ============================================
fn ownership_demo() {
    println!("\n=== 16. 所有权与借用 ===");

    // 所有权转移 (Move)
    let s1 = String::from("hello");
    let s2 = s1; // s1 不再有效
    // println!("{}", s1); // 错误!
    println!("Move: s2 = {}", s2);

    // Clone (深拷贝)
    let s3 = s2.clone();
    println!("Clone: s2 = {}, s3 = {}", s2, s3);

    // Copy (栈上数据)
    let x = 5;
    let y = x; // i32 实现了 Copy
    println!("Copy: x = {}, y = {}", x, y);

    // 借用 (&)
    fn calculate_length(s: &String) -> usize {
        s.len() // 不获取所有权
    }
    let len = calculate_length(&s2);
    println!("借用: len = {}, s2 still valid = {}", len, s2);

    // 可变借用 (&mut)
    fn append_world(s: &mut String) {
        s.push_str(", world!");
    }
    let mut s4 = String::from("hello");
    append_world(&mut s4);
    println!("可变借用: {}", s4);

    // 生命周期示例
    fn longest<'a>(x: &'a str, y: &'a str) -> &'a str {
        if x.len() > y.len() {
            x
        } else {
            y
        }
    }
    let result = longest("hello", "world!");
    println!("生命周期: longest = {}", result);
}

// ============================================
// main
// ============================================
fn main() {
    basic_types();
    arrays_and_vectors();
    tuples_demo();
    hashset_demo();
    hashmap_demo();
    string_operations();
    structs_demo();
    traits_demo();
    enums_demo();
    error_handling_demo();
    datetime_demo();
    regex_demo();
    iterators_demo();
    concurrency_demo();
    file_io_demo();
    ownership_demo();

    println!("\n✅ 所有示例执行完成!");
}
