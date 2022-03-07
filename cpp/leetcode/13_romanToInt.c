#include <string>
#include <iostream>
#include <map>

using namespace std;

int romanToInt(string s) {
	int romanNumber[256];
	romanNumber['I'] = 1;
	romanNumber['V'] = 5;
	romanNumber['X'] = 10;
	romanNumber['L'] = 50;
	romanNumber['C'] = 100;
	romanNumber['D'] = 500;
	romanNumber['M'] = 1000;
	romanNumber[0] = 0;
	int sum = 0;
	for (size_t i = 0; i < s.size(); ++i)
		if (romanNumber[ (int)s[i + 1]] > romanNumber[ (int)s[i]]) {
			sum += romanNumber[ (int)s[i + 1]] - romanNumber[ (int)s[i]];
			++i;
		}
		else
			sum += romanNumber[ (int)s[i]];


	return sum;
}

int main(int argc, char** argv)
{
	string s;
	while (cout << "please input roman number:", cin >> s)
		cout << "normal number is: " << romanToInt(s) << endl;

	return 0;
}
