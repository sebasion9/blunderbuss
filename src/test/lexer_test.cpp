#include <gtest/gtest.h>
#include "../lang/lexer.h"

TEST(LexerTest, TokenizeArithmetic) {
    auto lexer = Lexer("2 + 5;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "2"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::OPERATOR, "+"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "5"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI, ";"));
}

TEST(LexerTest, TokenizeIdentifier) {
    auto lexer = Lexer("abc = 3 + 1;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "abc"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "="));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "3"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::OPERATOR, "+"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "1"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI, ";"));
}

TEST(LexerTest, TokenizeComparison) {
    auto lexer = Lexer("identifier = 1 ==2;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "identifier"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "="));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "1"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "=="));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::INT, "2"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI, ";"));
}

TEST(LexerTest, TokenizeSymbol) {
    auto lexer = Lexer("\"abcd\"([{}])");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "\""));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "abcd"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "\""));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "("));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "["));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "{"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "}"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, "]"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SYMBOL, ")"));
}

TEST(LexerTest, TokenizeKeyword) {
    auto lexer = Lexer("for each franzl if return;return");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "for"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "each"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::IDENT, "franzl"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "if"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "return"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI, ";"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::KEYWORD, "return"));
}

TEST(LexerTest, TokenizeEOF) {
    auto lexer = Lexer("\0");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE, std::string{'\0'}));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE, std::string{'\0'}));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE, std::string{'\0'}));
}

TEST(LexerTest, TokenizeILLEGAL) {
    auto lexer = Lexer("#;");
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::ILLEGAL, "#"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::SEMI, ";"));
    EXPECT_EQ(lexer.nextToken(), Token(TokenType::END_OF_FILE, std::string{'\0'}));
}

int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
