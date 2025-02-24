package config

import (
	"regexp"
	"strings"

	"github.com/helmfile/helmfile/pkg/maputil"
)

func NewCLIConfigImpl(g *GlobalImpl) error {
	re := regexp.MustCompile(`(?:,|^)([^\s=]+)=(['"][^'"]*['"]|[^,]+)`)
	optsSet := g.RawStateValuesSetString()
	if len(optsSet) > 0 {
		set := map[string]any{}
		for i := range optsSet {
			ops := re.FindAllStringSubmatch(optsSet[i], -1)
			for j := range ops {
				op := ops[j]
				k := maputil.ParseKey(op[1])
				v := op[2]

				maputil.Set(set, k, v, true)
			}
		}
		g.SetSet(set)
	}
	optsSet = g.RawStateValuesSet()
	if len(optsSet) > 0 {
		set := map[string]any{}
		for i := range optsSet {
			ops := parseCommaSeparatedString(optsSet[i])
			for j := range ops {
				op := strings.SplitN(ops[j], "=", 2)
				k := maputil.ParseKey(op[0])
				v := op[1]

				maputil.Set(set, k, v, false)
			}
		}
		g.SetSet(set)
	}

	return nil
}

func parseCommaSeparatedString(input string) []string {
	var (
		result  []string
		current strings.Builder
		escaped bool
	)

	for i := 0; i < len(input); i++ {
		char := input[i]

		if escaped {
			if char == ',' {
				current.WriteByte(',')
			} else {
				current.WriteByte('\\')
				current.WriteByte(char)
			}
			escaped = false
			continue
		}

		if char == '\\' {
			escaped = true
			continue
		}

		if char == ',' {
			result = append(result, current.String())
			current.Reset()
		} else {
			current.WriteByte(char)
		}
	}
	result = append(result, current.String())
	return result
}
