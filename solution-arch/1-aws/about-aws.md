# Концепции географического распределения физических дата-центров AWS

## Задание

1. На официальном сайте вендора найти раздел документации, связанных с географией размещения центров данных вендора.
2. Записать определения уровней градации центров данных (региона, зона досутпности, edge point и т.д.)
3. Выписать минимальные требования отказоустойчивости, реализуемые вендором при планировании размещения центров данных.
4. Составить список сервисов провайдера, которые независимы от места размещения центра данных.
5. В результате следует получить - список терминов, описывающих географическое разделение инфраструктуры провайдера, минимальные характеристики отказоустойчивости и таблицы сервисов.

## Источники

При работе я опирался на следующие источники информации:
- https://aws.amazon.com/about-aws/global-infrastructure/regions_az/?p=ngi&loc=2
- https://stackoverflow.com/questions/63786743/aws-edge-locations-vs-local-zones
- https://www.youtube.com/watch?v=Uk2A9-JO-_w
- https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
- https://aws.amazon.com/cloudfront/features/?whats-new-cloudfront.sort-by=item.additionalFields.postDateTime&whats-new-cloudfront.sort-order=desc#Global_Edge_Network
-

## Определения уровней градации ДЦ в AWS

**Регион** - region - отдельная физическая локация, которая объединяет в себе Зоны доступности. Каждый регион географически отделён от других регионов.

**Зона доступности** - availability zone (AZ) - это один или больше отдельных центров данных с резервной инфраструктурой питания, сети и т.д. Отдельные зоны доступности внутри региона отделены друг от друга, но не более чем на 100 км.

**Локальная зона** - local zone - механизм расположения мощностей облачного провайдера ближе к конечному пользователю для снижения latency до однозначных величин. Распологаются как расширение соответствующего региона.

**Граничная зона** - edge location - внерегиональный механизм работы некоторых сервисов облачного провайдера (Amazon CloudFront, Amazon Route 53, AWS Firewall Manager и др.) через размещение внутри Global Edge Network для ещё большего сокращения расстояния до конечного пользователя.

**Операторские зоны** -  wavelength - механизм размещения инфраструктуры облачного провайдера у телекоммуникационных операторов для ещё большего снижения задержек (5G)

**Отдалённые зоны** - outpost - механихм размещения инфраструктуры облачного провайдера в любом центре данных (например, у заказчика) для решения сложных задач с высокими требованиями по задержкам и надёжности.

## Отказоустойчивость центров данных

При размещении центров данных внутри зон доступности облачный провайдер придерживается следующих критериев:
- наличие резервного питания
- наличие резервного сетевого соединения с другими датацентрами/регионами
- наличие шифрованного канала между ДЦ в регионе
- наличие сетевого канала достаточного для синхронной репликации между зонами доступности
- географическое удаление от других центров (но не более чем 100 км.)

## Списки сервисов

- Базовые сервисы AWS доступны при открытии каждого нового региона - Amazon API Gateway, Amazon Aurora, Amazon CloudWatch, Amazon DynamoDB, Amazon EC2 Auto Scaling, Amazon ElastiCache, Amazon Elastic Block Store (EBS), Amazon Elastic Compute Cloud (EC2), Amazon Elastic Container Registry (ECR), Amazon Elastic Container Service (ECS), Amazon Elastic MapReduce (EMR), Amazon OpenSearch Service, Amazon EventBridge, Amazon Kinesis Data Streams, Amazon Redshift, Amazon Relational Database Service (RDS), Amazon Route 53, Amazon Simple Notification Service (SNS), Amazon Simple Queue Service (SQS), Amazon Simple Storage Service (S3), Simple Workflow Service (SWF), Amazon Virtual Private Cloud (VPC), AWS Application Auto Scaling, AWS Certificate Manager, AWS CloudFormation, AWS CloudTrail, AWS CodeDeploy, AWS Config, AWS Database Migration Service, AWS Direct Connect, AWS Identity & Access Management (IAM), AWS Key Management Service, AWS Lambda, AWS Marketplace, AWS Health Dashboard, AWS Step Functions, AWS Support, AWS Systems Manager, AWS Trusted Advisor, AWS VPN, and Elastic Load Balancing (ELB).
- Часть сервисов планируется запускать только через год после открытия региона - Amazon Athena, Amazon CloudFront, Amazon Elastic File System (EFS), Amazon Elastic Kubernetes Services (EKS), Amazon GuardDuty, Amazon Kinesis Firehose, Amazon MQ, Amazon SageMaker, AWS Backup, AWS Batch, AWS Certificate Manager Private Certificate Authority, AWS Chatbot, AWS CodeBuild, AWS Console Mobile App, AWS Directory Service, AWS Fargate, AWS Glue, AWS LakeFormation, AWS License Manager, AWS Organizations, AWS Resource Access Manager (RAM), AWS Secrets Manager, AWS Security Hub, AWS Service Catalog, AWS Storage Gateway, AWS Transit Gateway, AWS WAF, and AWS X-Ray.
- Также существуют сервисы, доступность которых ограничивается:
  - географическим расположением - Amazon CloudFront, Amazon Route 53, AWS Firewall Manager, AWS Shield, и AWS WAF сервисы предоставляются через AWS Edge Locations
  - принадлежностью к определелённой организации - AWS GovCloud (US)
- Присутствует возможность получить список доступных сервисов для каждого региона отдельно - https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/
