FROM golang:1.16.5-buster AS builder

WORKDIR /app
ADD . .
RUN mkdir /out
RUN go build -o /out/main src/main.go

FROM scratch
WORKDIR /
COPY --from=build /out/main /main
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/main"]