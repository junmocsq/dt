#include "1-3.h"
#include "1-6.h"
#include "1-8.h"
#include <stdio.h>
#include <string.h>
#define N 30
int main()
{
    int list[N];
    // 选择排序
    sort = selection_sort;
    generate_arr_int(list, N);
    arr_printf_int(list, N);
    sort(list, N);
    printf("排序后\n");
    arr_printf_int(list, N);

    // 折半查找
    int index = binsearch(list, 20, 0, N - 1);
    int index2 = binsearch_recv(list, 20, 0, N - 1);
    if (index != index2)
    {
        printf("alg error!\n");
    }
    if (index == -1)
    {
        printf("find val not exists!\n");
    }
    else
    {
        printf("find val:%d\n", list[index]);
    }

    // 全排列
    char str[] = "abcd";
    perm(str, strlen(str));

    return 0;
}
