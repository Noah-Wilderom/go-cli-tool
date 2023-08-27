# Stage 1: Base image for both development and build
FROM php:8.2-fpm AS base

# Install system dependencies
RUN apt-get update && \
    apt-get install -y zlib1g-dev libpng-dev libzip-dev && \
    rm -rf /var/lib/apt/lists/*

# Install PHP extensions
RUN docker-php-ext-install exif gd zip pdo_mysql

# Install and enable APCu extension
RUN pecl install apcu && \
    docker-php-ext-enable apcu

# Stage 2: Development stage
FROM base AS dev

# Copy necessary files for development
COPY composer.json composer.lock /var/www/html/
COPY app /var/www/html/app
COPY bootstrap /var/www/html/bootstrap
COPY config /var/www/html/config
COPY artisan /var/www/html/artisan

# Stage 3: Build stage
FROM base AS build-fpm

# Set working directory
WORKDIR /var/www/html

# Install Composer from composer:2.5 image
COPY --from=composer:2.5 /usr/bin/composer /usr/bin/composer

# Copy application files and install dependencies
COPY composer.json composer.lock /var/www/html/
RUN composer install --prefer-dist --no-ansi --no-dev --no-autoloader
COPY artisan /var/www/html/artisan
COPY bootstrap /var/www/html/bootstrap
COPY app /var/www/html/app
COPY config /var/www/html/config
COPY routes /var/www/html/routes

# Copy the rest of the application
COPY . /var/www/html/

# Generate optimized autoload files
RUN composer dump-autoload -o

# Stage 4: Final stage for running the application
FROM base AS fpm

# Copy application files from the build-fpm stage
COPY --from=build-fpm /var/www/html /var/www/html