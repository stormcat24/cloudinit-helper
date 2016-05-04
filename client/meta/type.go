package meta

import "regexp"

type AvailabilityZone struct {
	Name string
}

func (z *AvailabilityZone) GetRegion() string {

	regionPattern := regexp.MustCompile(`^(.+-[0-9])[a-z]$`)
	tokens := regionPattern.FindStringSubmatch(z.Name)
	if len(tokens) == 2 {
		return tokens[1]
	}

	return ""
}