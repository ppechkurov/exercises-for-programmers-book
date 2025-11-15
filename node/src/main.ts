import { Console, Effect } from "effect";
import { exported } from "./another.ts";

const program = Console.log(exported);

Effect.runSync(program);
