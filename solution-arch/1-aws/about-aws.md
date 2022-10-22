# Концепции географическое распределения физических дата-центров AWS

Источники:
- https://aws.amazon.com/about-aws/global-infrastructure/regions_az/?p=ngi&loc=2
- https://stackoverflow.com/questions/63786743/aws-edge-locations-vs-local-zones
- https://www.youtube.com/watch?v=Uk2A9-JO-_w
- https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-regions-availability-zones.html
- https://aws.amazon.com/cloudfront/features/?whats-new-cloudfront.sort-by=item.additionalFields.postDateTime&whats-new-cloudfront.sort-order=desc#Global_Edge_Network
-

1. На официальном сайте вендора найти раздел документации, связанных с географией размещения центров данных вендора.
2. Записать определения уровней градации центров данных (региона, зона досутпности, edge point и т.д.)
3. Выписать минимальные требования отказоустойчивости, реализуемые вендором при планировании размещения центров данных.
4. Составить список сервисов провайдера, которые независимы от места размещения центра данных.
5. В результате следует получить - список терминов, описывающих географическое разделение инфраструктуры провайдера, минимальные характеристики отказоустойчивости и таблицы сервисов.

## Определения уровней градации ДЦ в AWS

**Регион** - region -

**Зона доступности** - availability zone (AZ) -

Region
Availability Zone
Local Zone
Edge Location
Wavelength
Outpost

