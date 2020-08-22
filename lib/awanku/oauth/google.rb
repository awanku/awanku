require 'googleauth'
require 'googleauth/user_authorizer'
require 'google/apis/oauth2_v2'

module Awanku
  module Oauth
    class GoogleProvider
      SCOPE = [
        'https://www.googleapis.com/auth/userinfo.email',
        'https://www.googleapis.com/auth/userinfo.profile'
      ].freeze

      def initialize(client_id, client_secret)
        google_client_id = Google::Auth::ClientId.new(client_id, client_secret)
        @authorizer = Google::Auth::UserAuthorizer.new(
          google_client_id,
          SCOPE,
          nil, # we don't need to store token
          '/auth/google/callback'
        )
      end

      def authorization_url(base_url)
        @authorizer.get_authorization_url(base_url: base_url)
      end

      def exchange_code(base_url, code)
        @authorizer.get_credentials_from_code(code: code, base_url: base_url)
      end

      def fetch_user_profile(creds)
        service = Google::Apis::Oauth2V2::Oauth2Service.new
        user_info = service.get_userinfo_v2(options: {authorization: creds})

        {
          name: user_info.name,
          email: user_info.email
        }
      end
    end
  end
end
