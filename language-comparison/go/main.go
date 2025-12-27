// Go 语言特性全面示例
package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

// ============================================
// 1. 基本类型与变量
// ============================================
func basicTypes() {
	fmt.Println("=== 1. 基本类型与变量 ===")

	// 显式声明
	var name string = "Alice"
	var age int = 25
	var height float64 = 1.75
	var isActive bool = true

	// 短变量声明 (类型推断)
	score := 100
	message := "Hello"

	// 常量
	const PI = 3.14159
	const (
		StatusOK    = 200
		StatusError = 500
	)

	// 零值
	var emptyString string // ""
	var emptyInt int       // 0
	var emptyBool bool     // false

	fmt.Printf("姓名: %s, 年龄: %d, 身高: %.2f, 活跃: %t\n", name, age, height, isActive)
	fmt.Printf("score: %d, message: %s\n", score, message)
	fmt.Printf("零值: string=%q, int=%d, bool=%t\n", emptyString, emptyInt, emptyBool)
}

// ============================================
// 2. 数组与切片
// ============================================
func arraysAndSlices() {
	fmt.Println("\n=== 2. 数组与切片 ===")

	// 数组 (固定长度)
	var arr [3]int = [3]int{1, 2, 3}
	fmt.Println("数组:", arr)

	// 切片 (动态长度)
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("原切片:", numbers)

	// 添加元素
	numbers = append(numbers, 6)
	fmt.Println("append(6):", numbers)

	// 切片操作
	fmt.Println("slice [1:3]:", numbers[1:3])
	fmt.Println("slice [:3]:", numbers[:3])
	fmt.Println("slice [3:]:", numbers[3:])

	// 长度和容量
	fmt.Printf("len: %d, cap: %d\n", len(numbers), cap(numbers))

	// make 创建切片
	s := make([]int, 3, 10) // len=3, cap=10
	fmt.Printf("make slice: %v, len=%d, cap=%d\n", s, len(s), cap(s))

	// 复制
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	copy(dst, src)
	fmt.Println("copy:", dst)

	// 手动实现 map/filter/reduce
	mapped := Map(numbers, func(x int) int { return x * 2 })
	fmt.Println("map (x2):", mapped)

	filtered := Filter(numbers, func(x int) bool { return x > 3 })
	fmt.Println("filter (>3):", filtered)

	sum := Reduce(numbers, 0, func(acc, x int) int { return acc + x })
	fmt.Println("reduce (sum):", sum)

	// 排序
	toSort := []int{5, 2, 8, 1, 9}
	sort.Ints(toSort)
	fmt.Println("sorted:", toSort)

	// 查找
	fmt.Println("contains(3):", Contains(numbers, 3))
}

// 泛型辅助函数 (Go 1.18+)
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

func Contains[T comparable](slice []T, target T) bool {
	for _, v := range slice {
		if v == target {
			return true
		}
	}
	return false
}

// ============================================
// 3. Map
// ============================================
func mapsDemo() {
	fmt.Println("\n=== 3. Map ===")

	// 创建 map
	person := map[string]interface{}{
		"name":  "Alice",
		"age":   25,
		"email": "alice@example.com",
	}
	fmt.Println("map:", person)

	// 操作
	person["phone"] = "123-4567"
	fmt.Println("添加 phone:", person)

	delete(person, "phone")
	fmt.Println("删除 phone:", person)

	// 检查存在
	if value, ok := person["name"]; ok {
		fmt.Println("name exists:", value)
	}

	// 遍历
	fmt.Println("遍历:")
	for key, value := range person {
		fmt.Printf("  %s: %v\n", key, value)
	}

	// make 创建
	scores := make(map[string]int)
	scores["Alice"] = 95
	scores["Bob"] = 87
	fmt.Println("scores:", scores)

	// 用 map 模拟 Set
	set := make(map[int]struct{})
	set[1] = struct{}{}
	set[2] = struct{}{}
	set[3] = struct{}{}

	// 检查 Set 成员
	if _, exists := set[2]; exists {
		fmt.Println("set contains 2")
	}
}

