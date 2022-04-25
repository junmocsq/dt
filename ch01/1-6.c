#include "1-6.h"

int binsearch(int list[], int findVal, int left, int right)
{
    int result = -1;
    int mid;
    while (left <= right)
    {
        mid = (left + right) / 2;
        if (list[mid] == findVal)
            return mid;
        else if (list[mid] < findVal)
            left = mid + 1;
        else
            right = mid - 1;
    }

    return result;
}
int binsearch_recv(int list[], int findVal, int left, int right)
{
    int result = -1;
    if (left > right)
    {
        return result;
    }
    int mid = (left + right) / 2;
    if (list[mid] == findVal)
    {
        return mid;
    }
    else if (list[mid] < findVal)
    {
        return binsearch(list, findVal, mid + 1, right);
    }
    else
    {
        return binsearch(list, findVal, left, mid - 1);
    }
    return result;
}
