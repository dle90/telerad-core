package utils

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetPaginationParams(c *fiber.Ctx) (int, int) {
	pageStr := c.Query("page", "1")
	sizeStr := c.Query("size", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(sizeStr)
	if err != nil || size < 1 {
		size = 10
	}

	return page, size
}

func GetInt64FromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*int64, error) {
	var output *int64 = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		value, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		output = &value
	}

	return output, nil
}

func GetInt16FromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*int16, error) {
	var output *int16 = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		value, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		int16Value := int16(value)
		output = &int16Value
	}

	return output, nil
}

func GetIntFromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*int, error) {
	var output *int = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		value, err := strconv.ParseInt(str, 10, 64)
		if err != nil {
			return nil, err
		}

		intValue := int(value)
		output = &intValue
	}

	return output, nil
}

func GetBoolFromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*bool, error) {
	var output *bool = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		boolValue, err := strconv.ParseBool(str)
		if err != nil {
			return nil, err
		}

		output = &boolValue
	}

	return output, nil
}

func GetUuidFromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*uuid.UUID, error) {
	var output *uuid.UUID = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		uuidValue, err := uuid.Parse(str)
		if err != nil {
			return nil, err
		}

		output = &uuidValue
	}

	return output, nil
}

func GetStringFromRequestParam(c *fiber.Ctx, paramName string) string {
	return c.Query(paramName, "")
}

// GetTimeFromRequestParam parses a query param as a time value. Accepts RFC3339
// (e.g. "2026-05-18T08:30:00+07:00") or date-only "2006-01-02". Returns zero
// time.Time when the param is missing/empty.
func GetTimeFromRequestParam(c *fiber.Ctx, paramName string, allowNull bool) (*time.Time, error) {
	var output *time.Time = nil
	str := c.Query(paramName, "")

	if str != "" || !allowNull {
		t, err := time.Parse(time.RFC3339, str)
		if err != nil {
			return nil, err
		}

		output = &t
	}

	return output, nil
}

func GetUuidFromRequestPath(c *fiber.Ctx, paramName string) (uuid.UUID, error) {
	str := c.Params(paramName)

	if _id, err := uuid.Parse(str); err != nil {
		return uuid.Nil, err
	} else {
		return _id, nil
	}
}

func GetInt64SliceFromRequestParam(c *fiber.Ctx, paramName string) ([]int64, error) {
	output := []int64{}
	queryParams := c.Context().QueryArgs()
	aParams := []string{}

	queryParams.VisitAll(func(key, value []byte) {
		if string(key) == paramName {
			valueStr := string(value)
			if valueStr != "" {
				aParams = append(aParams, string(value))
			}
		}
	})

	for _, v := range aParams {
		if value, err := strconv.ParseInt(v, 10, 64); err != nil {
			return output, err
		} else {
			output = append(output, value)
		}
	}

	return output, nil
}

func GetStringSliceFromRequestParam(c *fiber.Ctx, paramName string) []string {
	output := []string{}
	queryParams := c.Context().QueryArgs()

	queryParams.VisitAll(func(key, value []byte) {
		if string(key) == paramName {
			valueStr := string(value)
			if valueStr != "" {
				output = append(output, valueStr)
			}
		}
	})

	return output
}
