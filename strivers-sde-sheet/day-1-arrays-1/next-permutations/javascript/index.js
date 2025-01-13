/**
 * @param {number[]} nums
 * @return {void} Do not return anything, modify nums in-place instead.
 */
var nextPermutation = function(nums) {
    const n = nums.length;
    if (n <= 1) {
        return nums; // Handle edge cases: empty or single-element array
    }

    // 1. Find the first decreasing element from the right
    let i = n - 2;
    while (i >= 0 && nums[i] >= nums[i + 1]) {
        i--;
    }

    if (i < 0) {
        // Array is in descending order, reverse it to get the smallest permutation
        reverse(nums, 0, n - 1);
        return nums;
    }

    // 2. Find the smallest element to the right of nums[i] that is greater than nums[i]
    let j = n - 1;
    while (j > i && nums[j] <= nums[i]) {
        j--;
    }

    // 3. Swap nums[i] and nums[j]
    [nums[i], nums[j]] = [nums[j], nums[i]]; // Destructuring swap

    // 4. Reverse the subarray to the right of index i
    reverse(nums, i + 1, n - 1);
    return nums;
};


function reverse(nums, start, end) {
    while (start < end) {
        [nums[start], nums[end]] = [nums[end], nums[start]];
        start++;
        end--;
    }
}

console.log(nextPermutation([1,2,3])); // [1,3,2]
console.log(nextPermutation([3,2,1])); // [1,2,3]
console.log(nextPermutation([1,1,5])); // [1,5,1]