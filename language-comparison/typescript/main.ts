/**
 * TypeScript 语言特性全面示例
 */

// ============================================
// 1. 基本类型与变量
// ============================================
console.log("=== 1. 基本类型与变量 ===");

// 变量声明
let name: string = "Alice";
const age: number = 25;
const isActive: boolean = true;
let score = 100; // 类型推断

// 特殊类型
let nothing: null = null;
let notDefined: undefined = undefined;
let anything: any = "可以是任何类型";
let unknown1: unknown = 42; // 更安全的 any

console.log(`姓名: ${name}, 年龄: ${age}, 活跃: ${isActive}`);

// ============================================
// 2. 数组与元组
// ============================================
console.log("\n=== 2. 数组与元组 ===");

// 数组
const numbers: number[] = [1, 2, 3, 4, 5];
const mixed: Array<number | string> = [1, "two", 3];

// 数组方法
console.log("原数组:", numbers);
console.log("map (x2):", numbers.map(x => x * 2));
console.log("filter (>2):", numbers.filter(x => x > 2));
console.log("reduce (sum):", numbers.reduce((a, b) => a + b, 0));
console.log("find (>3):", numbers.find(x => x > 3));
console.log("includes(3):", numbers.includes(3));
console.log("slice(1,3):", numbers.slice(1, 3));

// 元组
const point: [number, number] = [10, 20];
const record: [string, number, boolean] = ["Alice", 25, true];
const [x, y] = point; // 解构
console.log(`Point: (${x}, ${y})`);

// ============================================
// 3. Set 与 Map
// ============================================
console.log("\n=== 3. Set 与 Map ===");

// Set
const set = new Set<number>([1, 2, 3, 3, 4]);
set.add(5);
set.delete(1);
console.log("Set:", [...set]);
console.log("has(2):", set.has(2));
console.log("size:", set.size);

// 集合运算
const set1 = new Set([1, 2, 3]);
const set2 = new Set([2, 3, 4]);
const union = new Set([...set1, ...set2]);
const intersection = new Set([...set1].filter(x => set2.has(x)));
const difference = new Set([...set1].filter(x => !set2.has(x)));
console.log("并集:", [...union]);
console.log("交集:", [...intersection]);
console.log("差集:", [...difference]);

// Map
const map = new Map<string, number>();
map.set("a", 1);
map.set("b", 2);
map.set("c", 3);
console.log("Map:", Object.fromEntries(map));
console.log("get('a'):", map.get("a"));
console.log("has('b'):", map.has("b"));

// Object 作为字典
const dict: Record<string, number> = { apple: 1, banana: 2 };
console.log("keys:", Object.keys(dict));
console.log("values:", Object.values(dict));
console.log("entries:", Object.entries(dict));

// ============================================
// 4. 字符串操作
// ============================================
console.log("\n=== 4. 字符串操作 ===");

const str = "  Hello, World!  ";
console.log("原字符串:", JSON.stringify(str));
console.log("length:", str.length);
console.log("trim():", JSON.stringify(str.trim()));
console.log("toUpperCase():", str.toUpperCase());
console.log("toLowerCase():", str.toLowerCase());
console.log("split(','):", str.trim().split(", "));
console.log("includes('World'):", str.includes("World"));
console.log("startsWith('  H'):", str.startsWith("  H"));
console.log("replace:", str.replace("World", "TypeScript"));
console.log("slice(2, 7):", str.slice(2, 7));

// 模板字符串
const greeting = `你好, ${name}! 你今年 ${age} 岁。`;
console.log("模板字符串:", greeting);

// ============================================
// 5. 类与接口
// ============================================
console.log("\n=== 5. 类与接口 ===");

// 接口
interface Greeter {
    greet(): string;
}

interface Named {
    readonly name: string;
}

// 类
class Person implements Greeter, Named {
    readonly name: string;
    private _age: number;
    
    constructor(name: string, age: number) {
        this.name = name;
        this._age = age;
    }
    
    greet(): string {
        return `Hello, I'm ${this.name}`;
    }
    
    get age(): number {
        return this._age;
    }
    
    set age(value: number) {
        if (value >= 0) this._age = value;
    }
    
