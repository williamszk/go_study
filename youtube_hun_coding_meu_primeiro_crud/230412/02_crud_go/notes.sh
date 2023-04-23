
# https://www.youtube.com/watch?v=-71BdpuaEks&list=PLm-xZWCprwYQ3gyCxJ8TR1L2ZnUOPvOpr&index=3&ab_channel=HunCoding


go get github.com/joho/godotenv

# to run the executable
make build
./primary

make build
./secondary


# Centralize errors standardization
mkdir -p config/rest_err
touch config/rest_err/rest_err.go


go get -u github.com/gin-gonic/gin

mkdir controller/routes
touch controller/routes/routes.go

touch controller/find_user.go
touch controller/create_user.go
touch controller/delete_user.go
touch controller/update_user.go

mkdir controller/model
mkdir controller/model/request
mkdir controller/model/response

touch controller/model/request/user_request.go
touch controller/model/response/user_response.go