#include <iostream>

using namespace std;

class Solution {
public:
    void fermatLagrange(unsigned int const& number) {
        size_t count = 0;
        size_t n = number;
        //for (unsigned int n = 0; n <= number; ++n)
        for (unsigned int one = 0; one * one<= n; ++one)
            for (unsigned int two = one; two * two <= n; ++two)
                for (unsigned int three = two; three * three <= n; ++three)
                    for (unsigned int four = three; four * four<= n; ++four)
                        if (n == one * one + two * two + three * three + four * four)
                            //                      ++count;
                            cout << one << ", " << two << ", " << three << ", " << four << endl;
        cout << count << endl;
    }
};


int main(int argc, char** argv)
{
    unsigned int n;
    while (cout << "input number please:", cin >> n)
        Solution().fermatLagrange(n);
    return 0;
}
            
