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
        size_t next_pos;
        char curr_char;

    public:
        // lexer fields
        Lexer(const std::string &input);
        void set_pos(size_t newpos);

        // creating tokens
        Token next_token();
        std::vector<Token> tokenize();

        // moving around input
        void advance_char();
        void retreat_char();
        void skip_ws();

        // abstracting text to tokens
        std::string read_alpha();
        TokenType read_op();
        TokenType read_sym();
        double* parse_double();
        int* parse_int();


};




