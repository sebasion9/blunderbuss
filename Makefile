CXX = g++
CXXFLAGS = -std=c++20 -Wall -Wextra -Iinclude
DEBUGFLAGS = -g -O0

SRC_DIR = src
BUILD_DIR = build

SRCS = $(wildcard $(SRC_DIR)/*.cpp)
OBJS = $(patsubst $(SRC_DIR)/%.cpp, $(BUILD_DIR)/%.o, $(SRCS))

TARGET = blunderbuss
DEBUG_TARGET = blunderbuss_debug

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
	rm -rf $(BUILD_DIR) $(TARGET) $(DEBUG_TARGET)

run: $(TARGET)
	./$(TARGET)

debug: $(DEBUG_TARGET)
	gdb ./$(DEBUG_TARGET)
