FROM kyma/docker-nginx
MAINTAINER Mayank Tahilramani and Brian Tannous
RUN mkdir /var/www
COPY ./cpx-blog /var/www
CMD 'nginx'