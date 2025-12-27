//! Rust HTTP Client 示例
//! 运行: cargo run --bin http_client

use serde::{Deserialize, Serialize};
use std::error::Error;

#[derive(Debug, Serialize, Deserialize)]
struct HttpBinResponse {
    url: String,
    origin: String,
    headers: std::collections::HashMap<String, String>,
    #[serde(default)]
    json: Option<serde_json::Value>,
}

#[derive(Debug, Serialize)]
struct PostData {
    name: String,
    age: i32,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn Error>> {
    println!("=== Rust HTTP Client 示例 ===\n");

    // 创建客户端
    let client = reqwest::Client::new();

    // GET 请求
    println!("GET 请求:");
    let response = client
        .get("https://httpbin.org/get")
        .send()
        .await?;

    println!("  状态码: {}", response.status());

    let body: HttpBinResponse = response.json().await?;
    println!("  URL: {}", body.url);
    println!("  Origin: {}", body.origin);

    // POST 请求
    println!("\nPOST 请求:");
    let data = PostData {
        name: "Alice".to_string(),
        age: 25,
    };

    let response = client
        .post("https://httpbin.org/post")
        .json(&data)
        .header("X-Custom-Header", "custom-value")
        .send()
        .await?;

    println!("  状态码: {}", response.status());

    let body: HttpBinResponse = response.json().await?;
    println!("  JSON: {:?}", body.json);

    // 超时设置
    println!("\n超时设置:");
    let timeout_client = reqwest::Client::builder()
        .timeout(std::time::Duration::from_secs(10))
        .build()?;

    let response = timeout_client
        .get("https://httpbin.org/delay/1")
        .send()
        .await?;
    println!("  延迟请求状态: {}", response.status());

    // 并发请求
    println!("\n并发请求:");
    let urls = vec![
        "https://httpbin.org/get?id=1",
        "https://httpbin.org/get?id=2",
        "https://httpbin.org/get?id=3",
    ];

    let futures: Vec<_> = urls
        .iter()
        .map(|url| client.get(*url).send())
        .collect();

    let results = futures::future::join_all(futures).await;
    for (i, result) in results.into_iter().enumerate() {
        match result {
            Ok(response) => println!("  请求 {}: {}", i + 1, response.status()),
            Err(e) => println!("  请求 {}: 错误 - {}", i + 1, e),
        }
    }

    println!("\n✅ HTTP Client 示例完成!");
    Ok(())
}
