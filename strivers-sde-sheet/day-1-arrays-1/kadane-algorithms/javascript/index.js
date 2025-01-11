/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    // if (nums.length === 0) return 0;

    let maxSum = nums[0];
    let currentSum = nums[0];

    // for (const num of nums) {
    //     currentSum = Math.max(num, currentSum + num);
    //     maxSum = Math.max(maxSum, currentSum);
    // }

    for (let i = 1; i < nums.length; i++) {
        currentSum = Math.max(currentSum + nums[i], nums[i]);
        maxSum = Math.max(maxSum, currentSum);
    }

    return maxSum;
    
};

console.log(maxSubArray([-2,1,-3,4,-1,2,1,-5,4])); // 6
console.log(maxSubArray([1])); // 1
console.log(maxSubArray([-1])); // -1
console.log(maxSubArray([-100000])); // -100000
console.log(maxSubArray([5,4,-1,7,8])); // 23
console.log(maxSubArray([1,-1,5,-2,3])); // 6
console.log(maxSubArray([-2,-3,-4])); // -2