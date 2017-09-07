package teardownLogic

/*TeardownLogic will be used to cancel subscription for given Observable*/
type TeardownLogic func()

/*AnonymousSubscription is simply an object with unsubscribe method on it*/
type AnonymousSubscription interface {
	Unsubscribe()
}

/*FromFunc convert func to TeardownLogic*/
func FromFunc(fn func()) TeardownLogic {
	return fn
}

/*FromSubscription convert AnonymousSubscription to TeardownLogic**/
func FromSubscription(subscription AnonymousSubscription) TeardownLogic {
	return func() { subscription.Unsubscribe() }
}
