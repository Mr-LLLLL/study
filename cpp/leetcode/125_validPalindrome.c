#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
	bool isPalindrome(string const& s) {
		for (int i = 0, j = s.size() - 1; i < j; ++i, --j) {
			while (!isvalid(s[i]) && i < j)
				++i;
			while (!isvalid(s[j]) && j > i)
				--j;
			if (!equal(s[i], s[j]))
				return false;
		}
		return true;
	}
	bool equal(char const& ch1, char const& ch2) {
		if (ch1 == ch2)
			return true;
		else if (ch1 - ch2 == 32 || ch1 - ch2 == -32)
			if (ch1 >= 'A' && ch2 >= 'A')
				return true;
		return false;
	}
	bool isvalid(char const& c) {
		if ((c >= '0' && c <= '9') || (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z'))
			return true;
		return false;
	}
};

int main(int argc, char** argv)
{
	string s;
	while (cout << "input string :", cin >> s)
		cout << Solution().isPalindrome(s) << endl;
	return 0;
}
