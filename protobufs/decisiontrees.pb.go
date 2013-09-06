// Code generated by protoc-gen-go.
// source: protobufs/decisiontrees.proto
// DO NOT EDIT!

package protobufs

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type LossFunction int32

const (
	LossFunction_LOGIT                    LossFunction = 1
	LossFunction_LEAST_ABSOLUTE_DEVIATION LossFunction = 2
)

var LossFunction_name = map[int32]string{
	1: "LOGIT",
	2: "LEAST_ABSOLUTE_DEVIATION",
}
var LossFunction_value = map[string]int32{
	"LOGIT":                    1,
	"LEAST_ABSOLUTE_DEVIATION": 2,
}

func (x LossFunction) Enum() *LossFunction {
	p := new(LossFunction)
	*p = x
	return p
}
func (x LossFunction) String() string {
	return proto.EnumName(LossFunction_name, int32(x))
}
func (x LossFunction) MarshalJSON() ([]byte, error) {
	return json.Marshal(x.String())
}
func (x *LossFunction) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(LossFunction_value, data, "LossFunction")
	if err != nil {
		return err
	}
	*x = LossFunction(value)
	return nil
}

type Feature struct {
	Feature          *int64   `protobuf:"varint,1,opt,name=feature" json:"feature,omitempty"`
	Value            *float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Feature) Reset()         { *m = Feature{} }
func (m *Feature) String() string { return proto.CompactTextString(m) }
func (*Feature) ProtoMessage()    {}

func (m *Feature) GetFeature() int64 {
	if m != nil && m.Feature != nil {
		return *m.Feature
	}
	return 0
}

func (m *Feature) GetValue() float64 {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return 0
}

type TreeNode struct {
	// feature to split on
	Feature *int64 `protobuf:"varint,1,opt,name=feature" json:"feature,omitempty"`
	// value to split on
	SplitValue       *float64  `protobuf:"fixed64,2,opt,name=splitValue" json:"splitValue,omitempty"`
	Left             *TreeNode `protobuf:"bytes,3,opt,name=left" json:"left,omitempty"`
	Right            *TreeNode `protobuf:"bytes,4,opt,name=right" json:"right,omitempty"`
	LeafValue        *float64  `protobuf:"fixed64,5,opt,name=leafValue" json:"leafValue,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *TreeNode) Reset()         { *m = TreeNode{} }
func (m *TreeNode) String() string { return proto.CompactTextString(m) }
func (*TreeNode) ProtoMessage()    {}

func (m *TreeNode) GetFeature() int64 {
	if m != nil && m.Feature != nil {
		return *m.Feature
	}
	return 0
}

func (m *TreeNode) GetSplitValue() float64 {
	if m != nil && m.SplitValue != nil {
		return *m.SplitValue
	}
	return 0
}

func (m *TreeNode) GetLeft() *TreeNode {
	if m != nil {
		return m.Left
	}
	return nil
}

func (m *TreeNode) GetRight() *TreeNode {
	if m != nil {
		return m.Right
	}
	return nil
}

func (m *TreeNode) GetLeafValue() float64 {
	if m != nil && m.LeafValue != nil {
		return *m.LeafValue
	}
	return 0
}

type Forest struct {
	Trees            []*TreeNode `protobuf:"bytes,1,rep,name=trees" json:"trees,omitempty"`
	XXX_unrecognized []byte      `json:"-"`
}

func (m *Forest) Reset()         { *m = Forest{} }
func (m *Forest) String() string { return proto.CompactTextString(m) }
func (*Forest) ProtoMessage()    {}

func (m *Forest) GetTrees() []*TreeNode {
	if m != nil {
		return m.Trees
	}
	return nil
}

type SplittingConstraints struct {
	MaximumLevels              *int64   `protobuf:"varint,1,opt,name=maximumLevels" json:"maximumLevels,omitempty"`
	MinimumAverageGain         *float64 `protobuf:"fixed64,2,opt,name=minimumAverageGain" json:"minimumAverageGain,omitempty"`
	MinimumSamplesAtLeaf       *int64   `protobuf:"varint,3,opt,name=minimumSamplesAtLeaf" json:"minimumSamplesAtLeaf,omitempty"`
	FeaturesConsideredFraction *float64 `protobuf:"fixed64,4,opt,name=featuresConsideredFraction" json:"featuresConsideredFraction,omitempty"`
	XXX_unrecognized           []byte   `json:"-"`
}

func (m *SplittingConstraints) Reset()         { *m = SplittingConstraints{} }
func (m *SplittingConstraints) String() string { return proto.CompactTextString(m) }
func (*SplittingConstraints) ProtoMessage()    {}

func (m *SplittingConstraints) GetMaximumLevels() int64 {
	if m != nil && m.MaximumLevels != nil {
		return *m.MaximumLevels
	}
	return 0
}

func (m *SplittingConstraints) GetMinimumAverageGain() float64 {
	if m != nil && m.MinimumAverageGain != nil {
		return *m.MinimumAverageGain
	}
	return 0
}

func (m *SplittingConstraints) GetMinimumSamplesAtLeaf() int64 {
	if m != nil && m.MinimumSamplesAtLeaf != nil {
		return *m.MinimumSamplesAtLeaf
	}
	return 0
}

func (m *SplittingConstraints) GetFeaturesConsideredFraction() float64 {
	if m != nil && m.FeaturesConsideredFraction != nil {
		return *m.FeaturesConsideredFraction
	}
	return 0
}

type PruningConstraints struct {
	CrossValidationFolds *int64 `protobuf:"varint,1,opt,name=crossValidationFolds" json:"crossValidationFolds,omitempty"`
	XXX_unrecognized     []byte `json:"-"`
}

func (m *PruningConstraints) Reset()         { *m = PruningConstraints{} }
func (m *PruningConstraints) String() string { return proto.CompactTextString(m) }
func (*PruningConstraints) ProtoMessage()    {}

func (m *PruningConstraints) GetCrossValidationFolds() int64 {
	if m != nil && m.CrossValidationFolds != nil {
		return *m.CrossValidationFolds
	}
	return 0
}

type ForestConfig struct {
	SplittingConstraints *SplittingConstraints `protobuf:"bytes,1,opt,name=splittingConstraints" json:"splittingConstraints,omitempty"`
	LossFunction         *LossFunction         `protobuf:"varint,2,opt,name=lossFunction,enum=protobufs.LossFunction" json:"lossFunction,omitempty"`
	NumWeakLearners      *int64                `protobuf:"varint,3,opt,name=numWeakLearners" json:"numWeakLearners,omitempty"`
	XXX_unrecognized     []byte                `json:"-"`
}

func (m *ForestConfig) Reset()         { *m = ForestConfig{} }
func (m *ForestConfig) String() string { return proto.CompactTextString(m) }
func (*ForestConfig) ProtoMessage()    {}

func (m *ForestConfig) GetSplittingConstraints() *SplittingConstraints {
	if m != nil {
		return m.SplittingConstraints
	}
	return nil
}

func (m *ForestConfig) GetLossFunction() LossFunction {
	if m != nil && m.LossFunction != nil {
		return *m.LossFunction
	}
	return 0
}

func (m *ForestConfig) GetNumWeakLearners() int64 {
	if m != nil && m.NumWeakLearners != nil {
		return *m.NumWeakLearners
	}
	return 0
}

func init() {
	proto.RegisterEnum("protobufs.LossFunction", LossFunction_name, LossFunction_value)
}
