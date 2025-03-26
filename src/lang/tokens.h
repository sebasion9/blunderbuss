#pragma once

#include <string>
#include <variant>

namespace LANG {
    const char OPERATORS[] = {
        '+', '-', '/', '*', '%',
    };
    // todo: add < > <= >=
    const char SYMBOLS[] = {
        '=', '[', ']', '(', ')', '{', '}', '"', '\'', ',', ';', '<', '>', '&', '|', '!', '^',
    };
    const std::string KEYWORDS[] = {
        "for", "if", "else", "return", "fn", "let",
    };
    // let ident assign expression
    // for assign_stmt semi condition(expression != 0) assign_stmt -> block
    // if expression != 0 -> block
    // else -> block
    // return expression
    // fn ident ???

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
    ASSIGN, LSPAREN, RSPAREN, LPAREN, RPAREN, LCPAREN, RCPAREN, QUOTE, SQUOTE, COMMA, SEMI,

    // comp
    EQUALS, LESS_THAN, GREATER_THAN, LESS_THAN_EQUAL, GREATER_THAN_EQUAL,

    // logic
    OR, AND, NOT, XOR,

    // other
    ILLEGAL,
    END_OF_FILE,

};

bool isSymbol(const char &token);
bool isExprOperator(const TokenType &type);


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
