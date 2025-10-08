#include <iostream>
#include <cstring>
#include <string>
#include <algorithm>
using namespace std;
int main(){
    string S;
    getline(cin,S);
    cout<<S.length()<<endl;
        int s=0;
    for (int i = 0; i <=S.length(); i++)
    {
     
        if ((S[i]=='\0')||(S[i]==' '))
        {
         int e=i-1;
            while (s<e)
            {
                swap(S[s++],S[e--]);
            }
            s=i+1;
        }

    }
    cout<<S<<endl;
    
    
    return 0;

}

/*

WE LOVE INDIA 
13
EW EVOL AIDNI

wewewe love love india indiaindia
33
ewewew evol evol aidni aidniaidni

*/