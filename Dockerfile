FROM golang:latest AS build

WORKDIR /work

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build $PWD/cmd/gateway.go


FROM scratch

LABEL maintainer=github/rickli-cloud

# ca-certificates are required to reach out to the OAuth Endpoint. 
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=build /work/gateway /hsgw

EXPOSE 8000

ENTRYPOINT [ "/hsgw serve" ]
