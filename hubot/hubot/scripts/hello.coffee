# Description:
#   Hello Hubot

module.exports = (robot) ->
 robot.respond /hello/i, (res) ->
  res.send "hubot -> hello"
