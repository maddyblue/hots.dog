# Install and run yarn

FROM gcr.io/google-appengine/debian8 AS static

# Install updates and dependencies
RUN apt-get update -y && apt-get install --no-install-recommends -y -q curl python build-essential git ca-certificates libkrb5-dev imagemagick && \
    apt-get clean && rm /var/lib/apt/lists/*_*

# Install the latest LTS release of nodejs
RUN mkdir /nodejs && curl https://nodejs.org/dist/v6.11.3/node-v6.11.3-linux-x64.tar.gz | tar xvzf - -C /nodejs --strip-components=1
ENV PATH $PATH:/nodejs/bin

# Install the latest stable release of Yarn
RUN mkdir /yarn && curl -L https://github.com/yarnpkg/yarn/releases/download/v1.0.2/yarn-v1.0.2.tar.gz | tar xvzf - -C /yarn --strip-components=1
ENV PATH $PATH:/yarn/bin

COPY /frontend /frontend

RUN cd frontend && yarn
RUN cd frontend && yarn run build

# Build Go app, install cockroach

FROM golang:alpine AS go

RUN apk add --update curl && \
	rm -rf /var/cache/apk/*

RUN \
	VERSION="v1.1-beta.20170921" && \
	NAME="cockroach-${VERSION}.linux-musl-amd64" && \
	curl https://s3.amazonaws.com/binaries-test.cockroachdb.com/${NAME}.tgz | \
	tar xzv && \
	mv ${NAME}/cockroach /cockroach

COPY . /go/src/website
RUN go install website

# Build final image

FROM alpine
COPY --from=static /frontend/build /static
COPY --from=go /go/bin/website /website
COPY --from=go /cockroach /cockroach
ENTRYPOINT ["/website"]
CMD [ \
	"-exec", "/cockroach start --background --insecure", \
	"-autocert", "hots.mattjibson.com" \
]
