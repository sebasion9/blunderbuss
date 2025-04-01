#include "parser.h"
#include "astnode.h"
#include "tokens.h"
#include <memory>
#include <variant>

Parser::Parser(Lexer &l) : lexer(l), curr_token(l.next_token()) {};

void Parser::advance_token() {
    this->curr_token = this->lexer.next_token();
}

std::unique_ptr<Expression> Parser::parse_single() {
    if(this->curr_token.type == TokenType::INT || this->curr_token.type == TokenType::DOUBLE || this->curr_token.type == TokenType::IDENT) {
        auto value = this->curr_token.value;
        this->advance_token();
        if(std::get_if<int>(&value)) {
            return std::make_unique<Literal>(std::get<int>(value));
        }
        if(std::get_if<std::string>(&value)) {
            return std::make_unique<Literal>(std::get<std::string>(value));
        }
        return std::make_unique<Literal>(std::get<double>(value));
    }
    return nullptr;
}

std::unique_ptr<Expression> Parser::parse_expr() {
    auto left = this->parse_single();
    auto currType = this->curr_token.type;
    while(is_expr_op(currType)) {
        TokenType op = currType;
        this->advance_token();
        auto right = parse_single();
        currType = this->curr_token.type;
        left = std::make_unique<BinaryExpression>(std::move(left), std::move(right), op);
    }
    return left;
}

std::unique_ptr<AstNode> Parser::parse_stmt() {
    bool parse_assign = false;
    bool parse_for = false;
    bool parse_if = false;
    std::string ident;
    if(this->curr_token.type == TokenType::KEYWORD) {
        auto keyword = std::get<std::string>(this->curr_token.value);
        this->advance_token();
        if(keyword == "let") {
            parse_assign = true;
        }
        if(keyword == "for") {
            parse_for = true;
        }
        if(keyword == "if") {
            parse_if = true;
        }
    } else if (this->curr_token.type == TokenType::IDENT) {
        parse_assign = true;
    }
    if(parse_assign && this->curr_token.type == TokenType::IDENT) {
        auto ident = std::get<std::string>(this->curr_token.value);
        this->advance_token();
        if(this->curr_token.type == TokenType::ASSIGN) {
            this->advance_token();
            auto expr = this->parse_expr();
            return std::make_unique<AssignStatement>(ident, std::move(expr));
        } else return nullptr;
    }
    if(parse_for) {
        auto stmt = parse_stmt();
        std::vector<std::unique_ptr<AstNode>> block;
        auto* assign = dynamic_cast<AssignStatement*>(stmt.get());
        if(assign == nullptr) {
            return nullptr;
        }
        this->advance_token();
        auto condition = parse_expr();
        this->advance_token();
        auto endstmt = parse_stmt();
        this->advance_token();

        if(this->curr_token.type == TokenType::LCPAREN) {
            // consume LCPAREN
            this->advance_token();
            block = parse_block();

            if(this->curr_token.type != TokenType::RCPAREN) {
                return nullptr;
            }
            // consume RCPAREN
            this->advance_token();

        }

        return std::make_unique<ForStatement>(std::move(stmt), std::move(condition), std::move(endstmt), std::move(block));
    }
    if(parse_if) {
        auto expr = parse_expr();
        std::vector<std::unique_ptr<AstNode>> then_block;
        std::vector<std::unique_ptr<AstNode>> else_block;
        if(expr == nullptr) {
            return nullptr;
        }

        if(this->curr_token.type == TokenType::LCPAREN) {
            // consume LCPAREN
            this->advance_token();
            then_block = parse_block();

            if(this->curr_token.type != TokenType::RCPAREN) {
                return nullptr;
            }
            // consume RCPAREN
            this->advance_token();

        } else {
        // single statement
            then_block.push_back(this->parse_stmt());
        }

        if(this->curr_token.type == TokenType::KEYWORD && std::get<std::string>(this->curr_token.value) == "else") {
            // consume keyword
            this->advance_token();
            if(this->curr_token.type == TokenType::LCPAREN) {
                // consume LCPAREN
                this->advance_token();
                else_block = parse_block();

                if(this->curr_token.type != TokenType::RCPAREN) {
                    return nullptr;
                }
                // consume RCPAREN
                this->advance_token();
            } else {
                else_block.push_back(this->parse_stmt());
            }
        }
        return std::make_unique<IfStatement>(std::move(expr), std::move(then_block), std::move(else_block));
    }
    return nullptr;
}

std::vector<std::unique_ptr<AstNode>> Parser::parse_block() {
    std::vector<std::unique_ptr<AstNode>>  block;
    while(auto stmt = parse_stmt()) {
        if(this->curr_token.type != TokenType::SEMI) {
            break;
        }
        this->advance_token();
        block.push_back(std::move(stmt));
    }
    return block;
}



