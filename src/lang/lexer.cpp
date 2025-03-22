#include "lexer.h"
#include "../util.h"
#include "tokens.h"
#include <cctype>
#include <cstdlib>
#include <string>


Lexer::Lexer(const std::string &input)
    : input(input), pos(0), nextPos(1) {
        if(this->input.length() < 2) {
            this->currChar = '\0';
        } else {
            this->currChar = input[0];
        }
    };
void Lexer::setPos(size_t newpos) {
    this->pos = newpos;
    this->nextPos = newpos+1;

}

Token Lexer::nextToken() {
    if(isspace(this->currChar)) this->skipWhitespace();

    if(this->currChar == '\0') {
        this->advanceChar();
        return Token(TokenType::END_OF_FILE);
    }

    if (isdigit(this->currChar)) {
        double* d = this->parseDouble();
        if(d != nullptr) {
            return Token(TokenType::DOUBLE, *d);
        }
        int* i = this->parseInt();
        if(i != nullptr) {
            return Token(TokenType::INT, *i);
        }
        return Token(TokenType::ILLEGAL);
    }

    if (isalpha(this->currChar)) {
        auto alpha = this->readAlpha();

        std::string keyword;
        for(char c : alpha) {
            keyword.push_back(tolower(c));
        }

        if(contains<std::string>(LANG::KEYWORDS, keyword)) {
            return Token(TokenType::KEYWORD, alpha);
        }
        return Token(TokenType::IDENT, alpha);
    }

    if(contains<char>(LANG::OPERATORS, this->currChar)) {
        auto op = this->readOperator();
        return Token(op);
    }

    if(contains<char>(LANG::SYMBOLS, this->currChar)) {
        auto symbol = this->readSymbol();
        return Token(symbol);
    }



    std::string val{this->currChar};
    this->advanceChar();
    return Token(TokenType::ILLEGAL, val);
}
std::vector<Token> Lexer::tokenize() {
    std::vector<Token> tokens;
    tokens.reserve(this->input.length());
    for(;;) {
        auto token = this->nextToken();
        if(token.type == TokenType::END_OF_FILE) {
            break;
        }
        tokens.push_back(token);
    }
    return tokens;
}

void Lexer::advanceChar() {
    if(this->nextPos >= this->input.length()) {
        this->currChar = '\0';
    } else {
        this->currChar = this->input[this->nextPos];
    }
    this->pos = this->nextPos;
    this->nextPos++;
}
void Lexer::retreatChar() {
    if(this->pos > 0) {
        this->pos--;
        this->nextPos--;
    }
    this->currChar = input[this->pos];
}

TokenType Lexer::readOperator() {
    auto currChar = this->currChar;
    this->advanceChar();
    switch(currChar) {
        case '+': return TokenType::PLUS;
        case '-': return TokenType::MINUS;
        case '/': return TokenType::DIV;
        case '*': return TokenType::MULT;
        case '%': return TokenType::MOD;
    }
    return TokenType::ILLEGAL;
}

TokenType Lexer::readSymbol() {
    auto currChar = this->currChar;
    this->advanceChar();
    switch(currChar) {
        case '[': return TokenType::LSPAREN;
        case ']': return TokenType::RSPAREN;
        case '(': return TokenType::LPAREN;
        case ')': return TokenType::RPAREN;
        case '{': return TokenType::LCPAREN;
        case '}': return TokenType::RCPAREN;
        case '"': return TokenType::QUOTE;
        case ';': return TokenType::SEMI;
        case ',': return TokenType::COMMA;
        case '\'': return TokenType::SQUOTE;
        case '=': {
            if(this->currChar == '=') {
                this->advanceChar();
                return TokenType::EQUALS;
            }
            return TokenType::ASSIGN;
        }
    }
    return TokenType::ILLEGAL;
}
double* Lexer::parseDouble() {
    size_t start = this->pos;
    while(isdigit(this->currChar)) {
        this->advanceChar();
    }
    if(this->currChar != '.') {
        for(size_t i = 0; i < this->pos - start; i++) {
            this->retreatChar();
        }
        return nullptr;
    } else this->advanceChar();
    while(isdigit(this->currChar)) {
        this->advanceChar();
    }
    auto str = this->input.substr(start, (this->pos - start));
    double d = stod(str);
    double* d_ptr = &d;
    return d_ptr;
}

int* Lexer::parseInt() {
    size_t start = this->pos;
    while(isdigit(this->currChar)) {
        this->advanceChar();
    }
    auto str = this->input.substr(start, (this->pos - start));
    int i = stoi(str);
    int* i_ptr = &i;
    return i_ptr;
}


std::string Lexer::readAlpha() {
    size_t start = this->pos;
    while(isalpha(this->currChar)) {
        this->advanceChar();
    }
    auto alpha = this->input.substr(start, (this->pos - start));
    return alpha;
}


void Lexer::skipWhitespace() {
    while(isspace(this->currChar)) {
        this->advanceChar();
    }
}











