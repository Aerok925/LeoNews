FROM golang:1.21.4 as builder

ENV CGO_ENABLED=0
ENV GO111MODULE=on

WORKDIR /go/src/app
COPY . .
RUN go build -mod vendor -o /bin/betting.api.golang .
RUN go build -mod vendor -o /bin/migration_refresh scripts/db/migration/refresh/refresh.go
RUN go build -mod vendor -o /bin/migration_roll scripts/db/migration/roll/roll.go
RUN go build -mod vendor -o /bin/migration_up scripts/db/migration/up/up.go



FROM golang:1.21.4-alpine
RUN apk update && apk add --no-cache git
COPY --from=builder /bin/betting.api.golang /bin/betting.api.golang
COPY --from=builder /bin/migration_refresh /bin/migration_refresh
COPY --from=builder /bin/migration_roll /bin/migration_roll
COPY --from=builder /bin/migration_up /bin/migration_up
CMD ["betting.api.golang"]

