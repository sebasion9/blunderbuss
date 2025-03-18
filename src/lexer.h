#pragma once

#include <span>
#include <string>
#include <concepts>
#include <iostream>
enum class TokenType {
    KEYWORD,
    IDENT,
    INT,
    OPERATOR,
    SYMBOL,
    SEMI,
    ILLEGAL,
    END_OF_FILE,
};

const char OPERATORS[] = {
    '+', '-', '/', '*', '%',
};
const char SYMBOLS[] = {
    '=', '[', ']', '(', ')', '{', '}'
};
const std::string KEYWORDS[] = {
    "for", "if", "else", "return", "fn"
};

template <typename T>
requires std::equality_comparable<T>
bool contains(std::span<const T> arr, const T& el) {
    for(const auto& e: arr) {
        if (e == el) return true;
    }
    return false;
}

struct Token {
    TokenType type;
    std::string value;
};



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


