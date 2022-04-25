#include <stdio.h>
#include "1-8.h"
#include "1-3.h"
void _perm(char list[],int i,int n){
    if(i==n){
        printf("%s\t",list);
        return;
    }
    char temp;
    for (int j=i;j<n;j++){
        SWAP(list[i],list[j],temp);
        _perm(list,i+1,n);
        SWAP(list[i],list[j],temp);
    }
}

void perm(char list[],int n){
   _perm(list,0,n);
   printf("\n");
}