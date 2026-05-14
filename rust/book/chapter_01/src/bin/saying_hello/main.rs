use run::run;
use std::io;

mod run;

fn main() -> io::Result<()> {
    run()?;
    Ok(())
}
