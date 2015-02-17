'use strict';

var gameServices = angular.module('bestbuys.services', ['ngResource']);

gameServices.factory('Game', ['$resource',
		     function($resource) {
		       return $resource('/games/:id', {id:'@id'});		       
		    }]);

gameServices.factory('ErrorMessage', [function() {
  return {
    generate: function(error) {
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
    }
  };
}]);
  
gameServices.factory('Command', ['$http', function($http) { 
  return {
    save: function (name, data, message) {	    
	    $http.post('/commands/', {name: name, data: data}).then(      
	      function () {	
		message = {
		  error: false,
		  show: true,
		  text: 'Successfully saved changes.'
		};
	      },
	      function(error) {	
		message = buildError(error);
	      }
	    );	    
	  }
  };
}]);