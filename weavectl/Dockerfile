FROM alpine

RUN apk --update add curl
RUN curl -L git.io/weave -o /usr/local/bin/weave && \
    chmod a+x /usr/local/bin/weave
CMD ["/bin/sh"]
