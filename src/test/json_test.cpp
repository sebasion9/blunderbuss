#include <gtest/gtest.h>
#include "color.h"
#include "../lang/parser.h"

TEST(json_serialize, literal) {
    auto input = 5;
    LOG_INPUT(input);
    auto literal = new Literal(5);
    std::string want = "{\"type\":\"literal\",\"value\":5}";

    EXPECT_EQ(want, literal->to_json().dump());
}
TEST(json_serialize, for_stmt) {
    auto input = "for let i = 0; i<1; i = i + 1; {\n\tif var == 111 {\n\t\tfoo = 1;\n\t} else {\n\t\tbar = 2;\n\t};\n};\n";
    LOG_INPUT(input);

    auto lexer = Lexer(input);
    auto parser = Parser(lexer);
    auto got = parser.parse_stmt()->to_json().dump();

}


int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