// ============================================
// 4. 字符串操作
// ============================================
func stringOperations() {
	fmt.Println("\n=== 4. 字符串操作 ===")

	s := "  Hello, World!  "
	fmt.Println("原字符串:", fmt.Sprintf("%q", s))
	fmt.Println("len():", len(s))
	fmt.Println("TrimSpace():", fmt.Sprintf("%q", strings.TrimSpace(s)))
	fmt.Println("ToUpper():", strings.ToUpper(s))
	fmt.Println("ToLower():", strings.ToLower(s))
	fmt.Println("Split(','):", strings.Split(strings.TrimSpace(s), ", "))
	fmt.Println("Contains('World'):", strings.Contains(s, "World"))
	fmt.Println("HasPrefix('  H'):", strings.HasPrefix(s, "  H"))
	fmt.Println("Replace():", strings.Replace(s, "World", "Go", 1))
	fmt.Println("slice [2:7]:", s[2:7])

	// 字符串构建 (高效)
	var builder strings.Builder
	builder.WriteString("Hello")
	builder.WriteString(", ")
	builder.WriteString("Go!")
	fmt.Println("Builder:", builder.String())

	// 格式化
	name := "Alice"
	age := 25
	greeting := fmt.Sprintf("你好, %s! 你今年 %d 岁。", name, age)
	fmt.Println("格式化:", greeting)

	// rune (Unicode 字符)
	chinese := "你好世界"
	fmt.Println("字节长度:", len(chinese))
	fmt.Println("字符长度:", len([]rune(chinese)))

	for i, r := range chinese {
		fmt.Printf("  [%d] %c (U+%04X)\n", i, r, r)
	}
}

// ============================================
// 5. 结构体与方法
// ============================================
type Person struct {
	name string // 小写 = 私有
	Age  int    // 大写 = 公开
}

// 方法 (值接收者)
func (p Person) Greet() string {
	return fmt.Sprintf("Hello, I'm %s", p.name)
}

// 方法 (指针接收者 - 可修改)
func (p *Person) SetAge(age int) {
	p.Age = age
}

func (p Person) Info() string {
	return fmt.Sprintf("%s, %d years old", p.name, p.Age)
}

// 构造函数 (约定)
func NewPerson(name string, age int) *Person {
	return &Person{name: name, Age: age}
}

// "继承" - 嵌入
type Student struct {
	Person // 嵌入
	School string
}

func NewStudent(name string, age int, school string) *Student {
	return &Student{
		Person: Person{name: name, Age: age},
		School: school,
	}
}

func (s Student) Greet() string {
	return fmt.Sprintf("%s, I study at %s", s.Person.Greet(), s.School)
}

func structsDemo() {
	fmt.Println("\n=== 5. 结构体与方法 ===")

	person := NewPerson("Alice", 25)
	student := NewStudent("Bob", 20, "MIT")

	fmt.Println("Person:", person.Greet())
	fmt.Println("Student:", student.Greet())
	fmt.Println("Info:", person.Info())

	// 修改
	person.SetAge(26)
	fmt.Println("修改后年龄:", person.Age)
}

// ============================================
// 6. 接口
// ============================================
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

func printGreeting(g Greeter) {
	fmt.Println("Greeting:", g.Greet())
}

// 空接口 (任意类型)
func printAny(v any) {
	fmt.Printf("  Type: %T, Value: %v\n", v, v)
}

func interfacesDemo() {
	fmt.Println("\n=== 6. 接口 ===")

	person := NewPerson("Alice", 25)
	student := NewStudent("Bob", 20, "MIT")

	// Person 和 Student 都实现了 Greeter (隐式)
	printGreeting(person)
	printGreeting(student)

	// 空接口
	fmt.Println("空接口:")
	printAny(42)
	printAny("hello")
	printAny(true)
	printAny([]int{1, 2, 3})

	// 类型断言
	var g Greeter = person
	if p, ok := g.(*Person); ok {
		fmt.Println("类型断言成功:", p.Info())
	}

	// 类型 switch
	values := []any{42, "hello", true, 3.14}
	fmt.Println("类型 switch:")
	for _, v := range values {
		switch val := v.(type) {
		case int:
			fmt.Printf("  int: %d\n", val)
		case string:
			fmt.Printf("  string: %s\n", val)
		case bool:
			fmt.Printf("  bool: %t\n", val)
		default:
			fmt.Printf("  unknown: %v\n", val)
		}
	}
}

// ============================================
// 7. 枚举 (iota)
// ============================================
type Direction int

const (
	Up Direction = iota // 0
	Down                // 1
	Left                // 2
	Right               // 3
)

func (d Direction) String() string {
	return [...]string{"Up", "Down", "Left", "Right"}[d]
}

type HttpStatus int

const (
	StatusOK            HttpStatus = 200
	StatusNotFound      HttpStatus = 404
	StatusInternalError HttpStatus = 500
)

func enumsDemo() {
	fmt.Println("\n=== 7. 枚举 (iota) ===")

	fmt.Println("Direction.Up:", Up)
	fmt.Println("Direction.Down:", Down)
	fmt.Println("HttpStatus.OK:", StatusOK)
	fmt.Println("HttpStatus.NotFound:", StatusNotFound)
}

