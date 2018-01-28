//
// This is only a SKELETON file for the "Leap" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Hamming = function(input) {
  this.year = input;
};

Hamming.prototype.compute = function(seq1, seq2) {
  if(seq1.length != seq2.length) {
    throw new Error('DNA strands must be of equal length.')
  }
  var distance = 0
  for (var i = 0, len = seq1.length; i < len; i++) {
    if(seq1[i] != seq2[i]){
      distance += 1;
    }
  }
  return distance;
};



module.exports = Hamming;
