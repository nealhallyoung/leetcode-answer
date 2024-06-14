//1
function finalValueAfterOperations(operations: string[]): number {
    let res = 0;
    for (const opt of operations) {
        res = opt[1] === "+" ? res + 1 : res - 1;
    }
    return res;
};

//2
function finalValueAfterOperations(operations: string[]): number {
    let res=0;
    operations.forEach((el)=>{
        el.includes('+')?res++:res--;
    })
    return res
  };
