#include<iostream>
#include"Queue.h"

using namespace std;

#define ElemType int

typedef struct BiTNode
{
    ElemType data;
    struct BiTNode *lchild, *rchild;
}BiTNode, *BiTree;

void visit(BiTree T){
    cout << T->data << endl;
}

void PreOrder(BiTree T){
    if(T != NULL){
        visit(T);
        PreOrder(T->lchild);
        PreOrder(T->rchild);
    }
}

void InOrder(BiTree T){
    if(T != NULL){
        InOrder(T->lchild);
        visit(T);
        InOrder(T->rchild);
    }
}

void PostOrder(BiTree T){
    if (T != NULL)
    {
        PostOrder(T->lchild);
        PostOrder(T->rchild);
        visit(T);
    }

}

void LevelOrder(BiTree T){
    SqQueue Q;
    InitQueue(Q);
    BiTree p;
    EnQueue(Q, T);
    while (!IsEmpty(Q))
    {
        DeQueue(Q, p);
        visit(p);
        if(!p->lchild){
            EnQueue(Q, p->lchild);
        }
        if(!p->rchild){
            EnQueue(Q, p->rchild);
        }
    }

}

int main(){
    BiTree T;

    return 0;
}
