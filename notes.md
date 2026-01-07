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
## Variables & Constants
## IOTA
## Programming Constructs
### if else
### switch case
### for
## Pointers
## Functions