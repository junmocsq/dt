#include <stdio.h>
#include <stdlib.h>
#include "1-3.h"
int func_2(int list[], int x, int length);
int func_3(int n);
void func_5(int list[], int n, int f(int));
int func_5_f(int n) { return n * n; }
int func_6(int n);
// 阶乘
int func_7(int n);
int func_7_recv(int n);
// 斐波那契
int func_8(int n);
int func_8_recv(int n);

// Hanoi 汉诺塔 移动
int func_11_recv(int n);
void func_12_recv(char *list, int n);
// 测试malloc用
void func_12_recv_2();
int main()
{
    int arr[] = {1, 2, 3, 4, 5, 6};
    printf("x:%d,res:%d\n", 1, func_2(arr, 1, 5));
    printf("x:%d,res:%d\n", 2, func_2(arr, 2, 5));
    printf("x:%d,res:%d\n", 3, func_2(arr, 3, 5));

    printf("prac3 n:%d res:%d\n", 4, func_3(4));

    int arr5[] = {1, 2, 3, 4, 5, 111, 22, 33, -1, -2, -3, -4, -5};

    func_5(arr5, 13, func_5_f);

    for (int i = 3; i < 1000; i++)
    {
        if (func_6(i) == 1)
        {
            printf("正整数%d的因子之和等于它自己\n", i);
        }
    }

    for (int i = 1; i < 20; i++)
    {
        if (func_7(i) != func_7_recv(i))
        {
            printf("n:%d 的阶乘错误\n", i);
        }
        if (func_8(i) != func_8_recv(i))
        {
            printf("n:%d 的斐波那契错误\n", i);
        }
    }

    func_11_recv(5);

    func_12_recv((char[]){'a', 'b', 'c', 'd'}, 4);
    func_12_recv_2();
}
int func_2(int list[], int x, int n)
{
    int result = 0;
    for (int i = n; i >= 0; i--)
    {
        result = result * x + list[i];
    }
    return result;
}

void _func_3(char *temp[], int n, int i, int *res)
{
    if (i == n)
    {
        printf("<");
        for (int j = 0; j < n; j++)
            printf("%s ", temp[j]);
        printf(">\t");
        *res += 1;
        return;
    }
    temp[i] = "TRUE";
    _func_3(temp, n, i + 1, res);
    temp[i] = "FALSE";
    _func_3(temp, n, i + 1, res);
}
int func_3(int n)
{
    char *temp[n + 1];
    temp[n] = 0;
    int result = 0;
    _func_3(temp, n, 0, &result);
    printf("\n");
    return result;
}

void func_5(int list[], int n, int f(int))
{
    int res[n];
    for (int i = 0; i < n; i++)
    {
        res[i] = f(list[i]);
    }

    int temp;
    for (int i = 0; i < n - 1; i++)
    {
        int min = i;
        for (int j = i + 1; j < n; j++)
        {
            if (res[min] > res[j])
                min = j;
        }
        SWAP(res[i], res[min], temp);
        SWAP(list[i], list[min], temp);
    }

    for (int i = 0; i < n - 1; i++)
    {
        if (res[i] == res[i + 1])
        {
            printf("%d 和 %d 的f值相等：%d\n", list[i], list[i + 1], res[i]);
        }
    }
}

int func_6(int n)
{
    int sum = 0;
    if (n < 3)
    {
        return 0;
    }
    int i = 2;
    // printf("%d的因子:",n);
    int temp = n;
    while (i < n && i <= temp)
    {
        if (temp % i == 0)
        {
            sum += i;
            // printf("%d ",i);
            temp /= i;
        }
        else
        {
            i++;
        }
    }
    // printf(" 和为:%d\n",sum);
    if (sum == n)
    {
        return 1;
    }
    return 0;
}

int func_7(int n)
{
    int result = 1;
    for (int i = n; i > 1; i--)
    {
        result *= i;
    }
    return result;
}

int func_7_recv(int n)
{
    if (n < 2)
        return 1;
    return n * func_7_recv(n - 1);
}

int func_8(int n)
{
    if (n == 0)
        return 0;
    if (n == 1)
        return 1;
    int f0 = 0, temp;
    int result = 1;
    for (int i = 2; i <= n; i++)
    {
        temp = result;
        result += f0;
        f0 = temp;
    }
    return result;
}
int func_8_recv(int n)
{
    if (n == 0)
        return 0;
    if (n == 1)
        return 1;
    return func_8_recv(n - 1) + func_8_recv(n - 2);
}

char func_11_recv_third(char s1, char s2)
{
    if (s1 == 'A')
    {
        if (s2 == 'B')
            return 'C';
        if (s2 == 'C')
            return 'B';
    }
    else if (s1 == 'B')
    {
        if (s2 == 'A')
            return 'C';
        if (s2 == 'C')
            return 'A';
    }
    else
    {
        if (s2 == 'A')
            return 'B';
        if (s2 == 'B')
            return 'A';
    }
    return 'A';
}

// A B C
int _func_11_recv(int n, char src, char dst)
{
    int num = 0;
    if (n == 0)
    {
        return num;
    }
    char third = func_11_recv_third(src, dst);

    num += _func_11_recv(n - 1, src, third);
    num++;
    printf("盘子 %d 从 %c 移动 %c \n", n, src, dst);
    num += _func_11_recv(n - 1, third, dst);
    return num;
}
int func_11_recv(int n)
{
    return _func_11_recv(n, 'A', 'C');
}

void _func_12_recv(char *list, char temp[], int n, int i, char ***res, int *resNum)
{
    if (i == n)
    {
        char *tmp = (char *)malloc(sizeof(char) * (n + 3));
        tmp[0] = '{';
        tmp[n + 1] = '}';
        tmp[n + 2] = 0;
        for (int j = 0; j < n; j++)
        {
            tmp[j + 1] = temp[j];
        }

        (*resNum)++;
        *res = (char **)realloc(*res, sizeof(char *) * *resNum);
        (*res)[*resNum - 1] = tmp;

        return;
    }
    temp[i] = ' ';
    _func_12_recv(list, temp, n, i + 1, res, resNum);
    temp[i] = list[i];
    _func_12_recv(list, temp, n, i + 1, res, resNum);
}

void func_12_recv(char *list, int n)
{
    char temp[n];
    char **res = NULL;
    int resNum = 0;
    _func_12_recv(list, temp, n, 0, &res, &resNum);
    printf("total:%d elements: ", resNum);
    for (int i = 0; i < resNum; i++)
    {
        printf("%s\t", res[i]);
        free(res[i]);
    }
    printf("\n");
    free(res);
}
void _func_12_recv_2()
{
    int n = 4;
    char **res = NULL;
    int i = 0;
    for (i = 1; i <= n; i++)
    {
        // 不可取，tmp为局部变量？？？
        char tmp[7];
        tmp[0] = '{';
        tmp[n + 1] = '}';
        tmp[n + 2] = 0;
        for (int j = 1; j <= n; j++)
        {
            tmp[j] = 'A' + i;
        }
        res = (char **)realloc(res, sizeof(char *) * i);
        printf("csapp--tmp:%s %p\n", tmp, tmp);
        res[i - 1] = tmp;
    }
    printf("total:%d elements: ", n);
    // output:
    // total:4 elements: {EEEE}        {EEEE}  {EEEE}  {EEEE}
    for (int i = 0; i < n; i++)
    {
        printf("%s\t", res[i]);
    }
    printf("\n");
    free(res);
}
void func_12_recv_2()
{
    _func_12_recv_2();
}