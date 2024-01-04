FROM nginx

#RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf
#COPY content /usr/share/nginx/html
COPY conf /etc/nginx

EXPOSE 80

STOPSIGNAL SIGQUIT

ENTRYPOINT ["/start.sh"]
