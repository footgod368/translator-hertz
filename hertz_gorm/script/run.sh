#!/bin/sh
SCRIPT_DIR=$(dirname "$0")
cd "$SCRIPT_DIR/.."
go build -o hertz_gorm && ./hertz_gorm $@