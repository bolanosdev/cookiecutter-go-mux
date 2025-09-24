#!/bin/bash

# Desired threshold
THRESHOLD=65.0

# Run tests and write coverage to file
task test:unit

# Extract the total coverage percentage
COVERAGE=$(go tool cover -func=.cover/cover.out | grep total | awk '{print substr($3, 1, length($3)-1)}')

# Compare coverage
if (($(echo "$COVERAGE < $THRESHOLD" | bc -l))); then
  echo "❌ Coverage $COVERAGE% is below threshold $THRESHOLD%"
  exit 1
else
  echo "✅ Coverage $COVERAGE% is above threshold $THRESHOLD%"
  exit 0
fi
