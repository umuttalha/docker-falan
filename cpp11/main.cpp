#include <iostream>
#include <algorithm>
#include <vector>
#include <thread>
#include <chrono>
int main() {
    std::vector<int> numbers = {1, 2, 3, 4, 5};
    std::random_shuffle(numbers.begin(), numbers.end());  // removed in C++17
    
    for(int n : numbers) {
        std::cout << n << " ";
    }

    std::cout << "Program continues after sleep." << std::endl;

    // Sleep for 2 seconds
    std::this_thread::sleep_for(std::chrono::seconds(10));



    return 0;
}