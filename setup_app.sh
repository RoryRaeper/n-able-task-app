go mod tidy
docker-compose up --build -d

cd frontend

npm install
npm start &