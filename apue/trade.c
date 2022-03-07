#include <iostream>
#include <typeinfo>
#include <fstream>
#include <unordered_map>
#include <string>
#include <sstream>
#include <vector>
#include <set>

#define NDEBUG

#include <cassert>

using std::cin;
using std::cout;
using std::endl;
using std::cerr;

/* delimit in order file */
char delim = ',';

/* order number */
unsigned int orderNO = 0;

/* stock's type */
enum class stock_type {BUY = 1, SELL = 2}; 

/* Order class */
class Order {
public:
	Order() = delete;

	Order(const Order&) = default;

	Order(unsigned int o, std::string c, double p, unsigned int n, stock_type t) : 
		order(o), code(c), price(p), number(n), type (t) {}
								
	Order(const Order&& arg) {
		order = std::move(arg.order);
		code = std::move(arg.code);
		price = std::move(arg.price);
		number = std::move(arg.number);
		type = std::move(arg.type);
	}
	/* the higher the price, the lower the priority */
	/* first come, first served */
	/* sell_store */
	bool operator <(const Order& arg) const {
		return price < arg.price;
	}
	/* the higher the price, the higher the priority */
	/* first come, first served */
	/* buy_store */
	bool operator >(const Order& arg) const {
		return price > arg.price;
	}

	/* member */
	unsigned int order;
	std::string code;
	double price;
	mutable unsigned int number;
	stock_type type;
};

std::unordered_map<std::string, std::multiset<Order, std::less<Order> > > sell_store;
std::unordered_map<std::string, std::multiset<Order, std::greater<Order> > > buy_store;


void getOrders(std::string file, std::vector<Order> &orders);
void trade_Contract(const Order&, const Order&, unsigned int);
void deal_order(const Order& order);
void read_repository();
void write_repository();
void get_orderNO();
void set_orderNO();

int main(int argc, char **argv)
{
	// save current order
	std::vector<Order> orders;

	// read orderNO
	get_orderNO();

	// read buy_store and sell_store
	read_repository();

	for (int i = 1; i < argc; ++i) {
		// read current order
		getOrders(argv[1], orders);
	}

	// order transaction
	for (auto &order : orders) {
		deal_order(order);
	}

	// save buy_sotre and sell_store to file
	write_repository();
	// set current orderNO
	set_orderNO();

	return 0;
}

void get_orderNO() {
	std::ifstream read("orderNO");
	std::string temp;
	read >> temp;
	if (!temp.empty())
		orderNO = stoul(temp);
	assert(cout << "orderNO = " << orderNO << endl);
}

void set_orderNO() {
	std::ofstream write("orderNO");
	write << orderNO << endl;
}

void read_repository() {
	std::vector<Order> orders;
	/* get the 	sell_store data */
	std::fstream is_exist("sell_store", std::fstream::in);
	if (!is_exist)				// if isn't exist, create it
		is_exist.open("sell_store", std::fstream::out);
	is_exist.close();
	getOrders("sell_store", orders);
	for (auto &order : orders)
		sell_store[order.code].insert(order);

	orders.clear();
	
	/* get the buy_store data */
	is_exist.open("buy_store", std::fstream::in);
	if (!is_exist)				// if isn't exist, create it
		is_exist.open("buy_store", std::fstream::out);
	getOrders("buy_store", orders);
	for (auto &order : orders)
		buy_store[order.code].insert(order);
}
	
void write_repository() {
	/* output sell store to sell_store file */
	std::ofstream sell_output("sell_store", std::ofstream::trunc);
	for (auto &orders : sell_store)
		for (auto &order : orders.second)
			sell_output << order.code << "," << order.price << "," << order.number << ",2," << order.order << endl;

	/* output buy store to buy_store file */
	std::ofstream buy_output("buy_store", std::ofstream::trunc);
	for (auto &orders : buy_store)
		for (auto &order : orders.second)
			sell_output << order.code << "," << order.price << "," << order.number << ",2," << order.order << endl;
}


