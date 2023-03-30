FROM golang:1.19 

WORKDIR /listapp

COPY . .

CMD ["go", "run", "main.go"]
