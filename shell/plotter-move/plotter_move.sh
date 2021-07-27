#!/usr/bin/env bash

function move_completed_file() {
  event_file=$1
  dest_dir=$2
  lsof "$event_file"
  if [[ $? -eq 1 ]]; then
    mv "$event_file" "$dest_dir"
  else
    echo "$event_file is opened."
  fi
}

# 监视的文件或目录
src_dir=$1
dest_dir=$2
old_dir=$3
if inotifywait; then
  sudo apt-get install inotify-tools -y
fi
inotifywait -mrq --format "%e %w%f" --event create,delete,modify,close,open "$src_dir" | while read event; do
  event_type=$(echo "$event" | awk '{print $1}')
  event_file=$(echo "$event" | awk '{print $2}')
  printf "%s :%s \n" "$(date)" "Change event:${event}"
  case $event_type in
  MODIFY | CREATE | OPEN)
    sleep 30s
    if [ -f "$event_file" ]; then
      lsof "$event_file" >/dev/null 2>&1
      if [[ $? -eq 1 ]]; then
        mv "$event_file" "$dest_dir"
        printf "%s :%s \n" "$(date)" "File {$event_file} has been moved."
        old_file=$(ls ${old_dir}/plot_* | tail -1)
        rm -f "$old_file"
        printf "%s :%s \n" "$(date)" "Old file [${old_file}] has been removed."
      else
        printf "%s :%s \n" "$(date)" "$event_file is opened."
      fi
    else
      continue
    fi
    ;;
  esac

done
