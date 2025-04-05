#!/bin/bash

# Define the application name and build path
APP_NAME="palplan"
BUILD_PATH="./cmd/${APP_NAME}/${APP_NAME}_executable"
MAIN_FILE="./cmd/${APP_NAME}/main.go"

# Build the Go executable
go build -o "${BUILD_PATH}" "${MAIN_FILE}"

# Check for build errors
if [ $? -ne 0 ]; then
  echo "Error building the application."
  exit 1
fi

# Run the Go executable
"${BUILD_PATH}"

# Optional: Add error handling for the application run
if [ $? -ne 0 ]; then
  echo "Error running the application."
  exit 1
fi
