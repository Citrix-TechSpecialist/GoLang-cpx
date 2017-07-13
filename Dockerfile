FROM fnichol/uhttpd
MAINTAINER Mayank Tahilramani and Brian Tannous
COPY ./cpx-blog /www
WORKDIR /www 
RUN find ./ -type f -exec sed -i "s/All rights reserved./Hosted by container: ${HOSTNAME}/g" {} \;
EXPOSE 80
ENTRYPOINT ["/usr/sbin/run_uhttpd", "-f", "-p", "80", "-h", "/www"]
CMD [""]
