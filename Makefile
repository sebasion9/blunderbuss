CXX = g++
CXXFLAGS = -std=c++20 -Wall -Wextra -Iinclude
DEBUGFLAGS = -g -O0
LDFLAGS = -lgtest -lgtest_main -pthread

SRC_DIR = src
BUILD_DIR = build
TEST_DIR = src/test

SRCS = $(wildcard $(SRC_DIR)/*.cpp)
SRCS_NO_MAIN = $(filter-out $(SRC_DIR)/main.cpp, $(SRCS))
OBJS = $(patsubst $(SRC_DIR)/%.cpp, $(BUILD_DIR)/%.o, $(SRCS))

TESTS = $(wildcard $(TEST_DIR)/*.cpp)
TEST_BINS = $(patsubst $(TEST_DIR)/%.cpp, $(BUILD_DIR)/%, $(TESTS))

TARGET = $(BUILD_DIR)/blunderbuss
DEBUG_TARGET = $(BUILD_DIR)/blunderbuss_debug

all: $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) $^ -o $@

$(DEBUG_TARGET): CXXFLAGS += $(DEBUGFLAGS)
$(DEBUG_TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) $^ -o $@

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp | $(BUILD_DIR)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIR):
	mkdir -p $(BUILD_DIR)

clean:
	rm -rf $(BUILD_DIR)

run: $(TARGET)
	./$(TARGET)

debug: $(DEBUG_TARGET)
	gdb ./$(DEBUG_TARGET)

test: $(TEST_BINS)
	@for test in $(TEST_BINS); do ./$$test || exit 1; done

$(BUILD_DIR)/%: $(TEST_DIR)/%.cpp $(SRCS_NO_MAIN) | $(BUILD_DIR)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS)

