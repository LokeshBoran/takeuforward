/**
 * @param {number[][]} intervals
 * @return {number[][]}
 */
var merge = function(intervals) {
    
    intervals.sort((a,b) => a[0] - b[0]);

    let result = [];
    let currentInterval = intervals[0];

    for (let i = 1; i < intervals.length; i++) {
        const current = intervals[i];
        if (current[0] <= currentInterval[1]) {
            currentInterval[1] = Math.max(currentInterval[1], current[1]);
        } else {
            result.push(currentInterval);
            currentInterval = current;
        }
    }
    result.push(currentInterval);
    return result;
};

console.log(merge([[1,3],[2,6],[8,10],[15,18]])); // [[1,6],[8,10],[15,18]]
console.log(merge([[1,4],[4,5]])); // [[1,5]]