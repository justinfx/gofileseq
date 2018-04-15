TEMPLATE = app
TARGET = test
CONFIG += console
CONFIG -= app_bundle
CONFIG -= qt

SOURCES += main.cpp

INCLUDEPATH += $$PWD/..
DEPENDPATH += $$PWD/..

unix: {
    LIBS += -lpthread
    LIBS += -L$$PWD/../ -lfileseq
    # LIBS += -Wl,-Bstatic -lfileseq -lgofileseq -Wl,-Bdynamic
    QMAKE_RPATHDIR += $$_PRO_FILE_PWD_/..
}

QMAKE_MACOSX_DEPLOYMENT_TARGET = 10.6
