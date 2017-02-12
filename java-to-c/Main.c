#include "Main.h"

#include <stdio.h>

/* The signature should be the same as that generated in Main.h */
JNIEXPORT void JNICALL Java_Main_printOkay(JNIEnv* env, jclass clazz)
{
    printf("Okay\n");
}
