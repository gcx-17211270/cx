char* Serialize(int argc, char** argv);
char* Serialize2(int argc, char** argv);

// 服务器端用来反序列化数据，argc用来接受参数个数
char** Deserialize2(int* argc, char* s);

struct Pack {
    int len;
    char* arg;
};

struct Info {
    int argc;
    struct Pack* argv;
};

