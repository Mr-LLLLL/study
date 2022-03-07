#include <iostream>
#include <stack>
#include <vector>

using namespace std;

enum op{no, add, sub, multi, devide, space, opend};

class Solution {
public:
	void calcu(vector<op> const& v, int const& n) {
		stack<op> stk;
		stack<int> stkn;
		stk.push(no);
		size_t i = 0;
		int temp1, temp2;
		stkn.push(i);
		while (!stk.empty()) {
			if (stk.top() < v[i]) {
				stk.push(v[i]);
				stkn.push(++i);
			}
			else if(stk.top() == no && v[i] == no)
				stk.pop();
			else if(stk.top() == add) {
				stk.pop();
				temp1 = stkn.top();
				stkn.pop();
				temp2 = stkn.top();
				stkn.pop();
				stkn.push(temp1 + temp2);
			}
			else if(stk.top() == multi) {
				stk.pop();
				temp1 = stkn.top();
				stkn.pop();
				temp2 = stkn.top();
				stkn.pop();
				stkn.push(temp1 * temp2);
			}
			else if(stk.top() == space) {
				stk.pop();
				temp1 = stkn.top();
				stkn.pop();
				temp2 = stkn.top();
				stkn.pop();
				stkn.push(temp2 * 10 + temp1);
			}
			else if(stk.top() == sub) {
				stk.pop();
				temp1 = stkn.top();
				stkn.pop();
				temp2 = stkn.top();
				stkn.pop();
				stkn.push(temp2 - temp1);
			}
			else if(stk.top() == devide) {
				stk.pop();
				temp1 = stkn.top();
				stkn.pop();
				temp2 = stkn.top();
				stkn.pop();
				stkn.push(temp2 / temp1);
			}
		}
		i = 0;
		if (stkn.top() == n) {
			for (op iter : v) {
				switch (iter) {
					case add:
						cout << i++ << " + ";
						break;
					case multi:
						cout << i++ << " * ";
						break;
					case space:
						cout << i++;
						break;
					case sub:
						cout << i++ << " - ";
						break;
					case devide:
						cout << i++ << " / ";
						break;
				}
			}
			cout << i << endl;
		} 
	}
	
	void operatorV(int const& n) {
		vector<op> v(9, no);
		v[9] = no;
		for (op iter = add; iter < opend; iter = op(iter + 1)) {
			v[0] = iter;
			for (op iter = add; iter < opend; iter = op(iter + 1)) {
				v[1] = iter;
				for (op iter = add; iter < opend; iter = op(iter + 1)) {
					v[2] = iter;
					for (op iter = add; iter < opend; iter = op(iter + 1)) {
						v[3] = iter;
						for (op iter = add; iter < opend; iter = op(iter + 1)) {
							v[4] = iter;
							for (op iter = add; iter < opend; iter = op(iter + 1)) {
								v[5] = iter;
								for (op iter = add; iter < opend; iter = op(iter + 1)) {
									v[6] = iter;
									for (op iter = add; iter < opend; iter = op(iter + 1)) {
										v[7] = iter;
										for (op iter = add; iter < opend; iter = op(iter + 1)) {
											v[8] = iter;
											calcu(v, n);
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
};

int main(int argc, char** argv)
{
	int n;
	while (cout << "input number please:", cin >> n)
		Solution().operatorV(n);

	return 0;
}
