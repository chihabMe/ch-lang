# ch-lang

ch-lang is an interpreted programming language built with Go. It is designed to
be simple, yet powerful, providing an easy-to-use syntax for both beginners and
experienced programmers.

## Example / Syntax

```ch
// Define a function to calculate the factorial of a number
fnc fact(num) {
    if num <= 1 {
        return  1 ;
    } else {
        return  num * fact(num - 1) ;
    }
}

// Define a function to find the maximum of two numbers
fnc max(num1, num2) {
    if (num1 > num2) {
        return num1 ;
    } else {
        return  num2 ;
    }
}

// Initialize variables
set i = 0;
set max = 10;

// Loop to calculate and echo the factorial of numbers from 0 to 'max'
while i < max {
    echo(fact(i))
    i++
}
```
