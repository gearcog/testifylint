package main

import (
	"io"
	"text/template"

	"github.com/Antonboom/testifylint/internal/checkers"
)

const header = "// Code generated by testifylint/internal/testgen. DO NOT EDIT."

var fm = template.FuncMap{
	"NewAssertionExpander": NewAssertionExpander,

	// NOTE(a.telyshev): Sub-template multiple arguments problem (example in BaseTestsGenerator):
	// - https://stackoverflow.com/a/18276262
	// - https://dev.to/moniquelive/passing-multiple-arguments-to-golang-templates-16h8
	"arr": func(elements ...any) []any { return elements },
}

type Executor interface {
	Execute(wr io.Writer, data any) error
}

// TestsGenerator provides test and golden file templates and data for them.
// The resulting (executed) files can be used with analysistest.RunWithSuggestedFixes.
//
// GoldenTemplate can be omitted if autocorrection is not possible.
// For example, non-obvious asserts that require human intervention, such as checkers.FloatCompare.
type TestsGenerator interface {
	TemplateData() any
	ErroredTemplate() Executor
	GoldenTemplate() Executor
}

// CheckerTestsGenerator is a TestsGenerator for specific Checker.
// One checker must be covered by one generator.
type CheckerTestsGenerator interface {
	TestsGenerator
	Checker() checkers.Checker
}
