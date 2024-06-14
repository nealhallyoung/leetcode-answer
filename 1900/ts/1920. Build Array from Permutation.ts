//1
function buildArray(nums: number[]): number[] {
    const ans: number[] = new Array(nums.length).fill(null);
    nums.forEach((n,i,arr) => {
        ans[i] = arr[n];
    })
    return ans;
};

//2
function buildArray(nums: number[]): number[] {
    let ans : number[] = []
    for(let i = 0; i < nums.length; i++) {
        ans.push(nums[nums[i]])
    }
    return ans
};
