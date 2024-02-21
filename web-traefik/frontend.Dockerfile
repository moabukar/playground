FROM node:16.10-alpine AS build

WORKDIR /frontend

COPY frontend/package.json frontend/package-lock.json ./

RUN npm install

COPY frontend .

RUN npm run build

### STAGE 2: Run ###
FROM nginx:1.17.1-alpine

COPY nginx/nginx.conf /etc/nginx/nginx.conf

COPY --from=build /frontend/dist /usr/share/nginx/html

#COPY --from=build /frontend /copiedfrontend
