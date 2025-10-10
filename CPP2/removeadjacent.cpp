#include <iostream>
#include <cstring>
#include <string>
using namespace std;

void RemoveAdj(string s)
{
    string ans = "";
    for (int i = 0; i < s.length(); i++)
    {
        if (ans.empty())
            ans.push_back(s[i]);
        else if(s[i] == ans.back())
            ans.pop_back();
        else if(s[i] != ans.back())
            ans.push_back(s[i]);
    }
    cout << ans;
}

int main()
{
    string s;
    getline(cin, s);
    cout << RemoveAdj(s) << endl;
    return 0;
}