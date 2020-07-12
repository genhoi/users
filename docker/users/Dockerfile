FROM golang:1.14-alpine as build_server
RUN apk --no-cache --update add \
    curl \
    git \
    openssh-client
ADD . /go/src/users
WORKDIR /go/src/users
RUN export GO111MODULE=on && \
    go get -v && \
    go build -v


FROM alpine:3.12.0 as build_ui
ADD ./ui /ui/
WORKDIR /ui/
RUN apk add --update --no-cache npm && \
    npm install && \
    npm run build

FROM alpine:3.12.0
WORKDIR /app
COPY --from=build_server /go/src/users/users ./
COPY --from=build_server /go/src/users/migrations ./migrations
COPY --from=build_ui  /ui/dist ./ui/dist
ENTRYPOINT ["./users"]