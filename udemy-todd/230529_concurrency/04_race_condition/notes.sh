

go run main.go



# to check if the are race conditions 
go run -race main.go

# OS       linux
# ARCH     amd64
# CPU      4
# Goroutines       1
# ==================
# WARNING: DATA RACE
# Read at 0x00c00001a0c8 by goroutine 8:
#   main.main.func1()
#       /home/ubuntu/go_study/udemy-todd/230529_concurrency/04_race_condition/main.go:23 +0x30

# Previous write at 0x00c00001a0c8 by goroutine 14:
#   main.main.func1()
#       /home/ubuntu/go_study/udemy-todd/230529_concurrency/04_race_condition/main.go:25 +0x44

# Goroutine 8 (running) created at:
#   main.main()
#       /home/ubuntu/go_study/udemy-todd/230529_concurrency/04_race_condition/main.go:22 +0x259

# Goroutine 14 (finished) created at:
#   main.main()
#       /home/ubuntu/go_study/udemy-todd/230529_concurrency/04_race_condition/main.go:22 +0x259
# ==================
# CPU      4
# Goroutines       9
# Found 1 data race(s)
# exit status 66


