#!/bin/bash

case "$OSTYPE" in
darwin*)
    # Mac (New Way)
    clang -shared -o libdemo.dylib c/libdemo.c
    ;;
linux*)
    # Linux
    clang -shared -o libdemo.so c/libdemo.c
    ;;
*)
    echo "Unknow OS"
    exit 1
esac
