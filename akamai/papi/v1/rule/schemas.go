package rule

type CriteriaMustSatisfyValue string

const (
	CriteriaMustSatisfyAll CriteriaMustSatisfyValue = "all"
	CriteriaMustSatisfyAny CriteriaMustSatisfyValue = "any"
)

// Tree is a collection of property rules.
type Tree struct {
	// RuleFormat represents current rule tree format.
	RuleFormat      string       `json:"ruleFormat"`

	// Rule specifies the executable logic to apply to cached edge content.
	Rule            *Rule        `json:"rules"`

	// Errors represents an array of validate errors returned for a rule.
	Errors          []*Error `json:"errors,omitempty"`
}

// TopLevelRuleOptions is a set of flags that could be applied to the top-level rule object.
type TopLevelRuleOptions struct {
	// IsSecure flag when enabled, serves the property’s content over SSL.
	// Doing so may enable additional rule behaviors.
	IsSecure *bool `json:"is_secure,omitempty"`
}

// Rule represents a property rule resource.
type Rule struct {
	// Name is a description of the rule. The top-level rule must be named default.
	Name                string                   `json:"name"`

	// Criteria is a series of criteria objects, which form a set of logical tests
	// that may prevent the rule’s behaviors from executing. Behavior and criteria
	// objects are structured identically.
	Criteria            []*Criteria              `json:"criteria,omitempty"`

	// Behaviors is a series of behavior objects, which execute after the set of criteria
	// evaluates. Behavior and criteria objects are structured identically.
	// Optional on nested rules.
	Behaviors           []*Behavior              `json:"behaviors,omitempty"`

	// Children is a series of nested Rule objects that execute after
	// the current rule’s criteria and behaviors.
	Children            []*Rule                  `json:"children,omitempty"`

	// Comment is a descriptive comment to help you track the rule’s function.
	Comment             string                   `json:"comment,omitempty"`

	// CriteriaLocked flag that prohibits any modifications to criteria objects.
	// Read-only.
	CriteriaLocked      *bool                     `json:"criteriaLocked,omitempty"`

	// CriteriaMustSatisfy field must be set if you define more than one,
	// the rule needs to set whether to match any or all criteria.
	CriteriaMustSatisfy CriteriaMustSatisfyValue `json:"criteriaMustSatisfy,omitempty"`

	// UUID is a data hash that indicates the rule contains at least one advanced
	// behavior or criteria, each identified with its own uuid member. When this member
	// is present on the rule, you may not remove any advanced features it contains,
	// or change their sequence. You must provide the same uuid value when modifying the rule tree.
	UUID                string                   `json:"uuid,omitempty"`

	// Variables represent an array of variables.
	Variables           []*Variable              `json:"variables,omitempty"`

	// AdvancedOverride is a block of post-processing XML metadata that your Akamai representative
	// can apply on your behalf. This must be configured separately for each property.
	AdvancedOverride    string                   `json:"advancedOverride,omitempty"`

	// Options is a set of flags that could be applied to the top-level rule object.
	Options *TopLevelRuleOptions `json:"options,omitempty"`
}

// Criteria represents a rule criteria resource
type Criteria struct {
	// Name is the name of the criteria.
	Name    string      `json:"name"`

	// Options is a variable set of options that configure each behavior.
	Options OptionValue `json:"options"`

	// UUID is a data hash that indicates an advanced behavior. When present,
	// you may not modify the contents of the behavior, and must provide the
	// same uuid value when modifying the rule tree.
	UUID    string      `json:"uuid,omitempty"`

	// Locked flag indicates a behavior or criteria that you arrange with your
	// Akamai representative to lock so that you can’t modify it, typically to protect
	// sensitive data from erroneous changes.
	Locked  *bool        `json:"locked,omitempty"`
}

// Behavior represents a rule behavior resource
type Behavior struct {
	// Name is the name of the behavior.
	Name    string      `json:"name"`

	// Options is a variable set of options that configure each behavior.
	Options OptionValue `json:"options"`

	// UUID is a data hash that indicates an advanced behavior. When present,
	// you may not modify the contents of the behavior, and must provide the
	// same uuid value when modifying the rule tree.
	UUID    string      `json:"uuid,omitempty"`

	// Locked flag indicates a behavior or criteria that you arrange with your
	// Akamai representative to lock so that you can’t modify it, typically to protect
	// sensitive data from erroneous changes.
	Locked  *bool        `json:"locked,omitempty"`
}

// OptionValue represents a generic option value
//
// OptionValue is a map with string keys, and any
// type of value. You can nest OptionValues as necessary
// to create more complex values.
type OptionValue map[string]interface{}

// Variable represents a rule variable.
type Variable struct {
	// Name is the underlying root name of the variable, which must be unique within
	// the set of variables.
	Name        string `json:"name"`

	// Value initializes a default value. Omitting this member initializes
	// the variable with an empty string.
	Value       string `json:"value"`

	// Description is a text to keep track of how you use each variable.
	Description string `json:"description"`

	// Hidden flag when enabled, the variable is suppressed from session response headers,
	// often used to test content.
	Hidden      bool   `json:"hidden"`

	// Sensitive flag when enabled, the variable is suppressed from session responses as
	// hidden ones, but it also can’t be invoked within any behaviors that assign values
	// to cookies or response headers. You also can’t assign a sensitive variable to another
	// one that is not sensitive, and you can’t add it to custom logging fields.
	// Use this more stringent option for any personally identifiable information, typically
	// after initially testing on the staging network.
	Sensitive   bool   `json:"sensitive"`
}

// Error represents an validate error returned for a rule.
type Error struct {
	// Type of the error.
	Type         string `json:"type"`

	// Title of the error.
	Title        string `json:"title"`

	// Detail is a description of the error.
	Detail       string `json:"detail"`

	// Instance reflects a URL where error was occurred.
	Instance     string `json:"instance"`

	// BehaviorName reflects a name of the behavior error
	// is related with.
	BehaviorName string `json:"behaviorName"`
}
