FROM mysql
ENV MYSQL_DATABASE dito
ENV MYSQL_PASSWORD dito  
ENV MYSQL_ROOT_PASSWORD dito
COPY ./dito.sql /docker-entrypoint-initdb.d
EXPOSE 3306