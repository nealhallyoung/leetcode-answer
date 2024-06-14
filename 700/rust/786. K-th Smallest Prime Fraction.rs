pub fn kth_smallest_prime_fraction(arr: Vec<i32>, k: i32) -> Vec<i32> {
    let (mut left,mut right)=(0.,1.);
    loop{
        let mid=(left+right)/2.0;
        let mut count=0;
        let mut ans=vec![0,1];
        let mut i=0;
        for j in 1..arr.len(){
            while(arr[i] as f64)<mid*(arr[j] as f64){
                if arr[i]*ans[1]>arr[j]*ans[0]{
                    ans=vec![arr[i],arr[j]];
                }
                i+=1;
            }
            count += i as i32;
        }
        if count == k{
            break ans;
        }else if count<k{
            left=mid;
        }else{
            right=mid;
        }
    }
}