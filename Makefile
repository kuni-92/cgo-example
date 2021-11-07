build:
	gcc -c libc/*.c
	ar rs libc.a *.o
	rm helloworld.o
	go build