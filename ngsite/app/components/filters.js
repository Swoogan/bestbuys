'use strict';

/*global angular*/

var bestbuys = angular.module('bestbuys');

bestbuys.filter('capitalize', function() {
  return function(input) {
    return (!!input) ? input.charAt(0).toUpperCase() + input.slice(1).toLowerCase() : '';
  }
});

