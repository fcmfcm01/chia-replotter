#!/usr/bin/env bash

cd hpool/miner
export VER=v1.4.0-2
envsubst < ./Dockerfile-template > Dockerfile
cd ..
docker-compose up -d