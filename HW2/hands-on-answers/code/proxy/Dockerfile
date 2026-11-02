FROM mirror.gcr.io/library/golang:1.22-alpine AS build
WORKDIR /src
COPY main.go go.mod ./
RUN CGO_ENABLED=0 go build -o /proxy .

FROM scratch
COPY --from=build /proxy /proxy
EXPOSE 80
ENTRYPOINT ["/proxy"]