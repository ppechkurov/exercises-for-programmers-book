use assert_cmd::Command;

#[test]
fn hello() {
    Command::cargo_bin("hello")
        .unwrap()
        .write_stdin("Peter")
        .assert()
        .stdout("hello\n");
}
