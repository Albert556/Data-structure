#include <stdio.h>
#include <iostream>

using namespace std;

#define ElemType int

#define MaxSize_Static 10
typedef struct
{
    ElemType data[MaxSize_Static];
    int length;
} SqList; //定义静态顺序表

#define initSize_SeqList 10
typedef struct
{
    ElemType *data;
    int MaxSize_Dynamic;
    int length;
} SeqList;

void Print_SqList(SqList &L)
{
    for (int i = 0; i < L.length; i++)
    {
        cout << L.data[i] << endl;
    }
}

void Init_SqList(SqList &L) //初始化
{
    for (int i = 0; i < MaxSize_Static; i++)
    {
        L.data[i] = 0;
    }
    L.length = 0;
}

void Init_SeqList(SeqList &L)
{
    L.data = (ElemType *)malloc(initSize_SeqList * sizeof(ElemType));
    L.length = 0;
    L.MaxSize_Dynamic = initSize_SeqList;
}

void IncreaseSize(SeqList &L, int len)
{
    int *p = L.data;
    L.data = (ElemType *)malloc((L.MaxSize_Dynamic + len) * sizeof(ElemType));
    for (int i = 0; i < L.length; i++)
    {
        L.data[i] = p[i];
    }
    L.MaxSize_Dynamic = L.MaxSize_Dynamic + len;
    free(p);
}

bool Insert_SqList(SqList &L, int local, ElemType e)
{
    if (local > L.length + 1 || local < 1 || L.length + 1 > MaxSize_Static)
    {
        cout << "error: bad local" << endl;
        return false;
    }

    for (int i = L.length; i >= local; i--)
    {
        L.data[i] = L.data[i - 1];
    }
    L.data[local - 1] = e;
    L.length++;
    return true;
}

bool Delet_SqList(SqList &L, int local, ElemType &e)
{
    if (local < 1 || local > L.length)
    {
        cout << "error: bad local";
        return false;
    }
    e = L.data[local - 1];
    for (int i = local - 1; i < L.length - 1; i++)
    {
        L.data[i] = L.data[i + 1];
    }
    L.length--;
    return true;
}

ElemType GetElem_SqList(SqList &L, int local)
{
    if (local < 1 || local > L.length)
    {
        cout << "error: bad local" << endl;
        return 0;
    }
    return L.data[local - 1];
}

int Find_SqList(SqList &L, ElemType e)
{
    for (int i = 0; i < L.length; i++)
    {
        if (e == L.data[i])
        {
            return i + 1;
        }
    }
    return 0;
}

int main()
{
    // //SqList 操作
    // SqList L;
    // Init_SqList(L);
    // Insert_SqList(L, 1, 1);
    // Insert_SqList(L, 2, 2);
    // Insert_SqList(L, 3, 3);
    // ElemType e = -1;
    // Delet_SqList(L, 2, e);
    // cout << e << endl;
    // Print_SqList(L);

    return 0;
}
