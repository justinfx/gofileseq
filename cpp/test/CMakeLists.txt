project(gtest)

find_package(Threads REQUIRED)

file(GLOB test_SOURCES
        main.cc
        Test*.cc)

set(gtest_SOURCES
        gtest/src/gtest-all.cc)

set(gtest_INCLUDES
        .
        gtest/include
        gtest
        )

include_directories(${gtest_INCLUDES})

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Wno-undef -g -ggdb")

add_definitions(-DFILESEQ_TEST_DIR=${CMAKE_CURRENT_SOURCE_DIR})

add_executable(fileseq_test ${test_SOURCES} ${gtest_SOURCES})

target_link_libraries(fileseq_test
        fileseq_static
        pthread
        )
