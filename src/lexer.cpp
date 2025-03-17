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
    Token token;
    if(this->currChar == ';') {
        token = Token {
            .type = TokenType::SEMI,
            .value = ";"
        };
    } else if (std::isdigit(this->currChar)) {
        // read whole number
        auto num = this->readNumber();
        return Token {
            .type = TokenType::INT,
            .value = num
        };
    } else {
        token = Token {
            .type = TokenType::ILLEGAL,
            .value = std::to_string(this->currChar),
        };
    }

    this->advanceChar();

    return token;
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
    while(std::isdigit(this->currChar)) {
        this->advanceChar();
    }
    auto num = this->input.substr(start, this->pos);
    return num;
}












