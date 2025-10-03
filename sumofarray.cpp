// Online C++ compiler to run C++ program online
#include <iostream>
#include<vector>
#include<algorithm>
using namespace std;
vector<int> Sum(vector <int>a, vector<int>b,int n,int m);
 //vector(vector<int> v);
 vector<int> my_reverse(vector<int> v);
int main() {
vector<int> arr1=  {2,3,5,6,8,9,8};
vector<int> arr2={2,4,5,7,2,3,5,6};
vector <int> ans;
int n=arr1.size(); 
int m=arr2.size();
vector<int> Sum(vector <int>a, vector<int>b,int n,int m);
 //vector(vector<int> v);
 vector<int> my_reverse(vector<int> &v);
 //Sum(arr)
 ans=Sum(arr1,arr2,n,m);
 for(const auto &x:ans){
     cout<<x<<" ";
 }

    return 0;
}
vector<int> Sum(vector <int>a, vector<int>b,int n,int m){
    int i=n-1;
    int j=m-1;
    int sum=0;
    int carry=0;
    vector<int> ans;
    while(i>=0&&j>=0){
        int val1=a[i]; int val2=b[j];
        sum =val1+val2+carry;
        carry=sum/10;
        sum=sum%10;
        ans.push_back(sum);i--;j--;
    }
    while(i>=0){
        sum=a[i]+carry;
        carry=sum/10;
        sum=sum%10;
        ans.push_back(sum);
        i--;
    }
    while(j>=0){
        sum=b[j]+carry;
        carry=sum/10;
        sum=sum%10;
        ans.push_back(sum);
        j--;
    }
   
    return my_reverse(ans);
}
vector<int> my_reverse(vector<int> v){
int s=0;
int e=v.size()-1;
while(s<e){
    swap(v[s++],v[e--]);
}
  return v;  
}
