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
        std::unique_ptr<ASTNode> parseSingle();
        std::unique_ptr<ASTNode> parseExpression();
        std::unique_ptr<ASTNode> parseStatement();
};
