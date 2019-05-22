# SQLSTR

String manipulation/helper for SQL Query.

## Examples

### Clean SQL Query from double white space, comment, etc.

```go
cleaned := sqlstr.Clean(`
  SELECT *
  FROM table -- some table comment 
  WHERE column1 = 'meh' /* request from tyrion*/`)

fmt.Println(cleaned) 

// Output: 
// SELECT * FROM table WHERE column1 = 'meh'
```
### Obscure value in SQL Query

```go
obsecured := sqlstr.Obscure(`SELECT * FROM table WHERE column1 = 'text' AND column2 = 1234 AND column3 = TRUE and column4 = 3.14`)

fmt.Println(obsecured)

// Output: 
// SELECT * FROM table WHERE column1 = ? AND column2 = ? AND column3 = ? and column4 = ?
```

## Author 

iman.tung@gmail.com
