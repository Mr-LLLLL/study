use std::env;
use std::io::{stdin, Read};
use std::net::SocketAddr;
use tokio::net::UdpSocket;

fn get_stdin_data() -> Result<Vec<u8>, Box<dyn std::error::Error>> {
    let mut buf = Vec::new();
    stdin().read_to_end(&mut buf)?;

    Ok(buf)
}

#[tokio::main]
pub async fn run() {
    let remote_addr: SocketAddr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:8080".into())
        .parse()
        .unwrap();

    let local_addr: SocketAddr = if remote_addr.is_ipv4() {
        "0.0.0.0:0"
    } else {
        "[::]:0"
    }
    .parse()
    .unwrap();

    let socket = UdpSocket::bind(local_addr).await.unwrap();
    const MAX_DATAGRAM_SIZE: usize = 65_507;
    socket.connect(&remote_addr).await.unwrap();
    let data = get_stdin_data().unwrap();
    socket.send(&data).await.unwrap();
    let mut data = vec![0u8; MAX_DATAGRAM_SIZE];
    let len = socket.recv(&mut data).await.unwrap();
    println!(
        "Received {} bytes:\n{}",
        len,
        String::from_utf8_lossy(&data[..len])
    );
}
