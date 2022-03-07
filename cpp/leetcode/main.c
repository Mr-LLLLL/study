#include <map>
#include <iostream>
#include <vector>
#include <sstream>
#include <string>
#include <algorithm>
#include <ctype.h>


using std::vector;
using std::string;
using std::cin;
using std::cout;
using std::endl;
using std::to_string;
using std::runtime_error;
using std::stringstream;
using std::map;

void trimLeftTraillingSpaces(string &input);
void trimRightTraillingSpaces(string &input);
vector<int> stringToIntegerVector(string input);
int stringToInteger(string input);
string integerVectorToString(vector<int> list, int length); 

void trimLeftTrailingSpaces(string &input) {
	input.erase(input.begin(), find_if(input.begin(), input.end(), [](int ch) {
				return !isspace(ch);
				}));
}

void trimRightTrailingSpaces(string &input) {
	input.erase(find_if(input.rbegin(), input.rend(), [](int ch) {
				return !isspace(ch);
				}).base(), input.end());
}

vector<int> stringToIntegerVector(string input) {
	vector<int> output;
	trimLeftTrailingSpaces(input);
	trimRightTrailingSpaces(input);
	input = input.substr(1, input.length() - 2);
	stringstream ss;
	ss.str(input);
	string item;
	char delim = ',';
	while (getline(ss, item, delim)) {
		output.push_back(stoi(item));
	}
	return output;
}

int stringToInteger(string input) {
	return stoi(input);
}

string integerVectorToString(vector<int> list, int length = -1) {
	if (length == -1) {
		length = list.size();
	}

	if (length == 0) {
		return "[]";
	}

	string result;
	for(int index = 0; index < length; ++index) {
		int number = list[index];
		result += to_string(number) + ", ";
	}
	return "[" + result.substr(0, result.length() - 2) + "]";
}

int main(int argc, char** argv)
{
	string line;
	while (getline(cin, line)) {
		vector<int> nums = stringToIntegerVector(line);
		getline(cin, line);
		int target = stringToInteger(line);

		vector<int> ret = Solution().twoSum(nums, target);

		string out = integerVectorToString(ret);
		cout << out << endl;
	}

	return 0;
}

