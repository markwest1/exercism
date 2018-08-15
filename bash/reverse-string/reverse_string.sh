#!/usr/bin/env bash

# Verify: one and only one argument
if [[ $# -ne 1 ]]; then
    echo "Error: enter one argument"
fi

# Initialize input and output varaibles
input=$1
output=

# Index starts at the number of characters in the input string
index=${#input}

# Decrement the index until it is zero
while [[ $index -gt 0 ]]; do
    (( index = index - 1 ))
    output="$output${input:$index:1}"
done

# Echo the output
echo $output
