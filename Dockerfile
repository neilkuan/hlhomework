FROM --platform=linux/amd64 public.ecr.aws/docker/library/golang:1.20.8-alpine3.17 AS builder
WORKDIR /worker
COPY . .
RUN go build -o app


FROM --platform=linux/amd64 public.ecr.aws/docker/library/alpine:3.17
WORKDIR /worker
COPY --from=builder /worker/app .
EXPOSE 8080
CMD [ "/worker/app" ] 