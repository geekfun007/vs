// Go HTTP Server 示例
// 运行: go run http_server.go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"sync"
	"time"
)

// User 用户模型
type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// 全局数据存储
var (
	users = []User{
		{ID: 1, Name: "Alice", Email: "alice@example.com"},
		{ID: 2, Name: "Bob", Email: "bob@example.com"},
	}
	nextID = 3
	mu     sync.RWMutex
)

// 辅助函数
func sendJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func parseJSON(r *http.Request, v interface{}) error {
	return json.NewDecoder(r.Body).Decode(v)
}

// 日志中间件
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s %v", r.RemoteAddr, r.Method, r.URL.Path, time.Since(start))
	})
}

// CORS 中间件
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

// 路由处理器
func indexHandler(w http.ResponseWriter, r *http.Request) {
	sendJSON(w, map[string]interface{}{
		"message": "Welcome to Go HTTP Server",
		"endpoints": []string{
			"GET /users - 获取所有用户",
			"GET /users/{id} - 获取单个用户",
			"POST /users - 创建用户",
			"PUT /users/{id} - 更新用户",
			"DELETE /users/{id} - 删除用户",
		},
	}, http.StatusOK)
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	mu.RLock()
	defer mu.RUnlock()
	sendJSON(w, users, http.StatusOK)
}

func getUserHandler(w http.ResponseWriter, r *http.Request, id int) {
	mu.RLock()
	defer mu.RUnlock()

	for _, user := range users {
		if user.ID == id {
			sendJSON(w, user, http.StatusOK)
			return
		}
	}
	sendJSON(w, map[string]string{"error": "User not found"}, http.StatusNotFound)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := parseJSON(r, &input); err != nil {
		sendJSON(w, map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Email == "" {
		sendJSON(w, map[string]string{"error": "Name and email required"}, http.StatusBadRequest)
		return
	}

	mu.Lock()
	newUser := User{
		ID:    nextID,
		Name:  input.Name,
		Email: input.Email,
	}
	nextID++
	users = append(users, newUser)
	mu.Unlock()

	sendJSON(w, newUser, http.StatusCreated)
}

func updateUserHandler(w http.ResponseWriter, r *http.Request, id int) {
	var input struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := parseJSON(r, &input); err != nil {
		sendJSON(w, map[string]string{"error": "Invalid JSON"}, http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	for i := range users {
		if users[i].ID == id {
			if input.Name != "" {
				users[i].Name = input.Name
			}
			if input.Email != "" {
				users[i].Email = input.Email
			}
			sendJSON(w, users[i], http.StatusOK)
			return
		}
	}
	sendJSON(w, map[string]string{"error": "User not found"}, http.StatusNotFound)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request, id int) {
	mu.Lock()
	defer mu.Unlock()

	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			sendJSON(w, map[string]interface{}{
				"message": "Deleted",
				"user":    user,
			}, http.StatusOK)
			return
		}
	}
	sendJSON(w, map[string]string{"error": "User not found"}, http.StatusNotFound)
}

// 路由器
func router() http.Handler {
	userIDPattern := regexp.MustCompile(`^/users/(\d+)$`)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path

		// GET /
		if path == "/" && r.Method == "GET" {
			indexHandler(w, r)
			return
		}

		// /users
		if path == "/users" {
			switch r.Method {
			case "GET":
				getUsersHandler(w, r)
			case "POST":
				createUserHandler(w, r)
			default:
				sendJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
			}
			return
		}

		// /users/{id}
		if matches := userIDPattern.FindStringSubmatch(path); matches != nil {
			var id int
			fmt.Sscanf(matches[1], "%d", &id)

			switch r.Method {
			case "GET":
				getUserHandler(w, r, id)
			case "PUT":
				updateUserHandler(w, r, id)
			case "DELETE":
				deleteUserHandler(w, r, id)
			default:
				sendJSON(w, map[string]string{"error": "Method not allowed"}, http.StatusMethodNotAllowed)
			}
			return
		}

		sendJSON(w, map[string]string{"error": "Not found"}, http.StatusNotFound)
	})
}

func main() {
	// 应用中间件
	handler := loggingMiddleware(corsMiddleware(router()))

	port := ":8080"
	fmt.Printf("🚀 Go HTTP Server running at http://localhost%s\n", port)
	fmt.Println("Available endpoints:")
	fmt.Println("  GET    /         - API 信息")
	fmt.Println("  GET    /users    - 获取所有用户")
	fmt.Println("  GET    /users/{id} - 获取单个用户")
	fmt.Println("  POST   /users    - 创建用户")
	fmt.Println("  PUT    /users/{id} - 更新用户")
	fmt.Println("  DELETE /users/{id} - 删除用户")

	log.Fatal(http.ListenAndServe(port, handler))
}
