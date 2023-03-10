//go:build simpletest || all
// +build simpletest all

package client

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	httpConfiguration := &HttpConfiguration{
		BaseURL:       "http://www.omdbapi.com/",
		TimeoutMillis: 120000,
		ApiKey:        os.Getenv("APIKEY"),
	}

	c := NewClient(httpConfiguration)

	res, err := c.Search("stem", nil)
	assert.Nil(t, err, "expecting nil err")
	assert.NotNil(t, res, "expecting non-nil result")

	t.Run("Assert that the result should contain at least 30 items", func(*testing.T) {
		assert.Greater(t, len(res), 30, "expecting more than 30 results")
	})

	t.Run("Assert that the result contains: The STEM Journal ,Activision: STEM - in the Videogame Industry ", func(*testing.T) {
		stringOne := "The STEM Journals"
		stringTwo := "Activision: STEM - in the Videogame Industry"

		foundStringOne := false
		foundStringTwo := false

		for _, item := range res {
			if strings.Contains(item.Title, stringOne) {
				foundStringOne = true
				// fmt.Println(item)
			}
			if strings.Contains(item.Title, stringTwo) {
				foundStringTwo = true
				// fmt.Println(item)
			}

			if foundStringTwo == true && foundStringOne == true {
				break
			}
		}

		assert.Equal(t, true, foundStringOne, "expecting string : The STEM Journals")
		assert.Equal(t, true, foundStringTwo, "Activision: STEM - in the Videogame Industry")
	})

	t.Run("Get movie by Id(Activision: STEM - in the Videogame Industry) detail using get_by_id ", func(*testing.T) {
		stringTwo := "Activision: STEM - in the Videogame Industry"
		stringReleaseDate := "23 Nov 2010"
		stringDirector := "Mike Feurstein"
		foundStringTwo := false

		var foundId string
		for _, item := range res {
			if strings.Contains(item.Title, stringTwo) {
				foundStringTwo = true
				foundId = item.ImdbID
				// fmt.Printf("id:%s", foundId)
			}
			if foundStringTwo == true {
				break
			}
		}
		assert.Equal(t, true, foundStringTwo, "Activision: STEM - in the Videogame Industry")
		res, err := c.get_by_id(foundId)
		assert.Nil(t, err, "expecting nil err")
		assert.NotNil(t, res, "expecting non-nil result")

		assert.Equal(t, res.Released, stringReleaseDate, "expecting the correct release date")
		assert.Equal(t, res.Director, stringDirector, "expecting the correct director ")

	})

	t.Run("Get movie by  Title  (The STEM Journals) detail using get_by_title ", func(*testing.T) {

		searchTitle := "The STEM Journals"
		res, err := c.get_by_title(searchTitle)
		assert.Nil(t, err, "expecting nil err")
		assert.NotNil(t, res, "expecting non-nil result")

		// check readme for why we lower case .
		plotstring := strings.ToLower("Science, Technology, Engineering and Math")
		durationstring := "22 min"

		assert.Contains(t, res.Plot, plotstring, "expecting the correct content in the plot")
		assert.Equal(t, res.Runtime, durationstring, "expecting the correct runtime ")

	})
}
