const BRANCH_3 = "├─";
const BRANCH_2 = "└─";
const BRANCH = "│";

function printTreeIntoArray(tree, treeNamePrefix, treePrefix, output) {
    output.push(`${treeNamePrefix}${tree.name}`);

    if (tree.items && tree.items.length > 0) {
        tree.items.forEach((leave, index) => {
            let leavePrefix = BRANCH_3;
            let nextTreePrefix = `${treePrefix}${BRANCH}  `;
            if (index === tree.items.length - 1) {
                leavePrefix = BRANCH_2;
                nextTreePrefix = `${treePrefix}   `;
            }

            const leaveNamePrefix = `${treePrefix}${leavePrefix} `;

            printTreeIntoArray(leave, leaveNamePrefix, nextTreePrefix, output);
        });
    }
}

export function printTree(tree) {
    const output = [];
    printTreeIntoArray(tree, "", "", output);
    return output;
}
