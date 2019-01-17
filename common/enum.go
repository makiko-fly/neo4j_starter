package common

type EnumRelation int

const (
	EnumRelation_PRODUCES       EnumRelation = 1001
	EnumRelation_HAS_DOWNSTREAM              = 1002
	EnumRelation_LISTED_AS                   = 1003
)

var EnumRelationDefaultStr = "UNDEFINED"

func (self EnumRelation) String() string {
	switch self {
	case EnumRelation_PRODUCES:
		return "PRODUCES"
	case EnumRelation_HAS_DOWNSTREAM:
		return "HAS_DOWNSTREAM"
	case EnumRelation_LISTED_AS:
		return "LISTED_AS"
	default:
		return EnumRelationDefaultStr
	}
}

func (self EnumRelation) IsValid() bool {
	return self.String() != EnumRelationDefaultStr
}
