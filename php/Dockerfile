FROM php:7.4-apache
WORKDIR /var/www/html
COPY --from=composer /usr/bin/composer /usr/bin/composer
COPY . .

RUN apt-get update \
    && apt-get install -y zlib1g-dev libzip-dev \
        git \
    && pecl install grpc \
    && docker-php-ext-install zip \
    && docker-php-ext-enable grpc

RUN composer install
RUN mv "$PHP_INI_DIR/php.ini-development" "$PHP_INI_DIR/php.ini"
