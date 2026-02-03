cmake_minimum_required(VERSION 3.15)
project(fileseq VERSION 3.0.0 LANGUAGES CXX)

# ANTLR4 4.13.1 requires C++17 (for std::optional, std::any, etc.)
# This is still widely supported: GCC 7+, Clang 5+, MSVC 2017+
set(CMAKE_CXX_STANDARD 17)
set(CMAKE_CXX_STANDARD_REQUIRED ON)
set(CMAKE_CXX_EXTENSIONS OFF)

# Compiler warnings
add_compile_options(-Wall -Wextra -Wpedantic)

# ANTLR4 parser generation (optional - only if Java is available)
find_program(JAVA_EXECUTABLE java)
if(JAVA_EXECUTABLE)
    set(ANTLR_JAR "${CMAKE_SOURCE_DIR}/../tools/antlr-4.13.1-complete.jar")
    set(GRAMMAR_FILE "${CMAKE_SOURCE_DIR}/../grammar/fileseq.g4")
    set(PARSER_OUTPUT_DIR "${CMAKE_SOURCE_DIR}/private/parser")

    add_custom_target(generate_parser
        COMMAND ${JAVA_EXECUTABLE} -jar ${ANTLR_JAR}
            -Dlanguage=Cpp
            -package fileseq
            -o ${PARSER_OUTPUT_DIR}
            -visitor
            -no-listener
            ${GRAMMAR_FILE}
        COMMENT "Generating ANTLR C++ parser from grammar"
        VERBATIM
    )
    message(STATUS "Java found - ANTLR parser can be regenerated with: make generate_parser")
else()
    message(STATUS "Java not found - using pre-generated ANTLR parser files")
endif()

# Build vendored ANTLR4 C++ runtime
set(ANTLR4_INSTALL OFF CACHE BOOL "Don't install ANTLR4 runtime" FORCE)
set(WITH_DEMO OFF CACHE BOOL "Don't build ANTLR4 demos" FORCE)
set(WITH_LIBCXX OFF CACHE BOOL "Don't use libc++" FORCE)
set(ANTLR_BUILD_CPP_TESTS OFF CACHE BOOL "Don't build ANTLR4 tests" FORCE)

# Force ANTLR4 to build with -fPIC so it can be linked into shared libraries
set(CMAKE_POSITION_INDEPENDENT_CODE ON)

add_subdirectory(ext/antlr4 EXCLUDE_FROM_ALL)

# ANTLR4 include directories (from vendored runtime)
set(ANTLR4_INCLUDE_DIRS ${CMAKE_SOURCE_DIR}/ext/antlr4/runtime/src)

# Include directories
include_directories(
    ${CMAKE_SOURCE_DIR}
    ${ANTLR4_INCLUDE_DIRS}
)

# Library sources
set(LIB_SOURCES
    sequence.cpp
    fileseq.cpp
    frameset.cpp
    pad.cpp
    error.cpp
    ranges/ranges.cpp
    private/strings.cpp
    private/frameset_p.cpp
    private/sequence_p.cpp
    private/parser/parse.cpp
    private/parser/parse_postprocess.cpp
)

# Generated parser sources
file(GLOB PARSER_SOURCES "private/parser/fileseq*.cpp")

# Create static library
add_library(fileseq_static STATIC ${LIB_SOURCES} ${PARSER_SOURCES})
target_compile_features(fileseq_static PUBLIC cxx_std_14)
set_target_properties(fileseq_static PROPERTIES OUTPUT_NAME fileseq)

# Link static library against vendored ANTLR4 runtime (static)
target_link_libraries(fileseq_static antlr4_static)

