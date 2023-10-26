pub mod api {
    include!(concat!(env!("OUT_DIR"), "/api.v1.rs"));
}

fn main() {
    println!("Hello, backend.")
}
