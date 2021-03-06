# spot-ninja

Software responsável pela inteligência de aumentar ou diminuir o tamanho do AutoScalingGroups, baseado na saúde do Spotfleet.

## Licença

spot ninja está licenciada por Apache License. Veja o arquivo LICENSE para mais detalhes ou o link [apache](https://www.apache.org/licenses/LICENSE-2.0).

## Requerimentos

- Docker >=v1.13.1
- Docker Compose >=1.23.1
- Go >= 1.12.7
- Modulo aws-sdk-go
- Modulo yaml.v2
- Polices:
  - CloudWatchReadOnlyAccess
  - AmazonEC2SpotFleetAutoscaleRole
  - AutoScalingConsoleFullAccess

## Arquitetura

![spot-ninja](assets/spot_PT_BR.png)

## Variáveis de ambiente

Para configurar o spot-ninja, é necessário configurar algumas variáveis de ambiente. As variáveis são:

- AWS_REGION
  - Default us-east-1
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
- PREFIX
  - Padrão ecs-
- TIME_TO_LIVE
  - Em segundos padrão 15 min.

## Instalando o projeto via o docker-compose

1 - Build via docker-compose

```bash
docker-compose build
```

2 - Execute o spot-ninja

```bash
docker-compose up
```

## Estrutura

- cmd/
  - main spot-ninja
- config/
  - configs gerais do spot ninja
- pkg/
  - libs

## Referências

- [gopherize](https://www.gopherize.me/)
- [cloudcraft](https://cloudcraft.co/)
- [project-layout](https://github.com/golang-standards/project-layout)
- [regex101](https://regex101.com/r/FwSMp7/1/)
- [sdk-for-go](https://docs.aws.amazon.com/sdk-for-go/api/)
