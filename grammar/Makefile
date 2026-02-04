# Configuration
# Different ANTLR versions are required for Go vs C++:
#   - Go 4.13.1: Generates code with antlr4-go/antlr import paths (matches v4.13.0 runtime)
#   - C++ 4.9.3:  Generates C++11-compatible code (no std::any) for GCC 4.8+ support
ANTLR_GO_JAR := ../tools/antlr-4.13.1-complete.jar
ANTLR_CPP_JAR := ../tools/antlr-4.9.3-complete.jar
OUTPUT_DIR := ../internal/parser
GO_PACKAGE := parser
CPP_OUTPUT_DIR := ../cpp/private/parser
CPP_NAMESPACE := fileseq

# Java runtime detection (override with: make JAVA=/path/to/java)
# Tries in order: env var, homebrew (macOS), system java, PATH
JAVA ?= $(shell command -v /opt/homebrew/opt/openjdk@21/bin/java 2>/dev/null || \
                command -v /usr/bin/java 2>/dev/null || \
                command -v java 2>/dev/null || \
                echo "java")

ANTLR_GO := $(JAVA) -jar $(ANTLR_GO_JAR)
ANTLR_CPP := $(JAVA) -jar $(ANTLR_CPP_JAR)

# Docker configuration (for Java-free generation)
JAVA_IMAGE := eclipse-temurin:21-jdk
WORKDIR := /work

.PHONY: all generate-go generate-go-docker generate-cpp generate-cpp-docker generate-all clean test-grammar check-java check-docker

all: generate-all

generate-go:
	@echo "Generating Go parser (ANTLR 4.13.1)..."
	$(ANTLR_GO) -Dlanguage=Go \
	         -package $(GO_PACKAGE) \
	         -o $(OUTPUT_DIR) \
	         -visitor \
	         -no-listener \
	         fileseq.g4
	@echo "Parser generated successfully in $(OUTPUT_DIR)"

generate-go-docker:
	@echo "Generating Go parser using Docker (ANTLR 4.13.1)..."
	@docker run --rm --platform linux/amd64 \
	  -v $(realpath ../tools):/tools \
	  -v $(PWD):$(WORKDIR) \
	  -w $(WORKDIR) \
	  $(JAVA_IMAGE) \
	  java -jar /tools/$(notdir $(ANTLR_GO_JAR)) \
	    -Dlanguage=Go \
	    -package $(GO_PACKAGE) \
	    -o $(OUTPUT_DIR) \
	    -visitor \
	    -no-listener \
	    fileseq.g4
	@echo "Parser generated successfully in $(OUTPUT_DIR)"

generate-cpp:
	@echo "Generating C++ parser (ANTLR 4.9.3)..."
	$(ANTLR_CPP) -Dlanguage=Cpp \
	         -package $(CPP_NAMESPACE) \
	         -o $(CPP_OUTPUT_DIR) \
	         -visitor \
	         -no-listener \
	         fileseq.g4
	@echo "C++ parser generated successfully in $(CPP_OUTPUT_DIR)"

generate-cpp-docker:
	@echo "Generating C++ parser using Docker (ANTLR 4.9.3)..."
	@docker run --rm --platform linux/amd64 \
	  -v $(realpath ../tools):/tools \
	  -v $(PWD):$(WORKDIR) \
	  -w $(WORKDIR) \
	  $(JAVA_IMAGE) \
	  java -jar /tools/$(notdir $(ANTLR_CPP_JAR)) \
	    -Dlanguage=Cpp \
	    -package $(CPP_NAMESPACE) \
	    -o $(CPP_OUTPUT_DIR) \
	    -visitor \
	    -no-listener \
	    fileseq.g4
	@echo "C++ parser generated successfully in $(CPP_OUTPUT_DIR)"

generate-all: generate-go generate-cpp
	@echo "Generated parsers for all languages"

clean:
	@echo "Cleaning generated parser files..."
	rm -rf $(OUTPUT_DIR)/*
	rm -rf $(CPP_OUTPUT_DIR)/*
	@echo "Clean complete"

test-grammar:
	@echo "Testing grammar compilation (Go version)..."
	$(ANTLR_GO) fileseq.g4
	@echo "Grammar test complete"

check-java:
	@echo "Java runtime detection:"
	@echo "  JAVA = $(JAVA)"
	@if [ "$(JAVA)" = "java" ]; then \
		echo "  Source: PATH (system default)"; \
	elif echo "$(JAVA)" | grep -q homebrew; then \
		echo "  Source: Homebrew (macOS)"; \
	elif echo "$(JAVA)" | grep -q "/usr/bin"; then \
		echo "  Source: System installation"; \
	else \
		echo "  Source: Custom/Environment"; \
	fi
	@echo ""
	@$(JAVA) -version 2>&1 | head -3 || echo "ERROR: Java not found!"
	@echo ""
	@echo "Override with: make JAVA=/path/to/java generate-go"

check-docker:
	@echo "Docker detection:"
	@if command -v docker >/dev/null 2>&1; then \
		echo "  Docker: installed"; \
		docker --version; \
		echo ""; \
		echo "Use: make generate-go-docker"; \
	else \
		echo "  Docker: not found"; \
		echo "  Install from: https://www.docker.com/"; \
	fi
