use assert_cmd::Command;

#[test]
fn chapter_one_saying_hello() {
    Command::cargo_bin("cli")
        .unwrap()
        .write_stdin("Brian")
        .assert()
        .stdout(
            "\
What is your name? Brian
Hello, Brian, nice to meet you!
",
        );
}
