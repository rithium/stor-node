FROM rithium/smartstackgo

WORKDIR $GOPATH/src/

COPY . $GOPATH/src/

RUN rm -rf /var/cache/apk/*

ADD nerve.conf.json /etc/nerve/
ADD nerve.conf.json /etc/

RUN chmod +x run.sh
RUN chmod +x /opt/startNerve.sh

ENTRYPOINT /go/src/run.sh