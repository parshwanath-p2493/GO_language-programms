#include <iostream>
#include <cstring>
#include <string>
using namespace std;

int main(){
 string S;
 getline(cin,S);
//  int len=strlen(S);   //// can be used only for array of char
 int k=S.length();
 string K="";
 //cout<<len<<endl;
 cout<<k<<endl;
 cout<<"Before"<<S<<endl;
 int i=0;
 while(S[i]!='\0'){
    if(S[i]==' '){
         K.push_back('@');
         K.push_back('4');
          K.push_back('0');
    }
    else
    K.push_back(S[i]);
    i++;
 }
 cout<<"After"<<K<<endl;


    return 0;
}