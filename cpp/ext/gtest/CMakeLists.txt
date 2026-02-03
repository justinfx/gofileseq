cmake_minimum_required(VERSION 3.15)

# Bundled Google Test
add_library(gtest STATIC
    src/gtest-all.cc
)

add_library(gtest_main STATIC
    src/gtest_main.cc
)

target_include_directories(gtest PUBLIC
    ${CMAKE_CURRENT_SOURCE_DIR}/include
    ${CMAKE_CURRENT_SOURCE_DIR}
)

target_include_directories(gtest_main PUBLIC
    ${CMAKE_CURRENT_SOURCE_DIR}/include
    ${CMAKE_CURRENT_SOURCE_DIR}
)

target_link_libraries(gtest_main gtest)

# Disable warnings for gtest
target_compile_options(gtest PRIVATE -w)
target_compile_options(gtest_main PRIVATE -w)
