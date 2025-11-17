import { Effect, Console, Schedule, pipe } from "effect";
import { NodeRuntime } from "@effect/platform-node";

const program = pipe(
	Effect.addFinalizer(() => Console.log("Exiting...")),
	Effect.andThen(Console.log("Application started!")),
	Effect.andThen(
		Effect.repeat(Console.log("still alive..."), {
			schedule: Schedule.spaced("1 second"),
		}),
	),
	Effect.scoped,
);

NodeRuntime.runMain(program);
