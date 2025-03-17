#pragma once


#include <string>
#include <string_view>
enum class TokenType {
    KEYWORD = 0,
    IDENTIFIER,
    OPERATOR,
    INT,
    SYMBOL,
    SEMI,
    ILLEGAL,
    END_OF_FILE,
};
constexpr std::string_view TokenTypes[] = {
    "KEYWORD",
    "IDENTIFIER",
    "OPERATOR",
    "INT",
    "SYMBOL",
    "SEMI",
    "ILLEGAL",
    "END_OF_FILE",
};

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
};


