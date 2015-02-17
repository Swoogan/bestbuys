'use strict';

var services = angular.module('bestbuys.gameServices', ['ngResource']);

services.factory('Game', ['$resource', function($resource) {
    return $resource('/games/:id', {id:'@id'});
}]);

services.service('Notification', [function() {  
  this.showSuccess = function () {	      
    return {
      error: false,
      show: true,
      text: 'Successfully saved changes.'
    };      
  };
  
  this.showError = function(error) {
    var message = {
      error: true,
      show: true,
      text: 'Error in: ' + error.config.url
    };
    
    if (error.status == 404)
      message.details = "The requested game data could not be found.";
    else if (error.status == 502)
      message.details = "The game data service appears unavailable. Try again later.";
    else
      message.details = error.statusText;
    
    return message;
  };
}]);
  
services.service('Command', ['$http', function($http) {   
  this.save = function (name, data) {
    return $http.post('/commands/', {name: name, data: data});
  };  
}]);