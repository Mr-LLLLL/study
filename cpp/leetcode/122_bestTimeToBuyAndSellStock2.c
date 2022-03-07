#include <vector>
#include <iostream>
#include <climits>

using namespace std;

class Solution {
public:
	int maxProfit(vector<int>& prices) {
		if (!prices.size())
			return 0;
		int minPrice = prices[0];
		int profit = 0;
		for (int i = 1; i < prices.size(); ++i)
			if (prices[i - 1] <= prices[i])
				continue;
			else {
				profit += prices[i - 1] - minPrice;
				minPrice = prices[i];
			}
		return profit += prices[prices.size() - 1] - minPrice;
	}
};

int main(int argc, char** argv)
{
	int price;
	while (cout << "input the price of the stock everyday: " << endl) {
		vector<int> v;
		while (cin >> price)
			v.push_back(price);
		cin.clear();
		cin.ignore();
		cout << "the stock max profit is :" << Solution().maxProfit(v) << endl;
	}
	return 0;
}
