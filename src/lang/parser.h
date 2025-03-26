#pragma once
#include "astnode.h"
#include "lexer.h"


class Parser {
    private:
        Lexer &lexer;
        Token currToken;


    public:
        Parser(Lexer &l);
        void advanceToken();
        std::unique_ptr<Expression> parseSingle();
        std::unique_ptr<Expression> parseExpression();
        std::unique_ptr<ASTNode> parseStatement();
};
