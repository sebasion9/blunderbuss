#pragma once
#include "tokens.h"
#include <memory>
#include <variant>
// abstract syntax tree node
struct ASTNode {
    virtual ~ASTNode() = default;
};


struct Literal : public ASTNode {
    std::variant<int, double, std::string> value;
    Literal(std::variant<int, double, std::string> val) : value(val) {};
};

struct BinaryExpression : public ASTNode {
    std::unique_ptr<ASTNode> left;
    std::unique_ptr<ASTNode> right;
    TokenType op;
    BinaryExpression(std::unique_ptr<ASTNode> l, std::unique_ptr<ASTNode> r, TokenType op) : left(std::move(l)), right(std::move(r)), op(op) {};
};

struct AssignStatement : public ASTNode {
    std::string ident;
    std::unique_ptr<BinaryExpression> expr;
    AssignStatement(std::string ident, std::unique_ptr<ASTNode> expr) : ident(ident), expr(dynamic_cast<BinaryExpression*>(expr.release())) {};
};

struct ForStatement : public ASTNode {
    std::unique_ptr<AssignStatement> assign;
    std::unique_ptr<BinaryExpression> condition;
    std::unique_ptr<AssignStatement> endstmt;
    ForStatement(std::unique_ptr<ASTNode> assign, std::unique_ptr<ASTNode> condition, std::unique_ptr<ASTNode> endstmt) :
        assign(dynamic_cast<AssignStatement*>(assign.release())),
        condition(dynamic_cast<BinaryExpression*>(condition.release())),
        endstmt(dynamic_cast<AssignStatement*>(endstmt.release())) {};
};



