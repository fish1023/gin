FROM centos:7.4.1708

RUN mkdir -p /data/log/

WORKDIR /data/www/captcha

COPY ./captcha ./
COPY ./conf  ./conf/

EXPOSE 80

ENTRYPOINT ["./captcha"]
