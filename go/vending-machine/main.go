// VendingMachine represents a vending machine with an infinite number of bins
// and keypads. The inventory of the bins is always changing as items are
// dispensed and replenished. At any given moment, zero or more users will
// activate a keypad. They will create a session with their requested item and a
// payment method.
//
// The vending machine must validate the requested items are available, and the
// user has enough funds to purchase the items. The machine must never dispense
// items if the user does not have the funds, nor may it charge a user for items
// it cannot dispense.
//
// The session is complete when the items are dispensed and the charge is
// complete. If the session cannot be completed, the session may be marked
// complete when all charges and dispense jobs have been canceled.
type VendingMachine struct {
	Dispenser
	PaymentService
	StateStore
}

// Reconcile will be called by the main event loop for each session repeatedly
// until a final status is returned. Multiple inflight sessions can exist at any
// given time and will be reconciled in an arbitrary order. The reconcile
// function should attempt to complete as much work as possible with each call
// and avoid blocking to ensure requests can be processed as concurrently as
// possible. Returning `false` will place the session back onto the work queue to
// be reconciled again later. Returning `true` will indicate there is no further
// work needed for the session, and it will not be reconciled again.  You may assume the session will be fully
// populated and immutable.
func (a *VendingMachine) Reconcile(session Session) bool {
	// TODO:
	panic("implement me")
}

// DO NOT EDIT ANY CODE BELOW THIS LINE

// Session represents the current state of input from a keypad.  Fields are immutable once set.
type Session struct {
	// ID is the unique identifier for the session.
	ID int64

	// Bin is the ID of the bin containing the item the user requested.
	Bin int

	// Quantity is the number of items the user requested.
	Quantity int

	// PricePerUnit is the price of a single item for the given bin.
	PricePerUnit int

	// PaymentInfo is the information required to charge the user's payment method.
	PaymentInfo PaymentInfo
}

// State tracks DispenseJobs and Charges associated with a given session
type State struct {
	DispenseJob
	Charge
}

// StateStore is a key value store that can persist a State for later retrieval.
type StateStore interface {
	// Get returns the State for the given session ID if one exists.
	Get(int64, *State)

	// Set stores the State for the given session ID.
	Set(int64, State)
}

// DispenseJob represents a request to dispense a quantity of a bin.
type DispenseJob int64

// DispenseJobStatus indicates the state of a DispenseJob.
type DispenseJobStatus int

const (
	// DispenseJobStatusPending the DispenseJob has not yet been processed by the Dispenser.
	DispenseJobStatusPending DispenseJobStatus = iota

	// DispenseJobStatusReserved the quantity requested is available and being held for
	// this job. This status will not change unless the items are dispensed or the
	// job is canceled.
	DispenseJobStatusReserved

	// DispenseJobStatusFailed the Dispenser was unable to dispense or reserve the quantity requested.  This status is final.
	DispenseJobStatusFailed

	// DispenseJobStatusCancelled the DispenseJob has been canceled and the reserved inventory has been freed. This status is final.
	DispenseJobStatusCancelled

	// DispenseJobStatusComplete the items were successfully dispensed to the user. This status is final.
	DispenseJobStatusComplete
)

// Dispenser manages the inventory of bins and physically dispenses items to users.
type Dispenser interface {
	// Reserve will request the given quantity of inventory reserved. This will prevent
	// the inventory from being allocated to other jobs unless the job is canceled.
	// If the requested quantity is not available or otherwise reserved, the job will
	// be marked as failed.
	Reserve(bin, quantity int) DispenseJob

	// Dispense will physically dispense items to the user. The system will dispense
	// the quantity reserved for this Job, If the quantity is not yet reserved, the
	// system will dispense zero items and mark the job as complete.
	Dispense(DispenseJob)

	// Cancel will cancel the job and free the reserved inventory.
	Cancel(DispenseJob)

	// GetStatus will return the current status of the job.
	GetStatus(DispenseJob) DispenseJobStatus
}

// PaymentInfo represents the information required to charge a user's payment method.
type PaymentInfo interface{}

// Charge represents a charge that has been placed on a user's payment method.
type Charge int64

// ChargeStatus represents the status of a charge.
type ChargeStatus int

const (
	// ChargeStatusRequested the charge has been requested but not yet authorized.
	ChargeStatusRequested ChargeStatus = iota

	// ChargeStatusAuthorized the charge has been authorized and the requested funds are available on the user's account.  The funds will be held until the charge is completed.
	ChargeStatusAuthorized

	// ChargeStatusDeclined the charge was declined due to insufficient funds in the account.  This status is final.
	ChargeStatusDeclined

	// ChargeStatusCanceled the charge was successfully canceled and the funds are no longer held on the user's account.  This status is final.
	ChargeStatusCanceled

	// ChargeStatusComplete the charge was successfully completed and the funds have been transferred to the vendor.  This status is final.
	ChargeStatusComplete
)

// PaymentService is responsible for transferring funds from a user's account to a vendor.
type PaymentService interface {
	// Authorize will attempt to place a hold for the requested funds.  If the hold is authorized, the funds will be held indefinitely.
	Authorize(amount int, paymentInfo PaymentInfo) Charge

	// Complete will transfer the held funds to the vendor. If the Charge is not yet authorized, the transaction will fail.
	Complete(Charge)

	// Cancel will cancel the charge and release the held funds.
	Cancel(Charge)

	// GetStatus will return the current status of the charge.
	GetStatus(Charge) ChargeStatus
}
