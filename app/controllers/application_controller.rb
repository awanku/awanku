class ApplicationController < ActionController::Base
  rescue_from ActionController::RoutingError, :with => :handle_not_found

  private

  def handle_not_found
    head :not_found
  end
end
