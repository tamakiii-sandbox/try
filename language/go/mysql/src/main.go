// https://qiita.com/taizo/items/54f5f49c6102f86194b8
package main

import (
  "database/sql"
  "fmt"
  _ "github.com/go-sql-driver/mysql"
)

func main() {
  db, err := sql.Open("mysql", "root:root@tcp(mysql)/mysql")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  rows, err := db.Query("SELECT * FROM proc")
  if err != nil {
    panic(err.Error())
  }

  columns, err := rows.Columns()
  if err != nil {
    panic(err.Error())
  }

  values := make([]sql.RawBytes, len(columns))

  scanArgs := make([]interface{}, len(values))
  for i := range values {
    scanArgs[i] = &values[i]
  }

  for rows.Next() {
    err = rows.Scan(scanArgs...)
    if err != nil {
      panic(err.Error())
    }

    var value string
    for i, col := range values {
      if col == nil {
        value = "NULL"
      } else {
        value = string(col)
      }
      fmt.Println(columns[i], ": ", value)
    }
    fmt.Println("---")
  }
}
