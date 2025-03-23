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

    auto* expr = stmt->expr.get();
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

    auto* expr = stmt->expr.get();
    ASSERT_NE(expr, nullptr);
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    auto* right = dynamic_cast<Literal*>(expr->right.get());

    EXPECT_EQ(expr->op, TokenType::MULT);
    EXPECT_EQ(std::get<int>(left->value), 3);
    EXPECT_EQ(std::get<int>(right->value), 2);

}


int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
