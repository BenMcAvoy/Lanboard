# Lanboard

A leaderboard application that is wrote with Go and Rust.

# Compiling from source

## Dependencies:

* [Just](https://github.com/casey/just) - For quick command running.
* [Go](https://go.dev) - For compiling the frontend.
* [Rust](https://www.rust-lang.org/) - For compiling the backend.

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
