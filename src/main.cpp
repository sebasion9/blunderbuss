#include "lexer.h"
#include <iostream>
#include <cstdio>

int main() {
    auto input = "10;";
    auto lexer = Lexer(input);
    auto token = lexer.nextToken();

    std::cout << token.value << '\n';
    token = lexer.nextToken();


    return 0;
}
