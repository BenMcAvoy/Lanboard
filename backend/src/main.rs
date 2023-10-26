pub mod api {
    include!(concat!(env!("OUT_DIR"), "/api.v1.rs"));
}

use tracing::{Level, info};

use structopt::StructOpt;

use tonic::transport::Server;
use tonic::{Request, Response, Status};

use api::leaderboard_server::{Leaderboard, LeaderboardServer};
use api::InsertionResponse;
use api::InsertionResult;
use api::Score;

#[derive(StructOpt)]
#[structopt(name = "lanboard")]
struct Config {
    /// The host to run on
    #[structopt(short, long, default_value = "0.0.0.0")]
    host: String,
    
    /// The port to run on
    #[structopt(short, long, default_value = 50051)]
    port: u32,
}

#[derive(Debug, Default)]
pub struct LeaderboardImpl {}

#[tonic::async_trait]
impl Leaderboard for LeaderboardImpl {
    async fn insert(&self, request: Request<Score>) -> Result<Response<InsertionResponse>, Status> {
        info!("New request: {:#?}", request.into_inner());

        Ok(Response::new(InsertionResponse {
            result: InsertionResult::Okay as i32,
        }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    tracing_subscriber::fmt()
        .with_thread_names(true)
        .with_line_number(true)
        .with_max_level(Level::INFO)
        .without_time()
        .init();
        
    let config = Config::from_args();

    let addr = format!("{}:{}", config.host, config.port).parse()?;
    let leaderboard = LeaderboardImpl::default();

    info!("Listening on {addr}");

    Server::builder()
        .add_service(LeaderboardServer::new(leaderboard))
        .serve(addr)
        .await?;

    Ok(())
}
