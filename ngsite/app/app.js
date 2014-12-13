'use strict';

// Declare app level module which depends on views, and components
angular.module('bestbuys', [
  'ngRoute',
  'bestbuys.game',
  'bestbuys.version'
]).
config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({redirectTo: '/game'});
}]).
directive('bbEditbox', function ($filter) {
  return {
    link: function ($scope, element, attrs) {
      /* On Focus */
      element.bind('focus', function () {
	element.addClass('shadow');
	element.html(parseCurrency(element.html()));
      });
      
      /* On Blur */
      element.bind('blur', function () {
	element.removeClass('shadow');
	element.html($filter('currency')(element.html(), undefined, 0));
      });
      
      /* On Keydown */
      element.bind('keydown', function(e) {
	var enter = e.which == 13;
	
	if (enter) {
	  e.preventDefault();
	  e.target.blur();	  
	}
      });
      
      /* On Keyup */
      element.bind('keyup', function(e) {
	var esc = e.which == 27;
	
	if (esc) {
	  document.execCommand("undo", false, null);
	  e.target.blur();	  
	} 
      });
    }    
  }
});



function parseCurrency(amount) {
  if (amount === "") return 0;
  
  if (amount.charAt(0) == "$")   
    amount = amount.slice(1);
  
  amount = amount.replace(/<br>/g, "");
  amount = amount.replace(/,/g, "");
  var i = parseInt(amount);
  
  if (isNaN(i)) return "0";
  
  return i;
}
