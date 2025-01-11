/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
    const triangle = [[1]];

    for (let i = 1; i < numRows; i++) {
        const newElement = [1]
        const lastRaw = triangle[i - 1];

        for (let j = 1; j < i; j++) {
            newElement.push(lastRaw[j - 1] + lastRaw[j]);
        }
        newElement.push(1);
        triangle.push(newElement);
    }

    return triangle;
};


console.log(generate(5));
console.log(generate(1));
console.log(generate(15));