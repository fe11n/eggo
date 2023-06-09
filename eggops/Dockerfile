# Build the manager binary
FROM golang:1.16 as builder

ARG ARCH

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN GOPROXY="https://goproxy.cn,direct" go mod download

# Copy the go source
COPY eggops/main.go eggops/main.go
COPY eggops/api/ eggops/api/
COPY eggops/controllers/ eggops/controllers/
COPY cmd/ cmd/
COPY pkg/ pkg/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${ARCH} go build -a -o manager eggops/main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/manager .
USER 65532:65532

ENTRYPOINT ["/manager"]
