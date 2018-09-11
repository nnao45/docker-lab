FROM alpine

RUN apk add --update python3 python3-dev git curl
RUN curl https://raw.githubusercontent.com/pypa/get-pip/master/get-pip.py | python3
#RUN curl https://bootstrap.pypa.io/get-pip.py | python3
#RUN pip install virtualenv
#RUN virtualenv venv
#RUN source venv/bin/activate
RUN pip install --force-reinstall pip==9.0.3
RUN pip install --upgrade setuptools
RUN pip install ansible-container[docker,k8s]

CMD ["sh"]
