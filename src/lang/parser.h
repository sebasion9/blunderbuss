#pragma once
#include "lexer.h"


class Parser {
    private:
        Lexer &lexer;
        Token currToken;


    public:
        Parser(Lexer &l);
        void advanceToken();
        void parseExpression();

};
