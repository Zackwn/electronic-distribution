package core

import "fmt"

type DistributionSublevel struct {
	value          Electrons
	representation rune
	layer          int
}

func formatDistributionSublevel(distributionSublevel DistributionSublevel) string {
	return fmt.Sprintf(
		"%v%v%v",
		distributionSublevel.layer,
		string(distributionSublevel.representation),
		distributionSublevel.value,
	)
}
