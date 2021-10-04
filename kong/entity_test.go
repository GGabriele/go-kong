package kong

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEntityValidate(T *testing.T) {
	assert := assert.New(T)

	client, err := NewTestClient(nil, nil)
	assert.Nil(err)
	assert.NotNil(client)

	isValid, err := client.Entity.Validate(defaultCtx, "acme_storage")
	assert.Nil(err)
	assert.True(isValid)

	isValid, err = client.Entity.Validate(defaultCtx, "not_existing_entity")
	assert.Equal(err.Error(), "HTTP status 404 (message: \"No entity named 'not_existing_entity'\")")
	assert.False(isValid)
}
