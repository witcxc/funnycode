#include    <stdio.h>
#include    <stdlib.h>
#include <time.h>
#include <sys/time.h>
#include <xlocale.h>
#include <string.h>
/*
 *
 * ===  FUNCTION  ======================================================================
 *         Name:  main
 *  Description:
 * =====================================================================================
 */
int main(int argc, char *argv[]) {
    struct tm tm_p;

    char buf [101];
    memset(buf,0,sizeof(buf));

    if ( fgets(buf, 40, stdin) != NULL ) {
        if (strcmp(argv[1], "git") == 0)
        {
            strptime(buf, "%a %b %e %T %Y", &tm_p);
            // puts(buf);
            // printf("%d %d %d %d %d %d\n",tm_p.tm_year,tm_p.tm_mon,tm_p.tm_mday,tm_p.tm_hour,tm_p.tm_min,tm_p.tm_sec);
            strftime (buf,40,"%Y-%m-%d %T",&tm_p);

            puts(buf);
        } else if(strcmp(argv[1], "fmt") == 0)
        {
            strptime(buf, "%Y-%m-%d %T", &tm_p);
            strftime (buf,40,"%Y%m%d%H%M",&tm_p);
            puts(buf);

        }

    }

    return EXIT_SUCCESS;
}
