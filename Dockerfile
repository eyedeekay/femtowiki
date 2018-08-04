FROM alpine:3.8
ENV args=args
# Setup distro and user
RUN apk update && apk upgrade
RUN apk add go git ca-certificates musl-dev
RUN mkdir -p /opt/femtowiki
RUN adduser -h /opt/femtowiki -g 'femtowiki,,,,' -D femtowiki

# Build femtowiki from source
COPY . /usr/src/femtowiki
WORKDIR /usr/src/femtowiki
RUN chown femtowiki -R /opt/femtowiki/
RUN go get -u github.com/s-gv/femtowiki/
RUN go get -u github.com/eyedeekay/sam-forwarder
RUN go build
RUN cp femtowiki /usr/bin/femtowiki

# Cleanup build and dependencies
RUN apk del go git musl-dev

# Setup and run femtowiki
USER femtowiki
WORKDIR /opt/femtowiki
RUN femtowiki -migrate
VOLUME /opt/femtowiki
CMD femtowiki -createsuperuser && \
    femtowiki $args
