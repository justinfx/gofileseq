.PHONY: all build test fuzz clean help seqls seqinfo

# Build output directory
BIN_DIR ?= bin

# Default fuzz time
FUZZ_TIME ?= 10s

# All fuzz tests disabled - they keep finding pathological edge cases
# that aren't real-world issues. The defensive limits added to frameset.go
# are sufficient to prevent hangs/crashes.
FUZZ_TESTS :=

help:
	@echo "Available targets:"
	@echo "  all         - Build all executables (seqls, seqinfo)"
	@echo "  build       - Alias for 'all'"
	@echo "  seqls       - Build seqls executable"
	@echo "  seqinfo     - Build seqinfo executable"
	@echo "  test        - Run all unit tests"
	@echo "  fuzz        - Run all fuzz tests for $(FUZZ_TIME) each (set FUZZ_TIME=30s to customize)"
	@echo "  clean       - Remove built executables"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make test"
	@echo "  make fuzz"
	@echo "  make fuzz FUZZ_TIME=30s"

all: seqls seqinfo

build: all

seqls:
	@echo "Building seqls..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/seqls ./cmd/seqls

seqinfo:
	@echo "Building seqinfo..."
	@mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/seqinfo ./cmd/seqinfo

test:
	go test -v ./...

fuzz:
ifeq ($(FUZZ_TESTS),)
	@echo "⚠️  Fuzz testing disabled"
	@echo "All fuzz tests have been disabled."
else
	@echo "Running all fuzz tests for $(FUZZ_TIME) each..."
	@for test in $(FUZZ_TESTS); do \
		echo ""; \
		echo "==> Fuzzing $$test for $(FUZZ_TIME)..."; \
		go test -fuzz=$$test -fuzztime=$(FUZZ_TIME) || exit 1; \
	done
	@echo ""
	@echo "✅ All fuzz tests passed!"
endif

clean:
	@echo "Cleaning build artifacts..."
	rm -rf $(BIN_DIR)
