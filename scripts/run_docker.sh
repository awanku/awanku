#!/usr/bin/env bash
yarn install --check-files

bundle exec rails db:create db:migrate
bundle exec rails server --binding=0.0.0.0
