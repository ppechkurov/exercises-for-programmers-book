use std::io::Write;

pub fn prompt(output: &mut impl Write, msg: &str) {
    write!(output, "{msg}").unwrap();
    output.flush().unwrap();
}
