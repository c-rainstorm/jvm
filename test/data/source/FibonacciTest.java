package me.rainstorm.jvm;

/**
 * @author baochen1.zhang
 * @date 2019.12.14
 */
public class FibonacciTest {
    public static void main(String[] args) {
        System.out.println(fibonacci(20));
    }

    private static long fibonacci(long n) {
        if (n <= 1) {
            return n;
        }
        long result = fibonacci(n - 1) + fibonacci(n - 2);
        System.out.println(result);
        return result;
    }
}
