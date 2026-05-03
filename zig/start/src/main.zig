const std = @import("std");
const Io = std.Io;

pub fn main(_: std.process.Init) !void {
    std.debug.print("test", .{});
    const a = "string";
    a = "another";
}
