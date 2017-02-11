public class Main {
    static {
        System.loadLibrary("demo");
    }

    public static void main(String[] args) {
        printOkay();
    }

    public static native void printOkay();
}
