package reflection

import (
	"testing"

	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/require"
)

func TestGatherInfo(t *testing.T) {
	t.Run("Gather Info", func(t *testing.T) {
		type testing struct {
			Exported1 string `emi:"default.exported1"`
		}

		test := testing{}
		namespace := "default"

		varInfos, err := GatherInfo(namespace, &test)
		require.NoError(t, err)

		require.NotEmpty(t, varInfos)

		varInfo := varInfos[0]

		log.Debug().Interface("varInfos", varInfos).Send()

		require.EqualValues(t, "Exported1", varInfo.Name)
		require.EqualValues(t, "default.exported1", varInfo.Alt)
		require.EqualValues(t, "default.default.exported1", varInfo.Key)
		require.EqualValues(t, "emi:\"default.exported1\"", varInfo.Tags)
	})
}
