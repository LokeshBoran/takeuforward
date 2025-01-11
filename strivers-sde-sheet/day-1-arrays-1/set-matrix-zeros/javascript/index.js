/**
 * @param {number[][]} matrix
 * @return {void} Do not return anything, modify matrix in-place instead.
 */
const setZeroes = (matrix) => {
    const n = matrix.length;
    const m = matrix[0].length;

    let hasFirstColZero = false;
    // Step 1: Traverse the matrix and mark 1st row & col accordingly:
    for (let i = 0; i < n; i++) {
        for (let j = 0; j < m; j++) {
            if (matrix[i][j] === 0) {
                // Mark i-th row:
                matrix[i][0] = 0;

                // Mark j-th column:
                if (j !== 0) {
                    matrix[0][j] = 0;
                } else {
                    hasFirstColZero = true;
                }
            }
        }
    }

    // Step 2: Mark with 0 from (1,1) to (n-1, m-1):
    for (let i = 1; i < n; i++) {
        for (let j = 1; j < m; j++) {
            if (matrix[i][j] !== 0) {
                // Check for col & row:
                if (matrix[i][0] === 0 || matrix[0][j] === 0) {
                    matrix[i][j] = 0;
                }
            }
        }
    }

    // Step 3: Finally mark the 1st col & then 1st row:
    if (matrix[0][0] === 0) {
        for (let j = 0; j < m; j++) {
            matrix[0][j] = 0;
        }
    }
    if (hasFirstColZero) {
        for (let i = 0; i < n; i++) {
            matrix[i][0] = 0;
        }
    }

    return matrix;
}



console.log(setZeroes([[1,1,1],[1,0,1],[1,1,1]]))
console.log(setZeroes([[0,1,2,0],[3,4,5,2],[1,3,1,5]]))