FROM danlynn/ember-cli:3.25.2

WORKDIR /app

COPY package.json ./
COPY package-lock.json ./

RUN npm install

# Copy all our app's directories
COPY . .

EXPOSE 4200 7357 9222
