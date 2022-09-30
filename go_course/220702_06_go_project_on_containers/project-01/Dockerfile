FROM golang:1.18

RUN mkdir -p /home/app

COPY  ./app /home/app

WORKDIR /home/app

CMD ["go", "run", "app.go"]
