package me.rainstorm.jvm;

/**
 * @author baochen1.zhang
 * @date 2019.12.14
 */
public class InvokeTest implements Runnable {

    public static void main(String[] args) {
        new InvokeTest().test();
    }

    public void test() {
        InvokeTest.staticMethod();                  // invokestatic
        InvokeTest invokeTest = new InvokeTest();   // invokespecial
        invokeTest.instanceMethod();                // invokevirtual

        super.equals(null);                         // invokespecial

        this.run();                                 // invokevirtual
        ((Runnable) invokeTest).run();              // invokeinterface
    }

    private void instanceMethod() {
    }

    private static void staticMethod() {
    }

    @Override
    public void run() {
    }
}
