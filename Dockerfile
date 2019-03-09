FROM node:10 AS builder-frontend

WORKDIR /usr/build
ADD frontend/package.json .
ADD frontend/package-lock.json .
RUN npm install
ADD frontend/public ./public
ADD frontend/src ./src
ADD frontend/tsconfig.json .
RUN npm run build



FROM golang:1.12.0-alpine3.9 AS builder-go

WORKDIR /usr
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh
ADD go.mod .
ADD go.sum .
RUN go mod download
ADD main.go .
ADD backend ./backend
RUN go build -o main .



FROM golang:1.12.0-alpine3.9

WORKDIR /usr
COPY --from=builder-frontend /usr/build/build frontend/build
COPY --from=builder-go /usr/main .
ENV GIN_MODE release
CMD ["./main"]