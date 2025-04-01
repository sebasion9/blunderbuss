#include "lexer.h"
#include "../util.h"
#include "tokens.h"
#include <cctype>
#include <cstdlib>
#include <string>


Lexer::Lexer(const std::string &input)
    : input(input), pos(0), next_pos(1) {
        if(this->input.length() < 2) {
            this->curr_char = '\0';
        } else {
            this->curr_char = input[0];
        }
    };
void Lexer::set_pos(size_t newpos) {
    this->pos = newpos;
    this->next_pos= newpos+1;

}

Token Lexer::next_token() {
    if(isspace(this->curr_char)) this->skip_ws();

    if(this->curr_char == '\0') {
        this->advance_char();
        return Token(TokenType::END_OF_FILE);
    }

    if (isdigit(this->curr_char)) {
        double* d = this->parse_double();
        if(d != nullptr) {
            return Token(TokenType::DOUBLE, *d);
        }
        int* i = this->parse_int();
        if(i != nullptr) {
            return Token(TokenType::INT, *i);
        }
        return Token(TokenType::ILLEGAL);
    }

    if (isalpha(this->curr_char)) {
        auto alpha = this->read_alpha();

        std::string keyword;
        for(char c : alpha) {
            keyword.push_back(tolower(c));
        }

        if(contains<std::string>(LANG::KEYWORDS, keyword)) {
            return Token(TokenType::KEYWORD, alpha);
        }
        return Token(TokenType::IDENT, alpha);
    }

    if(contains<char>(LANG::OPERATORS, this->curr_char)) {
        auto op = this->read_op();
        return Token(op);
    }

    if(is_symbol(this->curr_char)) {
        auto symbol = this->read_sym();
        return Token(symbol);
    }



    std::string val{this->curr_char};
    this->advance_char();
    return Token(TokenType::ILLEGAL, val);
}
std::vector<Token> Lexer::tokenize() {
    std::vector<Token> tokens;
    tokens.reserve(this->input.length());
    for(;;) {
        auto token = this->next_token();
        if(token.type == TokenType::END_OF_FILE) {
            break;
        }
        tokens.push_back(token);
    }
    return tokens;
}

void Lexer::advance_char() {
    if(this->next_pos >= this->input.length()) {
        this->curr_char = '\0';
    } else {
        this->curr_char = this->input[this->next_pos];
    }
    this->pos = this->next_pos;
    this->next_pos++;
}
void Lexer::retreat_char() {
    if(this->pos > 0) {
        this->pos--;
        this->next_pos--;
    }
    this->curr_char = input[this->pos];
}

TokenType Lexer::read_op() {
    auto curr_char = this->curr_char;
    this->advance_char();
    switch(curr_char) {
        case '+': return TokenType::PLUS;
        case '-': return TokenType::MINUS;
        case '/': return TokenType::DIV;
        case '*': return TokenType::MULT;
        case '%': return TokenType::MOD;
    }
    return TokenType::ILLEGAL;
}

TokenType Lexer::read_sym() {
    auto curr_char = this->curr_char;
    this->advance_char();
    switch(curr_char) {
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
        case '&': return TokenType::AND;
        case '|': return TokenType::OR;
        case '!': return TokenType::NOT;
        case '^': return TokenType::XOR;
        case '=': {
            if(this->curr_char == '=') {
                this->advance_char();
                return TokenType::EQUALS;
            }
            return TokenType::ASSIGN;
        }
        case '<': {
            if(this->curr_char == '=') {
                this->advance_char();
                return TokenType::LESS_THAN_EQUAL;
            }
            return TokenType::LESS_THAN;
        }
        case '>': {
            if(this->curr_char == '=') {
                this->advance_char();
                return TokenType::GREATER_THAN_EQUAL;
            }
            return TokenType::GREATER_THAN;
        }
    }
    return TokenType::ILLEGAL;
}
double* Lexer::parse_double() {
    size_t start = this->pos;
    while(isdigit(this->curr_char)) {
        this->advance_char();
    }
    if(this->curr_char != '.') {
        for(size_t i = 0; i < this->pos - start; i++) {
            this->retreat_char();
        }
        return nullptr;
    } else this->advance_char();
    while(isdigit(this->curr_char)) {
        this->advance_char();
    }
    auto str = this->input.substr(start, (this->pos - start));
    double d = stod(str);
    double* d_ptr = &d;
    return d_ptr;
}

int* Lexer::parse_int() {
    size_t start = this->pos;
    while(isdigit(this->curr_char)) {
        this->advance_char();
    }
    auto str = this->input.substr(start, (this->pos - start));
    int i = stoi(str);
    int* i_ptr = &i;
    return i_ptr;
}


std::string Lexer::read_alpha() {
    size_t start = this->pos;
    while(isalpha(this->curr_char)) {
        this->advance_char();
    }
    auto alpha = this->input.substr(start, (this->pos - start));
    return alpha;
}


void Lexer::skip_ws() {
    while(isspace(this->curr_char)) {
        this->advance_char();
    }
}











