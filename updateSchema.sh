go run github.com/99designs/gqlgen generate

 echo remember to put " gorm:"foreignKey:ID\;references:ID" "  on all foriegn keys and "gorm:"primary_key" " for all auto incremenet keys and  gorm:"unique_key" for unique keys like username and tigername in models_gen.go

