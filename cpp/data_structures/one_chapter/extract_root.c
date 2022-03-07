#include <iostream>

using namespace std;

int G(int n, int k) {           //if k == 0; the function is extract rooting function
    return n < 1 ? k : G(n - 2 * k - 1, k + 1);
}

int main(int argc, char** argv)
{
    int n, k;
    while(cout << "please input one number: ", cin >> n >> k)
        cout << "the result is :" << G(n, k) << endl;
    
    return 0;
}
