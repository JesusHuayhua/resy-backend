#!/usr/bin/env bash
set -euo pipefail

COMMAND=${1:-}

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
SOA_DIR="$SCRIPT_DIR/.."
ENTRYPOINTS_DIR="$SOA_DIR/cmd/entrypoints"
BIN_DIR="$SOA_DIR/cmd/bin"

function usage() {
  cat <<EOF
Usage: $(basename "$0") <command>
  build-local   Inicializa go.mod (si hace falta) y compila todos los entrypoints en $BIN_DIR
  run-local     Ejecuta build-local y luego arranca cada ejecutable en segundo plano
EOF
  exit 1
}

function init_mod() {
  pushd "$SOA_DIR" >/dev/null
    if [[ ! -f go.mod ]]; then
      echo "[GO MOD] go.mod no encontrado, inicializando módulo 'soa'..."
      go mod init soa
    else
      echo "[GO MOD] go.mod ya existe, omitiendo init."
    fi
    go mod tidy
  popd >/dev/null
}

function build_local() {
  init_mod
  mkdir -p "$BIN_DIR"
  for dir in "$ENTRYPOINTS_DIR"/*/; do
    svc_name="$(basename "$dir")"
    echo "[BUILD] $svc_name"
    pushd "$dir" >/dev/null
      go build -ldflags="-s -w" -o "$BIN_DIR/$svc_name"
    popd >/dev/null
  done
  echo "Compilacion completada."
}

function run_local() {
  build_local
  echo "Arrancando servicios..."
  mkdir -p /tmp/build_sh_logs
  for dir in "$ENTRYPOINTS_DIR"/*/; do
    svc_name="$(basename "$dir")"
    exe="$BIN_DIR/$svc_name"
    if [[ -x "$exe" ]]; then
      echo "[RUN] $svc_name (logs: /tmp/build_sh_logs/${svc_name}.log)"
      nohup "$exe" > "/tmp/build_sh_logs/${svc_name}.log" 2>&1 &
    else
      echo "WARNING: $exe no encontrado o no es ejecutable, se omite."
    fi
  done
  echo "Todos los servicios arrancados en segundo plano."
  echo "Puedes revisar los logs en /tmp/build_sh_logs/"
}

case "$COMMAND" in
  build-local) build_local ;;
  run-local)   run_local ;;
  *)           usage ;;
esac
