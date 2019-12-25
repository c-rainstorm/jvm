package me.rainstorm.jvm;

public class NativeTest {
    public static void main(String[] args) {
        System.out.println(void.class.getName());
        System.out.println(byte.class.getName());
        System.out.println(char.class.getName());
        System.out.println(short.class.getName());
        System.out.println(int.class.getName());
        System.out.println(long.class.getName());
        System.out.println(float.class.getName());
        System.out.println(double.class.getName());
        System.out.println(NativeTest.class.getName());
        System.out.println(Class.class.getName());
        System.out.println(int[].class.getName());
        System.out.println(int[][].class.getName());
        System.out.println(int[][][].class.getName());
        System.out.println(Object[].class.getName());
        System.out.println(Object[][].class.getName());
        System.out.println(Runnable.class.getName());
        System.out.println("abc".getClass().getName());
        System.out.println(new double[0].getClass().getName());
        System.out.println(new String[0].getClass().getName());
    }
}