/**
 * TypeScript HTTP Server 示例
 * 运行: npx tsx http-server.ts
 */

import { createServer, IncomingMessage, ServerResponse } from 'http';

// 简单的路由类型
interface Route {
    method: string;
    path: string | RegExp;
    handler: (req: IncomingMessage, res: ServerResponse, params?: Record<string, string>) => void;
}

// 用户数据
interface User {
    id: number;
    name: string;
    email: string;
}

const users: User[] = [
    { id: 1, name: "Alice", email: "alice@example.com" },
    { id: 2, name: "Bob", email: "bob@example.com" },
];

let nextId = 3;

// 辅助函数
function sendJSON(res: ServerResponse, data: unknown, statusCode = 200): void {
    res.writeHead(statusCode, { 'Content-Type': 'application/json' });
    res.end(JSON.stringify(data));
}

function parseBody(req: IncomingMessage): Promise<unknown> {
    return new Promise((resolve, reject) => {
        let body = '';
        req.on('data', chunk => body += chunk);
        req.on('end', () => {
            try {
                resolve(body ? JSON.parse(body) : {});
            } catch {
                reject(new Error('Invalid JSON'));
            }
        });
    });
}

// 路由定义
const routes: Route[] = [
    // GET /
    {
        method: 'GET',
        path: '/',
        handler: (req, res) => {
            sendJSON(res, {
                message: 'Welcome to TypeScript HTTP Server',
                endpoints: [
                    'GET /users - 获取所有用户',
                    'GET /users/:id - 获取单个用户',
                    'POST /users - 创建用户',
                    'PUT /users/:id - 更新用户',
                    'DELETE /users/:id - 删除用户',
                ]
            });
        }
    },
    // GET /users
    {
        method: 'GET',
        path: '/users',
        handler: (req, res) => {
            sendJSON(res, users);
        }
    },
    // GET /users/:id
    {
        method: 'GET',
        path: /^\/users\/(\d+)$/,
        handler: (req, res, params) => {
            const id = parseInt(params?.id || '0');
            const user = users.find(u => u.id === id);
            if (user) {
                sendJSON(res, user);
            } else {
                sendJSON(res, { error: 'User not found' }, 404);
            }
        }
    },
    // POST /users
    {
        method: 'POST',
        path: '/users',
        handler: async (req, res) => {
            try {
                const body = await parseBody(req) as Partial<User>;
                if (!body.name || !body.email) {
                    sendJSON(res, { error: 'Name and email required' }, 400);
                    return;
                }
                const newUser: User = {
                    id: nextId++,
                    name: body.name,
                    email: body.email,
                };
                users.push(newUser);
                sendJSON(res, newUser, 201);
            } catch (error) {
                sendJSON(res, { error: 'Invalid request body' }, 400);
            }
        }
    },
    // PUT /users/:id
    {
        method: 'PUT',
        path: /^\/users\/(\d+)$/,
        handler: async (req, res, params) => {
            const id = parseInt(params?.id || '0');
            const index = users.findIndex(u => u.id === id);
            if (index === -1) {
                sendJSON(res, { error: 'User not found' }, 404);
                return;
            }
            try {
                const body = await parseBody(req) as Partial<User>;
                users[index] = { ...users[index], ...body, id };
                sendJSON(res, users[index]);
            } catch {
                sendJSON(res, { error: 'Invalid request body' }, 400);
            }
        }
    },
    // DELETE /users/:id
    {
        method: 'DELETE',
        path: /^\/users\/(\d+)$/,
        handler: (req, res, params) => {
            const id = parseInt(params?.id || '0');
            const index = users.findIndex(u => u.id === id);
            if (index === -1) {
                sendJSON(res, { error: 'User not found' }, 404);
                return;
            }
            const deleted = users.splice(index, 1)[0];
            sendJSON(res, { message: 'Deleted', user: deleted });
        }
    },
];

// 路由匹配
function matchRoute(method: string, path: string): { route: Route; params: Record<string, string> } | null {
    for (const route of routes) {
        if (route.method !== method) continue;
        
        if (typeof route.path === 'string') {
            if (route.path === path) {
                return { route, params: {} };
            }
        } else {
            const match = path.match(route.path);
            if (match) {
                return { route, params: { id: match[1] } };
            }
        }
    }
    return null;
}

// 创建服务器
const server = createServer(async (req, res) => {
    const method = req.method || 'GET';
    const url = new URL(req.url || '/', `http://${req.headers.host}`);
    const path = url.pathname;
    
    // 日志
    console.log(`${new Date().toISOString()} ${method} ${path}`);
    
    // CORS 头
    res.setHeader('Access-Control-Allow-Origin', '*');
    res.setHeader('Access-Control-Allow-Methods', 'GET, POST, PUT, DELETE, OPTIONS');
    res.setHeader('Access-Control-Allow-Headers', 'Content-Type');
    
    if (method === 'OPTIONS') {
        res.writeHead(204);
        res.end();
        return;
    }
    
    // 路由匹配
    const matched = matchRoute(method, path);
    if (matched) {
        try {
            await matched.route.handler(req, res, matched.params);
        } catch (error) {
            console.error('Handler error:', error);
            sendJSON(res, { error: 'Internal Server Error' }, 500);
        }
    } else {
        sendJSON(res, { error: 'Not Found' }, 404);
    }
});

const PORT = 3000;
server.listen(PORT, () => {
    console.log(`🚀 TypeScript HTTP Server running at http://localhost:${PORT}`);
    console.log('Available endpoints:');
    console.log('  GET    /         - API 信息');
    console.log('  GET    /users    - 获取所有用户');
    console.log('  GET    /users/:id - 获取单个用户');
    console.log('  POST   /users    - 创建用户');
    console.log('  PUT    /users/:id - 更新用户');
    console.log('  DELETE /users/:id - 删除用户');
});
