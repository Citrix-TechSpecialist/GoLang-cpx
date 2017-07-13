FROM fnichol/uhttpd
MAINTAINER Mayank Tahilramani and Brian Tannous
COPY ./cpx-blog /www
RUN echo "find /www -type f -exec sed -i \"s/All rights reserved./Hosted by container: ${HOSTNAME}/g\" {} \\;" > /tmp/update.sh && chmod +x /tmp/update.sh
EXPOSE 80
ENTRYPOINT /tmp/update.sh && /usr/sbin/run_uhttpd -f -p 80 -h /www
CMD [""]
