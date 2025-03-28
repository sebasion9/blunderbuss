#include "gtest/gtest.h"
#include <cstdio>
#include <gtest/gtest.h>
#include "../lang/parser.h"
#include "color.h"

TEST(Expr, NumExpr1) {
    auto input = "1 + 3.14";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseExpression();

    auto* binExpr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(binExpr, nullptr);
    EXPECT_EQ(binExpr->op, TokenType::PLUS);

    auto* left = dynamic_cast<Literal*>(binExpr->left.get());
    ASSERT_NE(left, nullptr);
    auto left_val = std::get<int>(left->value);
    EXPECT_EQ(left_val, 1);

    auto* right = dynamic_cast<Literal*>(binExpr->right.get());
    ASSERT_NE(right, nullptr);
    auto right_val = std::get<double>(right->value);
    EXPECT_EQ(right_val, 3.14);
}

TEST(Expr, NumExpr2) {
    auto input = "3 * 1 + 3.14";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseExpression();

    auto* binExpr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(binExpr, nullptr);
    EXPECT_EQ(binExpr->op, TokenType::PLUS);

    auto* left = dynamic_cast<BinaryExpression*>(binExpr->left.get());
    ASSERT_NE(left, nullptr);
    EXPECT_EQ(left->op, TokenType::MULT);
    auto* left1 = dynamic_cast<Literal*>(left->left.get());
    auto* right1 = dynamic_cast<Literal*>(left->right.get());
    ASSERT_NE(left1, nullptr);
    ASSERT_NE(right1, nullptr);
    auto left1_val = std::get<int>(left1->value);
    auto right1_val = std::get<int>(right1->value);

    EXPECT_EQ(left1_val, 3);
    EXPECT_EQ(right1_val, 1);

    auto* right = dynamic_cast<Literal*>(binExpr->right.get());
    ASSERT_NE(right, nullptr);
    auto right_val = std::get<double>(right->value);
    EXPECT_EQ(right_val, 3.14);
}

TEST(Expr, SingleVal) {
    auto input = "3;";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseExpression();
    ASSERT_NE(ast, nullptr);

    auto* expr = dynamic_cast<Literal*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(std::get<int>(expr->value), 3);

}

TEST(Expr, IdentExpr) {
    auto input = "num + 3";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseExpression();

    auto* expr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(expr->op, TokenType::PLUS);

    auto* right = dynamic_cast<Literal*>(expr->right.get());
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    EXPECT_EQ(std::get<int>(right->value), 3);
    EXPECT_EQ(std::get<std::string>(left->value), "num");
}
TEST(Expr, CompExpr) {
    auto input = "num == 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseExpression();


    auto* expr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(expr->op, TokenType::EQUALS);

    auto* right = dynamic_cast<Literal*>(expr->right.get());
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    EXPECT_EQ(std::get<int>(right->value), 2);
    EXPECT_EQ(std::get<std::string>(left->value), "num");
}
TEST(Stmt, Let) {
    auto input = "let a = 3 * 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();

    auto* stmt = dynamic_cast<AssignStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);
    EXPECT_EQ(stmt->ident, "a");

    auto* expr = dynamic_cast<BinaryExpression*>(stmt->expr.get());
    ASSERT_NE(expr, nullptr);
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    auto* right = dynamic_cast<Literal*>(expr->right.get());

    EXPECT_EQ(expr->op, TokenType::MULT);
    EXPECT_EQ(std::get<int>(left->value), 3);
    EXPECT_EQ(std::get<int>(right->value), 2);
}
TEST(Stmt, LetNoKeyword) {
    auto input = "a = 3 * 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();

    auto* stmt = dynamic_cast<AssignStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);
    EXPECT_EQ(stmt->ident, "a");

    auto* expr = dynamic_cast<BinaryExpression*>(stmt->expr.get());
    ASSERT_NE(expr, nullptr);
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    auto* right = dynamic_cast<Literal*>(expr->right.get());

    EXPECT_EQ(expr->op, TokenType::MULT);
    EXPECT_EQ(std::get<int>(left->value), 3);
    EXPECT_EQ(std::get<int>(right->value), 2);

}

