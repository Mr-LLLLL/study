#include <pthread.h>
#include <stdio.h>
#include <unistd.h>

pthread_cond_t qready = PTHREAD_COND_INITIALIZER;
pthread_mutex_t qlock = PTHREAD_MUTEX_INITIALIZER;


void *dequeue(void *arg)
{
    while (1) {
        pthread_mutex_lock(&qlock);
        pthread_cond_wait(&qready, &qlock);
        printf("hello");
        pthread_mutex_unlock(&qlock);
    }
}

void *enqueue(void *arg)
{
    while (1) {
        pthread_mutex_lock(&qlock);
        getchar();
        pthread_cond_signal(&qready);
        pthread_mutex_unlock(&qlock);
    }
}

int main()
{
    pthread_t tid1, tid2;
    pthread_create(&tid1, NULL, enqueue, NULL);
    pthread_create(&tid2, NULL, dequeue, NULL);
    pthread_join(tid1, NULL);
    pthread_join(tid2, NULL);


    return 0;
}
