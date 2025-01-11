/**
 * @param {number[]} prices
 * @return {number}
 */
var maxProfit = function(prices) {
    let maxProfit = 0;
    let minPrice = prices[0];

    // for (let i = 1; i < prices.length; i++) {
    //     minPrice = Math.min(minPrice, prices[i]);
    //     maxProfit = Math.max(maxProfit, prices[i] - minPrice);
    // }
    // for (const p of prices) {
    //     minPrice = Math.min(minPrice, p);
    //     maxProfit = Math.max(maxProfit, p - minPrice);
    // }


    // const { maxProfit }= prices.reduce((acc, p) => {
    //     acc.minPrice = Math.min(acc.minPrice, p);
    //     acc.maxProfit = Math.max(acc.maxProfit, p - acc.minPrice);
    //     return acc;
    // }, { minPrice: prices[0], maxProfit: 0 });

    for (let i = 1; i < prices.length; i++) {
        minPrice = minPrice > prices[i] ? prices[i] : minPrice;
        maxProfit = maxProfit > prices[i] - minPrice ? maxProfit : prices[i] - minPrice;
    }

    return maxProfit;

};


console.log(maxProfit([7,1,5,3,6,4])); // 5
console.log(maxProfit([7,6,4,3,1])); // 0