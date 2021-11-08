FROM node:12.18.2-alpine

WORKDIR /app

ENV PATH /app/node_modules/.bin:$PATH
ENV REACT_APP_API_URL $REACT_APP_API_URL

COPY package.json ./
COPY yarn.lock ./

RUN yarn install
RUN yarn global add react-scripts@3.4.1

COPY . ./

EXPOSE 8002
CMD ["yarn", "start"]
