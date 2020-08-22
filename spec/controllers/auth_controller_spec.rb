require 'rails_helper'
require 'awanku/oauth/oauth'

RSpec.describe AuthController do

  describe "GET #index" do
    it "render index view" do
      get :index
      expect(response).to have_http_status(:ok)
      expect(response).to render_template("index")
    end
  end

  describe "GET #login" do
    context "unknown provider" do
      it "returns 404" do
        get :login, :params => { :provider => 'unknown' }
        expect(response).to have_http_status(:not_found)
      end
    end

    context "google provider" do
      it "returns redirect" do
        redirect_url = 'http://dummy.com'

        provider = Awanku::Oauth::GoogleProvider.new 'client_id', 'client_secret'
        allow(provider).to receive(:authorization_url).with('http://test.host').and_return(redirect_url)
        allow(Awanku::Oauth).to receive(:for).and_return(provider)

        get :login, :params => { :provider => Awanku::Oauth::PROVIDER_GOOGLE }
        expect(response).to redirect_to(redirect_url)
      end
    end

    context "github provider" do
      it "returns redirect" do
        redirect_url = 'http://dummy.com'

        provider = Awanku::Oauth::GithubProvider.new 'client_id', 'client_secret'
        allow(provider).to receive(:authorization_url).with('http://test.host').and_return(redirect_url)
        allow(Awanku::Oauth).to receive(:for).and_return(provider)

        get :login, :params => { :provider => Awanku::Oauth::PROVIDER_GITHUB }
        expect(response).to redirect_to(redirect_url)
      end
    end
  end

  describe "GET #callback" do
    it "requires code in param" do
      get :callback, :params => { :provider => Awanku::Oauth::PROVIDER_GITHUB, :code => nil }
      expect(response).to have_http_status(:bad_request)
    end

    context "github provider" do
      it "works" do
        code = 'somecode'
        creds = 'creds'
        user_profile = { name: 'somename', email: 'someemail@gmail.com' }

        provider = Awanku::Oauth::GithubProvider.new 'client_id', 'client_secret'
        allow(provider).to receive(:exchange_code).with('http://test.host', code).and_return(creds)
        allow(provider).to receive(:fetch_user_profile).with(creds).and_return(user_profile)
        allow(Awanku::Oauth).to receive(:for).and_return(provider)

        get :callback, :params => { :provider => Awanku::Oauth::PROVIDER_GITHUB, :code => code }
        expect(response).to redirect_to(root_path)
        expect(cookies[:user_id].to_i).to be > 0
      end
    end
  end

end
