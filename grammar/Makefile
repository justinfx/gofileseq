# Configuration
# Different ANTLR versions are required for Go vs C++:
#   - Go 4.13.1: Generates code with antlr4-go/antlr import paths (matches v4.13.0 runtime)
#   - C++ 4.9.3:  Generates C++11-compatible code (no std::any) for GCC 4.8+ support
# galop (https://github.com/martianoff/galop) then generates, next to ANTLR's Go
# parser, the faster parser ParseFileSequence uses: it builds the same contexts as
# ANTLR's (its "antlr" flavor). It is a native binary, downloaded into ../tools on
# first use at the pinned version below and checked against the release's checksums.
GALOP_VERSION := 0.5.1
GALOP ?= ../tools/galop-$(GALOP_VERSION)
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

generate-go: $(GALOP)
	@echo "Generating Go parser (ANTLR 4.13.1)..."
	$(ANTLR_GO) -Dlanguage=Go \
	         -package $(GO_PACKAGE) \
	         -o $(OUTPUT_DIR) \
	         -visitor \
	         -no-listener \
	         fileseq.g4
	@echo "Generating the galop parser (galop $(GALOP_VERSION))..."
	$(GALOP) -Dflavor=antlr \
	         -package $(GO_PACKAGE) \
	         -Xexact-output-dir \
	         -o $(OUTPUT_DIR) \
	         fileseq.g4
	@echo "Parser generated successfully in $(OUTPUT_DIR)"

generate-go-docker: $(GALOP)
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
	@echo "Generating the galop parser (galop $(GALOP_VERSION))..."
	$(GALOP) -Dflavor=antlr \
	         -package $(GO_PACKAGE) \
	         -Xexact-output-dir \
	         -o $(OUTPUT_DIR) \
	         fileseq.g4
	@echo "Parser generated successfully in $(OUTPUT_DIR)"

# Downloads the galop release binary for this OS and architecture and checks it
# against the release's SHA-256 checksums.
../tools/galop-$(GALOP_VERSION):
	@os=$$(uname -s | tr '[:upper:]' '[:lower:]'); arch=$$(uname -m); \
	case "$$arch" in x86_64|amd64) arch=amd64 ;; aarch64|arm64) arch=arm64 ;; esac; \
	case "$$os" in mingw*|msys*|cygwin*) os=windows; ext=.exe ;; *) ext= ;; esac; \
	name=galop-$$os-$$arch$$ext; \
	url=https://github.com/martianoff/galop/releases/download/$(GALOP_VERSION); \
	echo "Downloading $$name $(GALOP_VERSION)..."; \
	curl -fsSL -o $@.tmp "$$url/$$name" && \
	sum=$$(curl -fsSL "$$url/checksums-sha256.txt" | grep " \*\?$$name$$" | cut -d' ' -f1) && \
	got=$$( (sha256sum $@.tmp 2>/dev/null || shasum -a 256 $@.tmp) | cut -d' ' -f1) && \
	if [ -z "$$sum" ] || [ "$$sum" != "$$got" ]; then echo "checksum mismatch for $$name"; rm -f $@.tmp; exit 1; fi && \
	chmod +x $@.tmp && mv $@.tmp $@

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
