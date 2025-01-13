/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[], number[]}
 */
var mergeTwoSortedArraysWithoutExtraSpace = function(nums1, nums2) {
    let i = nums1.length - 1, j = 0;

    while (i >= 0 && j < nums2.length) {
        if (nums1[i] > nums2[j]){
            const tempNum1 = nums1[i];
            nums1[i] = nums2[j];
            nums2[j] = tempNum1;
        }
        i--;
        j++;
    }
    nums1.sort((a, b) => a - b);
    nums2.sort((a, b) => a - b);

    return [nums1, nums2];
    
};

console.log(mergeTwoSortedArraysWithoutExtraSpace([1,4,8,10],[2,3,9])); // Expected output: [1 2 3 4] [8 9 10]
console.log(mergeTwoSortedArraysWithoutExtraSpace([1,3,5,7], [0,2,6,8,9])); // Expected output: [0 1 2 3] [5 6 7 8 9]
console.log(mergeTwoSortedArraysWithoutExtraSpace([1,8,8], [2,3,4,5])); // Expected output: [1 2 3] [4 5 8 8]
console.log(mergeTwoSortedArraysWithoutExtraSpace([1,1,1,1], [2,2,3,3,5])); // Expected output: [1 1 1 1] [2 2 3 3 5]