void deal_order(const Order& order) {
	if (order.type == stock_type::BUY) {		// deal buy order
		auto iter = sell_store.find(order.code);
		if (iter == sell_store.end()){			// if sell_store without the code
			buy_store[order.code].insert(order);
		} else if (order.price < sell_store[order.code].begin()->price) {	// if the order's price lower than the sell_store lowerest price
			buy_store[order.code].insert(order);
		} else {
			auto &set = iter->second;
			while (!set.empty() && order.price >= set.begin()->price) {	// sell_store is not empty and order's price higher the lowerest price
				int remaind = order.number - set.begin()->number;
				if (remaind >= 0) {											// if remaind greate eqaul 0 mean the that order sell out in the store
					trade_Contract(order, *set.begin(), (*set.begin()).number);
					set.erase(set.begin());								
					order.number = remaind;
					if (remaind == 0) 
						break;
				} else {					// sell_store residue number 
					set.begin()->number = -remaind;
					trade_Contract(order, *set.begin(), order.number);
					order.number = 0;		// clear the current number of the order
					break;
				}
			}
			if (set.empty()) {			// if the code of the sell_store and clear it
				sell_store.erase(iter);
			}	
			if (order.number) {			// if current order remaind, save it to buy_store
				buy_store[order.code].insert(order);
			}
		} // else
	} else {	/* same as above */
		auto iter = buy_store.find(order.code);
		if (iter == buy_store.end()) {
			sell_store[order.code].insert(order);
		} else if (order.price > buy_store[order.code].begin()->price ) {
			sell_store[order.code].insert(order);
		} else {
			auto &set = iter->second;
			while (!set.empty() && order.price <= set.begin()->price) {
				int remaind = order.number - set.begin()->number;
				if (remaind >= 0) {
					trade_Contract(order, *set.begin(), (*set.begin()).number);
					set.erase(set.begin());
					order.number = remaind;
					if (remaind == 0) 
						break;
				} else {
					set.begin()->number = -remaind;
					trade_Contract(order, *set.begin(), order.number);
					order.number = 0;
					break;
				}
			}
			if (set.empty()) {
				buy_store.erase(iter);
			}
			if (order.number) {
				sell_store[order.code].insert(order);
			}
		} // else
	} // else
}

void getOrders(std::string file, std::vector<Order> &orders) {
	/* read the orders */
	std::ifstream in(file, std::ifstream::in);
	std::string line;
	std::string temp;

	unsigned int num;
	stock_type type;
	double price;
	std::string code;
	unsigned int order;

	if (!in)
		cerr << "can't open file: " << file << endl;

	
	while (getline(in, line)) {		// read the current order
		for (auto &ch : line)
			if (ch == delim)
				ch = ' ';
		std::istringstream is(line);
		int cnt = 0;
		/* require the orderNO, code, price, num, type of the current order */
		while (is >> temp) {
			++cnt;
			switch (cnt) {
				case 1:
					code = temp;
					break;
				case 2:
					price = stod(temp);
					break;
				case 3:
					num = stoul(temp);
					break;
				case 4:
					if (stoul(temp) == 1)
						type = stock_type::BUY;
					else
						type = stock_type::SELL;
					break;
				case 5:
					order = stoul(temp);
					break;
			}
		}
		if (cnt == 4) {				// create a new order from the new file
			++orderNO;
			orders.emplace_back(orderNO, code, price, num, type);
		} else if (cnt == 5) {		// read the order from buy_store and sell_store
			orders.emplace_back(order, code, price, num, type);
		}
#ifndef NDEBUG
		Order order = orders.back();
		cout << "current order: " << order.order << " " << order.code << " " << order.price << " " << order.number << " " << 
			(order.type == stock_type::BUY ? "buy" : "sell") << endl;
#endif
	}
}
	
void trade_Contract(const Order& first, const Order& second, unsigned int num) {
	std::ofstream output("trade_constract", std::ofstream::app);
	if (!output)
		cerr << "can't open trade_constract file" << endl;
	output << "Order " << first.order << " - Order " << second.order << ": 以 " << second.price << " 的价格成交 " << num << " 手" << endl;
	assert(cout << "Order " << first.order << " - Order " << second.order << ": 以 " << second.price << " 的价格成交 " << num << " 手" << endl);
	
}

