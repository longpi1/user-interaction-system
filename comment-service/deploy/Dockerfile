### build Artalk
ENV HOMEDIR /home/work/comment-service
ENV SERVICENAME comment-service
FROM golang:1.21.3-alpine3.18 as builder

WORKDIR ${HOMEDIR}
ENV CC=gcc
COPY . .
RUN apk add --no-cache gcc musl-dev \
    && go build ../ && mv ${SERVICENAME} /go/bin
###############
FROM alpine:3.6
COPY --from=builder /go/bin/${SERVICENAME} /usr/local/bin
COPY --from=builder  ${HOMEDIR} /${SERVICENAME}
COPY ../conf/online/web.yaml  ${HOMEDIR}/conf/web.yaml
WORKDIR /${SERVICENAME}
CMD "${SERVICENAME}"

