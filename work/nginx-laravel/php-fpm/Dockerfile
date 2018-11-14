#FROM php:7-fpm-alpine
FROM php:latest

RUN apt-get update \
    && apt-get install -y mysql-client \
    && apt-get install -y zip unzip git \
    && apt-get install -y vim

RUN apt-get update \
	&& apt-get install -y libmcrypt-dev \
	&& rm -rf /var/lib/apt/lists/* \
	&& pecl install mcrypt-1.0.1 \
	&& docker-php-ext-enable mcrypt

RUN git clone https://github.com/phpredis/phpredis.git /usr/src/php/ext/redis 

RUN docker-php-ext-install pdo_mysql session json mbstring redis

RUN curl -sS https://getcomposer.org/installer | php \
    && mv composer.phar /usr/local/bin/composer

#RUN mkdir -p /usr/share/nginx/app \
#    && cd /usr/share/nginx/app \
#    && composer create-project laravel/laravel --prefer-dist sample 5.6

WORKDIR /usr/share/nginx/app
