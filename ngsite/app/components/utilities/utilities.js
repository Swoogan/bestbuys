'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.factory('parseCurrency', [function () {
  return function (amount) {
    if (amount === "") return 0;
  
    if (amount.charAt(0) == "$")   
      amount = amount.slice(1);
  
    amount = amount.replace(/<br>/g, "");
    amount = amount.replace(/,/g, "");
    var i = parseInt(amount);
  
    if (isNaN(i)) return "0";
  
    return i;
  }
}]);

