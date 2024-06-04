int var[30000000] = {1};

int main(int artc, char **argv) {
  while (true)
    ;
  return 0;
}

struct {           // Structure declaration
  int myNum;       // Member (int variable)
  string myString; // Member (string variable)
} myStructure;     // Structure variable

class MyClass {    // The class
public:            // Access specifier
  void myMethod(); // Method/function declaration
};

// Method/function definition outside the class
void MyClass::myMethod() { cout << "Hello World!"; }

#define MACRO(X, Y)                                                            \
  do {                                                                         \
    cout << "1st arg is:" << (X) << endl;                                      \
    cout << "2nd arg is:" << (Y) << endl;                                      \
    cout << "Sum is:" << ((X) + (Y)) << endl;                                  \
  } while (0)
