pub fn search<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    let mut res = Vec::new();
    for l in contents.lines() {
        if l.contains(query) {
            res.push(l);
        }
    }

    res
}

pub fn search_case_insesitive<'a>(query: &str, contents: &'a str) -> Vec<&'a str> {
    let query = query.to_lowercase();
    let mut res = Vec::new();

    for line in contents.lines() {
        if line.to_lowercase().contains(&query) {
            res.push(line);
        }
    }

    res
}

#[cfg(test)]
mod test {
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
    fn case_insesitive() {
        let query = "rUsT";

        let contents = "\
Rust:
safe, fast, productive.
Pick three.
Trust me.";
        assert_eq!(
            vec!["Rust:", "Trust me."],
            search_case_insesitive(query, contents)
        )
    }
}
