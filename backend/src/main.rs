pub mod api {
    include!(concat!(env!("OUT_DIR"), "/api.v1.rs"));
}

use tonic::transport::Server;
use tonic::{Request, Response, Status};

use api::leaderboard_server::{Leaderboard, LeaderboardServer};
use api::InsertionResponse;
use api::InsertionResult;
use api::Score;

#[derive(Debug, Default)]
pub struct LeaderboardImpl {}

#[tonic::async_trait]
impl Leaderboard for LeaderboardImpl {
    async fn insert(&self, request: Request<Score>) -> Result<Response<InsertionResponse>, Status> {
        println!("New request: {:#?}", request.into_inner());

        Ok(Response::new(InsertionResponse {
            result: InsertionResult::Okay as i32,
        }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "[::1]:50051".parse()?;
    let leaderboard = LeaderboardImpl::default();

    Server::builder()
        .add_service(LeaderboardServer::new(leaderboard))
        .serve(addr)
        .await?;

    Ok(())
}
