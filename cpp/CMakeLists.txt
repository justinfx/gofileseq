cmake_minimum_required (VERSION 3.1)
project (fileseq)

#set(CMAKE_VERBOSE_MAKEFILE ON)

add_definitions(-DHAVE_REGEX=1)

set(CMAKE_CXX_STANDARD 11)
set(CMAKE_CXX_FLAGS "-fPIC -std=c++11")

set(fileseq_SOURCES
        error.cpp
        fileseq.cpp
        frameset.cpp
        pad.cpp
        private/frameset_p.cpp
        private/sequence_p.cpp
        private/strings.cpp
        ranges/ranges.cpp
        sequence.cpp
        )

set(fileseq_INCLUDES
        .
        ranges
        private
        )

add_subdirectory(test)

include_directories(${fileseq_INCLUDES})

add_library(fileseq_obj OBJECT ${fileseq_SOURCES})
set_property(TARGET fileseq_obj PROPERTY
        POSITION_INDEPENDENT_CODE 1
        )

add_library(fileseq_shared SHARED $<TARGET_OBJECTS:fileseq_obj>)
SET_TARGET_PROPERTIES(fileseq_shared PROPERTIES
        OUTPUT_NAME fileseq
        CLEAN_DIRECT_OUTPUT 1
        )

add_library(fileseq_static STATIC $<TARGET_OBJECTS:fileseq_obj>)
SET_TARGET_PROPERTIES(fileseq_static PROPERTIES
        OUTPUT_NAME fileseq
        CLEAN_DIRECT_OUTPUT 1
        )

