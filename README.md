# MOEXISS

Неофициальный [MOEX](https://iss.moex.com/iss/reference/) ISS API на Golang  
Unofficial Application Programming Interface for the Moscow Exchange([MOEX](https://iss.moex.com/iss/reference/)) ISS on Golang 

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

## Использование ##

Создайте новый клиент ISS MOEX, а затем используйте различные сервисы клиента 
для доступа к различным частям  MOEX ISS API.  
Например:

```go
client := moexiss.NewClient(nil)

// Получить сводные обороты по рынкам.
turnovers, err := client.Turnovers.Turnovers(context.Background(), nil)
```