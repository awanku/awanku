require 'awanku/oauth/oauth'

class AuthController < ApplicationController

  def index
    # display available auth providers
  end

  def login
    redirect_to fetch_provider.authorization_url(request.base_url)
  end

  def callback
    code = params[:code]
    if code.blank?
      head :bad_request
      return
    end

    # fetch authenticated user profile form provider
    provider = fetch_provider
    creds = provider.exchange_code(request.base_url, code)
    user_profile = provider.fetch_user_profile(creds)

    # create new user if not exist
    user = find_or_create_user(user_profile[:name], user_profile[:email])

    cookies[:user_id] = user.id
    redirect_to root_path
  end

  private

  def fetch_provider
    provider = params.require(:provider).to_sym
    raise ActionController::RoutingError.new('unknown provider') unless Awanku::Oauth::PROVIDERS.include?(provider)

    Awanku::Oauth.for provider
  end

  def find_or_create_user(name, email)
    query = <<~SQL
      insert into users(name, email)
      values (:name, :email)
      on conflict (email)
      do update set updated_at = now()
      returning *
    SQL
    returned = User.find_by_sql [query, { :name => name, :email => email }]
    returned.first
  end

end
