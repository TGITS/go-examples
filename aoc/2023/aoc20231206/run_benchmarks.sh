#!/usr/bin/env bash
set -u

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR" || exit 1

declare -a RESULT_NAMES=()
declare -a RESULT_CODES=()
declare -a RESULT_METRICS=()

run_go_benchmark() {
  local name="$1"
  local pattern="$2"

  echo
  echo "=== ${name} ==="

  local output
  output=$(go test ./boatrace -bench "$pattern" -benchmem -benchtime=1x -run '^$' 2>&1)
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

run_go_benchmark "NumberOfWaysToWin (small race)" "BenchmarkNumberOfWaysToWin_SmallRace$"
run_go_benchmark "RecordBreakingsProducts (test input)" "BenchmarkRecordBreakingsProducts_InputTest$"
run_go_benchmark "RecordBreaking (test input)" "BenchmarkRecordBreaking_InputTest$"

echo
echo "=== Summary ==="
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
