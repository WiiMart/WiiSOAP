package main

// LimitKinds represents various limits applied to the current ticket.
type LimitKinds int

const (
	// PR is presumably "purchased".
	PR LimitKinds = 0
	TR            = 1
	DR            = 2
	SR            = 3
	LR            = 4
	AT            = 10000
)

// LimitStruct returns a Limits struct filled for the given kind.
func LimitStruct(kind LimitKinds) Limits {
	names := map[LimitKinds]string{
		PR: "PR",
		TR: "TR",
		DR: "DR",
		SR: "SR",
		LR: "LR",
		AT: "AT",
	}

	return Limits{
		Limits:    kind,
		LimitKind: names[kind],
	}
}

// DeviceStatus represents the various statuses a device may have.
//
// These values do not appear to be directly checked by the client within the
// Wii Shop Channel, and are a generic string. We could utilize any value we wish.
// However, titles utilizing DLCs appear to check the raw values.
// For this reason, we mirror values from Nintendo.
type DeviceStatus string

const (
	DeviceStatusRegistered   = "R"
	DeviceStatusUnregistered = "U"
)
