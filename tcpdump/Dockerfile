FROM ubuntu:14.04

#Install nginx
RUN apt-get update && apt-get -y install nginx curl tcpdump
RUN echo "daemon off;" >> /etc/nginx/nginx.conf

#Install Supervisor and config
RUN apt-get install -y supervisor
RUN touch /etc/supervisord.conf
RUN echo '[supervisord]'  >> /etc/supervisord.conf
RUN echo 'nodaemon=true'  >> /etc/supervisord.conf
RUN echo '[program:nginx]' >> /etc/supervisord.conf
RUN echo 'command=nginx'   >> /etc/supervisord.conf
RUN echo 'stdout_logfile=/dev/fd/1'   >> /etc/supervisord.conf
RUN echo 'stdout_logfile_maxbytes=0'  >> /etc/supervisord.conf
RUN echo '[program:tcpdump]' >> /etc/supervisord.conf
RUN echo 'command=tcpdump -nn'   >> /etc/supervisord.conf
RUN echo 'stdout_logfile=/dev/fd/1'   >> /etc/supervisord.conf
RUN echo 'stdout_logfile_maxbytes=0'  >> /etc/supervisord.conf

EXPOSE 80
CMD /usr/bin/supervisord -c /etc/supervisord.conf

