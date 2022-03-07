#include <iostream>
#include <stack>

using namespace std;

class MinStack {
private:
	stack<int> s;
	stack<int> minStack;
public:
	MinStack() {}
	void push(int x) {
		if (s.empty()) {
			s.push(x);
			minStack.push(x);
		} else {
			s.push(x);
			if (x > minStack.top())
				minStack.push(minStack.top());
			else
				minStack.push(x);
		}
	}

	void pop() {
		s.pop();
		minStack.pop();
	}

	int top() {
		return s.top();
	}

	int getMin() {
		return minStack.top();
	}
};

int main(int argc, char** argv)
{
	MinStack s;
	s.push(1);
	s.push(-2);
	s.push(2);
	cout << s.getMin() << endl;
	s.pop();
	cout << s.getMin() << endl;
	s.pop();
	cout << s.getMin() << endl;
	s.pop();
	s.push(4);
	cout << s.getMin() << endl;
	return 0;
}
