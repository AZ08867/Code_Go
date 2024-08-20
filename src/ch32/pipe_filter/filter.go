package pipe_filter

// Request is the input data type of filter
type Request interface{}

// Response is the output data type of filter
type Response interface{}

// Filter is the interface of filter
type Filter interface {
	Process(data Request) (Response, error)
}
