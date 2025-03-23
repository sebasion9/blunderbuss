#include "parser.h"
#include "astnode.h"
#include "tokens.h"
#include <memory>
#include <variant>

Parser::Parser(Lexer &l) : lexer(l), currToken(l.nextToken()) {};

void Parser::advanceToken() {
    this->currToken = this->lexer.nextToken();
}

std::unique_ptr<ASTNode> Parser::parseSingle() {
    if(this->currToken.type == TokenType::INT || this->currToken.type == TokenType::DOUBLE || this->currToken.type == TokenType::IDENT) {
        auto value = this->currToken.value;
        this->advanceToken();
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

std::unique_ptr<ASTNode> Parser::parseExpression() {
    auto left = this->parseSingle();
    auto currType = this->currToken.type;
    while(isExprOperator(currType)) {
        TokenType op = currType;
        this->advanceToken();
        auto right = parseSingle();
        currType = this->currToken.type;
        left = std::make_unique<BinaryExpression>(std::move(left), std::move(right), op);
    }
    return left;
}

std::unique_ptr<ASTNode> Parser::parseStatement() {
    bool parseAssign;
    bool parseFor;
    std::string ident;
    if(this->currToken.type == TokenType::KEYWORD) {
        auto keyword = std::get<std::string>(this->currToken.value);
        this->advanceToken();
        if(keyword == "let") {
            parseAssign = true;
        }
        if(keyword == "for") {
            parseFor = true;
        }
    } else if (this->currToken.type == TokenType::IDENT) {
        parseAssign = true;
    }
    if(parseAssign && this->currToken.type == TokenType::IDENT) {
        auto ident = std::get<std::string>(this->currToken.value);
        this->advanceToken();
        if(this->currToken.type == TokenType::ASSIGN) {
            this->advanceToken();
            auto expr = this->parseExpression();
            return std::make_unique<AssignStatement>(ident, std::move(expr));
        } else return nullptr;
    }
    if(parseFor) {
        auto stmt = parseStatement();
        auto* assign = dynamic_cast<AssignStatement*>(stmt.get());
        if(assign == nullptr) {
            return nullptr;
        }
        this->advanceToken();
        auto condition = parseExpression();
        this->advanceToken();
        auto endstmt = parseStatement();
        this->advanceToken();
        return std::make_unique<ForStatement>(std::move(stmt), std::move(condition), std::move(endstmt));
    }

    return nullptr;
}


