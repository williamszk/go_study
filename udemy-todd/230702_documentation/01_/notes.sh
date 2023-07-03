

go mod init my_project

go run main.go

godoc -http=:8080

# --- seeing documentation in the command line --- 

# enter into the dir of a package and run go doc
cd mymath
go doc  # this will show the documentation of the package in the terminal
cd ..

go doc mymath
go doc mymath.Sum

go doc fmt
go doc fmt.Printf

go doc encoding/json
go doc encoding/json.Valid

go doc json.Number
go doc json.Number.Float64

# --- seeing documentation in the browser --- 

# show documention for package 
go doc <package-name>
go doc json


















