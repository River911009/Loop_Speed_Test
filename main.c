#include <stdio.h>

long long time(void){
  long long x=0;
  for(long long i=0;i<1000000000;i++){
    x=x+i;
  }
  return(x);
}

int main(void){
  printf("%lld",time());
}
