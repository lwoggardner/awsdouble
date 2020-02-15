AWSDouble
==============

Generated Test Doubles for aws-sdk-go

Implements the complete <svc>iface.<SVC>API interfaces as test doubles

API methods can be stubbed, mocked, spied on or faked

### Default Fakes
Default fake implementations are supplied for
* *action*WithContext() , returns ctx.Canceled if the context is canceled at the time the call is invoked
* *action*Pages() - repeatedly calls *action* paginating as per real requests
* WaitUntil*waiter* - repeatedly calls the underlying *action* methods as per real waiters

This allows the basic *action* methods to be stubbed without regard for whether the system under test is using
  the context, pagination or waiter methods on the API.

### Default Return Values
Default return values for stubs on the basic *action* methods that return (**action*Output, error) is a pointer
 to an empty struct, and a nil error.
 
For methods beginning with "Poll" the default return values are randomly delayed by 10ms to simulate a long poll


### Timewarp

The double has a Timewarp function with the same signature as time.After.  The default timewarp is to sleep for
 exactly 10ms regardless of the duration called in.

This means a waiter that has default of 20s between calls, will execute with actually 10ms between calls.

To test a waiter is waiting for the expected amount of time, use something like fakeclock