// ============================================
// 8. 错误处理
// ============================================
// 自定义错误
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func validateAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Message: "cannot be negative"}
	}
	return nil
}

func errorHandlingDemo() {
	fmt.Println("\n=== 8. 错误处理 ===")

	// 基本错误处理
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	_, err = divide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// 自定义错误
	err = validateAge(-5)
	if err != nil {
		var valErr *ValidationError
		if errors.As(err, &valErr) {
			fmt.Printf("ValidationError [%s]: %s\n", valErr.Field, valErr.Message)
		}
	}

	// 错误包装
	_, err = divide(10, 0)
	if err != nil {
		wrappedErr := fmt.Errorf("计算失败: %w", err)
		fmt.Println("Wrapped error:", wrappedErr)

		// 解包
		if errors.Is(wrappedErr, err) {
			fmt.Println("错误匹配")
		}
	}
}

// ============================================
// 9. 日期时间
// ============================================
func dateTimeDemo() {
	fmt.Println("\n=== 9. 日期时间 ===")

	now := time.Now()
	fmt.Println("当前时间:", now.Format(time.RFC3339))
	fmt.Println("年:", now.Year())
	fmt.Println("月:", now.Month())
	fmt.Println("日:", now.Day())
	fmt.Println("时间戳:", now.Unix())

	// Go 特有格式化 (使用参考时间 2006-01-02 15:04:05)
	fmt.Println("格式化:", now.Format("2006-01-02 15:04:05"))

	// 解析
	parsed, _ := time.Parse("2006-01-02", "2024-01-15")
	fmt.Println("解析:", parsed.Format("2006-01-02"))

	// 计算
	tomorrow := now.Add(24 * time.Hour)
	fmt.Println("明天:", tomorrow.Format("2006-01-02"))

	start := time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2024, 12, 31, 0, 0, 0, 0, time.UTC)
	diff := end.Sub(start)
	fmt.Printf("日期差: %.0f 天\n", diff.Hours()/24)

	// 计时器
	timer := time.Now()
	time.Sleep(100 * time.Millisecond)
	elapsed := time.Since(timer)
	fmt.Printf("耗时: %v\n", elapsed)
}

// ============================================
// 10. 正则表达式
// ============================================
func regexDemo() {
	fmt.Println("\n=== 10. 正则表达式 ===")

	text := "Call 123-4567 or 987-6543"
	pattern := regexp.MustCompile(`\d{3}-\d{4}`)

	fmt.Println("文本:", text)
	fmt.Println("MatchString():", pattern.MatchString(text))
	fmt.Println("FindString():", pattern.FindString(text))
	fmt.Println("FindAllString():", pattern.FindAllString(text, -1))
	fmt.Println("ReplaceAllString():", pattern.ReplaceAllString(text, "XXX-XXXX"))

	// 捕获组
	emailPattern := regexp.MustCompile(`(\w+)@(\w+\.\w+)`)
	match := emailPattern.FindStringSubmatch("test@example.com")
	if match != nil {
		fmt.Println("捕获组:")
		fmt.Println("  完整匹配:", match[0])
		fmt.Println("  用户:", match[1])
		fmt.Println("  域名:", match[2])
	}

	// 命名捕获组
	namedPattern := regexp.MustCompile(`(?P<user>\w+)@(?P<domain>\w+\.\w+)`)
	match = namedPattern.FindStringSubmatch("test@example.com")
	names := namedPattern.SubexpNames()
	fmt.Println("命名捕获组:")
	for i, name := range names {
		if i > 0 && name != "" {
			fmt.Printf("  %s: %s\n", name, match[i])
		}
	}
}

// ============================================
// 11. 并发 (Goroutine + Channel)
// ============================================
func concurrencyDemo() {
	fmt.Println("\n=== 11. 并发 (Goroutine + Channel) ===")

	// 基本 Goroutine
	fmt.Println("基本 Goroutine:")
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			time.Sleep(50 * time.Millisecond)
			fmt.Printf("  Goroutine %d 完成\n", n)
		}(i)
	}
	wg.Wait()

	// Channel
	fmt.Println("Channel:")
	ch := make(chan string, 3)

	go func() {
		ch <- "消息 1"
		ch <- "消息 2"
		ch <- "消息 3"
		close(ch)
	}()

	for msg := range ch {
		fmt.Println("  收到:", msg)
	}

	// Select
	fmt.Println("Select:")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "from ch1"
	}()

	go func() {
		time.Sleep(30 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-ch1:
			fmt.Println("  ", msg)
		case msg := <-ch2:
			fmt.Println("  ", msg)
		case <-time.After(100 * time.Millisecond):
			fmt.Println("  超时")
		}
	}

	// Mutex
	fmt.Println("Mutex:")
	var mu sync.Mutex
	counter := 0

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("  计数器:", counter)

	// Worker Pool
	fmt.Println("Worker Pool:")
	workerPoolDemo()
}

