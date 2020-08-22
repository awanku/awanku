# Add your own tasks in files placed in lib/tasks ending in .rake,
# for example lib/tasks/capistrano.rake, and they will automatically be available to Rake.

require_relative 'config/application'

Rails.application.load_tasks

namespace :docker do
  desc 'Bring up dependencies'
  task :up do
    sh <<~BASH
      docker-compose build
      docker-compose up -d postgres
      docker-compose run app bundle install
      docker-compose run app yarn install --check-files
    BASH
  end

  desc 'Serve app with docker'
  task :serve do
    sh <<~BASH
      docker-compose run --service-ports app scripts/run_docker.sh
    BASH
  end

  desc 'Run database migration'
  task 'db-migrate' do
    sh <<~BASH

    BASH
  end

  desc 'Run rails console'
  task :console do
    sh <<~BASH
      docker-compose run --service-ports app bundle exec rails console
    BASH
  end

  desc 'Run test with docker'
  task :test do
    sh <<~BASH
      docker-compose run app bundle exec rspec
    BASH
  end
end
