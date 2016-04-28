#!/bin/bash

# Build the Go wrapper lib
go build -buildmode=c-archive -o ./private/libgofileseq.a ./private


# Generate a Makefile for building the C++ lib
unamestr=`uname`

if [[ "$unamestr" == 'Linux' ]]; then
	qmake ./fileseq.pro

elif [[ "$unamestr" == 'Darwin' ]]; then
	qmake -spec macx-g++ ./fileseq.pro

else
	echo "Configure only support Linux/OSX"
	exit 1

fi

make



