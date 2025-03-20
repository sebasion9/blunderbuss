#include "lang/lexer.h"
#include <fstream>
#include <iostream>
#include <cstdio>
#include <sstream>

const std::string INPUT_PATH = "input.bs";

int main() {

    // read input
    std::ifstream t(INPUT_PATH);
    std::stringstream buffer;
    buffer << t.rdbuf();
    auto input = buffer.str();
    //

    auto lexer = Lexer(input);
    auto token = lexer.nextToken();
    std::cout << token.type;

    while(token.type != TokenType::ILLEGAL && token.type != TokenType::END_OF_FILE) {
        token = lexer.nextToken();
    }


    return 0;
}
