static class Singleton {
    private static volatile Singleton instance = null;
    String value = null;

    private Singleton() {
        
    }

    public static Singleton getInstance() {
        if(instance == null) {
            synchronized(Singleton.class) {
                instance = new Singleton();
            }
        }
        return instance;
    }

    public String getValue() {
        return value;
    }

    public void setValue(String value) {
        this.value = value;
    }
    
}