func workerPoolDemo() {
	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// 启动 3 个 worker
	for w := 0; w < 3; w++ {
		go func(id int) {
			for job := range jobs {
				time.Sleep(10 * time.Millisecond)
				results <- job * 2
			}
		}(w)
	}

	// 发送任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// 收集结果
	for i := 0; i < 5; i++ {
		fmt.Printf("  结果: %d\n", <-results)
	}
}

// ============================================
// 12. 文件 IO
// ============================================
func fileIODemo() {
	fmt.Println("\n=== 12. 文件 IO ===")

	// 创建测试目录
	testDir := filepath.Join(".", "test_output")
	os.MkdirAll(testDir, 0755)

	testFile := filepath.Join(testDir, "test.txt")

	// 写入文件
	content := []byte("Hello, Go!\n第二行内容\n第三行")
	err := os.WriteFile(testFile, content, 0644)
	if err != nil {
		fmt.Println("写入错误:", err)
		return
	}

	// 读取整个文件
	data, err := os.ReadFile(testFile)
	if err != nil {
		fmt.Println("读取错误:", err)
		return
	}
	fmt.Println("文件内容:")
	fmt.Println(string(data))

	// 逐行读取
	fmt.Println("\n逐行读取:")
	file, _ := os.Open(testFile)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineNum := 1
	for scanner.Scan() {
		fmt.Printf("  行 %d: %s\n", lineNum, scanner.Text())
		lineNum++
	}

	// 目录遍历
	fmt.Println("\n目录遍历:")
	entries, _ := os.ReadDir(testDir)
	for _, entry := range entries {
		fmt.Println("  ", entry.Name())
	}

	// 文件信息
	fmt.Println("\n文件信息:")
	info, _ := os.Stat(testFile)
	fmt.Println("  名称:", info.Name())
	fmt.Println("  大小:", info.Size(), "bytes")
	fmt.Println("  修改时间:", info.ModTime().Format("2006-01-02 15:04:05"))
}

// ============================================
// 13. HTTP Client
// ============================================
func httpClientDemo() {
	fmt.Println("\n=== 13. HTTP Client ===")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// GET 请求
	fmt.Println("GET 请求:")
	resp, err := client.Get("https://httpbin.org/get")
	if err != nil {
		fmt.Println("请求失败 (可能无网络):", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("  状态码:", resp.StatusCode)

	body, _ := io.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	fmt.Println("  URL:", result["url"])

	// POST 请求
	fmt.Println("POST 请求:")
	payload := map[string]interface{}{
		"name": "Alice",
		"age":  25,
	}
	jsonData, _ := json.Marshal(payload)

	resp, err = client.Post(
		"https://httpbin.org/post",
		"application/json",
		bytes.NewBuffer(jsonData),
	)
	if err != nil {
		fmt.Println("请求失败:", err)
		return
	}
	defer resp.Body.Close()

	body, _ = io.ReadAll(resp.Body)
	json.Unmarshal(body, &result)
	fmt.Println("  响应 JSON:", result["json"])
}

// ============================================
// 14. defer / panic / recover
// ============================================
func deferPanicRecoverDemo() {
	fmt.Println("\n=== 14. defer / panic / recover ===")

	// defer (后进先出)
	fmt.Println("defer 顺序:")
	for i := 0; i < 3; i++ {
		defer fmt.Printf("  defer %d\n", i)
	}
	fmt.Println("  (defer 将在函数返回前执行)")

	// panic / recover
	fmt.Println("panic / recover:")
	safeDivide := func(a, b int) (result int) {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("  recovered from:", r)
				result = 0
			}
		}()
		return a / b
	}

	fmt.Println("  10/2 =", safeDivide(10, 2))
	fmt.Println("  10/0 =", safeDivide(10, 0))
}

// ============================================
// main
// ============================================
func main() {
	basicTypes()
	arraysAndSlices()
	mapsDemo()
	stringOperations()
	structsDemo()
	interfacesDemo()
	enumsDemo()
	errorHandlingDemo()
	dateTimeDemo()
	regexDemo()
	concurrencyDemo()
	fileIODemo()
	httpClientDemo()
	deferPanicRecoverDemo()

	fmt.Println("\n✅ 所有示例执行完成!")
}
