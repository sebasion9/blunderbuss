#include "lang/lexer.h"
#include "lang/parser.h"
#include "util/json_util.h"
#include <fstream>
#include <iostream>
#include <cstdio>
#include <ostream>
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
    auto top_level = parser.parse_block();

    auto stringified = node_vec_to_json(top_level).dump(4);
    std::cout << stringified << std::endl;

    return 0;
}
