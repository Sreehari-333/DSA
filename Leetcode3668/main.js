

let order = [3,1,2,5,4]
let friends = [1,3,4]


var recoverOrder = function(order, friends) {
    const set = new Set(friends);
    return order.filter((val) => set.has(val));
};

console.log(recoverOrder(order,friends))