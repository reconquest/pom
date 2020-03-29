package pom

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnmarshal(t *testing.T) {
	test := assert.New(t)

	data := `
<project>
	<version>a</version>
</project>
`
	model, err := Unmarshal([]byte(data))
	test.NoError(err)
	test.NotNil(model)
}

func TestModel_Get(t *testing.T) {
	test := assert.New(t)

	data := `
<project>
	<version>${p1}.${a.b}.z</version>
	<broken>${p1}.${z}</broken>
	<properties>
		<p1>x${p.b}</p1>
		<p.b>x</p.b>
	</properties>
</project>
`
	model, err := Unmarshal([]byte(data))
	test.NoError(err)
	test.NotNil(model)

	model.SetProperty("a.b", "y")

	version, err := model.Get("version")
	test.NoError(err)
	test.Equal("xx.y.z", version)

	broken, err := model.Get("broken")
	test.Empty(broken)
	test.Error(err)
	test.IsType(ErrMissingField{}, err)
	test.True(IsMissingField(err))
}
