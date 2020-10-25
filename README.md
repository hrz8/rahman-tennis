# Rahman Tennis (Try Go 1.15)

## Quick Start

### Requirements

```bash
Go >= 1.15
MySQL
```

### Usage

``` yaml
# Setting up Apps and DB configurations in `app.config.dev.yml` file
APP_PORT: 8080
DB_HOST: somehost
DB_PORT: 3306
DB_USER: someuser
DB_PASSWORD: somepassword
DB_NAME: some_name
```

``` bash
# Create some_name db
msql> CREATE DATABASE some_name;
```

``` bash
# Run apps with migration and default seed data
$ go run main.go migrate

# Run apps only
$ go run main.go
```

### Available Endpoints

Apps run in: `HOST:APP_PORT`

1. `GET /api/v1/players`: Get list of all players

2. `GET /api/v1/players/:playerID`: Get player data by ID

3. `GET /api/v1/players/4db77dd4-09b0-4633-aed2-a8382e17a748`: Get player who is not ready to play yet 🙁

4. `GET /api/v1/players/b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755`: Get player who is ready to play 🥳

5. `POST /api/v1/players`: Create new player who absolutely not ready

    ``` json
    # Body
    {
        "name": "Some Name",
        "container_qty": 7,
        "container_capacity": 15
    }
    ```
6. `PUT /api/v1/players/:playerID`: Add 1 ball to a random player's container

7. `PUT /api/v1/players/4db77dd4-09b0-4633-aed2-a8382e17a748`: Add 1 ball to player's container who is not ready yet ✅

8. `PUT /api/v1/players/b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755`: Add 1 ball to player's container who is ready ❌ will response with error

9. `GET /api/v1/containers`: Get list of all container

10. `GET /api/v1/containers/player/:playerID`: Get list of player's containers

11. `GET /api/v1/containers/player/4db77dd4-09b0-4633-aed2-a8382e17a748`: Get list of new and not ready yet player's containers

    ```json
    {
        "data": {
            "containers": [
                {
                    "id": "260f1e55-2abc-42a1-8d4f-ffded1028650",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 0
                },
                {
                    "id": "92897446-654e-4baa-b066-586dc64392c3",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 0
                },
                {
                    "id": "a35c3ac1-1d61-4ae7-8e0c-10929705aac8",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 0
                },
                {
                    "id": "a4482532-6296-45ac-b470-c7693c30ebf6",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 0
                },
                {
                    "id": "bb9c78ea-af2e-400e-9dc4-c06915d9336b",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 1
                },
                {
                    "id": "eb020d63-7fcd-4b95-ab23-1d5a97020a8d",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 1
                },
                {
                    "id": "ff2a397b-25bd-46ad-9a23-5a985124de5d",
                    "playerId": "4db77dd4-09b0-4633-aed2-a8382e17a748",
                    "capacity": 7,
                    "ballQty": 0
                }
            ],
            "isVerified": false
        },
        "status": 200
    }
    ```

12. `GET /api/v1/containers/player/b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755`: Get list of ready and verified player's containers 🥳✅

    ```json
    {
        "data": {
            "containers": [
                {
                    "id": "0004101a-c546-45a6-81dc-786ba8250f9a",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 6
                },
                {
                    "id": "2559812e-1b7d-4810-8e35-a0ec18dee539",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 6
                },
                {
                    "id": "7395b744-6905-4383-b53c-e185d518aee7",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 2
                },
                {
                    "id": "7d05ee26-8620-429b-8556-4dcf8d5bba9d",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 2
                },
                {
                    "id": "7d47aafb-345e-4fe1-9bd6-4cf6d650dc33",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 7
                },
                {
                    "id": "a7241db8-4c63-47fc-b1fd-22aaaea54ec1",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 1
                },
                {
                    "id": "d0fa4320-40b3-4b7c-b5bc-1375465af655",
                    "playerId": "b26eb8e1-f55a-4041-8eb2-4b9f9ce6f755",
                    "capacity": 7,
                    "ballQty": 2
                }
            ],
            "isVerified": true
        },
        "status": 200
    }
    ```

### Test

``` bash
# test file in the `domains/container/service` path
$ go test
```

## App Info

### Authors

Hirzi Nurfakhrian

### Version

1.0.0