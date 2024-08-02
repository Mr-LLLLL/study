#include <condition_variable>
#include <functional>
#include <future>
#include <iostream>
#include <memory>
#include <mutex>
#include <string>
#include <thread>
#include <type_traits>
#include <unistd.h>
#include <vector>

using std::cout;
using std::endl;

class Test {
public:
  void print() {
    sleep(1);
    std::unique_lock<std::mutex> ulock(mutex);
    cond.wait(ulock, [this] { return i == 5; });
    cout << "print() id = " << std::this_thread::get_id() << endl;
  }
  void notify() {
    std::lock_guard<std::mutex> lock(mutex);
    cout << "start yell" << endl;
    cond.notify_all();
    i = 5;

    { ; }
  }

private:
  std::mutex mutex;
  std::condition_variable cond;
  int i = 0;
};

class A {
public:
  A(int j = 2) : i(j) { cout << "A()" << endl; }
  A(const A &) { cout << "A(const&)" << endl; }
  // A(const A &&r) {
  //	cout << "A(const&&)" << endl;
  //	i = std::move(r.i);
  // }
  void print() { cout << "A::print()" << endl; }
  virtual void printf() { cout << "A::printf()" << endl; }

  virtual ~A() { cout << "~A()" << endl; }

  int i = 0;
};

class B : public A {
public:
  B() { cout << "B()" << endl; }
  void print() { cout << "B::print()" << endl; }
  void printf() override { cout << "B:printf()" << endl; }
  ~B() { cout << "~B()" << endl; }
  explicit operator int() {
    cout << "operator int()" << endl;
    cout << "thread = " << std::this_thread::get_id() << endl;
    return 1;
  }
};

class Ta {
public:
  Ta() { cout << "Ta thread = " << std::this_thread::get_id() << endl; }
  void print() {
    cout << "print()" << endl;
    cout << "thread = " << std::this_thread::get_id() << endl;
    sleep(5);
  }
  int a;
  int b;
};

int test(std::promise<int> &temp) {
  temp.set_value(4);
  cout << "test()" << endl;
  return 1;
}

int main(int argc, char *argv[]) {
  Ta t;
  std::future<void> res = std::async(std::launch::async, &Ta::print, &t);
  cout << "main filished " << endl;
  // res.wait();
  // std::function<void()> f(test);
  // std::packaged_task<int()> f(test);
  // std::thread mythread(std::ref(f));
  // mythread.join();
  //
  //	std::promise<int> pro;
  //	std::thread mythread(test, std::ref(pro));
  //	mythread.join();
  //	std::future<int> res = pro.get_future();
  //	cout << res.get() << endl;
  //
  return 0;
}
