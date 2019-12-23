package me.rainstorm.jvm;

/**
 * @author baochen1.zhang
 * @date 2019.12.16
 */
public class PrimitiveArrayTest {
    public static void main(String[] args) {
        int[] intArray = new int[10];                      // newarray
        intArray[0] = 346543;
        System.out.println(intArray[0]);
        System.out.println(intArray.length);

        boolean[] booArray = new boolean[9];
        booArray[0] = true;
        System.out.println(booArray[0]);
        System.out.println(booArray.length);

        char[] chars = new char[8];
        chars[0] = 'h';
        System.out.println(chars[0]);
        System.out.println(chars.length);

        short[] shorts = new short[7];
        shorts[0] = 24543;
        System.out.println(shorts[0]);
        System.out.println(shorts.length);

        long[] longs = new long[6];
        longs[0] = 23454323543L;
        System.out.println(longs[0]);
        System.out.println(longs.length);

        float[] floats = new float[5];
        floats[0] = 3.14f;
        System.out.println(floats[0]);
        System.out.println(floats.length);

        double[] doubles = new double[4];
        doubles[0] = 3.141592653;
        System.out.println(doubles[0]);
        System.out.println(doubles.length);

        int[][][] ints = new int[3][4][5];
        ints[0][0][0] = 12345;
        System.out.println(ints[0][0][0]);
    }
}
