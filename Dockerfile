FROM golang:1.12

RUN go get github.com/oxequa/realize
ENTRYPOINT ["realize"]
CMD ["start"]
