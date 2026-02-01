# Configuration
ANTLR_JAR := ../tools/antlr-4.13.1-complete.jar
OUTPUT_DIR := ../internal/parser
GO_PACKAGE := parser

# Java runtime detection (override with: make JAVA=/path/to/java)
# Tries in order: env var, homebrew (macOS), system java, PATH
JAVA ?= $(shell command -v /opt/homebrew/opt/openjdk@21/bin/java 2>/dev/null || \
                command -v /usr/bin/java 2>/dev/null || \
                command -v java 2>/dev/null || \
                echo "java")

ANTLR := $(JAVA) -jar $(ANTLR_JAR)

.PHONY: generate-go clean test-grammar check-java

generate-go:
	@echo "Generating Go parser from ANTLR grammar..."
	$(ANTLR) -Dlanguage=Go \
	         -package $(GO_PACKAGE) \
	         -o $(OUTPUT_DIR) \
	         -visitor \
	         -no-listener \
	         fileseq.g4
	@echo "Parser generated successfully in $(OUTPUT_DIR)"

clean:
	@echo "Cleaning generated parser files..."
	rm -rf $(OUTPUT_DIR)/*
	@echo "Clean complete"

test-grammar:
	@echo "Testing grammar compilation..."
	$(ANTLR) fileseq.g4
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
