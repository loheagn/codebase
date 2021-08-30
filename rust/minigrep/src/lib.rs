use std::error::Error;
use std::fs;

pub fn run(config: Config) -> Result<(), Box<dyn Error>> {
    let contents = fs::read_to_string(config.filename)?;

    let results: Vec<&str> = if config.case_sensitive {
        search(&config.query, &contents)
    } else {
        search_case_insensitive(&config.query, &contents)
    };

    for line in results {
        println!("{}", line)
    }

    Ok(())
}

fn search<'a>(key: &str, contents: &'a str) -> Vec<&'a str> {
    let mut result = vec![];

    for line in contents.lines() {
        if line.contains(key) {
            result.push(line)
        }
    }

    result
}

fn search_case_insensitive<'a>(key: &str, contents: &'a str) -> Vec<&'a str> {
    let key = key.to_lowercase();
    let mut result = vec![];

    for line in contents.lines() {
        if line.to_lowercase().contains(&key) {
            result.push(line)
        }
    }

    result
}

pub struct Config {
    pub query: String,
    pub filename: String,
    pub case_sensitive: bool,
}

impl Config {
    pub fn new(args: &[String]) -> Result<Config, &str> {
        if args.len() < 3 {
            return Err("no enough arguments");
        }
        let case_sensitive = std::env::var("CASE_INSENSITIVE").is_err();
        Ok(Config {
            query: args[1].clone(),
            filename: args[2].clone(),
            case_sensitive,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn case_sensitive() {
        let query = "duct";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.
Duct tape.";
        assert_eq!(vec!["safe, fast, productive."], search(query, contents))
    }

    #[test]
    fn case_insensitive() {
        let query = "rUsT";
        let contents = "\
Rust:
safe, fast, productive.
Pick three.
Trust me.";
        assert_eq!(
            vec!["Rust:", "Trust me."],
            search_case_insensitive(query, contents)
        )
    }
}
