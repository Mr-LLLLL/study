use tokio::net::TcpListener;
use tokio_stream::StreamExt;
use tokio_util::codec::{BytesCodec, Decoder};

use std::env;

#[tokio::main]
pub async fn run() {
    let addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:8080".to_string());

    let listener = TcpListener::bind(&addr).await.unwrap();
    println!("Listening on:{}", addr);

    loop {
        let (socket, _) = listener.accept().await.unwrap();

        tokio::spawn(async move {
            let mut framed = BytesCodec::new().framed(socket);

            while let Some(message) = framed.next().await {
                match message {
                    Ok(bytes) => println!("bytes: {:?}", bytes),
                    Err(err) => println!("Socket closed with error: {:?}", err),
                }
            }

            println!("Socket received FIN packet and closed connectin");
        });
    }
}
