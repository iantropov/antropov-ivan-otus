const BRANCH_3 = "├─";
const BRANCH_2 = "└─";
const BRANCH_1 = "─";
const BRANCH = '│';

const TEST_INPUT = {
  name: 1,
  items: [
    { 
        name: 2, 
        items: [
            { name: 3 }, 
            { name: 4 }
        ] 
    },
    { 
        name: 5, 
        items: [
            { name: 6 }
        ] 
    },
  ],
};

function printTree(tree, treeNamePrefix = '', treePrefix = '') {
    console.log(`${treeNamePrefix}${tree.name}`);
    
    if (tree.items && tree.items.length > 0) {
        tree.items.forEach((leave, index) => {
            let leavePrefix = BRANCH_3;
            let nextTreePrefix = `${treePrefix}${BRANCH}  `;
            if (index === tree.items.length - 1) {
                leavePrefix = BRANCH_2;
                nextTreePrefix = `${treePrefix}   `
            }
            
            const leaveNamePrefix = `${treePrefix}${leavePrefix} `;
            
            printTree(leave, leaveNamePrefix, nextTreePrefix);
        });
    }
}

printTree(TEST_INPUT);
