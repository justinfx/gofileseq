# https://github.com/amireh/libalgol/blob/master/CMake/Modules/FindPCRECPP.cmake
# FindPCRECPP.cmake
# --
# Find the PCRECPP library
#
# This module defines:
#   PCRECPP_INCLUDE_DIRS - where to find pcrecpp.h
#   PCRECPP_LIBRARIES    - the pcrecpp library
#   PCRECPP_FOUND        - True if PCRECPP was found

Include(FindModule)
FIND_MODULE(PCRECPP pcrecpp.h "" "" pcrecpp "" "")