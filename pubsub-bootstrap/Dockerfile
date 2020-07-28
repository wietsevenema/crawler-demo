FROM node:14
ENV NODE_ENV=production

RUN apt-get update && apt-get install netcat -y

WORKDIR /app
COPY package.json .
COPY package-lock.json .
RUN npm install

COPY . . 
