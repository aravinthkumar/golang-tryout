# golang-tryout

## Sample Program to search for lines in log file with parameters

To run the `go` program

> go run .

To provide the cli parameters 

- To search the lines where the text is `ABDD`

> go run . -text=ABDD

- To provide the log file location

> go run . -path=app.log

- To give the documentation on the possible cli paramters

> go run . -help 