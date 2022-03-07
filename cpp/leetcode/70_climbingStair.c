#include <iostream>

using namespace std;

class Solution {
public:
	int climbStairs(int n) {
		int g = 0,
			f = 1;
		while (0 < n--) {
			f = f + g;
			g = f - g;
		}
		return f;
	}
};

int main(int argc, char** argv)
{
	int n;
	while (cout << "input number of stair's floor please: ", cin >> n)
		cout << Solution().climbStairs(n) << endl;
	return 0;
}
