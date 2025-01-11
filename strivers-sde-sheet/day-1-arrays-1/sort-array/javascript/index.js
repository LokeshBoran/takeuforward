/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */

var sortColors = function (nums) {
    const { noOfZeros,noOfOnes } = nums.reduce((acc, ele) => {
        if (ele === 0) {
            acc.noOfZeros++;
        } else if (ele === 1) {
            acc.noOfOnes++;
        }
        return acc;

    }, { noOfZeros: 0, noOfOnes: 0 });
    const noOfTwos = nums.length - noOfZeros - noOfOnes;

    nums.fill(0, 0, noOfZeros);
    nums.fill(1, noOfZeros, noOfZeros + noOfOnes);
    nums.fill(2, noOfZeros + noOfOnes, noOfZeros + noOfOnes + noOfTwos);
    
    return nums
};
sortColors([2,0,2,1,1,0]);
sortColors([2,0,1]);