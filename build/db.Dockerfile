FROM ubuntu
ENV TZ=Europe/Moscow
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

ENV DEBIAN_FRONTEND=noninteractive
ENV POSTGRES_HOST /var/run/postgresql/
ENV POSTGRES_PORT 5432
ENV POSTGRES_DB cw
ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD password

RUN apt-get update && apt-get install -y postgresql


USER postgres

COPY build/scripts.sql scripts.sql

RUN service postgresql start &&\
    psql -U postgres -c "ALTER USER postgres PASSWORD 'password';" &&\
    psql -U postgres -c 'CREATE DATABASE "cw";' &&\
    psql -U postgres -d cw -a -f scripts.sql &&\
    service postgresql stop

EXPOSE 5432
CMD service postgresql start
