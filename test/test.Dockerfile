FROM medtune/beta-platform:build

RUN go test -tags=cicd -v ./pkg/...

RUN go test -tags=cicd -v ./cmd/...
