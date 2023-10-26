pub mod api {
    include!(concat!(env!("OUT_DIR"), "/api.v1.rs"));
}

#[inline]
fn get_epoch() -> i32 {
    match std::time::SystemTime::now().duration_since(std::time::UNIX_EPOCH) {
        Ok(v) => v.as_secs_f32() as i32,
        Err(e) => {
            eprintln!("Failed to get unix epoch. Something is very wrong.\n    {e}");
            std::process::exit(-1);
        }
    }
}

fn main() {
    println!("Hello, backend.");

    let mut time = api::Time::default();

    time.name = String::from("Me");
    time.epoch = get_epoch();

    dbg!(time);
}
