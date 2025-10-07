#include<iostream>
#include<cstring>
using namespace std;

int main(){
    char c[100];
    cin>>c;
    char z[100];
    strcpy(z,c);
    //reverse the STRING

    int s=0; int e=strlen(z)-1;
    while (s<e)
    {
        /* code */
        swap(z[s++],z[e--])
    }
    cout<<z<<endl;
    
//length of string
int count =0;
for(int i=0;i=='\0';i++)
++count;
cout <<"length of string is "<<count<<endl;


// // check for palindrome

//before checking make all the characters small case
for (int i = 0; c[i]!='\0'; i++)
{
    c[i]=tolower(c[i]);
}

 int st=0; int e=strlen(z)-1;
 int palindrome=1;
    while (s<e)
    {
        /* code */
        if(c[st]!=c[en]){
            Palindrome =0;
            break;
        }
        else{
            st++;en--;
        }

    }
   if(palindrome)
   cout<<"The string is palindrome "<<endl;
   else
   cout<<"The string is not palindrome "<<endl;

    return 0;
}