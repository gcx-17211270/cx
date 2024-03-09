server_files = Main.c main.h clone.c clone.h close.c close.h serialize.c serialize.h
client_files = cxctl.c serialize.c serialize.h main.h
cxctl: $(client_files)
	gcc -o cxctl.exe $(client_files)  
cxd: 
	gcc -o cxd.exe $(server_files) -pthread
runcx:
	gcc -o runcx.exe runcx.c
clean:
	rm -f *.exe

all: cxctl cxd runcx

test:
	gcc -o ser_test serialize_test.c serialize.c serialize.h -std=c99
	./ser_test
    ifeq (0, $(shell echo $?) )
    	$(shell echo "test serialize false")
    endif
