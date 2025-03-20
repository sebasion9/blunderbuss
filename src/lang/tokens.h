#pragma once

#include <string>
#include <variant>

namespace LANG {
    const char OPERATORS[] = {
        '+', '-', '/', '*', '%',
    };
    // todo: add < > <= >=
    const char SYMBOLS[] = {
        '=', '[', ']', '(', ')', '{', '}', '"', '\'', ',', ';',
    };
    const std::string KEYWORDS[] = {
        "for", "if", "else", "return", "fn", "let",
    };

}

enum class TokenType {
    // str tokens
    KEYWORD,
    IDENT,

    INT,
    DOUBLE,

    // operators
    PLUS, MINUS, DIV, MULT, MOD,

    // symbols
    ASSIGN, EQUALS, LSPAREN, RSPAREN, LPAREN, RPAREN, LCPAREN, RCPAREN, QUOTE, SQUOTE, COMMA, SEMI,

    // other
    ILLEGAL,
    END_OF_FILE,
};


struct Token {
    TokenType type;
    std::variant<std::monostate, std::string, int, double> value;

    Token(TokenType type) : type(type), value(std::monostate{}) {};
    Token(TokenType type, std::string val) : type(type), value(val) {};
    Token(TokenType type, int val) : type(type), value(val) {};
    Token(TokenType type, double val) : type(type), value(val) {};

    bool operator==(const Token &other) const = default;
};

std::ostream& operator<<(std::ostream &os, TokenType type);
std::ostream& operator<<(std::ostream &os, const Token &token);
