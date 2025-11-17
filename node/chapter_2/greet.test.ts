import { describe, expect, it } from "@effect/vitest";
import { Effect } from "effect";
import { greet } from "./main.ts";

describe("Saying Hello", () => {
	it.effect("test success", () =>
		Effect.gen(function* () {
			const result = yield* greet("me");
			expect(result).toBe("Hello, me, nice to meet you!");
		}),
	);
});
