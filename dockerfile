FROM ALPINE AS build-stage

WORKDIR /app/easyhealth

COPY go.mod go.sum ./

COPY .. ./

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-gs-ping

FROM build-stage AS run-test-stage


EXPOSE 8080