#include <iostream>
#include <typeinfo>
#include <stdio.h>
#include <vector>
#include <unistd.h>
#include <algorithm>
#include <list>
#include <fstream>
#include <fcntl.h>
#include <utility>
#include <memory>

using namespace std;

class test;

class base {
public:
    virtual ~base() {
        cout << "destroy base" << endl;
    }
    shared_ptr<test> sp;
    int i = 8;
    
};

class test {
public:
    test() {
    }
    test(int i, int j) : _i(i) {
        cout << "test(int, itn)" << endl;
    }

    ~test() {
        cout << "destroy test" << endl;
    }
    int _i;
};

class derive : public base {
public:
	void print() {
		cout << "derived print()" << endl;
	}
};

template <typename Signature>
class SignalTrivial;

template <typename RET, typename... ARGS>
class SignalTrivial<RET(ARGS...)>
{
public:
    using Functor = function<void (ARGS...)>;

    void connect(Functor&& func)
    {
        functors_.push_back(std::forward<Functor>(func));
    }

    void call(ARGS&&... args)
    {
        for (const Functor& f : functors_)
        {
            f(args...);
        }
    }

    void print()
    {
        cout << sizeof...(ARGS) << endl;
    }
private:
    vector<Functor> functors_;
};

int main(int argc, char *argv[])
{
    base b;
    b.sp = make_shared<test>();

    shared_ptr<test> sp(new test(4, 2));
    weak_ptr<test> wp = sp;
    shared_ptr<test> sp1(new test(1, 2));
    wp = sp1;
    cout << (*sp)._i << endl;

    cout << 5 % 2 << endl;



		
	return 0;
}


