#!/bin/bash

javac Main.java
javah Main

case "$OSTYPE" in
darwin*)
    # Mac (New Way)
    clang -shared -I "${JAVA_HOME}/include" -I "${JAVA_HOME}/include/darwin" -o libdemo.dylib Main.c
    ;;
linux*)
    # Linux
    clang -shared -I "${JAVA_HOME}/include" -I "${JAVA_HOME}/include/linux" -o libdemo.so Main.c
    ;;
*)
    echo "Unknow OS"
    exit 1
esac
