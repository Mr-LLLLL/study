#include <iostream>
#include <vector>

using namespace std;

class Solution {
	public:
		vector<int> plusOne(vector<int> const &digits) {
			size_t rank = digits.size();
			vector<int> vec;
			int sign = 1;
			while (rank--){
				if (digits[rank] == 9)	
					vec.push_back(0);
				else {
					vec.push_back(digits[rank] + 1);
					while(rank--)
						vec.push_back(digits[rank]);
					sign = 0;
					break;
				}
			}
			vector<int> ret;
			if (sign)
				ret.push_back(1);
			for (int i = vec.size() - 1; i >= 0; --i)
				ret.push_back(vec[i]);


			return ret;
		}	
};


int main(int argc, char** argv)
{
	vector<int> vec;
	int number;
	cout << "please input vector<int>:";
	while (cin >> number)
		vec.push_back(number);
	for (int num : Solution().plusOne(vec))
		cout << num << ", ";
	cout << endl;


	return 0;
}
