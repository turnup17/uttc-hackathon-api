FROM golang:1.18 as build

EXPOSE 8080
COPY . /app
WORKDIR /app
CMD ["go", "run", "main.go"]

#make image with:   docker build -t my-hackathon-api .　
#Map port 8000(host) to 8000(container) with:    docker run -d -p 8000:8000 my-hackathon-api