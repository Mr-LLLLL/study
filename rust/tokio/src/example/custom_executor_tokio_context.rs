use tokio::net::TcpListener;
use tokio::runtime::Builder;
use tokio::sync::oneshot;
use tokio_util::context::RuntimeExt;

pub fn run() {
    let (tx, rx) = oneshot::channel();
    let rt1 = Builder::new_multi_thread()
        .worker_threads(1)
        .build()
        .unwrap();
    let rt2 = Builder::new_multi_thread()
        .worker_threads(1)
        .enable_all()
        .build()
        .unwrap();

    rt1.block_on(rt2.wrap(async move {
        let listener = TcpListener::bind("0.0.0.0:0").await.unwrap();
        println!("addr: {:?}", listener.local_addr());
        tx.send(()).unwrap();
    }));
    futures::executor::block_on(rx).unwrap();
}
