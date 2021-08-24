use std::error::Error;
use std::fs;

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let contents = fs::read_to_string(config.filename)?;

    Ok(())
}

fn search(key: String, contents: String) {}

pub struct Config {
    pub query: String,
    pub filename: String,
}

impl Config {
    pub fn new(args: &[String]) -> Result<Config, &str> {
        if args.len() < 3 {
            return Err("no enough arguments");
        }
        return Ok(Config {
            query: args[1].clone(),
            filename: args[2].clone(),
        });
    }
}

#[cfg(test)]
mod tests {}
