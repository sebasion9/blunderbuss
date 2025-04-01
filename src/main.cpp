#include "lang/lexer.h"
#include "lang/parser.h"
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
    // //


    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto topLevel = parser.parse_block();



    return 0;
}
