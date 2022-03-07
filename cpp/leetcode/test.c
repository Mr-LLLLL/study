#include <iostream>
#include <vector>
#include <stack>
#include <string>
#include <queue>

using namespace std;

int main()
{
	priority_queue<int, vector<int>, greater<int>> q;
	q.push(1);
	q.push(2);
	q.push(3);
	cout << q.top() << endl;
	return 0;
}