# Merge ANTLR4 static library into fileseq to make it self-contained
# This avoids requiring users to separately link libantlr4_static.a
if(APPLE)
    add_custom_command(TARGET fileseq_static POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E echo "Merging libantlr4_static.a into libfileseq.a..."
        COMMAND libtool -static -o libfileseq_combined.a
            $<TARGET_FILE:fileseq_static>
            $<TARGET_FILE:antlr4_static>
        COMMAND ${CMAKE_COMMAND} -E rename libfileseq_combined.a $<TARGET_FILE:fileseq_static>
        WORKING_DIRECTORY $<TARGET_FILE_DIR:fileseq_static>
        COMMENT "Combining fileseq and ANTLR4 static libraries"
    )
elseif(UNIX)
    # Linux: use ar/mri script to merge archives
    add_custom_command(TARGET fileseq_static POST_BUILD
        COMMAND ${CMAKE_COMMAND} -E echo "Merging libantlr4_static.a into libfileseq.a..."
        COMMAND echo "CREATE libfileseq_combined.a\nADDLIB $<TARGET_FILE:fileseq_static>\nADDLIB $<TARGET_FILE:antlr4_static>\nSAVE\nEND" | ar -M
        COMMAND ${CMAKE_COMMAND} -E rename libfileseq_combined.a $<TARGET_FILE:fileseq_static>
        WORKING_DIRECTORY $<TARGET_FILE_DIR:fileseq_static>
        COMMENT "Combining fileseq and ANTLR4 static libraries"
    )
endif()

# Create shared library
add_library(fileseq_shared SHARED ${LIB_SOURCES} ${PARSER_SOURCES})
target_compile_features(fileseq_shared PUBLIC cxx_std_14)
set_target_properties(fileseq_shared PROPERTIES OUTPUT_NAME fileseq)

# Link shared library against ANTLR4 static (embedded in the .so/.dylib)
target_link_libraries(fileseq_shared antlr4_static)

# Set version for shared library
set_target_properties(fileseq_shared PROPERTIES
    VERSION ${PROJECT_VERSION}
    SOVERSION ${PROJECT_VERSION_MAJOR}
)

# Testing
enable_testing()

# Try to find GTest
find_package(GTest QUIET)
if(NOT GTest_FOUND)
    # Use bundled gtest
    message(STATUS "Using bundled Google Test")
    add_subdirectory(ext/gtest EXCLUDE_FROM_ALL)
    set(GTEST_LIBRARIES gtest)
    set(GTEST_MAIN_LIBRARIES gtest_main)
    set(GTEST_INCLUDE_DIRS ${CMAKE_SOURCE_DIR}/ext/gtest/include)
else()
    message(STATUS "Using system Google Test")
    set(GTEST_LIBRARIES GTest::GTest)
    set(GTEST_MAIN_LIBRARIES GTest::Main)
endif()

# Test sources (following Test*.cc pattern)
file(GLOB TEST_SOURCES "test/Test*.cc")

# Main test runner
add_executable(fileseq_tests
    test/main.cc
    ${TEST_SOURCES}
)

target_include_directories(fileseq_tests PRIVATE ${GTEST_INCLUDE_DIRS})
target_link_libraries(fileseq_tests
    fileseq_static
    ${GTEST_LIBRARIES}
    pthread
)

add_test(NAME fileseq_tests COMMAND fileseq_tests)

# Install targets
install(TARGETS fileseq_static fileseq_shared
    ARCHIVE DESTINATION lib
    LIBRARY DESTINATION lib
)

install(FILES
    sequence.h
    fileseq.h
    frameset.h
    pad.h
    error.h
    DESTINATION include/fileseq
)

# Print configuration summary
message(STATUS "")
message(STATUS "FileSeq Configuration:")
message(STATUS "  C++ Standard:     C++${CMAKE_CXX_STANDARD}")
message(STATUS "  Compiler:         ${CMAKE_CXX_COMPILER_ID} ${CMAKE_CXX_COMPILER_VERSION}")
message(STATUS "  ANTLR4 runtime:   Vendored (static)")
message(STATUS "  Libraries:        Static + Shared")
message(STATUS "  Build type:       ${CMAKE_BUILD_TYPE}")
message(STATUS "")
