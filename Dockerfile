FROM node:16-alpine3.15 as js-builder

WORKDIR /grafana

ENV NODE_ENV production

FROM golang:1.17.6-alpine3.15 as go-builder

RUN apk add --no-cache gcc g++ make

WORKDIR /grafana

COPY go.mod Makefile build.go ./

COPY --from=js-builder /grafana/public ./public

EXPOSE 3000

USER grafana
ENTRYPOINT [ "/run.sh" ]