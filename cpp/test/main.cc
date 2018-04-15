#include <gtest/gtest.h>

#include <cstdlib>
#include <cstring>
#include <unistd.h>

#define xstr(a) str(a)
#define str(a) #a

int main( int argc, char** argv ) {

    #ifdef FILESEQ_TEST_DIR
    int ret = chdir(xstr(FILESEQ_TEST_DIR));
    if (ret != 0) {
        perror("Failed to change dir to: " xstr(FILESEQ_TEST_DIR));
        exit(1);
    }
    #endif

    testing::InitGoogleTest( &argc, argv );

    return RUN_ALL_TESTS();
}

