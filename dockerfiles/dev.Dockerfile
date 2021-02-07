# this is a dockerfile only for local development - do not use in production
FROM golang:1.15
RUN go get -v github.com/githubnemo/CompileDaemon
ENTRYPOINT CompileDaemon -directory=. -log-prefix=false -build="go build ./cmd/share-secret/" -command="./share-secret" -exclude-dir=.git -exclude-dir=./integration