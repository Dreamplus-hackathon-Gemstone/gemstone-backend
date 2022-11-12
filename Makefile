.PHONY: gemstone clean all test build

BUILD_DIR = build

BUILD_TARGET = ./cmd

BUILD_OUTPUT = gemstoned

all: build

build:
	@echo "\033[92mBuild gemstone daemon...\033[0m"
	@go build -o $(BUILD_DIR)/$(BUILD_OUTPUT) $(BUILD_TARGET)

clean:
	@echo "\033[92mClean daemon files...\033[0m"
	@rm $(BUILD_DIR)/*

fclean:
	@echo "\033[92mForce clean daemon files...\033[0m"
	@rm -f $(BUILD_DIR)/*

re:
	@make fclean
	@make all