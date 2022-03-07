#include <iostream>
#include <string>
#include <sstream>
#include <vector>

using namespace std;

class Solution {
public:
	string longestCommonPrefix(const vector<string> &strs) {
		if (strs.size() == 0)
			return "";
		return longestCommonPrefix(strs, 0, strs.size() - 1);
	}
private:
	string commonPrefix(const string left, const string right) {
		int min = left.size() < right.size() ? left.size() : right.size();
		for (int i = 0; i < min; ++i) {
			if (left[i] != right[i])
				return left.substr(0, i);
		}
		return left.substr(0, min);
	}

	string longestCommonPrefix(const vector<string> &strs, const int lo, const int hi){
		if (hi == lo)
			return strs[lo];
		int mi = (hi + lo) >> 1;
		string lcpLeft = longestCommonPrefix(strs, lo, mi);
		string lcpRight = longestCommonPrefix(strs, mi + 1, hi);
		return commonPrefix(lcpLeft, lcpRight);
	}
};

int main(int argc, char** argv)
{
	string line, word;
	stringstream ss;
	vector<string> strs;
	cout << "please input some words:" << endl;
	getline(cin, line);
	ss << line;
	while (ss >> word)
		strs.push_back(word);
	cout << "the words common prefix is: " << Solution().longestCommonPrefix(strs) << endl;

	return 0;
}
