
let s = "Hello how are you Contestant"
let k = 4

var truncateSentence = function(s, k) {
    return s.split(" ").slice(0,k).join(" ")
};

console.log(truncateSentence(s,k))