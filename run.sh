#!/bin/bash

# Determine the current day
today=$(date +"%d")

# Run the main.go of the current day
go run ./day$today/main.go