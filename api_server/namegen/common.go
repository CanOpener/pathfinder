package namegen

type NameGeneratorParameters interface {
	GetAllowDuplicates() bool
	GetMaximumSampleAttempts() int
}
