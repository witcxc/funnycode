/*
 * =====================================================================================
 *
 *       Filename:  quick_sort.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  2017/03/10 18时43分47秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Chen Xi (happyleaf), happyleaf.cx@alibaba-inc.com
 *   Organization:  
 *
 * =====================================================================================
 */


#include	<stdio.h>
#include	<stdlib.h>


/* 
 * ===  FUNCTION  ======================================================================
 *         Name:  quick_sort
 *  Description:  
 * =====================================================================================
 */
void quick_sort(int *arr, int size) {

    if (size <= 1) {
        return;
    }
    int mark = arr[0];
    int s_i = 0;
    
    for (int i = 1; i <= size - 1; ++i ) {

        if (arr[i] <= mark) {
            int tmp = arr[s_i + 1];
            arr[s_i + 1] = arr[i];
            arr[i] = tmp;
            s_i = s_i + 1;
        }
    }
    arr[0] = arr[s_i];
    arr[s_i] = mark;
    if(s_i > 1) {
        quick_sort(arr, s_i);
    }

    if ((size - s_i - 1) > 1) {
        quick_sort(arr + s_i + 1, size - s_i - 1);
    }
}		/* -----  end of function quick_sort  ----- */
/* 
 * ===  FUNCTION  ======================================================================
 *         Name:  main
 *  Description:  
 * =====================================================================================
 */
int main(int argc, char *argv[]) {
//    int arr[5] = {3, 4 , 2, 5, 1};
//    printf("%d %d %d %d %d\n", arr[0], arr[1], arr[2], arr[3], arr[4]);
//    quick_sort(arr, sizeof(arr)/sizeof(arr[0]));
//    printf("%d %d %d %d %d\n", arr[0], arr[1], arr[2], arr[3], arr[4]);
    int size = 0;
    int * arr = NULL;
    int i = 0;

    while(scanf("%d", &size) == 1){
        arr = (int *)malloc(sizeof(int) * size);

        for (i = 0; i < size; ++i)
        {
            scanf("%d", &(arr[i]));

        }
        for (i = 0; i < size; ++i) {
            printf("%d ", arr[i]);
        }
        printf("\n");
        quick_sort(arr, size);
        for (i = 0; i < size; ++i) {
            printf("%d ", arr[i]);
        }
        printf("\n");
    }

    return EXIT_SUCCESS;
}				/* ----------  end of function main  ---------- */
