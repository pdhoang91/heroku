## Running the Application
- Ensure that `docker & docker compose` compose are installed on your workstation.
- Some of the variables given below are set up as environment variables, so you may simply change them to suit your needs, modifying the .env file
```
PORTS = 80
BaseAPIURL = https://mfx-recruit-dev.herokuapp.com
SWAGGER_DOMAIN = 127.0.0.1
```
- `PORT` is the port on which the web application will execute
- `BaseAPIURL` is the domain of a third party, and it may vary based on the environment you are utilizing.
- `SWAGGER_DOMAIN` is the swagger domain that will be used for API testing.

To run the heroku, execute the following command:
```sh
git clone https://github.com/pdhoang91/heroku.git
`docker-compose up` or `docker-compose --env-file ./.env up`
```
# Testing API Endpoints with Swagger and Postman
We offer both public and private APIs. Please use the following token to access the secret API:
`Basic cGRob2FuZzkxQGdtYWlsLmNvbTptb25leV9mb3J3YXJkX3ZpZXRuYW0=`

##### Using Swagger
Follow these steps to test API endpoints using Swagger:
Swagger documentation may be found at `http://127.0.0.1:{PORT}/swagger/index.html`.

#####  Using Postman
To test API endpoints using Postman, follow these steps:

1. This API returns the name, account list, and balance of specific user’s:

   ```shell
   curl --location 'http://localhost:80/v1/users/1'
   
2. This API only the user admin may access, and it provides the name, account list, and balance of all users

   ```shell
    curl --location 'http://localhost:80/v1/admin/users' \
    --header 'Content-Type: application/json' \
    --header 'Authorization: Basic cGRob2FuZzkxQGdtYWlsLmNvbTptb25leV9mb3J3YXJkX3ZpZXRuYW0='
   
# Run Unittest
- Please stand at project root folder and run `go test ./...` or `go test --cover ./...`