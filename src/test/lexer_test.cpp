#include <gtest/gtest.h>
#include "../lang/lexer.h"

TEST(LexerTest, TokenizeArithmetic) {
    auto lexer = Lexer("2 + 5;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 2));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 5));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI));
}

TEST(LexerTest, TokenizeIdentifier) {
    auto lexer = Lexer("abc = 3 + 1;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "abc"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 3));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 1));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI));
}

TEST(LexerTest, TokenizeComparison) {
    auto lexer = Lexer("identifier = 1 ==2;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "identifier"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 1));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::EQUALS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, 2));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI));
}

TEST(LexerTest, TokenizeSymbol) {
    auto lexer = Lexer("\"abcd\"([{}])");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::QUOTE));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "abcd"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::QUOTE));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::LPAREN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::LSPAREN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::LCPAREN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::RCPAREN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::RSPAREN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::RPAREN));
}

TEST(LexerTest, TokenizeKeyword) {
    auto lexer = Lexer("for each franzl if return;return");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "for"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "each"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "franzl"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "if"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "return"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "return"));
}

TEST(LexerTest, TokenizeEOF) {
    auto lexer = Lexer("\0");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE));
}

TEST(LexerTest, TokenizeILLEGAL) {
    auto lexer = Lexer("#;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::ILLEGAL, "#"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE));
}

TEST(LexerTest, TokenizeDouble) {
    auto lexer = Lexer("someident = 3.14 + 1. - 2.0 + 0.1111;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "someident"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::ASSIGN));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::DOUBLE, 3.14));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::DOUBLE, 1.));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::MINUS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::DOUBLE, 2.0));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::PLUS));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::DOUBLE, 0.1111));
}


int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
