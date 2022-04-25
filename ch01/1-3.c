#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include "1-3.h"


void arr_printf_int(int list[], int length)
{
    printf("array:");
    for (int i = 0; i < length; i++)
    {
        printf("% 8d", list[i]);
    }
    printf("\n");
}

void generate_arr_int(int list[], int length)
{
    srand((unsigned int)time(NULL));
    for (int i = 0; i < length; i++)
    {
        list[i] = rand() % 100;
    }
}

void selection_sort(int list[], int length)
{
    int i, j, min, temp;
    for (i = 0; i < length - 1; i++)
    {
        min = i;
        for (j = i + 1; j < length; j++)
        {
            if (list[j] < list[min])
            {
                min = j;
            }
        }
        if (i != min)
        {
            SWAP(list[i], list[min], temp);
        }
    }
}
