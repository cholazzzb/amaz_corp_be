# amaz_corp_be

# 1. Generate SQLC

```
sqlc generate
```

## NOTES:

For Windows Only, sqlc postgres engine are not supported on windows, 
1. so install sqlc on docker 
`docker pull kjconroy/sqlc` 

2. and run this command
`docker run --rm -v ${pwd}:/src -w /src kjconroy/sqlc generate .\sqlc.yaml`. [Reference](https://docs.sqlc.dev/en/stable/overview/install.html)
