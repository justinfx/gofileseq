TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
CONFIG -= qt

SOURCES += main.cpp


unix: LIBS += -L$$PWD/../cpp/ -lfileseq

INCLUDEPATH += $$PWD/../cpp
DEPENDPATH += $$PWD/../cpp

unix: PRE_TARGETDEPS += $$PWD/../cpp/libfileseq.a
