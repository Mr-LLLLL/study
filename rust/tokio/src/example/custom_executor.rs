use tokio::net::TcpListener;
use tokio::sync::oneshot;

pub fn run() {
    let (tx, rx) = oneshot::channel();

    my_custom_runtime::spawn(async move {
        let listener = TcpListener::bind("0.0.0.0:0").await.unwrap();

        println!("addr: {:?}", listener.local_addr());

        tx.send(()).unwrap()
    });

    futures::executor::block_on(rx).unwrap();
}

mod my_custom_runtime {
    use once_cell::sync::Lazy;
    use std::future::Future;
    use tokio_util::context::TokioContext;

    pub fn spawn(f: impl Future<Output = ()> + Send + 'static) {
        EXECUTOR.spawn(f);
    }

    struct ThreadPool {
        inner: futures::executor::ThreadPool,
        rt: tokio::runtime::Runtime,
    }

    static EXECUTOR: Lazy<ThreadPool> = Lazy::new(|| {
        let rt = tokio::runtime::Builder::new_multi_thread()
            .enable_all()
            .build()
            .unwrap();
        let inner = futures::executor::ThreadPool::builder().create().unwrap();

        ThreadPool { inner, rt }
    });

    impl ThreadPool {
        fn spawn(&self, f: impl Future<Output = ()> + Send + 'static) {
            let handle = self.rt.handle().clone();
            self.inner.spawn_ok(TokioContext::new(f, handle));
        }
    }
}
