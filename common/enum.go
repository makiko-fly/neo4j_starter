package common

type EnumRelation int

const (
	EnumRelation_PRODUCES       EnumRelation = 1001 // company -> product
	EnumRelation_HAS_DOWNSTREAM              = 2001 // product -> product
	EnumRelation_HAS_CHILD                   = 2002 // product -> product
	EnumRelation_LISTED_AS                   = 3001 // company -> stock
	EnumRelation_WITHIN_CHAIN                = 4001 // product -> chain
)

var EnumRelationDefaultStr = "UNDEFINED"

func (self EnumRelation) String() string {
	switch self {
	case EnumRelation_PRODUCES:
		return "PRODUCES"
	case EnumRelation_HAS_DOWNSTREAM:
		return "HAS_DOWNSTREAM"
	case EnumRelation_HAS_CHILD:
		return "HAS_CHILD"
	case EnumRelation_LISTED_AS:
		return "LISTED_AS"
	case EnumRelation_WITHIN_CHAIN:
		return "WITHIN_CHAIN"
	default:
		return EnumRelationDefaultStr
	}
}

func (self EnumRelation) IsValid() bool {
	return self.String() != EnumRelationDefaultStr
}
