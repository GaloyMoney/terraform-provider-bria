#!/bin/bash

function gpt_files() {
  if [ $# -eq 0 ]; then
    echo "Usage: gpt_files glob_or_dir1 [glob_or_dir2 ...]"
    return 1
  fi

  process_file() {
    local file="$1"
    echo "=== File: $file ==="
    echo ""
    cat "$file"
    echo ""
    echo "=== End of file: $file ==="
    echo ""
  }

  for input in "$@"
  do
    if [ -d "$input" ]; then
      files=( $(find "$input" -type f) )
    else
      files=( $input )
    fi

    if [ ${#files[@]} -eq 0 ]; then
      echo "Error: No files match input '$input'."
      continue
    fi

    for file in "${files[@]}"
    do
      if [ -f "$file" ]; then
        process_file "$file"
      else
        echo "Error: $file is not a file."
      fi
    done
  done
}

