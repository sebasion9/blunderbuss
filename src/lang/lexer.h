#pragma once
#include "tokens.h"

#include <span>
#include <string>
#include <concepts>
#include <iostream>
#include <vector>


class Lexer {
    private:
        std::string input;
        size_t pos;
        size_t nextPos;
        char currChar;

    public:
        // lexer fields
        Lexer(const std::string &input);
        void setPos(size_t newpos);

        // creating tokens
        Token nextToken();
        std::vector<Token> tokenize();

        // moving around input
        void advanceChar();
        void retreatChar();
        void skipWhitespace();

        // abstracting text to tokens
        std::string readAlpha();
        TokenType readOperator();
        TokenType readSymbol();
        double* parseDouble();
        int* parseInt();


};




