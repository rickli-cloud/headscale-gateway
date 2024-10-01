# Develop

## Dev-container

1. Create headscale config

   ```sh
   docker volume create --driver local --opt type=none --opt o=bind --opt device=$PWD/.devcontainer/headscale headscale-config
   ```

2. Start dev-container

3. Start dev-server

   ```sh
   go run cmd/main.go
   ```
