package yammm_test

import "testing"

func Test_Validator_OptionalSpacevec(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary
		embedding Spacevector[3]
	}
	`
	instance := `
	{
		"Cars": [
			{ "regNbr": "ABC123" }	
		]
	}`
	validateMessages(t, model, instance)
}
func Test_Validator_RequiredMissingSpacevec(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary
		embedding Spacevector[3] required
	}
	`
	instance := `
	{
		"Cars": [
			{ "regNbr": "ABC123" }	
		]
	}`
	validateMessages(t, model, instance,
		"[Test_Validator_RequiredMissingSpacevec:4:4] Property value of 'Car.embedding' is required and is missing",
	)
}
func Test_Validator_SpacevecWithInstanceData(t *testing.T) {
	model := `
	schema "testing"
	type Car { regNbr String primary
		embedding Spacevector[3]
	}
	`
	instance := `
	{
		"Cars": [
			{ "regNbr": "ABC123", "embedding": [1.1, 2.1, 3.1] }	
		]
	}`
	validateMessages(t, model, instance)
}
