package util

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type Student struct {
	Name   string   `env:"STUDENT_NAME,required"`
	Class  string   `env:"STUDENT_CLASS" envDefault:"5-1"`
	Gender int      `env:"STUDENT_GENDET" envDefault:"1"`
	Hobby  []string `env:"STUDENT_HOBBY" envSeparator:":"` //  default envSeparator:","
}

func TestMustLoadConfig(t *testing.T) {
	var (
		name  = "jmz"
		hobby = "Basketball:book:go fishing"
	)
	err := os.Setenv("STUDENT_NAME", name)
	assert.NoError(t, err)
	err = os.Setenv("STUDENT_HOBBY", hobby)
	assert.NoError(t, err)
	s := &Student{}

	MustLoadConfig(s)
	assert.Equal(t, s.Name, name)
	assert.Equal(t, s.Hobby[0], "Basketball")
	assert.Equal(t, s.Gender, 1)
	assert.Equal(t, s.Class, "5-1")
}
