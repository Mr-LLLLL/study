#include <syslog.h>
#include <fcntl.h>
#include <sys/resource.h>
#include <stdio.h>
#include <signal.h>

void daemonize(const char *cmd) {
	int i, fd0, fd1, fd2;
	pid_t pid;
	struct rlimit r1;
	struct sigaction sa;
	
	// clear file creation mask.

	umask(0);
	// Get maximum number of file descriptors.

	if (getrlimit(RLIMIT_NOFILE, &r1) < 0)
		fprintf(stderr, "%s: can't get file limit", cmd);

	// Become a session leader to lose controlling TTY.
	
	if ((pid = fork()) < 0)
		fprintf(stderr, "%s: can't fork", cmd);
	else if (pid != 0)
		exit(0);
	setsid();

	// Ensure future opens won't allocate controlling TTYs.
	
	sa.sa_handler = SIG_IGN;
	sigemptyset(&sa.sa_mask);
	sa.sa_flags = 0;
	if (sigaction(SIGHUP, &sa, NULL) < 0)
		fprintf(stderr, "%s: can't ignore SIGHUP", cmd);
	if ((pid = fork())< 0)
		fprintf(stderr, "%s: can't fork", cmd);
	else if (pid != 0)
		exit(0);

	// Change the current working directory to the root so we won't prevent file systems from being unmounted.
	
	if (chdir("/") < 0)
		fprintf(stderr, "%s: can't change directory to /", cmd);

	// Close all open file descriptors.
	
	if (r1.rlim_max == RLIM_INFINITY)
		r1.rlim_max = 1024;
	for (i = 0; i < r1.rlim_max; i++)
		close(i);

	// Attach file descriptors 0, 1, and 2 to /dev/null.
	
	fd0 = open("/dev/null", O_RDWR);
	fd1 = dup(0);
	fd2 = dup(0);

	// Initialize the log file.
	
	openlog(cmd, LOG_CONS, LOG_DAEMON);
	if (fd0 != 0 || fd1 != 1 || fd2 != 2) {
		syslog(LOG_ERR, "unexpected file descriptors %d %d %d", fd0, fd1, fd2);
		exit(1);
	}
	sleep(100);
}

int main(int argc, char **argv)
{
	printf("create daemonize");
	daemonize("hello");
	return 0;
}


	
