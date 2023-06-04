# Solve race condition with mutex - mutual exclusion
# could we use a channel instead of a mutex to solve the problem? I think so.
# [ ] try to implement from the same code base, solve the race condition
#     using channels

go run main.go

# to check if the are race conditions 
go run -race main.go
# there is no race condition


