#include <iostream>
#include "Queue.h"

#define Elemtype int
#define MaxSize 50

typedef struct
{
    Elemtype data[MaxSize];
    int front, rear;
} SqQueue;

typedef struct
{
    Elemtype data;
    struct LinkNode *next;
} LinkNode;
typedef struct{
    LinkNode *front, *rear;
} LinkQueue;

void InitQueue(SqQueue &Q){
    Q.front = Q.rear = 0;
}

bool QueueEmpty(SqQueue &Q){
    if (Q.front == Q.rear){
        return true;
    }
    else
    {
        return false;
    }
}

bool EnQueue(SqQueue &Q, Elemtype x){
    if ((Q.rear + 1) % MaxSize == Q.front){
        return false;
    }
    Q.data[Q.rear] == x;
    Q.rear = (Q.rear + 1) % MaxSize;
    return true;
}

bool DeQueue(SqQueue &Q, Elemtype x){
    if (Q.rear == Q.front){
        return false;
    }
    x = Q.data[Q.front];\
    Q.front = (Q.front + 1) % MaxSize;
    return true;
}

void InitLinkQueue(LinkQueue &Q){
    Q.front = Q.rear = (LinkNode *)malloc(sizeof(LinkNode));
    Q.front->next = NULL;
}

int main(){


    return 0;
}
