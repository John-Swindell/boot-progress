#!/usr/bin/env bash
# progress.sh — chapter-by-chapter pass/fail scoreboard for the Go Lab.
#
# Walks every NN-* directory at the repo root, runs `go test ./NN-*` for each,
# and prints a green/red/blank checklist. "Cleared" = tests pass.

set -u

cd "$(dirname "$0")/.."

GREEN=$'\e[32m'
RED=$'\e[31m'
DIM=$'\e[2m'
RESET=$'\e[0m'

cleared=0
total=0

echo
echo "GOLAB PROGRESS"
echo "=============="

for chapter in [0-9][0-9]-*/; do
    chapter="${chapter%/}"
    total=$((total + 1))

    # Does this chapter have any tests yet (recursively)?
    if [ -z "$(find "$chapter" -name '*_test.go' -print -quit 2>/dev/null)" ]; then
        printf "[ ] %s ${DIM}(no tests yet)${RESET}\n" "$chapter"
        continue
    fi

    # Run tests with a short timeout (unimplemented exercises with nil channels
    # would otherwise hang). Swallow output, just want the exit code.
    if go test -timeout 30s "./${chapter}/..." >/dev/null 2>&1; then
        printf "[${GREEN}✓${RESET}] %s\n" "$chapter"
        cleared=$((cleared + 1))
    else
        printf "[${RED}✗${RESET}] %s\n" "$chapter"
    fi
done

echo
echo "Cleared: ${cleared}/${total}"
echo
