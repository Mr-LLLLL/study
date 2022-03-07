#include <iostream>

using namespace std;

struct ListNode {
	int val;
	ListNode *next;
	ListNode(int x) : val(x), next(nullptr) {}
};

class Solution {
public:
	bool hasCycle(ListNode *head) {
		if (nullptr == head || nullptr == head->next)
			return false;
		ListNode* slow = head;
		ListNode* fast = head->next;
		while (slow != fast) {
			if (nullptr == fast->next || nullptr == fast->next->next)
				return false;
			slow = slow->next;
			fast = fast->next->next;
		}
		return true;
	}
};

ListNode* getList(int posi) {
	int n;
	cout << "input List please: " << endl;
	cin >> n;
	ListNode* node;
	node = new ListNode(n);
	ListNode* head = node;
	while (cin >> n) {
		node->next = new ListNode(n);
		node = node->next;
	}
	if (-1 != posi) {
		ListNode* temp = head;
		while (posi--) {
			temp = head->next;
		}
		node->next = temp;
	}

	return head;
}
	


int main(int argc, char** argv)
{
	int posi;
	cout << "input Posi please: " << endl;
	cin >> posi;
	cout << Solution().hasCycle(getList(posi)) << endl;

	return 0;
}
