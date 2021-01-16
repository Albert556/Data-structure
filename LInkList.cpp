#include <iostream>

using namespace std;

#define ElemType int

typedef struct LNode
{
    ElemType data;
    LNode *next;
} LNode, *LinkList;

typedef struct DNode
{
    ElemType data;
    struct DNode *prior, *next;
} DNode, *DLinkList;

bool InitList_NonHead(LinkList &L)
{
    L = NULL;
    return true;
}

bool Empty_NonHead(LinkList &L)
{
    return (L == NULL);
}

bool InitList_Head(LinkList &L)
{
    L = (LNode *)malloc(sizeof(LNode));
    if (L == NULL)
    {
        return false;
    }
    L->next = NULL;
    return true;
}

bool Empty_Head(LinkList &L)
{
    return L->next == NULL;
}

bool InsertList_Head(LinkList &L, int local, ElemType e)
{
    if (local < 1)
    {
        return false;
    }
    LNode *p;
    int j = 0;
    p = L;
    while (p != NULL && j < local - 1)
    {
        p = p->next;
        j++;
    }
    if (p == NULL)
    {
        return false;
    }
    LNode *s = (LNode *)malloc(sizeof(LNode));
    s->data = e;
    s->next = p->next;
    p->next = s;
    return true;
}

bool InsertList_NonHead(LinkList &L, int local, ElemType e)
{
    if (local < 1)
    {
        return false;
    }
    if (local == 1)
    {
        LNode *s = (LNode *)malloc(sizeof(LNode));
        s->data = e;
        s->next = L;
        L = s;
        return true;
    }
    LNode *p = L;
    int j = 1;
    while (p != NULL && j < local - 1)
    {
        p = p->next;
        j++;
    }
    if (p == NULL)
    {
        return false;
    }
    LNode *s = (LNode *)malloc(sizeof(LNode));
    s->data = e;
    s->next = p->next;
    p->next = s;
    return true;
}

void OutLinkList(LinkList &L)
{
    LNode *p = L->next;
    while (p != NULL)
    {
        cout << p->data << endl;
        p = p->next;
    }
}

bool InitDLinkList(DLinkList &L)
{
    L = (DNode *)malloc(sizeof(DNode));
    if (L == NULL)
    {
        return false;
    }
    L->prior = NULL;
    L->next = NULL;
    return true;
}

bool Empty(DLinkList &L)
{
    if (L->next == NULL)
    {
        return true;
    }
    else
    {
        return false;
    }
}

bool InsertNextDNode(DNode *p, DNode *s)
{
    if (p == NULL || s == NULL)
    {
        return false;
    }
    s->next = p->next;
    if (p->next != NULL)
    {
        p->next->prior = s;
    }
    s->prior = p;
    p->next = s;
    return true;
}

bool DeletNextDNode(DNode *p)
{
    if (p == NULL)
    {
        return false;
    }
    DNode *q = p->next;
    if (p->next != NULL)
    {
        q->next->prior = p;
    }
    free(q);
    return true;
}

int main()
{
    cout << sizeof(int) << endl;

    return 0;
}
