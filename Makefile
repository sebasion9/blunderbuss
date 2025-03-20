CXX = g++
CXXFLAGS = -std=c++20 -Wall -Wextra -Isrc
DEBUGFLAGS = -g -O0
LDFLAGS = -lgtest -lgtest_main -pthread

SRC_DIR = src
BUILD_DIR = build
TEST_DIR = $(SRC_DIR)/test


SRCS = $(shell find $(SRC_DIR) -type f -name "*.cpp" ! -path "$(TEST_DIR)/*")
SRCS_NO_MAIN = $(filter-out $(SRC_DIR)/main.cpp, $(SRCS))


OBJS = $(patsubst $(SRC_DIR)/%.cpp, $(BUILD_DIR)/%.o, $(SRCS))


BUILD_DIRS = $(sort $(dir $(patsubst $(SRC_DIR)/%, $(BUILD_DIR)/%, $(SRCS))))

TESTS = $(wildcard $(TEST_DIR)/*.cpp)
TEST_BINS = $(patsubst $(TEST_DIR)/%.cpp, $(BUILD_DIR)/%, $(TESTS))

TARGET = $(BUILD_DIR)/blunderbuss
DEBUG_TARGET = $(BUILD_DIR)/blunderbuss_debug

all: $(BUILD_DIRS) $(TARGET)

$(TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) $^ -o $@

$(DEBUG_TARGET): CXXFLAGS += $(DEBUGFLAGS)
$(DEBUG_TARGET): $(OBJS)
	$(CXX) $(CXXFLAGS) $^ -o $@

$(BUILD_DIR)/%.o: $(SRC_DIR)/%.cpp | $(BUILD_DIRS)
	$(CXX) $(CXXFLAGS) -c $< -o $@

$(BUILD_DIRS):
	mkdir -p $@

clean:
	rm -rf $(BUILD_DIR)

run: $(TARGET)
	./$(TARGET)

debug: $(DEBUG_TARGET)
	gdb ./$(DEBUG_TARGET)

test: $(TEST_BINS)
	@for test in $(TEST_BINS); do ./$$test || exit 1; done

$(BUILD_DIR)/%: $(TEST_DIR)/%.cpp $(SRCS_NO_MAIN) | $(BUILD_DIRS)
	$(CXX) $(CXXFLAGS) $^ -o $@ $(LDFLAGS)

