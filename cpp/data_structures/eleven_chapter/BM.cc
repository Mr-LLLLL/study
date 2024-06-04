#include <cstring>
#include <iostream>
#include <climits>

int* buildBC(char*);
int* buildGS(char*);
int* buildSS(char*);

int match(char* P, char* T) {
	int* bc = buildBC(P);
	int* gs = buildGS(P);
	int n = strlen(T), i = 0;
	int m = strlen(P);
	while (i + m <= n) {
		int j = m - 1;
		while (P[j] == T[i + j])
			if (0 > --j)
				break;
			else
				i += std::max(gs[j], j - bc[T[i + j]]);
	}
	delete [] gs;
	delete [] bc;
	return i;
}

int* buildBC(char* P) {
	int* bc = new int[256];
	for (int j = 0; j < 256; ++j)
		bc[j] = -1;
	for (int m = strlen(P), j = 0; j < m; ++j)
		bc[P[j]] = j;
	return bc;
}

int* buildSS(char* P) {
	int m = strlen(P);
	int* ss = new int[m];
	ss[m - 1] = m;
	for (int lo = m - 1, hi = m - 1, j = lo - 1; j >= 0; --j)
		if (lo < j && ss[m - 1 - hi + j] <= j - lo)
			ss[j] = ss[m - 1 - hi + j];
		else {
			hi = j;
			lo = std::min(hi, lo);
			while (0 <= lo && P[m - 1 - hi + lo])
				--lo;
			ss[j] = hi - lo;
		}
	return ss;
}

int* buildGS(char* P) {
	int* ss = buildSS(P);
	int m = strlen(P);
	int* gs = new int[m];
	for (int j = 0; j < m; ++j)
		gs[j] = m;
	for (int i = 0, j = m - 1; j < UINT_MAX; --j)
		if (j + 1 == ss[j])
			while (i < m - j - 1)
				gs[i++] = m - j - 1;
	for (int j = 0; j < m - 1; ++j)
		gs[m - ss[j] - 1] = m - j - 1;
	delete [] ss;
	return gs;
}
