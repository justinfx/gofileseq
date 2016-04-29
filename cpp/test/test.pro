TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
CONFIG -= qt

SOURCES += main.cpp

unix: LIBS += -lpthread -L$$PWD/../ -lfileseq

INCLUDEPATH += $$PWD/..
DEPENDPATH += $$PWD/..

unix: QMAKE_LFLAGS += -Wl,-rpath,$$PWD/..
unix: PRE_TARGETDEPS += $$PWD/../libfileseq.so
