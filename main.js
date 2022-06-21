import { readAll } from "https://deno.land/std@0.144.0/streams/mod.ts";
import * as _ from "./wasm_exec.js"

let go = new window.Go()
let inst = await WebAssembly.instantiate(Deno.readFileSync("./y2j.wasm"), go.importObject)
go.run(inst.instance)


const line = new TextDecoder().decode(await readAll(Deno.stdin))
console.log(y2j(line))
