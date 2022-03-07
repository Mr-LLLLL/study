#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <pthread.h>
#include <unistd.h>
void* thread_summation(void *arg);
int sum = 0;

int main(int argc, char *argv[])
{
	pthread_t t_id1, t_id2;
	int range1[] = {1, 5};
	int range2[] = {2, 10000000};
	if (pthread_create(&t_id1, NULL, thread_summation, (void*)range1) != 0) {
		puts("pthread_creat() error");
		return -1;
	}
	if (pthread_create(&t_id2, NULL, thread_summation, (void*)range2) != 0) {
		puts("pthread_creat() error");
		return -1;
	}

	//if (pthread_join(t_id1, NULL) != 0) {
	//	puts("pthrea_join() error\n");
	//	return -1;
	//}		
	pthread_detach(t_id2);
	pthread_detach(t_id1);
	sleep(100);
	puts("finish first thread");

	printf("result: %d \n", sum);
	return 0;
}

void* thread_summation(void *arg) {
	int start = ((int*)arg)[0];
	int end = ((int*)arg)[1];
	sleep(3);
	printf("thread %d is running\n", start);
	while (start <= end) {
		sum += start;
		++start;
	}
	return NULL;
}
