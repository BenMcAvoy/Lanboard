# Lanboard

A leaderboard application that is wrote with Go and Rust.

# Compiling from source

## Dependencies:

* [Just](https://github.com/casey/just) - For quick command running.
* [Go](https://go.dev) - For compiling the frontend.
* [Rust](https://www.rust-lang.org/) - For compiling the backend.
* [Protobuf](https://protobuf.dev/) - For compiling `.proto` files.
* [Protobuf go](https://pkg.go.dev/google.golang.org/protobuf@v1.28.0/cmd/protoc-gen-go) - For using Protobuf with Go.
* [Protobuf go rpc](https://pkg.go.dev/google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0) - For using Protobuf with Go.

### Optional:
* [gow](https://github.com/mitranim/gow) - For watching the frontend code.
* [cargo watch](https://crates.io/crates/cargo-watch) - For watching the backend code.

## Compiling:

Go to the `backend` folder and simply run:
```
just <platform name>
```
Where `<platform name>` is either `linux` or `windows`.

Go to the `frontend` folder and simply run the same command.
