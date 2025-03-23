#include "tokens.h"
#include "ostream"
#include "../util.h"

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
        case TokenType::LESS_THAN: return os << "LESS_THAN";
        case TokenType::LESS_THAN_EQUAL: return os << "LESS_THAN_EQUAL";
        case TokenType::GREATER_THAN: return os << "GREATER_THAN";
        case TokenType::GREATER_THAN_EQUAL: return os << "GREATER_THAN_EQUAL";
        case TokenType::NOT: return os << "NOT";
        case TokenType::OR: return os << "OR";
        case TokenType::XOR: return os << "XOR";
        case TokenType::AND: return os << "AND";
    }
    return os << "UNKNOWN";
}

std::ostream& operator<<(std::ostream &os, const Token &token) {
    // return os << token.type << "\t" << token.value;
    return os << token.type;
}

bool isSymbol(const char &ch) {
    return contains<char>(LANG::SYMBOLS, ch);
}

bool isExprOperator(const TokenType &type) {
    switch(type) {
        case TokenType::AND: return true;
        case TokenType::OR: return true;
        case TokenType::NOT: return true;
        case TokenType::XOR: return true;
        case TokenType::EQUALS: return true;
        case TokenType::LESS_THAN_EQUAL: return true;
        case TokenType::LESS_THAN: return true;
        case TokenType::GREATER_THAN_EQUAL: return true;
        case TokenType::GREATER_THAN: return true;

        case TokenType::PLUS: return true;
        case TokenType::MINUS: return true;
        case TokenType::DIV: return true;
        case TokenType::MULT: return true;
        case TokenType::MOD: return true;
        default: return false;
    }
}



