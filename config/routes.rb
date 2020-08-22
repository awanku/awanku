# For details on the DSL available within this file, see https://guides.rubyonrails.org/routing.html
Rails.application.routes.draw do
  root to: 'landing#index'

  scope '/auth' do
    get '/', to: 'auth#index'
    get '/:provider/login', to: 'auth#login'
    get '/:provider/callback', to: 'auth#callback'
  end

  namespace :console do
  end
end
