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

