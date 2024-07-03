use std::{env, io};
use tokio::io::{AsyncReadExt, AsyncWriteExt};
use tokio::net::TcpListener;
use tokio::net::UdpSocket;

struct Server {
    addr: String,
}

impl Server {
    async fn run_udp(self) -> Result<(), io::Error> {
        let Server { addr } = self;
        let mut buf = vec![0; 1024];

        let socket = UdpSocket::bind(&addr).await.unwrap();
        println!("ListeningUdp on: {}", socket.local_addr().unwrap());

        loop {
            if let Some((size, peer)) = Some(socket.recv_from(&mut buf).await?) {
                let amt = socket.send_to(&buf[..size], &peer).await?;

                println!(
                    "Echoed {}/{} bytes to {}, content:{}",
                    amt,
                    size,
                    peer,
                    String::from_utf8(buf.clone()).expect(""),
                );
            }
        }
    }

    async fn run_tcp(self) -> Result<(), io::Error> {
        let Server { addr } = self;

        let listener = TcpListener::bind(&addr).await.unwrap();
        println!("ListeningTcp on: {}", addr);

        loop {
            let (mut socket, _) = listener.accept().await?;

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
}

#[tokio::main]
pub async fn run() {
    let addr = env::args()
        .nth(1)
        .unwrap_or_else(|| "127.0.0.1:8080".to_string());

    tokio::try_join!(
        Server { addr: addr.clone() }.run_tcp(),
        Server { addr }.run_udp()
    )
    .unwrap();
}
