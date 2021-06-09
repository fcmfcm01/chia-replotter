#!/usr/bin/env bash

function check_cmd() {
  cmd=$1
  if command -v $cmd >/dev/null 2>&1; then
    return 0
  else
    echo "$cmd not installed."
    return 1
  fi
}

function setup_docker_related() {
  check_cmd docker
  if [[ $? == "1" ]]; then
    echo "Installing docker ..."
    curl -sSL https://get.docker.com/ | sh
  else
    echo "Docker already installed, continue..."
  fi
  check_cmd docker-compose
    if [[ $? == "1" ]]; then
      echo "Installing docker-compose"
      curl -L \
      "https://github.com/docker/compose/releases/download/1.25.4/docker-compose-$(uname -s)-$(uname -m)" \
      -o /usr/local/bin/docker-compose
    else
      echo "docker-compose already installed, continue..."
    fi
}
setup_docker_related


cd docker
docker build --network=host -t chia-docker .
cd -
