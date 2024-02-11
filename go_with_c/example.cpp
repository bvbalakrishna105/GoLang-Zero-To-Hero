#include <iostream>

extern "C" {
    void hello() {
        std::cout << "Hello from C++" << std::endl;
        std::cout << "Beesetti Venkata Balakrushna" << std::endl;
    }

    void showName(){
        std::cout << "showName Beesetti" << std::endl;
    }

    int addNum(int a, int b) {
        return a+b;
    }

}
