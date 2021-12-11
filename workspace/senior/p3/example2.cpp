// example2.cpp
#include <iostream>

using namespace std;

/*函数changeValue使用引用传递*/
void changeValue(int &n) {
    n = 2;
}

int main() {
    int a = 1;
    /*
    b是引用变量，引用的是变量a
    */
    int &b = a;
    cout << "a=" << a << " address:" << &a << endl;
    cout << "b=" << b << " address:" << &b << endl;
    /*
    调用changeValue会改变外部实参a的值
    */
    changeValue(a);
    cout << "a=" << a << " address:" << &a << endl;
    cout << "b=" << b << " address:" << &b << endl;
}