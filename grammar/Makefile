ANTLR_JAR := ../tools/antlr-4.13.1-complete.jar
JAVA := /opt/homebrew/opt/openjdk@21/bin/java
ANTLR := $(JAVA) -jar $(ANTLR_JAR)
OUTPUT_DIR := ../internal/parser
GO_PACKAGE := parser

.PHONY: generate-go clean test-grammar

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
