# Go-Container-and-Iterator
This Go program implements a container and iterator interface and demonstrates their use with two different container types: a list and a vector. The program also includes functions for calculating the sum of the integers and float64s in a container.
## Usage
To run the program, use the following command:
```
go run main.go
```
The program will output the results of the container and iterator tests, as well as the sum of the integers and float64s in the containers.

## Example Output
```
Appending To List....
TESTING LIST:
50.5
10
11.1
11
15

Appending To Vector....
TESTING VECTOR:
10.66
11
12
13.25
14

SUM OF INTEGERS IN LIST: 36
SUM OF FLOAT64 IN LIST: 0
SUM OF INTEGERS IN VECTOR: 46
SUM OF FLOAT64 IN VECTOR: 48.91
```

## Note
The program defines the Container and Iterator interfaces, as well as two types that implement these interfaces: List and Vector. The List type is a linked list, while the Vector type is an array. The SumInt and SumFloat64 functions take a container as an argument and return the sum of the integers and float64s in the container, respectively.
