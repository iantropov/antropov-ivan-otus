import { printTree } from './print-tree.js';

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

const strings = printTree(TEST_INPUT);
strings.forEach((line) => console.log(line));
