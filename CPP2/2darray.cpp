#include <iostream>
#include <cstring>
#include <string>
using namespace std;

void WavePrint(int arr[][100], int nrow, int mcol)
{
    for (int col = 0; col < mcol; col++)
    {
        if (col & 1)
        {
            for (int row = nrow - 1; row >= 0; row--)
            {
                cout << arr[row][col];
            }
        }
        else
        {
            for (int row = 0; row < nrow; row++)
            {
                cout << arr[row][col];
            }
        }
    }
}
int main()
{
    int n, m;
    cin >> n >> m;
    int arr[n][100];
    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < m; j++)
        {

            cin >> arr[i][j];
        }
    }

    for (int i = 0; i < n; i++)
    {
        for (int j = 0; j < m; j++)
        {

            cout << arr[i][j] << " ";
        }
        cout << endl;
    }

    for (int j = 0; j < m; j++)
    {
        for (int i = 0; i < n; i++)
        {

            cout << arr[i][j] << " ";
        }
        cout << endl;
    }
    
    WavePrint(arr,n,m);
    return 0;
}