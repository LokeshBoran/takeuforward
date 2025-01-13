/**
 * @param {number[]} nums
 * @return {number}
 */
var getInversions = function(nums) {
    let inversions = 0;
    for (let i = 0; i < nums.length; i++) {
        for (let j = i + 1; j < nums.length; j++) {
            if (nums[i] > nums[j]) {
                inversions++;
            }
        }
    }
    
    return inversions;
};

console.log(getInversions([1, 2, 3, 4, 5])); // 0
console.log(getInversions([5, 4, 3, 2, 1])); // 10
console.log(getInversions([5, 3, 2, 1, 4])); // 7