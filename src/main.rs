use reqwest;
use scraper::{Html, Selector};
use std::fs;

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let body = reqwest::get("https://books.toscrape.com/").await?.text().await?;
    let doc = Html::parse_document(&body);
    let sel = Selector::parse("h3 a").unwrap();
    let count = doc.select(&sel).count();
    fs::write("data/Rust.json", format!("[{{ \"count\": {} }}]", count))?;
    println!("{}", count);
    Ok(())
}

