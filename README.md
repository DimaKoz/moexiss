# MOEXISS

Неофициальный клиент для получения информации и статистики Московской биржи [MOEX](https://iss.moex.com/iss/reference/) на Golang  
An unofficial client of Application Programming Interface for the Moscow Exchange([MOEX](https://iss.moex.com/iss/reference/)) ISS(Informational & Statistical Server) on Golang 

Under active development, pls, do not use it  
Находится в разработке, не стоит использовать

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