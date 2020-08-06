# Build Stage
FROM lacion/alpine-golang-buildimage:1.13 AS build-stage

LABEL app="build-yaig"
LABEL REPO="https://github.com/leonkozlowski/yaig"

ENV PROJPATH=/go/src/github.com/leonkozlowski/yaig

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/leonkozlowski/yaig
WORKDIR /go/src/github.com/leonkozlowski/yaig

RUN make build-alpine

# Final Stage
FROM lacion/alpine-base-image:latest

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/leonkozlowski/yaig"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/yaig/bin

WORKDIR /opt/yaig/bin

COPY --from=build-stage /go/src/github.com/leonkozlowski/yaig/bin/yaig /opt/yaig/bin/
RUN chmod +x /opt/yaig/bin/yaig

# Create appuser
RUN adduser -D -g '' yaig
USER yaig

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/yaig/bin/yaig"]
