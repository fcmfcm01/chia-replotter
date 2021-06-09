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
