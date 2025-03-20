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
        void setPos(size_t newpos);
        Token nextToken();
        void advanceChar();
        void retreatChar();
        void skipWhitespace();

        std::string readAlpha();
        TokenType readOperator();
        TokenType readSymbol();

        double* parseDouble();
        int* parseInt();

};




