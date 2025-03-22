#pragma once
#include "tokens.h"
#include <memory>
#include <variant>
struct ASTNode {
    virtual ~ASTNode() = default;
};


struct NumberLiteral : public ASTNode {
    std::variant<int, double> value;
    NumberLiteral(std::variant<int, double> val) : value(val) {};
};

struct BinaryExpression : public ASTNode {
    std::unique_ptr<ASTNode> left;
    std::unique_ptr<ASTNode> right;
    TokenType op;
    BinaryExpression(std::unique_ptr<ASTNode> l, std::unique_ptr<ASTNode> r, TokenType op) : left(std::move(l)), right(std::move(r)), op(op) {};
};



