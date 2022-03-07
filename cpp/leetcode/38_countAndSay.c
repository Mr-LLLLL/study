#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
	string countAndSay(int const& n) {
		string str("1");
		string ch;
		int count = 1, i = n;
		while (--i) {
			for (string::size_type j = 0; j < str.size(); ++j) {
				if (str[j] == str[j + 1])
					++count;
				else {
					ch = ch + to_string(count) + str[j];
					count = 1;
				}
			}
			str = ch;
			ch = "";
		}

		return str;
	}
};

int main(int argc, char** argv)
{
	int n;
	while (cout << "please input number n:", cin >> n)
		cout << Solution().countAndSay(n) << endl;

	return 0;
}
