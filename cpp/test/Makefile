CC=g++

target=app

ALL:$(target)

%.o:%.c
	$(CC) -c $< -o $@ 

$(target):*.o
	$(CC) $^ -o $@

clean:
	-rm -rf *.o

.PHONY:clean ALL
