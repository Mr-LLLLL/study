#include <string>
#include <iostream>
#include <stack>

using namespace std;

class Solution {
public:
	bool isparen(string const &s){
		stack<char> sta;
		for (size_t i = 0; i < s.size(); ++i){
			switch (s[i]){
				case '(':
				case '[':
				case '{':
					sta.push(s[i]);
					break;
				case ')':
					if (sta.empty() || sta.top() != '(')
						return false;
					else
						sta.pop();
					break;
				case '}':
					if (sta.empty() || sta.top() != '{')
						return false;
					else
						sta.pop();
					break;
				case ']':
					if (sta.empty() || sta.top() != '[')
						return false;
					else
						sta.pop();
					break;
				default:
					break;
			}
		}
		return sta.empty();
	}
};

int main(int argc, char** argv)
{
	string s;
	while (cin >> s)
		cout << Solution().isparen(s) << endl;


	return 0;
}
