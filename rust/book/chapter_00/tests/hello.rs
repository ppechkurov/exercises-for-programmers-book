use assert_cmd::Command;

#[test]
fn hello() {
    Command::cargo_bin("hello")
        .unwrap()
        .assert()
        .stdout("hello\n");
}
