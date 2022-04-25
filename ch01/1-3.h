#define SWAP(x, y, t) ((t) = (x), (x) = (y), (y) = (t))
void (*sort)(int list[], int length);
// 选择排序
void selection_sort(int list[], int length);

void arr_printf_int(int list[], int length);
void generate_arr_int(int list[], int length);