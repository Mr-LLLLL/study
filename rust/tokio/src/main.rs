#![allow(dead_code)]

use std::time::Duration;
use tokio::time;

mod example;

pub fn main() {
    // example::hello_world::run();
    // example::echo::run();
    // example::connect::run();
    // example::chat::run();
    // example::custom_executor_tokio_context::run();
    // example::dump::run();
    // example::proxy::run();
<<<<<<< HEAD
    // example::udp_client::run();
    example::udp_codec::run();
=======
    // example::print_each_packet::run();

    // tokio_main();
    // tokio_spawn();
    tokio_work_may_dropped();
}

fn tokio_work_may_dropped() {
    tokio::runtime::Runtime::new().unwrap().block_on(async {
        let handle = tokio::spawn(async { println!("test1") });
        time::sleep(Duration::from_secs(1)).await;
        tokio::spawn(async { println!("test2") });
        tokio::spawn(async { println!("test3") });
        tokio::spawn(async { println!("test4") });
        tokio::spawn(async { println!("test5") });
        tokio::spawn(process());

        handle.await.unwrap();
    })
}

#[tokio::main]
async fn tokio_spawn() {
    // here running process
    let handle = tokio::spawn(process());

    time::sleep(Duration::from_secs(1)).await;

    println!("hello");

    time::sleep(Duration::from_secs(1)).await;

    let out = handle.await.unwrap();
    println!("Main Finished Got {}", out);
}

fn tokio_main() {
    let rt = tokio::runtime::Runtime::new().unwrap();
    rt.block_on(async {
        let op = process();

        time::sleep(Duration::from_secs(1)).await;

        println!("hello");

        time::sleep(Duration::from_secs(1)).await;

        op.await;

        // here running process
        println!("Main finished");
    });
}

async fn process() -> i32 {
    println!("world");

    time::sleep(Duration::from_secs(1)).await;

    println!("coroutine finished");

    return 1;
>>>>>>> 233e8303b746fb76cadf033b3495c9873d77051b
}
