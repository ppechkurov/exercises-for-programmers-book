use std::fs::File;

fn main() {
    let res = File::open("/hello.txt");

    let f = match res {
        Ok(f) => f,
        Err(error) => match error.kind() {
            std::io::ErrorKind::NotFound => match File::create("/hello.txt") {
                Ok(fc) => fc,
                Err(e) => panic!("Problem creating a file: {e:?}"),
            },
            _ => panic!("Problem opening a file: {error:?}"),
        },
    };
}
