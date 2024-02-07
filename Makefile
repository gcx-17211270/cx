server_files = Main.c main.h clone.c clone.h close.c close.h serialize.c serialize.h
client_files = sendMsg.c serialize.c serialize.h main.h
client: $(client_files)
	gcc -o scxd.exe $(client_files)  
server: 
	gcc -o cxd.exe $(server_files) -pthread
clean:
	rm -f *.exe

all: server client

