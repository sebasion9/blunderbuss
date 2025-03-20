#include "tokens.h"
#include "ostream"

std::ostream& operator<<(std::ostream &os, TokenType type) {
    switch(type) {
        case TokenType::KEYWORD: return os << "KEYWORD";
        case TokenType::IDENT: return os << "IDENT";
        case TokenType::INT: return os << "INT";
        case TokenType::OPERATOR: return os << "OPERATOR";
        case TokenType::SYMBOL: return os << "SYMBOL";
        case TokenType::SEMI: return os << "SEMI";
        case TokenType::ILLEGAL: return os << "ILLEGAL";
        case TokenType::END_OF_FILE: return os << "END_OF_FILE";
    }
    return os << "UNKNOWN";
}

std::ostream& operator<<(std::ostream &os, const Token &token) {
    return os << token.type << "\t" << token.value;
}
