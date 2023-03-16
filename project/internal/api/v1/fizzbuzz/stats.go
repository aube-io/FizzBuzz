package fizzbuzz

import (
	"sync"

	"github.com/anotherhope/FizzBuzz/project/internal/model/fizzbuzz"
	"github.com/gofiber/fiber/v2"
)

var statistics = make(map[fizzbuzz.Request]uint)
var mutex = &sync.Mutex{}

// @Summary		Return FizzBuzz statistics.
// @Description	Return the parameters corresponding to the most used request, as well as the number of hits for this request.
// @Tags		FizzBuzz
// @Accept		*/*
// @Produce		json
// @Help		Name    ?        type    required   description
// @Success		200		{object}  fizzbuzz.Stats
// @Router		/api/v1/fizzbuzz/stats [get]
func Stats(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	var mostUsedRequest fizzbuzz.Request
	var maxHits uint

	for req, hits := range statistics {
		if hits > maxHits {
			maxHits = hits
			mostUsedRequest = req
		}
	}

	return c.JSON(fizzbuzz.Stats{
		Request: mostUsedRequest,
		Hits:    maxHits,
	})
}

func FizzBuzzHits(c *fiber.Ctx) error {
	mutex.Lock()
	defer mutex.Unlock()

	var fbr fizzbuzz.Request = c.Context().UserValue("fizzbuzz.Request").(fizzbuzz.Request)
	statistics[fbr]++

	return c.Next()
}
