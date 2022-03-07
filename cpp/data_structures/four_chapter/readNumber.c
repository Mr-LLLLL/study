#include <iostream>
#include <string>
#include <stack>

using namespace std;

void readNumber(char* const& s, stack<double>& stk) {
	stk.push((float)(*p - '0'));
	while (isdigit(*(++p))) {
		double temp = stk.top();
		stk.pop();
		stk.push(temp * 10 + (*p - '0'));
	}
	if ('.' != *p)
		return;
	double fraction = 1;
	while (isdigit(*(++p)))
		double temp = stk.top();
		stk.pop();
		stk.push(temp + (*p - '0') * (fraction /= 10));
}
	
