FROM mysql:8

MAINTAINER slinso

RUN chown -R mysql:root /var/lib/mysql/

ENV MYSQL_DATABASE=dvds
ENV MYSQL_USER=root
ENV MYSQL_PASSWORD=secret
ENV MYSQL_ROOT_PASSWORD=root

ADD data.sql /tmp/data.sql

RUN cp /tmp/data.sql /docker-entrypoint-initdb.d

EXPOSE 3306