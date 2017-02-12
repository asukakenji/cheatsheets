#!/bin/bash

case "$OSTYPE" in
darwin*)
    # Mac (New Way)
    clang -shared -o libdemo.dylib libdemo.c

    # Mac (Traditional Way)
    # Reference:
    # https://developer.apple.com/library/content/documentation/DeveloperTools/Conceptual/DynamicLibraries/100-Articles/CreatingDynamicLibraries.html
    #clang -dynamiclib -current_version 1.0 -compatibility_version 1.0 -o libdemo.A.dylib libdemo.c

    # Using this command gives the same result as using the "-shared" command (New Way)
    # The output only differs in meta data (timestamp, etc.)
    #clang -dynamiclib -o libdemo.dylib libdemo.c
    ;;
linux*)
    # Linux
    clang -shared -o libdemo.so libdemo.c
    ;;
*)
    echo "Unknow OS"
    exit 1
esac
