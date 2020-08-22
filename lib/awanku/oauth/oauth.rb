require 'awanku/oauth/github'
require 'awanku/oauth/google'

module Awanku
  module Oauth
    PROVIDER_GITHUB = :github
    PROVIDER_GOOGLE = :google
    PROVIDERS = [PROVIDER_GITHUB, PROVIDER_GOOGLE]

    @@cache = {}

    def self.for(provider)
      @@cache[provider] ||= self.create(provider)
      @@cache[provider]
    end

    def self.create(provider)
      case provider
      when PROVIDER_GITHUB
        client_id = Rails.application.config_for(:oauth)[:github][:client_id]
        client_secret = Rails.application.config_for(:oauth)[:github][:client_secret]
        GithubProvider.new client_id, client_secret
      when PROVIDER_GOOGLE
        client_id = Rails.application.config_for(:oauth)[:google][:client_id]
        client_secret = Rails.application.config_for(:oauth)[:google][:client_secret]
        GoogleProvider.new client_id, client_secret
      else
        raise ArgumentError.new "unknown provider: #{provider.to_s}"
      end
    end
  end
end
