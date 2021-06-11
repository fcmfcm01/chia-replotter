#!/bin/bash

function check_cmd() {
  cmd=$1
  if command -v $cmd >/dev/null 2>&1; then
    return 0
  else
    echo "$cmd not installed."
    return 1
  fi
}

function prepare_env() {
  if [[ $(check_cmd go) == "1" ]]; then
    echo "Installing go env...."
    wget -c https://golang.org/dl/go1.16.5.linux-amd64.tar.gz -O - | sudo tar -xz -C /usr/local
    echo "export PATH=$PATH:/usr/local/go/bin" | tee ~/.profile
    source ~/.profile
    go version
    check_cmd go
  fi
}
usage() {
  cat <<USAGE >&2
Usage:
    build.sh <--new=NEW_DIR> <--old=OLD_DIR> <-i|--interval=NUMBER_OF_SECONDS> [-h]
    --new=<NEW_DIR>     Directory contains new plots
    --old=<NEW_DIR>     Directory contains old plots
    -i|--interval       checking interval, in seconds, default 30
    -h                  print out usage
USAGE
  exit 1
}

NEW_DIR=""
OLD_DIR=""
INTERVAL=30

if [ $# -eq 0 ]; then
  usage
  exit 1
else
  # process arguments
  while [[ $# -gt 0 ]]; do
    case "$1" in
    --new=*)
      NEW_DIR="${1#*=}"
      shift 1
      ;;
    --old=*)
      OLD_DIR="${1#*=}"
      shift 1
      ;;
    -h)
      usage
      ;;
    -i | --interval)
      INTERVAL=$1
      shift 1
      ;;
    *)
      echo "Unknown argument: $1"
      usage
      ;;
    esac
  done
fi

if [[ "$NEW_DIR" == "" ]]; then
  echo "Error: you need to provide the new DIR."
  usage
fi
if [[ "$OLD_DIR" == "" ]]; then
  echo "Error: you need to provide the old DIR."
  usage
fi

prepare_env
go build -o plotter
./plotter --new=$NEW_DIR --old=$OLD_DIR --interval=$INTERVAL

