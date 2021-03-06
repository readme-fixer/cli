package types

import (
	"time"
)

const GINKGO_FOCUS_EXIT_CODE = 197

type SuiteSummary struct {
	SuiteDescription string
	SuiteSucceeded   bool
	SuiteID          string

	NumberOfSpecsBeforeParallelization int
	NumberOfTotalSpecs                 int
	NumberOfSpecsThatWillBeRun         int
	NumberOfPendingSpecs               int
	NumberOfSkippedSpecs               int
	NumberOfPassedSpecs                int
	NumberOfFailedSpecs                int
	RunTime                            time.Duration
}

type SpecSummary struct {
	ComponentTexts         []string
	ComponentCodeLocations []CodeLocation

	State           SpecState
	RunTime         time.Duration
	Failure         SpecFailure
	IsMeasurement   bool
	NumberOfSamples int
	Measurements    map[string]*SpecMeasurement

	CapturedOutput string
	SuiteID        string
}

type SetupSummary struct {
	ComponentType SpecComponentType
	CodeLocation  CodeLocation

	State   SpecState
	RunTime time.Duration
	Failure SpecFailure

	CapturedOutput string
	SuiteID        string
}

type SpecFailure struct {
	Message        string
	Location       CodeLocation
	ForwardedPanic interface{}

	ComponentIndex        int
	ComponentType         SpecComponentType
	ComponentCodeLocation CodeLocation
}

type SpecMeasurement struct {
	Name string
	Info interface{}

	Results []float64

	Smallest     float64
	Largest      float64
	Average      float64
	StdDeviation float64

	SmallestLabel string
	LargestLabel  string
	AverageLabel  string
	Units         string
}

type SpecState uint

const (
	SpecStateInvalid SpecState = iota

	SpecStatePending
	SpecStateSkipped
	SpecStatePassed
	SpecStateFailed
	SpecStatePanicked
	SpecStateTimedOut
)

type SpecComponentType uint

const (
	SpecComponentTypeInvalid SpecComponentType = iota

	SpecComponentTypeContainer
	SpecComponentTypeBeforeSuite
	SpecComponentTypeAfterSuite
	SpecComponentTypeBeforeEach
	SpecComponentTypeJustBeforeEach
	SpecComponentTypeAfterEach
	SpecComponentTypeIt
	SpecComponentTypeMeasure
)

type FlagType uint

const (
	FlagTypeNone FlagType = iota
	FlagTypeFocused
	FlagTypePending
)
