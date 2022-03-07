#include <iostream>
#include <string>

using namespace std;

struct Solution {
	string addBinary(string const& a, string const& b) {
		int a_size = a.size();
		int b_size = b.size();
		string s = "";
		char add = 0;
		char taila, tailb, tail;
		for (int i = 0; i < a_size || i < b_size; ++i) {
			taila = (a_size - 1 - i >= 0) ? (a[a_size - 1 - i] - '0') : 0;
			tailb = (b_size - 1 - i >= 0) ? (b[b_size - 1 - i] - '0') : 0;
			tail = (taila + tailb + add) % 2 + '0';
			s = tail + s;
			add = (taila + tailb + add) / 2;
		}
		if (1 == add)
			s = '1' + s;
		return s;
	}
};

int main(int argc, char** argv) 
{
	string s1, s2;
	while (cout << "input two Binary number:" << endl, cin >> s1 >> s2)
		cout << Solution().addBinary(s1, s2) << endl;

	return 0;
}
