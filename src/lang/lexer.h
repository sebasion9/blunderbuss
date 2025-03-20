#pragma once
#include "tokens.h"

#include <span>
#include <string>
#include <concepts>
#include <iostream>


class Lexer {
    std::string input;
    size_t pos;
    size_t nextPos;
    char currChar;

    public:
        Lexer(const std::string &input);
        Token nextToken();
        void advanceChar();
        std::string readNumber();
        std::string readSymbol();
        std::string readAlpha();
        void skipWhitespace();
};




