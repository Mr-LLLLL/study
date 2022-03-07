#include <errno.h>
#include <stdio.h>
#include <fcntl.h>

char buf[5000000];

int main(int argc, char **argv)
{
	int ntowrite, nwrite;
	char *ptr;

	ntowrite = read(fileno(stdin), buf, sizeof(buf));
	fprintf(stderr, "read %d bytes\n", ntowrite);

	int flag;
	flag = fcntl(fileno(stdout), F_GETFL, 0);
	fcntl(fileno(stdout), F_SETFL, flag | O_NONBLOCK);

	ptr = buf;
	while (ntowrite > 0) {
		errno = 0;
		nwrite = write(fileno(stdout), ptr, ntowrite);
		fprintf(stderr, "nwrite = %d, errno = %d\n", nwrite, errno);

		if (nwrite > 0) {
			ptr += nwrite;
			ntowrite -= nwrite;
		}
	}

	fcntl(fileno(stdout), F_SETFL, flag);



	return 0;
}
