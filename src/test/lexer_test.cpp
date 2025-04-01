#include <gtest/gtest.h>
#include "../lang/lexer.h"
#include "color.h"

TEST(lexer_test, tokenize_arithmetic) {
    auto input = "2 + 5;";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 2));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 5));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::SEMI));
}

TEST(lexer_test, tokenize_identifier) {
    auto input = "abc = 3 + 1;";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "abc"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 3));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 1));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::SEMI));
}

TEST(lexer_test, tokenize_comparison) {
    auto input = "identifier = 1 ==2;";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "identifier"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 1));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::EQUALS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::INT, 2));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::SEMI));
}

TEST(lexer_test, tokenize_symbol) {
    auto input = "\"abcd\"([{}])";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::QUOTE));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "abcd"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::QUOTE));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::LPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::LSPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::LCPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::RCPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::RSPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::RPAREN));
}

TEST(lexer_test, tokenize_keyword) {
    auto input = "for each franzl if return;return";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::KEYWORD, "for"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "each"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "franzl"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::KEYWORD, "if"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::KEYWORD, "return"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::SEMI));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::KEYWORD, "return"));
}

TEST(lexer_test, tokenize_eof) {
    auto input = "\0";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::END_OF_FILE));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::END_OF_FILE));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::END_OF_FILE));
}

TEST(lexer_test, tokenize_illegal) {
    auto input = "#;";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::ILLEGAL, "#"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::SEMI));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::END_OF_FILE));
}

TEST(lexer_test, tokenize_Double) {
    auto input = "someident = 3.14 + 1. - 2.0 + 0.1111;";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "someident"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::DOUBLE, 3.14));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::DOUBLE, 1.));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::MINUS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::DOUBLE, 2.0));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::DOUBLE, 0.1111));
}

TEST(lexer_test, condtions) {
    auto input = "(a > b & a <= c | !a == b)";
    LOG_INPUT(input);
    auto lexer = Lexer(input);
    EXPECT_EQ(lexer.next_token(), Token(TokenType::LPAREN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "a"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::GREATER_THAN));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "b"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::AND));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "a"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::LESS_THAN_EQUAL));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "c"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::OR));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::NOT));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "a"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::EQUALS));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::IDENT, "b"));
    EXPECT_EQ(lexer.next_token(), Token(TokenType::RPAREN));

}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
