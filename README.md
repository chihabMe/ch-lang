## example

```
fnc fact(num){
    -> num*fact(-1)
}
fnc max(num1,num2){
    if num1 > num2 {
        -> num1
    }else{
        -> num2
    }
}

be i = 0 ;
be max = 10;
while i < max{
    echo(fact(i))
    i++
}
```
# ch-lang
