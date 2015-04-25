#!/usr/bin/env sh
# By default go benchmarks for one second,
# my machine doesn't have enough memory for this
# so run a bit shorter.
go test -v -bench . -benchtime 500ms
