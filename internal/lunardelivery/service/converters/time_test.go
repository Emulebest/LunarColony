package converters_test

import (
	"LunarColony/internal/lunardelivery/service/converters"
	"github.com/stretchr/testify/assert"
	"testing"
)

const earthTime = "2021-11-26 11:28:45 UTC"

func TestConvertEarthToMoonTime(t *testing.T) {
	parsedTime, err := converters.StringToUTC(earthTime)
	assert.NoError(t, err)
	lunarElapsedSeconds := converters.LunarElapsedSeconds(parsedTime)
	lunarSeconds := converters.ConvertEarthToMoonSec(lunarElapsedSeconds)
	lst := converters.FormatLST(lunarSeconds)
	assert.Equal(t, "54-12-16 delta 07:18:10", lst)
}
