package me.rainstorm.jvm;

public class MyObject {
    public static int staticVar;
    public int instanceVar;

    public static void main(String[] args) {
        int x = 1234321;                      // ldc，这个值不能是 -1、0、1、2、3、4、5，会被替换为 IConst 指令
        MyObject myObj = new MyObject();      // new
        MyObject.staticVar = x + 1;           // putstatic
        x = MyObject.staticVar;               // getstatic

        myObj.instanceVar = x + 1;            // putfield
        x = myObj.instanceVar;                // getfield

        Object obj = myObj;
        if (obj instanceof MyObject) {        // instanceof
            myObj = (MyObject) obj;           // checkcast
            System.out.println(myObj.instanceVar); // getfield
        }
    }
}