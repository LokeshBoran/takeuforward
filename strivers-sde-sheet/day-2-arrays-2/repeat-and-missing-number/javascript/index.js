/**
 * @param {number[]} nums
 * @return {number, number}
 */
var repeatAndMissing = function(nums) {
    let total = nums.length * (nums.length + 1) / 2;
    let repeat;
    nums.forEach(n=>{
        total -= n;
        if(repeat == n) return;
        repeat = n
    });
    return [repeat, total + repeat];
    
};

console.log(repeatAndMissing([3,1,2,5,3])); // 3,4
console.log(repeatAndMissing([3,1,2,5,4,6,7,5])); // 5,8