#include <nlohmann/json.hpp>
#include "../lang/astnode.h"

nlohmann::json node_vec_to_json(const std::vector<std::unique_ptr<AstNode>> &vec);
