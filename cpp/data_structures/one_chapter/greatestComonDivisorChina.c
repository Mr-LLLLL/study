#include <iostream>

using namespace std;

int greatestCommonDivisorChina(int a, int b)
{
	int r = 0;	//a and b conmondivisor as 2^r
	while( ! ((a & 1) || (b & 1))) {	//a and b is even
		a >>= 1;						
		b >>= 1;
		++r;
	}									//afterward no more one even 
	while(1) {
		while ( ! (a & 1)) a >>= 1;		//a is even,be divided 2
		while ( ! (b & 1)) b >>= 1;		//b is even,be divided 2
		(a > b) ? a = a - b : b = b - a;//as normally as gcd(max(a, b) - min(a,b), min(a, b))
		if (0 == a)						//as normally as gcd(0, b) = b;
			return b << r;
		if (0 == b)						//as normally as gcd(a, 0) = a;
			return a << r;
	}
}

int main(int argc, char** argv)
{
	int x, y;
	while(cout << "input two number: ", cin >> x >> y)
		cout << x << " and " << y << "gcdCN is: " << greatestCommonDivisorChina(x, y) << endl; 
	

	return 0;
}
