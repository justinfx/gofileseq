#!/bin/bash

# Build the Go wrapper lib
export CGO_CXXFLAGS="-fPIC"

echo "Configuring build..."

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

echo "Building libgofileseq (static)..."
go build -buildmode=c-archive -o ./libgofileseq.a ./private

echo "Building libgofileseq (shared)..."
go build -buildmode=c-shared -o ./libgofileseq.so ./private

echo "Building libfileseq..."
make

echo "Done."



