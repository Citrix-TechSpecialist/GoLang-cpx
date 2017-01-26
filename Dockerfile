FROM kyma/docker-nginx
RUN apt-get -y update
RUN apt-get -y install wget
RUN mkdir /var/www ; cd /var/www ; wget --recursive --no-clobber --page-requisites --html-extension --convert-links --restrict-file-names=windows --random-wait --domains logic.bz --no-parent bt.logic.bz
RUN cp -R /var/www/bt.logic.bz/* /var/www/
RUN rm -R /var/www/bt.logic.bz
CMD 'nginx'
# Run the following command in the same directory as this file and hit your localhost's 8080 port to browse through the CPX blog.
# sudo docker build -t nginx-cpx-blog .
# sudo docker run -dt -p 8080:80 nginx-cpx-blog nginx
