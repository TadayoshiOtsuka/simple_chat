# ==========
# Dev
# ==========
FROM node:18 as Dev

WORKDIR /usr/app

COPY package*.json /usr/app/
COPY tsconfig*.json /usr/app/
COPY vite.config.ts /usr/app/
COPY index.html /usr/app/
RUN npm i

EXPOSE 5173
CMD ["npm", "run", "dev" ]


# ==========
# Builder
# ==========
FROM node:18-alpine as Builder

WORKDIR /usr/app/src

COPY . .

RUN npm i --only-production && \
    npm run build


#============
# Deploy
#============
FROM nginx:alpine

WORKDIR /usr/share/nginx/html

COPY --from=Builder /usr/app/src/dist /var/www
COPY ./nginx /etc/nginx/conf.d/

CMD ["nginx", "-g", "daemon off;"]
