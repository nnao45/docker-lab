FROM centos

ADD .zshrc /root

RUN yum -y install zsh

RUN usermod -s /bin/zsh root

CMD ["zsh"]
