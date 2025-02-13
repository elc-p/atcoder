#include <iostream>
#include <vector>
#include <sstream>
#include <string>
#include <iterator>

using namespace std;

int main() {
  string s;
  cin >> s;

  vector<string> result;

  for (int i = 0; i < (int)s.length(); i++) {
    if (s.at(i) == '1')
    {
      result.push_back("1");
    } else if (s.at(i) == '0')
    {
      result.push_back("0");
    } else if (s.at(i) == 'B' && result.empty())
    {
    } else
    {
      result.pop_back();
    }
  }

  ostringstream os;
  copy(result.begin(), result.end(), ostream_iterator<string>(os));
  cout << os.str() << endl;
}