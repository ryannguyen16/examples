require 'easypost'

EasyPost.api_key = ENV['EASYPOST_API_KEY']

user = EasyPost::User.retrieve('user_...')

user.recharge_threshold = '50.00'

user.save

puts user
