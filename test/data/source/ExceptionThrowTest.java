package me.rainstorm.jvm;

/**
 * @author baochen1.zhang
 * @date 2019.12.26
 */
public class ExceptionThrowTest {
    public static void main(String[] args) {
        try {
            for (String arg : args) {
                System.out.println(Integer.valueOf(arg).toString());
            }
        } catch (ArrayIndexOutOfBoundsException e) {
            System.out.println(ArrayIndexOutOfBoundsException.class.getName() + " " + e.getMessage());
        } catch (NumberFormatException e) {
            System.out.println(NumberFormatException.class.getName() + " " + e.getMessage());
        } catch (Exception e) {
            System.out.println(Exception.class.getName() + " " + e.getMessage());
        } finally {
            System.out.println("--- finally --");
        }
    }
}