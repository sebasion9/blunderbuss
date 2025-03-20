#pragma once

#include <string>

namespace LANG {
    const char OPERATORS[] = {
        '+', '-', '/', '*', '%',
    };
    const char SYMBOLS[] = {
        '=', '[', ']', '(', ')', '{', '}', '"', '\''
    };
    const std::string KEYWORDS[] = {
        "for", "if", "else", "return", "fn", "let"
    };

}

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

struct Token {
    TokenType type;
    std::string value;
    bool operator==(const Token &other) const = default;
};

std::ostream& operator<<(std::ostream &os, TokenType type);
std::ostream& operator<<(std::ostream &os, const Token &token);
