## Running the Application
- Make sure that your machine have been installed `docker & docker compose`
- I include some of the following variables as config in environment variables so you can easily change them depending on your intended use. Change if in file .env
```
PORTS = 80
BaseAPIURL = https://mfx-recruit-dev.herokuapp.com
SWAGGER_DOMAIN = 127.0.0.1
```
- `PORT` is port that the web application will run
- `BaseAPIURL` is the domain of the third party, I put it in env because it can change depending on the environment, for example the domain of DEV/QA/PROD
- `SWAGGER_DOMAIN` is the domain of swagger we will using for testing api

To run the heroku, execute the following command:
```sh
git clone https://github.com/pdhoang91/heroku.git
`docker-compose up` or `docker-compose --env-file ./.env up`
```
# Test application 
- Using swagger to test endpoints API. You can access `http://127.0.0.1:{PORT}/swagger/index.html`
- Using postman to testing: ```curl --location 'http://localhost:80/v1/users/1'```
- We provide an API get all user information for user admin. To call this api you need a basic token : `Basic cGRob2FuZzkxQGdtYWlsLmNvbTptb25leV9mb3J3YXJkX3ZpZXRuYW0=`
- Using postman to testing: ```curl --location 'http://localhost:80/v1/admin/users' \--header 'Content-Type: application/json' \--header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTptb25leV9mb3J3YXJkX3ZpZXRuYW0='```

# Run Unittest
- Please stand at project root folder and run `go test ./...` or `go test --cover ./...`