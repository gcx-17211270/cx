#ifndef __SERIALIZE_H__
#define __SERIALIZE_H__

// 用来将参数序列化传递给服务器端
char* Serialize1(int argc, char** argv, int* size_r);

// 服务器端用来反序列化数据，argc用来接受参数个数
char** Deserialize1(int* argc, char* s);

// 服务器端用来反序列化数据，argc用来接受参数个数
char** Deserialize2(char* s, int len, int* argc);

struct Pack {
    int len;
    char* arg;
};

struct Info {
    int argc;
    struct Pack* argv;
    int size;
};

struct Info* ArgSerialize(int argc, char** argv);

void printInfo(struct Info*);

#endif
