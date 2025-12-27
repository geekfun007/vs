#!/usr/bin/env python3
"""
Python HTTP Server 示例 (使用标准库)
运行: python http_server.py
"""

from http.server import HTTPServer, BaseHTTPRequestHandler
from urllib.parse import urlparse, parse_qs
import json
from typing import Any
import re

# 用户数据
users: list[dict[str, Any]] = [
    {"id": 1, "name": "Alice", "email": "alice@example.com"},
    {"id": 2, "name": "Bob", "email": "bob@example.com"},
]
next_id = 3


class APIHandler(BaseHTTPRequestHandler):
    """简单的 REST API 处理器"""
    
    def send_json(self, data: Any, status: int = 200) -> None:
        """发送 JSON 响应"""
        self.send_response(status)
        self.send_header('Content-Type', 'application/json')
        self.send_header('Access-Control-Allow-Origin', '*')
        self.end_headers()
        self.wfile.write(json.dumps(data, ensure_ascii=False).encode('utf-8'))
    
    def read_json(self) -> dict:
        """读取 JSON 请求体"""
        content_length = int(self.headers.get('Content-Length', 0))
        if content_length == 0:
            return {}
        body = self.rfile.read(content_length).decode('utf-8')
        return json.loads(body)
    
    def log_message(self, format: str, *args) -> None:
        """自定义日志格式"""
        print(f"{self.date_time_string()} {args[0]}")
    
    # GET 请求处理
    def do_GET(self) -> None:
        parsed = urlparse(self.path)
        path = parsed.path
        
        # GET /
        if path == '/':
            self.send_json({
                "message": "Welcome to Python HTTP Server",
                "endpoints": [
                    "GET /users - 获取所有用户",
                    "GET /users/<id> - 获取单个用户",
                    "POST /users - 创建用户",
                    "PUT /users/<id> - 更新用户",
                    "DELETE /users/<id> - 删除用户",
                ]
            })
            return
        
        # GET /users
        if path == '/users':
            self.send_json(users)
            return
        
        # GET /users/<id>
        match = re.match(r'^/users/(\d+)$', path)
        if match:
            user_id = int(match.group(1))
            user = next((u for u in users if u['id'] == user_id), None)
            if user:
                self.send_json(user)
            else:
                self.send_json({"error": "User not found"}, 404)
            return
        
        self.send_json({"error": "Not found"}, 404)
    
    # POST 请求处理
    def do_POST(self) -> None:
        global next_id
        
        if self.path == '/users':
            try:
                body = self.read_json()
                if not body.get('name') or not body.get('email'):
                    self.send_json({"error": "Name and email required"}, 400)
                    return
                
                new_user = {
                    "id": next_id,
                    "name": body['name'],
                    "email": body['email'],
                }
                next_id += 1
                users.append(new_user)
                self.send_json(new_user, 201)
            except json.JSONDecodeError:
                self.send_json({"error": "Invalid JSON"}, 400)
            return
        
        self.send_json({"error": "Not found"}, 404)
    
    # PUT 请求处理
    def do_PUT(self) -> None:
        match = re.match(r'^/users/(\d+)$', self.path)
        if match:
            user_id = int(match.group(1))
            user = next((u for u in users if u['id'] == user_id), None)
            
            if not user:
                self.send_json({"error": "User not found"}, 404)
                return
            
            try:
                body = self.read_json()
                if 'name' in body:
                    user['name'] = body['name']
                if 'email' in body:
                    user['email'] = body['email']
                self.send_json(user)
            except json.JSONDecodeError:
                self.send_json({"error": "Invalid JSON"}, 400)
            return
        
        self.send_json({"error": "Not found"}, 404)
    
    # DELETE 请求处理
    def do_DELETE(self) -> None:
        global users
        
        match = re.match(r'^/users/(\d+)$', self.path)
        if match:
            user_id = int(match.group(1))
            user = next((u for u in users if u['id'] == user_id), None)
            
            if not user:
                self.send_json({"error": "User not found"}, 404)
                return
            
            users = [u for u in users if u['id'] != user_id]
            self.send_json({"message": "Deleted", "user": user})
            return
        
        self.send_json({"error": "Not found"}, 404)
    
    # OPTIONS 请求处理 (CORS)
    def do_OPTIONS(self) -> None:
        self.send_response(204)
        self.send_header('Access-Control-Allow-Origin', '*')
        self.send_header('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS')
        self.send_header('Access-Control-Allow-Headers', 'Content-Type')
        self.end_headers()


def run_server(port: int = 8000) -> None:
    """启动服务器"""
    server = HTTPServer(('', port), APIHandler)
    print(f"🚀 Python HTTP Server running at http://localhost:{port}")
    print("Available endpoints:")
    print("  GET    /         - API 信息")
    print("  GET    /users    - 获取所有用户")
    print("  GET    /users/<id> - 获取单个用户")
    print("  POST   /users    - 创建用户")
    print("  PUT    /users/<id> - 更新用户")
    print("  DELETE /users/<id> - 删除用户")
    
    try:
        server.serve_forever()
    except KeyboardInterrupt:
        print("\n服务器已停止")
        server.shutdown()


if __name__ == '__main__':
    run_server()
