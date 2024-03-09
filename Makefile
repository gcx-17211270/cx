server_files = Main.c main.h clone.c clone.h close.c close.h serialize.c serialize.h log.c log.h
client_files = cxctl.c serialize.c serialize.h main.h log.c log.h
ctl:
	gcc -o cxctl.exe $(client_files) -std=c99 -D __UNIX_IPC__
cxd: 
	gcc -o cxd.exe $(server_files) -pthread -std=c99 -D __UNIX_IPC__
runcx:
	gcc -o runcx.exe runcx.c -std=c99 -D __UNIX_IPC__
clean:
	rm -f *.exe

all: ctl cxd runcx

test:
	gcc -o ser_test serialize_test.c serialize.c serialize.h -std=c99
	./ser_test
    ifeq (0, $(shell echo $?) )
    	$(shell echo "test serialize false")
    endif
