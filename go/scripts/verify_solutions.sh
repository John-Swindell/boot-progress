#!/usr/bin/env bash
# verify_solutions.sh — for each chapter, replace exercises.go with solutions/exercises.go
# (rewriting the package line to match the chapter package), run `go test`, restore.
#
# Used to confirm that the reference solutions actually pass the chapter tests.
# Run from the repo root.

set -u

cd "$(dirname "$0")/.."

GREEN=$'\e[32m'
RED=$'\e[31m'
RESET=$'\e[0m'

pass=0
fail=0
fail_chapters=()

for chapter in [0-9][0-9]-*/; do
    chapter="${chapter%/}"
    [ -f "${chapter}/exercises.go" ] || continue
    [ -f "${chapter}/solutions/exercises.go" ] || continue

    # detect chapter package name from the existing exercises.go
    pkg=$(awk '/^package /{print $2; exit}' "${chapter}/exercises.go")
    [ -n "$pkg" ] || { echo "no package in ${chapter}/exercises.go"; continue; }

    # backup, swap in solution rewritten to the chapter package
    cp "${chapter}/exercises.go" "${chapter}/exercises.go.bak"
    sed "s/^package solutions/package ${pkg}/" "${chapter}/solutions/exercises.go" > "${chapter}/exercises.go"

    if go test -timeout 30s "./${chapter}" >/dev/null 2>&1; then
        printf "[${GREEN}✓${RESET}] %s\n" "$chapter"
        pass=$((pass+1))
    else
        printf "[${RED}✗${RESET}] %s\n" "$chapter"
        fail=$((fail+1))
        fail_chapters+=("$chapter")
    fi

    # restore
    mv "${chapter}/exercises.go.bak" "${chapter}/exercises.go"
done

echo
echo "Passed: ${pass}, Failed: ${fail}"
if [ "$fail" -ne 0 ]; then
    echo "Failing chapters:"
    for c in "${fail_chapters[@]}"; do echo "  - $c"; done
    exit 1
fi
