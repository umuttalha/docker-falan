#include <iostream>
#include <tuple>
#include <thread>
#include <chrono>

int main() {
    auto tuple = std::make_tuple(1, "test", 3.14);
    auto [x, y, z] = tuple;  // structured binding, C++17 only
    std::cout << x << " " << y << " " << z << std::endl;

    std::cout << "Program continues after sleep." << std::endl;


    // Sleep for 2 seconds
    std::this_thread::sleep_for(std::chrono::seconds(10));


    return 0;
}