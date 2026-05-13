use assert_cmd::Command;

#[test]
fn saying_hello() {
    Command::cargo_bin("saying_hello")
        .unwrap()
        .write_stdin("Peter\n")
        .assert()
        .stdout("What is your name? Peter\nHello, Peter, nice to meet you!\n");
}

#[test]
fn saying_hello_for_different_people() {
    Command::cargo_bin("saying_hello")
        .unwrap()
        .write_stdin("Parker\n")
        .assert()
        .stdout("What is your name? Parker\nVecher v hatu, Peter Parker!\n");

    Command::cargo_bin("saying_hello")
        .unwrap()
        .write_stdin("Mechenyi\n")
        .assert()
        .stdout("What is your name? Mechenyi\nZdarova, Mechenyi!\n");
}
