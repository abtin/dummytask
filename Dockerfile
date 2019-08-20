FROM golang:1.12-alpine AS build

WORKDIR /src/
COPY main.go dummytask.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/dummytask

FROM scratch
COPY --from=build /bin/dummytask /bin/dummytask
ENTRYPOINT ["/bin/dummytask"]