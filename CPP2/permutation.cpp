#include <iostream>
#include <cstring>
#include <string>
using namespace std;
bool CheckEqual(int arr[26], int brr[26])
{
    for (int i = 0; i < 26; i++)
        if (arr[i] != brr[i])
            return false;

    return true;
}

bool CheckPalindrome(string s1, string s2)
{
    int windowsize = s1.length();
    int count1[26] = {0};
    for (int i = 0; i < s.length(); i++)
    {
        int index = s1[i] - 'a';
        count1[index]++;
    }
    int i = 0;
    int count2[26] = {0};
    while (i < windowsize && i < s2.length())
    {
        int index = s2[i] - 'a';
        count2[index]++;
        i++;
    }
    if (CheckEqual(count1, count2))
        return true;

    while (i < s2.length())
    {
        char newchar = s2[i];
        int index = newchar - 'a';
        count2[index]++;

        char oldchar = s2[i - windowsize];
        int index = oldchar - 'a';
        count2[index]--;

        i++;
        if (CheckEqual(count1, count2))
            return true;
    }
    return false;
}
int main()
{
    string S1;
    string S2;
    getline(cin, S1);
    getline(cin, S2);

    if (CheckPalindrome(S1,S2))
    cout<<"The string is palindrome"<<endl;
    else 
    cout<<"The string is not  palindrome"<<endl;
        return 0;
}