    get info(): string {
        return `${this.name}, ${this._age} years old`;
    }
    
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
    
    greet(): string {
        return `${super.greet()}, I study at ${this.school}`;
    }
}

const person = new Person("Alice", 25);
const student = new Student("Bob", 20, "MIT");
console.log("Person:", person.greet());
console.log("Student:", student.greet());
console.log("Info:", person.info);

// ============================================
// 6. 枚举
// ============================================
console.log("\n=== 6. 枚举 ===");

// 数字枚举
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

// 字符串枚举
enum Color {
    Red = "RED",
    Green = "GREEN",
    Blue = "BLUE",
}

console.log("Direction.Up:", Direction.Up);
console.log("HttpStatus.Ok:", HttpStatus.Ok);
console.log("Color.Red:", Color.Red);

// ============================================
// 7. 泛型
// ============================================
console.log("\n=== 7. 泛型 ===");

// 泛型函数
function identity<T>(arg: T): T {
    return arg;
}

// 泛型类
class Container<T> {
    private value: T;
    
    constructor(value: T) {
        this.value = value;
    }
    
    getValue(): T {
        return this.value;
    }
}

// 泛型约束
interface Lengthwise {
    length: number;
}

function logLength<T extends Lengthwise>(arg: T): number {
    console.log("Length:", arg.length);
    return arg.length;
}

console.log("identity(42):", identity(42));
console.log("identity('hello'):", identity("hello"));
const container = new Container<number>(100);
console.log("Container value:", container.getValue());
logLength("hello");
logLength([1, 2, 3]);

// ============================================
// 8. 错误处理
// ============================================
console.log("\n=== 8. 错误处理 ===");

// 自定义错误
class ValidationError extends Error {
    constructor(public field: string, message: string) {
        super(message);
        this.name = "ValidationError";
    }
}

function divide(a: number, b: number): number {
    if (b === 0) {
        throw new Error("Division by zero");
    }
    return a / b;
}

function validateAge(age: number): void {
    if (age < 0) {
        throw new ValidationError("age", "Age cannot be negative");
    }
}

try {
    console.log("10 / 2 =", divide(10, 2));
    console.log("10 / 0 =", divide(10, 0));
} catch (error) {
    if (error instanceof Error) {
        console.log("Error:", error.message);
    }
}

try {
    validateAge(-5);
} catch (error) {
    if (error instanceof ValidationError) {
        console.log(`ValidationError [${error.field}]:`, error.message);
    }
}

// ============================================
// 9. 日期时间
// ============================================
console.log("\n=== 9. 日期时间 ===");

import dayjs from 'dayjs';

const now = new Date();
console.log("当前时间:", now.toISOString());
console.log("年:", now.getFullYear());
console.log("月:", now.getMonth() + 1);
console.log("日:", now.getDate());
console.log("时间戳:", now.getTime());

// dayjs
console.log("dayjs 格式化:", dayjs().format("YYYY-MM-DD HH:mm:ss"));
console.log("明天:", dayjs().add(1, "day").format("YYYY-MM-DD"));
console.log("日期差:", dayjs("2024-12-31").diff(dayjs("2024-01-01"), "day"), "天");

// ============================================
// 10. 正则表达式
// ============================================
console.log("\n=== 10. 正则表达式 ===");

const text = "Call 123-4567 or 987-6543";
const pattern = /\d{3}-\d{4}/g;

console.log("文本:", text);
console.log("test():", pattern.test(text));
pattern.lastIndex = 0; // 重置
console.log("match():", text.match(pattern));
console.log("replace():", text.replace(pattern, "XXX-XXXX"));

// 捕获组
const emailPattern = /(?<user>\w+)@(?<domain>\w+\.\w+)/;
const emailMatch = "test@example.com".match(emailPattern);
console.log("邮箱匹配:", emailMatch?.groups);

// ============================================
// 11. 迭代器与生成器
// ============================================
console.log("\n=== 11. 迭代器与生成器 ===");

// Generator
function* range(start: number, end: number): Generator<number> {
    for (let i = start; i < end; i++) {
        yield i;
    }
}

console.log("Generator range(0, 5):", [...range(0, 5)]);

