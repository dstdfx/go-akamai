package activation

import "time"

// Value is used to create an "enum" of possible Activation.Type values
type Value string

// NetworkValue is used to create an "enum" of possible Activation.Network values
type NetworkValue string

// StatusValue is used to create an "enum" of possible Activation.Status values
type StatusValue string

const (
	// TypeActivate Activation.Type value ACTIVATE
	TypeActivate Value = "ACTIVATE"
	// TypeDeactivate Activation.Type value DEACTIVATE
	TypeDeactivate Value = "DEACTIVATE"

	// NetworkProduction Activation.Network value PRODUCTION
	NetworkProduction NetworkValue = "PRODUCTION"
	// NetworkStaging Activation.Network value STAGING
	NetworkStaging NetworkValue = "STAGING"

	// StatusActive Activation.Status value ACTIVE
	StatusActive StatusValue = "ACTIVE"
	// StatusInactive Activation.Status value INACTIVE
	StatusInactive StatusValue = "INACTIVE"
	// StatusPending Activation.Status value PENDING
	StatusPending StatusValue = "PENDING"
	// StatusZone1 Activation.Status value ZONE_1
	StatusZone1 StatusValue = "ZONE_1"
	// StatusZone2 Activation.Status value ZONE_2
	StatusZone2 StatusValue = "ZONE_2"
	// StatusZone3 Activation.Status value ZONE_3
	StatusZone3 StatusValue = "ZONE_3"
	// StatusAborted Activation.Status value ABORTED
	StatusAborted StatusValue = "ABORTED"
	// StatusFailed Activation.Status value FAILED
	StatusFailed StatusValue = "FAILED"
	// StatusDeactivated Activation.Status value DEACTIVATED
	StatusDeactivated StatusValue = "DEACTIVATED"
	// StatusPendingDeactivation Activation.Status value PENDING_DEACTIVATION
	StatusPendingDeactivation StatusValue = "PENDING_DEACTIVATION"
	// StatusNew Activation.Status value NEW
	StatusNew StatusValue = "NEW"
)

// Activation represents a property activation resource
type Activation struct {
	// ID is the activation’s unique identifier.
	ID                  string            `json:"activationId,omitempty"`

	// Type field. Either ACTIVATE or DEACTIVATE. The default is ACTIVATE.
	// Any new activation automatically deactivates the current activation.
	// Note that if you were to POST a DEACTIVATE type on an active property,
	// it would no longer serve any traffic. You would need to modify (de-Akamaize) your
	// DNS configuration and specify a different way to field the traffic.
	Type Value `json:"activationType,omitempty"`

	// AcknowledgeWarnings is a list of msg_-prefixed string IDs acknowledging any warnings
	// noted in responses to previous activation requests.
	AcknowledgeWarnings []string          `json:"acknowledgeWarnings,omitempty"`

	// FastPush flag that enables a fast metadata push when activating a new property,
	// true by default.
	FastPush            bool              `json:"fastPush,omitempty"`

	// IgnoreHTTPErrors flag that ignores any HTTP errors when pushing fast metadata activation,
	// true by default.
	IgnoreHTTPErrors    bool              `json:"ignoreHttpErrors,omitempty"`

	// PropertyName is the name of the property targeted with activation.
	PropertyName        string            `json:"propertyName,omitempty"`

	// PropertyID identifies property targeted with activation.
	PropertyID          string            `json:"propertyId,omitempty"`

	// PropertyVersion is the property version targeted with activation.
	// Once activated, you can no longer modify that version of the property.
	PropertyVersion     int               `json:"propertyVersion"`

	// Network represents a network to activate on, either STAGING or PRODUCTION.
	Network             NetworkValue      `json:"network"`

	// Status represents the activation’s status: ACTIVE if currently serving traffic;
	// INACTIVE if another activation has superseded this one;
	// NEW, PENDING, ZONE_1, ZONE_2, or ZONE_3 if not yet active;
	// ABORTED if the client followed up with a DELETE request in time;
	// FAILED if the activation causes a range of edge network errors that may cause a
	// fallback to the previous activation;
	// PENDING_DEACTIVATION or DEACTIVATED when the activationType is DEACTIVATE to no longer serve traffic.
	Status              StatusValue       `json:"status,omitempty"`

	// SubmitDate is a date stamp marking when the activation initiated.
	SubmitDate          *time.Time            `json:"submitDate,omitempty"`

	// UpdateDate is a date stamp marking when the status last changed.
	UpdateDate          *time.Time            `json:"updateDate,omitempty"`

	// Note assigns a log message to the activation request.
	Note                string            `json:"note,omitempty"`

	// NotifyEmails is a list of email address strings to notify when the activation status changes.
	NotifyEmails        []string          `json:"notifyEmails"`
}
