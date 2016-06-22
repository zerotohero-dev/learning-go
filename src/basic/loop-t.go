package main

func main() {
}

/*
RANNDOM MEANINGLESS NOTES BELOW:

          int count = 0;
        for (int i = N; i > 0; i /= 2) {
            for (int j = 0; j < i; j++) {
                count += 1;
            }
        }

                  int count = 0;
            for (int j = 0; j < N; j++) {
                for (int i = 0; i < j; i = j/2) {

                count += 1;
            }
        }

           int i, j, k = 0;
    for (i  = n/2; i <= n; i++) {
        for (j = 2; j <= n; j = j * 2) {
            k = k + n/2;
        }
    }

      int gcd(int n, int m) {
            if (n%m ==0) return m;
            if (n < m) swap(n, m);
            while (m > 0) {
                n = n%m;
                swap(n, m);
            }
            return n;
    }

    n    m
    15   2

    1
    2    1

    32   31
     31  1

     17 5

     252 105    ==> 21


Let us say n = fibonacci(N) and m = fibonacci(N - 1)

fibonacci(N) = fibonacci(N-1) + fibonacci(N-2)

OR n = m + k where k < m.

Therefore the step

n = n % m will make n = k

swap(n, m) will result in

n = fibonacci(N-1)

m = k = fibonacci(N-2)
So, it will take N steps before m becomes 0.

This means, in the worst case, this algorithm can take N step if n is Nth fibonacci number.

Think of whats the relation between N and n.








 */
