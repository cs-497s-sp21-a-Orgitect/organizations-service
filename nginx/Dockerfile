FROM nginx

COPY public /usr/share/nginx/html
RUN chown -R nginx:nginx /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf