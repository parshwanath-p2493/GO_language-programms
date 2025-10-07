#include <iostream>
#include <cstring>
#include <string>
using namespace std;

// check the palindrome for the phrases remove the special characters also

// remove the special characters and also the spaces

bool Valid(char c)
{
    if ((c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9'))
    {
        return 1;
    }
    return 0;
}
char tolowercase(char c)
{
    if ((c >= 'a' && c <= 'z') || (c >= '0' && c <= '9'))
        return c;
    else
        return (c - 'A' + 'a');
}
bool isPalindrome(string S){
    int s=0;
    int e=S.length()-1;
    while (s<e)
    {
        if (S[s]!=S[e])
        {
            return 0;
        }
        else{
            s++;e--;
        }
        
    }
    return 1;
    
}

int main()
{
    string S;
 getline(cin,S);
    string temp="";
    cout<<temp<<endl;
  //  cout<<S<<endl;
    //cout<<S[2]<<endl;
    for (int i = 0; i < S.length(); i++)
    {
        if (Valid(S[i])){
            char correct= tolowercase(S[i]);
            temp.push_back(correct);
        }
           // temp.push_back(S[i]);
    }
  
    cout<<S<<endl;
    cout<<temp<<endl;
if(isPalindrome(temp)){
    std::cout << "The string is palindrome" << std::endl;
}
else{
    cout<<" The phrase is not palindrome"<<endl;
}

    return 0;
}