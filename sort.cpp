#include <stdio.h>
#include <iostream>

using namespace std;

typedef int ElemType;

void Insertion_Sort(ElemType A[], int n)
{
    ElemType tmp;
    int i, p;
    for (p = 1; p < n; p++)
    {
        tmp = A[p];
        for (i = p; i > 0 && A[i - 1] > tmp; i--)
        {
            A[i] = A[i - 1];
        }
        A[i] = tmp;
    }
}

void InsertHalf_Sort(ElemType A[], int n)
{
    ElemType tmp;
    int i, j, low, mid, high;
    for (i = 1; i < n; i++)
    {
        low = 0;
        high = i - 1;
        tmp = A[i];
        while (low <= high)
        {
            mid = low + (high - low) / 2;
            if (A[mid] > tmp)
            {
                high = mid - 1;
            }
            else
                low = mid + 1;
        }
        for (j = i - 1; j >= high + 1; j--)
        {
            A[j + 1] = A[j];
        }
        A[high + 1] = tmp;
    }
}

void Bubble_Sort(ElemType A[], int n)
{
    for (int i = 0; i < n; i++)
    {
        bool flag = false;
        for (int j = n - 1; j < i; j--)
        {
            if (A[j - 1] > A[j])
            {
                int temp = A[j];
                A[j] = A[j - 1];
                A[j - 1] = temp;
                flag = true;
            }
        }
        if (flag == false)
        {
            return;
        }
    }
}

int Partition(ElemType A[], int low, int high)
{
    int pivot = high;
    int i, j = low;
    while (j < high)
    {
        if (A[j] < A[pivot])
        {
            int temp = A[j];
            A[j] = A[i];
            A[i] = temp;
            i++;
            j++;
        }
        else
        {
            j++;
        }
    }
    int temp = A[i];
    A[i] = A[high];
    A[high] = temp;

    return i;
}

void Quick_Sort_c(ElemType A[], int low, int high)
{
    if (low < high)
    {
        int pivot = Partition(A, low, high);
        Quick_Sort_c(A, low, pivot - 1);
        Quick_Sort_c(A, pivot + 1, high);
    }
}

void Quick_Sort(ElemType A[], int n)
{
    Quick_Sort_c(A, 0, n - 1);
}

int main()
{
    ElemType A[10] = {3, 4, 2, 6, 8, 1, 7, 9, 5};
    int length = 9;
    // Insertion_Sort(A, 9);
    // InsertHalf_Sort(A, length);
    Quick_Sort(A, length);
    for (int i = 0; i < length; i++)
    {
        cout << A[i] << endl;
    }
}
