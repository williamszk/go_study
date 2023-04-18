# https://www.youtube.com/watch?v=4i7_5NE6tlM&ab_channel=ArjunMahishi

# Build a Load Balancer



Round Robin



python server.py server-1 5000
python server.py server-2 5001
python server.py server-3 5002
python server.py server-4 5003
python server.py server-5 5004


go mod init my_load_balancer

go run main.go

curl 127.0.0.1:8000

for i in {1..20}; do curl 127.0.0.1:8000; done


# next video:
# https://www.youtube.com/watch?v=r9mcmZEhD9Q&ab_channel=ArjunMahishi
