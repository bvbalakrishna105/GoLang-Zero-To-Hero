#include <iostream>
#include <dlfcn.h>

int main() {
    void *handle;
    handle = dlopen("/home/vidkrix/Desktop/work/GoLang-Zero-To-Hero/go_with_c/libexample.so", RTLD_LAZY); // Replace "libexample.so" with your shared library name

    if (!handle) {
        std::cerr << "Error: " << dlerror() << std::endl;
        return 1;
    }

    // Use dlsym to get the address of a function or variable in the shared library
    // For example:
    void (*hello)() = (void (*)())dlsym(handle, "hello");

    // Call a function from the shared library
    hello(); 
    std::cout <<"Sucessfully loaded .so file" << std::endl;

    dlclose(handle);
    return 0;
}
