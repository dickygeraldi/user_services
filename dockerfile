# FROM golang:latest
# RUN mkdir /app
# ADD . /app
# WORKDIR /app
# RUN go build -o app .
# FROM scratch
# CMD ["/app/main"]

FROM golang:latest as build-env
RUN mkdir /app
WORKDIR /app
COPY go.mod . 
COPY go.sum .

RUN go mod download
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user_services_app

FROM scratch 
COPY --from=build-env /app/user_services_app /app/user_services_app
ENTRYPOINT ["/app/user_services_app"]