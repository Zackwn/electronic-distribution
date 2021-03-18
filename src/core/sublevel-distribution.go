package core

type SublevelDistribution = []DistributionSublevel

func runSublevelDistribution(electrons Electrons) SublevelDistribution {
	sublevelDistribution := SublevelDistribution{}

	index := 0

	for electrons > 0 {
		sublevel := sublevelsSequence[index]

		if sublevel.maxElectrons >= electrons {
			sublevelDistribution = append(sublevelDistribution, DistributionSublevel{
				value:          electrons,
				layer:          sublevel.layer,
				representation: sublevel.representation,
			})
			break
		}

		electrons = electrons - sublevel.maxElectrons
		sublevelDistribution = append(sublevelDistribution, DistributionSublevel{
			value:          sublevel.maxElectrons,
			layer:          sublevel.layer,
			representation: sublevel.representation,
		})
		index++
	}

	return sublevelDistribution
}

func formatSublevelDistribution(sublevelDistribution SublevelDistribution) string {
	result := ""
	for _, distributionSublevel := range sublevelDistribution {
		result += formatDistributionSublevel(distributionSublevel) + " "
	}
	return result
}
