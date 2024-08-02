use tokio::net::UdpSocket;
use tokio::{io, time};
use tokio_stream::StreamExt;
use tokio_util::codec::BytesCodec;
use tokio_util::udp::UdpFramed;

use bytes::Bytes;
use futures::{FutureExt, SinkExt};
use std::env;
use std::net::SocketAddr;
use std::time::Duration;

#[tokio::main]
pub async fn run() {
    let addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:0".to_string());

    let a = UdpSocket::bind(&addr).await.unwrap();
    let b = UdpSocket::bind(&addr).await.unwrap();

    let b_addr = b.local_addr().unwrap();

    let mut a = UdpFramed::new(a, BytesCodec::new());
    let mut b = UdpFramed::new(b, BytesCodec::new());

    let a = ping(&mut a, b_addr);
    let b = pong(&mut b);

    match tokio::try_join!(a, b) {
        Err(e) => println!("an error occured; error = {:?}", e),
        _ => println!("done!"),
    }
}

async fn ping(socket: &mut UdpFramed<BytesCodec>, b_addr: SocketAddr) -> Result<(), io::Error> {
    socket.send((Bytes::from(&b"PING"[..]), b_addr)).await?;

    for _ in 0..4usize {
        let (bytes, addr) = socket.next().map(|e| e.unwrap()).await?;

        println!("[a] recv: {}", String::from_utf8_lossy(&bytes));

        socket.send((Bytes::from(&b"PING"[..]), addr)).await?;
    }

    Ok(())
}

async fn pong(socket: &mut UdpFramed<BytesCodec>) -> Result<(), io::Error> {
    let timeout = Duration::from_millis(200);

    while let Ok(Some(Ok((bytes, addr)))) = time::timeout(timeout, socket.next()).await {
        println!("[b] recv: {}", String::from_utf8_lossy(&bytes));

        socket.send((Bytes::from(&b"PONG"[..]), addr)).await?;
    }

    Ok(())
}
