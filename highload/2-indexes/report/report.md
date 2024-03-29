# Работа с индексами

В рамках данного задания я работал с таблицей users, в которой было 1_000_000 записей

При нагрузочном тестировании доходил только до 100 одновременных запросов - дальше процессор захлёбывался.

Индекс добавлял через `create index users_name_idx on users(first_name,second_name) using BTREE;`

Индекс выбрал таким - по двум колонкам, поскольку MySQL не сможет использовать в одном запросе два одиночных и тестовые запросы показали лучшую работу данного индекса.

## Вывод EXPLAIN

### До добавления индексов

![](./explain-before-indexes.png)

### После добавления индексов

![](./explain-after-indexes.png)

## Графики latency/throughput

### До добавления индексов (1)

![](./tests-before-1.png)

### До добавления индексов (10)

![](./tests-before-10.png)

### До добавления индексов (100)

![](./tests-before-100.png)

### После добавления индексов (1)

![](./tests-after-1.png)

### После добавления индексов (10)

![](./tests-after-10.png)

### После добавления индексов (100)

![](./tests-after-100.png)
