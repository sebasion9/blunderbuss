#pragma once
#include <iostream>
#include <string>
#define ANSI_TXT_GRN "\033[0;32m"
#define ANSI_TXT_DFT "\033[0;0m"

#define LOG_INPUT(i) std::cout << ANSI_TXT_GRN << "[   INFO   ]" << ANSI_TXT_DFT << " INPUT: \"" << i << "\"\n"
#define LOG_MSG(msg) std::cout << ANSI_TXT_GRN << "[   INFO   ]" << ANSI_TXT_DFT << msg << "\n";
