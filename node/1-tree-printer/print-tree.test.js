import { printTree } from "./print-tree.js";

test("1 item", function () {
    expect(printTree({ name: "parent" })).toEqual(["parent"]);
});

test("2 items", function () {
    expect(
        printTree({
            name: "parent",
            items: [{ name: "child 1" }, { name: "child 2" }],
        })
    ).toEqual(["parent", "├─ child 1", "└─ child 2"]);
});
