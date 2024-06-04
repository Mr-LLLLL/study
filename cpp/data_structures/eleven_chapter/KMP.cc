#include <cstring>

int* buildNext(char*);

int match(char* P, char* T) {
	int* next = buildNext(P);
	int n = strlen(T), i = 0;
	int m = strlen(P), j = 0;
	if (!m)
		return 0;
	if (!n || m > n)
		return -1;
	while (j < m && i < n)
		if (j < 0 || P[j] == T[i]) {
			++i;
			++j;
		} else 
			j = next[j];
	delete [] next;
	return i - j > n - m ? - 1 : i - j;
}
int* buildNext(char* P) {
	int m = strlen(P), j = 0;
	int* N = new int[m];
	int t = N[0] = -1;
	while (j < m - 1)
		if (t < 0 || P[j] == P[t]) {
			++j;
			++t;
			N[j] = P[j] == P[t] ? N[t] : t;
		} else
			t = N[t];


	return N;
}
