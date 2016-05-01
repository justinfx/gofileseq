TEMPLATE = app
CONFIG += console
CONFIG -= app_bundle
CONFIG -= qt

SOURCES += main.cpp

INCLUDEPATH += $$PWD/..
DEPENDPATH += $$PWD/..

unix: {
    LIBS += -lpthread -L$$PWD/../ -lfileseq
    QMAKE_RPATHDIR += $$_PRO_FILE_PWD_/..
}

QMAKE_MACOSX_DEPLOYMENT_TARGET = 10.6
