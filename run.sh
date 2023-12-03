#!/bin/bash

# Determine the day
if [ "$#" -eq 1 ] && [ "$1" -ge 1 ] && [ "$1" -le 24 ]; then
    day=$(printf "%02d" "$1")
else
    day=$(date +"%d")
fi

# Run the main.go of the day
go run ./day$day/main.go