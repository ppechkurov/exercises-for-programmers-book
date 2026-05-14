use std::{io::ErrorKind, process};

pub fn run() -> std::io::Result<()> {
    let result = match io::prompt("What is the input string?") {
        Err(e) if e.kind() == ErrorKind::InvalidInput => {
            io::write("Empty input. Please enter something.\n");
            process::exit(2);
        }
        Err(e) => return Err(e),
        Ok(s) => s,
    };

    let count = result.len();
    io::write(format!("{result} has {count} characters\n").as_str());
    Ok(())
}
