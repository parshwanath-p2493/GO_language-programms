#include <iostream>
#include <cstring>
#include <string>
#include <vector>
using namespace std;

bool searchMatrix(vector<vector<int>> &matrix, int target)
{
    int rows = matrix.size();
    int cols = matrix[0].size();

    int left = 0;
    int right = rows * cols - 1;

    while (left <= right)
    {
        int mid = left + (right - left) / 2;
        int midElement = matrix[mid / cols][mid % cols];

        if (midElement == target)
            return true; // Found
        else if (midElement < target)
            left = mid + 1;
        else
            right = mid - 1;
    }

    return false;
}
int main()
{
    int n, m;
    cin >> m >> n;
    vector<vector<int>> matrix(m, vector<int>(n));
    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {

            cin >> matrix[i][j];
        }
    }

    for (int i = 0; i < m; i++)
    {
        for (int j = 0; j < n; j++)
        {

            cout << matrix[i][j] << " ";
        }
        cout << endl;
    }
   bool ans;
    ans = searchMatrix(matrix);
    if(ans){
        cout<<"the elemt is present ";
    }
    else{
        cout<<"Element is not present ";
    }

    return 0;
}
