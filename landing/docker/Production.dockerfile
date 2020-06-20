FROM node:12.18
WORKDIR /app/landing
COPY package.json .
COPY yarn.lock .
RUN yarn install
COPY . .
RUN yarn build
CMD yarn start
