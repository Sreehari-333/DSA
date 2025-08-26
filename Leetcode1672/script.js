var maximumWealth = function(accounts) {
    let richestAcc = 0;
    
    for(let i = 0; i<accounts.length;i++){
        let acc = 0;
        for(let j = 0; j<accounts[i].length;j++){
             acc += accounts[i][j]
            if (acc > richestAcc){
                richestAcc = acc
            }
        }
    }
    return richestAcc
};

let accounts = [[1,5],[7,3],[3,5]]
console.log(maximumWealth(accounts))