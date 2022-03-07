#include <iostream>
#include <mutex>
#include <condition_variable>
#include <thread>
#include <set>
#include <unistd.h>
#include <deque>

class Request;

class Inventory
{
public:
    void add(Request* req)
    {
        std::lock_guard<std::mutex> lock(mutex_);
        requests_.insert(req);
    }

    void remove(Request* req) 
    {
        std::lock_guard<std::mutex> lock(mutex_);
        requests_.erase(req);
    }

    void printAll() const;

private:
    mutable std::mutex mutex_;
    std::set<Request*> requests_;
};

Inventory g_inventory;

class Request
{
public:
    void process()
    {
        std::lock_guard<std::mutex> lock(mutex_);
        g_inventory.add(this);
    }

    ~Request() __attribute__ ((noinline))       // in case to auto inline
    {
        std::lock_guard<std::mutex> lock(mutex_);
        g_inventory.remove(this);
    }

    void print() const __attribute__ ((noinline))
    {
        std::lock_guard<std::mutex> lock(mutex_);
    }

private:
    mutable std::mutex mutex_;
};

void Inventory::printAll() const
{
    std::lock_guard<std::mutex> lock(mutex_);
    for (std::set<Request*>::const_iterator it = requests_.begin(); it != requests_.end(); ++it) {
        (*it)->print();
    }
    printf("Inventory::printAll() unlocked\n");
}

void threadFunc()
{
    Request* req = new Request;
    req->process();
    delete req;
}

std::mutex mutex;
std::condition_variable cond;
std::deque<int> queue;

int dequeue()
{
    std::lock_guard<std::mutex> lock(mutex);
    while (queue.empty())
        cond.wait(mutex);
    assert(!queue.empty());
    int top = queue.front();
    queue.pop_front();
    return pop
}

void enqueue(int x)
{
    std::lock_guard<std::mutex> lock(mutex);
    queue.push_back(x);
    cond.notify();
}

int main(int argc, char **argv)
{
    /*
    std::thread thread(threadFunc);
    g_inventory.printAll();
    thread.join();
    */

    std::thread thread(dequeue);

    while (true)
    {
        int x;
        std::cin >> x;
        enqueue(x);
    }

    return 0;
}
