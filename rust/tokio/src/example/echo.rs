use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpListener;

use std::env;

#[tokio::main]
pub async fn run() {
    let addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:8080".to_string());

    let listener = TcpListener::bind(&addr).await.unwrap();
    println!("Listening on: {}", addr);

    loop {
        let (mut socket, _) = listener.accept().await.unwrap();

        tokio::spawn(async move {
            let mut buf = vec![0; 1024];

            loop {
                let n = socket
                    .read(&mut buf)
                    .await
                    .expect("failed to read data from socket");

                if n == 0 {
                    return;
                }

                println!("{}", String::from_utf8(buf.clone()).expect(""));

                socket
                    .write_all(&buf[0..n])
                    .await
                    .expect("failed to write data t osocket");
            }
        });
    }
}
