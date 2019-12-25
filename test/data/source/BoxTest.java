package me.rainstorm.jvm;

import java.util.ArrayList;
import java.util.List;

/**
 * @author baochen1.zhang
 * @date 2019.12.25
 */
public class BoxTest {
    public static void main(String[] args) {
        List<Integer> integers = new ArrayList<>();

        for (int i = 1; i < 4; ++i) {
            integers.add(i);
        }

        System.out.println(integers.toString());
    }
}