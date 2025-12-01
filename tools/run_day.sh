#!/usr/bin/env bash
set -euo pipefail

# tools/run_day.sh <day> [--test] [--debug] [--build-dir build]
BUILD_DIR="build"
DAYS_DIR="data"
CMD_DIR="cmd"

if [[ $# -lt 1 ]]; then
  echo "Usage: $0 <day-number> [--test] [--debug] [--build-dir DIR] [--] [program args...]"
  exit 2
fi

day="$1"; shift

use_test=false
use_debug=false
prog_args=()

while [[ $# -gt 0 ]]; do
  case "$1" in
    --test) use_test=true; shift ;;
    --debug) use_debug=true; shift ;;
    --build-dir) BUILD_DIR="$2"; shift 2 ;;
    --) shift; prog_args=("$@"); break ;;
    -*) echo "Unknown option $1"; exit 3 ;;
    *) prog_args+=("$1"); shift ;;
  esac
done

if ! [[ "$day" =~ ^[0-9]+$ ]]; then
  echo "Day must be numeric"
  exit 4
fi

pad=$(printf "%02d" "$day")
pkg="cmd/day${pad}"
bin="${BUILD_DIR}/day${pad}"
data_dir="${DAYS_DIR}/day${pad}"
input_file="${data_dir}/input.txt"
test_file="${data_dir}/test_input.txt"

if [[ ! -d "${pkg}" ]]; then
  echo "Error: package ${pkg} not found"
  exit 5
fi

mkdir -p "${BUILD_DIR}"

echo "Building ${pkg} -> ${bin}"
go build -o "${bin}" "./${pkg}"

if $use_test; then
  if [[ -f "${test_file}" ]]; then
    in="${test_file}"
  else
    echo "Warning: test input not found (${test_file}), will run without stdin"
    in=""
  fi
else
  if [[ -f "${input_file}" ]]; then
    in="${input_file}"
  else
    echo "Warning: input not found (${input_file}), will run without stdin"
    in=""
  fi
fi

if $use_debug; then
  if ! command -v dlv >/dev/null 2>&1; then
    echo "Error: delve (dlv) not found. Install: go install github.com/go-delve/delve/cmd/dlv@latest"
    exit 6
  fi
  echo "Starting delve for ${bin} (interactive). Use 'continue' to run."
  if [[ -n "${in}" ]]; then
    exec dlv exec "${bin}" -- "${prog_args[@]}" < "${in}"
  else
    exec dlv exec "${bin}" -- "${prog_args[@]}"
  fi
else
  echo "Running ${bin} ..."
  if [[ -n "${in}" ]]; then
    exec "${bin}" "${prog_args[@]}" < "${in}"
  else
    exec "${bin}" "${prog_args[@]}"
  fi
fi
