#include "lexer.h"
#include <cctype>
#include <string>


Lexer::Lexer(const std::string &input)
    : input(input), pos(0), nextPos(1) {
        if(this->input.length() < 2) {
            this->currChar = '\0';
        } else {
            this->currChar = input[0];
        }
    };

Token Lexer::nextToken() {
    if(isspace(this->currChar)) this->skipWhitespace();

    if(this->currChar == ';') {
        this->advanceChar();
        return Token(TokenType::SEMI, ";");
    }

    if (isdigit(this->currChar)) {
        auto num = this->readNumber();
        return Token(TokenType::INT, num);
    }

    if (isalpha(this->currChar)) {
        auto alpha = this->readAlpha();

        // check if keyword
        std::string keyword;
        for(char c : alpha) {
            keyword.push_back(tolower(c));
        }

        if(contains<std::string>(KEYWORDS, keyword)) {
            return Token(TokenType::KEYWORD, alpha);
        }
        return Token(TokenType::IDENT, alpha);
    }

    if(contains<char>(OPERATORS, this->currChar)) {
        std::string val{this->currChar};
        this->advanceChar();
        return Token(TokenType::OPERATOR, val);
    }

    if(contains<char>(SYMBOLS, this->currChar)) {
        auto symbol = this->readSymbol();
        return Token(TokenType::SYMBOL, symbol);
    }


    if(this->currChar == '\0') {
        std::string val{this->currChar};
        this->advanceChar();
        return Token(TokenType::END_OF_FILE, val);
    }

    this->advanceChar();
    return Token(TokenType::ILLEGAL, std::string{this->currChar});
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

std::string Lexer::readNumber() {
    size_t start = this->pos;
    while(isdigit(this->currChar)) {
        this->advanceChar();
    }
    auto num = this->input.substr(start, (this->pos - start));
    return num;
}
std::string Lexer::readAlpha() {
    size_t start = this->pos;
    while(isalpha(this->currChar)) {
        this->advanceChar();
    }
    auto alpha = this->input.substr(start, (this->pos - start));
    return alpha;
}

std::string Lexer::readSymbol() {
    std::string sym = std::string{this->currChar};
    if(this->currChar == '=') {
        sym = std::string{this->currChar};
        this->advanceChar();
        if(this->currChar == '=') {
            sym.push_back(this->currChar);
            this->advanceChar();
        }
    } else {
        this->advanceChar();
    }
    return sym;
}

void Lexer::skipWhitespace() {
    while(isspace(this->currChar)) {
        this->advanceChar();
    }
}











