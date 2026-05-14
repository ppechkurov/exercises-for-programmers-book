use assert_cmd::Command;

#[test]
fn counting_number_chars() {
    Command::cargo_bin("counting_number_chars")
        .unwrap()
        .write_stdin("Homer\n")
        .assert()
        .stdout(
            "\
What is the input string? Homer
Homer has 5 characters
",
        );
}

#[test]
fn counting_number_chars_empty_input_loops() {
    Command::cargo_bin("counting_number_chars")
        .unwrap()
        .write_stdin("")
        .assert()
        .stdout(
            "\
What is the input string? 
Empty input. Please enter something.
",
        );
}
