#include <iostream>
#include <thread>
#include <functional>

using std::cin;
using std::cout;
using std::endl;


class Obj {
public:
	Obj() {
		cout << "Obj() thread id = " << std::this_thread::get_id() << endl;
	}
	Obj(const Obj&) {
		cout << "Obj(const Obj&) thread id = " << std::this_thread::get_id() << endl;
	}
	void print() {
		cout << "thread id = " << std::this_thread::get_id() << " starting..." << endl;
		cout << "thread id = " << std::this_thread::get_id() << " finilied" << endl;
	}
	~Obj() {
		cout << "~Obj() thread id = " << std::this_thread::get_id() << endl;
	}
};

void func(const Obj &) {
	cout << "thread id = " << std::this_thread::get_id() << " starting..." << endl;
	cout << "thread id = " << std::this_thread::get_id() << " finilied" << endl;
}

int main(int argc, char **argv)
{
	Obj obj;
	std::thread mythread(&Obj::print, &obj);	// use class function to call a thread (call to reference)
	/*
	 * if use call to value, linux system will be copy two object(windows system copy one), one of them like
	 * bind function, and another be create to the thread, and thread will be destruct this one;
	 */
	std::thread mythread1(func, obj);		// use function to call a thread (call to value)
	std::thread mythread2([]() {		// use lambda to call a thread
			cout << "thread id = " << std::this_thread::get_id() << " starting..." << endl;
			cout << "thread id = " << std::this_thread::get_id() << " finilied" << endl;
	});
	std::function<void(const Obj &)> f(func);
	std::thread mythread3(f, std::ref(obj));	// use function object to call a thread (call to reference)
	mythread.join();	// blocking the main thread until mythread finished 
	mythread1.join();	
	mythread2.detach(); // not block the main thread, and detach the mythread(don't use join after detach)
	mythread3.detach();	
	
	


	return 0;
}
