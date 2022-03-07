#include <iostream>
#include <algorithm>
#include <vector>

using namespace std;

struct Queen {
    int x, y;
    Queen(int xx = 0, int yy = 0) : x(xx), y(yy) {};
    bool operator==(Queen const &q) const {
        return (x == q.x)
                ||(y == q.y)
                ||(x + y == q.x + q.y)
                ||(x - y == q.x - q.y);
    }

    bool operator!=(Queen const &q) const {
        return !(*this == q);
    }
};

void placeQueens(int n) {
    vector<Queen> solu;
    Queen q(0, 0);
    do {
        if (n <= solu.size() || n <= q.y) {
            q = solu.back();
            ++q.y;
            solu.pop_back();
        }
        else {
            while (q.y < n && 0 < count(solu.begin(), solu.end(), q))
                ++q.y;
            if (n > q.y) {
                solu.push_back(q);
                if (n <= solu.size()) {
                    for (auto i : solu)
                        cout << i.y << ", ";
                    cout << endl;
                }
                ++q.x;
                q.y = 0;
            }
        }
    } while(0 < q.x || q.y < n);
}

int main(int argc, char** argv)
{
    int n;
    while (cout << "please input width of chess", cin >> n)
        placeQueens(n);

    return 0;
}
