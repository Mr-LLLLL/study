#include <iostream>
#include <unistd.h>
#include <mutex>
#include <thread>

using std::cin;
using std::cout;
using std::endl;

int g_i = 0;
std::mutex mutex;				// have three member function 
								// 1. mutex.lock();
								// 2. mutex.unlock();
								// 3. mutex.try_lock(); // return bool
std::recursive_mutex rmutex;		// recursive mutex; just like normal mtuex, std::lock_guard<std::recursive_mutex> lock
									// (rmutex), either std::unique<std::recursive_mutex> ulock(rmutex);
std::timed_mutex tmutex;		// time wait mutex, similar mutex and recursive_mutex;
std::recursive_timed_mutex rtmutex;	// recursive and timed mtuex
std::once_flag g_flag;

void only_one_call() {
	cout << "i am only one" << endl;
}

void add() {
	int j = 0;
	while (j < 1000000) {
		mutex.lock();	// call mutex class function lock
		++g_i;
		++j;
		mutex.unlock();	// manual call mutex class function unlock
		std::call_once(g_flag, only_one_call);	//only call once 
	}
}

void add1() {
	int j = 0;
	while (j < 1000000) {
		/*
		 * std::unique_lock<std::mutex> ulock(mutex, std::adopt_lock)	equal lock_guard lock(mutex, std::adopt_lock);
		 * std::unique_lock<std::mutex> ulock(mutex, std::try_to_lock)	try to catch the lock
		 * std::unique_lock<std::mutex> ulock(mutex, std::defer_lock)	defer to lock and the mutex not lock before
		 */
		std::unique_lock<std::mutex> ulock(mutex); // ulock bind the mutex 
		++g_i;
		++j;
		ulock.unlock();		// manual unlock the unique_lock, either don't be, similar lock_guard
							// previous unlock
	}
}

void add2() {
	int j = 0;
	while (j < 1000000) {
		std::unique_lock<std::mutex> ulock(mutex, std::defer_lock);
		/*
		 * if (ulock.try_lock() == true)	try to lock equal std::try_to_lock
		 * std::mutex *pm = ulock.release();	release the ulock and mutex and assign the pm
		 */
		ulock.lock();	// don't neet to manual unlock
		++g_i;
		++j;
	}
}

void print() {
	int j = 0;
	while (j < 100) {
		std::unique_lock<std::mutex> ulock(mutex, std::try_to_lock);
		if (ulock.owns_lock()) {		// caught the mutex
			cout << "current g_i = " << g_i << endl;
			++j;
		}
	}
}

void minus() {
	int j = 0;
	while (j < 1000000) {
		mutex.lock();
		std::lock_guard<std::mutex> lock(mutex, std::adopt_lock); // auto unlock the mutex but must be lock the mutex before
		--g_i;
		++j;
	}
	
}   

void minus1() {
	int j = 0;
	while (j < 1000000) {
		std::lock_guard<std::mutex> lock(mutex);  // unlock the mutex when the lock_guard variable distruct
		--g_i;
		++j;
	}
	
}   


void fireworks() {
	while (!tmutex.try_lock_for(std::chrono::milliseconds(2000))) 
		cout << "-";
	std::this_thread::sleep_for(std::chrono::milliseconds(10000));
	cout << "*" << endl;
	tmutex.unlock();
}



int main(int argc, char **argv)
{
	// std::lock(mutex1, mutex2);		// lock more mutex at least two and need to manual unlock
	// std::try_lock(mutex1, mutex2);	// try to more mutex lock, at least two and need to manuan unlock
	//std::thread mythread(test1);
	//std::thread mythread1(test);
	//std::thread mythread2(add);
	//std::thread mythread3(add1);
	//std::thread mythread4(print);
	//mythread.join();
	//mythread1.join();
	//mythread2.join();
	//mythread3.join();
	//mythread4.join();
	//cout << g_i << endl;
	std::thread threads[2];
	for (int i = 0; i < 2; ++i)
		threads[i] = std::thread(fireworks);
	for (auto& th : threads)
		th.join();
	
	
	return 0;
}
