#include <iostream>

using namespace std;

int main() {
  double n;
  cin >> n;

  double result = (1.0/2.0)*n*(1.0+n);

  cout << static_cast<int>(result) << endl;
}