TEST(Stmt, ForBlock) {
    auto input = "for let i = 0; i <= len; i = i + 1; { let a = 1; let b = 3; }";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();
    ASSERT_NE(ast, nullptr);
    auto* stmt = dynamic_cast<ForStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);

    auto* assign = dynamic_cast<AssignStatement*>(stmt->assign.get());
    ASSERT_NE(assign, nullptr);
    auto* con = dynamic_cast<BinaryExpression*>(stmt->condition.get());
    ASSERT_NE(con, nullptr);
    auto* endstmt = dynamic_cast<AssignStatement*>(stmt->endstmt.get());
    ASSERT_NE(endstmt, nullptr);


    auto block = std::move(stmt->block);
    EXPECT_EQ(block.size(), 2);

    auto* astmt1 = dynamic_cast<AssignStatement*>(block[0].get());
    auto* astmt2 = dynamic_cast<AssignStatement*>(block[1].get());

    ASSERT_NE(astmt1, nullptr);
    ASSERT_NE(astmt2, nullptr);

    auto ident1 = astmt1->ident;
    auto ident2 = astmt2->ident;
    EXPECT_EQ(ident1, "a");
    EXPECT_EQ(ident2, "b");

}

TEST(Stmt, IfSingleBlock) {
    auto input = "if 1 { let a = 3; let b = 3;} ";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();
    ASSERT_NE(ast, nullptr);

    auto* stmt = dynamic_cast<IfStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);

    auto* expr = dynamic_cast<Literal*>(stmt->expr.get());
    ASSERT_NE(expr, nullptr);
    auto val = std::get<int>(expr->value);
    EXPECT_EQ(val, 1);

    auto block = std::move(stmt->thenBlock);
    EXPECT_EQ(block.size(), 2);

    auto* astmt1 = dynamic_cast<AssignStatement*>(block[0].get());
    auto* astmt2 = dynamic_cast<AssignStatement*>(block[1].get());

    ASSERT_NE(astmt1, nullptr);
    ASSERT_NE(astmt2, nullptr);

    auto ident1 = astmt1->ident;
    auto ident2 = astmt2->ident;
    EXPECT_EQ(ident1, "a");
    EXPECT_EQ(ident2, "b");

}

TEST(Stmt, IfGreater) {
    auto input = "if a>5";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();
    ASSERT_NE(ast, nullptr);

    auto* stmt = dynamic_cast<IfStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);

    auto* expr = dynamic_cast<BinaryExpression*>(stmt->expr.get());
    ASSERT_NE(expr, nullptr);
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    auto* right= dynamic_cast<Literal*>(expr->right.get());
    auto op = expr->op;

    EXPECT_EQ(std::get<std::string>(left->value), "a");
    EXPECT_EQ(std::get<int>(right->value), 5);
    EXPECT_EQ(op, TokenType::GREATER_THAN);

}

TEST(Block, BlockAssign) {
    auto input = "let a = 5; let b = 3;";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto block = parser.parseBlock();
    ASSERT_EQ(block.size(), 2);
}
TEST(Block, ForIf) {
    auto input = "for let i = 0; i < 1; i = i + 1; { if 1 { let a = b; }; };";

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parseStatement();
    ASSERT_NE(ast, nullptr);


    auto* stmt = dynamic_cast<ForStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);

    auto* assign = dynamic_cast<AssignStatement*>(stmt->assign.get());
    ASSERT_NE(assign, nullptr);
    auto* con = dynamic_cast<BinaryExpression*>(stmt->condition.get());
    ASSERT_NE(con, nullptr);
    auto* endstmt = dynamic_cast<AssignStatement*>(stmt->endstmt.get());
    ASSERT_NE(endstmt, nullptr);

    auto block = std::move(stmt->block);
    ASSERT_EQ(block.size(), 1);
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
