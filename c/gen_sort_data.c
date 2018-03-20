/*
 * =====================================================================================
 *
 *       Filename:  gen_sort_data.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  2017/03/10 19时24分45秒
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
#include <string.h>
#include <errno.h>
#include <time.h>

/* 
 * ===  FUNCTION  ======================================================================
 *         Name:  main
 *  Description:  
 * =====================================================================================
 */

int main(int argc, char *argv[]) {

    FILE	*fp;										/* output-file pointer */
    char	*fp_file_name = "data.txt";		/* output-file name    */

    srand((int)time(0));
    fp	= fopen( fp_file_name, "w" );
    if ( fp == NULL ) {
        fprintf ( stderr, "couldn't open file '%s'; %s\n",
                fp_file_name, strerror(errno) );
        exit (EXIT_FAILURE);
    }
    int size  = rand()%1000000;

    fprintf(fp,"%d", size);
    for (int i = 0; i < size; ++i) {
        fprintf(fp," %d", rand());
    }
    if( fclose(fp) == EOF ) {			/* close output file   */
        fprintf ( stderr, "couldn't close file '%s'; %s\n",
                fp_file_name, strerror(errno) );
        exit (EXIT_FAILURE);
    }

    return EXIT_SUCCESS;
}				/* ----------  end of function main  ---------- */
