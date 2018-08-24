#!/usr/bin/env bash

# Set the path to the perl executable.
alias perl='/usr/bin/env perl'

# Exit with non-zero when an error occurs or an unset variable is evaluated.
set -o errexit
set -o nounset

# A square count greater than this threshold requires big-int capabilities of
# perl (or some other similarly-featured sub-program) to calculate the on-
# square and total grain counts.
export big_int_threshold=62
export chessboard_square_count=64

# Inform the user when bad input is received.
exit_with_error() {
    echo "Error: invalid input"
    exit 1
}

# The input value must be an integer with value 1.
is_one() {
    [[ $1 -eq 1 ]]
}

# The input value must be a an integer greater than zero.
is_positive_integer() {
    [[ $1 =~ ^[1-9][0-9]*$ ]]
}

# The input value must be an integer between 1 and 64.
is_valid_integer() {
    is_positive_integer $1 && [[ $1 -le 64 ]]
}

# Set the number of squares to evaluate.
set_square_count() {
    if !(is_positive_integer $1); then
        echo "$chessboard_square_count"
    else
        echo "$1"
    fi
}

# Constrain the number and content of command-line arguments:
#  - The first input value must be exactly one.
#  - The second input value must be an integer between 1 and 64, or the string
#    "total".
has_valid_args() {
    is_one $1 && is_valid_integer $2 || [[ $2 == "total" ]]
}

# Calculate the total number of grains of rice on a chessboard if each
# square contains two times the number of grains on the previous square
# and the first square contains a single grain of rice.
main() {
    is_total=$( [[ "$1" == "total" ]] && echo "true" || echo )
    square_count=$(set_square_count $1)

    current_square=1
    grains_on_square=1
    grains_on_board=1

    while [[ $current_square -lt $square_count ]]; do
        if [[ $current_square -le $big_int_threshold ]]; then
            (( grains_on_square = grains_on_square * 2 ))

            if [[ $is_total ]]; then
                (( grains_on_board = grains_on_board + grains_on_square ))
            fi
        else
            grains_on_square=$(perl -E "say $grains_on_square * 2")

            if [[ $is_total ]]; then
                grains_on_board=$(perl -E "say $grains_on_board + $grains_on_square")
            fi
        fi

        (( current_square = current_square + 1 ))
    done
  
    if [[ $is_total ]]; then
        echo $grains_on_board
    else
        echo $grains_on_square
    fi
}

# Validate the command-line arguments
if !(has_valid_args $# "$@")
then
    exit_with_error
fi

# Calls the main function passing all the arguments to it via '$@'
main "$@"
