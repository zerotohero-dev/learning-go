package main

func arrayPairSum(nums [10000]int) {
    arr := [20001]int
    lim := 10000

    for k, v: range nums {
        arr[v + lim]++
    }

    d, sum = 0;

    for i := -10000; i <= 10000; i++ {
        sum += ((arr[i + lim] + 1 -d)/2)*i
        d = (2 + arr[i+lim] - d) % 2
    }

    return sum
}