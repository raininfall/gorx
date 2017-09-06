package rx

/*TeardownLogic will be used to cancel subscription for given Observable*/
type TeardownLogic func()

/*AnonymousSubscription is simply an object with unsubscribe method on it*/
type AnonymousSubscription interface {
	Unsubscribe()
}

/*CreateTeardownLogicFromFunc convert func to TeardownLogic*/
func CreateTeardownLogicFromFunc(fn func()) TeardownLogic {
	return fn
}

/*CreateTeardownLogicFrom convert AnonymousSubscription to TeardownLogic**/
func CreateTeardownLogicFrom(subscription AnonymousSubscription) TeardownLogic {
	return func() { subscription.Unsubscribe() }
}
