#include <iostream>
#include <unistd.h>
#include <condition_variable>
#include <mutex>
#include <thread>

using std::cin;
using std::cout;
using std::endl;

class A {
public:
	void func() {
		std::unique_lock<std::mutex> ulock(mutex);
		cond.wait(ulock);			// unlock the ulock and block until notify
		cout << "func" << endl;
		i = 1;
	}

	void func2() {
		std::unique_lock<std::mutex> ulock(mutex);
		cond.wait(ulock, [this] () { return i == 1; });	// if second return is false, block and unlock ulock
		cout << "func2" << endl;						// each notify will check the result. until catch mutex first and
														// the result is true;
	}

	void notify() {
		sleep(1);
		cond.notify_all();	//cond.notify_one()  notify only one thread
		sleep(1);
		cond.notify_all();	// if notify one time, thread2 may be block;
	}
private:
	int i = 0;
	std::condition_variable cond;
	std::mutex mutex;
};


int main(int argc, char **argv)
{
	A a;
	std::thread mythread(&A::func, &a);
	std::thread mythread1(&A::func2, &a);
	std::thread mythread2(&A::notify, &a);
	mythread.join();
	mythread1.join();
	mythread2.join();

	return 0;
}
