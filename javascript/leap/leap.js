//
// This is only a SKELETON file for the "Leap" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Year = function(input) {
  this.year = input;
};

Year.prototype.evelyDividedBy = function(input) {
  return (this.year % input) == 0;
};

Year.prototype.isLeap = function() {
  if(!this.evelyDividedBy(4)){
    return false
  }

  if(this.evelyDividedBy(400)) {
    return true;
  }

  if(this.evelyDividedBy(100)) {
    return false;
  } else {
    return true;
  }
};



module.exports = Year;
