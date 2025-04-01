#pragma once
#include "astnode.h"
#include "lexer.h"


class Parser {
    private:
        Lexer &lexer;
        Token curr_token;


    public:
        Parser(Lexer &l);
        void advance_token();
        std::unique_ptr<Expression> parse_single();
        std::unique_ptr<Expression> parse_expr();
        std::unique_ptr<AstNode> parse_stmt();
        std::vector<std::unique_ptr<AstNode>> parse_block();
};
