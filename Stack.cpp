#include <iostream>

#define MaxSize 50
#define Elemtype int

typedef struct
{
    Elemtype data[MaxSize];
    int top;
} SqStack;

typedef struct Linknode
{
    Elemtype data;
    struct Linknode *next;
} *LiStack;


void InitStack(SqStack &S){
    S.top = -1;
}

bool StackEmpty(SqStack &S){
    if (S.top == -1){
        return true;
    }
    else {
        return false;
    }
}

bool Push(SqStack &S, Elemtype x){
    if (S.top == MaxSize-1){
        return false;
    }
    S.data[++S.top] = x;
    return true;
}

bool Pull(SqStack &S, Elemtype x){
    if (S.top == -1){
        return false;
    }
    x = S.data[S.top--];
    return true;
}

bool GetTop(SqStack &S, Elemtype x){
    if (S.top == -1){
        return false;
    }
    x = S.data[S.top];
    return true;
}

int main()
{


    return 0;
}
