/*
 * =====================================================================================
 *
 *       Filename:  code.c
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  2017/06/11 22时56分45秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Chen Xi (happyleaf), happyleaf.cx@alibaba-inc.com
 *   Organization:  
 *
 * =====================================================================================
 */

int accum = 0;
int sum(int x, int y)
{
    int t = x + y;
    accum += t;
    return t;
}
