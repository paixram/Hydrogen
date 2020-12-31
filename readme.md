# Hydrogen 

Hydrogen is a programming language designed for the versatility and ease that can offer static data types and a new concept called BLOCKS.

`Interprete not compiled`

---

### Version:
Beta 0.0.1

### Warning:
Hydrogen is still under development, it is just in its early stages, it is almost useless.

`It is available only for windows at the moment for technical reasons.`

### How to use:
- Install the executable file
- Create a file with extension `.hy`
- Write code
- Run executable file and type the enxt command:
```
run ${filepath}
Example:
run C:/Users/main.hy
```

## Tutorial
- <h5>Types</h5>
`int`: Example: 234
`bool`: Example: true; false
`string`: example: "sdffsdsfd"
- <h5>Variables</h5>
Variables are declared with the reserved word `dec` followed by the name and then the type of the variable, then the assignment operator` -> `followed by the value of the variable. 
Example:
```
dec number int -> 345;
println(number); 
// OUTPUT: 345
```
- <h5>If</h5>
The if's are the normal ones in every programming language. Example:
```
if(3 < 5){
    return 34
}else{
    return 10
}
```

- <h5>Functions</h5>
The functions in an initial state, they need to mature enough, they are not yet able to check the return type, the syntax of a function is the following:
```
fn callPrint(x int) {
    println(x);
}
callPrint(234);
// OUTPUT: 234
```
As I said, the functions are not yet able to evaluate the return value, you can even insert a string in the parameter, it does not verify the type aprameter either.

- <h5>Aritmetic operators</h5>
```
4 + 5

5 / 5 + 4

8 + 4 + 2
```
---

## Experimentals
- `STOP` keyword
- `Block` Blocks
- `String` String data Type
- `functions` functions return value and parameters type.
- `macros` a macros implementation