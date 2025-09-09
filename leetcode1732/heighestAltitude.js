var largestAltitude = function(gain) {
    let altitude=0;
    let heighestAltitude=0
    for(let i=0 ;i<gain.length; i++){
        altitude+=gain[i];
        if(altitude > heighestAltitude){
            heighestAltitude = altitude;
        }
    }
    return heighestAltitude;
    
};