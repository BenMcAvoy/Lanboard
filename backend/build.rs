fn main() -> std::io::Result<()> {
    prost_build::compile_protos(&["../shared/protobufs/api.proto"], &["../shared/protobufs/"])?;

    Ok(())
}
