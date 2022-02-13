# MOEXISS ![Golang](https://img.shields.io/badge/-Golang%20❤️-05122A?style=flat&logo=go&logoColor=white)&nbsp;

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