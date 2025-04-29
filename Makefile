# Makefile for Project Euler Go Scripts

# Define the Go command
GO=go

# Define the scripts directory
SCRIPTS_DIR=problems

# List of problem scripts
PROBLEMS=$(wildcard $(SCRIPTS_DIR)/*.go)

# Default target
all: run

# Run all problem scripts
run: $(PROBLEMS)
	@for script in $(PROBLEMS); do \
		echo "Running $$script..."; \
		$(GO) run $$script; \
	done

# Clean up compiled files (if any)
clean:
	@echo "Cleaning up..."
	@rm -f $(SCRIPTS_DIR)/*/*.out

# Help command
help:
	@echo "Makefile commands:"
	@echo "  make run    - Run all problem scripts"
	@echo "  make clean  - Clean up compiled files"
	@echo "  make help   - Show this help message"