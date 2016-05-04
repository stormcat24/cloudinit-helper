package meta

import "regexp"

type AvailabilityZone struct {
	Name string
}

func (z *AvailabilityZone) GetRegion() string {

	regionPattern := regexp.MustCompile(`^(.+)-(.+)$`)
	tokens := regionPattern.FindStringSubmatch(z.Name)
	if len(tokens) == 3 {
		return tokens[1]
	}

	return ""
}