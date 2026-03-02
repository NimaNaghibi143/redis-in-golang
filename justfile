default:
     @just --list # list all the available commands

# Compiles the Go code and creates the binary executable.
build:
  go build -o bin/goredis .
  ./bin/goredis

# connect to the server on a specified port
connect:
  telnet localhost 5001

# clean the bin directory
clean:
  rm -rf bin/*

