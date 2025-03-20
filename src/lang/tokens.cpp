#include "tokens.h"
#include "ostream"

std::ostream& operator<<(std::ostream &os, TokenType type) {
    switch(type) {
        case TokenType::KEYWORD: return os << "KEYWORD";
        case TokenType::IDENT: return os << "IDENT";
        case TokenType::INT: return os << "INT";
        case TokenType::DOUBLE: return os << "DOUBLE";
        case TokenType::PLUS: return os << "PLUS";
        case TokenType::MINUS: return os << "MINUS";
        case TokenType::DIV: return os << "DIV";
        case TokenType::MULT: return os << "MULT";
        case TokenType::MOD: return os << "MOD";
        case TokenType::ASSIGN: return os << "ASSIGN";
        case TokenType::EQUALS: return os << "EQUALS";
        case TokenType::LSPAREN: return os << "LSPAREN";
        case TokenType::RSPAREN: return os << "RSPAREN";
        case TokenType::LPAREN: return os << "LPAREN";
        case TokenType::RPAREN: return os << "RPAREN";
        case TokenType::LCPAREN: return os << "LCPAREN";
        case TokenType::RCPAREN: return os << "RCPAREN";
        case TokenType::QUOTE: return os << "QUOTE";
        case TokenType::SQUOTE: return os << "SQUOTE";
        case TokenType::COMMA: return os << "COMMA";
        case TokenType::SEMI: return os << "SEMI";
        case TokenType::ILLEGAL: return os << "ILLEGAL";
        case TokenType::END_OF_FILE: return os << "END_OF_FILE";
    }
    return os << "UNKNOWN";
}

std::ostream& operator<<(std::ostream &os, const Token &token) {
    // return os << token.type << "\t" << token.value;
    return os << token.type;
}
