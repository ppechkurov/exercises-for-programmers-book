use std::io::{BufRead, IsTerminal, Write};

fn prompt(output: &mut impl Write, msg: &str) {
    write!(output, "{msg}").unwrap();
    output.flush().unwrap();
}

fn read_name(mut input: impl BufRead) -> String {
    let mut name = String::new();
    input.read_line(&mut name).unwrap();
    name.trim().to_string()
}

fn say_hello(name: &str, output: &mut impl Write) {
    writeln!(output, "Hello, {name}, nice to meet you!").unwrap();
}

pub fn run(input: &mut impl BufRead, output: &mut impl Write) {
    prompt(output, "What is your name? ");
    let name = read_name(input);

    if !std::io::stdin().is_terminal() {
        writeln!(output, "{name}").unwrap();
    }

    say_hello(&name, output);
}
