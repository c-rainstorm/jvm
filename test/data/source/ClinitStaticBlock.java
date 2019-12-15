package me.rainstorm.jvm;

/**
 * @author baochen1.zhang
 * @date 2019.12.14
 */
public class ClinitStaticBlock {

    public static long staticLong;

    static {
        staticLong = 123432343234L;
    }
}
