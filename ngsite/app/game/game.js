'use strict';

angular.module('bestbuys.game', ['ngRoute'])

.config(['$routeProvider', function($routeProvider) {
  $routeProvider.when('/game', {
    templateUrl: 'game/game.html',
    controller: 'GameCtrl'
  });
}])

.controller('GameCtrl', ['$scope', '$http', function($scope, $http) {
  //$http.get('/games/54177f9ef047050f9200000').
  $http.get('/games/54177f9ef047050f92000004').then(
    function(result) {      
      $scope.finance = result.data;      
    },
    
    function(error) {
      $scope.message = {};
      
      $scope.message.error = true;
      $scope.message.text = 'Error in: ' + error.config.url;
      
      if (error.status == 404)	
	$scope.message.details = "The requested game data could not be found.";
      else if (error.status == 502)
	$scope.message.details = "The game data service appears unavailable. Try again later.";
      else
	$scope.message.details = error.statusText;
      
      $scope.message.show = true;
    });
}]); 
