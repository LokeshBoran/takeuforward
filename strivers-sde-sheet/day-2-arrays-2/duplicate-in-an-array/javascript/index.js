/**
 * @param {number[]} nums
 * @return {number}
 */
var findDuplicate = function(nums) {

    // Using Set
    // const set = new Set();
    // for (let i = 0; i < nums.length; i++) {
    //     if (set.has(nums[i])) {
    //         return nums[i];
    //     }
    //     set.add(nums[i]);
    // }
    // return -1;
    
    // Using Hashmap
    // const map = {};
    // for (let i = 0; i < nums.length; i++) {
    //     if (map[nums[i]]) {
    //         return nums[i];
    //     }
    //     map[nums[i]] = true;
    // }
    // return -1;
    
    // Using Fast and Slow pointers

    // let slow = nums[0];
    // let fast = nums[0];
    
    // do {
    //     slow = nums[slow];
    //     fast = nums[nums[fast]];
    // }
    // while (slow !== fast);
    
    // fast = nums[0];
    
    // while (slow !== fast) {
    //     slow = nums[slow];
    //     fast = nums[fast];
    // }
    
    // return slow;

    for (let i = 0; i < nums.length; i++) {
        const num = Math.abs(nums[i])
        if (nums[num] < 0) {
            return num
        }
        nums[num] = -nums[num]
    }
    return -1

};

console.log(findDuplicate([1,3,4,2,2])); // 2
console.log(findDuplicate([3,1,3,4,2])); // 3
console.log(findDuplicate([1,1])); // 1
console.log(findDuplicate([3,3,3,3,3])); // 3