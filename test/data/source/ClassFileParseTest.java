package me.rainstorm.jvm;

public final class ClassFileParseTest implements Interface1, Interface2 {
    @Anno
    public static final boolean FLAG = true;
    public static final byte BYTE = 123;
    public static final char X = 'X';
    public static final short SHORT = 12345;
    public static final int INT = 123456;
    public static final long LONG = 12345654321L;
    public static final float FLOAT = 3.14f;
    public static final double DOUBLE = 2.71828;

    public static void main(String[] args) throws RuntimeException {
        System.out.println("Hello world!");
        ClassFileParseTest test = new ClassFileParseTest();
        ((Interface1)test).sayHi();
        ((Interface2)test).saySomething();
    }
}