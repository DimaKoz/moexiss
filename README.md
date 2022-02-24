# MOEXISS ![Golang](https://img.shields.io/badge/-Golang%20❤️-05122A?style=flat&logo=go&logoColor=white)&nbsp; [![codecov](https://codecov.io/gh/DimaKoz/moexiss/branch/main/graph/badge.svg?token=B4TT54SCA1)](https://codecov.io/gh/DimaKoz/moexiss)&nbsp; [![Go Report Card](https://goreportcard.com/badge/github.com/dimakoz/moexiss)](https://goreportcard.com/report/github.com/dimakoz/moexiss)

An unofficial client of Application Programming Interface for the Moscow Exchange([MOEX](https://iss.moex.com/iss/reference/)) ISS(Informational & Statistical Server) on Golang.  
Неофициальный клиент для получения информации и статистики Московской биржи [MOEX](https://iss.moex.com/iss/reference/) на Golang.

[![stability-experimental](https://img.shields.io/badge/stability-experimental-orange.svg)](https://github.com/emersion/stability-badges#experimental)

**Under active development, pls, do not use it**  
**Находится в разработке, не стоит использовать**

## Usage ##

Construct a new MOEX ISS client, then use the various services on the client to
access different parts of the MOEX ISS API.  
For example:

```go
client := moexiss.NewClient(nil)

// Get current turnover on all the markets.
turnovers, err := client.Turnovers.GetTurnovers(context.Background(), nil)
```

### Getting ISS reference ###

Getting initial ISS reference:

```go
client := moexiss.NewClient(nil)

result, err := client.Index.List(context.Background(), nil)
```
Request options (not mandatory) for initial ISS reference:

* ```Engine()``` - options for the list of trading systems.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.


* ```Market()``` - options for the list of available markets.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.


* ```Board()``` - options for the list of trading modes.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.


* ```BoardGroup()``` - options for the list of board groups.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.
  - ```IsTraded(bool)``` - show traded only currently traded board groups. ```false``` by default.
  - ```WithEngine(EngineName)``` - filtering by an engine.


* ```Duration()``` - selection for the list of available calculation intervals of candles in the HLOCV format.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.


* ```SecurityType()``` - options for the list of securities.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.
  - ```WithEngine(EngineName)``` - filtering by an engine.


* ```SecurityGroup()``` - options for the list of security groups.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.
  - ```WithEngine(EngineName)``` - filtering by an engine.
  - ```HideInactive(bool)``` - filtering by activity.  ```false``` by default.


* ```SecurityCollection()``` - options of collections of the securities list.
  - ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.


An example:

```go
client := moexiss.NewClient(nil)
	
bld := moexiss.NewIndexReqOptionsBuilder()
bld.BoardGroup().Lang(moexiss.LangEn).WithEngine(moexiss.EngineStock).IsTraded(true)
bld.Engine().Lang(moexiss.LangEn)
options := bld.Build()
	
result, err := client.Index.List(context.Background(), options)
```

### Getting trading results ###

Getting aggregated trading results for the date by market:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
result, err := client.Aggregates.Aggregates(context.Background(), ticker, nil)
```

Optional query parameters:

- ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.
- ```Date(time.Date)``` - date of the results. The last date of aggregated trading results by default.

An example:
```go
client := moexiss.NewClient(nil)
ticker := "sberp"
opt := moexiss.NewAggregateReqOptionsBuilder().
Lang(moexiss.LangEn).
Date(time.Date(2021/*year*/, 2/*month*/, 24/*day*/, 12, 0, 0, 0, time.UTC)).
Build()
result, err := client.Aggregates.Aggregates(context.Background(), ticker, opt)
```

### Getting a list of indices that include a security  ###

How to get a list of indices that include a security:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
result, err := client.Indices.GetIndices(context.Background(), ticker, nil)
```

Optional query parameters:

- ```Lang(Language)``` - the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.

An example:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
opt := moexiss.NewIndicesReqOptionsBuilder().
Lang(moexiss.LangEn).
Build()
result, err := client.Indices.GetIndices(context.Background(), ticker, opt)
```


### Get turnovers on all the markets ###

How to get turnovers on all the markets:

```go
client := moexiss.NewClient(nil)
turnovers, err := client.Turnovers.GetTurnovers(context.Background(), nil)
```

Optional query parameters:

- ```Lang(Language)``` — the language of the result. Possible values ```moexiss.LangEn```, ```moexiss.LangRu```. By default, ```moexiss.LangRu```.
- ```Date(time.Date)``` — date for which you want to display data. Today — by default.
- ```IsTonightSession(bool)``` — show turnovers for the evening session. ```false``` — by default.

For example:

```go
client := moexiss.NewClient(nil)
options := moexiss.NewTurnoverReqOptionsBuilder().
Lang(moexiss.LangEn).
Date(time.Date(2021/*year*/, 2/*month*/, 24/*day*/, 12, 0, 0, 0, time.UTC)).
IsTonightSession(true).
Build()
result, err := client.Turnovers.GetTurnovers(context.Background(), options)
```


## Использование ##

Создайте новый MOEX ISS клиент, а затем используйте различные сервисы клиента 
для доступа к различным частям MOEX ISS API.  
Например:

```go
client := moexiss.NewClient(nil)

// Получить сводные обороты по рынкам.
turnovers, err := client.Turnovers.GetTurnovers(context.Background(), nil)
```


### Получение ISS справочников ###

Получить глобальные справочники ISS:

```go
client := moexiss.NewClient(nil)

result, err := client.Index.List(context.Background(), nil)
```


Опции запроса(не являются обязательными) для получения глобальных справочников ISS:

* ```Engine()``` - опции для списка торговых систем.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.  
    

* ```Market()``` - опции для списка доступных рынков.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.


* ```Board()``` - опции для списка режимов торгов.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.


* ```BoardGroup()``` - опции для списка групп режимов торгов.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.
  - ```IsTraded(bool)``` - показывать торгуемые только торгующиеся в настоящий момент группы режимов торгов. Значение по умолчанию - ```false```.
  - ```WithEngine(EngineName)``` - фильтрация по торговой системе.


* ```Duration()``` - опции для списка доступных расчетных интервалов свечей в формате HLOCV.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.


* ```SecurityType()``` - опции для списка инструментов.
  - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.
  - ```WithEngine(EngineName)``` - фильтрация по торговой системе.


* ```SecurityGroup()``` - опции для списка групп инструментов.
    - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.
    - ```WithEngine(EngineName)``` - фильтрация по торговой системе.
    - ```HideInactive(bool)``` - фильтрация по активности. Значение по умолчанию - ```false```.


* ```SecurityCollection()``` - опции для списка коллекций инструментов.
    - ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.
    
    
Пример:

```go
client := moexiss.NewClient(nil)
	
bld := moexiss.NewIndexReqOptionsBuilder()
bld.BoardGroup().Lang(moexiss.LangEn).WithEngine(moexiss.EngineStock).IsTraded(true)
bld.Engine().Lang(moexiss.LangEn)
options := bld.Build()
	
result, err := client.Index.List(context.Background(), options)
```
### Получение итогов торгов ###

Получить агрегированные итоги торгов за дату по рынкам:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
result, err := client.Aggregates.Aggregates(context.Background(), ticker, nil)
```

Опции запроса(не являются обязательными):

- ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.
- ```Date(time.Date)``` - дата за которую необходимо отобразить данные. По умолчанию за последнюю дату в итогах торгов.

Пример:
```go
client := moexiss.NewClient(nil)
ticker := "sberp"
opt := moexiss.NewAggregateReqOptionsBuilder().
Lang(moexiss.LangEn).
Date(time.Date(2021/*год*/, 2/*месяц*/, 24/*день*/, 12, 0, 0, 0, time.UTC)).
Build()
result, err := client.Aggregates.Aggregates(context.Background(), ticker, opt)
```

### Получение списка индексов по бумаге ###

Получить список индексов в которые входит бумага:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
result, err := client.Indices.GetIndices(context.Background(), ticker, nil)
```

Опции запроса(не являются обязательными):

- ```Lang(Language)``` - язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию - ```moexiss.LangRu```.

Пример:

```go
client := moexiss.NewClient(nil)
ticker := "sberp"
opt := moexiss.NewIndicesReqOptionsBuilder().
Lang(moexiss.LangEn).
Build()
result, err := client.Indices.GetIndices(context.Background(), ticker, opt)
```

### Получение сводных оборотов по рынкам ###

Получить сводные обороты по рынкам:

```go
client := moexiss.NewClient(nil)
turnovers, err := client.Turnovers.GetTurnovers(context.Background(), nil)
```

Опции запроса(не являются обязательными):

- ```Lang(Language)``` — язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию — ```moexiss.LangRu```.
- ```Date(time.Date)``` — дата за которую необходимо отобразить данные. По умолчанию — сегодня.
- ```IsTonightSession(bool)``` — показывать обороты за вечернюю сессию. Значение по умолчанию — ```false```.

Пример:

```go
client := moexiss.NewClient(nil)
options := moexiss.NewTurnoverReqOptionsBuilder().
Lang(moexiss.LangEn).
Date(time.Date(2021/*год*/, 2/*месяц*/, 24/*день*/, 12, 0, 0, 0, time.UTC)).
IsTonightSession(true).
Build()
result, err := client.Turnovers.GetTurnovers(context.Background(), options)
```

### Получение данных по листингу бумаг ###

 - Список неторгуемых/торгуемых инструментов с указанием интервалов торгуемости по режимам:

```go
client := moexiss.NewClient(nil)
engine := moexiss.EngineStock
market := "shares"
result, err := client.HistoryListing.
	GetListing(context.Background(), engine, market, nil)
```
 - Получить данные по листингу бумаг в историческом разрезе по указанному режиму:

```go
client := moexiss.NewClient(nil)
engine := moexiss.EngineStock
market := "shares"
board := "TQTD"
result, err := client.HistoryListing.
   GetListingByBoard(context.Background(), engine, market, board, nil)
```
 - Получить данные по листингу бумаг в историческом разрезе по указанной группе режимов:

```go
client := moexiss.NewClient(nil)
engine := moexiss.EngineStock
market := "shares"
boardGroupId := "6"
result, err := client.HistoryListing.
	GetListingByBoardGroup(context.Background(), engine, market, boardGroupId, nil)
```

Опции запроса(не являются обязательными):

- ```Lang(Language)``` — язык результата. Возможные значения ```moexiss.LangEn```, ```moexiss.LangRu```. Значение по умолчанию — ```moexiss.LangRu```.
- ```Status(HistoryListingTradingStatus)``` — фильтр торгуемости инструментов: ```moexiss.ListingTradingStatusTraded```, ```moexiss.ListingTradingStatusNotTraded``` или ```moexiss.ListingTradingStatusAll```. 
  По умолчанию — ```moexiss.ListingTradingStatusAll```.
- ```Start(int)``` — номер строки (отсчет с нуля), с которой следует начать порцию возвращаемых данных. Значение по умолчанию — 0. 
  Получение ответа без данных означает, что указанное значение превышает число строк, возвращаемых запросом.


Пример:

```go
client := moexiss.NewClient(nil)
engine := moexiss.EngineStock
market := "shares"
boardGroupId := "6"
opt := moexiss.NewHistoryListingReqOptionsBuilder().
    Status(moexiss.ListingTradingStatusNotTraded).
    Start(42).
    Lang(moexiss.LangEn).
    Build()
result, err := client.HistoryListing.
    GetListingByBoardGroup(context.Background(), engine, market, boardGroupId, opt)
```
