require 'json'
require 'faraday'
require 'octokit'

module Awanku
  module Oauth
    class GithubProvider
      SCOPE = ['read:user', 'user:email'].freeze

      def initialize(client_id, client_secret)
        @client_id = client_id
        @client_secret = client_secret
      end

      def authorization_url(_return_to)
        "https://github.com/login/oauth/authorize?client_id=#{@client_id}&scope=#{SCOPE.join(',')}"
      end

      def exchange_code(base_url, code)
        url = 'https://github.com/login/oauth/access_token'
        payload = {
          client_id: @client_id,
          client_secret: @client_secret,
          code: code
        }
        headers = {
          accept: 'application/json'
        }
        resp = Faraday.post(url, payload, headers)

        JSON.parse(resp.body)['access_token']
      end

      def fetch_user_profile(creds)
        client = Octokit::Client.new(access_token: creds)
        emails = client.emails
        primary_email = emails.select { |email| email[:primary] && email[:verified] }.first[:email]

        {
          name: client.user[:name],
          email: primary_email
        }
      end
    end
  end
end
