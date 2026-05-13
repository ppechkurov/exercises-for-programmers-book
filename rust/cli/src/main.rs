use std::io;

use cli::run;

fn main() {
    run(&mut io::stdin().lock(), &mut io::stdout().lock());
}
