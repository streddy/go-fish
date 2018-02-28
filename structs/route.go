package structs

// Route is a struct that contains information about a given route to be hit
type Route struct {
	Url										string
	Method								string
	Headers								string
	RequestBody 					string
	MandatoryDependencies	[]Route
	LikelyDependencies		[]Route
	Samples								int
}