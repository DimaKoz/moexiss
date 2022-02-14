# MOEXISS ![Golang](https://img.shields.io/badge/-Golang%20❤️-05122A?style=flat&logo=go&logoColor=white)&nbsp; [![codecov](https://codecov.io/gh/DimaKoz/moexiss/branch/main/graph/badge.svg?token=B4TT54SCA1)](https://codecov.io/gh/DimaKoz/moexiss)


Неофициальный клиент для получения информации и статистики Московской биржи [MOEX](https://iss.moex.com/iss/reference/) на Golang.  
An unofficial client of Application Programming Interface for the Moscow Exchange([MOEX](https://iss.moex.com/iss/reference/)) ISS(Informational & Statistical Server) on Golang. 

**Under active development, pls, do not use it**  
**Находится в разработке, не стоит использовать**

## Usage ##

Construct a new MOEX ISS client, then use the various services on the client to
access different parts of the MOEX ISS API.  
For example:

```go
client := moexiss.NewClient(nil)

// Get current turnover on all the markets.
turnovers, err := client.Turnovers.Turnovers(context.Background(), nil)
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


## Использование ##

Создайте новый MOEX ISS клиент, а затем используйте различные сервисы клиента 
для доступа к различным частям MOEX ISS API.  
Например:

```go
client := moexiss.NewClient(nil)

// Получить сводные обороты по рынкам.
turnovers, err := client.Turnovers.Turnovers(context.Background(), nil)
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
