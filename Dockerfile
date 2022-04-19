FROM golang:1.14.2-alpine3.11

RUN apk update && apk --no-cache add ca-certificates

RUN apk update && apk add tzdata
RUN cp /usr/share/zoneinfo/Asia/Kolkata /etc/localtime
RUN echo "Asia/Kolkata" > /etc/timezone

RUN apk update && apk add git
RUN mkdir -p /blood_bank_system_service/
ADD . /blood_bank_system_service/

ENV SERVICE=blood_bank_system_service
ENV NAMESPACE=
ENV CONFIG_DIR=/blood_bank_system_service/core/golang/config
ENV ENV=dev
WORKDIR /blood_bank_system_service/zerotouch/golang

RUN go build -o main .

EXPOSE 6000
EXPOSE 6005
CMD ["/blood_bank_system_service/zerotouch/golang/main"]
