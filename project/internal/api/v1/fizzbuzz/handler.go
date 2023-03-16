package fizzbuzz

import (
	"strconv"

	"github.com/anotherhope/FizzBuzz/project/internal/model/fizzbuzz"
	"github.com/gofiber/fiber/v2"
)

// @Summary		Return FizzBuzz result.
// @Description	Returns a list of strings with numbers from 1 to limit, where: \n all multiples of int1 are replaced by str1, all multiples of int2 are replaced by str2, all multiples of int1 and int2 are replaced by str1str2.
// @Tags		FizzBuzz
// @Accept		*/*
// @Produce		json
// @Help		Name    ?        type    required   description
// @Param       int1    path     int     true       "Give the first number"  default(3)
// @Param       int2    path     int     true       "Give the second number" default(5)
// @Param       limit   path     int     true       "Give limit of fizzbuzz" default(100)
// @Param       str1    path     string  true       "Give the first word"    default(fizz)
// @Param       str2    path     string  true       "Give the second word"   default(buzz)
// @Failure     400  	string  string
// @Success		200		{array}  string
// @Router		/api/v1/fizzbuzz/:int1/:int2/:limit/:str1/:str2 [get]
func FizzBuzz(c *fiber.Ctx) error {
	var fbr fizzbuzz.Request = c.Context().UserValue("fizzbuzz.Request").(fizzbuzz.Request)
	var result []string

	for i := 1; i <= fbr.Limit; i++ {
		if i%fbr.Int1 == 0 && i%fbr.Int2 == 0 {
			result = append(result, fbr.Str1+fbr.Str2)
		} else if i%fbr.Int1 == 0 {
			result = append(result, fbr.Str1)
		} else if i%fbr.Int2 == 0 {
			result = append(result, fbr.Str2)
		} else {
			result = append(result, strconv.Itoa(i))
		}
	}

	return c.JSON(result)
}

func FizzBuzzControls(c *fiber.Ctx) error {
	int1, err := strconv.Atoi(c.Params("int1"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("int1 parameter is not an integer")
	}

	int2, err := strconv.Atoi(c.Params("int2"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("int2 parameter is not an integer")
	}

	limit, err := strconv.Atoi(c.Params("limit"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("limit parameter is not an integer")
	}

	if int1 == 0 || int2 == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("int1 and int2 parameters must be greater than 0")
	}

	if limit == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("limit parameter must be greater than 0")
	}

	str1 := c.Params("str1")
	str2 := c.Params("str2")

	fbr := fizzbuzz.Request{
		Int1:  int1,
		Int2:  int2,
		Limit: limit,
		Str1:  str1,
		Str2:  str2,
	}

	c.Context().SetUserValue("fizzbuzz.Request", fbr)

	return c.Next()
}
