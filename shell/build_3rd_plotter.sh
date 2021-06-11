#!/usr/bin/env bash


git clone https://github.com/madMAx43v3r/chia-plotter.git
cd chia-plotter/
git submodule update --init
sudo apt-get install  make build-essential libsodium-dev  libgmp3-dev  g++ git -y
sudo apt remove --purge cmake -y
wget https://github.com/Kitware/CMake/releases/download/v3.20.3/cmake-3.20.3-linux-x86_64.sh
chmod +x cmake-3.20.3-linux-x86_64.sh
./cmake-3.20.3-linux-x86_64.sh  --skip-license
cp bin/cmake /usr/bin/
./make_devel.sh
./build/chia_plot  --help