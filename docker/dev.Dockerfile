FROM ruby:2.7-alpine
RUN apk add --no-cache bash
RUN apk add --no-cache postgresql-dev postgresql-client
RUN apk add --no-cache build-base
RUN apk add --no-cache nodejs yarn
