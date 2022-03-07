#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	vector<int> getRow(int rowIndex) {
		if (rowIndex < 0)
			return vector<int>();
		vector<int> res(rowIndex + 1, 0);
		res[0] = 1;
		for (int i = 1; i <= rowIndex; ++i)
			for (int j = i; j >= 1; --j) 
				res[j] += res[j - 1];
		return res;
	}
};

void print(vector<int> const & v) {
		for (int i : v)
			cout << i << ' ';
		cout << endl;
}

int main(int argc, char** argv)
{
	int n;
	while (cout << "input number of the Pascal's Triangle: ", cin >> n)
		print(Solution().getRow(n));


	return 0;
}

