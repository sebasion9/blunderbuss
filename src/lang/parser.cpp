#include "parser.h"

Parser::Parser(Lexer &l) : lexer(l), currToken(l.nextToken()) {};

void Parser::advanceToken() {
    this->currToken = this->lexer.nextToken();
}

