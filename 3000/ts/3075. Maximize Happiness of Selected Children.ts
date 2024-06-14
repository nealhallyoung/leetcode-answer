function maximumHappinessSum(happiness: number[], k: number): number {
    // 3075.maximize-happiness-of-selected-children
    let ans=0;
    happiness.sort((a,b)=>b-a);
    for(let i=0;i<k;i++){
        const x =happiness[i]-i;
        ans+=Math.max(x,0);
    }
    return ans
};