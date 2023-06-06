#!/bin/bash

bin_name=gedis
bin_path="bin"
version="1.0.0"
flags="-tags netgo -ldflags=-s -ldflags=-w"

declare -a os=("linux" "darwin")
declare -a arch=("amd64" "arm64" "arm")
declare -a incompatible=("darwin-armv7")

mkdir -p $bin_path

for os_target in "${os[@]}"
do
  for arch_target in "${arch[@]}"
  do
    if [[ " ${incompatible[*]} " == *"$os_target-$arch_target"* ]]; then
      continue
    fi

    export GOOS=$os_target
    export GOARCH=$arch_target

    output_path=$bin_path/$bin_name_$version-$os_target-$arch_target
    go build $flags -o $output_path/$bin_name

    if [ $? -eq 0 ]; then
      echo "Build successful for $os_target-$arch_target"
      cp LICENSE.md $output_path/LICENSE.md

      if [ -n "$(find "$output_path" -type f)" ]; then
        tar -czf "$output_path.tar.gz" $output_path
        echo "Compressed $os_target-$arch_target"
      else
        echo "Output folder is empty. Skipping compression."
      fi
    else
      echo "Build failed for $os_target-$arch_target"
    fi
  done
done