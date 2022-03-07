#include <iostream>
#include <functional>
#include <vector>
#include <list>
#include <algorithm>
#include <utility>

using std::cin;
using std::cout;
using std::endl;



//	make list sorted with insert sorted
template<typename T, typename F = std::less<T>>
class Sorted_list {
public:
	void insert(const T &t, F f = F()) {
		/* STL find_if and std::less, std::greater algorithm support first come first served */
		auto iter = find_if(_list.begin(), _list.end(), bind(f, std::placeholders::_1, t));
		_list.insert(iter, t);
	}
	/* make a deal */
	void deal_order() {
		_list.pop_front();
	}

	/* support range for statement */
	typename std::list<T>::iterator begin() {
		return _list.begin();
	}
	typename std::list<T>::iterator end() {
		return _list.end();
	}

private:
	std::list<T> _list;
};

int main(int artc, char** argv)
{
	// equal multiset<int, std::less<int>>, time: O(logn)
	Sorted_list<int, std::less<int>> buy_store;	// time: O(n)

	/* buy_store test */
	for (int i = 0; i < 10; i++)
		buy_store.insert(i);
	buy_store.insert(20);
	buy_store.insert(5);
	buy_store.deal_order();
	buy_store.deal_order();

	cout << "buy_store test result:" << endl;
	for (auto &i : buy_store)
		cout << i << endl;

	/* sell_store test */
	Sorted_list<int, std::greater<int>> sell_store;
	for (int i = 0; i < 10; i++)
		sell_store.insert(10 - i);
	sell_store.insert(3);
	sell_store.insert(20);
	sell_store.deal_order();
	sell_store.deal_order();

	cout << "\n\nsell_store test result:" << endl; 

	for (auto &i : sell_store)
		cout << i << endl;


	return 0;
}

