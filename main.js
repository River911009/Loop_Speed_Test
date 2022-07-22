function time(){
  var x=0;
  for(var i=0;i<1000000000;i++){
    x=x+i;
  }
  return(x);
}

console.log(time())
