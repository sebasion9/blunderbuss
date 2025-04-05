#include "gtest/gtest.h"
#include <cstdio>
#include <gtest/gtest.h>
#include "../lang/parser.h"
#include "color.h"

TEST(expr, num_expr1) {
    auto input = "1 + 3.14";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_expr();

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

TEST(expr, num_expr) {
    auto input = "3 * 1 + 3.14";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_expr();

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

TEST(expr, single_val) {
    auto input = "3;";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_expr();
    ASSERT_NE(ast, nullptr);

    auto* expr = dynamic_cast<Literal*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(std::get<int>(expr->value), 3);

}

TEST(expr, ident_expr) {
    auto input = "num + 3";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_expr();

    auto* expr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(expr->op, TokenType::PLUS);

    auto* right = dynamic_cast<Literal*>(expr->right.get());
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    EXPECT_EQ(std::get<int>(right->value), 3);
    EXPECT_EQ(std::get<std::string>(left->value), "num");
}
TEST(expr, comp_expr) {
    auto input = "num == 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_expr();


    auto* expr = dynamic_cast<BinaryExpression*>(ast.get());
    ASSERT_NE(expr, nullptr);
    EXPECT_EQ(expr->op, TokenType::EQUALS);

    auto* right = dynamic_cast<Literal*>(expr->right.get());
    auto* left = dynamic_cast<Literal*>(expr->left.get());
    EXPECT_EQ(std::get<int>(right->value), 2);
    EXPECT_EQ(std::get<std::string>(left->value), "num");
}
TEST(stmt, let) {
    auto input = "let a = 3 * 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();

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
TEST(stmt, let_no_keyword) {
    auto input = "a = 3 * 2";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();

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

TEST(stmt, for_block) {
    auto input = "for let i = 0; i <= len; i = i + 1; { let a = 1; let b = 3; }";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();
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

TEST(stmt, if_single_block) {
    auto input = "if 1 { let a = 3; let b = 3;} ";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();
    ASSERT_NE(ast, nullptr);

    auto* stmt = dynamic_cast<IfStatement*>(ast.get());
    ASSERT_NE(stmt, nullptr);

    auto* expr = dynamic_cast<Literal*>(stmt->expr.get());
    ASSERT_NE(expr, nullptr);
    auto val = std::get<int>(expr->value);
    EXPECT_EQ(val, 1);

    auto block = std::move(stmt->then_block);
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

TEST(stmt, if_greater) {
    auto input = "if a>5";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();
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

TEST(block, block_assign) {
    auto input = "let a = 5; let b = 3;";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto block = parser.parse_block();
    ASSERT_EQ(block.size(), 2);
}
TEST(block, for_if) {
    auto input = "for let i = 0; i < 1; i = i + 1; { if 1 { let a = b; }; };";

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto ast = parser.parse_stmt();
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
TEST(block, for_lf) {
    auto input = "let var = 132;\nvar = 111;\n\nfor let i = 0; i<1; i = i + 1; {\n\tif var == 111 {\n\t\tfoo = 1;\n\t} else {\n\t\tbar = 2;\n\t};\n};\n";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto stmts = parser.parse_block();
    ASSERT_EQ(stmts.size(), 3);

    for(size_t i = 0; i < stmts.size(); i++) {
        ASSERT_NE(stmts[i], nullptr);
    }
    auto* stmt1 = dynamic_cast<AssignStatement*>(stmts[0].get());
    auto* stmt2 = dynamic_cast<AssignStatement*>(stmts[1].get());
    auto* stmt3 = dynamic_cast<ForStatement*>(stmts[2].get());

    ASSERT_NE(stmt1, nullptr);
    ASSERT_NE(stmt2, nullptr);
    ASSERT_NE(stmt3, nullptr);


    // stmt1
    ASSERT_EQ(stmt1->ident, "var");
    ASSERT_EQ(std::get<int>(dynamic_cast<Literal*>(stmt1->expr.get())->value), 132);

    // stmt2
    ASSERT_EQ(stmt2->ident, "var");
    ASSERT_EQ(std::get<int>(dynamic_cast<Literal*>(stmt2->expr.get())->value), 111);

    // stmt3
    auto* assign = stmt3->assign.get();
    auto* condition = dynamic_cast<BinaryExpression*>(stmt3->condition.get());
    auto* endstmt = stmt3->endstmt.get();

    ASSERT_NE(assign, nullptr);
    ASSERT_NE(condition, nullptr);
    ASSERT_NE(endstmt, nullptr);

}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
