package arrays_and_string_manipulation

//My cake shop is so popular, I'm adding some tables and hiring wait staff so folks can have a cute sit-down cake-eating experience.
//
//I have two registers: one for take-out orders, and the other for the other folks eating inside the cafe. All the customer orders get combined into one list for the kitchen, where they should be handled first-come, first-served.
//
//Recently, some customers have been complaining that people who placed orders after them are getting their food first. Yikes—that's not good for business!
//
//To investigate their claims, one afternoon I sat behind the registers with my laptop and recorded:
//
//The take-out orders as they were entered into the system and given to the kitchen. (take_out_orders)
//The dine-in orders as they were entered into the system and given to the kitchen. (dine_in_orders)
//Each customer order (from either register) as it was finished by the kitchen. (served_orders)
//Given all three lists, write a function to check that my service is first-come, first-served. All food should come out in the same order customers requested it.
//
//We'll represent each customer order as a unique integer.
//
//As an example,
//
//  Take Out Orders: [1, 3, 5]
// Dine In Orders: [2, 4, 6]
//  Served Orders: [1, 2, 4, 6, 5, 3]
//would not be first-come, first-served, since order 3 was requested before order 5 but order 5 was served first.
//
//But,
//
//  Take Out Orders: [17, 8, 24]
// Dine In Orders: [12, 19, 2]
//  Served Orders: [17, 8, 12, 19, 24, 2]
//would be first-come, first-served.
//
//Note: Order numbers are arbitrary. They do NOT have to be in increasing order.
// Runtime: N
// Space:

// TODO: an alternative is to "throw out" the values we've seen in the slices instead of making new maps to keep track...
// See the other solution which comes from interview cake. By "throw out" they simply mean use index tracking....
// Before doing optimization, ask if they think it's worth it to save on memory space...
func ordersFirstComeFirstServe(takeOut []int, dineIn []int, servedOrder []int) bool {
	// Map the orders to their respective indexes so we can quickly check against them later...
	takeOutOrderToIndex := map[int]int{}
	for takeOutIndex, takeOutOrder := range takeOut {
		takeOutOrderToIndex[takeOutOrder] = takeOutIndex
	}

	dineInOrderToIndex := map[int]int{}
	for dineInIndex, dineInOrder := range dineIn {
		dineInOrderToIndex[dineInOrder] = dineInIndex
	}

	servedTakeoutIndex := 0
	servedDineinIndex := 0
	// For each order, check to see if matches its original list placement
	for _, o := range servedOrder {
		if _, ok := takeOutOrderToIndex[o]; ok {
			if takeOut[servedTakeoutIndex] != o {
				return false
			}
			servedTakeoutIndex++
		} else if _, ok := dineInOrderToIndex[o]; ok {
			if dineIn[servedDineinIndex] != o {
				return false
			}
			servedDineinIndex++
		}
	}
	return true
}

// This one is superior since it uses less space
func ordersFirstComeFirstServe2(takeOut []int, dineIn []int, servedOrder []int) bool {
	takeoutIndex := 0
	dineinIndex := 0

	takeoutOrderMaxIndex := len(takeOut) - 1
	dineInOrderMaxIndex := len(dineIn) - 1

	// For each order, check to see if matches its original list placement
	for _, order := range servedOrder {
		if takeoutIndex <= takeoutOrderMaxIndex && order == takeOut[takeoutIndex] {
			takeoutIndex++
		} else if dineinIndex <= dineInOrderMaxIndex && order == dineIn[dineinIndex] {
			dineinIndex++
		} else {
			//# If the current order in served_orders doesn't match the current
			//# order in take_out_orders or dine_in_orders, then we're not serving first-come,
			//# first-served.
			return false
		}
	}

	// What if there are extra orders in take_out_orders or dine_in_orders that don't appear in served_orders? Would our kitchen be first-come, first-served then?
	// Check for any extra orders at the end of take_out_orders or dine_in_orders
	if dineinIndex != len(dineIn) || takeoutIndex != len(takeOut) {
		return false
	}
	return true
}
