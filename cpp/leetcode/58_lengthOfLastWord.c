#include <iostream>
#include <string>

using namespace std;

class Solution {
public:
	int lengthOfLastWord(string const& s){
		if (!s.size())
			return 0;
		int i = s.size() - 1;
		while (s[i] == ' ')
			--i;
		int count = 0;
		while (i >= 0 && s[i] != ' ') {
			++count;
			--i;
		}
		return count;
	}
};

int main(int argc, char** argv)
{
	cout << "please input strings:";
	string line;
	getline(cin, line);
	cout << Solution().lengthOfLastWord(line) << endl;

	return 0;
}
