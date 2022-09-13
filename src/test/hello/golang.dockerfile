FROM alpine:latest

FROM golang:1.14-alpine AS build
RUN apk add chromium --repository=http://dl-cdn.alpinelinux.org/alpine/v3.10/main
RUN apk add bash
RUN apk add wget gnupg ca-certificates

COPY --from=golang:1.14-alpine /usr/local/go/ /usr/local/go/

ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo
#
#FROM scratch
#COPY --from=build /bin/demo /bin/demo
CMD ["/bin/demo"]
