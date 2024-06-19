use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpStream;

#[tokio::main]
pub async fn run() {
    let mut stream = TcpStream::connect("127.0.0.1:8080").await.unwrap();
    println!("created stream");

    let result = stream.write_all(b"hello wrold\n").await;
    println!("wrote to stream; success={:?}", result.is_ok());

    let mut buf = vec![0; 1024];
    stream
        .read(&mut buf)
        .await
        .expect("failed to read data from steam");

    println!("{}", String::from_utf8(buf).expect(""));
}
