#include "../lang/astnode.h"
#include <nlohmann/json_fwd.hpp>
#include <variant>
#include <sstream>

njson::json node_vec_to_json(const std::vector<std::unique_ptr<AstNode>> &vec) {
    auto j = njson::json::array();
    for(auto &node : vec) {
        j.push_back(node->to_json());
    }
    return j;
}

njson::json Literal::to_json() const {
    auto type = "literal";
    auto val = std::visit([](auto &&v) {return njson::json(v);}, this->value);
    njson::json j;
    j["type"] = type;
    j["value"] = val;
    return j;
}

njson::json ForStatement::to_json() const {
    auto type = "for_stmt";
    auto assign = this->assign->to_json();
    auto condition = this->condition->to_json();
    auto endstmt = this->endstmt->to_json();
    auto block = node_vec_to_json(this->block);
    njson::json j;
    j["type"] = type;
    j["assign_stmt"] = assign;
    j["condition"] = condition;
    j["end_stmt"] = endstmt;
    j["block"] = block;
    return {};
}

njson::json AssignStatement::to_json() const {
    auto type = "assign_stmt";
    auto ident = this->ident;
    auto expr = this->expr->to_json();
    njson::json j;
    j["type"] = type;
    j["ident"] = ident;
    j["expr"] = expr;
    return j;
}

njson::json BinaryExpression::to_json() const {
    auto type = "bin_expr";
    auto left = this->left->to_json();
    auto right = this->left->to_json();
    std::stringstream opss;
    opss << this->op;
    std::string op = opss.str();

    njson::json j;
    j["type"] = type;
    j["left"] = left;
    j["right"] = right;
    j["operator"] = op;

    return j;
}

njson::json IfStatement::to_json() const {
    auto type = "if_stmt";
    auto expr = this->expr->to_json();
    auto then_block = node_vec_to_json(this->then_block);
    auto else_block = node_vec_to_json(this->else_block);

    njson::json j;
    j["type"] = type;
    j["expr"] = expr;
    j["then_block"] = then_block;
    j["else_blockj"] = else_block;
    return j;
}




