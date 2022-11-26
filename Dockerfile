FROM golang:1.19.3-bullseye
RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build
CMD ["/app/discordbot"]