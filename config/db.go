package config

type Db struct {
driver string = "postgres"
 source string = "postgres://postgres:freemail@localhost/test?sslmode=disable"
}