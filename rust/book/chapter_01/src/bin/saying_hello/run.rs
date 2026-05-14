use std::io::Result;

pub fn run() -> Result<()> {
    let name = io::prompt("What is your name?")?;

    Ok(greet(&name))
}

fn greet(name: &str) {
    let greet = match name {
        "Parker" => format!("Vecher v hatu, Peter {name}!\n"),
        "Mechenyi" => format!("Zdarova, {name}!\n"),
        _ => format!("Hello, {name}, nice to meet you!\n"),
    };

    io::write(&greet);
}
