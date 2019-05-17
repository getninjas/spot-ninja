# spot-ninja

Available translations:

- [pt-BR](#pt-bR)

This software is responsible for all the intelligence in increasing and decreasing AutoScalingGroups size based on Spotfleet health.

## Requirements

- Docker >=v1.13.1
- Docker Compose >=1.23.1
- Go >= 1.11.2
- aws-sdk-go
- yaml.v2
- Polices:
  - CloudWatchReadOnlyAccess
  - AmazonEC2SpotFleetAutoscaleRole
  - AutoScalingConsoleFullAccess

## ENVs configuration

- DATA_POINT_CONFIG
  - Default 30
- STATISTIC_CONFIG
  - Default Average
- FLEET_TYPE_CONFIG
  - Default FleetRequestId
- ID_METRIC_CONFIG
  - Default metric
- SCAN_TYPE_CONFIG
  - Default TimestampDescending
- METRIC_NAME_CONFIG
  - Default PendingCapacity
- NAMESPACE_CONFIG
  - Default AWS/EC2Spot
- UNIT_CONFIG
  - Default Count
- TIME_TO_CONFIG
  - Default 5
- SPEED_CONFIG
  - Default 10
- DIVIDER_CONFIG
  - Default 4
- FLEET_IGNORED
  - Default none
- ENABLE_EVENTS_ON_SQS
  - Default false
- SQS_URL
  - Default none

## Installing with Docker

1 - Set your AWS_REGION env variable:

```bash
export AWS_REGION=us-east-1
```

2 - Build

```bash
docker-compose build
```

3 - Run spot-ninja

```bash
docker-compose up
```

## Installing without Docker

1 - Set your AWS_REGION env variable:

```bash
export AWS_REGION=us-east-1
```

2 - Install the binary

```bash
go get github.com/getninjas/spot-ninja
```

3 - Run spot-ninja

## Manual build

1 - Clone this repository:

```bash
git clone git@github.com:getninjas/spot-ninja -b master
```

2 - Modify config/config.go according to your needs:

3 - Set your AWS_REGION env variable:

```bash
export AWS_REGION=us-east-1
```

3 - After testing, run the build:

```bash
 go build -o spot-ninja ./cmd/spot-ninja/main.go
 ```

## Project Structure

- assets/
  - images and icons.
- config/
  - spot-fleet config file.
- internal/
  - source code.
- Parameter Store:
  - /{CLUSTER}/spot-ninja/SPOT_CONFIG

### Refs

- [gopherize](https://www.gopherize.me/)
- [project-layout](https://github.com/golang-standards/project-layout)
- [regex101](https://regex101.com/r/FwSMp7/1/)
- [sdk-for-go](https://docs.aws.amazon.com/sdk-for-go/api/)

## pt-BR

Software responsável pela inteligência de aumentar ou diminuir o tamanho do AutoScalingGroups, baseado na saúde do Spotfleet.

## Requerimentos

- Docker >=v1.13.1
- Docker Compose >=1.23.1
- Go >= 1.11.2
- aws-sdk-go
- yaml.v2
- Polices:
  - CloudWatchReadOnlyAccess
  - AmazonEC2SpotFleetAutoscaleRole
  - AutoScalingConsoleFullAccess

- DATA_POINT_CONFIG
  - Padrão 30
- STATISTIC_CONFIG
  - Padrão Average
- FLEET_TYPE_CONFIG
  - Padrão FleetRequestId
- ID_METRIC_CONFIG
  - Padrão metric
- SCAN_TYPE_CONFIG
  - Padrão TimestampDescending
- METRIC_NAME_CONFIG
  - Padrão PendingCapacity
- NAMESPACE_CONFIG
  - Padrão AWS/EC2Spot
- UNIT_CONFIG
  - Padrão Count
- TIME_TO_CONFIG
  - Padrão 5
- SPEED_CONFIG
  - Padrão 10
- DIVIDER_CONFIG
  - Padrão 4
- FLEET_IGNORED
  - Padrão nenhum
- ENABLE_EVENTS_ON_SQS
  - Padrão false
- SQS_URL
  - Padrão nenhum

## Instalando o projeto via o Docker

1 - Tenha o AWS_REGION configurado:

```bash
export AWS_REGION=us-east-1
```

2 - Build via  docker-compose

```bash
docker-compose build
```

3 - Execute o spot-ninja

```bash
docker-compose up
```

## Instalando o projeto sem o Docker

1 - Tenha o AWS_REGION configurado:

```bash
export AWS_REGION=us-east-1
```

2 - Instale o binário

```bash
go get github.com/getninjas/spot-ninja
```

3 - Execute o spot-ninja

## Build o projeto manualmente

1 - Clone o repositório:

```bash
git clone git@github.com:getninjas/spot-ninja -b master
```

2 - Altere o config/config.go para as suas configurações personalizadas:

3 - Tenha o AWS_REGION configurado:

```bash
export AWS_REGION=us-east-1
```

3 - Após test, execute o build:

```bash
 go build -o spot-ninja ./cmd/spot-ninja/main.go
 ```

## Estrutura

- assets/
  - imagens e icones.
- config/
  - arquivos de configuração do spot-fleet.
- internal/
  - código fonte do projeto.
- Parameter Store:
  - /{CLUSTER}/spot-ninja/SPOT_CONFIG

## Ref

- [gopherize](https://www.gopherize.me/)
- [project-layout](https://github.com/golang-standards/project-layout)
- [regex101](https://regex101.com/r/FwSMp7/1/)
- [sdk-for-go](https://docs.aws.amazon.com/sdk-for-go/api/)
