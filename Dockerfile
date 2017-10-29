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

FROM golang:1.9-alpine AS go
COPY . /go/src/website
RUN go install website

# Build final image

FROM alpine:3.6
# Add ssl certs for Go
RUN apk add --no-cache ca-certificates curl
COPY hots.json /
ENV GOOGLE_APPLICATION_CREDENTIALS /hots.json
COPY --from=static /frontend/build /static
COPY --from=go /go/bin/website /website
ENTRYPOINT ["/website"]
