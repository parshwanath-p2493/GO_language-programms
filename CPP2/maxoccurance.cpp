#include <iostream>
#include <cstring>
#include <string>
using namespace std;

char MaxOccurance(string S){
int l=S.length();
int arr[26]={0};
cout<<arr<<endl<<endl;
for (int i=0;i<l;i++){
    int num=S[i]-'a';
    arr[num]++;
}

int max=-1;int index=-1;
for(int i=0;i<26;i++){
    if(arr[i]>max){
        max=arr[i];
        index=i;
    }
}
char x=index+'a';
return x;
}

int main(){
string S;
getline(cin,S);
string Temp="";
for (char c:S){
    if(isalpha(c)){
        tolower(c);
        Temp+=c;
    
    }

        
}
char H=MaxOccurance(Temp);
cout<<"Hoghest repeating character is "<<H<<endl;

    return 0;
}