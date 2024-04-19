# GO Learning
For example, the module path might be github.com/mymodule.\
`go mod edit -replace go_learning/greetings=../greetings`\
`go mod tidy`

## Compile and install the application
From the hello directory\
`$ go build`

To run the program\
`$ ./hello` (on Linux or Mac)
`$ hello.exe` (on Windiws)

You can see the install path\
`$ go list -f '{{.Target}}'`\

