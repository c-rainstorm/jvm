package me.rainstorm.jvm;

public interface Interface1 {
    default void sayHi() {
        System.out.println("hello from interface1");
    }
}