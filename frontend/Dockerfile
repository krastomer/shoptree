FROM node:14-alpine3.14 AS builder

WORKDIR /app

COPY package.json .
RUN npm install
RUN npm install react-scripts -g

COPY . .

RUN npm run build

FROM nginx:1.21.3

COPY --from=builder /app/build /usr/share/nginx/html
WORKDIR /usr/share/nginx/html

CMD [ "nginx", "-g", "daemon off;" ]