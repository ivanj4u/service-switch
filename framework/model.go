package framework

type object interface{}

type restUrl struct {
	urlId           string
	isTransaction   string
	isReffSwitching string
	amountField     string
	keyField        string
	routingField    string
	requestType     string
	urlName         string
	isExisting      string
}

type restField struct {
	fieldName string
	minLength int
	maxLength int
}
