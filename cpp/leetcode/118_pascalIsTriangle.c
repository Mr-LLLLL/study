#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
	vector<vector<int>> generate(int numRows) {
		vector<vector<int>> triangle;
		if (0 == numRows)
			return triangle;
		triangle.push_back(vector<int>{1});
		for (int i = 2; i <= numRows; ++i) {
			vector<int> v;
			v.push_back(1);
			for (int j = 0; j < i - 2; ++j) 
				v.push_back(triangle[i - 2][j] + triangle[i - 2][j + 1]);
			v.push_back(1);
			triangle.push_back(v);

		}
		return triangle;
	}
};

void print(vector<vector<int>> const & vv) {
	for (vector<int> vi : vv) {
		for (int j = vv.size(); j > vi.size(); --j)
			cout << ' ';
		for (int i : vi)
			cout << i << ' ';
		cout << endl;
	}
}

int main(int argc, char** argv)
{
	int n;
	while (cout << "input number of the Pascal's Triangle: ", cin >> n)
		print(Solution().generate(n));


	return 0;
}

