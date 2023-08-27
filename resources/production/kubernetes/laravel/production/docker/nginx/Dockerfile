FROM node:20-alpine AS assets-build

WORKDIR /var/www/html

COPY . /var/www/html/

RUN npm ci
RUN npm run build

FROM nginx:1.25.2-alpine AS nginx

COPY /production/docker/nginx/vhost.conf /etc/nginx/conf.d/default.conf
COPY --from=assets-build /var/www/html/public /var/www/html/