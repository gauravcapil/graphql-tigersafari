go run github.com/99designs/gqlgen generate

 echo remember to put " gorm:"foreignKey:ID\;references:ID" "  on all foriegn keys and "gorm:"primary_key" " for all auto incremenet keys and  gorm:"unique_key" for unique keys like username and tigername in models_gen.go

echo  "The below code is not going to be present in generated.go so paste it there"
echo "

func unmarshalNDateTime2string(v interface{}) (string, error) {
	switch v := v.(type) {
	case string:
		_, err := time.Parse("2006-01-02 15:04:05.999", v)
	    if err != nil {
        	return "", fmt.Errorf("%T is not a DateTime string", v)
    	}
		return v, nil
	case nil:
		return "null", nil
	default:
		return "", fmt.Errorf("%T is not a string", v)
	}
}


func (ec *executionContext) unmarshalNDateTime2string(ctx context.Context, v interface{}) (string, error) {
	res, err := unmarshalNDateTime2string(v)
	return res, graphql.ErrorOnPath(ctx, err)
}
"