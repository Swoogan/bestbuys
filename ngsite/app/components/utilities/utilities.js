'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');


bestbuys.service('utilities', [function () {
  return {
    parseCurrency: function (amount) {
      if (amount === "") return 0;

      if (amount.charAt(0) == "$")   
	amount = amount.slice(1);
  
      amount = amount.replace(/<br>/g, "");
      amount = amount.replace(/,/g, "");
      var i = parseInt(amount);
  
      if (isNaN(i)) return "0";
  
      return i;
    },
    handleEscape: function (e) {
      var esc = e.which == 27;
  
      if (esc) {
	document.execCommand("undo", false, null);
	e.target.blur();
      } 
    },
    handleEnter: function (e) {
      var enter = e.which == 13;

      if (enter) {
	e.preventDefault();
	e.target.blur();
      }
    }
  }
}]);


/*
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

function handleEscape(e) {
  var esc = e.which == 27;
  
  if (esc) {
    document.execCommand("undo", false, null);
    e.target.blur();
  } 
}

function handleEnter(e) {
  var enter = e.which == 13;

  if (enter) {
    e.preventDefault();
    e.target.blur();
  }
}
*/
