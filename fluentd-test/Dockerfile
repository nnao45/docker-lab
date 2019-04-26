FROM fluent/fluentd:v0.12

WORKDIR /

RUN mkdir -p /var/log/mill && \
    mkdir -p /var/log/td-agent/position && \
    chmod 755 /var/log -R

COPY fluentd.conf /fluentd/etc/fluentd.conf

ENV FLUENTD_CONF fluentd.conf
