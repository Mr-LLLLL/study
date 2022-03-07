#include <exception>
#include <iostream>

using namespace std;

class ListNode{
public:
	ListNode(int i = 0) : ival(i), next(nullptr) {}
	int ival;
	ListNode *next;
};

class Solution {
public:
	ListNode* listMergeSort(ListNode *p) {
		if (p == nullptr && p->next == nullptr)
			return p;
		return mergeSort(p);
	}
	ListNode* mergeSort(ListNode *p) {
		if (p->next == nullptr)
			return p;
		ListNode *pre, *mi, *tail;
		mi = tail = p;
		pre = nullptr;
		while (tail != nullptr && tail->next != nullptr) {
			tail = tail->next;
			if(tail->next)
				tail = tail->next;
			pre = mi;
			mi = mi->next;
		}
		pre->next = nullptr;
		ListNode *left, *right;
		left = mergeSort(p);
		right = mergeSort(mi);
		return merge(left, right);
	}
	ListNode *merge(ListNode *left, ListNode *right) {
		ListNode *pRes = new ListNode(0);
		ListNode *temp = pRes;
		while (left != nullptr && right != nullptr) {
			if (left->ival < right->ival) {
				temp->next = left;
				temp = temp->next;
				left = left->next;
			}
			else {
				temp->next = right;
				temp = temp->next;
				right = right->next;
			}
		}
		if (left != nullptr)
			temp->next = left;
		else if (right != nullptr)
			temp->next = right;
		temp = pRes->next;
		delete pRes;
		return temp;
	}
};

int main(int argc, char** argv)
{
	ListNode *head = new ListNode(0);
	ListNode *test = head;
	for (size_t  i = 11; i > 0; --i) {
		test->next = new ListNode(i);
		test = test->next;
	}
	ListNode *head1 = new ListNode(0);
	test = head;
	while (test) {
		cout << test->ival << ", ";
		test = test->next;
	}
	cout << endl;
	test = head;
	Solution().listMergeSort(test);

	while (test) {
		cout << test->ival << ", ";
		test = test->next;
	}
	cout << endl;
	test = head;
	while(test) {
		head = test->next;
		delete test;
		test = head;
	}


	return 0;
}
