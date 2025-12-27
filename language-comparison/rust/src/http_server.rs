//! Rust HTTP Server 示例 (使用 Axum)
//! 运行: cargo run --bin http_server

use axum::{
    extract::{Path, State},
    http::StatusCode,
    response::IntoResponse,
    routing::{delete, get, post, put},
    Json, Router,
};
use serde::{Deserialize, Serialize};
use std::{
    collections::HashMap,
    net::SocketAddr,
    sync::{Arc, RwLock},
};

// ============================================
// 数据模型
// ============================================

#[derive(Debug, Clone, Serialize, Deserialize)]
struct User {
    id: u32,
    name: String,
    email: String,
}

#[derive(Debug, Deserialize)]
struct CreateUser {
    name: String,
    email: String,
}

#[derive(Debug, Deserialize)]
struct UpdateUser {
    name: Option<String>,
    email: Option<String>,
}

#[derive(Debug, Serialize)]
struct ErrorResponse {
    error: String,
}

#[derive(Debug, Serialize)]
struct DeleteResponse {
    message: String,
    user: User,
}

#[derive(Debug, Serialize)]
struct IndexResponse {
    message: String,
    endpoints: Vec<String>,
}

// ============================================
// 应用状态
// ============================================

struct AppState {
    users: RwLock<HashMap<u32, User>>,
    next_id: RwLock<u32>,
}

impl AppState {
    fn new() -> Self {
        let mut users = HashMap::new();
        users.insert(
            1,
            User {
                id: 1,
                name: "Alice".to_string(),
                email: "alice@example.com".to_string(),
            },
        );
        users.insert(
            2,
            User {
                id: 2,
                name: "Bob".to_string(),
                email: "bob@example.com".to_string(),
            },
        );

        AppState {
            users: RwLock::new(users),
            next_id: RwLock::new(3),
        }
    }
}

type SharedState = Arc<AppState>;

// ============================================
// 路由处理器
// ============================================

// GET /
async fn index() -> Json<IndexResponse> {
    Json(IndexResponse {
        message: "Welcome to Rust HTTP Server (Axum)".to_string(),
        endpoints: vec![
            "GET /users - 获取所有用户".to_string(),
            "GET /users/:id - 获取单个用户".to_string(),
            "POST /users - 创建用户".to_string(),
            "PUT /users/:id - 更新用户".to_string(),
            "DELETE /users/:id - 删除用户".to_string(),
        ],
    })
}

// GET /users
async fn get_users(State(state): State<SharedState>) -> Json<Vec<User>> {
    let users = state.users.read().unwrap();
    let users_vec: Vec<User> = users.values().cloned().collect();
    Json(users_vec)
}

// GET /users/:id
async fn get_user(
    State(state): State<SharedState>,
    Path(id): Path<u32>,
) -> Result<Json<User>, (StatusCode, Json<ErrorResponse>)> {
    let users = state.users.read().unwrap();

    match users.get(&id) {
        Some(user) => Ok(Json(user.clone())),
        None => Err((
            StatusCode::NOT_FOUND,
            Json(ErrorResponse {
                error: "User not found".to_string(),
            }),
        )),
    }
}

// POST /users
async fn create_user(
    State(state): State<SharedState>,
    Json(input): Json<CreateUser>,
) -> Result<(StatusCode, Json<User>), (StatusCode, Json<ErrorResponse>)> {
    if input.name.is_empty() || input.email.is_empty() {
        return Err((
            StatusCode::BAD_REQUEST,
            Json(ErrorResponse {
                error: "Name and email required".to_string(),
            }),
        ));
    }

    let mut next_id = state.next_id.write().unwrap();
    let id = *next_id;
    *next_id += 1;

    let user = User {
        id,
        name: input.name,
        email: input.email,
    };

    let mut users = state.users.write().unwrap();
    users.insert(id, user.clone());

    Ok((StatusCode::CREATED, Json(user)))
}

// PUT /users/:id
async fn update_user(
    State(state): State<SharedState>,
    Path(id): Path<u32>,
    Json(input): Json<UpdateUser>,
) -> Result<Json<User>, (StatusCode, Json<ErrorResponse>)> {
    let mut users = state.users.write().unwrap();

    match users.get_mut(&id) {
        Some(user) => {
            if let Some(name) = input.name {
                user.name = name;
            }
            if let Some(email) = input.email {
                user.email = email;
            }
            Ok(Json(user.clone()))
        }
        None => Err((
            StatusCode::NOT_FOUND,
            Json(ErrorResponse {
                error: "User not found".to_string(),
            }),
        )),
    }
}

// DELETE /users/:id
async fn delete_user(
    State(state): State<SharedState>,
    Path(id): Path<u32>,
) -> Result<Json<DeleteResponse>, (StatusCode, Json<ErrorResponse>)> {
    let mut users = state.users.write().unwrap();

    match users.remove(&id) {
        Some(user) => Ok(Json(DeleteResponse {
            message: "Deleted".to_string(),
            user,
        })),
        None => Err((
            StatusCode::NOT_FOUND,
            Json(ErrorResponse {
                error: "User not found".to_string(),
            }),
        )),
    }
}

// ============================================
// 主函数
// ============================================

#[tokio::main]
async fn main() {
    // 创建共享状态
    let state = Arc::new(AppState::new());

    // 构建路由
    let app = Router::new()
        .route("/", get(index))
        .route("/users", get(get_users).post(create_user))
        .route(
            "/users/:id",
            get(get_user).put(update_user).delete(delete_user),
        )
        .with_state(state);

    // 启动服务器
    let addr = SocketAddr::from(([0, 0, 0, 0], 3000));
    println!("🚀 Rust HTTP Server (Axum) running at http://localhost:3000");
    println!("Available endpoints:");
    println!("  GET    /         - API 信息");
    println!("  GET    /users    - 获取所有用户");
    println!("  GET    /users/:id - 获取单个用户");
    println!("  POST   /users    - 创建用户");
    println!("  PUT    /users/:id - 更新用户");
    println!("  DELETE /users/:id - 删除用户");

    let listener = tokio::net::TcpListener::bind(addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
