#include <iostream>

using namespace std;

void reverse(int *A, int lo, int hi) {	//array reverse recursive version
	if (lo < hi) {
		swap(A[lo], A[hi]);				//exchange A[lo] and A[hi]
		reverse(A, lo + 1, hi - 1);		//recursive reverse
	}
}

int shift(int *A, int n, int k) {	//shift left k bit 
	k %= n;							//ensure k <= n
	reverse(A, 0, k - 1);					//exchange A[k, n) O(3k/2)
	reverse(A, k, n - 1);			//exchange A[k, n) O(3(n-k)/2)
	reverse(A, 0, n - 1);					//exchange A[0, n) O(3n/2)

	return 3 * n;					//return O(3/2*(k+(n-k)+n)=O(3n)
}

int main(int argc, char** argv)
{
	int arry[] = {1, 2, 3, 4, 5, 6};
	shift(arry, 6, 2);
	for(int iter : arry)
		cout << iter << ", ";
	cout << endl;


	return 0;
}
