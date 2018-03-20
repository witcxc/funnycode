/*
 * =====================================================================================
 *
 *       Filename:  url_encode.cpp
 *
 *    Description:  
 *
 *        Version:  1.0
 *        Created:  2017/02/28 00时13分39秒
 *       Revision:  none
 *       Compiler:  gcc
 *
 *         Author:  Chen Xi (happyleaf), happyleaf.cx@alibaba-inc.com
 *   Organization:  
 *
 * =====================================================================================
 */
#include <iostream>
#include <string>
#include <assert.h>
using namespace std;

unsigned char ToHex(unsigned char x) 
{ 
	return  x > 9 ? x + 55 : x + 48; 
}

unsigned char FromHex(unsigned char x) 
{ 
	unsigned char y;
	if (x >= 'A' && x <= 'Z') y = x - 'A' + 10;
	else if (x >= 'a' && x <= 'z') y = x - 'a' + 10;
	else if (x >= '0' && x <= '9') y = x - '0';
	else assert(0);
	return y;
}

std::string UrlEncode(const std::string& str)
{
	std::string strTemp = "";
	size_t length = str.length();
	for (size_t i = 0; i < length; i++)
	{
		if (isalnum((unsigned char)str[i]) || 
				(str[i] == '-') ||
				(str[i] == '_') || 
				(str[i] == '.') || 
				(str[i] == '~'))
			strTemp += str[i];
		else if (str[i] == ' ')
			strTemp += "+";
		else
		{
			strTemp += '%';
			strTemp += ToHex((unsigned char)str[i] >> 4);
			strTemp += ToHex((unsigned char)str[i] % 16);
		}
	}
	return strTemp;
}

std::string UrlDecode(const std::string& str)
{
	std::string strTemp = "";
	size_t length = str.length();
	for (size_t i = 0; i < length; i++)
	{
		if (str[i] == '+') strTemp += ' ';
		else if (str[i] == '%')
		{
			assert(i + 2 < length);
			unsigned char high = FromHex((unsigned char)str[++i]);
			unsigned char low = FromHex((unsigned char)str[++i]);
			strTemp += high*16 + low;
		}
		else strTemp += str[i];
	}
	return strTemp;
}

#include	<stdio.h>
#include	<stdlib.h>

/* 
 * ===  FUNCTION  ======================================================================
 *         Name:  main
 *  Description:  
 * =====================================================================================
 */
int main(int argc, char *argv[]) {
    string src = "my=apples&are=green+and+red";
    string ret = UrlEncode(src);
    cout<<"src:["<<src<<"]"<<endl<<"encode:["<<ret<<"]"<<endl;
    string dec = UrlDecode(ret);
    cout<<dec<<endl;
    cout<<"dec:["<<ret<<"]"<<endl<<"src:["<<dec<<"]"<<endl;

    
    return EXIT_SUCCESS;
}				/* ----------  end of function main  ---------- */
