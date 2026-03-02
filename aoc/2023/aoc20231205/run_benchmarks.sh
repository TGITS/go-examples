#!/usr/bin/env bash
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR" || exit 1

SKIP_BRUTEFORCE=0
if [[ "${1:-}" == "--skip-bruteforce" ]]; then
  SKIP_BRUTEFORCE=1
fi

declare -a RESULT_NAMES=()
declare -a RESULT_CODES=()
declare -a RESULT_METRICS=()

run_go_benchmark() {
  local name="$1"
  local pattern="$2"

  echo
  echo "=== ${name} ==="

  local output
  output=$(go test ./almanac -bench "$pattern" -benchmem -benchtime=1x -run '^$' 2>&1)
  local exit_code=$?

  printf '%s\n' "$output"

  local metrics_line
  metrics_line=$(printf '%s\n' "$output" | grep -E '^Benchmark[^[:space:]]+' | tail -n 1 || true)

  local metrics=""
  if [[ -n "$metrics_line" ]]; then
    if [[ "$metrics_line" =~ ^(Benchmark[^[:space:]]+)[[:space:]]+[0-9]+[[:space:]]+([0-9.]+)[[:space:]]+ns/op[[:space:]]+([0-9.]+)[[:space:]]+B/op[[:space:]]+([0-9.]+)[[:space:]]+allocs/op ]]; then
      metrics="${BASH_REMATCH[1]}|${BASH_REMATCH[2]}|${BASH_REMATCH[3]}|${BASH_REMATCH[4]}"
    fi
  fi

  RESULT_NAMES+=("$name")
  RESULT_CODES+=("$exit_code")
  RESULT_METRICS+=("$metrics")
}

run_go_benchmark "Optimized (full input)" "BenchmarkPart2Optimized_Input$"
if [[ "$SKIP_BRUTEFORCE" -eq 0 ]]; then
  run_go_benchmark "Brute force (full input)" "BenchmarkPart2BruteForce_Input$"
fi
run_go_benchmark "Optimized vs BruteForce (test input)" "BenchmarkPart2(Optimized|BruteForce)_InputTest$"

echo
echo "=== Summary ==="
if [[ "$SKIP_BRUTEFORCE" -eq 1 ]]; then
  echo "Brute force (full input): skipped by --skip-bruteforce"
fi
for i in "${!RESULT_NAMES[@]}"; do
  name="${RESULT_NAMES[$i]}"
  code="${RESULT_CODES[$i]}"
  metrics="${RESULT_METRICS[$i]}"

  if [[ "$code" -eq 0 ]]; then
    if [[ -n "$metrics" ]]; then
      IFS='|' read -r bench ns bytes allocs <<< "$metrics"
      echo "${name}: ${ns} ns/op, ${bytes} B/op, ${allocs} allocs/op"
    else
      echo "${name}: completed, metrics not parsed"
    fi
  else
    echo "${name}: failed (exit code ${code})"
  fi
done
