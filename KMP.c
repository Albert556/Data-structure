/*
 * @Author: Albert
 * @Date: 2020-12-06 08:55:02
 * @LastEditors: Albert
 * @LastEditTime: 2020-12-24 12:02:28
 * @Desctiption:
 */
#include <stdio.h>
#include <string.h>
#include <stdlib.h>

typedef int Position;
#define NotFound -1

void BuildNext(char *pattern, int *next)
{
    Position i, j;
    int m = strlen(pattern);
    next[0] = -1;

    for (j = 1; j < m; j++)
    {
        i = next[j - 1];
        while ((i >= 0) && (pattern[i + 1] != pattern[j]))
            i = next[i];
        if (pattern[i + 1] == pattern[j])
            next[j] = i + 1;
        else
            next[j] = -1;
    }
    for (int x = 0; x < m; x++)
    {
        printf("%d ", next[x]);
    }
    printf("\n");
}

Position KMP(char *string, char *pattern)
{
    int n = strlen(string);
    int m = strlen(pattern);
    Position s, p, *next;

    if (n < m)
        return NotFound;
    next = (Position *)malloc(sizeof(Position) * m);
    BuildNext(pattern, next);
    s = p = 0;
    while (s < n && p < m)
    {
        if (string[s] == pattern[p])
        {
            s++;
            p++;
        }
        else if (p > 0)
            p = next[p - 1] + 1;
        else
            s++;
    }
    return (p == m) ? (s - m) : NotFound;
}

int main()
{
    char string[] = "This is a simple example.";
    char pattern[] = "abacdababc";
    Position p = KMP(string, pattern);
    if (p == NotFound)
        printf("Not Found.\n");
    else
        printf("%d\n", p);
    return 0;
}
