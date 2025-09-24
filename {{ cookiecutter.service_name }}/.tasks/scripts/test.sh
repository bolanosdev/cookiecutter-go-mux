#!/bin/bash

API_EXCLUDE_DIRS=("entities")
API_PACKAGES=$(go list ./service/...)

INTERNAL_EXCLUDE_DIRS=("config" "consts" "mocks" "queries" "models")
INTERNAL_PACKAGES=$(go list ./internal/...)

for dir in "${API_EXCLUDE_DIRS[@]}"; do
  API_PACKAGES=$(echo "$API_PACKAGES" | grep -v "$dir")
done

for dir in "${INTERNAL_EXCLUDE_DIRS[@]}"; do
  INTERNAL_PACKAGES=$(echo "$INTERNAL_PACKAGES" | grep -v "$dir")
done

go test $API_PACKAGES $INTERNAL_PACKAGES -coverprofile .cover/cover.out
