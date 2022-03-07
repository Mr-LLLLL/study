#include <cstring>

#define M 97
#define R 256
#define DIGIT(S, i) ((S)[i] - '0')

typedef long long HashCode;

bool check1by1(char* P, char* T, int i);
HashCode prepareDm(int m);
void updateHash(HashCode& hashT, char* T, int m, int k, HashCode Dm);

int match(char* P, char* T) {
	int m = strlen(P), n = strlen(T);
	HashCode Dm = prepareDm(m);
	HashCode hashP = 0;
	HashCode hashT = 0;
	for (int i = 0; i < m; ++i) {
		hashP = (hashP * R + DIGIT(P, i)) % M;
		hashT = (hashT * R + DIGIT(P, i)) % M;
	}
	for (int k = 0;;) {
		if ((hashT == hashP))
			if (check1by1(P, T, k))
				return k;
		if (++k > n - m)
			return k;
		else
			updateHash(hashT, T, m, k, Dm);
	}
}

				

bool check1by1(char* P, char* T, int i) {
	for (int m = strlen(P), j = 0; j < m; ++j, ++i)
		if (P[j] != T[i])
			return false;
	return true;
}

void updateHash(HashCode hashT, char* T, int m, int k, HashCode Dm) {
	hashT = (hashT - DIGIT(T, k - 1) * Dm) % M;
	hashT = (hashT * R + DIGIT(T, k + m - 1)) % M;
	if (0 > hashT)
		hashT += M;
}

HashCode prepareDm (int m) {
	HashCode Dm = 1;
	for (int i = 1; i < m; ++i)
		Dm = (R * Dm) % M;
	return Dm;
}

