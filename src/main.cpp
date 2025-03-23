#include "lang/lexer.h"
#include "lang/parser.h"
#include <fstream>
#include <iostream>
#include <cstdio>
#include <sstream>

const std::string INPUT_PATH = "input.bs";

int main() {

    // read input
    // std::ifstream t(INPUT_PATH);
    // std::stringstream buffer;
    // buffer << t.rdbuf();
    // auto input = buffer.str();
    // //

    auto input = "for let a = 2; a < 1; a == a + 1;";

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();

    return 0;
}
