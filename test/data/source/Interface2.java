package me.rainstorm.jvm;

public interface Interface2 {
    default void saySomething() {
        System.out.println("world from interface2");
    }
}