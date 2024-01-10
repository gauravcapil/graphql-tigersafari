go run github.com/99designs/gqlgen generate

echo remember to put " gorm:"foreignKey:ID;references:ID"` "  on all foriegn keys and "gorm:"primary_key" " for all auto incremenet keys in models_gen.go