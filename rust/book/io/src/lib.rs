use std::io;
use std::io::{ErrorKind, IsTerminal, Write};

pub fn write(msg: &str) {
    write!(io::stdout(), "{msg}").unwrap();
    io::stdout().flush().unwrap();
}

pub fn prompt(msg: &str) -> io::Result<String> {
    write(&format!("{msg} "));

    let result = read_once()?;
    if result.is_empty() {
        return Err(ErrorKind::InvalidInput.into());
    }

    Ok(result)
}

pub fn read_once() -> io::Result<String> {
    let mut buf = String::new();
    let input = io::stdin();
    input.read_line(&mut buf)?;
    let res = buf.trim().to_string();

    if !io::stdin().is_terminal() {
        writeln!(io::stdout(), "{res}")?;
    }

    Ok(res)
}
