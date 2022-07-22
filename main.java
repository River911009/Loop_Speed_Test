public class main {
	public static long time(){
		long x=0;
		for(long i=0;i<1000000000;i++){
			x=x+i;
		}
		return(x);
	}

	public static void main(String[] args) {
		System.out.println(time());
	}
}
