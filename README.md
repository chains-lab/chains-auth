after add to ```"scripts"``` - ```"start": "http-server -c-1 -p 7272"```
### Start docs server
```bash
  cd docs 
  npm run start
```

### Build and Start project
```bash
docker compose chains-auth up -d

# Down project 
docker compose chains-auth down -v
```

### Connecting to network in Docker
```bash
# Create a Docker network named `chains-net` if it does not already exist.
docker network create chains-net
# or connect to existing network
docker network connect chains-net chains-auth
```
