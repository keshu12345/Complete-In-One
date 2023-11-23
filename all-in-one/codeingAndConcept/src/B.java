class B extends A{
    int x;

    public void setX(int x) {
        this.x = x;
    }

    @Override
    public int getX(){
        return 10;
    }
}