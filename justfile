
# list justfile commands
default:
    just --list

alias list := default

# run main.go
run:
    go run main.go
