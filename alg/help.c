#include "help.h"

void print_int_array(const int *arr,int len)  {
    int i;
    for(i=0;i<len;i++)
        printf("%d ",arr[i]);
    printf("\n");
}

void print_str_array(const char *arr[],int len)  {
    int i;
    for(i=0;i<len;i++)
        printf("%s ",arr[i]);
    printf("\n");
}

void print_float_array(const float *arr,int len) {
    int i;
    for(i=0;i<len;i++)
        printf("%f ",arr[i]);
    printf("\n");
}

void generate_random_int_array(int * arr,int len){
    int i;
    srand(time(NULL));
    for(i=0;i<len;i++)
        arr[i]=rand()%1000;
        //printf("%d ",arr[i]);
        //printf("\n");
}

void generate_random_string(char *arr,int len){
    int s1 = 'z'-'a'+1;
    int s2 = 'Z'-'A'+1;
    printf("%d %d \n",'a','A');
    int i;
    srand(time(NULL));
    for(i=0;i<len-1;i++)
    {
        int index = rand()%(s1+s2);
        arr[i]=index<s1?'a'+index:'A'+index-s1;
    }
    arr[i]='\0';
    printf("%s\n",arr);
        
}

// cc -DMM help.c && ./a.out

#ifdef MM
int main(){
    int arr[10];
    generate_random_int_array(arr,10);
    print_int_array(arr,10);

    char strArr[20];
    generate_random_string(strArr,20) ; 
    return 0;
}
#endif