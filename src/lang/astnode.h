#pragma once
#include "tokens.h"
#include <memory>
#include <variant>
#include <vector>
#include <nlohmann/json.hpp>

namespace njson = nlohmann;

// abstract syntax tree node
struct AstNode {
    virtual ~AstNode () = default;
    virtual njson::json to_json() const = 0;
};


struct Expression : public AstNode {};

struct Literal : public Expression {
    std::variant<int, double, std::string> value;
    Literal(std::variant<int, double, std::string> val) : value(val) {};

    njson::json to_json() const;
};

struct BinaryExpression : public Expression {
    std::unique_ptr<AstNode> left;
    std::unique_ptr<AstNode> right;
    TokenType op;
    BinaryExpression(std::unique_ptr<AstNode> l, std::unique_ptr<AstNode> r, TokenType op) : left(std::move(l)), right(std::move(r)), op(op) {};

    njson::json to_json() const;
};

struct AssignStatement : public AstNode {
    std::string ident;
    std::unique_ptr<Expression> expr;
    AssignStatement(std::string ident, std::unique_ptr<Expression> expr) : ident(ident), expr(std::move(expr)) {};

    njson::json to_json() const;
};

struct ForStatement : public AstNode {
    std::unique_ptr<AssignStatement> assign;
    std::unique_ptr<Expression> condition;
    std::unique_ptr<AssignStatement> endstmt;
    std::vector<std::unique_ptr<AstNode>> block;

    ForStatement(std::unique_ptr<AstNode> assign, std::unique_ptr<Expression> condition, std::unique_ptr<AstNode> endstmt, std::vector<std::unique_ptr<AstNode>> block) :
        assign(dynamic_cast<AssignStatement*>(assign.release())),
        condition(std::move(condition)),
        endstmt(dynamic_cast<AssignStatement*>(endstmt.release())),
        block(std::move(block)){};
    njson::json to_json() const;

};

struct IfStatement : public AstNode {
    std::unique_ptr<Expression> expr;
    std::vector<std::unique_ptr<AstNode>> then_block;
    std::vector<std::unique_ptr<AstNode>> else_block;
    IfStatement(std::unique_ptr<Expression> expr, std::vector<std::unique_ptr<AstNode>> then_block, std::vector<std::unique_ptr<AstNode>> else_block)
        : expr(std::move(expr)), then_block(std::move(then_block)), else_block(std::move(else_block)) {};

    njson::json to_json() const;
};



