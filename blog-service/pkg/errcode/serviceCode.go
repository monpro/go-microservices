package errcode

var (
	ErrorGetTagListFail = NewError(20010001, "get tag list service fail")
	ErrorCreateTagFail  = NewError(20010002, "create tag service fail ")
	ErrorUpdateTagFail  = NewError(20010003, "update tag service fail")
	ErrorDeleteTagFail  = NewError(20010004, "delete tag service fail")
	ErrorCountTagFail   = NewError(20010005, "count tag service fail")
)
