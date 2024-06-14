function kthSmallestPrimeFraction(arr: number[], k: number): number[] {
    let left = 0, right = 1, n = arr.length;
    while (true) {
        const mid = (left + right) / 2.0;
        let i = 0, count = 0;
        let x = 0, y = 1;
        for (let j = 1; j < n; j++) {
            while (arr[i] < mid * arr[j]){
                if (arr[i] * y > arr[j] * x){
                    x = arr[i];
                    y = arr[j];
                }
                i++;
            }
            count += i;
        }
        if (count == k) return [x, y];
        if (count > k) right = mid;
        else left = mid;
    }
};