// 自定义迭代器
const iterable = {
    data: [1, 2, 3],
    [Symbol.iterator]() {
        let index = 0;
        const data = this.data;
        return {
            next() {
                if (index < data.length) {
                    return { value: data[index++], done: false };
                }
                return { value: undefined, done: true };
            }
        };
    }
};

console.log("自定义迭代器:", [...iterable]);

// 链式操作
const chainResult = [1, 2, 3, 4, 5]
    .filter(x => x % 2 === 0)
    .map(x => x * 10)
    .reduce((a, b) => a + b, 0);
console.log("链式操作结果:", chainResult);

// ============================================
// 12. 异步编程
// ============================================
console.log("\n=== 12. 异步编程 ===");

// Promise
function delay(ms: number): Promise<void> {
    return new Promise(resolve => setTimeout(resolve, ms));
}

async function fetchData(id: number): Promise<{ id: number; data: string }> {
    await delay(100);
    return { id, data: `Data for ${id}` };
}

// async/await
async function asyncDemo(): Promise<void> {
    console.log("开始异步操作...");
    
    // 顺序执行
    const result1 = await fetchData(1);
    console.log("顺序执行:", result1);
    
    // 并行执行
    const [r1, r2, r3] = await Promise.all([
        fetchData(1),
        fetchData(2),
        fetchData(3),
    ]);
    console.log("并行执行:", [r1, r2, r3]);
    
    // Promise.race
    const fastest = await Promise.race([
        delay(100).then(() => "slow"),
        delay(50).then(() => "fast"),
    ]);
    console.log("竞争结果:", fastest);
}

await asyncDemo();

// ============================================
// 13. 文件 IO
// ============================================
console.log("\n=== 13. 文件 IO ===");

import { readFileSync, writeFileSync, existsSync, mkdirSync } from 'fs';
import { readFile, writeFile, mkdir } from 'fs/promises';
import { join, dirname } from 'path';

const testDir = join(dirname(new URL(import.meta.url).pathname), 'test_output');
const testFile = join(testDir, 'test.txt');

// 确保目录存在
if (!existsSync(testDir)) {
    mkdirSync(testDir, { recursive: true });
}

// 同步写入
writeFileSync(testFile, 'Hello, TypeScript!\n第二行内容\n第三行');

// 同步读取
const content = readFileSync(testFile, 'utf-8');
console.log("文件内容:");
console.log(content);

// 异步操作
await writeFile(join(testDir, 'async.txt'), 'Async content');
const asyncContent = await readFile(join(testDir, 'async.txt'), 'utf-8');
console.log("异步读取:", asyncContent);

// ============================================
// 14. HTTP Client
// ============================================
console.log("\n=== 14. HTTP Client ===");

// 使用 fetch (Node.js 18+)
async function httpClientDemo(): Promise<void> {
    try {
        // GET 请求
        const response = await fetch('https://httpbin.org/get');
        const data = await response.json();
        console.log("GET 响应状态:", response.status);
        console.log("GET 响应 URL:", data.url);
        
        // POST 请求
        const postResponse = await fetch('https://httpbin.org/post', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ name: 'Alice', age: 25 }),
        });
        const postData = await postResponse.json();
        console.log("POST 响应:", postData.json);
    } catch (error) {
        console.log("HTTP 请求失败 (可能无网络):", (error as Error).message);
    }
}

await httpClientDemo();

// ============================================
// 类型工具
// ============================================
console.log("\n=== 类型工具 ===");

// Partial - 所有属性可选
interface User {
    id: number;
    name: string;
    email: string;
}

type PartialUser = Partial<User>;

// Pick - 选择属性
type UserName = Pick<User, 'name'>;

// Omit - 排除属性
type UserWithoutEmail = Omit<User, 'email'>;

// Record - 创建对象类型
type StringMap = Record<string, number>;

// Union 与交叉类型
type A = { a: number };
type B = { b: string };
type AB = A & B; // 交叉
type AorB = A | B; // 联合

// 条件类型
type NonNullable<T> = T extends null | undefined ? never : T;

console.log("TypeScript 类型系统演示完成!");

// ============================================
// 完成
// ============================================
console.log("\n✅ 所有示例执行完成!");
