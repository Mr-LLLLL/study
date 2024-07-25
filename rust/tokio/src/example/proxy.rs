#![warn(rust_2018_idioms)]

use tokio::io::copy_bidirectional;
use tokio::net::{TcpListener, TcpStream};

use futures::FutureExt;
use std::env;

#[tokio::main]
pub async fn run() {
    let listen_addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:8081".to_string());
    let server_addr = env::args()
        .nth(2)
        .unwrap_or_else(|| "127.0.0.1:8080".to_string());

    println!("Listening on: {listen_addr}");
    println!("Proxying on: {server_addr}");

    let listener = TcpListener::bind(listen_addr).await.unwrap();

    while let Ok((mut inbound, _)) = listener.accept().await {
        let mut outbound = TcpStream::connect(server_addr.clone()).await.unwrap();

        tokio::spawn(async move {
            copy_bidirectional(&mut inbound, &mut outbound)
                .map(|r| {
                    if let Err(e) = r {
                        println!("Failed to transfer; error={e}");
                    }
                })
                .await
        });
    }
}
