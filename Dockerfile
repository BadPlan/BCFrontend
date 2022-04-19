FROM golang:1.18
RUN mkdir /app
COPY . /app
WORKDIR /app/cmd
EXPOSE 8080
RUN go build -o BCFrontend .
CMD ["/app/cmd/BCFrontend"]