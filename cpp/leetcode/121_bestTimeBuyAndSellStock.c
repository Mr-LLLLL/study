#include <vector>
#include <iostream>
#include <climits>

using namespace std;

class Solution {
public:
	int maxProfit(vector<int>& prices) {
		int minPrice = INT_MAX;
		int maxProfit = 0;
		for (int i = 0; i < prices.size(); ++i)
			if (prices[i] < minPrice)
				minPrice = prices[i];
			else if (prices[i] - minPrice > maxProfit)
				maxProfit = prices[i] - minPrice;
		return maxProfit;
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
		cin.ignore(1,EOF);
		cout << "the stock max profit is :" << Solution().maxProfit(v) << endl;
	}
	return 0;
}
