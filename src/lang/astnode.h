#pragma once
#include "tokens.h"
#include <memory>
#include <variant>
// abstract syntax tree node
struct ASTNode {
    virtual ~ASTNode() = default;
};


struct Expression : public ASTNode {};

struct Literal : public Expression {
    std::variant<int, double, std::string> value;
    Literal(std::variant<int, double, std::string> val) : value(val) {};
};

struct BinaryExpression : public Expression {
    std::unique_ptr<ASTNode> left;
    std::unique_ptr<ASTNode> right;
    TokenType op;
    BinaryExpression(std::unique_ptr<ASTNode> l, std::unique_ptr<ASTNode> r, TokenType op) : left(std::move(l)), right(std::move(r)), op(op) {};
};

struct AssignStatement : public ASTNode {
    std::string ident;
    std::unique_ptr<Expression> expr;
    AssignStatement(std::string ident, std::unique_ptr<Expression> expr) : ident(ident), expr(std::move(expr)) {};
};

struct ForStatement : public ASTNode {
    std::unique_ptr<AssignStatement> assign;
    std::unique_ptr<Expression> condition;
    std::unique_ptr<AssignStatement> endstmt;
    ForStatement(std::unique_ptr<ASTNode> assign, std::unique_ptr<Expression> condition, std::unique_ptr<ASTNode> endstmt) :
        assign(dynamic_cast<AssignStatement*>(assign.release())),
        condition(std::move(condition)),
        endstmt(dynamic_cast<AssignStatement*>(endstmt.release())) {};
};

struct IfStatement : public ASTNode {
    std::unique_ptr<Expression> expr;
    IfStatement(std::unique_ptr<Expression> expr) : expr(std::move(expr)) {};
};



