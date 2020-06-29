package server

import "errors"
import "github.com/harishb2k/gox-base"

// This is a helper interface to extract info from request e.g. extract header, query param etc
type RequestExtractor interface {
	// Get value from headers
	GetIntHeader(name string) (value int, err error)
	GetStringHeader(name string) (value string, err error)
	GetBoolHeader(name string) (value bool, err error)

	// Get value from path param
	GetIntPathParam(name string) (value int, err error)
	GetStringPathParam(name string) (value string, err error)
	GetBoolPathParam(name string) (value bool, err error)
	GetFloatPathParam(name string) (value float64, err error)

	// Get value from query param
	GetIntQueryParam(name string) (value int, err error)
	GetFloatQueryParam(name string) (value float64, err error)
	GetStringQueryParam(name string) (value string, err error)
	GetBoolQueryParam(name string) (value bool, err error)

	// Get value array from query param
	GetIntQueryParams(name string) (value []int, err error)
	GetFloatQueryParams(name string) (value []float64, err error)
	GetStringQueryParams(name string) (value []string, err error)
	GetBoolQueryParams(name string) (value []bool, err error)
}

type DefaultRequestExtractor struct {
	Request *Request
}

func getFirstInt(strings []string) (value int, err error) {
	if intArray, e := gox.StrArrayToIntArray(strings); e == nil && intArray != nil && len(intArray) > 0 {
		return intArray[0], nil
	} else if e != nil {
		return 0, e
	}
	return 0, errors.New("error in GetIntHeader")
}

func getFirstBool(strings []string) (value bool, err error) {
	if intArray, e := gox.StrArrayToBoolArray(strings); e == nil && intArray != nil && len(intArray) > 0 {
		return intArray[0], nil
	} else if e != nil {
		return false, e
	}
	return false, errors.New("error in GetIntHeader")
}

func getFirstFloat(strings []string) (value float64, err error) {
	if intArray, e := gox.StrArrayToFloatArray(strings); e == nil && intArray != nil && len(intArray) > 0 {
		return intArray[0], nil
	} else if e != nil {
		return 0.0, e
	}
	return 0.0, errors.New("error in GetIntHeader")
}

func getFirstString(strings []string) (value string, err error) {
	if strings != nil && len(strings) > 0 {
		return strings[0], nil
	}
	return "", errors.New("error in GetStringHeader")
}

func (re *DefaultRequestExtractor) GetIntHeader(name string) (value int, err error) {
	return getFirstInt(re.Request.HttpRequest.Header[name])
}

func (re *DefaultRequestExtractor) GetStringHeader(name string) (value string, err error) {
	return getFirstString(re.Request.HttpRequest.Header[name])
}

func (re *DefaultRequestExtractor) GetBoolHeader(name string) (value bool, err error) {
	return getFirstBool(re.Request.HttpRequest.Header[name])
}

func (re *DefaultRequestExtractor) GetIntPathParam(name string) (value int, err error) {
	panic("implement me")
}

func (re *DefaultRequestExtractor) GetStringPathParam(name string) (value string, err error) {
	panic("implement me")
}

func (re *DefaultRequestExtractor) GetBoolPathParam(name string) (value bool, err error) {
	panic("implement me")
}

func (re *DefaultRequestExtractor) GetFloatPathParam(name string) (value float64, err error) {
	panic("implement me")
}

func (re *DefaultRequestExtractor) GetIntQueryParam(name string) (value int, err error) {
	return getFirstInt(re.Request.HttpRequest.Form[name])
}

func (re *DefaultRequestExtractor) GetFloatQueryParam(name string) (value float64, err error) {
	panic("implement me")
}

func (re *DefaultRequestExtractor) GetStringQueryParam(name string) (value string, err error) {
	return getFirstString(re.Request.HttpRequest.Form[name])
}

func (re *DefaultRequestExtractor) GetBoolQueryParam(name string) (value bool, err error) {
	return getFirstBool(re.Request.HttpRequest.Form[name])
}

func (re *DefaultRequestExtractor) GetIntQueryParams(name string) (value []int, err error) {
	return gox.StrArrayToIntArray(re.Request.HttpRequest.Form[name])
}

func (re *DefaultRequestExtractor) GetFloatQueryParams(name string) (value []float64, err error) {
	return gox.StrArrayToFloatArray(re.Request.HttpRequest.Form[name])
}

func (re *DefaultRequestExtractor) GetStringQueryParams(name string) (value []string, err error) {
	return re.Request.HttpRequest.Form[name], nil
}

func (re *DefaultRequestExtractor) GetBoolQueryParams(name string) (value []bool, err error) {
	return gox.StrArrayToBoolArray(re.Request.HttpRequest.Form[name])
}
