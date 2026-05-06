const std = @import("std");
const Io = std.Io;

pub fn main(init: std.process.Init) !void {
    const cwd = std.Io.Dir.cwd();
    const file = try cwd.openFile(
        init.io,
        "doesnt_exist.txt",
        .{ .mode = .write_only },
    );
    std.log.info("test: {}", .{file});
}
