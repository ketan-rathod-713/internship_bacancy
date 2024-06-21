FROM node:14-alpine

# install nodemon globally inside cotnainer
# because in future we want to use it.
RUN npm install -g nodemon

# usr folder is safe space to put our application
WORKDIR /usr/app

# first copy package json and install all depencdencies
COPY ./package.json ./
RUN npm install

# Now add other changes that occured.
COPY ./ ./

# default command
CMD ["npm", "run", "dev"]
