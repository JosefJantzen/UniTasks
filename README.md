# UniTasks
UniTasks is a todo web app specifically for your university todos. But it's not limited to that. <br>
The backend is written in [Go](https://go.dev/) and the frontend in [Vue.js 3](https://vuejs.org/). The databsae is a [CockroachDB](https://www.cockroachlabs.com/). <br>
You can test UniTasks [here](https://unitasks.josefjantzen.de).

## Installation

1. Create a new directory for example uni-tasks
2. Create a .env file and define the following attributes:
```env
BACKEND_CONFIG="/home/<user>/uni-tasks/config.json" # Full path to your config.json

DB_DATABASE="unitasks" # The name of the database that is created if not existing.
DB_CERTS_DIR="/home/<user>/uni-tasks/certs" # Full path to your generated certificates

SERVER_ALIASES="localhost 127.0.0.1 db" # All the possible host names of your database. These are used to generate the certificates. 
```

3. Create your own config.json file as described [here](#config)
4. Run the following command to install the server. ***Note:*** You have to install [cockraoch](https://www.cockroachlabs.com/docs/v22.2/install-cockroachdb-linux.html) before that and add it to the PATH
```bash
curl -s https://raw.githubusercontent.com/JosefJantzen/UniTasks/main/gen-certs.sh -O ; chmod +x gen-certs.sh ; ./gen-certs.sh ; curl -s -o docker-compose.yml https://raw.githubusercontent.com/JosefJantzen/UniTasks/main/docker-compose.prod.yml ; sudo docker compose up -d
```

5. Connect to your databse instance
```bash
docker exec -it uni-tasks-db ./cockroach sql --certs-dir=/certs
```

6. Create a user with password for your database and grant him access to the database:
```sql
CREATE DATABASE unitasks;
CREATE USER api WITH PASSWORD '<password>';
GRANT ALL ON DATABASE unitasks TO api;
```

7. Restart your rest server to connect again to the database
```bash
sudo docker restart uni-tasks-rest-server
```

8. You should now find the frontend at `localhost:8082`, the api at `localhost:8081` and the database dashboard at `localhost:8180`.

## Config
If you don't have a config file or it doesn't contains all the attributes the default values are defined in [here](https://github.com/JosefJantzen/UniTasks/blob/main/backend/config.sample.json).
Your `config.json` should look like this:
```json
{
    "port": "8080",
    "jwtKey": "SecretYouShouldHide",
    "jwtExpireMin": 5,
    "frontendUrl": "http://localhost:8080"
    "DB": {
        "user": "totoro",
        "pwd": "whatever",
        "host": "db",
        "port": "26257",
        "database": "unitasks",
        "initial": "DB-initial.pgsql",
        "testData": "DB-test-data.pgsql"
    },
}
```
| Attribute | Description | Default value |
| ---       | ----------- | ------------- |
| port      | Port of the backend. You don't need to change it normaly because it's only the port inside the docker container and the outside port is configured in the docker file anyways. | 8080 |
| jwtKey    | The secret key used to generate the json web tokens. | SecretYouShouldHide |
| jwtExpireMin | The time in minutes after the generated jwt will expire. You should choose a relatively small time here. | 5 |
| frontendUrl | The url your frontend has. It's used in the backend for the CORS-Policy. | http://localhost:8080 |

### DB config

| Attribute | Description | Default value |
| ---       | ----------- | ------------- |
| user      | The user you created in your database for the api. | totoro |
| pwd       | The password of the databse user. This is only in production mode necessary with a secured database. | whatever | 
| host | The address to the database instance. You can use here the internal hostnames of the docker containers. | db |
| port | The port of the database instance. | 26257 |
| database | The name of the database that should be used. | unitasks | 
| initial | A  sql file that is executed on startup of the rest server | DB-initial.pgsql |
| testData | A sql file that is executed after the initial file if the enviroment variable `DEBUG` is set to true. | DB-test-data.pgsql |
