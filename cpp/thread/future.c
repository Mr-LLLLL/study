#include <iostream>
#include <future>
#include <mutex>
#include <unistd.h>

using std::cin;
using std::cout;
using std::endl;

std::atomic<int> g_i(0);	// atomic operator and better mutex, but atomic just support this operator
							// ++, *=, +=, *=, /=, &= ...
int g_ii = 0;
std::mutex mtx;

void test2() {
	for (int i = 0; i < 10000000; i++) {
		g_i += 1;
	}
}

int test() {
	cout << "thread start" << endl;
	sleep(1);
	return 1;
}

void test1(std::promise<int> &pro) {
	pro.set_value(1);
	sleep(1);
}

int main(int argc, char** argv) 
{
	//std::future<int> res = std::async(std::launch::async, test);	// async mean create a new thread
	//std::future<int> res = std::async(std::launch::deferred, test);	// if use res.wait() or res.get() will be call the
	//																	// test function in the thread
	//std::future_status stat = res.wait_for(std::chrono::seconds(1));	
	//if (stat == std::future_status::timeout) {
	//	cout << "timeout" << endl;
	//}
	//else if (stat == std::future_status::ready) {
	//	cout << "completed" << endl;
	//}
	//else if (stat == std::future_status::deferred) {
	//	cout << "thread be deferred" << endl;
	//}
	
	std::thread mythreads[10];
	for (auto &iter : mythreads)
		iter = std::thread(test2);
	for (auto &iter : mythreads)
		iter.join();
	cout << g_i << endl;

	//std::shared_future<int> res3 = std::async(std::launch::deferred, test);
	//cout << "first" << endl;
	//cout << res3.get() << endl;
	//cout << "second" << endl;
	//cout << res3.get() << endl;

	//std::packaged_task<int()> mypack(test);
	//std::thread mythread(std::ref(mypack));
	//mythread.detach();
	//std::future<int> res1 = mypack.get_future();
	//cout << res1.get() << endl;

	//std::promise<int> pro;			// require a value from other thread
	//std::thread mythread1(test1, std::ref(pro));
	//mythread1.join();
	//std::future<int> res2 = pro.get_future();
	//cout << res2.get() << endl;


	return 0;
}
