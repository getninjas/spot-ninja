# ------------------------------------------------------------------------------
# configure env and workdir - golang:1.12.7
# ------------------------------------------------------------------------------
FROM golang:1.12.7 AS pre-build
LABEL SQUAD="getninjas"
RUN useradd spot-ninja
ENV APP /home/spot-ninja
ENV FLEETIGNORED "fleet1,fleet2"
ENV SQSURL "https://sqs.example.com"
WORKDIR ${APP}/src/spot-ninja
COPY . ${APP}/src/spot-ninja
# ------------------------------------------------------------------------------
# test - test and conver
# ------------------------------------------------------------------------------
FROM pre-build AS test
LABEL SQUAD="getninjas"
RUN go mod download
RUN go test -cover ./config/ \
    && go test -cover ./pkg/logic/ \
    && go test -cover ./pkg/structure/ \
    && go test -cover ./pkg/api/ \
    && go vet ./pkg/api/ \
    && go vet ./cmd/spot-ninja/ \
    && go vet ./config/ \
    && go vet ./pkg/logic/ \
    && go vet ./pkg/structure/ \
    && go vet ./pkg/api/ 
# ------------------------------------------------------------------------------
# builder - build the binary 
# ------------------------------------------------------------------------------
FROM test AS builder
LABEL SQUAD="getninjas"
ENV APP /home/spot-ninja
RUN CGO_ENABLED=0 GOOS=linux \
    go build -o ${APP}/spot-ninja ./cmd/spot-ninja/main.go \
    && chmod +x ${APP}/spot-ninja \
    && echo "nobody:x:65534:65534:Nobody:/:" > ${APP}/etc_passwd
## ------------------------------------------------------------------------------
## runner - daemon image
## ------------------------------------------------------------------------------
FROM scratch AS runner
ENV APP /home/spot-ninja
COPY --from=builder /etc/ssl /etc/ssl
COPY --from=builder ${APP}/etc_passwd /etc/passwd
COPY --from=builder ${APP}/spot-ninja /spot-ninja
USER nobody
CMD ["/spot-ninja"]
