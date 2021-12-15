# Design Patterns
Add the following in the settings.json file:

    "gopls": {
        "experimentalWorkspaceModule": true
    }

Adding the above allows the user to to be able to open the design patterns folder in and not each module separately.
### **Modules** :
* design_patterns/creational_patterns/singleton
* design_patterns/creational_patterns/builder

### **How to run a module :** 
ex. Let's say we want to run Builder, navigate inside the builder folder. Then run the following :

    go test -v -run=Builder .

You should get something like this :

    === RUN   TestBuilderPattern
    --- PASS: TestBuilderPattern (0.00s)
    PASS
    ok      design_patterns/creational_patterns/builder

You can try changing the values being set for a certain type of vehicle to see that testing works:

    func (b *BusBuilder) SetSeats() BuildProcess {
	    b.v.Seats = 31 // instead of 30
	    return b
    }

You should get the following :

    === RUN   TestBuilderPattern
        builder_test.go:51: Seats on a bus must be 30 and they were 31
    --- FAIL: TestBuilderPattern (0.00s)
    FAIL
    FAIL    design_patterns/creational_patterns/builder     0.411s
    FAIL

