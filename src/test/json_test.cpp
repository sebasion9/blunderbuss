#include <gtest/gtest.h>
#include "color.h"
#include "../lang/astnode.h"

TEST(json_serialize, literal) {
    auto input = 5;
    LOG_INPUT(input);
    auto literal = new Literal(5);
    std::string want = "{\"type\":\"Literal\",\"value\":5}";

    EXPECT_EQ(want, literal->to_json().dump());
}


int main(int argc, char **argv) {
    ::testing::InitGoogleTest(&argc, argv);
    return RUN_ALL_TESTS();
}
