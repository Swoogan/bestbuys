'use strict';

// Declare app level module which depends on views, and components
var bestbuyApp = angular.module('bestbuys', [
  'ngRoute',
  'ngResource',
  'ngAnimate',
]);

bestbuyApp.config(['$routeProvider', function($routeProvider) {
  $routeProvider.otherwise({redirectTo: '/game'});
}]);



bestbuyApp.filter('capitalize', function() {
  return function(input) {
    return (!!input) ? input.charAt(0).toUpperCase() + input.slice(1).toLowerCase() : '';
  }
});



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
