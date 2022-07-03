docker pull golang:1.18

# Run this in:
# go_study/go_course/220702_06_go_project_on_containers/project-01
docker build -t my-simple-go-app:0.1 .

# Sending build context to Docker daemon  10.75kB
# Step 1/5 : FROM golang:1.18
#  ---> 46ae95f04a69
# Step 2/5 : RUN mkdir -p /home/app
#  ---> Running in 257788d653e0
# Removing intermediate container 257788d653e0
#  ---> 1a991ac98bd8
# Step 3/5 : COPY  ./app /home/app
#  ---> 8e9a9aab1e71
# Step 4/5 : WORKDIR /home/app
#  ---> Running in 9dd2bd30baac
# Removing intermediate container 9dd2bd30baac
#  ---> 8696f7f82fdb
# Step 5/5 : CMD ["go", "run", "main.go"]
#  ---> Running in d16191c5b80f
# Removing intermediate container d16191c5b80f
#  ---> a57f3cce2963
# Successfully built a57f3cce2963
# Successfully tagged my-simple-go-app:0.1

docker images

# REPOSITORY                                                 TAG         IMAGE ID       CREATED          SIZE
# my-simple-go-app                                           0.1         a57f3cce2963   47 seconds ago   964MB


docker run -d \
    -p 8080:8080 \
    --name my-simple-container \
    my-simple-go-app:0.1


docker stop my-simple-container
docker rm my-simple-container


# push the image to AWS ECR
# instead of using the local name let's use the AWS name
# create a new repository in ECR
# my-go-simple-app
# to run this command we need to have the AWS CLI configured in the host machine
aws ecr get-login-password --region sa-east-1 | docker login --username AWS --password-stdin 613363943992.dkr.ecr.sa-east-1.amazonaws.com

# Build docker image
docker build -t my-go-simple-app:0.1 .

# After the build completes, tag your image so you can push the image to this repository:
docker tag my-go-simple-app:0.1 613363943992.dkr.ecr.sa-east-1.amazonaws.com/my-go-simple-app:0.1

# Run the following command to push this image to your newly created AWS repository:
docker push 613363943992.dkr.ecr.sa-east-1.amazonaws.com/my-go-simple-app:0.1

docker run -d \
    -p 8080:8080 \
    --name my-go-simple-container \
    613363943992.dkr.ecr.sa-east-1.amazonaws.com/my-go-simple-app:0.1

docker stop my-go-simple-container
docker rm my-go-simple-container


