#include <iostream>
#include <sstream>
#include <algorithm>
#include <ctype.h>
#include <string>
#include <vector>

using namespace std;

struct ListNode {
	int val;
	ListNode* next;
	ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
public:
	ListNode* deleteDuplicates(ListNode* head) {
		if (!head)
			return head;
		ListNode* p = head;
		ListNode* q;
		while (q = p->next) {
			if (p->val != q->val)
				p = q;
			else {
				p->next = q->next;
				delete q;
			}
		}
		return head;
	}
};

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

ListNode* stringToListNode(string input) {
	// Generate list from the input
	vector<int> list = stringToIntegerVector(input);

	// Now convert that list into linked list
	ListNode* dummyRoot = new ListNode(0);
	ListNode* ptr = dummyRoot;
	for (int iter : list) {
		ptr->next = new ListNode(iter);
		ptr = ptr->next;
	}
	ptr = dummyRoot->next;
	delete dummyRoot;
	return ptr;
}

int main(int argc, char** argv)
{
	string line;
	while (getline(cin, line)) {
		ListNode* head = stringToListNode(line);

		ListNode* ret = Solution().deleteDuplicates(head);

		do {
			cout << ret->val << ' ';
		} while (ret = ret->next);
		cout << endl;
	}
		

	return 0;
}
