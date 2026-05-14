use run::run;
use std::error::Error;

mod run;

fn main() -> Result<(), Box<dyn Error>> {
    run()?;

    Ok(())
}
