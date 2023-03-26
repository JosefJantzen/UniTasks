# UniTasks
The project is still be developed. If I'm finished there will be a  more detailed documentation.

## Installation

1. Create a new directory for example uni-tasks
2. Create a .env file and define the following attributes:
```env
BACKEND_CONFIG="/home/<user>/uni-tasks/config.json"

DB_DATABASE="unitasks"
DB_CERTS_DIR="/home/<user>/uni-tasks/certs"

SERVER_ALIASES="localhost 127.0.0.1 db"
```

3. Optional: Create your own config.json file
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

8. You should now find the api at `localhost:8081` and the database dashboard at `localhost:8180`