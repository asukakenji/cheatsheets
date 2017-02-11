#include "Main.h"

#include <stdio.h>

JNIEXPORT void JNICALL Java_Main_printOkay(JNIEnv* env, jclass clazz)
{
    printf("Okay\n");
}
