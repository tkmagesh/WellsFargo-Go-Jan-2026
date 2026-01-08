# Golang

## Magesh Kuppan
- tkmagesh77@gmail.com
- 99019-11221

## Schedule
| What | When |
|------|------|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins)|
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up | 5:00 PM |

## Methodology
- No powerpoint presentation
- Discussion & Code

## Repository
- https://github.com/tkmagesh/WellsFargo-Go-Jan-2026


# Day-01 
## Why Go?
- Simplicity
    ONLY 25 keywords
    - No access modifiers (private/public/protected)
    - No reference types (Everything is a value by default)
    - No pointer arithmatic
    - No class (Only structs)
    - No Inheritance (Only composition)
    - No Exceptions (Only errors)
    - No try..catch..finally
    - No implicit type conversions
- Performance
    - On par with C++
    - Close to the hardware
        - Compiled to machine code
- Concurrency
    - Cheap & Efficient


## Program Structure
### Compilation
```shell
go build <filename.go>

# To create a binary in a different name
go build -o <binary_name> <filename.go>
```

### Cross Compilation
#### To get the list of supported platforms
```shell
go tool dist list
```

#### To list all the environment variables (go tool)
```shell
go env
```

#### To list specific environment variables
```shell
go env <var_1> <var_2> ....
```

#### To change the env variables
```shell
go env -w <var_1>=<new_value_1> <var_2>=<new_value_2> ...
```

#### Cross Compile (mac/unix/linux)
```shell
GOOS=<target_os> GOARCH=<target_arch> go build <filename.go>
```

#### In windows (powershell)
```powershell
$env:GOOS="windows"; $env:GOARCH="amd64"; go build <filename.go>
```

### Compile & Execute 
```shell
go run <filename.go>
```

### Standard Library Documentation
- https://pkg.go.dev/std



## Data Types
- string
- bool
- integers family
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers family
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points family
    - float32
    - float64
- complex family
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- aliases
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |


## Variables & Constants

### Function Scope
- CAN use ":="
- CANNOT have unused variables
### Package Scope
- CANNOT use ":="
- CAN have unused variables

## IOTA
## Programming Constructs
### if else
### switch case
### for
## Pointers
## Functions
    - Named Results
    - Variadic functions
    - Anonymous functions
        - nested functions
        - cannot have any name
        - they have to be immediately executed
    - Higher Order Functions (treat functions as data)
        - Assign "function" as a value to a variable
        - Pass functions as arguments to other functions
        - Return functions as return values
    - Deferred functions

## Collections
### Array
- Fixed sized typed collection
### Slice
- Varrying sized typed collection
- Slice maintains a pointer to an underlying array
![image](./images/slices.png)
### Map
- Typed collection of key/value pairs

## Error Handling
- Errors are values returned from a function (non "thrown")
- In practice, error objects in Go must implement "error" interface
- Creation
    - errors.New()
    - fmt.Errorf()
    - Custom type implementing the "error" interface

## Panic & Recovery
### Panic
- Represents the state of the application where the application execution cannot proceed further
- When the panic occurs, the application is shutdown AFTER all the deferred functions are executed
- use "panic(error)" to raise a panic

### Recovery
- Use "recover()" to recover from a panic

## Modules & Packages
### Module
- Any code that need to be versioned and deployed together
- Typically a folder with "go.mod" file
- `go.mod` file (manifest file of the application)
    - name of the module
        - should include the repo path
    - go runtime version targetted
    - references to other open source libraries & framewors used by the module

#### Create a module
```shell
go mod init <module_name>
```

#### Execute a module
```shell
go run .
```

#### Build a module
```shell
go build .
# 
go build -o <binary_name> .
```


### Package
- Internal orgaization of code in a module
- typically folders
