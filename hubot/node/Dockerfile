FROM node:10.16.1

# hubot
RUN npm install -g hubot yo generator-hubot coffee-script

RUN useradd hubot -m
USER hubot
WORKDIR /home/hubot

