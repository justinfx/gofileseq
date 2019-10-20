cmake_minimum_required (VERSION 3.1)
project (fileseq)

set ( CMAKE_MODULE_PATH ${CMAKE_SOURCE_DIR}/cmake )

set(CMAKE_VERBOSE_MAKEFILE ON)

if("$ENV{HAVE_REGEX}" STREQUAL "0")
    message("Disabling <regex> compiler support (HAVE_REGEX=0)")
    find_package(PCRECPP)
    add_definitions(-DHAVE_REGEX=0)
    include_directories(${PCRECPP_INCLUDE_DIRS})

else()
    message("Enabling <regex> compiler support (HAVE_REGEX=1)")
    add_definitions(-DHAVE_REGEX=1)

endif()

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

if("${PCRECPP_FOUND}")
    target_link_libraries(fileseq_shared ${PCRECPP_LIBRARIES})
    target_link_libraries(fileseq_static ${PCRECPP_LIBRARIES})
endif()