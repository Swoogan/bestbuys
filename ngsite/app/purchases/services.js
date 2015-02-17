'use strict';

var services = angular.module('bestbuys.purchasesServices', ['ngResource']);

services.service('Command', ['$http', function($http) {   
  this.save = function (name, data) {
    return $http.post('/commands/', {name: name, data: data});
  };  
}]);