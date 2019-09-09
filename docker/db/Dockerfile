FROM mysql
ENV MYSQL_DATABASE dito
ENV MYSQL_USER root
ENV MYSQL_PASSWORD password  
ENV MYSQL_ROOT_PASSWORD password
COPY ./dito.sql /docker-entrypoint-initdb.d
#CMD  mysql -u -ppassword dito < /docker-entrypoint-initdb.d/dito.sql
EXPOSE 3306