package iconoflix_test

import (
	"testing"

	"github.com/derailed/iconoflix"
	"github.com/stretchr/testify/assert"
)

func TestLoadFile(t *testing.T) {
	data, err := iconoflix.LoadFile("./assets/movies.yml")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 2)
	assert.Equal(t, "m1", data.Movies[0].Name)
	assert.Equal(t, "ic1", data.Movies[0].Icons[0].Emoji)
}

func TestLoadFileFail(t *testing.T) {
	_, err := iconoflix.LoadFile("./assets/blee.yml")
	assert.NotNil(t, err)
}

func TestLoadMem_BoxOffice(t *testing.T) {
	data, err := iconoflix.LoadMem("v1")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 9)
	assert.Equal(t, "Home Alone", data.Movies[0].Name)
	assert.Equal(t, "🏡", data.Movies[0].Icons[0].Emoji)
}

func TestLoadMem_BMovies(t *testing.T) {
	data, err := iconoflix.LoadMem("v2")
	assert.Nil(t, err)

	assert.Equal(t, len(data.Movies), 1)
	assert.Equal(t, "Cobra", data.Movies[0].Name)
	assert.Equal(t, "🕶", data.Movies[0].Icons[0].Emoji)
}

func TestRandMovie(t *testing.T) {
	m := iconoflix.RandMovie("v1")
	assert.NotEqual(t, "", m.Name)
}
