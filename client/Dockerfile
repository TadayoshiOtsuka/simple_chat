FROM node:18 as Dev

WORKDIR /usr/app

COPY package*.json /usr/app/
COPY tsconfig*.json /usr/app/
COPY vite.config.ts /usr/app/
COPY index.html /usr/app/
RUN npm i

EXPOSE 5173
CMD ["npm", "run", "dev" ]
