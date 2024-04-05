package namegen

import "fmt"

type NameGeneratorParameters interface {
	GetAllowDuplicates() bool
	GetMaximumSampleAttempts() int
}

func newName(parameters NameGeneratorParameters, takenNames map[string]int, generationFunc func() string) (string, error) {
	if parameters.GetAllowDuplicates() {
		return newNameWithDuplicates(takenNames, generationFunc), nil
	}
	return newNameWithoutDuplicates(takenNames, parameters.GetMaximumSampleAttempts(), generationFunc)
}

func newNameWithDuplicates(takenNames map[string]int, generationFunc func() string) string {
	originalNewName := generationFunc()
	finalNewName := originalNewName
	count := takenNames[originalNewName]
	if count > 0 {
		finalNewName = fmt.Sprintf("%s%d", originalNewName, count)
	}
	takenNames[originalNewName] += 1
	return finalNewName
}

func newNameWithoutDuplicates(takenNames map[string]int, maxSamples int, generationFunc func() string) (string, error) {
	attemptCount := 0
	for attemptCount < maxSamples {
		potentialName := generationFunc()
		if _, ok := takenNames[potentialName]; ok {
			attemptCount += 1
			continue
		}

		takenNames[potentialName] = 1
		return potentialName, nil
	}

	return "", fmt.Errorf("failed to find an unused name after %d attempts. total taken names: %d",
		attemptCount, len(takenNames))
}
