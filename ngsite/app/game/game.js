'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$http', function($scope, $http) {
  //$http.get('/games/54177f9ef047050f92000004').
  $http.get('/games/54177f9ef047050f9200000').then(
    function(result) {      
      $scope.finance = result.data;
    },
    
    function(error) {
      console.log(error.statusText);
      $scope.error = 'Error in: ' + error.config.url;
      
      if (error.status == 404)	
	$scope.exception = "The requested game data could not be found.";
      else if (error.status == 502)
	$scope.exception = "The game data service appears unavailable. Try again later.";
      else
	$scope.exception = error.statusText;
      
      $scope.showError = true;

      /*
      $(this).removeClass('success');
      $(this).addClass('error');
      $(this).fadeToggle('slow');
      $('#error').text("Error in: '" + settings.url + "'");
      $('#exception').text('Message: ' + exception);	  
      */
    });
}]); 
