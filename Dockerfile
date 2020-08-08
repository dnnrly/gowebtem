FROM golang:1.14-alpine as builder

ENV GO111MODULE=on
EXPOSE 8080

RUN mkdir /project
COPY . /project
WORKDIR /project

RUN go build ./cmd/gowebtem

FROM golang:1.14-alpine

RUN mkdir /app
RUN addgroup -S appuser && adduser -S appuser -G appuser

WORKDIR /app
COPY --from=builder /project/gowebtem .

RUN chown -R appuser:appuser /app
USER appuser

CMD ./gowebtem