package goi

type Validatable interface {
	Validate(data *any) error
}

type Validator struct {
	value     *any
	rules     []func(*any) error
	ruleNames []string
	name      string
	Validator Validatable
}

func findIndex(arr []string, el string) int {
	for i, v := range arr {
		if el == v {
			return i
		}
	}
	return -1
}
