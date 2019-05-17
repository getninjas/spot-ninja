# ------------------------------------------------------------------------------
# configure env and workdir - golang:1.12.3
# ------------------------------------------------------------------------------
FROM golang@sha256:55f89a93dde69671d902f5c205ff49299cb4dfa3480961773dd8f61696a3aa02 AS pre-build
LABEL SQUAD="getninjas"
RUN useradd spot-advisor
ENV APP /home/spot-advisor
ENV FLEETIGNORED "fleet1,fleet2"
ENV SQSURL "https://sqs.example.com"
WORKDIR ${APP}/src/spot-advisor
COPY . ${APP}/src/spot-advisor
# ------------------------------------------------------------------------------
# test - test and conver
# ------------------------------------------------------------------------------
FROM pre-build AS test
LABEL SQUAD="getninjas"
RUN go mod download
RUN go test -cover ./config/ \
    && go test -cover ./pkg/logic/ \
    && go test -cover ./pkg/structure/ \
    && go vet ./pkg/api/ \
    && go vet ./cmd/spot-advisor/ \
    && go vet ./config/ \
    && go vet ./pkg/logic/ \
    && go vet ./pkg/structure/
# ------------------------------------------------------------------------------
# builder - build the binary 
# ------------------------------------------------------------------------------
FROM test AS builder
LABEL SQUAD="getninjas"
ENV APP /home/spot-advisor
RUN CGO_ENABLED=0 GOOS=linux \
    go build -o ${APP}/spot-advisor ./cmd/spot-advisor/main.go \
    && chmod +x ${APP}/spot-advisor \
    && echo "nobody:x:65534:65534:Nobody:/:" > ${APP}/etc_passwd
## ------------------------------------------------------------------------------
## runner - daemon image
## ------------------------------------------------------------------------------
FROM scratch AS runner
ENV APP /home/spot-advisor
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder ${APP}/etc_passwd /etc/passwd
COPY --from=builder ${APP}/spot-advisor /spot-advisor
USER nobody
CMD ["/spot-advisor"]
